/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
    corev1 "k8s.io/api/core/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Foo is a specification for a Foo resource
type CustomDeployment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CustomDeploymentSpec   `json:"spec"`
	Status CustomDeploymentStatus `json:"status"`
}

// CustomDeploymentSpec is the spec for a CustomDeployment resource
type CustomDeploymentSpec struct {
    PodName string `json:"podName"`
    Labels map[string]string `json:"labels"`
    Containers []CustomDeploymentContainers `json:"containers"`
//	DeploymentName string `json:"deploymentName"`
//	Freeze         bool `json:"freeze"`
}

type CustomDeploymentContainers struct {
    Name string `json:"name"`
    Image string `json:"image"`
    Ports []corev1.ContainerPort `json:"ports"`
}

// CustomDeploymentStatus is the status for a CustomDeployment resource
type CustomDeploymentStatus struct {
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// CustomDeploymentList is a list of CustomDeployment resources
type CustomDeploymentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []CustomDeployment `json:"items"`
}
