// +build integration

//Package ssm provides gucumber integration tests support.
package ssm

import (
	"github.com/stowelly/aws-sdk-go/awstesting/integration/smoke"
	"github.com/stowelly/aws-sdk-go/service/ssm"
	"github.com/gucumber/gucumber"
)

func init() {
	gucumber.Before("@ssm", func() {
		gucumber.World["client"] = ssm.New(smoke.Session)
	})
}
