package config

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "github.com/pip-services/pip-services-runtime-go/config"    
    "github.com/pip-services/pip-services-runtime-go/util"    
)

func TestJsonConfigRead(t *testing.T) {
    c := config.NewJsonConfigFromFile("./options.json")

    err := c.Open()
    assert.Nil(t, err)
    
    v, err := c.Read()
    assert.NotNil(t, v)
    assert.Nil(t, err)
    assert.Equal(t, 123, v.GetInteger("test")) 
    
    a := v.Get("array").([]interface{})
    assert.Len(t, a, 2)
    assert.Equal(t, 111, util.Converter.ToInteger(a[0]))
    assert.Equal(t, 222, util.Converter.ToInteger(a[1]))
}