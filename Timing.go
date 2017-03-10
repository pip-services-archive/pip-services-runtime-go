package runtime

import "time"

type ITimingSetter interface {
	SetTiming(name string, elapsed float32)
}

type Timing struct {
	start  time.Time
	setter ITimingSetter
	name   string
}

func EmptyTiming() *Timing {
	return &Timing{setter: nil, name: ""}
}

func NewTiming(setter ITimingSetter, name string) *Timing {
	return &Timing{
		start:  time.Now(),
		setter: setter,
		name:   name}
}

func (t *Timing) EndTiming() {
	if t.setter != nil {
		elapsed := (float32)(time.Since(t.start) / time.Millisecond)
		t.setter.SetTiming(t.name, elapsed)
	}
}
