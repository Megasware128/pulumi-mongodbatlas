// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package mongodbatlas

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// `CloudProviderSnapshotRestoreJob` provides a Cloud Backup Snapshot Restore Job datasource. Gets all the cloud backup snapshot restore jobs for the specified cluster.
//
// > **NOTE:** Groups and projects are synonymous terms. You may find `groupId` in the official documentation.
func LookupCloudProviderSnapshotRestoreJob(ctx *pulumi.Context, args *LookupCloudProviderSnapshotRestoreJobArgs, opts ...pulumi.InvokeOption) (*LookupCloudProviderSnapshotRestoreJobResult, error) {
	var rv LookupCloudProviderSnapshotRestoreJobResult
	err := ctx.Invoke("mongodbatlas:index/getCloudProviderSnapshotRestoreJob:getCloudProviderSnapshotRestoreJob", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCloudProviderSnapshotRestoreJob.
type LookupCloudProviderSnapshotRestoreJobArgs struct {
	// The name of the Atlas cluster for which you want to retrieve the restore job.
	ClusterName string `pulumi:"clusterName"`
	// The unique identifier of the restore job to retrieve.
	JobId string `pulumi:"jobId"`
	// The unique identifier of the project for the Atlas cluster.
	ProjectId string `pulumi:"projectId"`
}

// A collection of values returned by getCloudProviderSnapshotRestoreJob.
type LookupCloudProviderSnapshotRestoreJobResult struct {
	// Indicates whether the restore job was canceled.
	Cancelled   bool   `pulumi:"cancelled"`
	ClusterName string `pulumi:"clusterName"`
	// UTC ISO 8601 formatted point in time when Atlas created the restore job.
	CreatedAt string `pulumi:"createdAt"`
	// Type of restore job to create. Possible values are: automated and download.
	DeliveryType string `pulumi:"deliveryType"`
	// One or more URLs for the compressed snapshot files for manual download. Only visible if deliveryType is download.
	DeliveryUrls []string `pulumi:"deliveryUrls"`
	// Indicates whether the restore job expired.
	Expired bool `pulumi:"expired"`
	// UTC ISO 8601 formatted point in time when the restore job expires.
	ExpiresAt string `pulumi:"expiresAt"`
	// UTC ISO 8601 formatted point in time when the restore job completed.
	FinishedAt string `pulumi:"finishedAt"`
	// The provider-assigned unique ID for this managed resource.
	Id                    string `pulumi:"id"`
	JobId                 string `pulumi:"jobId"`
	OplogInc              int    `pulumi:"oplogInc"`
	OplogTs               int    `pulumi:"oplogTs"`
	PointInTimeUtcSeconds int    `pulumi:"pointInTimeUtcSeconds"`
	ProjectId             string `pulumi:"projectId"`
	// Unique identifier of the source snapshot ID of the restore job.
	SnapshotId string `pulumi:"snapshotId"`
	// Name of the target Atlas cluster to which the restore job restores the snapshot. Only visible if deliveryType is automated.
	TargetClusterName string `pulumi:"targetClusterName"`
	TargetProjectId   string `pulumi:"targetProjectId"`
	// Timestamp in ISO 8601 date and time format in UTC when the snapshot associated to snapshotId was taken.
	Timestamp string `pulumi:"timestamp"`
}

func LookupCloudProviderSnapshotRestoreJobOutput(ctx *pulumi.Context, args LookupCloudProviderSnapshotRestoreJobOutputArgs, opts ...pulumi.InvokeOption) LookupCloudProviderSnapshotRestoreJobResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupCloudProviderSnapshotRestoreJobResult, error) {
			args := v.(LookupCloudProviderSnapshotRestoreJobArgs)
			r, err := LookupCloudProviderSnapshotRestoreJob(ctx, &args, opts...)
			return *r, err
		}).(LookupCloudProviderSnapshotRestoreJobResultOutput)
}

// A collection of arguments for invoking getCloudProviderSnapshotRestoreJob.
type LookupCloudProviderSnapshotRestoreJobOutputArgs struct {
	// The name of the Atlas cluster for which you want to retrieve the restore job.
	ClusterName pulumi.StringInput `pulumi:"clusterName"`
	// The unique identifier of the restore job to retrieve.
	JobId pulumi.StringInput `pulumi:"jobId"`
	// The unique identifier of the project for the Atlas cluster.
	ProjectId pulumi.StringInput `pulumi:"projectId"`
}

func (LookupCloudProviderSnapshotRestoreJobOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCloudProviderSnapshotRestoreJobArgs)(nil)).Elem()
}

// A collection of values returned by getCloudProviderSnapshotRestoreJob.
type LookupCloudProviderSnapshotRestoreJobResultOutput struct{ *pulumi.OutputState }

func (LookupCloudProviderSnapshotRestoreJobResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCloudProviderSnapshotRestoreJobResult)(nil)).Elem()
}

func (o LookupCloudProviderSnapshotRestoreJobResultOutput) ToLookupCloudProviderSnapshotRestoreJobResultOutput() LookupCloudProviderSnapshotRestoreJobResultOutput {
	return o
}

func (o LookupCloudProviderSnapshotRestoreJobResultOutput) ToLookupCloudProviderSnapshotRestoreJobResultOutputWithContext(ctx context.Context) LookupCloudProviderSnapshotRestoreJobResultOutput {
	return o
}

// Indicates whether the restore job was canceled.
func (o LookupCloudProviderSnapshotRestoreJobResultOutput) Cancelled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupCloudProviderSnapshotRestoreJobResult) bool { return v.Cancelled }).(pulumi.BoolOutput)
}

func (o LookupCloudProviderSnapshotRestoreJobResultOutput) ClusterName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudProviderSnapshotRestoreJobResult) string { return v.ClusterName }).(pulumi.StringOutput)
}

// UTC ISO 8601 formatted point in time when Atlas created the restore job.
func (o LookupCloudProviderSnapshotRestoreJobResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudProviderSnapshotRestoreJobResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

// Type of restore job to create. Possible values are: automated and download.
func (o LookupCloudProviderSnapshotRestoreJobResultOutput) DeliveryType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudProviderSnapshotRestoreJobResult) string { return v.DeliveryType }).(pulumi.StringOutput)
}

// One or more URLs for the compressed snapshot files for manual download. Only visible if deliveryType is download.
func (o LookupCloudProviderSnapshotRestoreJobResultOutput) DeliveryUrls() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupCloudProviderSnapshotRestoreJobResult) []string { return v.DeliveryUrls }).(pulumi.StringArrayOutput)
}

// Indicates whether the restore job expired.
func (o LookupCloudProviderSnapshotRestoreJobResultOutput) Expired() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupCloudProviderSnapshotRestoreJobResult) bool { return v.Expired }).(pulumi.BoolOutput)
}

// UTC ISO 8601 formatted point in time when the restore job expires.
func (o LookupCloudProviderSnapshotRestoreJobResultOutput) ExpiresAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudProviderSnapshotRestoreJobResult) string { return v.ExpiresAt }).(pulumi.StringOutput)
}

// UTC ISO 8601 formatted point in time when the restore job completed.
func (o LookupCloudProviderSnapshotRestoreJobResultOutput) FinishedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudProviderSnapshotRestoreJobResult) string { return v.FinishedAt }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupCloudProviderSnapshotRestoreJobResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudProviderSnapshotRestoreJobResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupCloudProviderSnapshotRestoreJobResultOutput) JobId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudProviderSnapshotRestoreJobResult) string { return v.JobId }).(pulumi.StringOutput)
}

func (o LookupCloudProviderSnapshotRestoreJobResultOutput) OplogInc() pulumi.IntOutput {
	return o.ApplyT(func(v LookupCloudProviderSnapshotRestoreJobResult) int { return v.OplogInc }).(pulumi.IntOutput)
}

func (o LookupCloudProviderSnapshotRestoreJobResultOutput) OplogTs() pulumi.IntOutput {
	return o.ApplyT(func(v LookupCloudProviderSnapshotRestoreJobResult) int { return v.OplogTs }).(pulumi.IntOutput)
}

func (o LookupCloudProviderSnapshotRestoreJobResultOutput) PointInTimeUtcSeconds() pulumi.IntOutput {
	return o.ApplyT(func(v LookupCloudProviderSnapshotRestoreJobResult) int { return v.PointInTimeUtcSeconds }).(pulumi.IntOutput)
}

func (o LookupCloudProviderSnapshotRestoreJobResultOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudProviderSnapshotRestoreJobResult) string { return v.ProjectId }).(pulumi.StringOutput)
}

// Unique identifier of the source snapshot ID of the restore job.
func (o LookupCloudProviderSnapshotRestoreJobResultOutput) SnapshotId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudProviderSnapshotRestoreJobResult) string { return v.SnapshotId }).(pulumi.StringOutput)
}

// Name of the target Atlas cluster to which the restore job restores the snapshot. Only visible if deliveryType is automated.
func (o LookupCloudProviderSnapshotRestoreJobResultOutput) TargetClusterName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudProviderSnapshotRestoreJobResult) string { return v.TargetClusterName }).(pulumi.StringOutput)
}

func (o LookupCloudProviderSnapshotRestoreJobResultOutput) TargetProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudProviderSnapshotRestoreJobResult) string { return v.TargetProjectId }).(pulumi.StringOutput)
}

// Timestamp in ISO 8601 date and time format in UTC when the snapshot associated to snapshotId was taken.
func (o LookupCloudProviderSnapshotRestoreJobResultOutput) Timestamp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudProviderSnapshotRestoreJobResult) string { return v.Timestamp }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCloudProviderSnapshotRestoreJobResultOutput{})
}
