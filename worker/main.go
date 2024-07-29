package main

import (
	app "testapp"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func main() {
	c, err := client.Dial(client.Options{})
	if err != nil {
		panic(err)
	}
	defer c.Close()

	w := worker.New(c, app.TaskQueue, worker.Options{})
	w.RegisterWorkflow(app.SignalTestWorkflow)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		panic(err)
	}
}
