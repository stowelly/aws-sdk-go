// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package appstream

import (
	"github.com/stowelly/aws-sdk-go/private/waiter"
)

// WaitUntilFleetStarted uses the Amazon AppStream API operation
// DescribeFleets to wait for a condition to be met before returning.
// If the condition is not meet within the max attempt window an error will
// be returned.
func (c *AppStream) WaitUntilFleetStarted(input *DescribeFleetsInput) error {
	waiterCfg := waiter.Config{
		Operation:   "DescribeFleets",
		Delay:       30,
		MaxAttempts: 40,
		Acceptors: []waiter.WaitAcceptor{
			{
				State:    "success",
				Matcher:  "pathAll",
				Argument: "fleets[].state",
				Expected: "ACTIVE",
			},
			{
				State:    "failure",
				Matcher:  "pathAny",
				Argument: "fleets[].state",
				Expected: "PENDING_DEACTIVATE",
			},
			{
				State:    "failure",
				Matcher:  "pathAny",
				Argument: "fleets[].state",
				Expected: "INACTIVE",
			},
		},
	}

	w := waiter.Waiter{
		Client: c,
		Input:  input,
		Config: waiterCfg,
	}
	return w.Wait()
}

// WaitUntilFleetStopped uses the Amazon AppStream API operation
// DescribeFleets to wait for a condition to be met before returning.
// If the condition is not meet within the max attempt window an error will
// be returned.
func (c *AppStream) WaitUntilFleetStopped(input *DescribeFleetsInput) error {
	waiterCfg := waiter.Config{
		Operation:   "DescribeFleets",
		Delay:       30,
		MaxAttempts: 40,
		Acceptors: []waiter.WaitAcceptor{
			{
				State:    "success",
				Matcher:  "pathAll",
				Argument: "fleets[].state",
				Expected: "INACTIVE",
			},
			{
				State:    "failure",
				Matcher:  "pathAny",
				Argument: "fleets[].state",
				Expected: "PENDING_ACTIVATE",
			},
			{
				State:    "failure",
				Matcher:  "pathAny",
				Argument: "fleets[].state",
				Expected: "ACTIVE",
			},
		},
	}

	w := waiter.Waiter{
		Client: c,
		Input:  input,
		Config: waiterCfg,
	}
	return w.Wait()
}
