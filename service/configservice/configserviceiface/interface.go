// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package configserviceiface provides an interface to enable mocking the AWS Config service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package configserviceiface

import (
	"github.com/stowelly/aws-sdk-go/aws/request"
	"github.com/stowelly/aws-sdk-go/service/configservice"
)

// ConfigServiceAPI provides an interface to enable mocking the
// configservice.ConfigService service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // AWS Config.
//    func myFunc(svc configserviceiface.ConfigServiceAPI) bool {
//        // Make svc.DeleteConfigRule request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := configservice.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockConfigServiceClient struct {
//        configserviceiface.ConfigServiceAPI
//    }
//    func (m *mockConfigServiceClient) DeleteConfigRule(input *configservice.DeleteConfigRuleInput) (*configservice.DeleteConfigRuleOutput, error) {
//        // mock response/functionality
//    }
//
//    TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockConfigServiceClient{}
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
type ConfigServiceAPI interface {
	DeleteConfigRuleRequest(*configservice.DeleteConfigRuleInput) (*request.Request, *configservice.DeleteConfigRuleOutput)

	DeleteConfigRule(*configservice.DeleteConfigRuleInput) (*configservice.DeleteConfigRuleOutput, error)

	DeleteConfigurationRecorderRequest(*configservice.DeleteConfigurationRecorderInput) (*request.Request, *configservice.DeleteConfigurationRecorderOutput)

	DeleteConfigurationRecorder(*configservice.DeleteConfigurationRecorderInput) (*configservice.DeleteConfigurationRecorderOutput, error)

	DeleteDeliveryChannelRequest(*configservice.DeleteDeliveryChannelInput) (*request.Request, *configservice.DeleteDeliveryChannelOutput)

	DeleteDeliveryChannel(*configservice.DeleteDeliveryChannelInput) (*configservice.DeleteDeliveryChannelOutput, error)

	DeleteEvaluationResultsRequest(*configservice.DeleteEvaluationResultsInput) (*request.Request, *configservice.DeleteEvaluationResultsOutput)

	DeleteEvaluationResults(*configservice.DeleteEvaluationResultsInput) (*configservice.DeleteEvaluationResultsOutput, error)

	DeliverConfigSnapshotRequest(*configservice.DeliverConfigSnapshotInput) (*request.Request, *configservice.DeliverConfigSnapshotOutput)

	DeliverConfigSnapshot(*configservice.DeliverConfigSnapshotInput) (*configservice.DeliverConfigSnapshotOutput, error)

	DescribeComplianceByConfigRuleRequest(*configservice.DescribeComplianceByConfigRuleInput) (*request.Request, *configservice.DescribeComplianceByConfigRuleOutput)

	DescribeComplianceByConfigRule(*configservice.DescribeComplianceByConfigRuleInput) (*configservice.DescribeComplianceByConfigRuleOutput, error)

	DescribeComplianceByResourceRequest(*configservice.DescribeComplianceByResourceInput) (*request.Request, *configservice.DescribeComplianceByResourceOutput)

	DescribeComplianceByResource(*configservice.DescribeComplianceByResourceInput) (*configservice.DescribeComplianceByResourceOutput, error)

	DescribeConfigRuleEvaluationStatusRequest(*configservice.DescribeConfigRuleEvaluationStatusInput) (*request.Request, *configservice.DescribeConfigRuleEvaluationStatusOutput)

	DescribeConfigRuleEvaluationStatus(*configservice.DescribeConfigRuleEvaluationStatusInput) (*configservice.DescribeConfigRuleEvaluationStatusOutput, error)

	DescribeConfigRulesRequest(*configservice.DescribeConfigRulesInput) (*request.Request, *configservice.DescribeConfigRulesOutput)

	DescribeConfigRules(*configservice.DescribeConfigRulesInput) (*configservice.DescribeConfigRulesOutput, error)

	DescribeConfigurationRecorderStatusRequest(*configservice.DescribeConfigurationRecorderStatusInput) (*request.Request, *configservice.DescribeConfigurationRecorderStatusOutput)

	DescribeConfigurationRecorderStatus(*configservice.DescribeConfigurationRecorderStatusInput) (*configservice.DescribeConfigurationRecorderStatusOutput, error)

	DescribeConfigurationRecordersRequest(*configservice.DescribeConfigurationRecordersInput) (*request.Request, *configservice.DescribeConfigurationRecordersOutput)

	DescribeConfigurationRecorders(*configservice.DescribeConfigurationRecordersInput) (*configservice.DescribeConfigurationRecordersOutput, error)

	DescribeDeliveryChannelStatusRequest(*configservice.DescribeDeliveryChannelStatusInput) (*request.Request, *configservice.DescribeDeliveryChannelStatusOutput)

	DescribeDeliveryChannelStatus(*configservice.DescribeDeliveryChannelStatusInput) (*configservice.DescribeDeliveryChannelStatusOutput, error)

	DescribeDeliveryChannelsRequest(*configservice.DescribeDeliveryChannelsInput) (*request.Request, *configservice.DescribeDeliveryChannelsOutput)

	DescribeDeliveryChannels(*configservice.DescribeDeliveryChannelsInput) (*configservice.DescribeDeliveryChannelsOutput, error)

	GetComplianceDetailsByConfigRuleRequest(*configservice.GetComplianceDetailsByConfigRuleInput) (*request.Request, *configservice.GetComplianceDetailsByConfigRuleOutput)

	GetComplianceDetailsByConfigRule(*configservice.GetComplianceDetailsByConfigRuleInput) (*configservice.GetComplianceDetailsByConfigRuleOutput, error)

	GetComplianceDetailsByResourceRequest(*configservice.GetComplianceDetailsByResourceInput) (*request.Request, *configservice.GetComplianceDetailsByResourceOutput)

	GetComplianceDetailsByResource(*configservice.GetComplianceDetailsByResourceInput) (*configservice.GetComplianceDetailsByResourceOutput, error)

	GetComplianceSummaryByConfigRuleRequest(*configservice.GetComplianceSummaryByConfigRuleInput) (*request.Request, *configservice.GetComplianceSummaryByConfigRuleOutput)

	GetComplianceSummaryByConfigRule(*configservice.GetComplianceSummaryByConfigRuleInput) (*configservice.GetComplianceSummaryByConfigRuleOutput, error)

	GetComplianceSummaryByResourceTypeRequest(*configservice.GetComplianceSummaryByResourceTypeInput) (*request.Request, *configservice.GetComplianceSummaryByResourceTypeOutput)

	GetComplianceSummaryByResourceType(*configservice.GetComplianceSummaryByResourceTypeInput) (*configservice.GetComplianceSummaryByResourceTypeOutput, error)

	GetResourceConfigHistoryRequest(*configservice.GetResourceConfigHistoryInput) (*request.Request, *configservice.GetResourceConfigHistoryOutput)

	GetResourceConfigHistory(*configservice.GetResourceConfigHistoryInput) (*configservice.GetResourceConfigHistoryOutput, error)

	GetResourceConfigHistoryPages(*configservice.GetResourceConfigHistoryInput, func(*configservice.GetResourceConfigHistoryOutput, bool) bool) error

	ListDiscoveredResourcesRequest(*configservice.ListDiscoveredResourcesInput) (*request.Request, *configservice.ListDiscoveredResourcesOutput)

	ListDiscoveredResources(*configservice.ListDiscoveredResourcesInput) (*configservice.ListDiscoveredResourcesOutput, error)

	PutConfigRuleRequest(*configservice.PutConfigRuleInput) (*request.Request, *configservice.PutConfigRuleOutput)

	PutConfigRule(*configservice.PutConfigRuleInput) (*configservice.PutConfigRuleOutput, error)

	PutConfigurationRecorderRequest(*configservice.PutConfigurationRecorderInput) (*request.Request, *configservice.PutConfigurationRecorderOutput)

	PutConfigurationRecorder(*configservice.PutConfigurationRecorderInput) (*configservice.PutConfigurationRecorderOutput, error)

	PutDeliveryChannelRequest(*configservice.PutDeliveryChannelInput) (*request.Request, *configservice.PutDeliveryChannelOutput)

	PutDeliveryChannel(*configservice.PutDeliveryChannelInput) (*configservice.PutDeliveryChannelOutput, error)

	PutEvaluationsRequest(*configservice.PutEvaluationsInput) (*request.Request, *configservice.PutEvaluationsOutput)

	PutEvaluations(*configservice.PutEvaluationsInput) (*configservice.PutEvaluationsOutput, error)

	StartConfigRulesEvaluationRequest(*configservice.StartConfigRulesEvaluationInput) (*request.Request, *configservice.StartConfigRulesEvaluationOutput)

	StartConfigRulesEvaluation(*configservice.StartConfigRulesEvaluationInput) (*configservice.StartConfigRulesEvaluationOutput, error)

	StartConfigurationRecorderRequest(*configservice.StartConfigurationRecorderInput) (*request.Request, *configservice.StartConfigurationRecorderOutput)

	StartConfigurationRecorder(*configservice.StartConfigurationRecorderInput) (*configservice.StartConfigurationRecorderOutput, error)

	StopConfigurationRecorderRequest(*configservice.StopConfigurationRecorderInput) (*request.Request, *configservice.StopConfigurationRecorderOutput)

	StopConfigurationRecorder(*configservice.StopConfigurationRecorderInput) (*configservice.StopConfigurationRecorderOutput, error)
}

var _ ConfigServiceAPI = (*configservice.ConfigService)(nil)
