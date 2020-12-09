# mode

一个用于暂停控制台的 golang 工具库

`import "github.com/minph/waiter"`

- [Overview](#pkg-overview)
- [Index](#pkg-index)

## <a name="pkg-overview">Overview</a>

## <a name="pkg-index">Index</a>

- [type KeyTrigger](#KeyTrigger)
  - [func (k \*KeyTrigger) OnExit(onexit func()) Waiter](#KeyTrigger.OnExit)
  - [func (k \*KeyTrigger) Tip(tip string) Waiter](#KeyTrigger.Tip)
  - [func (k \*KeyTrigger) Wait()](#KeyTrigger.Wait)
- [type MouseTrigger](#MouseTrigger)
  - [func (m \*MouseTrigger) OnExit(onexit func()) Waiter](#MouseTrigger.OnExit)
  - [func (m \*MouseTrigger) Tip(tip string) Waiter](#MouseTrigger.Tip)
  - [func (m \*MouseTrigger) Wait()](#MouseTrigger.Wait)
- [type ResizeTrigger](#ResizeTrigger)
  - [func (r \*ResizeTrigger) OnExit(onexit func()) Waiter](#ResizeTrigger.OnExit)
  - [func (r \*ResizeTrigger) Tip(tip string) Waiter](#ResizeTrigger.Tip)
  - [func (r \*ResizeTrigger) Wait()](#ResizeTrigger.Wait)
- [type Trigger](#Trigger)
- [type Waiter](#Waiter)
  - [func New(trigger Trigger) Waiter](#New)
  - [func NewKey() Waiter](#NewKey)
  - [func NewMouse() Waiter](#NewMouse)
  - [func NewResize() Waiter](#NewResize)

## <a name="KeyTrigger">type</a> KeyTrigger

```go
type KeyTrigger struct {
    // contains filtered or unexported fields
}

```

KeyTrigger 按键触发动作

### <a name="KeyTrigger.OnExit">func</a> (\*KeyTrigger) OnExit

```go
func (k *KeyTrigger) OnExit(onexit func()) Waiter
```

OnExit 设置离开动作

### <a name="KeyTrigger.Tip">func</a> (\*KeyTrigger) Tip

```go
func (k *KeyTrigger) Tip(tip string) Waiter
```

Tip 设置 Tips

### <a name="KeyTrigger.Wait">func</a> (\*KeyTrigger) Wait

```go
func (k *KeyTrigger) Wait()
```

Wait 等待按键触发

## <a name="MouseTrigger">type</a> MouseTrigger

```go
type MouseTrigger struct {
    // contains filtered or unexported fields
}

```

MouseTrigger 鼠标触发动作

### <a name="MouseTrigger.OnExit">func</a> (\*MouseTrigger) OnExit

```go
func (m *MouseTrigger) OnExit(onexit func()) Waiter
```

OnExit 设置离开动作

### <a name="MouseTrigger.Tip">func</a> (\*MouseTrigger) Tip

```go
func (m *MouseTrigger) Tip(tip string) Waiter
```

Tip 设置 Tips

### <a name="MouseTrigger.Wait">func</a> (\*MouseTrigger) Wait

```go
func (m *MouseTrigger) Wait()
```

Wait 等待鼠标触发

## <a name="ResizeTrigger">type</a> ResizeTrigger

```go
type ResizeTrigger struct {
    // contains filtered or unexported fields
}

```

ResizeTrigger 缩放触发动作

### <a name="ResizeTrigger.OnExit">func</a> (\*ResizeTrigger) OnExit

```go
func (r *ResizeTrigger) OnExit(onexit func()) Waiter
```

OnExit 设置离开动作

### <a name="ResizeTrigger.Tip">func</a> (\*ResizeTrigger) Tip

```go
func (r *ResizeTrigger) Tip(tip string) Waiter
```

Tip 设置 Tips

### <a name="ResizeTrigger.Wait">func</a> (\*ResizeTrigger) Wait

```go
func (r *ResizeTrigger) Wait()
```

Wait 等待缩放触发

## <a name="Trigger">type</a> Trigger

```go
type Trigger int
```

Trigger 触发类型

```go
const (
    // Key 按键类型
    Key Trigger = iota
    // Mouse 鼠标类型
    Mouse
    // Resize 缩放类型
    Resize
)
```

## <a name="Waiter">type</a> Waiter

```go
type Waiter interface {
    Wait()
    Tip(tip string) Waiter
    OnExit(onexit func()) Waiter
}
```

Waiter 等待触发

### <a name="New">func</a> New

```go
func New(trigger Trigger) Waiter
```

New 创建 Waiter

### <a name="NewKey">func</a> NewKey

```go
func NewKey() Waiter
```

NewKey 创建按键 Waiter

### <a name="NewMouse">func</a> NewMouse

```go
func NewMouse() Waiter
```

NewMouse 创建鼠标 Waiter

### <a name="NewResize">func</a> NewResize

```go
func NewResize() Waiter
```

NewResize 创建缩放 Waiter
