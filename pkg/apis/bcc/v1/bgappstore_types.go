package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// BgAppStoreSpec defines the desired state of BgAppStore
type BgAppStoreSpec struct {
	Name      string `json:"name" protobuf:"bytes,1,opt,name=name"`           // 中文名称
	Logo      string `json:"logo" protobuf:"bytes,2,opt,name=logo"`           // 图标
	Version   string `json:"version" protobuf:"bytes,3,opt,name=version"`     // 版本号
	Step      string `json:"step" protobuf:"bytes,4,opt,name=step"`           // step 文件
	Values    string `json:"values" protobuf:"bytes,5,opt,name=values"`       // values 文件
	ChartName string `json:"chartName" protobuf:"bytes,6,opt,name=chartName"` // chart 名称
	Readme    string `json:"readme" protobuf:"bytes,7,opt,name=readme"`       // readme 文件
	Chart     string `json:"chart" protobuf:"bytes,8,opt,name=chart"`         // chart 内容
	State     int32  `json:"state" protobuf:"varint,9,opt,name=state"`        // 状态 1 上线  0下线
}

// BgAppStoreStatus defines the observed state of BgAppStore
type BgAppStoreStatus struct {
	// +optional
	Created metav1.Time `json:"created,omitempty" protobuf:"bytes,1,opt,name=created"`

	// 运行状态， Running , Completed
	// +optional
	Phase BgPhase `json:"phase,omitempty" protobuf:"bytes,2,opt,name=phase,casttype=BgPhase"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// BgAppStore is the Schema for the bgappstores API
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=bgappstores,scope=Cluster
type BgAppStore struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	Spec   BgAppStoreSpec   `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
	Status BgAppStoreStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// BgAppStoreList contains a list of BgAppStore
type BgAppStoreList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items           []BgAppStore `json:"items" protobuf:"bytes,2,rep,name=items"`
}
