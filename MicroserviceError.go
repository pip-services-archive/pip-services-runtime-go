package runtime

type MicroserviceError struct {
    Component string        `json:"component"`
    Code string             `json:"code"`
    Message string          `json:"message"`
    Status int              `json:"status"`
    Details []interface{}   `json:"details"`
    Correlation string      `json:"correlation"`
    Cause error             `json:"cause"`
}

func (e *MicroserviceError) Error() string {
    return e.Message
}

func (e *MicroserviceError) ForComponent(component string) *MicroserviceError {
    e.Component = component
    return e
}

func (e *MicroserviceError) WithCode(code string) *MicroserviceError {
    e.Code = code
    return e
}

func (e *MicroserviceError) WithStatus(status int) *MicroserviceError {
    e.Status = status
    return e
}

func (e *MicroserviceError) WithDetails(details ...interface{}) *MicroserviceError {
    e.Details = details
    return e
}

func (e *MicroserviceError) WithCause(cause error) *MicroserviceError {
    e.Cause = cause
    return e
}

func (e *MicroserviceError) WithCorrelation(correlation string) *MicroserviceError {
    e.Correlation = correlation
    return e
}

func WrapError(err error) *MicroserviceError {
    return WrapErrorWithMessage(err, "");
}

func WrapErrorWithMessage(err error, message string) *MicroserviceError {
    if merr, ok := err.(*MicroserviceError); ok == true {
        return merr;
    }
    
    return NewInternalError(message).WithCause(err);
}

func NewError(message string) *MicroserviceError {
    if message == "" { message = "Internal error" }
    return &MicroserviceError{ Code: "InternalError", Message: message, Status: 500 }
}

func NewInternalError(message string) *MicroserviceError {
    if message == "" { message = "Internal error" }
    return &MicroserviceError{ Code: "InternalError", Message: message, Status: 500 }
}

func NewBuildError(message string) *MicroserviceError {
    if message == "" { message = "Build failed" }
    return &MicroserviceError{ Code: "BuildError", Message: message, Status: 500 }
}

func NewConfigError(message string) *MicroserviceError {
    if message == "" { message = "Wrong configuration" }
    return &MicroserviceError{ Code: "ConfigError", Message: message, Status: 500 }
}

func NewStateError(message string) *MicroserviceError {
    if message == "" { message = "Illegal state" }
    return &MicroserviceError{ Code: "StateError", Message: message, Status: 500 }
}

func NewOpenError(message string, cause error) *MicroserviceError {
    if message == "" { message = "Failed to open" }
    return &MicroserviceError{ Code: "OpenError", Message: message, Status: 500, Cause: cause }
}

func NewCloseError(message string, cause error) *MicroserviceError {
    if message == "" { message = "Failed to close" }
    return &MicroserviceError{ Code: "CloseError", Message: message, Status: 500, Cause: cause }
}

func NewCallError(message string, cause error) *MicroserviceError {
    if message == "" { message = "Failed to call" }
    return &MicroserviceError{ Code: "CallError", Message: message, Status: 500, Cause: cause }
}

func NewReadError(message string, cause error) *MicroserviceError {
    if message == "" { message = "Failed to read" }
    return &MicroserviceError{ Code: "ReadError", Message: message, Status: 500, Cause: cause }
}

func NewWriteError(message string, cause error) *MicroserviceError {
    if message == "" { message = "Failed to write" }
    return &MicroserviceError{ Code: "WriteError", Message: message, Status: 500, Cause: cause }
}

func NewBadRequest(message string) *MicroserviceError {
    if message == "" { message = "Bad request" }
    return &MicroserviceError{ Code: "BadRequest", Message: message, Status: 400 }
}

func NewUnauthorized(message string) *MicroserviceError {
    if message == "" { message = "Unauthorized access" }
    return &MicroserviceError{ Code: "Unauthorized", Message: message, Status: 401 }
}

func NewNotFound(message string) *MicroserviceError {
    if message == "" { message = "Object was not found" }
    return &MicroserviceError{ Code: "NotFound", Message: message, Status: 404 }
}

func NewConflict(message string) *MicroserviceError {
    if message == "" { message = "Conflict detected" }
    return &MicroserviceError{ Code: "Conflict", Message: message, Status: 409 }
}
