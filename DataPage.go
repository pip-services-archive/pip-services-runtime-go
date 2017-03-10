package runtime

type DataPage struct {
    Total *int32         `json:"total"`
    Data []interface{}  `json:"data"`
}

func NewEmptyDataPage() *DataPage {
    return &DataPage{}
}

func NewDataPage(total *int32, data []interface{}) *DataPage {
    return &DataPage{ Total: total, Data: data }
}