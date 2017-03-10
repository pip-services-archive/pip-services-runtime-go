package runtime

type IConfig interface{
    IComponent
    
    Read() (*DynamicMap, error)
}