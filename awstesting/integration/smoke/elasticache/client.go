// +build integration

//Package elasticache provides gucumber integration tests support.
package elasticache

import (
	"github.com/stowelly/aws-sdk-go/awstesting/integration/smoke"
	"github.com/stowelly/aws-sdk-go/service/elasticache"
	"github.com/gucumber/gucumber"
)

func init() {
	gucumber.Before("@elasticache", func() {
		gucumber.World["client"] = elasticache.New(smoke.Session)
	})
}
