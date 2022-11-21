// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mongodbatlas

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// `OnlineArchive` describes an Online Archive
//
// > **NOTE:** Groups and projects are synonymous terms. You may find groupId in the official documentation.
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
//			_, err := mongodbatlas.LookupOnlineArchive(ctx, &GetOnlineArchiveArgs{
//				ProjectId:   _var.Project_id,
//				ClusterName: _var.Cluster_name,
//				ArchiveId:   "5ebad3c1fe9c0ab8d37d61e1",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ## Attributes reference
//
//   - `dbName`          -  Name of the database that contains the collection.
//   - `collName`        -  Name of the collection.
//   - `criteria`         -  Criteria to use for archiving data.
//   - `criteria.type`          - Type of criteria (DATE, CUSTOM)
//   - `criteria.date_field`    - Name of an already indexed date field from the documents. Data is archived when the current date is greater than the value of the date field specified here plus the number of days specified via the `expireAfterDays` parameter.
//   - `criteria.date_format`   - the date format. Valid values:  ISODATE (default), EPOCH_SECONDS, EPOCH_MILLIS, EPOCH_NANOSECONDS
//   - `criteria.expire_after_days` - Number of days that specifies the age limit for the data in the live Atlas cluster.
//   - `criteria.query` - JSON query to use to select documents for archiving. Only for `CUSTOM` type
//   - `partitionFields` -  Fields to use to partition data.
//   - `partition_fields.field_name` - Name of the field. To specify a nested field, use the dot notation.
//   - `partition_fields.order` - Position of the field in the partition. Value can be: 0,1,2
//     By default, the date field specified in the criteria.dateField parameter is in the first position of the partition.
//   - `partitio_fields.field_type` - Type of the partition field
//   - `state`    - Status of the online archive. Valid values are: Pending, Archiving, Idle, Pausing, Paused, Orphaned and Deleted
//
// See [MongoDB Atlas API](https://docs.atlas.mongodb.com/reference/api/online-archive-get-one/) Documentation for more information.
func LookupOnlineArchive(ctx *pulumi.Context, args *LookupOnlineArchiveArgs, opts ...pulumi.InvokeOption) (*LookupOnlineArchiveResult, error) {
	var rv LookupOnlineArchiveResult
	err := ctx.Invoke("mongodbatlas:index/getOnlineArchive:getOnlineArchive", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getOnlineArchive.
type LookupOnlineArchiveArgs struct {
	// ID of the online archive.
	ArchiveId string `pulumi:"archiveId"`
	// Name of the cluster that contains the collection.
	ClusterName string `pulumi:"clusterName"`
	// The unique ID for the project.
	ProjectId string `pulumi:"projectId"`
}

// A collection of values returned by getOnlineArchive.
type LookupOnlineArchiveResult struct {
	ArchiveId   string                     `pulumi:"archiveId"`
	ClusterName string                     `pulumi:"clusterName"`
	CollName    string                     `pulumi:"collName"`
	Criterias   []GetOnlineArchiveCriteria `pulumi:"criterias"`
	DbName      string                     `pulumi:"dbName"`
	// The provider-assigned unique ID for this managed resource.
	Id              string                           `pulumi:"id"`
	PartitionFields []GetOnlineArchivePartitionField `pulumi:"partitionFields"`
	Paused          bool                             `pulumi:"paused"`
	ProjectId       string                           `pulumi:"projectId"`
	State           string                           `pulumi:"state"`
}

func LookupOnlineArchiveOutput(ctx *pulumi.Context, args LookupOnlineArchiveOutputArgs, opts ...pulumi.InvokeOption) LookupOnlineArchiveResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupOnlineArchiveResult, error) {
			args := v.(LookupOnlineArchiveArgs)
			r, err := LookupOnlineArchive(ctx, &args, opts...)
			var s LookupOnlineArchiveResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupOnlineArchiveResultOutput)
}

// A collection of arguments for invoking getOnlineArchive.
type LookupOnlineArchiveOutputArgs struct {
	// ID of the online archive.
	ArchiveId pulumi.StringInput `pulumi:"archiveId"`
	// Name of the cluster that contains the collection.
	ClusterName pulumi.StringInput `pulumi:"clusterName"`
	// The unique ID for the project.
	ProjectId pulumi.StringInput `pulumi:"projectId"`
}

func (LookupOnlineArchiveOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOnlineArchiveArgs)(nil)).Elem()
}

// A collection of values returned by getOnlineArchive.
type LookupOnlineArchiveResultOutput struct{ *pulumi.OutputState }

func (LookupOnlineArchiveResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOnlineArchiveResult)(nil)).Elem()
}

func (o LookupOnlineArchiveResultOutput) ToLookupOnlineArchiveResultOutput() LookupOnlineArchiveResultOutput {
	return o
}

func (o LookupOnlineArchiveResultOutput) ToLookupOnlineArchiveResultOutputWithContext(ctx context.Context) LookupOnlineArchiveResultOutput {
	return o
}

func (o LookupOnlineArchiveResultOutput) ArchiveId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOnlineArchiveResult) string { return v.ArchiveId }).(pulumi.StringOutput)
}

func (o LookupOnlineArchiveResultOutput) ClusterName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOnlineArchiveResult) string { return v.ClusterName }).(pulumi.StringOutput)
}

func (o LookupOnlineArchiveResultOutput) CollName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOnlineArchiveResult) string { return v.CollName }).(pulumi.StringOutput)
}

func (o LookupOnlineArchiveResultOutput) Criterias() GetOnlineArchiveCriteriaArrayOutput {
	return o.ApplyT(func(v LookupOnlineArchiveResult) []GetOnlineArchiveCriteria { return v.Criterias }).(GetOnlineArchiveCriteriaArrayOutput)
}

func (o LookupOnlineArchiveResultOutput) DbName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOnlineArchiveResult) string { return v.DbName }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupOnlineArchiveResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOnlineArchiveResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupOnlineArchiveResultOutput) PartitionFields() GetOnlineArchivePartitionFieldArrayOutput {
	return o.ApplyT(func(v LookupOnlineArchiveResult) []GetOnlineArchivePartitionField { return v.PartitionFields }).(GetOnlineArchivePartitionFieldArrayOutput)
}

func (o LookupOnlineArchiveResultOutput) Paused() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupOnlineArchiveResult) bool { return v.Paused }).(pulumi.BoolOutput)
}

func (o LookupOnlineArchiveResultOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOnlineArchiveResult) string { return v.ProjectId }).(pulumi.StringOutput)
}

func (o LookupOnlineArchiveResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOnlineArchiveResult) string { return v.State }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupOnlineArchiveResultOutput{})
}
