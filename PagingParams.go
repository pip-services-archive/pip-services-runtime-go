package runtime

import (
    "github.com/pip-services/pip-services-runtime-go/util"
)

type PagingParams struct {
    Skip *int
    Take *int
    Paging bool
}

func NewEmptyPagingParams() *PagingParams {
    return &PagingParams { Skip: nil, Take: nil, Paging: true }
}

func NewPagingParams(skip, take, paging interface{}) *PagingParams {
    c := PagingParams {}
    
    c.Skip = util.Converter.ToNullableInteger(skip)
    c.Take = util.Converter.ToNullableInteger(take)
    c.Paging = util.Converter.ToBooleanWithDefault(paging, true)
    
    return &c
}

func NewPagingParamsFromMap(value *DynamicMap) *PagingParams {
    c := PagingParams {}
    
    c.Skip = value.GetNullableInteger("skip")
    c.Take = value.GetNullableInteger("take")
    c.Paging = value.GetBooleanWithDefault("paging", true)
    
    return &c
}

func (c *PagingParams) GetSkip(minSkip int) int {
    if c.Skip == nil { return minSkip }
    if *c.Skip < minSkip { return minSkip }
    return *c.Skip
}

func (c *PagingParams) GetTake(maxTake int) int {
    if c.Take == nil { return maxTake }
    if *c.Take < 0 { return 0 }
    if *c.Take > maxTake { return maxTake }
    return *c.Take
} 