package log

import "time"

type LogEntry struct {
    Time            time.Time           `json:"time"`
    Component       string              `json:"component"`
    Level           int                 `json:"level"`
    CorrelationId   string              `json:"correlationId"`
    Message         []interface{}       `json:"message"`  
}

func MakeLogEntry(component string, level int, correlationId string, message []interface{}) LogEntry {
    return LogEntry {
        Time: time.Now().UTC(),
        Component: component,
        Level: level,
        CorrelationId: correlationId,
        Message: message, 
    }
}