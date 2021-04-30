// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package mongodbatlas

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Import
//
// The Cloud Provider Access resource can be imported using project ID and the provider name and mongodbatlas role id, in the format `project_id`-`provider_name`-`role_id`, e.g.
//
// ```sh
//  $ pulumi import mongodbatlas:index/cloudProviderAccess:CloudProviderAccess my_role 1112222b3bf99403840e8934-AWS-5fc17d476f7a33224f5b224e
// ```
type CloudProviderAccess struct {
	pulumi.CustomResourceState

	// Unique external ID Atlas uses when assuming the IAM role in your AWS account.
	AtlasAssumedRoleExternalId pulumi.StringOutput `pulumi:"atlasAssumedRoleExternalId"`
	// ARN associated with the Atlas AWS account used to assume IAM roles in your AWS account.
	AtlasAwsAccountArn pulumi.StringOutput `pulumi:"atlasAwsAccountArn"`
	// Date on which this role was authorized.
	AuthorizedDate pulumi.StringOutput `pulumi:"authorizedDate"`
	// Date on which this role was created.
	CreatedDate pulumi.StringOutput `pulumi:"createdDate"`
	// Atlas features this AWS IAM role is linked to.
	FeatureUsages CloudProviderAccessFeatureUsageArrayOutput `pulumi:"featureUsages"`
	// ARN of the IAM Role that Atlas assumes when accessing resources in your AWS account. This value is required after the creation (register of the role) as part of [Set Up Unified AWS Access](https://docs.atlas.mongodb.com/security/set-up-unified-aws-access/#set-up-unified-aws-access).
	IamAssumedRoleArn pulumi.StringPtrOutput `pulumi:"iamAssumedRoleArn"`
	// The unique ID for the project
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// The cloud provider for which to create a new role. Currently only AWS is supported.
	ProviderName pulumi.StringOutput `pulumi:"providerName"`
	// Unique ID of this role returned by mongodb atlas api
	RoleId pulumi.StringOutput `pulumi:"roleId"`
}

// NewCloudProviderAccess registers a new resource with the given unique name, arguments, and options.
func NewCloudProviderAccess(ctx *pulumi.Context,
	name string, args *CloudProviderAccessArgs, opts ...pulumi.ResourceOption) (*CloudProviderAccess, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ProjectId == nil {
		return nil, errors.New("invalid value for required argument 'ProjectId'")
	}
	if args.ProviderName == nil {
		return nil, errors.New("invalid value for required argument 'ProviderName'")
	}
	var resource CloudProviderAccess
	err := ctx.RegisterResource("mongodbatlas:index/cloudProviderAccess:CloudProviderAccess", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCloudProviderAccess gets an existing CloudProviderAccess resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCloudProviderAccess(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CloudProviderAccessState, opts ...pulumi.ResourceOption) (*CloudProviderAccess, error) {
	var resource CloudProviderAccess
	err := ctx.ReadResource("mongodbatlas:index/cloudProviderAccess:CloudProviderAccess", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CloudProviderAccess resources.
type cloudProviderAccessState struct {
	// Unique external ID Atlas uses when assuming the IAM role in your AWS account.
	AtlasAssumedRoleExternalId *string `pulumi:"atlasAssumedRoleExternalId"`
	// ARN associated with the Atlas AWS account used to assume IAM roles in your AWS account.
	AtlasAwsAccountArn *string `pulumi:"atlasAwsAccountArn"`
	// Date on which this role was authorized.
	AuthorizedDate *string `pulumi:"authorizedDate"`
	// Date on which this role was created.
	CreatedDate *string `pulumi:"createdDate"`
	// Atlas features this AWS IAM role is linked to.
	FeatureUsages []CloudProviderAccessFeatureUsage `pulumi:"featureUsages"`
	// ARN of the IAM Role that Atlas assumes when accessing resources in your AWS account. This value is required after the creation (register of the role) as part of [Set Up Unified AWS Access](https://docs.atlas.mongodb.com/security/set-up-unified-aws-access/#set-up-unified-aws-access).
	IamAssumedRoleArn *string `pulumi:"iamAssumedRoleArn"`
	// The unique ID for the project
	ProjectId *string `pulumi:"projectId"`
	// The cloud provider for which to create a new role. Currently only AWS is supported.
	ProviderName *string `pulumi:"providerName"`
	// Unique ID of this role returned by mongodb atlas api
	RoleId *string `pulumi:"roleId"`
}

type CloudProviderAccessState struct {
	// Unique external ID Atlas uses when assuming the IAM role in your AWS account.
	AtlasAssumedRoleExternalId pulumi.StringPtrInput
	// ARN associated with the Atlas AWS account used to assume IAM roles in your AWS account.
	AtlasAwsAccountArn pulumi.StringPtrInput
	// Date on which this role was authorized.
	AuthorizedDate pulumi.StringPtrInput
	// Date on which this role was created.
	CreatedDate pulumi.StringPtrInput
	// Atlas features this AWS IAM role is linked to.
	FeatureUsages CloudProviderAccessFeatureUsageArrayInput
	// ARN of the IAM Role that Atlas assumes when accessing resources in your AWS account. This value is required after the creation (register of the role) as part of [Set Up Unified AWS Access](https://docs.atlas.mongodb.com/security/set-up-unified-aws-access/#set-up-unified-aws-access).
	IamAssumedRoleArn pulumi.StringPtrInput
	// The unique ID for the project
	ProjectId pulumi.StringPtrInput
	// The cloud provider for which to create a new role. Currently only AWS is supported.
	ProviderName pulumi.StringPtrInput
	// Unique ID of this role returned by mongodb atlas api
	RoleId pulumi.StringPtrInput
}

func (CloudProviderAccessState) ElementType() reflect.Type {
	return reflect.TypeOf((*cloudProviderAccessState)(nil)).Elem()
}

type cloudProviderAccessArgs struct {
	// ARN of the IAM Role that Atlas assumes when accessing resources in your AWS account. This value is required after the creation (register of the role) as part of [Set Up Unified AWS Access](https://docs.atlas.mongodb.com/security/set-up-unified-aws-access/#set-up-unified-aws-access).
	IamAssumedRoleArn *string `pulumi:"iamAssumedRoleArn"`
	// The unique ID for the project
	ProjectId string `pulumi:"projectId"`
	// The cloud provider for which to create a new role. Currently only AWS is supported.
	ProviderName string `pulumi:"providerName"`
}

// The set of arguments for constructing a CloudProviderAccess resource.
type CloudProviderAccessArgs struct {
	// ARN of the IAM Role that Atlas assumes when accessing resources in your AWS account. This value is required after the creation (register of the role) as part of [Set Up Unified AWS Access](https://docs.atlas.mongodb.com/security/set-up-unified-aws-access/#set-up-unified-aws-access).
	IamAssumedRoleArn pulumi.StringPtrInput
	// The unique ID for the project
	ProjectId pulumi.StringInput
	// The cloud provider for which to create a new role. Currently only AWS is supported.
	ProviderName pulumi.StringInput
}

func (CloudProviderAccessArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*cloudProviderAccessArgs)(nil)).Elem()
}

type CloudProviderAccessInput interface {
	pulumi.Input

	ToCloudProviderAccessOutput() CloudProviderAccessOutput
	ToCloudProviderAccessOutputWithContext(ctx context.Context) CloudProviderAccessOutput
}

func (*CloudProviderAccess) ElementType() reflect.Type {
	return reflect.TypeOf((*CloudProviderAccess)(nil))
}

func (i *CloudProviderAccess) ToCloudProviderAccessOutput() CloudProviderAccessOutput {
	return i.ToCloudProviderAccessOutputWithContext(context.Background())
}

func (i *CloudProviderAccess) ToCloudProviderAccessOutputWithContext(ctx context.Context) CloudProviderAccessOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudProviderAccessOutput)
}

func (i *CloudProviderAccess) ToCloudProviderAccessPtrOutput() CloudProviderAccessPtrOutput {
	return i.ToCloudProviderAccessPtrOutputWithContext(context.Background())
}

func (i *CloudProviderAccess) ToCloudProviderAccessPtrOutputWithContext(ctx context.Context) CloudProviderAccessPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudProviderAccessPtrOutput)
}

type CloudProviderAccessPtrInput interface {
	pulumi.Input

	ToCloudProviderAccessPtrOutput() CloudProviderAccessPtrOutput
	ToCloudProviderAccessPtrOutputWithContext(ctx context.Context) CloudProviderAccessPtrOutput
}

type cloudProviderAccessPtrType CloudProviderAccessArgs

func (*cloudProviderAccessPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudProviderAccess)(nil))
}

func (i *cloudProviderAccessPtrType) ToCloudProviderAccessPtrOutput() CloudProviderAccessPtrOutput {
	return i.ToCloudProviderAccessPtrOutputWithContext(context.Background())
}

func (i *cloudProviderAccessPtrType) ToCloudProviderAccessPtrOutputWithContext(ctx context.Context) CloudProviderAccessPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudProviderAccessPtrOutput)
}

// CloudProviderAccessArrayInput is an input type that accepts CloudProviderAccessArray and CloudProviderAccessArrayOutput values.
// You can construct a concrete instance of `CloudProviderAccessArrayInput` via:
//
//          CloudProviderAccessArray{ CloudProviderAccessArgs{...} }
type CloudProviderAccessArrayInput interface {
	pulumi.Input

	ToCloudProviderAccessArrayOutput() CloudProviderAccessArrayOutput
	ToCloudProviderAccessArrayOutputWithContext(context.Context) CloudProviderAccessArrayOutput
}

type CloudProviderAccessArray []CloudProviderAccessInput

func (CloudProviderAccessArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*CloudProviderAccess)(nil))
}

func (i CloudProviderAccessArray) ToCloudProviderAccessArrayOutput() CloudProviderAccessArrayOutput {
	return i.ToCloudProviderAccessArrayOutputWithContext(context.Background())
}

func (i CloudProviderAccessArray) ToCloudProviderAccessArrayOutputWithContext(ctx context.Context) CloudProviderAccessArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudProviderAccessArrayOutput)
}

// CloudProviderAccessMapInput is an input type that accepts CloudProviderAccessMap and CloudProviderAccessMapOutput values.
// You can construct a concrete instance of `CloudProviderAccessMapInput` via:
//
//          CloudProviderAccessMap{ "key": CloudProviderAccessArgs{...} }
type CloudProviderAccessMapInput interface {
	pulumi.Input

	ToCloudProviderAccessMapOutput() CloudProviderAccessMapOutput
	ToCloudProviderAccessMapOutputWithContext(context.Context) CloudProviderAccessMapOutput
}

type CloudProviderAccessMap map[string]CloudProviderAccessInput

func (CloudProviderAccessMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*CloudProviderAccess)(nil))
}

func (i CloudProviderAccessMap) ToCloudProviderAccessMapOutput() CloudProviderAccessMapOutput {
	return i.ToCloudProviderAccessMapOutputWithContext(context.Background())
}

func (i CloudProviderAccessMap) ToCloudProviderAccessMapOutputWithContext(ctx context.Context) CloudProviderAccessMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudProviderAccessMapOutput)
}

type CloudProviderAccessOutput struct {
	*pulumi.OutputState
}

func (CloudProviderAccessOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CloudProviderAccess)(nil))
}

func (o CloudProviderAccessOutput) ToCloudProviderAccessOutput() CloudProviderAccessOutput {
	return o
}

func (o CloudProviderAccessOutput) ToCloudProviderAccessOutputWithContext(ctx context.Context) CloudProviderAccessOutput {
	return o
}

func (o CloudProviderAccessOutput) ToCloudProviderAccessPtrOutput() CloudProviderAccessPtrOutput {
	return o.ToCloudProviderAccessPtrOutputWithContext(context.Background())
}

func (o CloudProviderAccessOutput) ToCloudProviderAccessPtrOutputWithContext(ctx context.Context) CloudProviderAccessPtrOutput {
	return o.ApplyT(func(v CloudProviderAccess) *CloudProviderAccess {
		return &v
	}).(CloudProviderAccessPtrOutput)
}

type CloudProviderAccessPtrOutput struct {
	*pulumi.OutputState
}

func (CloudProviderAccessPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudProviderAccess)(nil))
}

func (o CloudProviderAccessPtrOutput) ToCloudProviderAccessPtrOutput() CloudProviderAccessPtrOutput {
	return o
}

func (o CloudProviderAccessPtrOutput) ToCloudProviderAccessPtrOutputWithContext(ctx context.Context) CloudProviderAccessPtrOutput {
	return o
}

type CloudProviderAccessArrayOutput struct{ *pulumi.OutputState }

func (CloudProviderAccessArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CloudProviderAccess)(nil))
}

func (o CloudProviderAccessArrayOutput) ToCloudProviderAccessArrayOutput() CloudProviderAccessArrayOutput {
	return o
}

func (o CloudProviderAccessArrayOutput) ToCloudProviderAccessArrayOutputWithContext(ctx context.Context) CloudProviderAccessArrayOutput {
	return o
}

func (o CloudProviderAccessArrayOutput) Index(i pulumi.IntInput) CloudProviderAccessOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) CloudProviderAccess {
		return vs[0].([]CloudProviderAccess)[vs[1].(int)]
	}).(CloudProviderAccessOutput)
}

type CloudProviderAccessMapOutput struct{ *pulumi.OutputState }

func (CloudProviderAccessMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]CloudProviderAccess)(nil))
}

func (o CloudProviderAccessMapOutput) ToCloudProviderAccessMapOutput() CloudProviderAccessMapOutput {
	return o
}

func (o CloudProviderAccessMapOutput) ToCloudProviderAccessMapOutputWithContext(ctx context.Context) CloudProviderAccessMapOutput {
	return o
}

func (o CloudProviderAccessMapOutput) MapIndex(k pulumi.StringInput) CloudProviderAccessOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) CloudProviderAccess {
		return vs[0].(map[string]CloudProviderAccess)[vs[1].(string)]
	}).(CloudProviderAccessOutput)
}

func init() {
	pulumi.RegisterOutputType(CloudProviderAccessOutput{})
	pulumi.RegisterOutputType(CloudProviderAccessPtrOutput{})
	pulumi.RegisterOutputType(CloudProviderAccessArrayOutput{})
	pulumi.RegisterOutputType(CloudProviderAccessMapOutput{})
}
