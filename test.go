package main

import (
	"context"
	"fmt"
	"go.temporal.io/sdk/client"
	"log"
)

//func GreetingWorkflow(ctx workflow.Context, name string) (string, error) {
//	options := workflow.ActivityOptions{
//		StartToCloseTimeout: time.Second * 5,
//	}
//
//	ctx = workflow.WithActivityOptions(ctx, options)
//
//	fmt.Printf("Execute Activity - ComposeGreeting: %s, \n", name)
//	var result string
//	err := workflow.ExecuteActivity(ctx, ComposeGreeting, name).Get(ctx, &result)
//	if err != nil {
//		return result, fmt.Errorf("ComposeGreeting failed: %w", err)
//	}
//
//	return result, err
//}
//
//func ComposeGreeting(ctx context.Context, name string) (string, error) {
//	greeting := fmt.Sprintf("Hello %s!", name)
//	fmt.Printf("Execute ComposeGreeting: %s, \n", greeting)
//	return greeting, nil
//}

func main() {

	// Create the client object just once per process
	c, err := client.Dial(client.Options{
		HostPort:  "192.168.1.233:7233", // Default for local Temporal server
		Namespace: "network-usecases",   // Specify the namespace
	})
	if err != nil {
		log.Fatalln("unable to create Temporal client", err)
	}
	defer c.Close()

	options := client.StartWorkflowOptions{
		ID:        "greeting-workflow",
		TaskQueue: "tdws-demo",
	}

	// Start the Workflow
	name := "World"
	we, err := c.ExecuteWorkflow(context.Background(), options, GreetingWorkflow, name)
	if err != nil {
		log.Fatalln("unable to complete Workflow", err)
	}

	// Get the results
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
