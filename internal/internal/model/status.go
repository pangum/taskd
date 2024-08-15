package model

const (
	// StatusCreated 已创建
	StatusCreated = 1
	// StatusRunning 执行中
	StatusRunning = 2
	// StatusRetrying 重试中
	StatusRetrying = 3

	// StatusFailed 失败
	StatusFailed = 10

	// StatusSuccess 成功
	StatusSuccess = 20
)

// Status 类型
type Status uint8
