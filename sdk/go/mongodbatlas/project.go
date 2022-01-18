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
// Project must be imported using project ID, e.g.
//
// ```sh
//  $ pulumi import mongodbatlas:index/project:Project my_project 5d09d6a59ccf6445652a444a
// ```
//
//  For more information see[MongoDB Atlas API Reference.](https://docs.atlas.mongodb.com/reference/api/projects/) - [and MongoDB Atlas API - Teams](https://docs.atlas.mongodb.com/reference/api/teams/) Documentation for more information.
type Project struct {
	pulumi.CustomResourceState

	ApiKeys ProjectApiKeyArrayOutput `pulumi:"apiKeys"`
	// The number of Atlas clusters deployed in the project..
	ClusterCount pulumi.IntOutput `pulumi:"clusterCount"`
	// The ISO-8601-formatted timestamp of when Atlas created the project..
	Created pulumi.StringOutput `pulumi:"created"`
	// The name of the project you want to create. (Cannot be changed via this Provider after creation.)
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the organization you want to create the project within.
	OrgId pulumi.StringOutput `pulumi:"orgId"`
	// Unique 24-hexadecimal digit string that identifies the Atlas user account to be granted the [Project Owner](https://docs.atlas.mongodb.com/reference/user-roles/#mongodb-authrole-Project-Owner) role on the specified project. If you set this parameter, it overrides the default value of the oldest [Organization Owner](https://docs.atlas.mongodb.com/reference/user-roles/#mongodb-authrole-Organization-Owner).
	ProjectOwnerId pulumi.StringPtrOutput `pulumi:"projectOwnerId"`
	Teams          ProjectTeamArrayOutput `pulumi:"teams"`
	// It allows users to disable the creation of the default alert settings. By default, this flag is set to true.
	WithDefaultAlertsSettings pulumi.BoolPtrOutput `pulumi:"withDefaultAlertsSettings"`
}

// NewProject registers a new resource with the given unique name, arguments, and options.
func NewProject(ctx *pulumi.Context,
	name string, args *ProjectArgs, opts ...pulumi.ResourceOption) (*Project, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.OrgId == nil {
		return nil, errors.New("invalid value for required argument 'OrgId'")
	}
	var resource Project
	err := ctx.RegisterResource("mongodbatlas:index/project:Project", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProject gets an existing Project resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProject(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProjectState, opts ...pulumi.ResourceOption) (*Project, error) {
	var resource Project
	err := ctx.ReadResource("mongodbatlas:index/project:Project", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Project resources.
type projectState struct {
	ApiKeys []ProjectApiKey `pulumi:"apiKeys"`
	// The number of Atlas clusters deployed in the project..
	ClusterCount *int `pulumi:"clusterCount"`
	// The ISO-8601-formatted timestamp of when Atlas created the project..
	Created *string `pulumi:"created"`
	// The name of the project you want to create. (Cannot be changed via this Provider after creation.)
	Name *string `pulumi:"name"`
	// The ID of the organization you want to create the project within.
	OrgId *string `pulumi:"orgId"`
	// Unique 24-hexadecimal digit string that identifies the Atlas user account to be granted the [Project Owner](https://docs.atlas.mongodb.com/reference/user-roles/#mongodb-authrole-Project-Owner) role on the specified project. If you set this parameter, it overrides the default value of the oldest [Organization Owner](https://docs.atlas.mongodb.com/reference/user-roles/#mongodb-authrole-Organization-Owner).
	ProjectOwnerId *string       `pulumi:"projectOwnerId"`
	Teams          []ProjectTeam `pulumi:"teams"`
	// It allows users to disable the creation of the default alert settings. By default, this flag is set to true.
	WithDefaultAlertsSettings *bool `pulumi:"withDefaultAlertsSettings"`
}

type ProjectState struct {
	ApiKeys ProjectApiKeyArrayInput
	// The number of Atlas clusters deployed in the project..
	ClusterCount pulumi.IntPtrInput
	// The ISO-8601-formatted timestamp of when Atlas created the project..
	Created pulumi.StringPtrInput
	// The name of the project you want to create. (Cannot be changed via this Provider after creation.)
	Name pulumi.StringPtrInput
	// The ID of the organization you want to create the project within.
	OrgId pulumi.StringPtrInput
	// Unique 24-hexadecimal digit string that identifies the Atlas user account to be granted the [Project Owner](https://docs.atlas.mongodb.com/reference/user-roles/#mongodb-authrole-Project-Owner) role on the specified project. If you set this parameter, it overrides the default value of the oldest [Organization Owner](https://docs.atlas.mongodb.com/reference/user-roles/#mongodb-authrole-Organization-Owner).
	ProjectOwnerId pulumi.StringPtrInput
	Teams          ProjectTeamArrayInput
	// It allows users to disable the creation of the default alert settings. By default, this flag is set to true.
	WithDefaultAlertsSettings pulumi.BoolPtrInput
}

func (ProjectState) ElementType() reflect.Type {
	return reflect.TypeOf((*projectState)(nil)).Elem()
}

type projectArgs struct {
	ApiKeys []ProjectApiKey `pulumi:"apiKeys"`
	// The name of the project you want to create. (Cannot be changed via this Provider after creation.)
	Name *string `pulumi:"name"`
	// The ID of the organization you want to create the project within.
	OrgId string `pulumi:"orgId"`
	// Unique 24-hexadecimal digit string that identifies the Atlas user account to be granted the [Project Owner](https://docs.atlas.mongodb.com/reference/user-roles/#mongodb-authrole-Project-Owner) role on the specified project. If you set this parameter, it overrides the default value of the oldest [Organization Owner](https://docs.atlas.mongodb.com/reference/user-roles/#mongodb-authrole-Organization-Owner).
	ProjectOwnerId *string       `pulumi:"projectOwnerId"`
	Teams          []ProjectTeam `pulumi:"teams"`
	// It allows users to disable the creation of the default alert settings. By default, this flag is set to true.
	WithDefaultAlertsSettings *bool `pulumi:"withDefaultAlertsSettings"`
}

// The set of arguments for constructing a Project resource.
type ProjectArgs struct {
	ApiKeys ProjectApiKeyArrayInput
	// The name of the project you want to create. (Cannot be changed via this Provider after creation.)
	Name pulumi.StringPtrInput
	// The ID of the organization you want to create the project within.
	OrgId pulumi.StringInput
	// Unique 24-hexadecimal digit string that identifies the Atlas user account to be granted the [Project Owner](https://docs.atlas.mongodb.com/reference/user-roles/#mongodb-authrole-Project-Owner) role on the specified project. If you set this parameter, it overrides the default value of the oldest [Organization Owner](https://docs.atlas.mongodb.com/reference/user-roles/#mongodb-authrole-Organization-Owner).
	ProjectOwnerId pulumi.StringPtrInput
	Teams          ProjectTeamArrayInput
	// It allows users to disable the creation of the default alert settings. By default, this flag is set to true.
	WithDefaultAlertsSettings pulumi.BoolPtrInput
}

func (ProjectArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*projectArgs)(nil)).Elem()
}

type ProjectInput interface {
	pulumi.Input

	ToProjectOutput() ProjectOutput
	ToProjectOutputWithContext(ctx context.Context) ProjectOutput
}

func (*Project) ElementType() reflect.Type {
	return reflect.TypeOf((**Project)(nil)).Elem()
}

func (i *Project) ToProjectOutput() ProjectOutput {
	return i.ToProjectOutputWithContext(context.Background())
}

func (i *Project) ToProjectOutputWithContext(ctx context.Context) ProjectOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectOutput)
}

// ProjectArrayInput is an input type that accepts ProjectArray and ProjectArrayOutput values.
// You can construct a concrete instance of `ProjectArrayInput` via:
//
//          ProjectArray{ ProjectArgs{...} }
type ProjectArrayInput interface {
	pulumi.Input

	ToProjectArrayOutput() ProjectArrayOutput
	ToProjectArrayOutputWithContext(context.Context) ProjectArrayOutput
}

type ProjectArray []ProjectInput

func (ProjectArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Project)(nil)).Elem()
}

func (i ProjectArray) ToProjectArrayOutput() ProjectArrayOutput {
	return i.ToProjectArrayOutputWithContext(context.Background())
}

func (i ProjectArray) ToProjectArrayOutputWithContext(ctx context.Context) ProjectArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectArrayOutput)
}

// ProjectMapInput is an input type that accepts ProjectMap and ProjectMapOutput values.
// You can construct a concrete instance of `ProjectMapInput` via:
//
//          ProjectMap{ "key": ProjectArgs{...} }
type ProjectMapInput interface {
	pulumi.Input

	ToProjectMapOutput() ProjectMapOutput
	ToProjectMapOutputWithContext(context.Context) ProjectMapOutput
}

type ProjectMap map[string]ProjectInput

func (ProjectMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Project)(nil)).Elem()
}

func (i ProjectMap) ToProjectMapOutput() ProjectMapOutput {
	return i.ToProjectMapOutputWithContext(context.Background())
}

func (i ProjectMap) ToProjectMapOutputWithContext(ctx context.Context) ProjectMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectMapOutput)
}

type ProjectOutput struct{ *pulumi.OutputState }

func (ProjectOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Project)(nil)).Elem()
}

func (o ProjectOutput) ToProjectOutput() ProjectOutput {
	return o
}

func (o ProjectOutput) ToProjectOutputWithContext(ctx context.Context) ProjectOutput {
	return o
}

type ProjectArrayOutput struct{ *pulumi.OutputState }

func (ProjectArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Project)(nil)).Elem()
}

func (o ProjectArrayOutput) ToProjectArrayOutput() ProjectArrayOutput {
	return o
}

func (o ProjectArrayOutput) ToProjectArrayOutputWithContext(ctx context.Context) ProjectArrayOutput {
	return o
}

func (o ProjectArrayOutput) Index(i pulumi.IntInput) ProjectOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Project {
		return vs[0].([]*Project)[vs[1].(int)]
	}).(ProjectOutput)
}

type ProjectMapOutput struct{ *pulumi.OutputState }

func (ProjectMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Project)(nil)).Elem()
}

func (o ProjectMapOutput) ToProjectMapOutput() ProjectMapOutput {
	return o
}

func (o ProjectMapOutput) ToProjectMapOutputWithContext(ctx context.Context) ProjectMapOutput {
	return o
}

func (o ProjectMapOutput) MapIndex(k pulumi.StringInput) ProjectOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Project {
		return vs[0].(map[string]*Project)[vs[1].(string)]
	}).(ProjectOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectInput)(nil)).Elem(), &Project{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectArrayInput)(nil)).Elem(), ProjectArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectMapInput)(nil)).Elem(), ProjectMap{})
	pulumi.RegisterOutputType(ProjectOutput{})
	pulumi.RegisterOutputType(ProjectArrayOutput{})
	pulumi.RegisterOutputType(ProjectMapOutput{})
}
