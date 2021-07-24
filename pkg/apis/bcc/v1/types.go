package v1

// BgPhase , 表示运行时状态
type BgPhase string

// These are the valid statuses of BgInstance.
const (
	// BgRunning 表示示例处于运行状态
	BgRunning BgPhase = "Running"
	// BgCompleted 表示实例运行结束
	BgCompleted BgPhase = "Completed"
)
