package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// BgSettingStatus defines the observed state of BgSetting
type BgSettingStatus struct {
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// BgSetting is the Schema for the bgsettings API
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=bgsettings,scope=Cluster
type BgSetting struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	Spec   map[string]string `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
	Status BgSettingStatus   `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// BgSettingList contains a list of BgSetting
type BgSettingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items           []BgSetting `json:"items" protobuf:"bytes,2,rep,name=items"`
}
