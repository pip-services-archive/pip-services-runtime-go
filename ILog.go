package runtime

const (
    LogLevelNone = 0
    LogLevelFatal = 1
    LogLevelError = 2
    LogLevelWarn = 3
    LogLevelInfo = 4
    LogLevelDebug = 5
    LogLevelTrace = 6
)

type ILog interface{
    IComponent
    
    Level() int
    SetLevel(value int)
    
    Fatal(message ...interface{})
    Error(message ...interface{})
    Warn(message ...interface{})
    Info(message ...interface{})    
    Debug(message ...interface{})
    Trace(message ...interface{})
}


