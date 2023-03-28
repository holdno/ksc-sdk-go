// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cdnv2

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
)

const opGetClientRequestDataGet = "GetClientRequestData"

// GetClientRequestDataGetRequest generates a "ksc/request.Request" representing the
// client's request for the GetClientRequestDataGet operation. The "output" return
// value will be populated with the request's response once the request completes
// successfully.
//
// Use "Send" method on the returned Request to send the API call to the service.
// the "output" return value is not valid until after Send returns without error.
//
// See GetClientRequestDataGet for more information on using the GetClientRequestDataGet
// API call, and error handling.
//
// This method is useful when you want to inject custom logic or configuration
// into the SDK's request lifecycle. Such as custom headers, or retry logic.
//
//
//    // Example sending a request using the GetClientRequestDataGetRequest method.
//    req, resp := client.GetClientRequestDataGetRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/cdn-2020-06-30/GetClientRequestDataGet
func (c *Cdnv2) GetClientRequestDataGetRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opGetClientRequestDataGet,
		HTTPMethod: "GET",
		HTTPPath:   "/2020-06-30/statistics/GetClientRequestData",
	}

	if input == nil {
		input = &map[string]interface{}{}
	}

	output = &map[string]interface{}{}
	req = c.newRequest(op, input, output)

	return
}

func (c *Cdnv2) GetRealTimeHitRateDataPostRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       "GetRealTimeHitRateData",
		HTTPMethod: "POST",
		HTTPPath:   "/2020-06-30/statistics/GetRealTimeHitRateData",
	}

	if input == nil {
		input = &map[string]interface{}{}
	}

	output = &map[string]interface{}{}
	req = c.newRequest(op, input, output)
	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")
	
	return
}

func (c *Cdnv2) GetFlowHitRateDataPostRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       "GetFlowHitRateData",
		HTTPMethod: "POST",
		HTTPPath:   "/2020-06-30/statistics/GetFlowHitRateData",
	}

	if input == nil {
		input = &map[string]interface{}{}
	}

	output = &map[string]interface{}{}
	req = c.newRequest(op, input, output)
	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")
	
	return
}

func (c *Cdnv2) GetReqHitRateDataPostRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       "GetReqHitRateData",
		HTTPMethod: "POST",
		HTTPPath:   "/2020-06-30/statistics/GetReqHitRateData",
	}

	if input == nil {
		input = &map[string]interface{}{}
	}

	output = &map[string]interface{}{}
	req = c.newRequest(op, input, output)
	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")
	
	return
}

// GetClientRequestDataGet API operation for cdnv2.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the KSC API reference guide for cdnv2's
// API operation GetClientRequestDataGet for usage and error information.
// See also, https://docs.aws.amazon.com/goto/WebAPI/cdn-2020-06-30/GetClientRequestDataGet
func (c *Cdnv2) GetClientRequestDataGet(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.GetClientRequestDataGetRequest(input)
	return out, req.Send()
}

// GetClientRequestDataGetWithContext is the same as GetClientRequestDataGet with the addition of
// the ability to pass a context and additional request options.
//
// See GetClientRequestDataGet for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *Cdnv2) GetClientRequestDataGetWithContext(ctx aws.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.GetClientRequestDataGetRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opGetClientRequestDataPost = "GetClientRequestData"

// GetClientRequestDataPostRequest generates a "ksc/request.Request" representing the
// client's request for the GetClientRequestDataPost operation. The "output" return
// value will be populated with the request's response once the request completes
// successfully.
//
// Use "Send" method on the returned Request to send the API call to the service.
// the "output" return value is not valid until after Send returns without error.
//
// See GetClientRequestDataPost for more information on using the GetClientRequestDataPost
// API call, and error handling.
//
// This method is useful when you want to inject custom logic or configuration
// into the SDK's request lifecycle. Such as custom headers, or retry logic.
//
//
//    // Example sending a request using the GetClientRequestDataPostRequest method.
//    req, resp := client.GetClientRequestDataPostRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/cdn-2020-06-30/GetClientRequestDataPost
func (c *Cdnv2) GetClientRequestDataPostRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opGetClientRequestDataPost,
		HTTPMethod: "POST",
		HTTPPath:   "/2020-06-30/statistics/GetClientRequestData",
	}

	if input == nil {
		input = &map[string]interface{}{}
	}

	output = &map[string]interface{}{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// GetClientRequestDataPost API operation for cdnv2.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the KSC API reference guide for cdnv2's
// API operation GetClientRequestDataPost for usage and error information.
// See also, https://docs.aws.amazon.com/goto/WebAPI/cdn-2020-06-30/GetClientRequestDataPost
func (c *Cdnv2) GetClientRequestDataPost(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.GetClientRequestDataPostRequest(input)
	return out, req.Send()
}

// GetClientRequestDataPostWithContext is the same as GetClientRequestDataPost with the addition of
// the ability to pass a context and additional request options.
//
// See GetClientRequestDataPost for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *Cdnv2) GetClientRequestDataPostWithContext(ctx aws.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.GetClientRequestDataPostRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opGetServerDataGet = "GetServerData"

// GetServerDataGetRequest generates a "ksc/request.Request" representing the
// client's request for the GetServerDataGet operation. The "output" return
// value will be populated with the request's response once the request completes
// successfully.
//
// Use "Send" method on the returned Request to send the API call to the service.
// the "output" return value is not valid until after Send returns without error.
//
// See GetServerDataGet for more information on using the GetServerDataGet
// API call, and error handling.
//
// This method is useful when you want to inject custom logic or configuration
// into the SDK's request lifecycle. Such as custom headers, or retry logic.
//
//
//    // Example sending a request using the GetServerDataGetRequest method.
//    req, resp := client.GetServerDataGetRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/cdn-2020-06-30/GetServerDataGet
func (c *Cdnv2) GetServerDataGetRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opGetServerDataGet,
		HTTPMethod: "GET",
		HTTPPath:   "/2020-06-30/statistics/GetServerData",
	}

	if input == nil {
		input = &map[string]interface{}{}
	}

	output = &map[string]interface{}{}
	req = c.newRequest(op, input, output)

	return
}

// GetServerDataGet API operation for cdnv2.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the KSC API reference guide for cdnv2's
// API operation GetServerDataGet for usage and error information.
// See also, https://docs.aws.amazon.com/goto/WebAPI/cdn-2020-06-30/GetServerDataGet
func (c *Cdnv2) GetServerDataGet(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.GetServerDataGetRequest(input)
	return out, req.Send()
}

// GetServerDataGetWithContext is the same as GetServerDataGet with the addition of
// the ability to pass a context and additional request options.
//
// See GetServerDataGet for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *Cdnv2) GetServerDataGetWithContext(ctx aws.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.GetServerDataGetRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opGetServerDataPost = "GetServerData"

// GetServerDataPostRequest generates a "ksc/request.Request" representing the
// client's request for the GetServerDataPost operation. The "output" return
// value will be populated with the request's response once the request completes
// successfully.
//
// Use "Send" method on the returned Request to send the API call to the service.
// the "output" return value is not valid until after Send returns without error.
//
// See GetServerDataPost for more information on using the GetServerDataPost
// API call, and error handling.
//
// This method is useful when you want to inject custom logic or configuration
// into the SDK's request lifecycle. Such as custom headers, or retry logic.
//
//
//    // Example sending a request using the GetServerDataPostRequest method.
//    req, resp := client.GetServerDataPostRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/cdn-2020-06-30/GetServerDataPost
func (c *Cdnv2) GetServerDataPostRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opGetServerDataPost,
		HTTPMethod: "POST",
		HTTPPath:   "/2020-06-30/statistics/GetServerData",
	}

	if input == nil {
		input = &map[string]interface{}{}
	}

	output = &map[string]interface{}{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// GetServerDataPost API operation for cdnv2.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the KSC API reference guide for cdnv2's
// API operation GetServerDataPost for usage and error information.
// See also, https://docs.aws.amazon.com/goto/WebAPI/cdn-2020-06-30/GetServerDataPost
func (c *Cdnv2) GetServerDataPost(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.GetServerDataPostRequest(input)
	return out, req.Send()
}

// GetServerDataPostWithContext is the same as GetServerDataPost with the addition of
// the ability to pass a context and additional request options.
//
// See GetServerDataPost for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *Cdnv2) GetServerDataPostWithContext(ctx aws.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.GetServerDataPostRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}


func (c *Cdnv2) GetRealTimeHitRateData(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.GetRealTimeHitRateDataPostRequest(input)
	return out, req.Send()
}

func (c *Cdnv2) GetReqHitRateData(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.GetReqHitRateDataPostRequest(input)
	return out, req.Send()
}

func (c *Cdnv2) GetFlowHitRateData(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.GetFlowHitRateDataPostRequest(input)
	return out, req.Send()
}
