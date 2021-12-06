// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package mongodbatlas

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// `getCloudProviderSnapshotRestoreJobs` provides a Cloud Backup Snapshot Restore Jobs datasource. Gets all the cloud backup snapshot restore jobs for the specified cluster.
//
// > **NOTE:** Groups and projects are synonymous terms. You may find `groupId` in the official documentation.
//
// ## Example Usage
//
// First create a snapshot of the desired cluster. Then request that snapshot be restored in an automated fashion to the designated cluster and project.
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
// 		testCloudProviderSnapshot, err := mongodbatlas.NewCloudProviderSnapshot(ctx, "testCloudProviderSnapshot", &mongodbatlas.CloudProviderSnapshotArgs{
// 			ProjectId:       pulumi.String("5cf5a45a9ccf6400e60981b6"),
// 			ClusterName:     pulumi.String("MyCluster"),
// 			Description:     pulumi.String("MyDescription"),
// 			RetentionInDays: pulumi.Int(1),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		testCloudProviderSnapshotRestoreJob, err := mongodbatlas.NewCloudProviderSnapshotRestoreJob(ctx, "testCloudProviderSnapshotRestoreJob", &mongodbatlas.CloudProviderSnapshotRestoreJobArgs{
// 			ProjectId:   pulumi.String("5cf5a45a9ccf6400e60981b6"),
// 			ClusterName: pulumi.String("MyCluster"),
// 			SnapshotId:  testCloudProviderSnapshot.ID(),
// 			DeliveryType: pulumi.StringMap{
// 				pulumi.String{
// 					Automated:         true,
// 					TargetClusterName: "MyCluster",
// 					TargetProjectId:   "5cf5a45a9ccf6400e60981b6",
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupCloudProviderSnapshotRestoreJobs(ctx *pulumi.Context, args *LookupCloudProviderSnapshotRestoreJobsArgs, opts ...pulumi.InvokeOption) (*LookupCloudProviderSnapshotRestoreJobsResult, error) {
	var rv LookupCloudProviderSnapshotRestoreJobsResult
	err := ctx.Invoke("mongodbatlas:index/getCloudProviderSnapshotRestoreJobs:getCloudProviderSnapshotRestoreJobs", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCloudProviderSnapshotRestoreJobs.
type LookupCloudProviderSnapshotRestoreJobsArgs struct {
	// The name of the Atlas cluster for which you want to retrieve restore jobs.
	ClusterName string `pulumi:"clusterName"`
	// Number of items to return per page, up to a maximum of 500. Defaults to `100`.
	ItemsPerPage *int `pulumi:"itemsPerPage"`
	// The page to return. Defaults to `1`.
	PageNum *int `pulumi:"pageNum"`
	// The unique identifier of the project for the Atlas cluster.
	ProjectId string `pulumi:"projectId"`
}

// A collection of values returned by getCloudProviderSnapshotRestoreJobs.
type LookupCloudProviderSnapshotRestoreJobsResult struct {
	ClusterName string `pulumi:"clusterName"`
	// The provider-assigned unique ID for this managed resource.
	Id           string `pulumi:"id"`
	ItemsPerPage *int   `pulumi:"itemsPerPage"`
	PageNum      *int   `pulumi:"pageNum"`
	ProjectId    string `pulumi:"projectId"`
	// Includes cloudProviderSnapshotRestoreJob object for each item detailed in the results array section.
	Results    []GetCloudProviderSnapshotRestoreJobsResult `pulumi:"results"`
	TotalCount int                                         `pulumi:"totalCount"`
}

func LookupCloudProviderSnapshotRestoreJobsOutput(ctx *pulumi.Context, args LookupCloudProviderSnapshotRestoreJobsOutputArgs, opts ...pulumi.InvokeOption) LookupCloudProviderSnapshotRestoreJobsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupCloudProviderSnapshotRestoreJobsResult, error) {
			args := v.(LookupCloudProviderSnapshotRestoreJobsArgs)
			r, err := LookupCloudProviderSnapshotRestoreJobs(ctx, &args, opts...)
			return *r, err
		}).(LookupCloudProviderSnapshotRestoreJobsResultOutput)
}

// A collection of arguments for invoking getCloudProviderSnapshotRestoreJobs.
type LookupCloudProviderSnapshotRestoreJobsOutputArgs struct {
	// The name of the Atlas cluster for which you want to retrieve restore jobs.
	ClusterName pulumi.StringInput `pulumi:"clusterName"`
	// Number of items to return per page, up to a maximum of 500. Defaults to `100`.
	ItemsPerPage pulumi.IntPtrInput `pulumi:"itemsPerPage"`
	// The page to return. Defaults to `1`.
	PageNum pulumi.IntPtrInput `pulumi:"pageNum"`
	// The unique identifier of the project for the Atlas cluster.
	ProjectId pulumi.StringInput `pulumi:"projectId"`
}

func (LookupCloudProviderSnapshotRestoreJobsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCloudProviderSnapshotRestoreJobsArgs)(nil)).Elem()
}

// A collection of values returned by getCloudProviderSnapshotRestoreJobs.
type LookupCloudProviderSnapshotRestoreJobsResultOutput struct{ *pulumi.OutputState }

func (LookupCloudProviderSnapshotRestoreJobsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCloudProviderSnapshotRestoreJobsResult)(nil)).Elem()
}

func (o LookupCloudProviderSnapshotRestoreJobsResultOutput) ToLookupCloudProviderSnapshotRestoreJobsResultOutput() LookupCloudProviderSnapshotRestoreJobsResultOutput {
	return o
}

func (o LookupCloudProviderSnapshotRestoreJobsResultOutput) ToLookupCloudProviderSnapshotRestoreJobsResultOutputWithContext(ctx context.Context) LookupCloudProviderSnapshotRestoreJobsResultOutput {
	return o
}

func (o LookupCloudProviderSnapshotRestoreJobsResultOutput) ClusterName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudProviderSnapshotRestoreJobsResult) string { return v.ClusterName }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupCloudProviderSnapshotRestoreJobsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudProviderSnapshotRestoreJobsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupCloudProviderSnapshotRestoreJobsResultOutput) ItemsPerPage() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupCloudProviderSnapshotRestoreJobsResult) *int { return v.ItemsPerPage }).(pulumi.IntPtrOutput)
}

func (o LookupCloudProviderSnapshotRestoreJobsResultOutput) PageNum() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupCloudProviderSnapshotRestoreJobsResult) *int { return v.PageNum }).(pulumi.IntPtrOutput)
}

func (o LookupCloudProviderSnapshotRestoreJobsResultOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudProviderSnapshotRestoreJobsResult) string { return v.ProjectId }).(pulumi.StringOutput)
}

// Includes cloudProviderSnapshotRestoreJob object for each item detailed in the results array section.
func (o LookupCloudProviderSnapshotRestoreJobsResultOutput) Results() GetCloudProviderSnapshotRestoreJobsResultArrayOutput {
	return o.ApplyT(func(v LookupCloudProviderSnapshotRestoreJobsResult) []GetCloudProviderSnapshotRestoreJobsResult {
		return v.Results
	}).(GetCloudProviderSnapshotRestoreJobsResultArrayOutput)
}

func (o LookupCloudProviderSnapshotRestoreJobsResultOutput) TotalCount() pulumi.IntOutput {
	return o.ApplyT(func(v LookupCloudProviderSnapshotRestoreJobsResult) int { return v.TotalCount }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCloudProviderSnapshotRestoreJobsResultOutput{})
}
