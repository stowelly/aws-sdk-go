// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package lambdaiface provides an interface to enable mocking the AWS Lambda service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package lambdaiface

import (
	"github.com/stowelly/aws-sdk-go/aws/request"
	"github.com/stowelly/aws-sdk-go/service/lambda"
)

// LambdaAPI provides an interface to enable mocking the
// lambda.Lambda service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // AWS Lambda.
//    func myFunc(svc lambdaiface.LambdaAPI) bool {
//        // Make svc.AddPermission request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := lambda.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockLambdaClient struct {
//        lambdaiface.LambdaAPI
//    }
//    func (m *mockLambdaClient) AddPermission(input *lambda.AddPermissionInput) (*lambda.AddPermissionOutput, error) {
//        // mock response/functionality
//    }
//
//    TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockLambdaClient{}
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
type LambdaAPI interface {
	AddPermissionRequest(*lambda.AddPermissionInput) (*request.Request, *lambda.AddPermissionOutput)

	AddPermission(*lambda.AddPermissionInput) (*lambda.AddPermissionOutput, error)

	CreateAliasRequest(*lambda.CreateAliasInput) (*request.Request, *lambda.AliasConfiguration)

	CreateAlias(*lambda.CreateAliasInput) (*lambda.AliasConfiguration, error)

	CreateEventSourceMappingRequest(*lambda.CreateEventSourceMappingInput) (*request.Request, *lambda.EventSourceMappingConfiguration)

	CreateEventSourceMapping(*lambda.CreateEventSourceMappingInput) (*lambda.EventSourceMappingConfiguration, error)

	CreateFunctionRequest(*lambda.CreateFunctionInput) (*request.Request, *lambda.FunctionConfiguration)

	CreateFunction(*lambda.CreateFunctionInput) (*lambda.FunctionConfiguration, error)

	DeleteAliasRequest(*lambda.DeleteAliasInput) (*request.Request, *lambda.DeleteAliasOutput)

	DeleteAlias(*lambda.DeleteAliasInput) (*lambda.DeleteAliasOutput, error)

	DeleteEventSourceMappingRequest(*lambda.DeleteEventSourceMappingInput) (*request.Request, *lambda.EventSourceMappingConfiguration)

	DeleteEventSourceMapping(*lambda.DeleteEventSourceMappingInput) (*lambda.EventSourceMappingConfiguration, error)

	DeleteFunctionRequest(*lambda.DeleteFunctionInput) (*request.Request, *lambda.DeleteFunctionOutput)

	DeleteFunction(*lambda.DeleteFunctionInput) (*lambda.DeleteFunctionOutput, error)

	GetAccountSettingsRequest(*lambda.GetAccountSettingsInput) (*request.Request, *lambda.GetAccountSettingsOutput)

	GetAccountSettings(*lambda.GetAccountSettingsInput) (*lambda.GetAccountSettingsOutput, error)

	GetAliasRequest(*lambda.GetAliasInput) (*request.Request, *lambda.AliasConfiguration)

	GetAlias(*lambda.GetAliasInput) (*lambda.AliasConfiguration, error)

	GetEventSourceMappingRequest(*lambda.GetEventSourceMappingInput) (*request.Request, *lambda.EventSourceMappingConfiguration)

	GetEventSourceMapping(*lambda.GetEventSourceMappingInput) (*lambda.EventSourceMappingConfiguration, error)

	GetFunctionRequest(*lambda.GetFunctionInput) (*request.Request, *lambda.GetFunctionOutput)

	GetFunction(*lambda.GetFunctionInput) (*lambda.GetFunctionOutput, error)

	GetFunctionConfigurationRequest(*lambda.GetFunctionConfigurationInput) (*request.Request, *lambda.FunctionConfiguration)

	GetFunctionConfiguration(*lambda.GetFunctionConfigurationInput) (*lambda.FunctionConfiguration, error)

	GetPolicyRequest(*lambda.GetPolicyInput) (*request.Request, *lambda.GetPolicyOutput)

	GetPolicy(*lambda.GetPolicyInput) (*lambda.GetPolicyOutput, error)

	InvokeRequest(*lambda.InvokeInput) (*request.Request, *lambda.InvokeOutput)

	Invoke(*lambda.InvokeInput) (*lambda.InvokeOutput, error)

	InvokeAsyncRequest(*lambda.InvokeAsyncInput) (*request.Request, *lambda.InvokeAsyncOutput)

	InvokeAsync(*lambda.InvokeAsyncInput) (*lambda.InvokeAsyncOutput, error)

	ListAliasesRequest(*lambda.ListAliasesInput) (*request.Request, *lambda.ListAliasesOutput)

	ListAliases(*lambda.ListAliasesInput) (*lambda.ListAliasesOutput, error)

	ListEventSourceMappingsRequest(*lambda.ListEventSourceMappingsInput) (*request.Request, *lambda.ListEventSourceMappingsOutput)

	ListEventSourceMappings(*lambda.ListEventSourceMappingsInput) (*lambda.ListEventSourceMappingsOutput, error)

	ListEventSourceMappingsPages(*lambda.ListEventSourceMappingsInput, func(*lambda.ListEventSourceMappingsOutput, bool) bool) error

	ListFunctionsRequest(*lambda.ListFunctionsInput) (*request.Request, *lambda.ListFunctionsOutput)

	ListFunctions(*lambda.ListFunctionsInput) (*lambda.ListFunctionsOutput, error)

	ListFunctionsPages(*lambda.ListFunctionsInput, func(*lambda.ListFunctionsOutput, bool) bool) error

	ListVersionsByFunctionRequest(*lambda.ListVersionsByFunctionInput) (*request.Request, *lambda.ListVersionsByFunctionOutput)

	ListVersionsByFunction(*lambda.ListVersionsByFunctionInput) (*lambda.ListVersionsByFunctionOutput, error)

	PublishVersionRequest(*lambda.PublishVersionInput) (*request.Request, *lambda.FunctionConfiguration)

	PublishVersion(*lambda.PublishVersionInput) (*lambda.FunctionConfiguration, error)

	RemovePermissionRequest(*lambda.RemovePermissionInput) (*request.Request, *lambda.RemovePermissionOutput)

	RemovePermission(*lambda.RemovePermissionInput) (*lambda.RemovePermissionOutput, error)

	UpdateAliasRequest(*lambda.UpdateAliasInput) (*request.Request, *lambda.AliasConfiguration)

	UpdateAlias(*lambda.UpdateAliasInput) (*lambda.AliasConfiguration, error)

	UpdateEventSourceMappingRequest(*lambda.UpdateEventSourceMappingInput) (*request.Request, *lambda.EventSourceMappingConfiguration)

	UpdateEventSourceMapping(*lambda.UpdateEventSourceMappingInput) (*lambda.EventSourceMappingConfiguration, error)

	UpdateFunctionCodeRequest(*lambda.UpdateFunctionCodeInput) (*request.Request, *lambda.FunctionConfiguration)

	UpdateFunctionCode(*lambda.UpdateFunctionCodeInput) (*lambda.FunctionConfiguration, error)

	UpdateFunctionConfigurationRequest(*lambda.UpdateFunctionConfigurationInput) (*request.Request, *lambda.FunctionConfiguration)

	UpdateFunctionConfiguration(*lambda.UpdateFunctionConfigurationInput) (*lambda.FunctionConfiguration, error)
}

var _ LambdaAPI = (*lambda.Lambda)(nil)
