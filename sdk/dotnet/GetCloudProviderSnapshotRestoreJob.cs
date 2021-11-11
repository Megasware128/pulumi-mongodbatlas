// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi.Utilities;

namespace Pulumi.Mongodbatlas
{
    public static class GetCloudProviderSnapshotRestoreJob
    {
        /// <summary>
        /// `mongodbatlas.CloudProviderSnapshotRestoreJob` provides a Cloud Backup Snapshot Restore Job datasource. Gets all the cloud backup snapshot restore jobs for the specified cluster.
        /// 
        /// &gt; **NOTE:** Groups and projects are synonymous terms. You may find `groupId` in the official documentation.
        /// </summary>
        public static Task<GetCloudProviderSnapshotRestoreJobResult> InvokeAsync(GetCloudProviderSnapshotRestoreJobArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetCloudProviderSnapshotRestoreJobResult>("mongodbatlas:index/getCloudProviderSnapshotRestoreJob:getCloudProviderSnapshotRestoreJob", args ?? new GetCloudProviderSnapshotRestoreJobArgs(), options.WithVersion());

        /// <summary>
        /// `mongodbatlas.CloudProviderSnapshotRestoreJob` provides a Cloud Backup Snapshot Restore Job datasource. Gets all the cloud backup snapshot restore jobs for the specified cluster.
        /// 
        /// &gt; **NOTE:** Groups and projects are synonymous terms. You may find `groupId` in the official documentation.
        /// </summary>
        public static Output<GetCloudProviderSnapshotRestoreJobResult> Invoke(GetCloudProviderSnapshotRestoreJobInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetCloudProviderSnapshotRestoreJobResult>("mongodbatlas:index/getCloudProviderSnapshotRestoreJob:getCloudProviderSnapshotRestoreJob", args ?? new GetCloudProviderSnapshotRestoreJobInvokeArgs(), options.WithVersion());
    }


    public sealed class GetCloudProviderSnapshotRestoreJobArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the Atlas cluster for which you want to retrieve the restore job.
        /// </summary>
        [Input("clusterName", required: true)]
        public string ClusterName { get; set; } = null!;

        /// <summary>
        /// The unique identifier of the restore job to retrieve.
        /// </summary>
        [Input("jobId", required: true)]
        public string JobId { get; set; } = null!;

        /// <summary>
        /// The unique identifier of the project for the Atlas cluster.
        /// </summary>
        [Input("projectId", required: true)]
        public string ProjectId { get; set; } = null!;

        public GetCloudProviderSnapshotRestoreJobArgs()
        {
        }
    }

    public sealed class GetCloudProviderSnapshotRestoreJobInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the Atlas cluster for which you want to retrieve the restore job.
        /// </summary>
        [Input("clusterName", required: true)]
        public Input<string> ClusterName { get; set; } = null!;

        /// <summary>
        /// The unique identifier of the restore job to retrieve.
        /// </summary>
        [Input("jobId", required: true)]
        public Input<string> JobId { get; set; } = null!;

        /// <summary>
        /// The unique identifier of the project for the Atlas cluster.
        /// </summary>
        [Input("projectId", required: true)]
        public Input<string> ProjectId { get; set; } = null!;

        public GetCloudProviderSnapshotRestoreJobInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetCloudProviderSnapshotRestoreJobResult
    {
        /// <summary>
        /// Indicates whether the restore job was canceled.
        /// </summary>
        public readonly bool Cancelled;
        public readonly string ClusterName;
        /// <summary>
        /// UTC ISO 8601 formatted point in time when Atlas created the restore job.
        /// </summary>
        public readonly string CreatedAt;
        /// <summary>
        /// Type of restore job to create. Possible values are: automated and download.
        /// </summary>
        public readonly string DeliveryType;
        /// <summary>
        /// One or more URLs for the compressed snapshot files for manual download. Only visible if deliveryType is download.
        /// </summary>
        public readonly ImmutableArray<string> DeliveryUrls;
        /// <summary>
        /// Indicates whether the restore job expired.
        /// </summary>
        public readonly bool Expired;
        /// <summary>
        /// UTC ISO 8601 formatted point in time when the restore job expires.
        /// </summary>
        public readonly string ExpiresAt;
        /// <summary>
        /// UTC ISO 8601 formatted point in time when the restore job completed.
        /// </summary>
        public readonly string FinishedAt;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string JobId;
        public readonly int OplogInc;
        public readonly int OplogTs;
        public readonly int PointInTimeUtcSeconds;
        public readonly string ProjectId;
        /// <summary>
        /// Unique identifier of the source snapshot ID of the restore job.
        /// </summary>
        public readonly string SnapshotId;
        /// <summary>
        /// Name of the target Atlas cluster to which the restore job restores the snapshot. Only visible if deliveryType is automated.
        /// </summary>
        public readonly string TargetClusterName;
        public readonly string TargetProjectId;
        /// <summary>
        /// Timestamp in ISO 8601 date and time format in UTC when the snapshot associated to snapshotId was taken.
        /// </summary>
        public readonly string Timestamp;

        [OutputConstructor]
        private GetCloudProviderSnapshotRestoreJobResult(
            bool cancelled,

            string clusterName,

            string createdAt,

            string deliveryType,

            ImmutableArray<string> deliveryUrls,

            bool expired,

            string expiresAt,

            string finishedAt,

            string id,

            string jobId,

            int oplogInc,

            int oplogTs,

            int pointInTimeUtcSeconds,

            string projectId,

            string snapshotId,

            string targetClusterName,

            string targetProjectId,

            string timestamp)
        {
            Cancelled = cancelled;
            ClusterName = clusterName;
            CreatedAt = createdAt;
            DeliveryType = deliveryType;
            DeliveryUrls = deliveryUrls;
            Expired = expired;
            ExpiresAt = expiresAt;
            FinishedAt = finishedAt;
            Id = id;
            JobId = jobId;
            OplogInc = oplogInc;
            OplogTs = oplogTs;
            PointInTimeUtcSeconds = pointInTimeUtcSeconds;
            ProjectId = projectId;
            SnapshotId = snapshotId;
            TargetClusterName = targetClusterName;
            TargetProjectId = targetProjectId;
            Timestamp = timestamp;
        }
    }
}
