// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mongodbatlas

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// `privatelinkEndpointServiceAdl` Provides an Atlas Data Lake (ADL) and Online Archive PrivateLink endpoint resource.
//
// > **NOTE:** Groups and projects are synonymous terms. You may find groupId in the official documentation.
//
// ## Example Usage
func LookupPrivatelinkEndpointServiceAdl(ctx *pulumi.Context, args *LookupPrivatelinkEndpointServiceAdlArgs, opts ...pulumi.InvokeOption) (*LookupPrivatelinkEndpointServiceAdlResult, error) {
	var rv LookupPrivatelinkEndpointServiceAdlResult
	err := ctx.Invoke("mongodbatlas:index/getPrivatelinkEndpointServiceAdl:getPrivatelinkEndpointServiceAdl", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getPrivatelinkEndpointServiceAdl.
type LookupPrivatelinkEndpointServiceAdlArgs struct {
	// Unique 22-character alphanumeric string that identifies the private endpoint. Atlas supports AWS private endpoints using the [|aws| PrivateLink](https://aws.amazon.com/privatelink/) feature.
	EndpointId string `pulumi:"endpointId"`
	// Unique 24-digit hexadecimal string that identifies the project.
	ProjectId string `pulumi:"projectId"`
}

// A collection of values returned by getPrivatelinkEndpointServiceAdl.
type LookupPrivatelinkEndpointServiceAdlResult struct {
	// Human-readable string to associate with this private endpoint.
	Comment    string `pulumi:"comment"`
	EndpointId string `pulumi:"endpointId"`
	// The provider-assigned unique ID for this managed resource.
	Id        string `pulumi:"id"`
	ProjectId string `pulumi:"projectId"`
	// Human-readable label that identifies the cloud provider for this endpoint.
	ProviderName string `pulumi:"providerName"`
	// Human-readable label that identifies the type of resource to associate with this private endpoint.
	Type string `pulumi:"type"`
}

func LookupPrivatelinkEndpointServiceAdlOutput(ctx *pulumi.Context, args LookupPrivatelinkEndpointServiceAdlOutputArgs, opts ...pulumi.InvokeOption) LookupPrivatelinkEndpointServiceAdlResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPrivatelinkEndpointServiceAdlResult, error) {
			args := v.(LookupPrivatelinkEndpointServiceAdlArgs)
			r, err := LookupPrivatelinkEndpointServiceAdl(ctx, &args, opts...)
			var s LookupPrivatelinkEndpointServiceAdlResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPrivatelinkEndpointServiceAdlResultOutput)
}

// A collection of arguments for invoking getPrivatelinkEndpointServiceAdl.
type LookupPrivatelinkEndpointServiceAdlOutputArgs struct {
	// Unique 22-character alphanumeric string that identifies the private endpoint. Atlas supports AWS private endpoints using the [|aws| PrivateLink](https://aws.amazon.com/privatelink/) feature.
	EndpointId pulumi.StringInput `pulumi:"endpointId"`
	// Unique 24-digit hexadecimal string that identifies the project.
	ProjectId pulumi.StringInput `pulumi:"projectId"`
}

func (LookupPrivatelinkEndpointServiceAdlOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPrivatelinkEndpointServiceAdlArgs)(nil)).Elem()
}

// A collection of values returned by getPrivatelinkEndpointServiceAdl.
type LookupPrivatelinkEndpointServiceAdlResultOutput struct{ *pulumi.OutputState }

func (LookupPrivatelinkEndpointServiceAdlResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPrivatelinkEndpointServiceAdlResult)(nil)).Elem()
}

func (o LookupPrivatelinkEndpointServiceAdlResultOutput) ToLookupPrivatelinkEndpointServiceAdlResultOutput() LookupPrivatelinkEndpointServiceAdlResultOutput {
	return o
}

func (o LookupPrivatelinkEndpointServiceAdlResultOutput) ToLookupPrivatelinkEndpointServiceAdlResultOutputWithContext(ctx context.Context) LookupPrivatelinkEndpointServiceAdlResultOutput {
	return o
}

// Human-readable string to associate with this private endpoint.
func (o LookupPrivatelinkEndpointServiceAdlResultOutput) Comment() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivatelinkEndpointServiceAdlResult) string { return v.Comment }).(pulumi.StringOutput)
}

func (o LookupPrivatelinkEndpointServiceAdlResultOutput) EndpointId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivatelinkEndpointServiceAdlResult) string { return v.EndpointId }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupPrivatelinkEndpointServiceAdlResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivatelinkEndpointServiceAdlResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupPrivatelinkEndpointServiceAdlResultOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivatelinkEndpointServiceAdlResult) string { return v.ProjectId }).(pulumi.StringOutput)
}

// Human-readable label that identifies the cloud provider for this endpoint.
func (o LookupPrivatelinkEndpointServiceAdlResultOutput) ProviderName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivatelinkEndpointServiceAdlResult) string { return v.ProviderName }).(pulumi.StringOutput)
}

// Human-readable label that identifies the type of resource to associate with this private endpoint.
func (o LookupPrivatelinkEndpointServiceAdlResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivatelinkEndpointServiceAdlResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPrivatelinkEndpointServiceAdlResultOutput{})
}
