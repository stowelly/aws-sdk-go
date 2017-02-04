// +build integration

//Package glacier provides gucumber integration tests support.
package glacier

import (
	"github.com/stowelly/aws-sdk-go/awstesting/integration/smoke"
	"github.com/stowelly/aws-sdk-go/service/glacier"
	"github.com/gucumber/gucumber"
)

func init() {
	gucumber.Before("@glacier", func() {
		gucumber.World["client"] = glacier.New(smoke.Session)
	})
}
