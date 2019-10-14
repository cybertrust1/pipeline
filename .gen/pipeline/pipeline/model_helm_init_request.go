/*
 * Pipeline API
 *
 * Pipeline is a feature rich application platform, built for containers on top of Kubernetes to automate the DevOps experience, continuous application development and the lifecycle of deployments. 
 *
 * API version: latest
 * Contact: info@banzaicloud.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package pipeline

type HelmInitRequest struct {

	KubeContext string `json:"kube_context,omitempty"`

	Namespace string `json:"namespace"`

	Upgrade bool `json:"upgrade,omitempty"`

	ForceUpgrade bool `json:"force_upgrade,omitempty"`

	ServiceAccount string `json:"service_account"`

	CanaryImage bool `json:"canary_image,omitempty"`

	TillerImage string `json:"tiller_image,omitempty"`

	HistoryMax int32 `json:"history_max,omitempty"`
}