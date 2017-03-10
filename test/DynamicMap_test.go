package test

import ( 
    "testing"
    "github.com/stretchr/testify/assert"
    "github.com/pip-services/pip-services-runtime-go"
)

func TestNewMap(t *testing.T) {
    m := runtime.NewEmptyMap()
    assert.Len(t, *m, 0)
    
    m = runtime.NewMapFromJson(`{ "value1": 123, "value2": 456 }`)
    assert.Len(t, *m, 2)
    
    v := map[string]int { "value1": 123, "value2": 456 }
    m = runtime.NewMapOf(v) 
    assert.Len(t, *m, 2)
    
    dm := runtime.DynamicMap{ "value1": 123, "value2": 456 }
    assert.Len(t, dm, 2)

    m = runtime.NewMapAndSet("value1", 123, "value2", 456)
    assert.Len(t, *m, 2)
}

func TestMapGet(t *testing.T) {
    m := runtime.NewMapFromJson(`{ "value1": 123, "value2": { "value21": 111, "value22": 222 } }`)
    assert.NotNil(t, m)
    assert.Len(t, *m, 2)
    
    v := m.Get("")
    assert.Nil(t, v)

    v = m.Get("value1")
    assert.Equal(t, 123., v)

    v = m.Get("value2")
    assert.NotNil(t, v)
    assert.Len(t, v, 2)

    v = m.Get("value3")
    assert.Nil(t, v)

    v = m.Get("value2.value21")
    assert.Equal(t, 111., v)

    v = m.Get("value2.value31")
    assert.Nil(t, v)

    v = m.Get("value2.value21.value211")
    assert.Nil(t, v)

    v = m.Get("valueA.valueB.valueC")
    assert.Nil(t, v)
}

func TestMapSet(t *testing.T) {
    m := runtime.DynamicMap {}
    
    m.Set("", 123)
    assert.Len(t, m, 0)
    
    m.Set("field1", 123)
    assert.Len(t, m, 1)
    assert.Equal(t, 123, m["field1"])
    
    m.Set("field2", "ABC")
    assert.Len(t, m, 2)
    assert.Equal(t, "ABC", m["field2"])

    m.Set("field2.field21", 123)
    assert.Equal(t, "ABC", m["field2"])
    
    m.Set("field3.field31", 456)
    assert.Len(t, m, 3)
    m3 := m["field3"].(map[string]interface{})
    assert.Len(t, m3, 1)    
    assert.Equal(t, 456, m3["field31"])

    m.Set("field3.field32", "XYZ")
    assert.Len(t, m3, 2)    
    assert.Equal(t, "XYZ", m3["field32"])    
}

func TestMapSetDefaults(t *testing.T) {
    m1 := runtime.NewMapFromJson(`{ "value1":123, "value2": 234 }`)
    m2 := runtime.NewMapFromJson(`{ "value2": 432, "value3": 345 }`)
    m := m1.SetDefaults(m2, false)
    
    assert.Len(t, *m, 3)
    assert.Equal(t, 123., (*m)["value1"])
    assert.Equal(t, 234., (*m)["value2"])
    assert.Equal(t, 345., (*m)["value3"])
}

func TestMapSetDefaultsDeep(t *testing.T) {
    m1 := runtime.NewMapFromJson(`{ "value1":123, "value2": { "value21": 111, "value22": 222 } }`)
    m2 := runtime.NewMapFromJson(`{ "value2": { "value22": 777, "value23": 333 }, "value3": 345 }`)
    m := m1.SetDefaultsDeep(m2)

    assert.Len(t, *m, 3)
    assert.Equal(t, 123., (*m)["value1"])
    assert.Equal(t, 345., (*m)["value3"])

    v := (*m)["value2"].(map[string]interface{})
    assert.Len(t, v, 3)
    assert.Equal(t, 111., v["value21"])
    assert.Equal(t, 222., v["value22"])
    assert.Equal(t, 333., v["value23"])
}

func TestSetDefaultsWithNulls(t *testing.T) {
    m1 := runtime.NewMapFromJson(`{ "value1":123, "value2": 234 }`)
    m := m1.SetDefaultsDeep(nil)

    assert.Len(t, *m, 2)    
    assert.Equal(t, 123., (*m)["value1"])
    assert.Equal(t, 234., (*m)["value2"])
}

func TestAssign(t *testing.T) {
    type testStruct struct { Value1 float64; _value2 string }
    
    v := testStruct{}
    m := runtime.NewMapFromJson(`{ "value1": 123, "value2": "ABC", "value3": 456 }`)
    
    m.AssignTo(&v)
    assert.Equal(t, 123., v.Value1)
    assert.Empty(t, v._value2)
}

