package resources

import (
	"context"
	"fmt"
	"strings"

	autuscaling "github.com/containers-ai/alameda/operator/pkg/apis/autoscaling/v1alpha1"
	logUtil "github.com/containers-ai/alameda/operator/pkg/utils/log"
	appsapi_v1 "github.com/openshift/api/apps/v1"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	listResourcesScope = logUtil.RegisterScope("listresources", "List resources", 0)
)

// ListResources define resource list functions
type ListResources struct {
	client client.Client
}

// NewListResources return ListResources instance
func NewListResources(client client.Client) *ListResources {
	return &ListResources{
		client: client,
	}
}

// ListAllNodes return all nodes in cluster
func (listResources *ListResources) ListAllNodes() ([]corev1.Node, error) {
	nodeList := &corev1.NodeList{}
	if err := listResources.listAllResources(nodeList); err != nil {
		return []corev1.Node{}, err
	}
	return nodeList.Items, nil
}

// ListPodsByLabels return pods by labels
func (listResources *ListResources) ListPodsByLabels(labels map[string]string) ([]corev1.Pod, error) {
	podList := &corev1.PodList{}
	if err := listResources.listResourcesByLabels(podList, labels); err != nil {
		return []corev1.Pod{}, err
	}

	return podList.Items, nil
}

// ListDeploymentsByLabels return deployments by labels
func (listResources *ListResources) ListDeploymentsByLabels(labels map[string]string) ([]appsv1.Deployment, error) {
	deploymentList := &appsv1.DeploymentList{}
	if err := listResources.listResourcesByLabels(deploymentList, labels); err != nil {
		return []appsv1.Deployment{}, err
	}

	return deploymentList.Items, nil
}

// ListDeploymentsByNamespaceLabels return deployments by namespace and labels
func (listResources *ListResources) ListDeploymentsByNamespaceLabels(namespace string, labels map[string]string) ([]appsv1.Deployment, error) {
	deploymentList := &appsv1.DeploymentList{}

	if err := listResources.listResourcesByNamespaceLabels(deploymentList, namespace, labels); err != nil {
		return []appsv1.Deployment{}, err
	}

	return deploymentList.Items, nil
}

// ListDeploymentConfigsByNamespaceLabels return deploymentconfigs by namespace and labels
func (listResources *ListResources) ListDeploymentConfigsByNamespaceLabels(namespace string, labels map[string]string) ([]appsapi_v1.DeploymentConfig, error) {
	deploymentConfigList := &appsapi_v1.DeploymentConfigList{}

	if err := listResources.listResourcesByNamespaceLabels(deploymentConfigList, namespace, labels); err != nil {
		return []appsapi_v1.DeploymentConfig{}, err
	}

	return deploymentConfigList.Items, nil
}

// ListDeploymentConfigsByLabels return DeploymentConfigs by labels
func (listResources *ListResources) ListDeploymentConfigsByLabels(labels map[string]string) ([]appsapi_v1.DeploymentConfig, error) {
	deploymentConfigList := &appsapi_v1.DeploymentConfigList{}
	if err := listResources.listResourcesByLabels(deploymentConfigList, labels); err != nil {
		return []appsapi_v1.DeploymentConfig{}, err
	}

	return deploymentConfigList.Items, nil
}

// ListPodsByDeployment return pods by deployment namespace and name
func (listResources *ListResources) ListPodsByDeployment(deployNS, deployName string) ([]corev1.Pod, error) {
	deploymentPods := []corev1.Pod{}
	podList := &corev1.PodList{}
	if err := listResources.listResourcesByNamespace(podList, deployNS); err != nil {
		return deploymentPods, err
	}

	for _, pod := range podList.Items {
		podName := pod.GetName()
		for _, or := range pod.GetOwnerReferences() {
			if *or.Controller && strings.ToLower(or.Kind) == "replicaset" && strings.HasPrefix(podName, fmt.Sprintf("%s-", deployName)) && strings.HasPrefix(podName, fmt.Sprintf("%s-", or.Name)) {
				deploymentPods = append(deploymentPods, pod)
				break
			}
		}
	}

	return deploymentPods, nil
}

// ListPodsByDeploymentConfig return pods by deployment namespace and name
func (listResources *ListResources) ListPodsByDeploymentConfig(deployNS, deployConfigName string) ([]corev1.Pod, error) {
	deploymentConfigPods := []corev1.Pod{}
	podList := &corev1.PodList{}
	if err := listResources.listResourcesByNamespace(podList, deployNS); err != nil {
		return deploymentConfigPods, err
	}

	for _, pod := range podList.Items {
		podName := pod.GetName()
		for _, or := range pod.GetOwnerReferences() {
			if *or.Controller && strings.ToLower(or.Kind) == "replicationcontroller" && strings.HasPrefix(podName, fmt.Sprintf("%s-", deployConfigName)) && strings.HasPrefix(podName, fmt.Sprintf("%s-", or.Name)) {
				deploymentConfigPods = append(deploymentConfigPods, pod)
				break
			}
		}
	}

	return deploymentConfigPods, nil
}

// ListAllAlamedaScaler return all nodes in cluster
func (listResources *ListResources) ListAllAlamedaScaler() ([]autuscaling.AlamedaScaler, error) {
	alamedaScalerList := &autuscaling.AlamedaScalerList{}
	if err := listResources.listAllResources(alamedaScalerList); err != nil {
		return []autuscaling.AlamedaScaler{}, err
	}
	return alamedaScalerList.Items, nil
}

func (listResources *ListResources) listAllResources(resourceList runtime.Object) error {
	if err := listResources.client.List(context.TODO(),
		&client.ListOptions{},
		resourceList); err != nil {
		listResourcesScope.Error(err.Error())
		return err
	}
	return nil
}

func (listResources *ListResources) listResourcesByNamespace(resourceList runtime.Object, namespace string) error {
	if err := listResources.client.List(context.TODO(),
		&client.ListOptions{
			Namespace: namespace,
		}, resourceList); err != nil {
		listResourcesScope.Error(err.Error())
		return err
	}
	return nil
}

func (listResources *ListResources) listResourcesByLabels(resourceList runtime.Object, lbls map[string]string) error {
	if err := listResources.client.List(context.TODO(),
		client.MatchingLabels(lbls),
		resourceList); err != nil {
		listResourcesScope.Error(err.Error())
		return err
	}
	return nil
}

func (listResources *ListResources) listResourcesByNamespaceLabels(resourceList runtime.Object, namespace string, lbls map[string]string) error {
	if err := listResources.client.List(context.TODO(),
		client.InNamespace(namespace).MatchingLabels(lbls),
		resourceList); err != nil {
		listResourcesScope.Debug(err.Error())
		return err
	}
	return nil
}