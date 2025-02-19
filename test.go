package main

import (
	"context"
	"fmt"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/workflow"
	"log"
	"time"
)

func ComposeGreeting(ctx context.Context, name string) (string, error) {
	greeting := fmt.Sprintf("Hello %s!", name)
	return greeting, nil
}

func GreetingWorkflow(ctx workflow.Context, name string) (string, error) {
	//options := workflow.ActivityOptions{
	//	StartToCloseTimeout: time.Second * 5,
	//}
	options := workflow.ActivityOptions{}

	ctx = workflow.WithActivityOptions(ctx, options)

	var result string
	err := workflow.ExecuteActivity(ctx, ComposeGreeting, name).Get(ctx, &result)

	return result, err
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

	// Start the Workflow
	name := "World"
	we, err := c.ExecuteWorkflow(context.Background(), options, GreetingWorkflow, name)
	if err != nil {
		log.Fatalln("unable to start Workflow", err)
	}

	// Get the Results
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
