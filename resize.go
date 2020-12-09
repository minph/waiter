package waiter

import (
	"fmt"

	"github.com/nsf/termbox-go"
)

// ResizeTrigger 缩放触发动作
type ResizeTrigger struct {
	tip    string
	onexit func()
}

// Wait 等待缩放触发
func (r *ResizeTrigger) Wait() {
	if len(r.tip) > 0 {
		fmt.Println(r.tip)
	}
Loop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventResize:
			if r.onexit != nil {
				r.onexit()
			}
			break Loop
		}
	}
}

// Tip 设置 Tips
func (r *ResizeTrigger) Tip(tip string) Waiter {
	r.tip = tip
	return r
}

// OnExit 设置离开动作
func (r *ResizeTrigger) OnExit(onexit func()) Waiter {
	r.onexit = onexit
	return r
}
