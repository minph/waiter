package waiter

// OnKey 等待按键动作
func OnKey() {
	NewKey().Wait()
}

// OnMouse 等待鼠标动作
func OnMouse() {
	NewMouse().Wait()
}

// OnResize 等待缩放动作
func OnResize() {
	NewResize().Wait()
}
