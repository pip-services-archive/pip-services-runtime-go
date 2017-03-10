package log

import (
    "fmt"
    "github.com/pip-services/pip-services-runtime-go"
)

type AbstractLog struct {
    runtime.AbstractComponent
    
    level int
}

func NewAbstractLog(id string, config *runtime.DynamicMap) *AbstractLog {
    defaultConfig := runtime.DynamicMap { "level": runtime.LogLevelInfo }
    config = runtime.NewMapWithDefaults(config, &defaultConfig)

    c := AbstractLog { AbstractComponent: *runtime.NewAbstractComponent(id, config) }
    c.level = int(config.GetInteger("level"))
    
    return &c
}

func (c *AbstractLog) Level() int {
    return c.level
}

func (c *AbstractLog) SetLevel(value int) {
    c.level = value
}

func (c *AbstractLog) Fatal(message ...interface{}) {}
func (c *AbstractLog) Error(message ...interface{}) {}
func (c *AbstractLog) Warn(message ...interface{}) {}
func (c *AbstractLog) Info(message ...interface{}) {}
func (c *AbstractLog) Debug(message ...interface{}) {}
func (c *AbstractLog) Trace(message ...interface{}) {}

func formatMessage(message []interface{}) string {
    if len(message) == 0 { return "" }
    
    output := fmt.Sprint(message[0]) 
    for _, m := range message[1:] {
        output += "," + fmt.Sprint(m)
    }
    return output
}
