package main

import (
	"context"
	"fmt"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
	"log"
	"time"
)

func GreetingWorkflow(ctx workflow.Context, name string) (string, error) {
	// Increased timeout and added better retry configuration
	options := workflow.ActivityOptions{
		StartToCloseTimeout:    time.Second * 30, // Increased from 5 seconds to 30 seconds
		ScheduleToCloseTimeout: time.Minute * 5,  // Added overall timeout
		RetryPolicy: &temporal.RetryPolicy{
			InitialInterval:    time.Second,
			BackoffCoefficient: 2.0,
			MaximumInterval:    time.Second * 100,
			MaximumAttempts:    3,
		},
	}

	ctx = workflow.WithActivityOptions(ctx, options)

	var result string
	logger := workflow.GetLogger(ctx)
	logger.Info("Starting ComposeGreeting activity", "name", name)

	err := workflow.ExecuteActivity(ctx, ComposeGreeting, name).Get(ctx, &result)
	if err != nil {
		logger.Error("ComposeGreeting activity failed", "error", err)
		return "", fmt.Errorf("ComposeGreeting failed: %w", err)
	}

	return result, nil
}

func ComposeGreeting(ctx context.Context, name string) (string, error) {
	// Using simple logging for the activity since we're in a non-workflow context
	greeting := fmt.Sprintf("Hello %s!", name)
	log.Printf("Composing greeting: %s", greeting)
	return greeting, nil
}

func main() {
	c, err := client.Dial(client.Options{
		HostPort:  "192.168.1.233:7233",
		Namespace: "network-usecases",
	})
	if err != nil {
		log.Fatalln("unable to create Temporal client", err)
	}
	defer c.Close()

	options := client.StartWorkflowOptions{
		ID:        "greeting-workflow",
		TaskQueue: "tdws-demo",
		// Added workflow timeout
		WorkflowExecutionTimeout: time.Minute * 10,
	}

	//// Create a worker
	//w := worker.New(c, "tdws-demo", worker.Options{})
	//
	//// Register workflow and activity
	//w.RegisterWorkflow(GreetingWorkflow)
	//w.RegisterActivity(ComposeGreeting)
	//
	//// Start listening to the Task Queue
	//err = w.Run(worker.InterruptCh())
	//if err != nil {
	//	log.Fatalln("unable to start Worker", err)
	//}

	name := "World"
	we, err := c.ExecuteWorkflow(context.Background(), options, GreetingWorkflow, name)
	if err != nil {
		log.Fatalln("unable to start Workflow", err)
	}

	var greeting string
	err = we.Get(context.Background(), &greeting)
	if err != nil {
		log.Fatalln("unable to get Workflow result", err)
	}

	printResults(greeting, we.GetID(), we.GetRunID())
}

func printResults(greeting string, workflowID, runID string) {
	fmt.Printf("\nWorkflowID: %s RunID: %s\n", workflowID, runID)
	fmt.Printf("\n%s\n\n", greeting)
}
