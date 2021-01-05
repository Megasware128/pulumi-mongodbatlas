// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package mongodbatlas

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// `MaintenanceWindow` provides a resource to schedule a maintenance window for your MongoDB Atlas Project and/or set to defer a scheduled maintenance up to two times.
//
// > **NOTE:** Groups and projects are synonymous terms. You may find `groupId` in the official documentation.
//
// ## Maintenance Window Considerations:
//
// - Urgent Maintenance Activities Cannot Wait: Urgent maintenance activities such as security patches cannot wait for your chosen window. Atlas will start those maintenance activities when needed.
//
// Once maintenance is scheduled for your cluster, you cannot change your maintenance window until the current maintenance efforts have completed.
// - Maintenance Requires Replica Set Elections: Atlas performs maintenance the same way as the manual maintenance procedure. This requires at least one replica set election during the maintenance window per replica set.
// - Maintenance Starts As Close to the Hour As Possible: Maintenance always begins as close to the scheduled hour as possible, but in-progress cluster updates or expected system issues could delay the start time.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-mongodbatlas/sdk/go/mongodbatlas"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := mongodbatlas.NewMaintenanceWindow(ctx, "test", &mongodbatlas.MaintenanceWindowArgs{
// 			DayOfWeek: pulumi.Int(3),
// 			HourOfDay: pulumi.Int(4),
// 			ProjectId: pulumi.String("<your-project-id>"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-mongodbatlas/sdk/go/mongodbatlas"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := mongodbatlas.NewMaintenanceWindow(ctx, "test", &mongodbatlas.MaintenanceWindowArgs{
// 			Defer:     pulumi.Bool(true),
// 			ProjectId: pulumi.String("<your-project-id>"),
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
// Maintenance Window entries can be imported using project project_id, in the format `PROJECTID`, e.g.
//
// ```sh
//  $ pulumi import mongodbatlas:index/maintenanceWindow:MaintenanceWindow test 5d0f1f73cf09a29120e173cf
// ```
//
//  For more information see[MongoDB Atlas API Reference.](https://docs.atlas.mongodb.com/reference/api/maintenance-windows/)
type MaintenanceWindow struct {
	pulumi.CustomResourceState

	// Day of the week when you would like the maintenance window to start as a 1-based integer: S=1, M=2, T=3, W=4, T=5, F=6, S=7.
	DayOfWeek pulumi.IntOutput `pulumi:"dayOfWeek"`
	// Defer maintenance for the given project for one week.
	Defer pulumi.BoolOutput `pulumi:"defer"`
	// Hour of the day when you would like the maintenance window to start. This parameter uses the 24-hour clock, where midnight is 0, noon is 12 (Time zone is UTC).
	HourOfDay pulumi.IntOutput `pulumi:"hourOfDay"`
	// Number of times the current maintenance event for this project has been deferred, you can set a maximum of 2 deferrals.
	NumberOfDeferrals pulumi.IntOutput `pulumi:"numberOfDeferrals"`
	// The unique identifier of the project for the Maintenance Window.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// Flag indicating whether project maintenance has been directed to start immediately. If you request that maintenance begin immediately, this field returns true from the time the request was made until the time the maintenance event completes.
	StartAsap pulumi.BoolOutput `pulumi:"startAsap"`
}

// NewMaintenanceWindow registers a new resource with the given unique name, arguments, and options.
func NewMaintenanceWindow(ctx *pulumi.Context,
	name string, args *MaintenanceWindowArgs, opts ...pulumi.ResourceOption) (*MaintenanceWindow, error) {
	if args == nil || args.ProjectId == nil {
		return nil, errors.New("missing required argument 'ProjectId'")
	}
	if args == nil {
		args = &MaintenanceWindowArgs{}
	}
	var resource MaintenanceWindow
	err := ctx.RegisterResource("mongodbatlas:index/maintenanceWindow:MaintenanceWindow", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMaintenanceWindow gets an existing MaintenanceWindow resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMaintenanceWindow(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MaintenanceWindowState, opts ...pulumi.ResourceOption) (*MaintenanceWindow, error) {
	var resource MaintenanceWindow
	err := ctx.ReadResource("mongodbatlas:index/maintenanceWindow:MaintenanceWindow", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MaintenanceWindow resources.
type maintenanceWindowState struct {
	// Day of the week when you would like the maintenance window to start as a 1-based integer: S=1, M=2, T=3, W=4, T=5, F=6, S=7.
	DayOfWeek *int `pulumi:"dayOfWeek"`
	// Defer maintenance for the given project for one week.
	Defer *bool `pulumi:"defer"`
	// Hour of the day when you would like the maintenance window to start. This parameter uses the 24-hour clock, where midnight is 0, noon is 12 (Time zone is UTC).
	HourOfDay *int `pulumi:"hourOfDay"`
	// Number of times the current maintenance event for this project has been deferred, you can set a maximum of 2 deferrals.
	NumberOfDeferrals *int `pulumi:"numberOfDeferrals"`
	// The unique identifier of the project for the Maintenance Window.
	ProjectId *string `pulumi:"projectId"`
	// Flag indicating whether project maintenance has been directed to start immediately. If you request that maintenance begin immediately, this field returns true from the time the request was made until the time the maintenance event completes.
	StartAsap *bool `pulumi:"startAsap"`
}

type MaintenanceWindowState struct {
	// Day of the week when you would like the maintenance window to start as a 1-based integer: S=1, M=2, T=3, W=4, T=5, F=6, S=7.
	DayOfWeek pulumi.IntPtrInput
	// Defer maintenance for the given project for one week.
	Defer pulumi.BoolPtrInput
	// Hour of the day when you would like the maintenance window to start. This parameter uses the 24-hour clock, where midnight is 0, noon is 12 (Time zone is UTC).
	HourOfDay pulumi.IntPtrInput
	// Number of times the current maintenance event for this project has been deferred, you can set a maximum of 2 deferrals.
	NumberOfDeferrals pulumi.IntPtrInput
	// The unique identifier of the project for the Maintenance Window.
	ProjectId pulumi.StringPtrInput
	// Flag indicating whether project maintenance has been directed to start immediately. If you request that maintenance begin immediately, this field returns true from the time the request was made until the time the maintenance event completes.
	StartAsap pulumi.BoolPtrInput
}

func (MaintenanceWindowState) ElementType() reflect.Type {
	return reflect.TypeOf((*maintenanceWindowState)(nil)).Elem()
}

type maintenanceWindowArgs struct {
	// Day of the week when you would like the maintenance window to start as a 1-based integer: S=1, M=2, T=3, W=4, T=5, F=6, S=7.
	DayOfWeek *int `pulumi:"dayOfWeek"`
	// Defer maintenance for the given project for one week.
	Defer *bool `pulumi:"defer"`
	// Hour of the day when you would like the maintenance window to start. This parameter uses the 24-hour clock, where midnight is 0, noon is 12 (Time zone is UTC).
	HourOfDay *int `pulumi:"hourOfDay"`
	// Number of times the current maintenance event for this project has been deferred, you can set a maximum of 2 deferrals.
	NumberOfDeferrals *int `pulumi:"numberOfDeferrals"`
	// The unique identifier of the project for the Maintenance Window.
	ProjectId string `pulumi:"projectId"`
}

// The set of arguments for constructing a MaintenanceWindow resource.
type MaintenanceWindowArgs struct {
	// Day of the week when you would like the maintenance window to start as a 1-based integer: S=1, M=2, T=3, W=4, T=5, F=6, S=7.
	DayOfWeek pulumi.IntPtrInput
	// Defer maintenance for the given project for one week.
	Defer pulumi.BoolPtrInput
	// Hour of the day when you would like the maintenance window to start. This parameter uses the 24-hour clock, where midnight is 0, noon is 12 (Time zone is UTC).
	HourOfDay pulumi.IntPtrInput
	// Number of times the current maintenance event for this project has been deferred, you can set a maximum of 2 deferrals.
	NumberOfDeferrals pulumi.IntPtrInput
	// The unique identifier of the project for the Maintenance Window.
	ProjectId pulumi.StringInput
}

func (MaintenanceWindowArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*maintenanceWindowArgs)(nil)).Elem()
}

type MaintenanceWindowInput interface {
	pulumi.Input

	ToMaintenanceWindowOutput() MaintenanceWindowOutput
	ToMaintenanceWindowOutputWithContext(ctx context.Context) MaintenanceWindowOutput
}

func (MaintenanceWindow) ElementType() reflect.Type {
	return reflect.TypeOf((*MaintenanceWindow)(nil)).Elem()
}

func (i MaintenanceWindow) ToMaintenanceWindowOutput() MaintenanceWindowOutput {
	return i.ToMaintenanceWindowOutputWithContext(context.Background())
}

func (i MaintenanceWindow) ToMaintenanceWindowOutputWithContext(ctx context.Context) MaintenanceWindowOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MaintenanceWindowOutput)
}

type MaintenanceWindowOutput struct {
	*pulumi.OutputState
}

func (MaintenanceWindowOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MaintenanceWindowOutput)(nil)).Elem()
}

func (o MaintenanceWindowOutput) ToMaintenanceWindowOutput() MaintenanceWindowOutput {
	return o
}

func (o MaintenanceWindowOutput) ToMaintenanceWindowOutputWithContext(ctx context.Context) MaintenanceWindowOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(MaintenanceWindowOutput{})
}
