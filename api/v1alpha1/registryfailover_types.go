/*
Copyright 2021.

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
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// RegistryFailoverSpec defines the desired state of RegistryFailover
type RegistryFailoverSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of RegistryFailover. Edit registryfailover_types.go to remove/update
	Foo string `json:"foo,omitempty"`
}

// RegistryFailoverStatus defines the observed state of RegistryFailover
type RegistryFailoverStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// RegistryFailover is the Schema for the registryfailovers API
type RegistryFailover struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   RegistryFailoverSpec   `json:"spec,omitempty"`
	Status RegistryFailoverStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// RegistryFailoverList contains a list of RegistryFailover
type RegistryFailoverList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RegistryFailover `json:"items"`
}

func init() {
	SchemeBuilder.Register(&RegistryFailover{}, &RegistryFailoverList{})
}
