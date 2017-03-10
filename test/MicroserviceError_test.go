package test

import ( 
    "testing"
    "github.com/stretchr/testify/assert"
    "github.com/pip-services/pip-services-runtime-go"
)

func TestMicroserviceError(t *testing.T) {
    error := runtime.NewError("Test error").ForComponent("TestComponent").WithCode("TestError")
    assert.Equal(t, "TestComponent", error.Component)
    assert.Equal(t, "TestError", error.Code)
    assert.Equal(t, "Test error", error.Message)
    
    error = runtime.NewError("").ForComponent("TestComponent")
    assert.Equal(t, "InternalError", error.Code)
    assert.Equal(t, "Internal error", error.Message)
}