package main

import (
	"go.temporal.io/sdk/worker"
)

// == Workflow ==

func TdwsRegister(w worker.Worker) {
	w.RegisterWorkflow(GreetingWorkflow)
	w.RegisterWorkflow(GoodbyeWorkflow)
	w.RegisterActivity(ComposeGreeting)
	w.RegisterActivity(ComposeGoodbye)
}
