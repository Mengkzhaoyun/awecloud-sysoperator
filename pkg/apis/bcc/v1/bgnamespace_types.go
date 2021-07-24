package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// BgNamespaceSpec defines the desired state of BgNamespace
type BgNamespaceSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
	Namespace  string                `json:"namespace,omitempty" protobuf:"bytes,1,opt,name=namespace"`
	Alias      string                `json:"alias,omitempty" protobuf:"bytes,2,opt,name=alias"`
	Des        string                `json:"des,omitempty" protobuf:"bytes,3,opt,name=des"`
	Limit      BgNamespaceLimit      `json:"limit,omitempty" protobuf:"bytes,4,opt,name=limit"`
	Pod        int64                 `json:"pod,omitempty" protobuf:"varint,5,opt,name=pod"`
	LimitRange BgNamespaceLimitRange `json:"limit_range" protobuf:"bytes,6,opt,name=limit_range,json=limitRange"`
	Group      string                `json:"group,omitempty" protobuf:"bytes,7,opt,name=group"`
	Class      string                `json:"class,omitempty" protobuf:"bytes,8,opt,name=class"`
	Approval   string                `json:"approval,omitempty" protobuf:"bytes,9,opt,name=approval"` // 审批状态
	// 创建人
	// +optional
	Creator string `json:"creator,omitempty" protobuf:"bytes,10,opt,name=creator"`
}

// BgNamespaceLimitRange
type BgNamespaceLimitRange struct {
	LimitCpu      string `json:"limit_cpu" protobuf:"bytes,1,opt,name=limit_cpu,json=limitCpu"`
	LimitMemory   string `json:"limit_memory" protobuf:"bytes,2,opt,name=limit_memory,json=limitMemory"`
	RequestCpu    string `json:"request_cpu" protobuf:"bytes,3,opt,name=request_cpu,json=requestCpu"`
	RequestMemory string `json:"request_memory" protobuf:"bytes,4,opt,name=request_memory,json=requestMemory"`
}

// BgNamespaceLimit
type BgNamespaceLimit struct {
	Cpu    string `json:"cpu" protobuf:"bytes,1,opt,name=cpu"`
	Memory string `json:"memory" protobuf:"bytes,2,opt,name=memory"`
}

// BgNamespaceStatus defines the observed state of BgNamespace
type BgNamespaceStatus struct {
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

// BgNamespace is the Schema for the bgnamespaces API
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=bgnamespaces,scope=Cluster
type BgNamespace struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	Spec   BgNamespaceSpec   `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
	Status BgNamespaceStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// BgNamespaceList contains a list of BgNamespace
type BgNamespaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items           []BgNamespace `json:"items" protobuf:"bytes,2,rep,name=items"`
}
