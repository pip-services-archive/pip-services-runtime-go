package runtime

import "time"

type Counter struct {
    Name string    `json:"name"`
    Type int       `json:"type"`
    Last *float32   `json:"last"`
    Count *int      `json:"count"`
    Min *float32    `json:"min"`
    Max *float32    `json:"max"`
    Avg *float32    `json:"avg"`
    Time *time.Time `json:"time"`
}

func NewCounter(name string, typ int) *Counter {
    return &Counter{ Name: name, Type: typ }
}
