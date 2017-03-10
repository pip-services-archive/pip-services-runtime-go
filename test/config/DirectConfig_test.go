package config

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "github.com/pip-services/pip-services-runtime-go"    
    "github.com/pip-services/pip-services-runtime-go/config"    
)

func TestDirectConfigRead(t *testing.T) {
    m := &runtime.DynamicMap { "test": 123 }
    c := config.NewDirectConfig(m)
    
    v, err := c.Read()
    assert.NotNil(t, v)
    assert.Nil(t, err)
    assert.Equal(t, 123, v.Get("test")) 
}