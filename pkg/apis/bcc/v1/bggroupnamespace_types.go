package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// BgGroupNamespace , 分组与命名空间关联对象
// NOTE: 会创建User对象

// BgGroupNamespaceSpec defines the desired state of BgGroupNamespace
// +k8s:openapi-gen=true
type BgGroupNamespaceSpec struct {
	Group     string `json:"group" protobuf:"bytes,1,opt,name=group"`
	Namespace string `json:"namespace" protobuf:"bytes,2,opt,name=namespace"`
}

// BgGroupNamespaceStatus defines the observed state of BgGroupNamespace
// +k8s:openapi-gen=true
type BgGroupNamespaceStatus struct {
	// 创建人
	// +optional
	Creator string `json:"creator,omitempty" protobuf:"bytes,1,opt,name=creator"`

	// 创建时间
	// +optional
	Created metav1.Time `json:"created,omitempty" protobuf:"bytes,2,opt,name=created"`

	// 运行状态， Running , Completed
	// +optional
	Phase BgPhase `json:"phase,omitempty" protobuf:"bytes,3,opt,name=phase,casttype=BgPhase"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// BgGroupNamespace is the Schema for the bggroupnamespaces API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=bggroupnamespaces,scope=Cluster
type BgGroupNamespace struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	Spec   BgGroupNamespaceSpec   `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
	Status BgGroupNamespaceStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// BgGroupNamespaceList contains a list of BgGroupNamespace
type BgGroupNamespaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items           []BgGroupNamespace `json:"items" protobuf:"bytes,2,rep,name=items"`
}
