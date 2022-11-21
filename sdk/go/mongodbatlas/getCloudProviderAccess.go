// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mongodbatlas

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// `CloudProviderAccess` allows you to get the list of cloud provider access roles, currently only AWS is supported.
//
// > **NOTE:** Groups and projects are synonymous terms. You may find `groupId` in the official documentation.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-mongodbatlas/sdk/go/mongodbatlas"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			testRole, err := mongodbatlas.NewCloudProviderAccess(ctx, "testRole", &mongodbatlas.CloudProviderAccessArgs{
//				ProjectId:    pulumi.String("<PROJECT-ID>"),
//				ProviderName: pulumi.String("AWS"),
//			})
//			if err != nil {
//				return err
//			}
//			_ = mongodbatlas.LookupCloudProviderAccessOutput(ctx, GetCloudProviderAccessOutputArgs{
//				ProjectId: testRole.ProjectId,
//			}, nil)
//			return nil
//		})
//	}
//
// ```
func LookupCloudProviderAccess(ctx *pulumi.Context, args *LookupCloudProviderAccessArgs, opts ...pulumi.InvokeOption) (*LookupCloudProviderAccessResult, error) {
	var rv LookupCloudProviderAccessResult
	err := ctx.Invoke("mongodbatlas:index/getCloudProviderAccess:getCloudProviderAccess", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCloudProviderAccess.
type LookupCloudProviderAccessArgs struct {
	// The unique ID for the project to get all Cloud Provider Access
	ProjectId string `pulumi:"projectId"`
}

// A collection of values returned by getCloudProviderAccess.
type LookupCloudProviderAccessResult struct {
	// A list where each represents a Cloud Provider Access Role.
	AwsIamRoles []GetCloudProviderAccessAwsIamRole `pulumi:"awsIamRoles"`
	// The provider-assigned unique ID for this managed resource.
	Id        string `pulumi:"id"`
	ProjectId string `pulumi:"projectId"`
}

func LookupCloudProviderAccessOutput(ctx *pulumi.Context, args LookupCloudProviderAccessOutputArgs, opts ...pulumi.InvokeOption) LookupCloudProviderAccessResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupCloudProviderAccessResult, error) {
			args := v.(LookupCloudProviderAccessArgs)
			r, err := LookupCloudProviderAccess(ctx, &args, opts...)
			var s LookupCloudProviderAccessResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupCloudProviderAccessResultOutput)
}

// A collection of arguments for invoking getCloudProviderAccess.
type LookupCloudProviderAccessOutputArgs struct {
	// The unique ID for the project to get all Cloud Provider Access
	ProjectId pulumi.StringInput `pulumi:"projectId"`
}

func (LookupCloudProviderAccessOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCloudProviderAccessArgs)(nil)).Elem()
}

// A collection of values returned by getCloudProviderAccess.
type LookupCloudProviderAccessResultOutput struct{ *pulumi.OutputState }

func (LookupCloudProviderAccessResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCloudProviderAccessResult)(nil)).Elem()
}

func (o LookupCloudProviderAccessResultOutput) ToLookupCloudProviderAccessResultOutput() LookupCloudProviderAccessResultOutput {
	return o
}

func (o LookupCloudProviderAccessResultOutput) ToLookupCloudProviderAccessResultOutputWithContext(ctx context.Context) LookupCloudProviderAccessResultOutput {
	return o
}

// A list where each represents a Cloud Provider Access Role.
func (o LookupCloudProviderAccessResultOutput) AwsIamRoles() GetCloudProviderAccessAwsIamRoleArrayOutput {
	return o.ApplyT(func(v LookupCloudProviderAccessResult) []GetCloudProviderAccessAwsIamRole { return v.AwsIamRoles }).(GetCloudProviderAccessAwsIamRoleArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupCloudProviderAccessResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudProviderAccessResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupCloudProviderAccessResultOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudProviderAccessResult) string { return v.ProjectId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCloudProviderAccessResultOutput{})
}
