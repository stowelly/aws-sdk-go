// +build integration

//Package cloudwatch provides gucumber integration tests support.
package cloudwatch

import (
	"github.com/stowelly/aws-sdk-go/awstesting/integration/smoke"
	"github.com/stowelly/aws-sdk-go/service/cloudwatch"
	"github.com/gucumber/gucumber"
)

func init() {
	gucumber.Before("@cloudwatch", func() {
		gucumber.World["client"] = cloudwatch.New(smoke.Session)
	})
}
