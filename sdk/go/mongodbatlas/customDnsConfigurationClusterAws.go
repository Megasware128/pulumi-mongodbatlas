// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mongodbatlas

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// `CustomDnsConfigurationClusterAws` provides a Custom DNS Configuration for Atlas Clusters on AWS resource. This represents a Custom DNS Configuration for Atlas Clusters on AWS that can be updated in an Atlas project.
//
// > **IMPORTANT:**You must have one of the following roles to successfully handle the resource:
//   - Organization Owner
//   - Project Owner
//
// > **NOTE:** Groups and projects are synonymous terms. You may find groupId in the official documentation.
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
//			_, err := mongodbatlas.NewCustomDnsConfigurationClusterAws(ctx, "test", &mongodbatlas.CustomDnsConfigurationClusterAwsArgs{
//				Enabled:   pulumi.Bool(true),
//				ProjectId: pulumi.String("<PROJECT-ID>"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// Custom DNS Configuration for Atlas Clusters on AWS must be imported using auditing ID, e.g.
//
// ```sh
//
//	$ pulumi import mongodbatlas:index/customDnsConfigurationClusterAws:CustomDnsConfigurationClusterAws test 1112222b3bf99403840e8934
//
// ```
//
//	See detailed information for arguments and attributes[MongoDB API Custom DNS Configuration for Atlas Clusters on AWS](https://docs.atlas.mongodb.com/reference/api/aws-custom-dns)
type CustomDnsConfigurationClusterAws struct {
	pulumi.CustomResourceState

	// Indicates whether the project's clusters deployed to AWS use custom DNS. If `true`, the `Get All Clusters` and `Get One Cluster` endpoints return the `connectionStrings.private` and `connectionStrings.privateSrv` fields for clusters deployed to AWS .
	Enabled pulumi.BoolOutput `pulumi:"enabled"`
	// Required 	Unique identifier for the project.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
}

// NewCustomDnsConfigurationClusterAws registers a new resource with the given unique name, arguments, and options.
func NewCustomDnsConfigurationClusterAws(ctx *pulumi.Context,
	name string, args *CustomDnsConfigurationClusterAwsArgs, opts ...pulumi.ResourceOption) (*CustomDnsConfigurationClusterAws, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Enabled == nil {
		return nil, errors.New("invalid value for required argument 'Enabled'")
	}
	if args.ProjectId == nil {
		return nil, errors.New("invalid value for required argument 'ProjectId'")
	}
	var resource CustomDnsConfigurationClusterAws
	err := ctx.RegisterResource("mongodbatlas:index/customDnsConfigurationClusterAws:CustomDnsConfigurationClusterAws", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCustomDnsConfigurationClusterAws gets an existing CustomDnsConfigurationClusterAws resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCustomDnsConfigurationClusterAws(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CustomDnsConfigurationClusterAwsState, opts ...pulumi.ResourceOption) (*CustomDnsConfigurationClusterAws, error) {
	var resource CustomDnsConfigurationClusterAws
	err := ctx.ReadResource("mongodbatlas:index/customDnsConfigurationClusterAws:CustomDnsConfigurationClusterAws", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CustomDnsConfigurationClusterAws resources.
type customDnsConfigurationClusterAwsState struct {
	// Indicates whether the project's clusters deployed to AWS use custom DNS. If `true`, the `Get All Clusters` and `Get One Cluster` endpoints return the `connectionStrings.private` and `connectionStrings.privateSrv` fields for clusters deployed to AWS .
	Enabled *bool `pulumi:"enabled"`
	// Required 	Unique identifier for the project.
	ProjectId *string `pulumi:"projectId"`
}

type CustomDnsConfigurationClusterAwsState struct {
	// Indicates whether the project's clusters deployed to AWS use custom DNS. If `true`, the `Get All Clusters` and `Get One Cluster` endpoints return the `connectionStrings.private` and `connectionStrings.privateSrv` fields for clusters deployed to AWS .
	Enabled pulumi.BoolPtrInput
	// Required 	Unique identifier for the project.
	ProjectId pulumi.StringPtrInput
}

func (CustomDnsConfigurationClusterAwsState) ElementType() reflect.Type {
	return reflect.TypeOf((*customDnsConfigurationClusterAwsState)(nil)).Elem()
}

type customDnsConfigurationClusterAwsArgs struct {
	// Indicates whether the project's clusters deployed to AWS use custom DNS. If `true`, the `Get All Clusters` and `Get One Cluster` endpoints return the `connectionStrings.private` and `connectionStrings.privateSrv` fields for clusters deployed to AWS .
	Enabled bool `pulumi:"enabled"`
	// Required 	Unique identifier for the project.
	ProjectId string `pulumi:"projectId"`
}

// The set of arguments for constructing a CustomDnsConfigurationClusterAws resource.
type CustomDnsConfigurationClusterAwsArgs struct {
	// Indicates whether the project's clusters deployed to AWS use custom DNS. If `true`, the `Get All Clusters` and `Get One Cluster` endpoints return the `connectionStrings.private` and `connectionStrings.privateSrv` fields for clusters deployed to AWS .
	Enabled pulumi.BoolInput
	// Required 	Unique identifier for the project.
	ProjectId pulumi.StringInput
}

func (CustomDnsConfigurationClusterAwsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*customDnsConfigurationClusterAwsArgs)(nil)).Elem()
}

type CustomDnsConfigurationClusterAwsInput interface {
	pulumi.Input

	ToCustomDnsConfigurationClusterAwsOutput() CustomDnsConfigurationClusterAwsOutput
	ToCustomDnsConfigurationClusterAwsOutputWithContext(ctx context.Context) CustomDnsConfigurationClusterAwsOutput
}

func (*CustomDnsConfigurationClusterAws) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomDnsConfigurationClusterAws)(nil)).Elem()
}

func (i *CustomDnsConfigurationClusterAws) ToCustomDnsConfigurationClusterAwsOutput() CustomDnsConfigurationClusterAwsOutput {
	return i.ToCustomDnsConfigurationClusterAwsOutputWithContext(context.Background())
}

func (i *CustomDnsConfigurationClusterAws) ToCustomDnsConfigurationClusterAwsOutputWithContext(ctx context.Context) CustomDnsConfigurationClusterAwsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomDnsConfigurationClusterAwsOutput)
}

// CustomDnsConfigurationClusterAwsArrayInput is an input type that accepts CustomDnsConfigurationClusterAwsArray and CustomDnsConfigurationClusterAwsArrayOutput values.
// You can construct a concrete instance of `CustomDnsConfigurationClusterAwsArrayInput` via:
//
//	CustomDnsConfigurationClusterAwsArray{ CustomDnsConfigurationClusterAwsArgs{...} }
type CustomDnsConfigurationClusterAwsArrayInput interface {
	pulumi.Input

	ToCustomDnsConfigurationClusterAwsArrayOutput() CustomDnsConfigurationClusterAwsArrayOutput
	ToCustomDnsConfigurationClusterAwsArrayOutputWithContext(context.Context) CustomDnsConfigurationClusterAwsArrayOutput
}

type CustomDnsConfigurationClusterAwsArray []CustomDnsConfigurationClusterAwsInput

func (CustomDnsConfigurationClusterAwsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CustomDnsConfigurationClusterAws)(nil)).Elem()
}

func (i CustomDnsConfigurationClusterAwsArray) ToCustomDnsConfigurationClusterAwsArrayOutput() CustomDnsConfigurationClusterAwsArrayOutput {
	return i.ToCustomDnsConfigurationClusterAwsArrayOutputWithContext(context.Background())
}

func (i CustomDnsConfigurationClusterAwsArray) ToCustomDnsConfigurationClusterAwsArrayOutputWithContext(ctx context.Context) CustomDnsConfigurationClusterAwsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomDnsConfigurationClusterAwsArrayOutput)
}

// CustomDnsConfigurationClusterAwsMapInput is an input type that accepts CustomDnsConfigurationClusterAwsMap and CustomDnsConfigurationClusterAwsMapOutput values.
// You can construct a concrete instance of `CustomDnsConfigurationClusterAwsMapInput` via:
//
//	CustomDnsConfigurationClusterAwsMap{ "key": CustomDnsConfigurationClusterAwsArgs{...} }
type CustomDnsConfigurationClusterAwsMapInput interface {
	pulumi.Input

	ToCustomDnsConfigurationClusterAwsMapOutput() CustomDnsConfigurationClusterAwsMapOutput
	ToCustomDnsConfigurationClusterAwsMapOutputWithContext(context.Context) CustomDnsConfigurationClusterAwsMapOutput
}

type CustomDnsConfigurationClusterAwsMap map[string]CustomDnsConfigurationClusterAwsInput

func (CustomDnsConfigurationClusterAwsMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CustomDnsConfigurationClusterAws)(nil)).Elem()
}

func (i CustomDnsConfigurationClusterAwsMap) ToCustomDnsConfigurationClusterAwsMapOutput() CustomDnsConfigurationClusterAwsMapOutput {
	return i.ToCustomDnsConfigurationClusterAwsMapOutputWithContext(context.Background())
}

func (i CustomDnsConfigurationClusterAwsMap) ToCustomDnsConfigurationClusterAwsMapOutputWithContext(ctx context.Context) CustomDnsConfigurationClusterAwsMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomDnsConfigurationClusterAwsMapOutput)
}

type CustomDnsConfigurationClusterAwsOutput struct{ *pulumi.OutputState }

func (CustomDnsConfigurationClusterAwsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomDnsConfigurationClusterAws)(nil)).Elem()
}

func (o CustomDnsConfigurationClusterAwsOutput) ToCustomDnsConfigurationClusterAwsOutput() CustomDnsConfigurationClusterAwsOutput {
	return o
}

func (o CustomDnsConfigurationClusterAwsOutput) ToCustomDnsConfigurationClusterAwsOutputWithContext(ctx context.Context) CustomDnsConfigurationClusterAwsOutput {
	return o
}

// Indicates whether the project's clusters deployed to AWS use custom DNS. If `true`, the `Get All Clusters` and `Get One Cluster` endpoints return the `connectionStrings.private` and `connectionStrings.privateSrv` fields for clusters deployed to AWS .
func (o CustomDnsConfigurationClusterAwsOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *CustomDnsConfigurationClusterAws) pulumi.BoolOutput { return v.Enabled }).(pulumi.BoolOutput)
}

// Required 	Unique identifier for the project.
func (o CustomDnsConfigurationClusterAwsOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *CustomDnsConfigurationClusterAws) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

type CustomDnsConfigurationClusterAwsArrayOutput struct{ *pulumi.OutputState }

func (CustomDnsConfigurationClusterAwsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CustomDnsConfigurationClusterAws)(nil)).Elem()
}

func (o CustomDnsConfigurationClusterAwsArrayOutput) ToCustomDnsConfigurationClusterAwsArrayOutput() CustomDnsConfigurationClusterAwsArrayOutput {
	return o
}

func (o CustomDnsConfigurationClusterAwsArrayOutput) ToCustomDnsConfigurationClusterAwsArrayOutputWithContext(ctx context.Context) CustomDnsConfigurationClusterAwsArrayOutput {
	return o
}

func (o CustomDnsConfigurationClusterAwsArrayOutput) Index(i pulumi.IntInput) CustomDnsConfigurationClusterAwsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *CustomDnsConfigurationClusterAws {
		return vs[0].([]*CustomDnsConfigurationClusterAws)[vs[1].(int)]
	}).(CustomDnsConfigurationClusterAwsOutput)
}

type CustomDnsConfigurationClusterAwsMapOutput struct{ *pulumi.OutputState }

func (CustomDnsConfigurationClusterAwsMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CustomDnsConfigurationClusterAws)(nil)).Elem()
}

func (o CustomDnsConfigurationClusterAwsMapOutput) ToCustomDnsConfigurationClusterAwsMapOutput() CustomDnsConfigurationClusterAwsMapOutput {
	return o
}

func (o CustomDnsConfigurationClusterAwsMapOutput) ToCustomDnsConfigurationClusterAwsMapOutputWithContext(ctx context.Context) CustomDnsConfigurationClusterAwsMapOutput {
	return o
}

func (o CustomDnsConfigurationClusterAwsMapOutput) MapIndex(k pulumi.StringInput) CustomDnsConfigurationClusterAwsOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *CustomDnsConfigurationClusterAws {
		return vs[0].(map[string]*CustomDnsConfigurationClusterAws)[vs[1].(string)]
	}).(CustomDnsConfigurationClusterAwsOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CustomDnsConfigurationClusterAwsInput)(nil)).Elem(), &CustomDnsConfigurationClusterAws{})
	pulumi.RegisterInputType(reflect.TypeOf((*CustomDnsConfigurationClusterAwsArrayInput)(nil)).Elem(), CustomDnsConfigurationClusterAwsArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CustomDnsConfigurationClusterAwsMapInput)(nil)).Elem(), CustomDnsConfigurationClusterAwsMap{})
	pulumi.RegisterOutputType(CustomDnsConfigurationClusterAwsOutput{})
	pulumi.RegisterOutputType(CustomDnsConfigurationClusterAwsArrayOutput{})
	pulumi.RegisterOutputType(CustomDnsConfigurationClusterAwsMapOutput{})
}
