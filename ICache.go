package runtime

type ICache interface{
    IComponent
    
    Get(key string, value interface{}) (interface{}, error)
    Set(key string, value interface{}) (interface{}, error)
    Unset(key string) error
}