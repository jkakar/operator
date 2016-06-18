// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package apigateway

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/client"
	"github.com/aws/aws-sdk-go/aws/client/metadata"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/private/protocol/restjson"
	"github.com/aws/aws-sdk-go/private/signer/v4"
)

// Amazon API Gateway helps developers deliver robust, secure and scalable mobile
// and web application backends. Amazon API Gateway allows developers to securely
// connect mobile and web applications to APIs that run on AWS Lambda, Amazon
// EC2, or other publicly addressable web services that are hosted outside of
// AWS.
//The service client's operations are safe to be used concurrently.
// It is not safe to mutate any of the client's properties though.
type APIGateway struct {
	*client.Client
}

// Used for custom client initialization logic
var initClient func(*client.Client)

// Used for custom request initialization logic
var initRequest func(*request.Request)

// A ServiceName is the name of the service the client will make API calls to.
const ServiceName = "apigateway"

// New creates a new instance of the APIGateway client with a session.
// If additional configuration is needed for the client instance use the optional
// aws.Config parameter to add your extra config.
//
// Example:
//     // Create a APIGateway client from just a session.
//     svc := apigateway.New(mySession)
//
//     // Create a APIGateway client with additional configuration
//     svc := apigateway.New(mySession, aws.NewConfig().WithRegion("us-west-2"))
func New(p client.ConfigProvider, cfgs ...*aws.Config) *APIGateway {
	c := p.ClientConfig(ServiceName, cfgs...)
	return newClient(*c.Config, c.Handlers, c.Endpoint, c.SigningRegion)
}

// newClient creates, initializes and returns a new service client instance.
func newClient(cfg aws.Config, handlers request.Handlers, endpoint, signingRegion string) *APIGateway {
	svc := &APIGateway{
		Client: client.New(
			cfg,
			metadata.ClientInfo{
				ServiceName:   ServiceName,
				SigningRegion: signingRegion,
				Endpoint:      endpoint,
				APIVersion:    "2015-07-09",
			},
			handlers,
		),
	}

	// Handlers
	svc.Handlers.Sign.PushBack(v4.Sign)
	svc.Handlers.Build.PushBackNamed(restjson.BuildHandler)
	svc.Handlers.Unmarshal.PushBackNamed(restjson.UnmarshalHandler)
	svc.Handlers.UnmarshalMeta.PushBackNamed(restjson.UnmarshalMetaHandler)
	svc.Handlers.UnmarshalError.PushBackNamed(restjson.UnmarshalErrorHandler)

	// Run custom client initialization if present
	if initClient != nil {
		initClient(svc.Client)
	}

	return svc
}

// newRequest creates a new request for a APIGateway operation and runs any
// custom request initialization.
func (c *APIGateway) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(req)
	}

	return req
}
