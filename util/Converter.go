package util

import (
    "encoding/json"
    "fmt"
    "reflect"
    "strings"
    "strconv"
    "time"
)

type TConverter struct{}

var Converter *TConverter = &TConverter{}

func (c *TConverter) ToNullableString(value interface{}) *string {
    if value == nil { return nil }
    
    switch value.(type) {
        case string:
            r := value.(string)
            return &r
        
        // case int
        // case int32:
        // case int64:
        //     r := strconv.FormatInt(value.(int64), 10)
        //     return &r
             
        // case byte:        
        // case uint:
        // case uint32:
        // case uint64:
        //     r := strconv.FormatUint(value.(uint64), 10)
        //     return &r
            
        case bool:
            r := "false"
            if value.(bool) { r = "true" }
            return &r 
            
        case time.Time:
            r := value.(time.Time).Format(time.RFC3339)
            return &r
            
        default:
            r := fmt.Sprint(value)
            return &r
    }
}

func (c *TConverter) ToString(value interface{}) string {
    return c.ToStringWithDefault(value, "") 
}

func (c *TConverter) ToStringWithDefault(value interface{}, defaultValue string) string {
    r := c.ToNullableString(value)
    if r == nil { return defaultValue }
    return *r 
}

func (c *TConverter) ToNullableBoolean(value interface{}) *bool {
    if value == nil { return nil }
    
    var v string
    
    switch value.(type) {
        case bool:
            r := value.(bool)
            return &r
            
        case string:
            v = strings.ToLower(value.(string))
            
        default:
            v = strings.ToLower(fmt.Sprint(value))
    }
    
    if (v == "1" || v == "true" || v == "t" || v == "yes" || v == "y") {
        r := true
        return &r        
    }
    
    if (v == "0" || v == "false" || v == "f" || v == "no" || v == "n") {
        r := false
        return &r        
    }
    
    return nil
}

func (c *TConverter) ToBoolean(value interface{}) bool {
    return c.ToBooleanWithDefault(value, false)
}

func (c *TConverter) ToBooleanWithDefault(value interface{}, defaultValue bool) bool {
    r := c.ToNullableBoolean(value)
    if r == nil { return defaultValue }
    return *r 
}

func (c *TConverter) ToNullableLong(value interface{}) *int64 {
    if value == nil { return nil }
    
    var r int64 = 0
    
    switch value.(type) {
        case int8:
            r = (int64)(value.(int8))
        case uint8:
            r = (int64)(value.(uint8))
        case int:
            r = (int64)(value.(int))
        case int16:
            r = (int64)(value.(int16))
        case uint16:
            r = (int64)(value.(uint16))
        case int32:
            r = (int64)(value.(int32))
        case uint32:
            r = (int64)(value.(uint32))
        case int64:
            r = (int64)(value.(int64))
        case uint64:
            r = (int64)(value.(uint64))
        case float32:
            r = (int64)(value.(float32))
        case float64:
            r = (int64)(value.(float64))

        case bool:
            v := value.(bool)
            if v == true { r = 1 }

        case time.Time:
            r = value.(time.Time).Unix()

        case string:
            v, ok := strconv.ParseFloat(value.(string), 0)
            if ok != nil { return nil }
            r = int64(v)
            
        default:
            return nil
    }
    
    return &r
}

func (c *TConverter) ToLong(value interface{}) int64 {
    return c.ToLongWithDefault(value, 0)
}

func (c *TConverter) ToLongWithDefault(value interface{}, defaultValue int64) int64 {
    r := c.ToNullableLong(value)
    if r == nil { return defaultValue }
    return *r 
}

func (c *TConverter) ToNullableInteger(value interface{}) *int {
    v := c.ToNullableLong(value)
    if v == nil { return nil }
    r := int(*v) 
    return &r
}

func (c *TConverter) ToInteger(value interface{}) int {
    return c.ToIntegerWithDefault(value, 0)
}

func (c *TConverter) ToIntegerWithDefault(value interface{}, defaultValue int) int {
    r := c.ToNullableInteger(value)
    if r == nil { return defaultValue }
    return *r 
}

func (c *TConverter) ToNullableFloat(value interface{}) *float64 {
    if value == nil { return nil }
    
    var r float64 = 0
    
    switch value.(type) {
        case int8:
            r = (float64)(value.(int8))
        case uint8:
            r = (float64)(value.(uint8))
        case int:
            r = (float64)(value.(int))
        case int16:
            r = (float64)(value.(int16))
        case uint16:
            r = (float64)(value.(uint16))
        case int32:
            r = (float64)(value.(int32))
        case uint32:
            r = (float64)(value.(uint32))
        case int64:
            r = (float64)(value.(int64))
        case uint64:
            r = (float64)(value.(uint64))
        case float32:
            r = (float64)(value.(float32))
        case float64:
            r = (float64)(value.(float64))

        case bool:
            v := value.(bool)
            if v == true { r = 1.0 }

        case time.Time:
            r = float64(value.(time.Time).Unix())

        case string:
            var ok error
            r, ok = strconv.ParseFloat(value.(string), 0)
            if ok != nil { return nil }
            
        default:
            return nil
    }
    
    return &r
}

func (c *TConverter) ToFloat(value interface{}) float64 {
    return c.ToFloatWithDefault(value, 0)
}

func (c *TConverter) ToFloatWithDefault(value interface{}, defaultValue float64) float64 {
    r := c.ToNullableFloat(value)
    if r == nil { return defaultValue }
    return *r 
}

func (c *TConverter) ToNullableDate(value interface{}) *time.Time {
    if value == nil { return nil }

    var r time.Time
    
    switch value.(type) {
        case int8:
            r = time.Unix((int64)(value.(int8)), 0)
        case uint8:
            r = time.Unix((int64)(value.(uint8)), 0)
        case int:
            r = time.Unix((int64)(value.(int)), 0)
        case int16:
            r = time.Unix((int64)(value.(int16)), 0)
        case uint16:
            r = time.Unix((int64)(value.(uint16)), 0)
        case int32:
            r = time.Unix((int64)(value.(int32)), 0)
        case uint32:
            r = time.Unix((int64)(value.(uint32)), 0)
        case int64:
            r = time.Unix((int64)(value.(int64)), 0)
        case uint64:
            r = time.Unix((int64)(value.(uint64)), 0)
        case float32:
            r = time.Unix((int64)(value.(float32)), 0)
        case float64:
            r = time.Unix((int64)(value.(float64)), 0)

        case time.Time:
            r = value.(time.Time)

        case string:
            v := value.(string)
            var ok error
            r, ok = time.Parse(time.RFC3339, v)
            if ok != nil {
                r, ok = time.Parse(time.RFC3339Nano, v) 
            }
            if (ok != nil) { return nil }
            
        default:
            return nil
    }
    
    return &r
}

func (c *TConverter) ToDate(value interface{}) time.Time {
    return c.ToDateWithDefault(value, time.Now())
}

func (c *TConverter) ToDateWithDefault(value interface{}, defaultValue time.Time) time.Time {
    r := c.ToNullableDate(value)
    if r == nil { return defaultValue }
    return *r 
}

func (c *TConverter) FromMultiString(value map[string]string, language string) string {
    r, ok := value[language]
    if ok { return r }
    
    r, ok = value["en"]
    if ok { return r }
    
    for _, r = range value {
        return r
    } 
    
    return ""
}

func (c *TConverter) arrayToMap(value reflect.Value) interface{} {
    r := []interface{}{}

    for i := 0; i < value.Len(); i++ {
        v := value.Index(i)
        r = append(r, c.valueToMap(v))
    }
           
    return r
}

func (c *TConverter) mapToMap(value reflect.Value) interface{} {
    r := map[string]interface{}{}

    for _, key := range value.MapKeys() {
        k := c.ToString(c.valueToMap(key))
        v := c.valueToMap(value.MapIndex(key))
        r[k] = v 
    }
    
    return r
}

func (c *TConverter) structToMap(value reflect.Value) interface{} {
    r := map[string]interface{}{}
    
    t := value.Type()
        
    if t.Name() == "Time" {
        return value.Interface()
    }
    
    for i := 0; i < value.NumField(); i++ {
        k := t.Field(i).Name
        v := c.valueToMap(value.Field(i))
        r[k] = v
    }
    
    return r
}

func (c *TConverter) valueToMap(value reflect.Value) interface{} {    
    switch value.Kind() {
        case reflect.Invalid:
            return nil
        case reflect.String:
            return value.String()
        case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
            return int64(value.Int())
        case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
            return int64(value.Uint())
        case reflect.Float32, reflect.Float64:
            return float64(value.Float())
        case reflect.Bool:
            return value.Bool()
            
        case reflect.Map:
            return c.mapToMap(value)
        case reflect.Array, reflect.Slice:
            return c.arrayToMap(value)
        case reflect.Struct:
            return c.structToMap(value)   
        case reflect.Interface:
            if value.IsNil() { return nil }
            return c.valueToMap(value.Elem())
    }
        
    return nil
}

func (c *TConverter) ToNullableMap(value interface{}) *map[string]interface{} {
    if value == nil { return nil }
    
    // Parse JSON
    if s, ok := value.(string); ok == true {
        var m map[string]interface{}
        if err := json.Unmarshal([]byte(s), &m); err != nil { 
            return nil 
        }
        value = m
    }

    // Convert to map
    value = c.valueToMap(reflect.ValueOf(value))
    
    if m, ok := value.(map[string]interface{}); ok {
        return &m
    }
    
    return nil  
} 

func (c *TConverter) ToMap(value interface{}) map[string]interface{} {
    return c.ToMapWithDefault(value, map[string]interface{} {})
}

func (c *TConverter) ToMapWithDefault(value interface{}, defaultValue map[string]interface{}) map[string]interface{} {
    if m := c.ToNullableMap(value); m != nil { return *m } 
    return map[string]interface{} {}
} 