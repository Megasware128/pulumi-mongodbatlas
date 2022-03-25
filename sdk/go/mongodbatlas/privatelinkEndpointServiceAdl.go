// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package mongodbatlas

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// `privatelinkEndpointServiceAdl` Provides an Atlas Data Lake (ADL) and Online Archive PrivateLink endpoint resource.   The same configuration will provide a PrivateLink connection for either Atlas Data Lake or Online Archive.
//
// ## Example Usage
// ### Basic
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
// 		_, err := mongodbatlas.NewPrivatelinkEndpointServiceAdl(ctx, "adlTest", &mongodbatlas.PrivatelinkEndpointServiceAdlArgs{
// 			Comment:      pulumi.String("comments for private link endpoint adl"),
// 			EndpointId:   pulumi.String("<ENDPOINT_ID>"),
// 			ProjectId:    pulumi.String("<PROJECT_ID>"),
// 			ProviderName: pulumi.String("AWS"),
// 			Type:         pulumi.String("DATA_LAKE"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// ADL privatelink endpoint can be imported using project ID and endpoint ID, in the format `project_id`--`endpoint_id`, e.g.
//
// ```sh
//  $ pulumi import mongodbatlas:index/privatelinkEndpointServiceAdl:PrivatelinkEndpointServiceAdl test 1112222b3bf99403840e8934--vpce-jjg5e24qp93513h03
// ```
//
//  For more information see[MongoDB Atlas API - DataLake](https://docs.mongodb.com/datalake/reference/api/datalakes-api/)
//
// and [MongoDB Atlas API - Online Archive](https://docs.atlas.mongodb.com/reference/api/online-archive/) Documentation.
type PrivatelinkEndpointServiceAdl struct {
	pulumi.CustomResourceState

	// Human-readable string to associate with this private endpoint.
	Comment pulumi.StringPtrOutput `pulumi:"comment"`
	// Unique 22-character alphanumeric string that identifies the private endpoint. Atlas supports AWS private endpoints using the [|aws| PrivateLink](https://aws.amazon.com/privatelink/) feature.
	EndpointId pulumi.StringOutput `pulumi:"endpointId"`
	// Unique 24-digit hexadecimal string that identifies the project.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// Human-readable label that identifies the cloud provider for this endpoint. Atlas supports AWS only. If empty, defaults to AWS.
	ProviderName pulumi.StringOutput `pulumi:"providerName"`
	// Human-readable label that identifies the type of resource to associate with this private endpoint. Atlas supports `DATA_LAKE` only. If empty, defaults to `DATA_LAKE`.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewPrivatelinkEndpointServiceAdl registers a new resource with the given unique name, arguments, and options.
func NewPrivatelinkEndpointServiceAdl(ctx *pulumi.Context,
	name string, args *PrivatelinkEndpointServiceAdlArgs, opts ...pulumi.ResourceOption) (*PrivatelinkEndpointServiceAdl, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EndpointId == nil {
		return nil, errors.New("invalid value for required argument 'EndpointId'")
	}
	if args.ProjectId == nil {
		return nil, errors.New("invalid value for required argument 'ProjectId'")
	}
	if args.ProviderName == nil {
		return nil, errors.New("invalid value for required argument 'ProviderName'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	var resource PrivatelinkEndpointServiceAdl
	err := ctx.RegisterResource("mongodbatlas:index/privatelinkEndpointServiceAdl:PrivatelinkEndpointServiceAdl", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPrivatelinkEndpointServiceAdl gets an existing PrivatelinkEndpointServiceAdl resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPrivatelinkEndpointServiceAdl(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PrivatelinkEndpointServiceAdlState, opts ...pulumi.ResourceOption) (*PrivatelinkEndpointServiceAdl, error) {
	var resource PrivatelinkEndpointServiceAdl
	err := ctx.ReadResource("mongodbatlas:index/privatelinkEndpointServiceAdl:PrivatelinkEndpointServiceAdl", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PrivatelinkEndpointServiceAdl resources.
type privatelinkEndpointServiceAdlState struct {
	// Human-readable string to associate with this private endpoint.
	Comment *string `pulumi:"comment"`
	// Unique 22-character alphanumeric string that identifies the private endpoint. Atlas supports AWS private endpoints using the [|aws| PrivateLink](https://aws.amazon.com/privatelink/) feature.
	EndpointId *string `pulumi:"endpointId"`
	// Unique 24-digit hexadecimal string that identifies the project.
	ProjectId *string `pulumi:"projectId"`
	// Human-readable label that identifies the cloud provider for this endpoint. Atlas supports AWS only. If empty, defaults to AWS.
	ProviderName *string `pulumi:"providerName"`
	// Human-readable label that identifies the type of resource to associate with this private endpoint. Atlas supports `DATA_LAKE` only. If empty, defaults to `DATA_LAKE`.
	Type *string `pulumi:"type"`
}

type PrivatelinkEndpointServiceAdlState struct {
	// Human-readable string to associate with this private endpoint.
	Comment pulumi.StringPtrInput
	// Unique 22-character alphanumeric string that identifies the private endpoint. Atlas supports AWS private endpoints using the [|aws| PrivateLink](https://aws.amazon.com/privatelink/) feature.
	EndpointId pulumi.StringPtrInput
	// Unique 24-digit hexadecimal string that identifies the project.
	ProjectId pulumi.StringPtrInput
	// Human-readable label that identifies the cloud provider for this endpoint. Atlas supports AWS only. If empty, defaults to AWS.
	ProviderName pulumi.StringPtrInput
	// Human-readable label that identifies the type of resource to associate with this private endpoint. Atlas supports `DATA_LAKE` only. If empty, defaults to `DATA_LAKE`.
	Type pulumi.StringPtrInput
}

func (PrivatelinkEndpointServiceAdlState) ElementType() reflect.Type {
	return reflect.TypeOf((*privatelinkEndpointServiceAdlState)(nil)).Elem()
}

type privatelinkEndpointServiceAdlArgs struct {
	// Human-readable string to associate with this private endpoint.
	Comment *string `pulumi:"comment"`
	// Unique 22-character alphanumeric string that identifies the private endpoint. Atlas supports AWS private endpoints using the [|aws| PrivateLink](https://aws.amazon.com/privatelink/) feature.
	EndpointId string `pulumi:"endpointId"`
	// Unique 24-digit hexadecimal string that identifies the project.
	ProjectId string `pulumi:"projectId"`
	// Human-readable label that identifies the cloud provider for this endpoint. Atlas supports AWS only. If empty, defaults to AWS.
	ProviderName string `pulumi:"providerName"`
	// Human-readable label that identifies the type of resource to associate with this private endpoint. Atlas supports `DATA_LAKE` only. If empty, defaults to `DATA_LAKE`.
	Type string `pulumi:"type"`
}

// The set of arguments for constructing a PrivatelinkEndpointServiceAdl resource.
type PrivatelinkEndpointServiceAdlArgs struct {
	// Human-readable string to associate with this private endpoint.
	Comment pulumi.StringPtrInput
	// Unique 22-character alphanumeric string that identifies the private endpoint. Atlas supports AWS private endpoints using the [|aws| PrivateLink](https://aws.amazon.com/privatelink/) feature.
	EndpointId pulumi.StringInput
	// Unique 24-digit hexadecimal string that identifies the project.
	ProjectId pulumi.StringInput
	// Human-readable label that identifies the cloud provider for this endpoint. Atlas supports AWS only. If empty, defaults to AWS.
	ProviderName pulumi.StringInput
	// Human-readable label that identifies the type of resource to associate with this private endpoint. Atlas supports `DATA_LAKE` only. If empty, defaults to `DATA_LAKE`.
	Type pulumi.StringInput
}

func (PrivatelinkEndpointServiceAdlArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*privatelinkEndpointServiceAdlArgs)(nil)).Elem()
}

type PrivatelinkEndpointServiceAdlInput interface {
	pulumi.Input

	ToPrivatelinkEndpointServiceAdlOutput() PrivatelinkEndpointServiceAdlOutput
	ToPrivatelinkEndpointServiceAdlOutputWithContext(ctx context.Context) PrivatelinkEndpointServiceAdlOutput
}

func (*PrivatelinkEndpointServiceAdl) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivatelinkEndpointServiceAdl)(nil)).Elem()
}

func (i *PrivatelinkEndpointServiceAdl) ToPrivatelinkEndpointServiceAdlOutput() PrivatelinkEndpointServiceAdlOutput {
	return i.ToPrivatelinkEndpointServiceAdlOutputWithContext(context.Background())
}

func (i *PrivatelinkEndpointServiceAdl) ToPrivatelinkEndpointServiceAdlOutputWithContext(ctx context.Context) PrivatelinkEndpointServiceAdlOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivatelinkEndpointServiceAdlOutput)
}

// PrivatelinkEndpointServiceAdlArrayInput is an input type that accepts PrivatelinkEndpointServiceAdlArray and PrivatelinkEndpointServiceAdlArrayOutput values.
// You can construct a concrete instance of `PrivatelinkEndpointServiceAdlArrayInput` via:
//
//          PrivatelinkEndpointServiceAdlArray{ PrivatelinkEndpointServiceAdlArgs{...} }
type PrivatelinkEndpointServiceAdlArrayInput interface {
	pulumi.Input

	ToPrivatelinkEndpointServiceAdlArrayOutput() PrivatelinkEndpointServiceAdlArrayOutput
	ToPrivatelinkEndpointServiceAdlArrayOutputWithContext(context.Context) PrivatelinkEndpointServiceAdlArrayOutput
}

type PrivatelinkEndpointServiceAdlArray []PrivatelinkEndpointServiceAdlInput

func (PrivatelinkEndpointServiceAdlArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PrivatelinkEndpointServiceAdl)(nil)).Elem()
}

func (i PrivatelinkEndpointServiceAdlArray) ToPrivatelinkEndpointServiceAdlArrayOutput() PrivatelinkEndpointServiceAdlArrayOutput {
	return i.ToPrivatelinkEndpointServiceAdlArrayOutputWithContext(context.Background())
}

func (i PrivatelinkEndpointServiceAdlArray) ToPrivatelinkEndpointServiceAdlArrayOutputWithContext(ctx context.Context) PrivatelinkEndpointServiceAdlArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivatelinkEndpointServiceAdlArrayOutput)
}

// PrivatelinkEndpointServiceAdlMapInput is an input type that accepts PrivatelinkEndpointServiceAdlMap and PrivatelinkEndpointServiceAdlMapOutput values.
// You can construct a concrete instance of `PrivatelinkEndpointServiceAdlMapInput` via:
//
//          PrivatelinkEndpointServiceAdlMap{ "key": PrivatelinkEndpointServiceAdlArgs{...} }
type PrivatelinkEndpointServiceAdlMapInput interface {
	pulumi.Input

	ToPrivatelinkEndpointServiceAdlMapOutput() PrivatelinkEndpointServiceAdlMapOutput
	ToPrivatelinkEndpointServiceAdlMapOutputWithContext(context.Context) PrivatelinkEndpointServiceAdlMapOutput
}

type PrivatelinkEndpointServiceAdlMap map[string]PrivatelinkEndpointServiceAdlInput

func (PrivatelinkEndpointServiceAdlMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PrivatelinkEndpointServiceAdl)(nil)).Elem()
}

func (i PrivatelinkEndpointServiceAdlMap) ToPrivatelinkEndpointServiceAdlMapOutput() PrivatelinkEndpointServiceAdlMapOutput {
	return i.ToPrivatelinkEndpointServiceAdlMapOutputWithContext(context.Background())
}

func (i PrivatelinkEndpointServiceAdlMap) ToPrivatelinkEndpointServiceAdlMapOutputWithContext(ctx context.Context) PrivatelinkEndpointServiceAdlMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivatelinkEndpointServiceAdlMapOutput)
}

type PrivatelinkEndpointServiceAdlOutput struct{ *pulumi.OutputState }

func (PrivatelinkEndpointServiceAdlOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivatelinkEndpointServiceAdl)(nil)).Elem()
}

func (o PrivatelinkEndpointServiceAdlOutput) ToPrivatelinkEndpointServiceAdlOutput() PrivatelinkEndpointServiceAdlOutput {
	return o
}

func (o PrivatelinkEndpointServiceAdlOutput) ToPrivatelinkEndpointServiceAdlOutputWithContext(ctx context.Context) PrivatelinkEndpointServiceAdlOutput {
	return o
}

type PrivatelinkEndpointServiceAdlArrayOutput struct{ *pulumi.OutputState }

func (PrivatelinkEndpointServiceAdlArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PrivatelinkEndpointServiceAdl)(nil)).Elem()
}

func (o PrivatelinkEndpointServiceAdlArrayOutput) ToPrivatelinkEndpointServiceAdlArrayOutput() PrivatelinkEndpointServiceAdlArrayOutput {
	return o
}

func (o PrivatelinkEndpointServiceAdlArrayOutput) ToPrivatelinkEndpointServiceAdlArrayOutputWithContext(ctx context.Context) PrivatelinkEndpointServiceAdlArrayOutput {
	return o
}

func (o PrivatelinkEndpointServiceAdlArrayOutput) Index(i pulumi.IntInput) PrivatelinkEndpointServiceAdlOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *PrivatelinkEndpointServiceAdl {
		return vs[0].([]*PrivatelinkEndpointServiceAdl)[vs[1].(int)]
	}).(PrivatelinkEndpointServiceAdlOutput)
}

type PrivatelinkEndpointServiceAdlMapOutput struct{ *pulumi.OutputState }

func (PrivatelinkEndpointServiceAdlMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PrivatelinkEndpointServiceAdl)(nil)).Elem()
}

func (o PrivatelinkEndpointServiceAdlMapOutput) ToPrivatelinkEndpointServiceAdlMapOutput() PrivatelinkEndpointServiceAdlMapOutput {
	return o
}

func (o PrivatelinkEndpointServiceAdlMapOutput) ToPrivatelinkEndpointServiceAdlMapOutputWithContext(ctx context.Context) PrivatelinkEndpointServiceAdlMapOutput {
	return o
}

func (o PrivatelinkEndpointServiceAdlMapOutput) MapIndex(k pulumi.StringInput) PrivatelinkEndpointServiceAdlOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *PrivatelinkEndpointServiceAdl {
		return vs[0].(map[string]*PrivatelinkEndpointServiceAdl)[vs[1].(string)]
	}).(PrivatelinkEndpointServiceAdlOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PrivatelinkEndpointServiceAdlInput)(nil)).Elem(), &PrivatelinkEndpointServiceAdl{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrivatelinkEndpointServiceAdlArrayInput)(nil)).Elem(), PrivatelinkEndpointServiceAdlArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrivatelinkEndpointServiceAdlMapInput)(nil)).Elem(), PrivatelinkEndpointServiceAdlMap{})
	pulumi.RegisterOutputType(PrivatelinkEndpointServiceAdlOutput{})
	pulumi.RegisterOutputType(PrivatelinkEndpointServiceAdlArrayOutput{})
	pulumi.RegisterOutputType(PrivatelinkEndpointServiceAdlMapOutput{})
}