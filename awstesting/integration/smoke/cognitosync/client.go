// +build integration

//Package cognitosync provides gucumber integration tests support.
package cognitosync

import (
	"github.com/stowelly/aws-sdk-go/awstesting/integration/smoke"
	"github.com/stowelly/aws-sdk-go/service/cognitosync"
	"github.com/gucumber/gucumber"
)

func init() {
	gucumber.Before("@cognitosync", func() {
		gucumber.World["client"] = cognitosync.New(smoke.Session)
	})
}
