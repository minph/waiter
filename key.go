package mode

import (
	"fmt"
	"github.com/nsf/termbox-go"
)

// KeyTrigger 按键触发动作
type KeyTrigger struct {
	tip    string
	onexit func()
}


// Wait 等待按键触发
func (k *KeyTrigger) Wait() {
	if len(k.tip) > 0 {
		fmt.Println(k.tip)
	}
	Loop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			if k.onexit != nil {
				k.onexit()
			}
			break Loop
		}
	}
}

// Tip 设置 Tips
func (k *KeyTrigger) Tip(tip string) Waiter {
	k.tip = tip
	return k
}


// OnExit 设置离开动作
func (k *KeyTrigger) OnExit(onexit func()) Waiter {
	k.onexit = onexit
	return k
}
