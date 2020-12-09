package mode

import (
	"fmt"
	"github.com/nsf/termbox-go"
)

// MouseTrigger 鼠标触发动作
type MouseTrigger struct {
	tip    string
	onexit func()
}

// Wait 等待鼠标触发
func (m *MouseTrigger) Wait() {
	if len(m.tip) > 0 {
		fmt.Println(m.tip)
	}
	Loop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventMouse:
			if m.onexit != nil {
				m.onexit()
			}
			break Loop
		}
	}
}

// Tip 设置 Tips
func (m *MouseTrigger) Tip(tip string) Waiter {
	m.tip = tip
	return m
}


// OnExit 设置离开动作
func (m *MouseTrigger) OnExit(onexit func()) Waiter {
	m.onexit = onexit
	return m
}
