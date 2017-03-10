package run

import (
    "fmt"
    "os"
    "os/signal"    
    "github.com/pip-services/pip-services-runtime-go/build"
)

type ProcessRunner struct {
    microservice *Microservice
    signalChannel chan os.Signal
    waitChannel chan bool
}

func NewProcessRunner(builder *build.Builder) *ProcessRunner {
    c := ProcessRunner { microservice: NewMicroservice(builder) }
    c.signalChannel = make(chan os.Signal, 1)
    c.waitChannel = make(chan bool)
    return &c
}

func (c *ProcessRunner) printErrorAndExit(err error) bool {
    if err != nil {
        c.microservice.Fatal(err)
        c.microservice.Info("Process is terminated")
        return true
    }
    return false
}

func (c *ProcessRunner) SetConfig(config interface{}) {
    c.microservice.SetConfig(config)
}

func (c *ProcessRunner) LoadConfig(configPath string) error {
    return c.microservice.LoadConfig(configPath)
}

func (c *ProcessRunner) LoadConfigWithDefault(defaultConfigPath string) error {
    configPath := defaultConfigPath
    if len(os.Args) > 1 { configPath = os.Args[1] }
    return c.microservice.LoadConfig(configPath)
}

func (c *ProcessRunner) captureExit() {
    c.microservice.Info("Press Control-C to stop the microservice...")
    
    signal.Notify(c.signalChannel, os.Interrupt)
    go func() {
        for _ = range c.signalChannel {
            c.microservice.Info("Goodbye!")
            c.waitChannel <- true
        }
    }()
    <- c.waitChannel
}

func (c *ProcessRunner) Run() (err error) {
    // Intercept panic errors
    defer func() {
        if p := recover(); p != nil {
            err = fmt.Errorf("Internal error: %v", p)
            c.printErrorAndExit(err) 
        }
    }()

    err = c.microservice.Start()
    if c.printErrorAndExit(err) { return err }
        
    c.captureExit()
    
    return c.microservice.Stop()
}

func (c *ProcessRunner) RunWithConfig(config interface{}) error {
    c.SetConfig(config)
    return c.Run()
}

func (c *ProcessRunner) RunWithConfigFile(configPath string) error {
    if err := c.LoadConfig(configPath); err != nil { return err }
    return c.Run()
}

func (c *ProcessRunner) RunWithDefaultConfigFile(defaultConfigPath string) error {
    if err := c.LoadConfigWithDefault(defaultConfigPath); err != nil { return err }
    return c.Run()
}
