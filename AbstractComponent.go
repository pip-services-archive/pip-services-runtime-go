package runtime

import "time"

type AbstractComponent struct {
    id string
    config *DynamicMap
    refs *References
    log ILog
    counters ICounters
}

func NewAbstractComponent(id string, config *DynamicMap) *AbstractComponent {
    if id == "" { panic("Component id is not set") }    
    if config == nil { config = NewEmptyMap() }
    
    return &AbstractComponent { id: id, config: config }
}

func (c *AbstractComponent) Id() string {
    return c.id
}

func (c *AbstractComponent) Config() *DynamicMap {
    return c.config
}

func (c *AbstractComponent) Refs() *References {
    return c.refs
}

/* Life cycle management */

func (c *AbstractComponent) Init(refs *References) error {
    if refs == nil { panic("Component references are not defined") }
    
    c.refs = refs
    c.log = refs.Log
    c.counters = refs.Counters
    
    return nil
}

func (c *AbstractComponent) Open() error {
    c.Trace("Component " + c.id + " opened")
    return nil
}

func (c *AbstractComponent) Close() error {
    c.Trace("Component " + c.id + " closed")
    return nil
}

/* Error handling */

func (c *AbstractComponent) NewError(message string) *MicroserviceError {
    return NewError(message).ForComponent(c.id)
}

func (c *AbstractComponent) NewInternalError(message string) *MicroserviceError {
    return NewInternalError(message).ForComponent(c.id)
}

func (c *AbstractComponent) NewConfigError(message string) *MicroserviceError {
    return NewConfigError(message).ForComponent(c.id)
}

func (c *AbstractComponent) NewStateError(message string) *MicroserviceError {
    return NewStateError(message).ForComponent(c.id)
}

func (c *AbstractComponent) NewOpenError(message string, err error) *MicroserviceError {
    return NewOpenError(message, err).ForComponent(c.id)
}

func (c *AbstractComponent) NewCloseError(message string, err error) *MicroserviceError {
    return NewCloseError(message, err).ForComponent(c.id)
}
func (c *AbstractComponent) NewReadError(message string, err error) *MicroserviceError {
    return NewWriteError(message, err).ForComponent(c.id)
}
func (c *AbstractComponent) NewWriteError(message string, err error) *MicroserviceError {
    return NewWriteError(message, err).ForComponent(c.id)
}
    
func (c *AbstractComponent) NewBadRequest(message string) *MicroserviceError {
    return NewBadRequest(message).ForComponent(c.id)
}

func (c *AbstractComponent) NewUnauthorized(message string) *MicroserviceError {
    return NewUnauthorized(message).ForComponent(c.id)
}

func (c *AbstractComponent) NewNotFound(message string) *MicroserviceError {
    return NewNotFound(message).ForComponent(c.id)
}

func (c *AbstractComponent) NewConflict(message string) *MicroserviceError {
    return NewConflict(message).ForComponent(c.id)
}
    
/* Logging */

func (c *AbstractComponent) Fatal(message ...interface{}) {
    if c.log != nil { c.log.Fatal(message) }
}

func (c *AbstractComponent) Error(message ...interface{}) {
    if c.log != nil { c.log.Error(message) }
}

func (c *AbstractComponent) Warn(message ...interface{}) {
    if c.log != nil { c.log.Warn(message) }
}

func (c *AbstractComponent) Info(message ...interface{}) {
    if c.log != nil { c.log.Info(message) }
}

func (c *AbstractComponent) Debug(message ...interface{}) {
    if c.log != nil { c.log.Debug(message) }
}

func (c *AbstractComponent) Trace(message ...interface{}) {
    if c.log != nil { c.log.Trace(message) }
}

/* Performance monitoring */

func (c *AbstractComponent) BeginTiming(name string) *Timing {
    if c.counters != nil {
        return c.counters.BeginTiming(name)
    } else {
        return EmptyTiming()
    }
}

func (c *AbstractComponent) Stats(name string, value float32) {
    if c.counters != nil { c.counters.Stats(name, value) }
}

func (c *AbstractComponent) Last(name string, value float32) {
    if c.counters != nil { c.counters.Last(name, value) }
}

func (c *AbstractComponent) TimestampNow(name string) {
    if c.counters != nil { c.counters.TimestampNow(name) }
}

func (c *AbstractComponent) Timestamp(name string, time time.Time) {
    if c.counters != nil { c.counters.Timestamp(name, time) }
}

func (c *AbstractComponent) IncrementOne(name string) {
    if c.counters != nil { c.counters.IncrementOne(name) }
}

func (c *AbstractComponent) Increment(name string, value int) {
    if c.counters != nil { c.counters.Increment(name, value) }
}
