package util

import (
    "time"
    "testing"
    "github.com/stretchr/testify/assert"
    "github.com/pip-services/pip-services-runtime-go/util"
)

func TestToString(t *testing.T) {
    assert.Nil(t, util.Converter.ToNullableString(nil))

    assert.Equal(t, "xyz", util.Converter.ToString("xyz"))
    assert.Equal(t, "123", util.Converter.ToString(123))
    assert.Equal(t, "true", util.Converter.ToString(true))
    
    value := struct{ prop string }{ "xyz" }
    assert.Equal(t, "{xyz}", util.Converter.ToString(value))

    assert.Equal(t, "xyz", util.Converter.ToStringWithDefault(nil, "xyz"))
}

func TestToBoolean(t *testing.T) {
    assert.Nil(t, util.Converter.ToNullableBoolean(nil))
    
    assert.True(t, util.Converter.ToBoolean(true))
    assert.True(t, util.Converter.ToBoolean(1))
    assert.True(t, util.Converter.ToBoolean("True"))
    assert.True(t, util.Converter.ToBoolean("yes"))
    assert.True(t, util.Converter.ToBoolean("1"))
    assert.True(t, util.Converter.ToBoolean("Y"))

    assert.False(t, util.Converter.ToBoolean(false))
    assert.False(t, util.Converter.ToBoolean(0))
    assert.False(t, util.Converter.ToBoolean("False"))
    assert.False(t, util.Converter.ToBoolean("no"))
    assert.False(t, util.Converter.ToBoolean("0"))
    assert.False(t, util.Converter.ToBoolean("N"))
    
    assert.False(t, util.Converter.ToBoolean(123))
    assert.True(t, util.Converter.ToBooleanWithDefault("XYZ", true))
}

func TestToInteger(t *testing.T) {
    assert.Nil(t, util.Converter.ToNullableInteger(nil))
    
    assert.Equal(t, int(123), util.Converter.ToInteger(123))
    assert.Equal(t, int(123), util.Converter.ToInteger(123.456))
    assert.Equal(t, int(123), util.Converter.ToInteger("123"))
    assert.Equal(t, int(123), util.Converter.ToInteger("123.456"))

    assert.Equal(t, int(123), util.Converter.ToIntegerWithDefault(nil, 123))
    assert.Equal(t, int(0), util.Converter.ToIntegerWithDefault(false, 123))
    assert.Equal(t, int(123), util.Converter.ToIntegerWithDefault("ABC", 123))
}

func TestToLong(t *testing.T) {
    assert.Nil(t, util.Converter.ToNullableLong(nil))
    
    assert.Equal(t, int64(123), util.Converter.ToLong(123))
    assert.Equal(t, int64(123), util.Converter.ToLong(123.456))
    assert.Equal(t, int64(123), util.Converter.ToLong("123"))
    assert.Equal(t, int64(123), util.Converter.ToLong("123.456"))

    assert.Equal(t, int64(123), util.Converter.ToLongWithDefault(nil, 123))
    assert.Equal(t, int64(0), util.Converter.ToLongWithDefault(false, 123))
    assert.Equal(t, int64(123), util.Converter.ToLongWithDefault("ABC", 123))
}

func TestToFloat(t *testing.T) {
	assert.Nil(t, util.Converter.ToNullableFloat(nil))

	assert.Equal(t, 123., util.Converter.ToFloat(123))
	assert.Equal(t, 123.456, util.Converter.ToFloat(123.456))
	assert.Equal(t, 123., util.Converter.ToFloat("123"))
	assert.Equal(t, 123.456, util.Converter.ToFloat("123.456"))

	assert.Equal(t, 123., util.Converter.ToFloatWithDefault(nil, 123))
	assert.Equal(t, 0., util.Converter.ToFloatWithDefault(false, 123))
	assert.Equal(t, 123., util.Converter.ToFloatWithDefault("ABC", 123))
}

func TestToDate(t *testing.T) {
    assert.Nil(t, util.Converter.ToNullableDate(nil))

    date1 := time.Date(1975, time.April, 8, 0, 0, 0, 0, time.UTC)
    assert.Equal(t, date1, util.Converter.ToDateWithDefault(nil, date1))
    assert.Equal(t, date1, util.Converter.ToDate(date1))
    assert.Equal(t, date1, util.Converter.ToDate("1975-04-08T00:00:00Z"))
    assert.Equal(t, date1, util.Converter.ToDate("1975-04-08T00:00:00.00Z"))

    date2 := time.Unix(123, 0)
    assert.Equal(t, date2, util.Converter.ToDate(123))
    assert.Equal(t, date2, util.Converter.ToDate(123.456))
}

func TestObjectToMap(t *testing.T) {
    assert.Nil(t, util.Converter.ToNullableMap(nil))

    v1 := struct{ value1, value2 float64 }{ 123, 234 }
    m := util.Converter.ToMap(v1)
    assert.Len(t, m, 2)
    assert.Equal(t, 123., m["value1"])
    assert.Equal(t, 234., m["value2"])
    
    v2 := map[string]interface{} { "value1": 123 }
    m = util.Converter.ToMap(v2)
    assert.Len(t, m, 1)
    assert.Equal(t, int64(123), m["value1"])
}

func TestJsonToMap(t *testing.T) {
    // Handling simple objects
    v := `{ "value1":123, "value2":234 }`
    m := util.Converter.ToMap(v)
    assert.NotNil(t, m)
    assert.Len(t, m, 2)
    assert.Equal(t, 123., m["value1"])
    assert.Equal(t, 234., m["value2"])
    
    // Recursive conversion
    v = `{ "value1":123, "value2": { "value21": 111, "value22": 222} }`
    m = util.Converter.ToMap(v)
    assert.NotNil(t, m)
    assert.Len(t, m, 2)
    assert.Equal(t, 123., m["value1"])

    m2 := m["value2"].(map[string]interface{})
    assert.Len(t, m2, 2) 
    assert.Equal(t, 111., m2["value21"])
    assert.Equal(t, 222., m2["value22"])
    
    // Handling arrays
    v = `{ "value1":123, "value2": [{ "value21": 111, "value22": 222}] }`
    m = util.Converter.ToMap(v)
    assert.NotNil(t, m)
    assert.Len(t, m, 2)
    assert.Equal(t, 123., m["value1"])

    a2 := m["value2"].([]interface{})
    assert.NotNil(t, a2)
    assert.Len(t, a2, 1) 

    m2 = a2[0].(map[string]interface{})
    assert.Len(t, m2, 2) 
    assert.Equal(t, 111., m2["value21"])
    assert.Equal(t, 222., m2["value22"])
}