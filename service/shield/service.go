// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package shield

import (
	"github.com/stowelly/aws-sdk-go/aws"
	"github.com/stowelly/aws-sdk-go/aws/client"
	"github.com/stowelly/aws-sdk-go/aws/client/metadata"
	"github.com/stowelly/aws-sdk-go/aws/request"
	"github.com/stowelly/aws-sdk-go/aws/signer/v4"
	"github.com/stowelly/aws-sdk-go/private/protocol/jsonrpc"
)

// This is the AWS Shield Advanced API Reference. This guide is for developers
// who need detailed information about the AWS Shield Advanced API actions,
// data types, and errors. For detailed information about AWS WAF and AWS Shield
// Advanced features and an overview of how to use the AWS WAF and AWS Shield
// Advanced APIs, see the AWS WAF and AWS Shield Developer Guide (http://docs.aws.amazon.com/waf/latest/developerguide/).
// The service client's operations are safe to be used concurrently.
// It is not safe to mutate any of the client's properties though.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/shield-2016-06-02
type Shield struct {
	*client.Client
}

// Used for custom client initialization logic
var initClient func(*client.Client)

// Used for custom request initialization logic
var initRequest func(*request.Request)

// Service information constants
const (
	ServiceName = "shield"    // Service endpoint prefix API calls made to.
	EndpointsID = ServiceName // Service ID for Regions and Endpoints metadata.
)

// New creates a new instance of the Shield client with a session.
// If additional configuration is needed for the client instance use the optional
// aws.Config parameter to add your extra config.
//
// Example:
//     // Create a Shield client from just a session.
//     svc := shield.New(mySession)
//
//     // Create a Shield client with additional configuration
//     svc := shield.New(mySession, aws.NewConfig().WithRegion("us-west-2"))
func New(p client.ConfigProvider, cfgs ...*aws.Config) *Shield {
	c := p.ClientConfig(EndpointsID, cfgs...)
	return newClient(*c.Config, c.Handlers, c.Endpoint, c.SigningRegion, c.SigningName)
}

// newClient creates, initializes and returns a new service client instance.
func newClient(cfg aws.Config, handlers request.Handlers, endpoint, signingRegion, signingName string) *Shield {
	svc := &Shield{
		Client: client.New(
			cfg,
			metadata.ClientInfo{
				ServiceName:   ServiceName,
				SigningName:   signingName,
				SigningRegion: signingRegion,
				Endpoint:      endpoint,
				APIVersion:    "2016-06-02",
				JSONVersion:   "1.1",
				TargetPrefix:  "AWSShield_20160616",
			},
			handlers,
		),
	}

	// Handlers
	svc.Handlers.Sign.PushBackNamed(v4.SignRequestHandler)
	svc.Handlers.Build.PushBackNamed(jsonrpc.BuildHandler)
	svc.Handlers.Unmarshal.PushBackNamed(jsonrpc.UnmarshalHandler)
	svc.Handlers.UnmarshalMeta.PushBackNamed(jsonrpc.UnmarshalMetaHandler)
	svc.Handlers.UnmarshalError.PushBackNamed(jsonrpc.UnmarshalErrorHandler)

	// Run custom client initialization if present
	if initClient != nil {
		initClient(svc.Client)
	}

	return svc
}

// newRequest creates a new request for a Shield operation and runs any
// custom request initialization.
func (c *Shield) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(req)
	}

	return req
}
