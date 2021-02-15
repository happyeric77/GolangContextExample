package main

import (
	"contextExample/withTimeout"
	"contextExample/withValue"
)

type Context interface {
	Demo()
}

func NewContext(kind string) Context {
	switch kind{
	case "withTimeout":
		return withTimeout.WithTimeoutContext{}
	case "withValue":
		return withValue.WithValueContext{}
	}
	return nil
}

func main() {
	// option: "withTimeout", "withValue"
	contextType := "withValue"
	Context := NewContext(contextType)
	Context.Demo()
}
