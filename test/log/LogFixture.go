package log

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "github.com/pip-services/pip-services-runtime-go"    
)

type LogFixture struct {
    log runtime.ILog
}

func NewLogFixture(log runtime.ILog) *LogFixture {
    c := LogFixture{ log: log }
    return &c
}

func (c *LogFixture) TestLogLevel(t *testing.T) {
    assert.True(t, c.log.Level() >= runtime.LogLevelNone)
    assert.True(t, c.log.Level() <= runtime.LogLevelTrace)
}

func (c *LogFixture) TestTextOutput(t *testing.T) {
    c.log.Fatal("Fatal error...")
    c.log.Error("Recoverable error...")
    c.log.Warn("Warning...")
    c.log.Info("Information message...")
    c.log.Debug("Debug message...")
    c.log.Trace("Trace message...")
}

func (c *LogFixture) TestMixedOutput(t *testing.T) {
    type testClass struct { abc string }
    obj := testClass{ "ABC" }
    
    c.log.Fatal(123, "ABC", obj)
    c.log.Error(123, "ABC", obj)
    c.log.Warn(123, "ABC", obj)
    c.log.Info(123, "ABC", obj)
    c.log.Debug(123, "ABC", obj)
    c.log.Trace(123, "ABC", obj)
}