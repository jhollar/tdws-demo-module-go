package main

import (
	"context"
	"fmt"
	"go.temporal.io/sdk/worker"
	"go.temporal.io/sdk/workflow"
	"time"
)

// == Workflow ==

//func GreetingWorkflow(ctx workflow.Context, name string) (string, error) {
//	options := workflow.ActivityOptions{
//		StartToCloseTimeout: time.Second * 5,
//	}
//
//	ctx = workflow.WithActivityOptions(ctx, options)
//
//	var result string
//	err := workflow.ExecuteActivity(ctx, ComposeGreeting, name).Get(ctx, &result)
//
//	return result, err
//}

func GoodbyeWorkflow(ctx workflow.Context, name string) (string, error) {
	options := workflow.ActivityOptions{
		StartToCloseTimeout: time.Second * 5,
	}

	ctx = workflow.WithActivityOptions(ctx, options)

	var result string
	err := workflow.ExecuteActivity(ctx, ComposeGoodbye, name).Get(ctx, &result)

	return result, err
}

// == Activity ==
func ComposeGoodbye(ctx context.Context, name string) (string, error) {
	greeting := fmt.Sprintf("Goodbye %s!", name)
	return greeting, nil
}

//func ComposeGreeting(ctx context.Context, name string) (string, error) {
//	greeting := fmt.Sprintf("Hello %s!", name)
//	return greeting, nil
//}

func TdwsRegister(w worker.Worker) {

	fmt.Println("TdwsRegister Invoked")

	w.RegisterWorkflow(GreetingWorkflow)
	w.RegisterWorkflow(GoodbyeWorkflow)
	w.RegisterActivity(ComposeGreeting)
	w.RegisterActivity(ComposeGoodbye)

	fmt.Println("TdwsRegister Return")
}
