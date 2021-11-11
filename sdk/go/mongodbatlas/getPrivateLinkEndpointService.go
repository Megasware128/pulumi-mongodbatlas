// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package mongodbatlas

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// `PrivateLinkEndpointService` describe a Private Endpoint Link. This represents a Private Endpoint Link Connection that wants to retrieve details in an Atlas project.
//
// > **NOTE:** Groups and projects are synonymous terms. You may find groupId in the official documentation.
func LookupPrivateLinkEndpointService(ctx *pulumi.Context, args *LookupPrivateLinkEndpointServiceArgs, opts ...pulumi.InvokeOption) (*LookupPrivateLinkEndpointServiceResult, error) {
	var rv LookupPrivateLinkEndpointServiceResult
	err := ctx.Invoke("mongodbatlas:index/getPrivateLinkEndpointService:getPrivateLinkEndpointService", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getPrivateLinkEndpointService.
type LookupPrivateLinkEndpointServiceArgs struct {
	// Unique identifier of the `AWS` or `AZURE` resource.
	EndpointServiceId string `pulumi:"endpointServiceId"`
	// Unique identifier of the private endpoint service for which you want to retrieve a private endpoint.
	PrivateLinkId string `pulumi:"privateLinkId"`
	// Unique identifier for the project.
	ProjectId string `pulumi:"projectId"`
	// Cloud provider for which you want to create a private endpoint. Atlas accepts `AWS` or `AZURE`.
	ProviderName string `pulumi:"providerName"`
}

// A collection of values returned by getPrivateLinkEndpointService.
type LookupPrivateLinkEndpointServiceResult struct {
	// Status of the interface endpoint for AWS.
	// Returns one of the following values:
	AwsConnectionStatus string `pulumi:"awsConnectionStatus"`
	// Status of the interface endpoint for AZURE.
	// Returns one of the following values:
	AzureStatus string `pulumi:"azureStatus"`
	// Indicates if Atlas received a request to remove the interface endpoint from the private endpoint connection.
	DeleteRequested   bool   `pulumi:"deleteRequested"`
	EndpointServiceId string `pulumi:"endpointServiceId"`
	// Error message pertaining to the interface endpoint. Returns null if there are no errors.
	ErrorMessage string `pulumi:"errorMessage"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Unique identifier of the interface endpoint.
	InterfaceEndpointId string `pulumi:"interfaceEndpointId"`
	// Name of the connection for this private endpoint that Atlas generates.
	PrivateEndpointConnectionName string `pulumi:"privateEndpointConnectionName"`
	// Private IP address of the private endpoint network interface.
	PrivateEndpointIpAddress string `pulumi:"privateEndpointIpAddress"`
	// Unique identifier of the private endpoint.
	PrivateEndpointResourceId string `pulumi:"privateEndpointResourceId"`
	PrivateLinkId             string `pulumi:"privateLinkId"`
	ProjectId                 string `pulumi:"projectId"`
	ProviderName              string `pulumi:"providerName"`
}

func LookupPrivateLinkEndpointServiceOutput(ctx *pulumi.Context, args LookupPrivateLinkEndpointServiceOutputArgs, opts ...pulumi.InvokeOption) LookupPrivateLinkEndpointServiceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPrivateLinkEndpointServiceResult, error) {
			args := v.(LookupPrivateLinkEndpointServiceArgs)
			r, err := LookupPrivateLinkEndpointService(ctx, &args, opts...)
			return *r, err
		}).(LookupPrivateLinkEndpointServiceResultOutput)
}

// A collection of arguments for invoking getPrivateLinkEndpointService.
type LookupPrivateLinkEndpointServiceOutputArgs struct {
	// Unique identifier of the `AWS` or `AZURE` resource.
	EndpointServiceId pulumi.StringInput `pulumi:"endpointServiceId"`
	// Unique identifier of the private endpoint service for which you want to retrieve a private endpoint.
	PrivateLinkId pulumi.StringInput `pulumi:"privateLinkId"`
	// Unique identifier for the project.
	ProjectId pulumi.StringInput `pulumi:"projectId"`
	// Cloud provider for which you want to create a private endpoint. Atlas accepts `AWS` or `AZURE`.
	ProviderName pulumi.StringInput `pulumi:"providerName"`
}

func (LookupPrivateLinkEndpointServiceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPrivateLinkEndpointServiceArgs)(nil)).Elem()
}

// A collection of values returned by getPrivateLinkEndpointService.
type LookupPrivateLinkEndpointServiceResultOutput struct{ *pulumi.OutputState }

func (LookupPrivateLinkEndpointServiceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPrivateLinkEndpointServiceResult)(nil)).Elem()
}

func (o LookupPrivateLinkEndpointServiceResultOutput) ToLookupPrivateLinkEndpointServiceResultOutput() LookupPrivateLinkEndpointServiceResultOutput {
	return o
}

func (o LookupPrivateLinkEndpointServiceResultOutput) ToLookupPrivateLinkEndpointServiceResultOutputWithContext(ctx context.Context) LookupPrivateLinkEndpointServiceResultOutput {
	return o
}

// Status of the interface endpoint for AWS.
// Returns one of the following values:
func (o LookupPrivateLinkEndpointServiceResultOutput) AwsConnectionStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateLinkEndpointServiceResult) string { return v.AwsConnectionStatus }).(pulumi.StringOutput)
}

// Status of the interface endpoint for AZURE.
// Returns one of the following values:
func (o LookupPrivateLinkEndpointServiceResultOutput) AzureStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateLinkEndpointServiceResult) string { return v.AzureStatus }).(pulumi.StringOutput)
}

// Indicates if Atlas received a request to remove the interface endpoint from the private endpoint connection.
func (o LookupPrivateLinkEndpointServiceResultOutput) DeleteRequested() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupPrivateLinkEndpointServiceResult) bool { return v.DeleteRequested }).(pulumi.BoolOutput)
}

func (o LookupPrivateLinkEndpointServiceResultOutput) EndpointServiceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateLinkEndpointServiceResult) string { return v.EndpointServiceId }).(pulumi.StringOutput)
}

// Error message pertaining to the interface endpoint. Returns null if there are no errors.
func (o LookupPrivateLinkEndpointServiceResultOutput) ErrorMessage() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateLinkEndpointServiceResult) string { return v.ErrorMessage }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupPrivateLinkEndpointServiceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateLinkEndpointServiceResult) string { return v.Id }).(pulumi.StringOutput)
}

// Unique identifier of the interface endpoint.
func (o LookupPrivateLinkEndpointServiceResultOutput) InterfaceEndpointId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateLinkEndpointServiceResult) string { return v.InterfaceEndpointId }).(pulumi.StringOutput)
}

// Name of the connection for this private endpoint that Atlas generates.
func (o LookupPrivateLinkEndpointServiceResultOutput) PrivateEndpointConnectionName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateLinkEndpointServiceResult) string { return v.PrivateEndpointConnectionName }).(pulumi.StringOutput)
}

// Private IP address of the private endpoint network interface.
func (o LookupPrivateLinkEndpointServiceResultOutput) PrivateEndpointIpAddress() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateLinkEndpointServiceResult) string { return v.PrivateEndpointIpAddress }).(pulumi.StringOutput)
}

// Unique identifier of the private endpoint.
func (o LookupPrivateLinkEndpointServiceResultOutput) PrivateEndpointResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateLinkEndpointServiceResult) string { return v.PrivateEndpointResourceId }).(pulumi.StringOutput)
}

func (o LookupPrivateLinkEndpointServiceResultOutput) PrivateLinkId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateLinkEndpointServiceResult) string { return v.PrivateLinkId }).(pulumi.StringOutput)
}

func (o LookupPrivateLinkEndpointServiceResultOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateLinkEndpointServiceResult) string { return v.ProjectId }).(pulumi.StringOutput)
}

func (o LookupPrivateLinkEndpointServiceResultOutput) ProviderName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateLinkEndpointServiceResult) string { return v.ProviderName }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPrivateLinkEndpointServiceResultOutput{})
}
