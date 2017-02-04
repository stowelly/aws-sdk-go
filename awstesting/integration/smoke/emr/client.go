// +build integration

//Package emr provides gucumber integration tests support.
package emr

import (
	"github.com/stowelly/aws-sdk-go/awstesting/integration/smoke"
	"github.com/stowelly/aws-sdk-go/service/emr"
	"github.com/gucumber/gucumber"
)

func init() {
	gucumber.Before("@emr", func() {
		gucumber.World["client"] = emr.New(smoke.Session)
	})
}
