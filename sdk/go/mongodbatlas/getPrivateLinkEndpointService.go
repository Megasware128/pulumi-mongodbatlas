// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package mongodbatlas

import (
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
