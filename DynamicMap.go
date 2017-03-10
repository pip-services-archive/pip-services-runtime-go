package runtime

import (
    "reflect"
    "strings"
    "time"
    "github.com/pip-services/pip-services-runtime-go/util"
)

type DynamicMap map[string]interface{}

/********* Constructors ***********/

func NewEmptyMap() *DynamicMap {
    return &DynamicMap {}
}

func NewMapOf(value interface{}) *DynamicMap {
    c := DynamicMap(util.Converter.ToMap(value))
    return &c
}

func NewMapFromJson(json string) *DynamicMap {
    c := DynamicMap(util.Converter.ToMap(json))
    return &c
}

func NewMapAndSet(params ...interface{}) *DynamicMap {
    c := DynamicMap {}
    c.SetArray(params)
    return &c
}

func NewMapWithDefaults(m *DynamicMap, defaults *DynamicMap) *DynamicMap {
    if m == nil { m = new(DynamicMap) }
    return m.SetDefaultsDeep(defaults)
}

/********** Getters ************/

func (c *DynamicMap) Get(prop string) interface{} {
    if prop == "" { return nil }
    
    props := strings.Split(prop, ".")
    var r interface{} = *(*map[string]interface{})(c)
    
    for _, p := range props {
        m, ok := r.(map[string]interface{})
        if ok == false { return nil }
        
        r, ok = m[p]
        if ok == false { return nil }
    }
    
    return r
}

func (c *DynamicMap) Has(prop string) bool {
    return c.Get(prop) != nil
}

func (c *DynamicMap) HasNot(prop string) bool {
    return c.Get(prop) == nil
}

func (c *DynamicMap) GetNullableMap(prop string) *DynamicMap {    
    if v := c.Get(prop); v != nil {
        if r, ok := v.(map[string]interface{}); ok == true {
            var m DynamicMap = r 
            return &m
        }
    }
    return nil
}

func (c *DynamicMap) GetMap(prop string) *DynamicMap {
    r := c.GetNullableMap(prop)
    if r != nil { return r }
    return NewEmptyMap()
}

func (c *DynamicMap) GetMapWithDefault(prop string, defaultValue *DynamicMap) *DynamicMap {
    r := c.GetNullableMap(prop)
    if r != nil { return r }
    return defaultValue
}

func (c *DynamicMap) GetNullableString(prop string) *string {
    return util.Converter.ToNullableString(c.Get(prop))        
}

func (c *DynamicMap) GetString(prop string) string {
    return util.Converter.ToString(c.Get(prop))        
}

func (c *DynamicMap) GetStringWithDefault(prop string, defaultValue string) string {
    return util.Converter.ToStringWithDefault(c.Get(prop), defaultValue)        
}

func (c *DynamicMap) GetNullableBoolean(prop string) *bool {
    return util.Converter.ToNullableBoolean(c.Get(prop))        
}

func (c *DynamicMap) GetBoolean(prop string) bool {
    return util.Converter.ToBoolean(c.Get(prop))        
}

func (c *DynamicMap) GetBooleanWithDefault(prop string, defaultValue bool) bool {
    return util.Converter.ToBooleanWithDefault(c.Get(prop), defaultValue)        
}

func (c *DynamicMap) GetNullableInteger(prop string) *int {
    return util.Converter.ToNullableInteger(c.Get(prop))        
}

func (c *DynamicMap) GetInteger(prop string) int {
    return util.Converter.ToInteger(c.Get(prop))        
}

func (c *DynamicMap) GetIntegerWithDefault(prop string, defaultValue int) int {
    return util.Converter.ToIntegerWithDefault(c.Get(prop), defaultValue)        
}

func (c *DynamicMap) GetNullableLong(prop string) *int64 {
    return util.Converter.ToNullableLong(c.Get(prop))        
}

func (c *DynamicMap) GetLong(prop string) int64 {
    return util.Converter.ToLong(c.Get(prop))        
}

func (c *DynamicMap) GetLongWithDefault(prop string, defaultValue int64) int64 {
    return util.Converter.ToLongWithDefault(c.Get(prop), defaultValue)        
}

func (c *DynamicMap) GetNullableFloat(prop string) *float64 {
    return util.Converter.ToNullableFloat(c.Get(prop))        
}

func (c *DynamicMap) GetFloat(prop string) float64 {
    return util.Converter.ToFloat(c.Get(prop))        
}

func (c *DynamicMap) GetFloatWithDefault(prop string, defaultValue float64) float64 {
    return util.Converter.ToFloatWithDefault(c.Get(prop), defaultValue)        
}

func (c *DynamicMap) GetNullableDate(prop string) *time.Time {
    return util.Converter.ToNullableDate(c.Get(prop))        
}

func (c *DynamicMap) GetDate(prop string) time.Time {
    return util.Converter.ToDate(c.Get(prop))        
}

func (c *DynamicMap) GetDateWithDefault(prop string, defaultValue time.Time) time.Time {
    return util.Converter.ToDateWithDefault(c.Get(prop), defaultValue)        
}

/********* Setters *********/

func (c *DynamicMap) Set(prop string, value interface{}) {
    if prop == "" { return }
    
    props := strings.Split(prop, ".")       
    var v *map[string]interface{} = (*map[string]interface{})(c)
    
    for _, p := range props[:len(props) - 1] {
        r, ok1 := (*v)[p]
        if ok1 == false { 
            // Create a map when necessary
            m := map[string]interface{} {}
            (*v)[p], r = m, m
        }
        
        // Get a map
        m, ok2 := r.(map[string]interface{})
        if ok2 == false { return }
        
        v = &m
    }
    
    // Set value to the map property
    (*v)[props[len(props) - 1]] = value
}

func (c *DynamicMap) SetArray(params []interface{}) {
    for i := 0; i < len(params) - 1; i += 2 {
        prop := util.Converter.ToString(params[i])
        value := params[i + 1]
        c.Set(prop, value)
    }
}

func (c *DynamicMap) SetAll(params ...interface{}) {
    c.SetArray(params)
}

func (c *DynamicMap) Remove(prop string) {
    delete(*c, prop)
}

func (c *DynamicMap) RemoveAll(props ...string) {
    for _, p := range props {
        delete(*c, p)
    }
}

/************* Merging ************/

func (c *DynamicMap) Merge(other *DynamicMap, deep bool) *DynamicMap {
    r := DynamicMap {}
        
    // Copy all from original map
    for k, v := range *c {
        r[k] = v
    }
    
    if other == nil { return &r }
    
    // Add from other map
    for k, v2 := range *other {        
        if v1 := r[k]; v1 != nil {
            m1, ok1 := v1.(map[string]interface{})
            m2, ok2 := v2.(map[string]interface{})
            
            if deep && ok1 && ok2 {
                c1 := DynamicMap(m1) 
                c2 := DynamicMap(m2)
                c3 := *(&c1).Merge(&c2, deep)
                r[k] = (map[string]interface{})(c3)                
            }
        } else {
            r[k] = v2
        }
    }
    
    return &r    
}

func (c *DynamicMap) MergeDeep(other *DynamicMap) *DynamicMap {
    return c.Merge(other, true)
}

func (c *DynamicMap) SetDefaults(defaults *DynamicMap, deep bool) *DynamicMap {
    return c.Merge(defaults, deep)
}

func (c *DynamicMap) SetDefaultsDeep(defaults *DynamicMap) *DynamicMap {
    return c.Merge(defaults, true)
}

/********* Other Utilities ***********/

func (c *DynamicMap) AssignTo(value interface{}) {
    if value == nil || len(*c) == 0 { return }

    // Value must be a pointer to a struct    
    v := reflect.ValueOf(value)
    if v.Kind() != reflect.Ptr { return }

    // Get the struct
    v = v.Elem()
    if v.Kind() != reflect.Struct { return }

    // Assign matching fields from the 1st level
    for fn, fv := range *c {
        f := v.FieldByNameFunc(func (n string) bool { return strings.EqualFold(fn, n) })
        if f.CanSet() {
            f.Set(reflect.ValueOf(fv))
        }
    }
}

func (c *DynamicMap) Pick(props ...string) *DynamicMap {
    r := DynamicMap {}    

    for _, p := range props {
        if v, ok := (*c)[p]; ok == true {
            r[p] = v
        }
    }    

    return &r
}

func (c *DynamicMap) Omit(props ...string) *DynamicMap {
    var r DynamicMap = DynamicMap {}    

    // Copy over
    for k, v := range *c {
        r[k] = v
    }

    for _, p := range props {
        delete(r, p)
    }    

    return &r
}