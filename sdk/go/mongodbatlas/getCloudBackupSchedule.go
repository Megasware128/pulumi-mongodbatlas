// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package mongodbatlas

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// `CloudBackupSchedule` provides a Cloud Backup Schedule datasource. An Atlas Cloud Backup Schedule provides the current cloud backup schedule for the cluster.
//
// > **NOTE:** Groups and projects are synonymous terms. You may find `groupId` in the official documentation.
func LookupCloudBackupSchedule(ctx *pulumi.Context, args *LookupCloudBackupScheduleArgs, opts ...pulumi.InvokeOption) (*LookupCloudBackupScheduleResult, error) {
	var rv LookupCloudBackupScheduleResult
	err := ctx.Invoke("mongodbatlas:index/getCloudBackupSchedule:getCloudBackupSchedule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCloudBackupSchedule.
type LookupCloudBackupScheduleArgs struct {
	// The name of the Atlas cluster that contains the snapshots backup policy you want to retrieve.
	ClusterName string `pulumi:"clusterName"`
	// The unique identifier of the project for the Atlas cluster.
	ProjectId string `pulumi:"projectId"`
}

// A collection of values returned by getCloudBackupSchedule.
type LookupCloudBackupScheduleResult struct {
	// Unique identifier of the Atlas cluster.
	ClusterId   string `pulumi:"clusterId"`
	ClusterName string `pulumi:"clusterName"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Unique identifier of the backup policy.
	IdPolicy string `pulumi:"idPolicy"`
	// UTC ISO 8601 formatted point in time when Atlas will take the next snapshot.
	NextSnapshot string `pulumi:"nextSnapshot"`
	// Daily policy item
	PolicyItemDailies []GetCloudBackupSchedulePolicyItemDaily `pulumi:"policyItemDailies"`
	// Hourly policy item
	PolicyItemHourlies []GetCloudBackupSchedulePolicyItemHourly `pulumi:"policyItemHourlies"`
	// Monthly policy item
	PolicyItemMonthlies []GetCloudBackupSchedulePolicyItemMonthly `pulumi:"policyItemMonthlies"`
	// Weekly policy item
	PolicyItemWeeklies []GetCloudBackupSchedulePolicyItemWeekly `pulumi:"policyItemWeeklies"`
	ProjectId          string                                   `pulumi:"projectId"`
	// UTC Hour of day between 0 and 23 representing which hour of the day that Atlas takes a snapshot.
	ReferenceHourOfDay int `pulumi:"referenceHourOfDay"`
	// UTC Minute of day between 0 and 59 representing which minute of the `referenceHourOfDay` that Atlas takes the snapshot.
	ReferenceMinuteOfHour int `pulumi:"referenceMinuteOfHour"`
	// Specifies a restore window in days for cloud backup to maintain.
	RestoreWindowDays int `pulumi:"restoreWindowDays"`
}

func LookupCloudBackupScheduleOutput(ctx *pulumi.Context, args LookupCloudBackupScheduleOutputArgs, opts ...pulumi.InvokeOption) LookupCloudBackupScheduleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupCloudBackupScheduleResult, error) {
			args := v.(LookupCloudBackupScheduleArgs)
			r, err := LookupCloudBackupSchedule(ctx, &args, opts...)
			return *r, err
		}).(LookupCloudBackupScheduleResultOutput)
}

// A collection of arguments for invoking getCloudBackupSchedule.
type LookupCloudBackupScheduleOutputArgs struct {
	// The name of the Atlas cluster that contains the snapshots backup policy you want to retrieve.
	ClusterName pulumi.StringInput `pulumi:"clusterName"`
	// The unique identifier of the project for the Atlas cluster.
	ProjectId pulumi.StringInput `pulumi:"projectId"`
}

func (LookupCloudBackupScheduleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCloudBackupScheduleArgs)(nil)).Elem()
}

// A collection of values returned by getCloudBackupSchedule.
type LookupCloudBackupScheduleResultOutput struct{ *pulumi.OutputState }

func (LookupCloudBackupScheduleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCloudBackupScheduleResult)(nil)).Elem()
}

func (o LookupCloudBackupScheduleResultOutput) ToLookupCloudBackupScheduleResultOutput() LookupCloudBackupScheduleResultOutput {
	return o
}

func (o LookupCloudBackupScheduleResultOutput) ToLookupCloudBackupScheduleResultOutputWithContext(ctx context.Context) LookupCloudBackupScheduleResultOutput {
	return o
}

// Unique identifier of the Atlas cluster.
func (o LookupCloudBackupScheduleResultOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudBackupScheduleResult) string { return v.ClusterId }).(pulumi.StringOutput)
}

func (o LookupCloudBackupScheduleResultOutput) ClusterName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudBackupScheduleResult) string { return v.ClusterName }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupCloudBackupScheduleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudBackupScheduleResult) string { return v.Id }).(pulumi.StringOutput)
}

// Unique identifier of the backup policy.
func (o LookupCloudBackupScheduleResultOutput) IdPolicy() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudBackupScheduleResult) string { return v.IdPolicy }).(pulumi.StringOutput)
}

// UTC ISO 8601 formatted point in time when Atlas will take the next snapshot.
func (o LookupCloudBackupScheduleResultOutput) NextSnapshot() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudBackupScheduleResult) string { return v.NextSnapshot }).(pulumi.StringOutput)
}

// Daily policy item
func (o LookupCloudBackupScheduleResultOutput) PolicyItemDailies() GetCloudBackupSchedulePolicyItemDailyArrayOutput {
	return o.ApplyT(func(v LookupCloudBackupScheduleResult) []GetCloudBackupSchedulePolicyItemDaily {
		return v.PolicyItemDailies
	}).(GetCloudBackupSchedulePolicyItemDailyArrayOutput)
}

// Hourly policy item
func (o LookupCloudBackupScheduleResultOutput) PolicyItemHourlies() GetCloudBackupSchedulePolicyItemHourlyArrayOutput {
	return o.ApplyT(func(v LookupCloudBackupScheduleResult) []GetCloudBackupSchedulePolicyItemHourly {
		return v.PolicyItemHourlies
	}).(GetCloudBackupSchedulePolicyItemHourlyArrayOutput)
}

// Monthly policy item
func (o LookupCloudBackupScheduleResultOutput) PolicyItemMonthlies() GetCloudBackupSchedulePolicyItemMonthlyArrayOutput {
	return o.ApplyT(func(v LookupCloudBackupScheduleResult) []GetCloudBackupSchedulePolicyItemMonthly {
		return v.PolicyItemMonthlies
	}).(GetCloudBackupSchedulePolicyItemMonthlyArrayOutput)
}

// Weekly policy item
func (o LookupCloudBackupScheduleResultOutput) PolicyItemWeeklies() GetCloudBackupSchedulePolicyItemWeeklyArrayOutput {
	return o.ApplyT(func(v LookupCloudBackupScheduleResult) []GetCloudBackupSchedulePolicyItemWeekly {
		return v.PolicyItemWeeklies
	}).(GetCloudBackupSchedulePolicyItemWeeklyArrayOutput)
}

func (o LookupCloudBackupScheduleResultOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudBackupScheduleResult) string { return v.ProjectId }).(pulumi.StringOutput)
}

// UTC Hour of day between 0 and 23 representing which hour of the day that Atlas takes a snapshot.
func (o LookupCloudBackupScheduleResultOutput) ReferenceHourOfDay() pulumi.IntOutput {
	return o.ApplyT(func(v LookupCloudBackupScheduleResult) int { return v.ReferenceHourOfDay }).(pulumi.IntOutput)
}

// UTC Minute of day between 0 and 59 representing which minute of the `referenceHourOfDay` that Atlas takes the snapshot.
func (o LookupCloudBackupScheduleResultOutput) ReferenceMinuteOfHour() pulumi.IntOutput {
	return o.ApplyT(func(v LookupCloudBackupScheduleResult) int { return v.ReferenceMinuteOfHour }).(pulumi.IntOutput)
}

// Specifies a restore window in days for cloud backup to maintain.
func (o LookupCloudBackupScheduleResultOutput) RestoreWindowDays() pulumi.IntOutput {
	return o.ApplyT(func(v LookupCloudBackupScheduleResult) int { return v.RestoreWindowDays }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCloudBackupScheduleResultOutput{})
}
