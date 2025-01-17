// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mongodbatlas

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// **WARNING:** This data source is deprecated, use `CloudBackupSchedule`
//
// `CloudProviderSnapshotBackupPolicy` provides a Cloud Backup Snapshot Backup Policy datasource. An Atlas Cloud Backup Snapshot Policy provides the current snapshot schedule and retention settings for the cluster.
//
// > **NOTE:** Groups and projects are synonymous terms. You may find `groupId` in the official documentation.
func LookupCloudProviderSnapshotBackupPolicy(ctx *pulumi.Context, args *LookupCloudProviderSnapshotBackupPolicyArgs, opts ...pulumi.InvokeOption) (*LookupCloudProviderSnapshotBackupPolicyResult, error) {
	var rv LookupCloudProviderSnapshotBackupPolicyResult
	err := ctx.Invoke("mongodbatlas:index/getCloudProviderSnapshotBackupPolicy:getCloudProviderSnapshotBackupPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCloudProviderSnapshotBackupPolicy.
type LookupCloudProviderSnapshotBackupPolicyArgs struct {
	// The name of the Atlas cluster that contains the snapshots backup policy you want to retrieve.
	ClusterName string `pulumi:"clusterName"`
	// The unique identifier of the project for the Atlas cluster.
	ProjectId string `pulumi:"projectId"`
}

// A collection of values returned by getCloudProviderSnapshotBackupPolicy.
type LookupCloudProviderSnapshotBackupPolicyResult struct {
	// Unique identifier of the Atlas cluster.
	ClusterId   string `pulumi:"clusterId"`
	ClusterName string `pulumi:"clusterName"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// UTC ISO 8601 formatted point in time when Atlas will take the next snapshot.
	NextSnapshot string `pulumi:"nextSnapshot"`
	// A list of policy definitions for the cluster.
	// * `policies.#.id` - Unique identifier of the backup policy.
	Policies  []GetCloudProviderSnapshotBackupPolicyPolicy `pulumi:"policies"`
	ProjectId string                                       `pulumi:"projectId"`
	// UTC Hour of day between 0 and 23 representing which hour of the day that Atlas takes a snapshot.
	ReferenceHourOfDay int `pulumi:"referenceHourOfDay"`
	// UTC Minute of day between 0 and 59 representing which minute of the referenceHourOfDay that Atlas takes the snapshot.
	ReferenceMinuteOfHour int `pulumi:"referenceMinuteOfHour"`
	// Specifies a restore window in days for cloud backup to maintain.
	RestoreWindowDays int  `pulumi:"restoreWindowDays"`
	UpdateSnapshots   bool `pulumi:"updateSnapshots"`
}

func LookupCloudProviderSnapshotBackupPolicyOutput(ctx *pulumi.Context, args LookupCloudProviderSnapshotBackupPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupCloudProviderSnapshotBackupPolicyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupCloudProviderSnapshotBackupPolicyResult, error) {
			args := v.(LookupCloudProviderSnapshotBackupPolicyArgs)
			r, err := LookupCloudProviderSnapshotBackupPolicy(ctx, &args, opts...)
			var s LookupCloudProviderSnapshotBackupPolicyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupCloudProviderSnapshotBackupPolicyResultOutput)
}

// A collection of arguments for invoking getCloudProviderSnapshotBackupPolicy.
type LookupCloudProviderSnapshotBackupPolicyOutputArgs struct {
	// The name of the Atlas cluster that contains the snapshots backup policy you want to retrieve.
	ClusterName pulumi.StringInput `pulumi:"clusterName"`
	// The unique identifier of the project for the Atlas cluster.
	ProjectId pulumi.StringInput `pulumi:"projectId"`
}

func (LookupCloudProviderSnapshotBackupPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCloudProviderSnapshotBackupPolicyArgs)(nil)).Elem()
}

// A collection of values returned by getCloudProviderSnapshotBackupPolicy.
type LookupCloudProviderSnapshotBackupPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupCloudProviderSnapshotBackupPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCloudProviderSnapshotBackupPolicyResult)(nil)).Elem()
}

func (o LookupCloudProviderSnapshotBackupPolicyResultOutput) ToLookupCloudProviderSnapshotBackupPolicyResultOutput() LookupCloudProviderSnapshotBackupPolicyResultOutput {
	return o
}

func (o LookupCloudProviderSnapshotBackupPolicyResultOutput) ToLookupCloudProviderSnapshotBackupPolicyResultOutputWithContext(ctx context.Context) LookupCloudProviderSnapshotBackupPolicyResultOutput {
	return o
}

// Unique identifier of the Atlas cluster.
func (o LookupCloudProviderSnapshotBackupPolicyResultOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudProviderSnapshotBackupPolicyResult) string { return v.ClusterId }).(pulumi.StringOutput)
}

func (o LookupCloudProviderSnapshotBackupPolicyResultOutput) ClusterName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudProviderSnapshotBackupPolicyResult) string { return v.ClusterName }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupCloudProviderSnapshotBackupPolicyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudProviderSnapshotBackupPolicyResult) string { return v.Id }).(pulumi.StringOutput)
}

// UTC ISO 8601 formatted point in time when Atlas will take the next snapshot.
func (o LookupCloudProviderSnapshotBackupPolicyResultOutput) NextSnapshot() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudProviderSnapshotBackupPolicyResult) string { return v.NextSnapshot }).(pulumi.StringOutput)
}

// A list of policy definitions for the cluster.
// * `policies.#.id` - Unique identifier of the backup policy.
func (o LookupCloudProviderSnapshotBackupPolicyResultOutput) Policies() GetCloudProviderSnapshotBackupPolicyPolicyArrayOutput {
	return o.ApplyT(func(v LookupCloudProviderSnapshotBackupPolicyResult) []GetCloudProviderSnapshotBackupPolicyPolicy {
		return v.Policies
	}).(GetCloudProviderSnapshotBackupPolicyPolicyArrayOutput)
}

func (o LookupCloudProviderSnapshotBackupPolicyResultOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudProviderSnapshotBackupPolicyResult) string { return v.ProjectId }).(pulumi.StringOutput)
}

// UTC Hour of day between 0 and 23 representing which hour of the day that Atlas takes a snapshot.
func (o LookupCloudProviderSnapshotBackupPolicyResultOutput) ReferenceHourOfDay() pulumi.IntOutput {
	return o.ApplyT(func(v LookupCloudProviderSnapshotBackupPolicyResult) int { return v.ReferenceHourOfDay }).(pulumi.IntOutput)
}

// UTC Minute of day between 0 and 59 representing which minute of the referenceHourOfDay that Atlas takes the snapshot.
func (o LookupCloudProviderSnapshotBackupPolicyResultOutput) ReferenceMinuteOfHour() pulumi.IntOutput {
	return o.ApplyT(func(v LookupCloudProviderSnapshotBackupPolicyResult) int { return v.ReferenceMinuteOfHour }).(pulumi.IntOutput)
}

// Specifies a restore window in days for cloud backup to maintain.
func (o LookupCloudProviderSnapshotBackupPolicyResultOutput) RestoreWindowDays() pulumi.IntOutput {
	return o.ApplyT(func(v LookupCloudProviderSnapshotBackupPolicyResult) int { return v.RestoreWindowDays }).(pulumi.IntOutput)
}

func (o LookupCloudProviderSnapshotBackupPolicyResultOutput) UpdateSnapshots() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupCloudProviderSnapshotBackupPolicyResult) bool { return v.UpdateSnapshots }).(pulumi.BoolOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCloudProviderSnapshotBackupPolicyResultOutput{})
}
