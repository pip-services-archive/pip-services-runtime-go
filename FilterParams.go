package runtime

type FilterParams DynamicMap

func NewEmptyFilterParams() *FilterParams {
    return (*FilterParams)(NewEmptyMap())
}

func NewFilterParamsOf(value interface{}) *FilterParams {
    return (*FilterParams)(NewMapOf(value))
}

func NewFilterParamsAndSet(params ...interface{}) *FilterParams {
    c := &DynamicMap {}
    c.SetArray(params)
    return (*FilterParams)(c)
}

func NewFilterParamsFromJson(json string) *FilterParams {
    return (*FilterParams)(NewMapFromJson(json))
}
