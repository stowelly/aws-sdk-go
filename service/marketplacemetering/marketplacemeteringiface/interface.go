// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package marketplacemeteringiface provides an interface to enable mocking the AWSMarketplace Metering service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package marketplacemeteringiface

import (
	"github.com/stowelly/aws-sdk-go/aws/request"
	"github.com/stowelly/aws-sdk-go/service/marketplacemetering"
)

// MarketplaceMeteringAPI provides an interface to enable mocking the
// marketplacemetering.MarketplaceMetering service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // AWSMarketplace Metering.
//    func myFunc(svc marketplacemeteringiface.MarketplaceMeteringAPI) bool {
//        // Make svc.BatchMeterUsage request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := marketplacemetering.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockMarketplaceMeteringClient struct {
//        marketplacemeteringiface.MarketplaceMeteringAPI
//    }
//    func (m *mockMarketplaceMeteringClient) BatchMeterUsage(input *marketplacemetering.BatchMeterUsageInput) (*marketplacemetering.BatchMeterUsageOutput, error) {
//        // mock response/functionality
//    }
//
//    TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockMarketplaceMeteringClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type MarketplaceMeteringAPI interface {
	BatchMeterUsageRequest(*marketplacemetering.BatchMeterUsageInput) (*request.Request, *marketplacemetering.BatchMeterUsageOutput)

	BatchMeterUsage(*marketplacemetering.BatchMeterUsageInput) (*marketplacemetering.BatchMeterUsageOutput, error)

	MeterUsageRequest(*marketplacemetering.MeterUsageInput) (*request.Request, *marketplacemetering.MeterUsageOutput)

	MeterUsage(*marketplacemetering.MeterUsageInput) (*marketplacemetering.MeterUsageOutput, error)

	ResolveCustomerRequest(*marketplacemetering.ResolveCustomerInput) (*request.Request, *marketplacemetering.ResolveCustomerOutput)

	ResolveCustomer(*marketplacemetering.ResolveCustomerInput) (*marketplacemetering.ResolveCustomerOutput, error)
}

var _ MarketplaceMeteringAPI = (*marketplacemetering.MarketplaceMetering)(nil)
