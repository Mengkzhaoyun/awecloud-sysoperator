package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// BgLoginLogSpec defines the desired state of BgLoginLog
type BgLoginLogSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
	Logs []Log `json:"logs" protobuf:"bytes,1,rep,name=logs"`
}

type Log struct {
	IP           string      `json:"ip" protobuf:"bytes,1,opt,name=ip"`
	Username     string      `json:"username" protobuf:"bytes,2,opt,name=username"`
	OperatorType int32       `json:"operatorType" protobuf:"varint,3,opt,name=operatorType"` // 操作类型  1 登录 | 2 登出  | 3 新增 | 4 修改 | 5 删除
	Content      string      `json:"content" protobuf:"bytes,4,opt,name=content"`            // 操作内容
	Result       int32       `json:"result" protobuf:"varint,5,opt,name=result"`             // 操作结果  0 失败 | 1 成功
	Time         metav1.Time `json:"created,omitempty" protobuf:"bytes,6,opt,name=created"`  // 操作时间
}

// BgLoginLogStatus defines the observed state of BgLoginLog
type BgLoginLogStatus struct {
	// 创建时间
	// +optional
	Created metav1.Time `json:"created,omitempty" protobuf:"bytes,1,opt,name=created"`

	// 运行状态， Running , Completed
	// +optional
	Phase BgPhase `json:"phase,omitempty" protobuf:"bytes,2,opt,name=phase,casttype=BgPhase"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// BgLoginLog is the Schema for the bgloginlogs API
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=bgloginlogs,scope=Cluster
type BgLoginLog struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	Spec   BgLoginLogSpec   `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
	Status BgLoginLogStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// BgLoginLogList contains a list of BgLoginLog
type BgLoginLogList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items           []BgLoginLog `json:"items" protobuf:"bytes,2,rep,name=items"`
}
