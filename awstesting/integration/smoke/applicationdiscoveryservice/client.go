// +build integration

//Package applicationdiscoveryservice provides gucumber integration tests support.
package applicationdiscoveryservice

import (
	"github.com/stowelly/aws-sdk-go/aws"
	"github.com/stowelly/aws-sdk-go/awstesting/integration/smoke"
	"github.com/stowelly/aws-sdk-go/service/applicationdiscoveryservice"
	"github.com/gucumber/gucumber"
)

func init() {
	gucumber.Before("@applicationdiscoveryservice", func() {
		gucumber.World["client"] = applicationdiscoveryservice.New(
			smoke.Session, &aws.Config{Region: aws.String("us-west-2")},
		)
	})
}
