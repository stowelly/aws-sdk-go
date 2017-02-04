// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package supportiface provides an interface to enable mocking the AWS Support service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package supportiface

import (
	"github.com/stowelly/aws-sdk-go/aws/request"
	"github.com/stowelly/aws-sdk-go/service/support"
)

// SupportAPI provides an interface to enable mocking the
// support.Support service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // AWS Support.
//    func myFunc(svc supportiface.SupportAPI) bool {
//        // Make svc.AddAttachmentsToSet request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := support.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockSupportClient struct {
//        supportiface.SupportAPI
//    }
//    func (m *mockSupportClient) AddAttachmentsToSet(input *support.AddAttachmentsToSetInput) (*support.AddAttachmentsToSetOutput, error) {
//        // mock response/functionality
//    }
//
//    TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockSupportClient{}
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
type SupportAPI interface {
	AddAttachmentsToSetRequest(*support.AddAttachmentsToSetInput) (*request.Request, *support.AddAttachmentsToSetOutput)

	AddAttachmentsToSet(*support.AddAttachmentsToSetInput) (*support.AddAttachmentsToSetOutput, error)

	AddCommunicationToCaseRequest(*support.AddCommunicationToCaseInput) (*request.Request, *support.AddCommunicationToCaseOutput)

	AddCommunicationToCase(*support.AddCommunicationToCaseInput) (*support.AddCommunicationToCaseOutput, error)

	CreateCaseRequest(*support.CreateCaseInput) (*request.Request, *support.CreateCaseOutput)

	CreateCase(*support.CreateCaseInput) (*support.CreateCaseOutput, error)

	DescribeAttachmentRequest(*support.DescribeAttachmentInput) (*request.Request, *support.DescribeAttachmentOutput)

	DescribeAttachment(*support.DescribeAttachmentInput) (*support.DescribeAttachmentOutput, error)

	DescribeCasesRequest(*support.DescribeCasesInput) (*request.Request, *support.DescribeCasesOutput)

	DescribeCases(*support.DescribeCasesInput) (*support.DescribeCasesOutput, error)

	DescribeCasesPages(*support.DescribeCasesInput, func(*support.DescribeCasesOutput, bool) bool) error

	DescribeCommunicationsRequest(*support.DescribeCommunicationsInput) (*request.Request, *support.DescribeCommunicationsOutput)

	DescribeCommunications(*support.DescribeCommunicationsInput) (*support.DescribeCommunicationsOutput, error)

	DescribeCommunicationsPages(*support.DescribeCommunicationsInput, func(*support.DescribeCommunicationsOutput, bool) bool) error

	DescribeServicesRequest(*support.DescribeServicesInput) (*request.Request, *support.DescribeServicesOutput)

	DescribeServices(*support.DescribeServicesInput) (*support.DescribeServicesOutput, error)

	DescribeSeverityLevelsRequest(*support.DescribeSeverityLevelsInput) (*request.Request, *support.DescribeSeverityLevelsOutput)

	DescribeSeverityLevels(*support.DescribeSeverityLevelsInput) (*support.DescribeSeverityLevelsOutput, error)

	DescribeTrustedAdvisorCheckRefreshStatusesRequest(*support.DescribeTrustedAdvisorCheckRefreshStatusesInput) (*request.Request, *support.DescribeTrustedAdvisorCheckRefreshStatusesOutput)

	DescribeTrustedAdvisorCheckRefreshStatuses(*support.DescribeTrustedAdvisorCheckRefreshStatusesInput) (*support.DescribeTrustedAdvisorCheckRefreshStatusesOutput, error)

	DescribeTrustedAdvisorCheckResultRequest(*support.DescribeTrustedAdvisorCheckResultInput) (*request.Request, *support.DescribeTrustedAdvisorCheckResultOutput)

	DescribeTrustedAdvisorCheckResult(*support.DescribeTrustedAdvisorCheckResultInput) (*support.DescribeTrustedAdvisorCheckResultOutput, error)

	DescribeTrustedAdvisorCheckSummariesRequest(*support.DescribeTrustedAdvisorCheckSummariesInput) (*request.Request, *support.DescribeTrustedAdvisorCheckSummariesOutput)

	DescribeTrustedAdvisorCheckSummaries(*support.DescribeTrustedAdvisorCheckSummariesInput) (*support.DescribeTrustedAdvisorCheckSummariesOutput, error)

	DescribeTrustedAdvisorChecksRequest(*support.DescribeTrustedAdvisorChecksInput) (*request.Request, *support.DescribeTrustedAdvisorChecksOutput)

	DescribeTrustedAdvisorChecks(*support.DescribeTrustedAdvisorChecksInput) (*support.DescribeTrustedAdvisorChecksOutput, error)

	RefreshTrustedAdvisorCheckRequest(*support.RefreshTrustedAdvisorCheckInput) (*request.Request, *support.RefreshTrustedAdvisorCheckOutput)

	RefreshTrustedAdvisorCheck(*support.RefreshTrustedAdvisorCheckInput) (*support.RefreshTrustedAdvisorCheckOutput, error)

	ResolveCaseRequest(*support.ResolveCaseInput) (*request.Request, *support.ResolveCaseOutput)

	ResolveCase(*support.ResolveCaseInput) (*support.ResolveCaseOutput, error)
}

var _ SupportAPI = (*support.Support)(nil)
