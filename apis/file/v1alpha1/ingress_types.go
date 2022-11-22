/*
Copyright 2020 The Crossplane Authors.

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
	"reflect"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

// IngressParameters are the configurable fields of a Ingress.
type IngressParameters struct {
	FileName string `json:"fileName"`
	
}

// IngressObservation are the observable fields of a Ingress.
type IngressObservation struct {
	ObservableField string `json:"observableField,omitempty"`
}

// A IngressSpec defines the desired state of a Ingress.
type IngressSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       IngressParameters `json:"forProvider"`
}

// A IngressStatus represents the observed state of a Ingress.
type IngressStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          IngressObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// A Ingress is an example API type.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,git}
type Ingress struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   IngressSpec   `json:"spec"`
	Status IngressStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IngressList contains a list of Ingress
type IngressList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Ingress `json:"items"`
}

// Ingress type metadata.
var (
	IngressKind             = reflect.TypeOf(Ingress{}).Name()
	IngressGroupKind        = schema.GroupKind{Group: Group, Kind: IngressKind}.String()
	IngressKindAPIVersion   = IngressKind + "." + SchemeGroupVersion.String()
	IngressGroupVersionKind = SchemeGroupVersion.WithKind(IngressKind)
)

func init() {
	SchemeBuilder.Register(&Ingress{}, &IngressList{})
}
