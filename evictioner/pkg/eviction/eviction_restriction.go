package eviction

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	datahubutils "github.com/containers-ai/alameda/datahub/pkg/utils"
	autoscalingv1alpha1 "github.com/containers-ai/alameda/operator/api/v1alpha1"
	utilsresource "github.com/containers-ai/alameda/operator/pkg/utils/resources"
	datahub_common "github.com/containers-ai/api/alameda_api/v1alpha1/datahub/common"
	datahub_recommendations "github.com/containers-ai/api/alameda_api/v1alpha1/datahub/recommendations"
	datahub_resources "github.com/containers-ai/api/alameda_api/v1alpha1/datahub/resources"

	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
	core_v1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type EvictionRestriction interface {
	IsEvictabel(pod *core_v1.Pod) (isEvictabel bool, err error)
}

type podReplicaStatus struct {
	preservedPodCount float64

	evictedPodCount float64
	runningPodCount float64
}

// NewPodReplicaStatus build podReplicaStatus by pods and maxUnavailable
func NewPodReplicaStatus(pods []core_v1.Pod, replicasCount int32, maxUnavailable string) (podReplicaStatus, error) {

	podReplicaStatus := podReplicaStatus{}

	for _, pod := range pods {
		if pod.Status.Phase == core_v1.PodRunning && pod.ObjectMeta.DeletionTimestamp.IsZero() {
			podReplicaStatus.runningPodCount++
		}
	}
	maxUnavailableCount := float64(0)
	if strings.Contains(maxUnavailable, "%") {
		maxUnavailableValue, err := strconv.ParseFloat(strings.TrimSuffix(maxUnavailable, "%"), 64)
		if err != nil {
			return podReplicaStatus, errors.Errorf("%s", err.Error())
		}
		maxUnavailableCount = math.Ceil(float64(replicasCount) * (maxUnavailableValue / 100))
	} else {
		maxUnavailableValue, err := strconv.ParseFloat(maxUnavailable, 64)
		if err != nil {
			return podReplicaStatus, errors.Errorf("%s", err.Error())
		}
		maxUnavailableCount = math.Ceil(maxUnavailableValue)
	}
	podReplicaStatus.preservedPodCount = float64(replicasCount) - maxUnavailableCount
	if podReplicaStatus.preservedPodCount < 0 {
		podReplicaStatus.preservedPodCount = 0
	}

	return podReplicaStatus, nil
}

type evictionRestriction struct {
	triggerThreshold triggerThreshold

	alamedaScalerMap map[string]*autoscalingv1alpha1.AlamedaScaler

	podIDToPodRecommendationMap            map[string]*datahub_recommendations.PodRecommendation
	podIDToAlamedaResourceIDMap            map[string]string
	alamedaResourceIDToPodReplicaStatusMap map[string]*podReplicaStatus
}

func NewEvictionRestriction(client client.Client, maxUnavailable string, triggerThreshold triggerThreshold, podRecommendations []*datahub_recommendations.PodRecommendation) EvictionRestriction {

	podIDToPodRecommendationMap := make(map[string]*datahub_recommendations.PodRecommendation)
	podIDToAlamedaResourceIDMap := make(map[string]string)
	alamedaResourceIDToPodReplicaStatusMap := make(map[string]*podReplicaStatus)
	for _, podRecommendation := range podRecommendations {

		copyPodRecommendation := proto.Clone(podRecommendation)
		podRecommendation = copyPodRecommendation.(*datahub_recommendations.PodRecommendation)

		if podRecommendation.ObjectMeta == nil {
			scope.Warnf("skip PodRecommendation due to get nil ObjectMeta")
			continue
		}

		podRecommendationNamespace := podRecommendation.ObjectMeta.Namespace
		podRecommendationName := podRecommendation.ObjectMeta.Name
		podNamespace := podRecommendationNamespace
		podName := podRecommendationName
		podID := fmt.Sprintf("%s/%s", podNamespace, podName)
		podIDToPodRecommendationMap[podID] = podRecommendation

		topController := podRecommendation.TopController
		if topController == nil || topController.ObjectMeta == nil {
			scope.Warnf("skip PodRecommendation (%s/%s) due to get empty topController", podRecommendationNamespace, podRecommendationName)
			continue
		}
		alamedaResourceNamespace := topController.ObjectMeta.Namespace
		alamedaResourceName := topController.ObjectMeta.Name
		alamedaResourceKind := topController.Kind
		alamedaResourceID := fmt.Sprintf("%s.%s.%s", alamedaResourceNamespace, alamedaResourceName, alamedaResourceKind)
		podIDToAlamedaResourceIDMap[podID] = alamedaResourceID

		if _, exist := alamedaResourceIDToPodReplicaStatusMap[alamedaResourceID]; !exist {

			getResource := utilsresource.NewGetResource(client)
			listResource := utilsresource.NewListResources(client)

			controllerKind := datahub_resources.Kind_name[int32(alamedaResourceKind)]
			replicasCount, err := getResource.GetReplicasCountByController(alamedaResourceNamespace, alamedaResourceName, strings.ToLower(controllerKind))
			if err != nil {
				if err != nil {
					scope.Warnf("skip PodRecommendation (%s/%s) due to get replicas count by controller failed: %s", podRecommendationNamespace, podRecommendationName, err.Error())
					continue
				}
			}
			pods, err := listResource.ListPodsByController(alamedaResourceNamespace, alamedaResourceName, strings.ToLower(controllerKind))
			if err != nil {
				scope.Warnf("skip PodRecommendation (%s/%s) due to list pods by controller failed: %s", podRecommendationNamespace, podRecommendationName, err.Error())
				continue
			}
			podReplicaStatus, err := NewPodReplicaStatus(pods, replicasCount, maxUnavailable)
			if err != nil {
				scope.Warnf("skip PodRecommendation (%s/%s) due to build PodReplicaStatus failed: %s", podRecommendationNamespace, podRecommendationName, err.Error())
				continue
			}
			alamedaResourceIDToPodReplicaStatusMap[alamedaResourceID] = &podReplicaStatus
		}
	}

	e := &evictionRestriction{
		triggerThreshold: triggerThreshold,

		podIDToPodRecommendationMap:            podIDToPodRecommendationMap,
		podIDToAlamedaResourceIDMap:            podIDToAlamedaResourceIDMap,
		alamedaResourceIDToPodReplicaStatusMap: alamedaResourceIDToPodReplicaStatusMap,
	}
	return e
}

func (e *evictionRestriction) IsEvictabel(pod *core_v1.Pod) (bool, error) {

	podNamespace := pod.Namespace
	podName := pod.Name
	podID := fmt.Sprintf("%s/%s", podNamespace, podName)

	podRecommendation := e.podIDToPodRecommendationMap[podID]
	if !e.isPodEvictable(pod, podRecommendation) {
		return false, nil
	}

	ok, err := e.canRollingUpdatePod(podID)
	if err != nil {
		scope.Errorf("check if rolling update can perform on pod (%s) failed: %s", podID, err.Error())
		return false, err
	} else if !ok {
		return false, nil
	}

	return true, nil
}

func (e *evictionRestriction) canRollingUpdatePod(podID string) (bool, error) {

	alamedaResourceID, exist := e.podIDToAlamedaResourceIDMap[podID]
	if !exist {
		return false, errors.Errorf("topController owning pod does not exist")
	}
	podReplicaStatus, exist := e.alamedaResourceIDToPodReplicaStatusMap[alamedaResourceID]
	if !exist {
		return false, errors.Errorf("PodReplicaStatus of pod does not exit")
	}
	if podReplicaStatus.runningPodCount-podReplicaStatus.evictedPodCount > podReplicaStatus.preservedPodCount {
		podReplicaStatus.evictedPodCount++
		return true, nil
	} else {
		podRecommendationID := podID
		scope.Infof("Pod (%s) is not evictable, current running replicas count %.f is not greater then preseved replicas count %.f , ignore PodRecommendation (%s)",
			podID,
			podReplicaStatus.runningPodCount-podReplicaStatus.evictedPodCount,
			podReplicaStatus.preservedPodCount,
			podRecommendationID)
		return false, nil
	}

}

func (e *evictionRestriction) isPodEvictable(pod *core_v1.Pod, podRecomm *datahub_recommendations.PodRecommendation) bool {
	ctRecomms := podRecomm.GetContainerRecommendations()
	containers := pod.Spec.Containers
	for _, container := range containers {
		for _, recContainer := range ctRecomms {
			if container.Name != recContainer.GetName() {
				continue
			}
			if e.isContainerEvictable(pod, &container, recContainer) {
				return true
			}
		}
	}
	return false
}

func (e *evictionRestriction) isContainerEvictable(pod *core_v1.Pod, container *core_v1.Container, recContainer *datahub_recommendations.ContainerRecommendation) bool {
	cpuTriggerThreshold := e.triggerThreshold.CPU
	memoryTriggerThreshold := e.triggerThreshold.Memory

	if &container.Resources == nil || container.Resources.Limits == nil || container.Resources.Requests == nil {
		scope.Infof("Pod %s/%s selected to evict due to some resource of container %s not defined.",
			pod.GetNamespace(), pod.GetName(), recContainer.GetName())
		return true
	}

	for _, resourceType := range []core_v1.ResourceName{
		core_v1.ResourceMemory,
		core_v1.ResourceCPU,
	} {
		// resource limit check
		if _, ok := container.Resources.Limits[resourceType]; !ok {
			scope.Infof("Pod %s/%s selected to evict due to resource limit %s of container %s not defined.",
				pod.GetNamespace(), pod.GetName(), resourceType, recContainer.GetName())
			return true
		}

		for _, limitRec := range recContainer.GetLimitRecommendations() {
			if resourceType == core_v1.ResourceMemory && limitRec.GetMetricType() == datahub_common.MetricType_MEMORY_USAGE_BYTES && len(limitRec.GetData()) > 0 {
				if limitRecVal, err := datahubutils.StringToFloat64(limitRec.GetData()[0].GetNumValue()); err == nil {
					limitRecVal = math.Ceil(limitRecVal)
					limitQuan := container.Resources.Limits[resourceType]
					delta := (math.Abs(float64(100*(limitRecVal-float64(limitQuan.Value())))) / float64(limitQuan.Value()))
					scope.Infof("Resource limit of %s pod %s/%s container %s checking eviction threshold (%v perentage). Current setting: %v, Recommended setting: %v",
						resourceType, pod.GetNamespace(), pod.GetName(), recContainer.GetName(), memoryTriggerThreshold, limitQuan.Value(), limitRecVal)
					if delta >= memoryTriggerThreshold {
						scope.Infof("Decide to evict pod %s/%s due to delta is %v >= %v (threshold)", pod.GetNamespace(), pod.GetName(), delta, memoryTriggerThreshold)
						return true
					}
				}
			}
			if resourceType == core_v1.ResourceCPU && limitRec.GetMetricType() == datahub_common.MetricType_CPU_USAGE_SECONDS_PERCENTAGE && len(limitRec.GetData()) > 0 {
				if limitRecVal, err := datahubutils.StringToFloat64(limitRec.GetData()[0].GetNumValue()); err == nil {
					limitRecVal = math.Ceil(limitRecVal)
					limitQuan := container.Resources.Limits[resourceType]
					delta := (math.Abs(float64(100*(limitRecVal-float64(limitQuan.MilliValue())))) / float64(limitQuan.MilliValue()))
					scope.Infof("Resource limit of %s pod %s/%s container %s checking eviction threshold (%v perentage). Current setting: %v, Recommended setting: %v",
						resourceType, pod.GetNamespace(), pod.GetName(), recContainer.GetName(), cpuTriggerThreshold, limitQuan.MilliValue(), limitRecVal)
					if delta >= cpuTriggerThreshold {
						scope.Infof("Decide to evict pod %s/%s due to delta is %v >= %v (threshold)", pod.GetNamespace(), pod.GetName(), delta, cpuTriggerThreshold)
						return true
					}
				}
			}
		}

		// resource request check
		if _, ok := container.Resources.Requests[resourceType]; !ok {
			scope.Infof("Pod %s/%s selected to evict due to resource request %s of container %s not defined.",
				pod.GetNamespace(), pod.GetName(), resourceType, recContainer.GetName())
			return true
		}
		for _, reqRec := range recContainer.GetRequestRecommendations() {
			if resourceType == core_v1.ResourceMemory && reqRec.GetMetricType() == datahub_common.MetricType_MEMORY_USAGE_BYTES && len(reqRec.GetData()) > 0 {
				if requestRecVal, err := datahubutils.StringToFloat64(reqRec.GetData()[0].GetNumValue()); err == nil {
					requestRecVal = math.Ceil(requestRecVal)
					requestQuan := container.Resources.Requests[resourceType]
					delta := (math.Abs(float64(100*(requestRecVal-float64(requestQuan.Value())))) / float64(requestQuan.Value()))
					scope.Infof("Resource request of %s pod %s/%s container %s checking eviction threshold (%v perentage). Current setting: %v, Recommended setting: %v",
						resourceType, pod.GetNamespace(), pod.GetName(), recContainer.GetName(), memoryTriggerThreshold, requestQuan.Value(), requestRecVal)
					if delta >= memoryTriggerThreshold {
						scope.Infof("Decide to evict pod %s/%s due to delta is %v >= %v (threshold)", pod.GetNamespace(), pod.GetName(), delta, memoryTriggerThreshold)
						return true
					}
				}
			}
			if resourceType == core_v1.ResourceCPU && reqRec.GetMetricType() == datahub_common.MetricType_CPU_USAGE_SECONDS_PERCENTAGE && len(reqRec.GetData()) > 0 {
				if requestRecVal, err := datahubutils.StringToFloat64(reqRec.GetData()[0].GetNumValue()); err == nil {
					requestRecVal = math.Ceil(requestRecVal)
					requestQuan := container.Resources.Requests[resourceType]
					delta := (math.Abs(float64(100*(requestRecVal-float64(requestQuan.MilliValue())))) / float64(requestQuan.MilliValue()))
					scope.Infof("Resource request of %s pod %s/%s container %s checking eviction threshold (%v perentage). Current setting: %v, Recommended setting: %v",
						resourceType, pod.GetNamespace(), pod.GetName(), recContainer.GetName(), cpuTriggerThreshold, requestQuan.MilliValue(), requestRecVal)
					if delta >= cpuTriggerThreshold {
						scope.Infof("Decide to evict pod %s/%s due to delta is %v >= %v (threshold)", pod.GetNamespace(), pod.GetName(), delta, cpuTriggerThreshold)
						return true
					}
				}
			}
		}
	}
	return false
}
