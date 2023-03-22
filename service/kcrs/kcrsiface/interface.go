// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package kcrsiface provides an interface to enable mocking the kcrs service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package kcrsiface

import (
	"github.com/KscSDK/ksc-sdk-go/service/kcrs"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
)

// KcrsAPI provides an interface to enable mocking the
// kcrs.Kcrs service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // kcrs.
//    func myFunc(svc kcrsiface.KcrsAPI) bool {
//        // Make svc.CreateInstance request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := kcrs.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockKcrsClient struct {
//        kcrsiface.KcrsAPI
//    }
//    func (m *mockKcrsClient) CreateInstance(input *map[string]interface{}) (*map[string]interface{}, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockKcrsClient{}
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
type KcrsAPI interface {
	CreateInstance(*map[string]interface{}) (*map[string]interface{}, error)
	CreateInstanceWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateInstanceRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateInstanceToken(*map[string]interface{}) (*map[string]interface{}, error)
	CreateInstanceTokenWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateInstanceTokenRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateInternalEndpoint(*map[string]interface{}) (*map[string]interface{}, error)
	CreateInternalEndpointWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateInternalEndpointRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateInternalEndpointDns(*map[string]interface{}) (*map[string]interface{}, error)
	CreateInternalEndpointDnsWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateInternalEndpointDnsRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateNamespace(*map[string]interface{}) (*map[string]interface{}, error)
	CreateNamespaceWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateNamespaceRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateWebhookTrigger(*map[string]interface{}) (*map[string]interface{}, error)
	CreateWebhookTriggerWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateWebhookTriggerRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteImages(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteImagesWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteImagesRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteInstance(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteInstanceWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteInstanceRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteInstanceToken(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteInstanceTokenWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteInstanceTokenRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteInternalEndpoint(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteInternalEndpointWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteInternalEndpointRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteInternalEndpointDns(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteInternalEndpointDnsWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteInternalEndpointDnsRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteNamespace(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteNamespaceWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteNamespaceRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteRepoTag(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteRepoTagWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteRepoTagRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteRepository(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteRepositoryWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteRepositoryRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteWebhookTrigger(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteWebhookTriggerWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteWebhookTriggerRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeImageScan(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeImageScanWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeImageScanRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeImages(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeImagesWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeImagesRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeInstance(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeInstanceWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeInstanceRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeInstanceToken(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeInstanceTokenWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeInstanceTokenRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeInstanceUsage(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeInstanceUsageWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeInstanceUsageRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeInternalEndpoint(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeInternalEndpointWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeInternalEndpointRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeInternalEndpointDns(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeInternalEndpointDnsWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeInternalEndpointDnsRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeNamespace(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeNamespaceWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeNamespaceRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeRepository(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeRepositoryWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeRepositoryRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeWebhookTrigger(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeWebhookTriggerWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeWebhookTriggerRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifyInstanceTokenInformation(*map[string]interface{}) (*map[string]interface{}, error)
	ModifyInstanceTokenInformationWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifyInstanceTokenInformationRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifyInstanceTokenStatus(*map[string]interface{}) (*map[string]interface{}, error)
	ModifyInstanceTokenStatusWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifyInstanceTokenStatusRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifyNamespaceType(*map[string]interface{}) (*map[string]interface{}, error)
	ModifyNamespaceTypeWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifyNamespaceTypeRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifyRepoDesc(*map[string]interface{}) (*map[string]interface{}, error)
	ModifyRepoDescWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifyRepoDescRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifyWebhookTrigger(*map[string]interface{}) (*map[string]interface{}, error)
	ModifyWebhookTriggerWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifyWebhookTriggerRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	StartImageScan(*map[string]interface{}) (*map[string]interface{}, error)
	StartImageScanWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	StartImageScanRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})
}

var _ KcrsAPI = (*kcrs.Kcrs)(nil)