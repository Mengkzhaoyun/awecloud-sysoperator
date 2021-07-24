package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// BgStorageSpec defines the desired state of BgStorage
type BgStorageSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
	StorageName  string   `json:"storageName" protobuf:"bytes,1,opt,name=storageName"`
	Volume       string   `json:"volume,omitempty" protobuf:"bytes,2,opt,name=volume"`
	StorageClass string   `json:"storageClass" protobuf:"bytes,3,opt,name=storageClass"`
	OldStorage   string   `json:"oldStorage,omitempty" protobuf:"bytes,4,opt,name=oldStorage"`
	NewStorage   string   `json:"newStorage" protobuf:"bytes,5,opt,name=newStorage"`
	Class        string   `json:"class" protobuf:"bytes,6,opt,name=class"`
	AccessModes  []string `json:"accessModes" protobuf:"bytes,7,rep,name=accessModes"`
	Approval     string   `json:"approval,omitempty" protobuf:"bytes,8,opt,name=approval"`
	// 创建人
	// +optional
	Creator string `json:"creator,omitempty" protobuf:"bytes,9,opt,name=creator"`
}

// BgStorageStatus defines the observed state of BgStorage
type BgStorageStatus struct {
	// 创建人
	// +optional
	Creator string `json:"creator,omitempty" protobuf:"bytes,1,opt,name=creator"`

	// 创建时间
	// +optional
	Created metav1.Time `json:"created,omitempty" protobuf:"bytes,2,opt,name=created"`

	// 运行状态， Running , Completed
	// +optional
	Phase BgPhase `json:"phase,omitempty" protobuf:"bytes,3,opt,name=phase"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// BgStorage is the Schema for the bgstorages API
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=bgstorages,scope=Namespaced
type BgStorage struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	Spec   BgStorageSpec   `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
	Status BgStorageStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// BgStorageList contains a list of BgStorage
type BgStorageList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items           []BgStorage `json:"items" protobuf:"bytes,2,rep,name=items"`
}
