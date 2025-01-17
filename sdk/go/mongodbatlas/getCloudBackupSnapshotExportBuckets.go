// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mongodbatlas

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// `getCloudBackupSnapshotExportBuckets` datasource allows you to retrieve all the buckets for the specified project.
//
// > **NOTE:** Groups and projects are synonymous terms. You may find `groupId` in the official documentation.
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
//			_, err := mongodbatlas.NewCloudBackupSnapshotExportBucket(ctx, "testCloudBackupSnapshotExportBucket", &mongodbatlas.CloudBackupSnapshotExportBucketArgs{
//				BucketName:    pulumi.String("example-bucket"),
//				CloudProvider: pulumi.String("AWS"),
//				IamRoleId:     pulumi.String("{IAM_ROLE_ID}"),
//				ProjectId:     pulumi.String("{PROJECT_ID}"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = mongodbatlas.LookupCloudBackupSnapshotExportBuckets(ctx, &GetCloudBackupSnapshotExportBucketsArgs{
//				ProjectId: "{PROJECT_ID}",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupCloudBackupSnapshotExportBuckets(ctx *pulumi.Context, args *LookupCloudBackupSnapshotExportBucketsArgs, opts ...pulumi.InvokeOption) (*LookupCloudBackupSnapshotExportBucketsResult, error) {
	var rv LookupCloudBackupSnapshotExportBucketsResult
	err := ctx.Invoke("mongodbatlas:index/getCloudBackupSnapshotExportBuckets:getCloudBackupSnapshotExportBuckets", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCloudBackupSnapshotExportBuckets.
type LookupCloudBackupSnapshotExportBucketsArgs struct {
	// Number of items to return per page, up to a maximum of 500. Defaults to `100`.
	ItemsPerPage *int `pulumi:"itemsPerPage"`
	// The page to return. Defaults to `1`.
	PageNum *int `pulumi:"pageNum"`
	// The unique identifier of the project for the Atlas cluster.
	ProjectId string `pulumi:"projectId"`
}

// A collection of values returned by getCloudBackupSnapshotExportBuckets.
type LookupCloudBackupSnapshotExportBucketsResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id           string `pulumi:"id"`
	ItemsPerPage *int   `pulumi:"itemsPerPage"`
	PageNum      *int   `pulumi:"pageNum"`
	// The unique identifier of the project for the Atlas cluster.
	ProjectId string `pulumi:"projectId"`
	// Includes CloudProviderSnapshotExportBucket object for each item detailed in the results array section.
	Results    []GetCloudBackupSnapshotExportBucketsResult `pulumi:"results"`
	TotalCount int                                         `pulumi:"totalCount"`
}

func LookupCloudBackupSnapshotExportBucketsOutput(ctx *pulumi.Context, args LookupCloudBackupSnapshotExportBucketsOutputArgs, opts ...pulumi.InvokeOption) LookupCloudBackupSnapshotExportBucketsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupCloudBackupSnapshotExportBucketsResult, error) {
			args := v.(LookupCloudBackupSnapshotExportBucketsArgs)
			r, err := LookupCloudBackupSnapshotExportBuckets(ctx, &args, opts...)
			var s LookupCloudBackupSnapshotExportBucketsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupCloudBackupSnapshotExportBucketsResultOutput)
}

// A collection of arguments for invoking getCloudBackupSnapshotExportBuckets.
type LookupCloudBackupSnapshotExportBucketsOutputArgs struct {
	// Number of items to return per page, up to a maximum of 500. Defaults to `100`.
	ItemsPerPage pulumi.IntPtrInput `pulumi:"itemsPerPage"`
	// The page to return. Defaults to `1`.
	PageNum pulumi.IntPtrInput `pulumi:"pageNum"`
	// The unique identifier of the project for the Atlas cluster.
	ProjectId pulumi.StringInput `pulumi:"projectId"`
}

func (LookupCloudBackupSnapshotExportBucketsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCloudBackupSnapshotExportBucketsArgs)(nil)).Elem()
}

// A collection of values returned by getCloudBackupSnapshotExportBuckets.
type LookupCloudBackupSnapshotExportBucketsResultOutput struct{ *pulumi.OutputState }

func (LookupCloudBackupSnapshotExportBucketsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCloudBackupSnapshotExportBucketsResult)(nil)).Elem()
}

func (o LookupCloudBackupSnapshotExportBucketsResultOutput) ToLookupCloudBackupSnapshotExportBucketsResultOutput() LookupCloudBackupSnapshotExportBucketsResultOutput {
	return o
}

func (o LookupCloudBackupSnapshotExportBucketsResultOutput) ToLookupCloudBackupSnapshotExportBucketsResultOutputWithContext(ctx context.Context) LookupCloudBackupSnapshotExportBucketsResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o LookupCloudBackupSnapshotExportBucketsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudBackupSnapshotExportBucketsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupCloudBackupSnapshotExportBucketsResultOutput) ItemsPerPage() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupCloudBackupSnapshotExportBucketsResult) *int { return v.ItemsPerPage }).(pulumi.IntPtrOutput)
}

func (o LookupCloudBackupSnapshotExportBucketsResultOutput) PageNum() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupCloudBackupSnapshotExportBucketsResult) *int { return v.PageNum }).(pulumi.IntPtrOutput)
}

// The unique identifier of the project for the Atlas cluster.
func (o LookupCloudBackupSnapshotExportBucketsResultOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudBackupSnapshotExportBucketsResult) string { return v.ProjectId }).(pulumi.StringOutput)
}

// Includes CloudProviderSnapshotExportBucket object for each item detailed in the results array section.
func (o LookupCloudBackupSnapshotExportBucketsResultOutput) Results() GetCloudBackupSnapshotExportBucketsResultArrayOutput {
	return o.ApplyT(func(v LookupCloudBackupSnapshotExportBucketsResult) []GetCloudBackupSnapshotExportBucketsResult {
		return v.Results
	}).(GetCloudBackupSnapshotExportBucketsResultArrayOutput)
}

func (o LookupCloudBackupSnapshotExportBucketsResultOutput) TotalCount() pulumi.IntOutput {
	return o.ApplyT(func(v LookupCloudBackupSnapshotExportBucketsResult) int { return v.TotalCount }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCloudBackupSnapshotExportBucketsResultOutput{})
}
