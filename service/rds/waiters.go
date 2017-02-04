// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package rds

import (
	"github.com/stowelly/aws-sdk-go/private/waiter"
)

// WaitUntilDBInstanceAvailable uses the Amazon RDS API operation
// DescribeDBInstances to wait for a condition to be met before returning.
// If the condition is not meet within the max attempt window an error will
// be returned.
func (c *RDS) WaitUntilDBInstanceAvailable(input *DescribeDBInstancesInput) error {
	waiterCfg := waiter.Config{
		Operation:   "DescribeDBInstances",
		Delay:       30,
		MaxAttempts: 60,
		Acceptors: []waiter.WaitAcceptor{
			{
				State:    "success",
				Matcher:  "pathAll",
				Argument: "DBInstances[].DBInstanceStatus",
				Expected: "available",
			},
			{
				State:    "failure",
				Matcher:  "pathAny",
				Argument: "DBInstances[].DBInstanceStatus",
				Expected: "deleted",
			},
			{
				State:    "failure",
				Matcher:  "pathAny",
				Argument: "DBInstances[].DBInstanceStatus",
				Expected: "deleting",
			},
			{
				State:    "failure",
				Matcher:  "pathAny",
				Argument: "DBInstances[].DBInstanceStatus",
				Expected: "failed",
			},
			{
				State:    "failure",
				Matcher:  "pathAny",
				Argument: "DBInstances[].DBInstanceStatus",
				Expected: "incompatible-restore",
			},
			{
				State:    "failure",
				Matcher:  "pathAny",
				Argument: "DBInstances[].DBInstanceStatus",
				Expected: "incompatible-parameters",
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

// WaitUntilDBInstanceDeleted uses the Amazon RDS API operation
// DescribeDBInstances to wait for a condition to be met before returning.
// If the condition is not meet within the max attempt window an error will
// be returned.
func (c *RDS) WaitUntilDBInstanceDeleted(input *DescribeDBInstancesInput) error {
	waiterCfg := waiter.Config{
		Operation:   "DescribeDBInstances",
		Delay:       30,
		MaxAttempts: 60,
		Acceptors: []waiter.WaitAcceptor{
			{
				State:    "success",
				Matcher:  "pathAll",
				Argument: "DBInstances[].DBInstanceStatus",
				Expected: "deleted",
			},
			{
				State:    "success",
				Matcher:  "error",
				Argument: "",
				Expected: "DBInstanceNotFound",
			},
			{
				State:    "failure",
				Matcher:  "pathAny",
				Argument: "DBInstances[].DBInstanceStatus",
				Expected: "creating",
			},
			{
				State:    "failure",
				Matcher:  "pathAny",
				Argument: "DBInstances[].DBInstanceStatus",
				Expected: "modifying",
			},
			{
				State:    "failure",
				Matcher:  "pathAny",
				Argument: "DBInstances[].DBInstanceStatus",
				Expected: "rebooting",
			},
			{
				State:    "failure",
				Matcher:  "pathAny",
				Argument: "DBInstances[].DBInstanceStatus",
				Expected: "resetting-master-credentials",
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
