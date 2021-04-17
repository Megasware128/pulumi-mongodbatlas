// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package mongodbatlas

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

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
