package deps

import (
    "bytes"
    "encoding/json"
    "io"
    "net/http"
    "strings"
    "github.com/pip-services/pip-services-runtime-go"
    "github.com/pip-services/pip-services-runtime-go/util"            
)

type RestClient struct {
    AbstractClient
    url string
    client *http.Client
}

func NewRestClient(id string, config *runtime.DynamicMap) *RestClient {
    defaultConfig := runtime.NewMapAndSet(
        "transport.type", "http",
        //"transport.host", "localhost",
        //"transport.port", 3000,
    )
    config = runtime.NewMapWithDefaults(config, defaultConfig)

    return &RestClient { AbstractClient: *NewAbstractClient(id, config) }
}

func (c *RestClient) Url() string {
    return c.url
}

func (c *RestClient) Client() *http.Client {
    return c.client
}

func (c *RestClient) Init(refs *runtime.References) error {
    err := c.AbstractClient.Init(refs)
    if err != nil { return err }
            
    // Check for type
    transportType := c.Config().GetString("transport.type")
    if transportType != "http" {
        return c.NewConfigError("Protocol is not supported by REST transport").WithDetails(transportType)
    }

    // Check for host
    transportHost := c.Config().GetString("transport.host")
    if transportHost == "" {
        return c.NewConfigError("Host is not configured in REST transport")
    }

    // Check for port
    transportPort := c.Config().GetString("transport.port")
    if transportPort == "" {
        return c.NewConfigError("Port is not configured in REST transport")
    }    
                  
    return nil
}

func (c *RestClient) Open() error {
    if c.client != nil { return nil }
        
    transportHost := c.Config().GetString("transport.host")
    transportPort := c.Config().GetString("transport.port")
    
    c.url = "http://" + transportHost + ":" + transportPort
    c.client = &http.Client{}
                        
    return c.AbstractClient.Open()    
}
    
func (c *RestClient) Close() error {
    if c.client != nil  {
        c.url = ""
        c.client = nil
    }
    
    return c.AbstractClient.Close()
} 

func (c *RestClient) QueryParams(params map[string] string) string {
    r := ""
    
    for key, value := range params {
        if r != "" { r = r + "&" }
        r = r + key + "=" + value
    }
    
    if r != "" { r = "?" + r }
    
    return r
}

func (c *RestClient) GetPagingAndFilterParams(filter *runtime.FilterParams, paging *runtime.PagingParams) string {
    r := map[string] string {}
    
    if filter != nil {
        for key, value := range *filter {
            r[key] = util.Converter.ToString(value)
        }
    }
    
    if paging != nil && paging.Paging {
        r["paging"] = util.Converter.ToString(paging.Paging)
        
        if paging.Skip != nil {
            r["skip"] = util.Converter.ToString(*paging.Skip)
        }
        
        if paging.Take != nil {
            r["take"] = util.Converter.ToString(*paging.Take)
        }
    }
    
    return c.QueryParams(r)
}

func (c *RestClient) Call(method string, route string,  input interface{}, output interface{}) (interface{}, error) {
    var body io.Reader
    
    // Prepare input data
    if input != nil {
        data, err1 := json.Marshal(input);
        if err1 != nil { 
            return nil, c.NewReadError("Failed to marshal JSON object", err1)  
        }
        body = bytes.NewBuffer(data)
    }
    
    // Create a request with parameters
    method = strings.ToUpper(method)
    req, err2 := http.NewRequest(method, c.url + route, body)
    if err2 != nil { 
        return nil, c.NewReadError("Call to REST API failed", err2)    
    }
    req.Header.Add("Content-Type", "application/json; charset=UTF-8")

    // Make a call to the service    
    res, err3 := c.client.Do(req)
    if err3 != nil { 
        return nil, c.NewReadError("Call to REST API failed", err3)    
    }

    // Exit if no result is expected
    if output == nil { return nil, nil }

    // Exit if no result was found
    if res.StatusCode == http.StatusNotFound { 
        return nil, nil 
    }

    // Exit if no result was send
    if res.StatusCode == http.StatusNoContent { 
        return nil, nil 
    }

    // Extract output data
    defer func() { res.Body.Close() }()

    err4 := json.NewDecoder(res.Body).Decode(output)
    if err4 != nil {
        return nil, c.NewReadError("Failed to unmarshal JSON object", err4)    
    }
        
    return output, nil
}
