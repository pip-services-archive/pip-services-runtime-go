package api

import (
    "encoding/json"
    "net"
    "net/http"
    "reflect"
    "strings"
    "time"
    "github.com/gorilla/mux"
    "github.com/pip-services/pip-services-runtime-go"
    "github.com/pip-services/pip-services-runtime-go/util"            
)

type RestService struct {
    AbstractService
    
    router *mux.Router
    listener net.Listener
    server *http.Server
    registrations func()
}

func NewRestService(id string, config *runtime.DynamicMap) *RestService {
    defaultConfig := runtime.NewMapAndSet(
        "transport.type", "http",
        "transport.host", "0.0.0.0",
        //"transport.port", 3000,
        "transport.requestMaxSize", 1024 * 1024,
        "transport.connectTimeout", 60000,
        "transport.debug", true,
    )
    config = runtime.NewMapWithDefaults(config, defaultConfig)
        
    return &RestService { AbstractService: *NewAbstractService(id, config) }
}

func (c *RestService) SetRegistrations(registrations func()) {
    c.registrations = registrations
}

func (c *RestService) Init(refs *runtime.References) error {
    err := c.AbstractService.Init(refs)
    if err != nil { return err }
            
    // Check for type
    transportType := c.Config().GetString("transport.type")
    if transportType != "http" {
        return c.NewConfigError("Protocol is not supported by REST transport").WithDetails(transportType)
    }

    // Check for port
    transportPort := c.Config().GetString("transport.port")
    if transportPort == "" {
        return c.NewConfigError("Port is not configured in REST transport")
    }
    
    c.router = mux.NewRouter()
    c.registrations()
               
    return nil
}

func (c *RestService) Open() error {
    // Check for initialization
    if c.router == nil {
        return c.NewStateError("REST service is not initialized")
    }

    // Check for previous opening
    if c.server != nil { return nil }
        
    // Get transport config parameters
    transportHost := c.Config().GetString("transport.host")
    transportPort := c.Config().GetString("transport.port")
    transportTimeout := c.Config().GetIntegerWithDefault("transport.connectTimeout", 60000)
    transportMaxSize := c.Config().GetIntegerWithDefault("transport.requestMaxSize", 1024 * 1024)
    
    // Create TCP listener
    address := transportHost + ":" + transportPort
    listener, err := net.Listen("tcp", address)
    if err != nil {
        return c.NewOpenError("Failed to open server port", err).WithDetails(address)
    }

    // Create HTTP server
    c.listener = listener    
    c.server = &http.Server{
        Handler: c.router,
        ReadTimeout: time.Duration(transportTimeout) * time.Millisecond,
        WriteTimeout: time.Duration(transportTimeout) * time.Millisecond,
        MaxHeaderBytes: int(transportMaxSize), 
    }

    // Start accepting calls in a separate thread
    go func() {
        err := c.server.Serve(c.listener)
        if err != nil {
            c.Error(err)                    
        }
    }()
               
    c.Info("REST service started listening at http://" + transportHost + ":" + transportPort)
                        
    return c.AbstractService.Open()    
}
    
func (c *RestService) Close() error {
    if c.listener != nil  {
        // Close TCP listener
        err := c.listener.Close()
        if err != nil {
            c.Error(err)
        }
        
        // Clean up references
        c.listener = nil
        c.server = nil
    }
    
    return c.AbstractService.Close()
} 

func (c *RestService) Server() *http.Server {
    return c.server
}

func (c *RestService) Router() *mux.Router {
    return c.router
}

func (c *RestService) Register(method string, route string, handler func(w http.ResponseWriter, r *http.Request)) {
    method = strings.ToUpper(method)
    c.router.HandleFunc(route, handler).Methods(method)
}

func (c *RestService) GetRouteParam(req *http.Request, name string) string {
    return mux.Vars(req)[name]
}

func (c *RestService) GetQueryParam(req *http.Request, name string) string {
    return req.URL.Query().Get(name)
}

func (c *RestService) GetPagingParams(req *http.Request) *runtime.PagingParams {
    r := runtime.NewEmptyPagingParams()

    paging := req.URL.Query().Get("paging")
    if paging != "" { r.Paging = util.Converter.ToBoolean(paging) }

    skip := req.URL.Query().Get("skip")
    if skip != "" { r.Skip = util.Converter.ToNullableInteger(skip) }

    take := req.URL.Query().Get("take")
    if take != "" { r.Take = util.Converter.ToNullableInteger(take) }
    
    return r
}

func (c *RestService) GetFilterParams(req *http.Request) *runtime.FilterParams {
    r := runtime.NewEmptyFilterParams()
    
    for key, values := range req.URL.Query() {
        if key == "paging" || key == "skip" || key == "take" {
            // Skip paging param
        } else if len(values) == 0 { 
            // Skip empty param
        } else if len(values) == 1 {
            // Extract single value
            (*r)[key] = values[0]
        } else {
            // Convert multiple values into proper array
            array := []interface{} {}
            for _, value := range values {
                array = append(array, value)
            }
            (*r)[key] = array
        }
    }
    
    return r
}

func (c *RestService) GetInputData(req *http.Request, data interface{}) error {
    defer func() { req.Body.Close() }()

    err := json.NewDecoder(req.Body).Decode(data)
    if err != nil {
        return c.NewBadRequest("Failed to unmarshal request data").WithStatus(http.StatusBadRequest).WithCause(err)
    }
    
    return nil
}

func (c *RestService) SendError(w http.ResponseWriter, err error) {
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")

    merror, ok := err.(*runtime.MicroserviceError)
    if ok == false {
        merror = runtime.NewError(err.Error()).WithCause(err)
    }

    w.WriteHeader(merror.Status)
    json.NewEncoder(w).Encode(&merror) 
}

func (c *RestService) SendResult(w http.ResponseWriter, result interface{}, err error) {
    if err != nil {
        c.SendError(w, err)
        return
    }
        
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    if result != nil && reflect.ValueOf(result).IsNil() == false {
        w.WriteHeader(http.StatusOK)
        json.NewEncoder(w).Encode(result)
    } else {
        w.WriteHeader(http.StatusNotFound)
    }
}

func (c *RestService) SendCreatedResult(w http.ResponseWriter, result interface{}, err error) {
    if err != nil {
        c.SendError(w, err)
        return
    }
        
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    if result != nil {
        w.WriteHeader(http.StatusCreated)
        json.NewEncoder(w).Encode(result)
    } else {
        w.WriteHeader(http.StatusNotFound)
    }
}

func (c *RestService) SendDeletedResult(w http.ResponseWriter, err error) {
    if err != nil {
        c.SendError(w, err)
        return
    }
        
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusNoContent)
}
