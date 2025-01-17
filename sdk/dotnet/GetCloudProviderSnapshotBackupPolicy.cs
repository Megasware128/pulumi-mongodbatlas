// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Mongodbatlas
{
    public static class GetCloudProviderSnapshotBackupPolicy
    {
        /// <summary>
        /// **WARNING:** This data source is deprecated, use `mongodbatlas.CloudBackupSchedule`
        /// 
        /// `mongodbatlas.CloudProviderSnapshotBackupPolicy` provides a Cloud Backup Snapshot Backup Policy datasource. An Atlas Cloud Backup Snapshot Policy provides the current snapshot schedule and retention settings for the cluster. 
        /// 
        /// &gt; **NOTE:** Groups and projects are synonymous terms. You may find `groupId` in the official documentation.
        /// </summary>
        public static Task<GetCloudProviderSnapshotBackupPolicyResult> InvokeAsync(GetCloudProviderSnapshotBackupPolicyArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetCloudProviderSnapshotBackupPolicyResult>("mongodbatlas:index/getCloudProviderSnapshotBackupPolicy:getCloudProviderSnapshotBackupPolicy", args ?? new GetCloudProviderSnapshotBackupPolicyArgs(), options.WithDefaults());

        /// <summary>
        /// **WARNING:** This data source is deprecated, use `mongodbatlas.CloudBackupSchedule`
        /// 
        /// `mongodbatlas.CloudProviderSnapshotBackupPolicy` provides a Cloud Backup Snapshot Backup Policy datasource. An Atlas Cloud Backup Snapshot Policy provides the current snapshot schedule and retention settings for the cluster. 
        /// 
        /// &gt; **NOTE:** Groups and projects are synonymous terms. You may find `groupId` in the official documentation.
        /// </summary>
        public static Output<GetCloudProviderSnapshotBackupPolicyResult> Invoke(GetCloudProviderSnapshotBackupPolicyInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetCloudProviderSnapshotBackupPolicyResult>("mongodbatlas:index/getCloudProviderSnapshotBackupPolicy:getCloudProviderSnapshotBackupPolicy", args ?? new GetCloudProviderSnapshotBackupPolicyInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetCloudProviderSnapshotBackupPolicyArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the Atlas cluster that contains the snapshots backup policy you want to retrieve.
        /// </summary>
        [Input("clusterName", required: true)]
        public string ClusterName { get; set; } = null!;

        /// <summary>
        /// The unique identifier of the project for the Atlas cluster.
        /// </summary>
        [Input("projectId", required: true)]
        public string ProjectId { get; set; } = null!;

        public GetCloudProviderSnapshotBackupPolicyArgs()
        {
        }
        public static new GetCloudProviderSnapshotBackupPolicyArgs Empty => new GetCloudProviderSnapshotBackupPolicyArgs();
    }

    public sealed class GetCloudProviderSnapshotBackupPolicyInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the Atlas cluster that contains the snapshots backup policy you want to retrieve.
        /// </summary>
        [Input("clusterName", required: true)]
        public Input<string> ClusterName { get; set; } = null!;

        /// <summary>
        /// The unique identifier of the project for the Atlas cluster.
        /// </summary>
        [Input("projectId", required: true)]
        public Input<string> ProjectId { get; set; } = null!;

        public GetCloudProviderSnapshotBackupPolicyInvokeArgs()
        {
        }
        public static new GetCloudProviderSnapshotBackupPolicyInvokeArgs Empty => new GetCloudProviderSnapshotBackupPolicyInvokeArgs();
    }


    [OutputType]
    public sealed class GetCloudProviderSnapshotBackupPolicyResult
    {
        /// <summary>
        /// Unique identifier of the Atlas cluster.
        /// </summary>
        public readonly string ClusterId;
        public readonly string ClusterName;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// UTC ISO 8601 formatted point in time when Atlas will take the next snapshot.
        /// </summary>
        public readonly string NextSnapshot;
        /// <summary>
        /// A list of policy definitions for the cluster.
        /// * `policies.#.id` - Unique identifier of the backup policy.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetCloudProviderSnapshotBackupPolicyPolicyResult> Policies;
        public readonly string ProjectId;
        /// <summary>
        /// UTC Hour of day between 0 and 23 representing which hour of the day that Atlas takes a snapshot.
        /// </summary>
        public readonly int ReferenceHourOfDay;
        /// <summary>
        /// UTC Minute of day between 0 and 59 representing which minute of the referenceHourOfDay that Atlas takes the snapshot.
        /// </summary>
        public readonly int ReferenceMinuteOfHour;
        /// <summary>
        /// Specifies a restore window in days for cloud backup to maintain.
        /// </summary>
        public readonly int RestoreWindowDays;
        public readonly bool UpdateSnapshots;

        [OutputConstructor]
        private GetCloudProviderSnapshotBackupPolicyResult(
            string clusterId,

            string clusterName,

            string id,

            string nextSnapshot,

            ImmutableArray<Outputs.GetCloudProviderSnapshotBackupPolicyPolicyResult> policies,

            string projectId,

            int referenceHourOfDay,

            int referenceMinuteOfHour,

            int restoreWindowDays,

            bool updateSnapshots)
        {
            ClusterId = clusterId;
            ClusterName = clusterName;
            Id = id;
            NextSnapshot = nextSnapshot;
            Policies = policies;
            ProjectId = projectId;
            ReferenceHourOfDay = referenceHourOfDay;
            ReferenceMinuteOfHour = referenceMinuteOfHour;
            RestoreWindowDays = restoreWindowDays;
            UpdateSnapshots = updateSnapshots;
        }
    }
}
