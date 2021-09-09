// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package mongodbatlas

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// `Auditing` provides an Auditing resource. This allows auditing to be created.
//
// ## Example Usage
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
// 		_, err := mongodbatlas.NewAuditing(ctx, "test", &mongodbatlas.AuditingArgs{
// 			AuditAuthorizationSuccess: pulumi.Bool(false),
// 			AuditFilter:               pulumi.String("{ 'atype': 'authenticate', 'param': {   'user': 'auditAdmin',   'db': 'admin',   'mechanism': 'SCRAM-SHA-1' }}"),
// 			Enabled:                   pulumi.Bool(true),
// 			ProjectId:                 pulumi.String("<project-id>"),
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
// Auditing must be imported using auditing ID, e.g.
//
// ```sh
//  $ pulumi import mongodbatlas:index/auditing:Auditing my_auditing 5d09d6a59ccf6445652a444a
// ```
//
//  For more information see[MongoDB Atlas API Reference.](https://docs.atlas.mongodb.com/reference/api/auditing/)
type Auditing struct {
	pulumi.CustomResourceState

	// JSON-formatted audit filter used by the project
	AuditAuthorizationSuccess pulumi.BoolOutput `pulumi:"auditAuthorizationSuccess"`
	// Indicates whether the auditing system captures successful authentication attempts for audit filters using the "atype" : "authCheck" auditing event. For more information, see auditAuthorizationSuccess
	AuditFilter pulumi.StringOutput `pulumi:"auditFilter"`
	// Denotes the configuration method for the audit filter. Possible values are:
	// * NONE - auditing not configured for the project.
	// * FILTER_BUILDER - auditing configured via Atlas UI filter builder.
	// * FILTER_JSON - auditing configured via Atlas custom filter or API.
	ConfigurationType pulumi.StringOutput `pulumi:"configurationType"`
	// Denotes whether or not the project associated with the {project_id} has database auditing enabled.
	Enabled pulumi.BoolOutput `pulumi:"enabled"`
	// The unique ID for the project to configure auditing.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
}

// NewAuditing registers a new resource with the given unique name, arguments, and options.
func NewAuditing(ctx *pulumi.Context,
	name string, args *AuditingArgs, opts ...pulumi.ResourceOption) (*Auditing, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ProjectId == nil {
		return nil, errors.New("invalid value for required argument 'ProjectId'")
	}
	var resource Auditing
	err := ctx.RegisterResource("mongodbatlas:index/auditing:Auditing", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAuditing gets an existing Auditing resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAuditing(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AuditingState, opts ...pulumi.ResourceOption) (*Auditing, error) {
	var resource Auditing
	err := ctx.ReadResource("mongodbatlas:index/auditing:Auditing", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Auditing resources.
type auditingState struct {
	// JSON-formatted audit filter used by the project
	AuditAuthorizationSuccess *bool `pulumi:"auditAuthorizationSuccess"`
	// Indicates whether the auditing system captures successful authentication attempts for audit filters using the "atype" : "authCheck" auditing event. For more information, see auditAuthorizationSuccess
	AuditFilter *string `pulumi:"auditFilter"`
	// Denotes the configuration method for the audit filter. Possible values are:
	// * NONE - auditing not configured for the project.
	// * FILTER_BUILDER - auditing configured via Atlas UI filter builder.
	// * FILTER_JSON - auditing configured via Atlas custom filter or API.
	ConfigurationType *string `pulumi:"configurationType"`
	// Denotes whether or not the project associated with the {project_id} has database auditing enabled.
	Enabled *bool `pulumi:"enabled"`
	// The unique ID for the project to configure auditing.
	ProjectId *string `pulumi:"projectId"`
}

type AuditingState struct {
	// JSON-formatted audit filter used by the project
	AuditAuthorizationSuccess pulumi.BoolPtrInput
	// Indicates whether the auditing system captures successful authentication attempts for audit filters using the "atype" : "authCheck" auditing event. For more information, see auditAuthorizationSuccess
	AuditFilter pulumi.StringPtrInput
	// Denotes the configuration method for the audit filter. Possible values are:
	// * NONE - auditing not configured for the project.
	// * FILTER_BUILDER - auditing configured via Atlas UI filter builder.
	// * FILTER_JSON - auditing configured via Atlas custom filter or API.
	ConfigurationType pulumi.StringPtrInput
	// Denotes whether or not the project associated with the {project_id} has database auditing enabled.
	Enabled pulumi.BoolPtrInput
	// The unique ID for the project to configure auditing.
	ProjectId pulumi.StringPtrInput
}

func (AuditingState) ElementType() reflect.Type {
	return reflect.TypeOf((*auditingState)(nil)).Elem()
}

type auditingArgs struct {
	// JSON-formatted audit filter used by the project
	AuditAuthorizationSuccess *bool `pulumi:"auditAuthorizationSuccess"`
	// Indicates whether the auditing system captures successful authentication attempts for audit filters using the "atype" : "authCheck" auditing event. For more information, see auditAuthorizationSuccess
	AuditFilter *string `pulumi:"auditFilter"`
	// Denotes whether or not the project associated with the {project_id} has database auditing enabled.
	Enabled *bool `pulumi:"enabled"`
	// The unique ID for the project to configure auditing.
	ProjectId string `pulumi:"projectId"`
}

// The set of arguments for constructing a Auditing resource.
type AuditingArgs struct {
	// JSON-formatted audit filter used by the project
	AuditAuthorizationSuccess pulumi.BoolPtrInput
	// Indicates whether the auditing system captures successful authentication attempts for audit filters using the "atype" : "authCheck" auditing event. For more information, see auditAuthorizationSuccess
	AuditFilter pulumi.StringPtrInput
	// Denotes whether or not the project associated with the {project_id} has database auditing enabled.
	Enabled pulumi.BoolPtrInput
	// The unique ID for the project to configure auditing.
	ProjectId pulumi.StringInput
}

func (AuditingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*auditingArgs)(nil)).Elem()
}

type AuditingInput interface {
	pulumi.Input

	ToAuditingOutput() AuditingOutput
	ToAuditingOutputWithContext(ctx context.Context) AuditingOutput
}

func (*Auditing) ElementType() reflect.Type {
	return reflect.TypeOf((*Auditing)(nil))
}

func (i *Auditing) ToAuditingOutput() AuditingOutput {
	return i.ToAuditingOutputWithContext(context.Background())
}

func (i *Auditing) ToAuditingOutputWithContext(ctx context.Context) AuditingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuditingOutput)
}

func (i *Auditing) ToAuditingPtrOutput() AuditingPtrOutput {
	return i.ToAuditingPtrOutputWithContext(context.Background())
}

func (i *Auditing) ToAuditingPtrOutputWithContext(ctx context.Context) AuditingPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuditingPtrOutput)
}

type AuditingPtrInput interface {
	pulumi.Input

	ToAuditingPtrOutput() AuditingPtrOutput
	ToAuditingPtrOutputWithContext(ctx context.Context) AuditingPtrOutput
}

type auditingPtrType AuditingArgs

func (*auditingPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Auditing)(nil))
}

func (i *auditingPtrType) ToAuditingPtrOutput() AuditingPtrOutput {
	return i.ToAuditingPtrOutputWithContext(context.Background())
}

func (i *auditingPtrType) ToAuditingPtrOutputWithContext(ctx context.Context) AuditingPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuditingPtrOutput)
}

// AuditingArrayInput is an input type that accepts AuditingArray and AuditingArrayOutput values.
// You can construct a concrete instance of `AuditingArrayInput` via:
//
//          AuditingArray{ AuditingArgs{...} }
type AuditingArrayInput interface {
	pulumi.Input

	ToAuditingArrayOutput() AuditingArrayOutput
	ToAuditingArrayOutputWithContext(context.Context) AuditingArrayOutput
}

type AuditingArray []AuditingInput

func (AuditingArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*Auditing)(nil))
}

func (i AuditingArray) ToAuditingArrayOutput() AuditingArrayOutput {
	return i.ToAuditingArrayOutputWithContext(context.Background())
}

func (i AuditingArray) ToAuditingArrayOutputWithContext(ctx context.Context) AuditingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuditingArrayOutput)
}

// AuditingMapInput is an input type that accepts AuditingMap and AuditingMapOutput values.
// You can construct a concrete instance of `AuditingMapInput` via:
//
//          AuditingMap{ "key": AuditingArgs{...} }
type AuditingMapInput interface {
	pulumi.Input

	ToAuditingMapOutput() AuditingMapOutput
	ToAuditingMapOutputWithContext(context.Context) AuditingMapOutput
}

type AuditingMap map[string]AuditingInput

func (AuditingMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*Auditing)(nil))
}

func (i AuditingMap) ToAuditingMapOutput() AuditingMapOutput {
	return i.ToAuditingMapOutputWithContext(context.Background())
}

func (i AuditingMap) ToAuditingMapOutputWithContext(ctx context.Context) AuditingMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuditingMapOutput)
}

type AuditingOutput struct {
	*pulumi.OutputState
}

func (AuditingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Auditing)(nil))
}

func (o AuditingOutput) ToAuditingOutput() AuditingOutput {
	return o
}

func (o AuditingOutput) ToAuditingOutputWithContext(ctx context.Context) AuditingOutput {
	return o
}

func (o AuditingOutput) ToAuditingPtrOutput() AuditingPtrOutput {
	return o.ToAuditingPtrOutputWithContext(context.Background())
}

func (o AuditingOutput) ToAuditingPtrOutputWithContext(ctx context.Context) AuditingPtrOutput {
	return o.ApplyT(func(v Auditing) *Auditing {
		return &v
	}).(AuditingPtrOutput)
}

type AuditingPtrOutput struct {
	*pulumi.OutputState
}

func (AuditingPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Auditing)(nil))
}

func (o AuditingPtrOutput) ToAuditingPtrOutput() AuditingPtrOutput {
	return o
}

func (o AuditingPtrOutput) ToAuditingPtrOutputWithContext(ctx context.Context) AuditingPtrOutput {
	return o
}

type AuditingArrayOutput struct{ *pulumi.OutputState }

func (AuditingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Auditing)(nil))
}

func (o AuditingArrayOutput) ToAuditingArrayOutput() AuditingArrayOutput {
	return o
}

func (o AuditingArrayOutput) ToAuditingArrayOutputWithContext(ctx context.Context) AuditingArrayOutput {
	return o
}

func (o AuditingArrayOutput) Index(i pulumi.IntInput) AuditingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Auditing {
		return vs[0].([]Auditing)[vs[1].(int)]
	}).(AuditingOutput)
}

type AuditingMapOutput struct{ *pulumi.OutputState }

func (AuditingMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Auditing)(nil))
}

func (o AuditingMapOutput) ToAuditingMapOutput() AuditingMapOutput {
	return o
}

func (o AuditingMapOutput) ToAuditingMapOutputWithContext(ctx context.Context) AuditingMapOutput {
	return o
}

func (o AuditingMapOutput) MapIndex(k pulumi.StringInput) AuditingOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Auditing {
		return vs[0].(map[string]Auditing)[vs[1].(string)]
	}).(AuditingOutput)
}

func init() {
	pulumi.RegisterOutputType(AuditingOutput{})
	pulumi.RegisterOutputType(AuditingPtrOutput{})
	pulumi.RegisterOutputType(AuditingArrayOutput{})
	pulumi.RegisterOutputType(AuditingMapOutput{})
}
