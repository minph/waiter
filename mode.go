package waiter

// Trigger 触发类型
type Trigger int

const (
	// Key 按键类型
	Key Trigger = iota
	// Mouse 鼠标类型
	Mouse
	// Resize 缩放类型
	Resize
)

// Waiter 等待触发
type Waiter interface {
	Wait()
	Tip(tip string) Waiter
	OnExit(onexit func()) Waiter
}

// New 创建 Waiter
func New(trigger Trigger) Waiter {
	switch trigger {
	case Key:
		return &KeyTrigger{}
	case Mouse:
		return &MouseTrigger{}
	case Resize:
		return &ResizeTrigger{}
	}

	return &KeyTrigger{}
}

// NewKey 创建按键 Waiter
func NewKey() Waiter {
	return &KeyTrigger{}
}

// NewMouse 创建鼠标 Waiter
func NewMouse() Waiter {
	return &MouseTrigger{}
}

// NewResize 创建缩放 Waiter
func NewResize() Waiter {
	return &ResizeTrigger{}
}
