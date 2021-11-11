// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package mongodbatlas

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// `getNetworkContainers` describes all Network Peering Containers. The data source requires your Project ID.
//
// > **NOTE:** Groups and projects are synonymous terms. You may find **group_id** in the official documentation.
//
// ## Example Usage
// ### Basic Example.
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-mongodbatlas/sdk/v3/go/mongodbatlas"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		testNetworkContainer, err := mongodbatlas.NewNetworkContainer(ctx, "testNetworkContainer", &mongodbatlas.NetworkContainerArgs{
// 			ProjectId:      pulumi.String("<YOUR-PROJECT-ID>"),
// 			AtlasCidrBlock: pulumi.String("10.8.0.0/21"),
// 			ProviderName:   pulumi.String("AWS"),
// 			RegionName:     pulumi.String("US_EAST_1"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupNetworkContainers(ctx *pulumi.Context, args *LookupNetworkContainersArgs, opts ...pulumi.InvokeOption) (*LookupNetworkContainersResult, error) {
	var rv LookupNetworkContainersResult
	err := ctx.Invoke("mongodbatlas:index/getNetworkContainers:getNetworkContainers", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getNetworkContainers.
type LookupNetworkContainersArgs struct {
	// The unique ID for the project to create the database user.
	ProjectId string `pulumi:"projectId"`
	// Cloud provider for this Network peering container. Accepted values are AWS, GCP, and Azure.
	ProviderName string `pulumi:"providerName"`
}

// A collection of values returned by getNetworkContainers.
type LookupNetworkContainersResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id        string `pulumi:"id"`
	ProjectId string `pulumi:"projectId"`
	// Cloud provider for this Network Peering connection. If omitted, Atlas sets this parameter to AWS.
	ProviderName string `pulumi:"providerName"`
	// A list where each represents a Network Peering Container.
	Results []GetNetworkContainersResult `pulumi:"results"`
}

func LookupNetworkContainersOutput(ctx *pulumi.Context, args LookupNetworkContainersOutputArgs, opts ...pulumi.InvokeOption) LookupNetworkContainersResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupNetworkContainersResult, error) {
			args := v.(LookupNetworkContainersArgs)
			r, err := LookupNetworkContainers(ctx, &args, opts...)
			return *r, err
		}).(LookupNetworkContainersResultOutput)
}

// A collection of arguments for invoking getNetworkContainers.
type LookupNetworkContainersOutputArgs struct {
	// The unique ID for the project to create the database user.
	ProjectId pulumi.StringInput `pulumi:"projectId"`
	// Cloud provider for this Network peering container. Accepted values are AWS, GCP, and Azure.
	ProviderName pulumi.StringInput `pulumi:"providerName"`
}

func (LookupNetworkContainersOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNetworkContainersArgs)(nil)).Elem()
}

// A collection of values returned by getNetworkContainers.
type LookupNetworkContainersResultOutput struct{ *pulumi.OutputState }

func (LookupNetworkContainersResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNetworkContainersResult)(nil)).Elem()
}

func (o LookupNetworkContainersResultOutput) ToLookupNetworkContainersResultOutput() LookupNetworkContainersResultOutput {
	return o
}

func (o LookupNetworkContainersResultOutput) ToLookupNetworkContainersResultOutputWithContext(ctx context.Context) LookupNetworkContainersResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o LookupNetworkContainersResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkContainersResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupNetworkContainersResultOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkContainersResult) string { return v.ProjectId }).(pulumi.StringOutput)
}

// Cloud provider for this Network Peering connection. If omitted, Atlas sets this parameter to AWS.
func (o LookupNetworkContainersResultOutput) ProviderName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkContainersResult) string { return v.ProviderName }).(pulumi.StringOutput)
}

// A list where each represents a Network Peering Container.
func (o LookupNetworkContainersResultOutput) Results() GetNetworkContainersResultArrayOutput {
	return o.ApplyT(func(v LookupNetworkContainersResult) []GetNetworkContainersResult { return v.Results }).(GetNetworkContainersResultArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNetworkContainersResultOutput{})
}
