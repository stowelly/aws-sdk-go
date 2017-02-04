// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package elasticbeanstalk

import (
	"github.com/stowelly/aws-sdk-go/aws"
	"github.com/stowelly/aws-sdk-go/aws/client"
	"github.com/stowelly/aws-sdk-go/aws/client/metadata"
	"github.com/stowelly/aws-sdk-go/aws/request"
	"github.com/stowelly/aws-sdk-go/aws/signer/v4"
	"github.com/stowelly/aws-sdk-go/private/protocol/query"
)

// AWS Elastic Beanstalk makes it easy for you to create, deploy, and manage
// scalable, fault-tolerant applications running on the Amazon Web Services
// cloud.
//
// For more information about this product, go to the AWS Elastic Beanstalk
// (http://aws.amazon.com/elasticbeanstalk/) details page. The location of the
// latest AWS Elastic Beanstalk WSDL is http://elasticbeanstalk.s3.amazonaws.com/doc/2010-12-01/AWSElasticBeanstalk.wsdl
// (http://elasticbeanstalk.s3.amazonaws.com/doc/2010-12-01/AWSElasticBeanstalk.wsdl).
// To install the Software Development Kits (SDKs), Integrated Development Environment
// (IDE) Toolkits, and command line tools that enable you to access the API,
// go to Tools for Amazon Web Services (http://aws.amazon.com/tools/).
//
// Endpoints
//
// For a list of region-specific endpoints that AWS Elastic Beanstalk supports,
// go to Regions and Endpoints (http://docs.aws.amazon.com/general/latest/gr/rande.html#elasticbeanstalk_region)
// in the Amazon Web Services Glossary.
// The service client's operations are safe to be used concurrently.
// It is not safe to mutate any of the client's properties though.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticbeanstalk-2010-12-01
type ElasticBeanstalk struct {
	*client.Client
}

// Used for custom client initialization logic
var initClient func(*client.Client)

// Used for custom request initialization logic
var initRequest func(*request.Request)

// Service information constants
const (
	ServiceName = "elasticbeanstalk" // Service endpoint prefix API calls made to.
	EndpointsID = ServiceName        // Service ID for Regions and Endpoints metadata.
)

// New creates a new instance of the ElasticBeanstalk client with a session.
// If additional configuration is needed for the client instance use the optional
// aws.Config parameter to add your extra config.
//
// Example:
//     // Create a ElasticBeanstalk client from just a session.
//     svc := elasticbeanstalk.New(mySession)
//
//     // Create a ElasticBeanstalk client with additional configuration
//     svc := elasticbeanstalk.New(mySession, aws.NewConfig().WithRegion("us-west-2"))
func New(p client.ConfigProvider, cfgs ...*aws.Config) *ElasticBeanstalk {
	c := p.ClientConfig(EndpointsID, cfgs...)
	return newClient(*c.Config, c.Handlers, c.Endpoint, c.SigningRegion, c.SigningName)
}

// newClient creates, initializes and returns a new service client instance.
func newClient(cfg aws.Config, handlers request.Handlers, endpoint, signingRegion, signingName string) *ElasticBeanstalk {
	svc := &ElasticBeanstalk{
		Client: client.New(
			cfg,
			metadata.ClientInfo{
				ServiceName:   ServiceName,
				SigningName:   signingName,
				SigningRegion: signingRegion,
				Endpoint:      endpoint,
				APIVersion:    "2010-12-01",
			},
			handlers,
		),
	}

	// Handlers
	svc.Handlers.Sign.PushBackNamed(v4.SignRequestHandler)
	svc.Handlers.Build.PushBackNamed(query.BuildHandler)
	svc.Handlers.Unmarshal.PushBackNamed(query.UnmarshalHandler)
	svc.Handlers.UnmarshalMeta.PushBackNamed(query.UnmarshalMetaHandler)
	svc.Handlers.UnmarshalError.PushBackNamed(query.UnmarshalErrorHandler)

	// Run custom client initialization if present
	if initClient != nil {
		initClient(svc.Client)
	}

	return svc
}

// newRequest creates a new request for a ElasticBeanstalk operation and runs any
// custom request initialization.
func (c *ElasticBeanstalk) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(req)
	}

	return req
}
