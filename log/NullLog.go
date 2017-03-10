package log

import (
    "github.com/pip-services/pip-services-runtime-go"    
)

type NullLog struct {
    AbstractLog
}

func NewNullLog(config *runtime.DynamicMap) *NullLog {
    return &NullLog { AbstractLog: *NewAbstractLog("NullLog", config) }
}

func (c *NullLog) Fatal(message ...interface{}) {}
func (c *NullLog) Error(message ...interface{}) {}
func (c *NullLog) Warn(message ...interface{}) {}
func (c *NullLog) Info(message ...interface{}) {}
func (c *NullLog) Debug(message ...interface{}) {}
func (c *NullLog) Trace(message ...interface{}) {}
