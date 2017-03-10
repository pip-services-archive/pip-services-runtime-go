package build

import (
    "io"
    "os"
    "github.com/pip-services/pip-services-runtime-go"
    "github.com/pip-services/pip-services-runtime-go/log"
)

type TLifeCycleManager struct {}

var LifeCycleManager *TLifeCycleManager = &TLifeCycleManager{}

func (c *TLifeCycleManager) Fatal(refs *runtime.References, message ...interface{}) {
    if refs != nil && refs.Log != nil {
        refs.Log.Fatal(message)
    } else {
        io.WriteString(os.Stderr, log.FormatConsoleMessage("FATAL", message))
    }
}

func (c *TLifeCycleManager) Error(refs *runtime.References, message ...interface{}) {
    if refs != nil && refs.Log != nil {
        refs.Log.Error(message)
    } else {
        io.WriteString(os.Stderr, log.FormatConsoleMessage("ERROR", message))
    }
}

func (c *TLifeCycleManager) Warn(refs *runtime.References, message ...interface{}) {
    if refs != nil && refs.Log != nil {
        refs.Log.Warn(message)
    } else {
        io.WriteString(os.Stdout, log.FormatConsoleMessage("WARN", message))
    }
}

func (c *TLifeCycleManager) Info(refs *runtime.References, message ...interface{}) {
    if refs != nil && refs.Log != nil {
        refs.Log.Info(message)
    } else {
        io.WriteString(os.Stdout, log.FormatConsoleMessage("INFO", message))
    }
}

func (c *TLifeCycleManager) Debug(refs *runtime.References, message ...interface{}) {
    if refs != nil && refs.Log != nil {
        refs.Log.Debug(message)
    } else {
        io.WriteString(os.Stdout, log.FormatConsoleMessage("DEBUG", message))
    }
}

func (c *TLifeCycleManager) Trace(refs *runtime.References, message ...interface{}) {
    if refs != nil && refs.Log != nil {
        refs.Log.Trace(message)
    } else {
        io.WriteString(os.Stdout, log.FormatConsoleMessage("TRACE", message))
    }
}

// We need this check since IComponent is identical to IController
func isLogic(component runtime.IComponent, refs *runtime.References) bool {
    if component == refs.Ctrl { return true }
    
    for _, ints := range refs.Ints {
        if component == ints { return true }
    }
    
    return false
}

func (c *TLifeCycleManager) InitComponents(refs *runtime.References, components []runtime.IComponent) error {
    if len(components) == 0 { return nil }
    
    for _, component := range components {
        err := component.Init(refs)
        if err != nil { return err }
        
        // Assign logic reference
        logic, ok := component.(runtime.IController)
        if ok && isLogic(component, refs) { refs.WithLogic(logic) }
    }
    
    return nil
}

func (c *TLifeCycleManager) Init(refs *runtime.References) error {
    components := refs.GetComponents()
    return c.InitComponents(refs, components)
}

func (c *TLifeCycleManager) InitAndOpenComponents(refs *runtime.References, components []runtime.IComponent) error {
    if len(components) == 0 { return nil }

    err := c.InitComponents(refs, components)
    if err != nil { return err }
    
    return c.OpenComponents(refs, components)    
}

func (c *TLifeCycleManager) InitAndOpen(refs *runtime.References) error {
    components := refs.GetComponents()
    return c.InitAndOpenComponents(refs, components)
}

func (c *TLifeCycleManager) OpenComponents(refs *runtime.References, components []runtime.IComponent) error {
    if len(components) == 0 { return nil }
    
    var err error = nil
    
    for _, component := range components {
        err = component.Open()
        if err != nil { break }
    }
    
    if err != nil {
        c.Trace(refs, "Microservice opening failed with error " + err.Error()) 
        c.ForceCloseComponents(refs, components) 
    }
    
    return err
}

func (c *TLifeCycleManager) Open(refs *runtime.References) error {
    components := refs.GetComponents()
    return c.OpenComponents(refs, components)
}

func (c *TLifeCycleManager) CloseComponents(refs *runtime.References, components []runtime.IComponent) error {
    if len(components) == 0 { return nil }
        
    // Process components in the reverse order
    for i := len(components) - 1; i >= 0; i-- {
        err := components[i].Close()
        
        if err != nil {
            c.Trace(refs, "Microservice closure failed with error " + err.Error()) 
            return err 
        }
    }
    
    return nil
}

func (c *TLifeCycleManager) Close(refs *runtime.References) error {
    components := refs.GetComponents()
    return c.CloseComponents(refs, components)
}

func (c *TLifeCycleManager) ForceCloseComponents(refs *runtime.References, components []runtime.IComponent) error {
    if len(components) == 0 { return nil }
    
    var firstError error
    
    // Process components in the reverse order
    for i := len(components) - 1; i >= 0; i-- {
        err := components[i].Close()
        
        if err != nil && firstError == nil { 
            c.Trace(refs, "Microservice closure failed with error " + err.Error())
            firstError = err 
        }
    }
    
    return firstError
}

func (c *TLifeCycleManager) ForceClose(refs *runtime.References) error {
    components := refs.GetComponents()
    return c.ForceCloseComponents(refs, components)
}
