package main

import (
    "github.com/pip-services/pip-services-runtime-go/run"
    "github.com/pip-services/pip-services-runtime-go/test/build"        
)

type DummyProcessRunner struct {
    run.ProcessRunner    
}

func NewDummyProcessRunner() *DummyProcessRunner {
    builder := build.NewDummyBuilder()
    return &DummyProcessRunner { ProcessRunner: *run.NewProcessRunner(builder) }
}

func main() {
    runner := NewDummyProcessRunner()
    runner.RunWithDefaultConfigFile("config.json")
}