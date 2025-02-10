package components

import (
	"context"
	"fmt"
	"go.temporal.io/sdk/workflow"
	"time"
)

func GreetingWorkflow(ctx workflow.Context, name string) (string, error) {
	options := workflow.ActivityOptions{
		StartToCloseTimeout: time.Second * 5,
	}

	ctx = workflow.WithActivityOptions(ctx, options)

	var result string
	err := workflow.ExecuteActivity(ctx, ComposeGreeting, name).Get(ctx, &result)

	return result, err
}
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

func ComposeGreeting(ctx context.Context, name string) (string, error) {
	greeting := fmt.Sprintf("Hello %s!", name)
	return greeting, nil
}
