package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type HyperNode struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec HyperNodeSpec `json:"spec,omitempty"`
}

type HyperNodeSpec struct {
	SupervisorName string `json:"supervisorName,omitempty"`

	MyAreaName string `json:"myAreaName,omitempty"`

	NodeAreaType string `json:"nodeAreaType,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type HyperNodeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []HyperNode `json:"items"`
}
