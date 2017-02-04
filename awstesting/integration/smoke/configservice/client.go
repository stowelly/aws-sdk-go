// +build integration

//Package configservice provides gucumber integration tests support.
package configservice

import (
	"github.com/stowelly/aws-sdk-go/awstesting/integration/smoke"
	"github.com/stowelly/aws-sdk-go/service/configservice"
	"github.com/gucumber/gucumber"
)

func init() {
	gucumber.Before("@configservice", func() {
		gucumber.World["client"] = configservice.New(smoke.Session)
	})
}
