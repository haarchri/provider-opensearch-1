/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type ComponentTemplateObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type ComponentTemplateParameters struct {

	// The JSON body of the template.
	// +kubebuilder:validation:Required
	Body *string `json:"body" tf:"body,omitempty"`

	// Name of the component template to create.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`
}

// ComponentTemplateSpec defines the desired state of ComponentTemplate
type ComponentTemplateSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ComponentTemplateParameters `json:"forProvider"`
}

// ComponentTemplateStatus defines the observed state of ComponentTemplate.
type ComponentTemplateStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ComponentTemplateObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ComponentTemplate is the Schema for the ComponentTemplates API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,opensearch}
type ComponentTemplate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ComponentTemplateSpec   `json:"spec"`
	Status            ComponentTemplateStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ComponentTemplateList contains a list of ComponentTemplates
type ComponentTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ComponentTemplate `json:"items"`
}

// Repository type metadata.
var (
	ComponentTemplate_Kind             = "ComponentTemplate"
	ComponentTemplate_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ComponentTemplate_Kind}.String()
	ComponentTemplate_KindAPIVersion   = ComponentTemplate_Kind + "." + CRDGroupVersion.String()
	ComponentTemplate_GroupVersionKind = CRDGroupVersion.WithKind(ComponentTemplate_Kind)
)

func init() {
	SchemeBuilder.Register(&ComponentTemplate{}, &ComponentTemplateList{})
}