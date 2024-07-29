package app

import (
	"go.temporal.io/sdk/workflow"
)

type TestSignalType struct {
	Action string
}

const TestSignalChannel = "test-signal-channel"

func SignalTestWorkflow(ctx workflow.Context) error {
	sigChan := workflow.GetSignalChannel(ctx, TestSignalChannel)
	selector := workflow.NewSelector(ctx)

	done := false
	selector.AddReceive(sigChan, func(c workflow.ReceiveChannel, more bool) {
		var sig TestSignalType
		c.Receive(ctx, &sig)

		if sig.Action == "stop" {
			done = true
		}
	})

	for !done {
		selector.Select(ctx)
	}

	return nil
}
