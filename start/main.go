package main

import (
	"context"
	app "testapp"

	"go.temporal.io/sdk/client"
)

func main() {
	c, err := client.Dial(client.Options{})
	if err != nil {
		panic(err)
	}
	defer c.Close()

	options := client.StartWorkflowOptions{
		ID: "test-workflow",
		TaskQueue: app.TaskQueue,
	}

	_, err = c.ExecuteWorkflow(context.Background(), options, app.SignalTestWorkflow)
	if err != nil {
		panic(err)
	}
}
