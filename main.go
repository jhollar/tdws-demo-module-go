package main

import (
	"go.temporal.io/sdk/worker"
	"tdws-demo-module-go/components"
)

// == Workflow ==

func TdwsRegister(w worker.Worker) {
	w.RegisterWorkflow(components.GreetingWorkflow)
	w.RegisterWorkflow(components.GoodbyeWorkflow)
	w.RegisterActivity(components.ComposeGreeting)
	w.RegisterActivity(components.ComposeGoodbye)
}
