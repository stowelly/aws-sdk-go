// +build integration

//Package dynamodb provides gucumber integration tests support.
package dynamodb

import (
	"github.com/stowelly/aws-sdk-go/awstesting/integration/smoke"
	"github.com/stowelly/aws-sdk-go/service/dynamodb"
	"github.com/gucumber/gucumber"
)

func init() {
	gucumber.Before("@dynamodb", func() {
		gucumber.World["client"] = dynamodb.New(smoke.Session)
	})
}
