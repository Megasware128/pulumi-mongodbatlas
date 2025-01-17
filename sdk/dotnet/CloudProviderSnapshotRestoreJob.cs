// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Mongodbatlas
{
    /// <summary>
    /// **WARNING:** This resource is deprecated, use `mongodbatlas.CloudBackupSnapshotRestoreJob`
    /// 
    /// `mongodbatlas.CloudProviderSnapshotRestoreJob` provides a resource to create a new restore job from a cloud backup snapshot of a specified cluster. The restore job can be one of three types:
    /// * **automated:** Atlas automatically restores the snapshot with snapshotId to the Atlas cluster with name targetClusterName in the Atlas project with targetGroupId.
    /// 
    /// * **download:** Atlas provides a URL to download a .tar.gz of the snapshot with snapshotId. The contents of the archive contain the data files for your Atlas cluster.
    /// 
    /// * **pointInTime:**  Atlas performs a Continuous Cloud Backup restore.
    /// 
    /// &gt; **Important:** If you specify `deliveryType` : `automated` or `deliveryType` : `pointInTime` in your request body to create an automated restore job, Atlas removes all existing data on the target cluster prior to the restore.
    /// 
    /// &gt; **NOTE:** Groups and projects are synonymous terms. You may find `groupId` in the official documentation.
    /// 
    /// ## Example Usage
    /// ### Example automated delivery type.
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using Mongodbatlas = Pulumi.Mongodbatlas;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var myCluster = new Mongodbatlas.Cluster("myCluster", new()
    ///     {
    ///         ProjectId = "5cf5a45a9ccf6400e60981b6",
    ///         DiskSizeGb = 5,
    ///         ProviderName = "AWS",
    ///         ProviderRegionName = "EU_WEST_2",
    ///         ProviderInstanceSizeName = "M10",
    ///         CloudBackup = true,
    ///     });
    /// 
    ///     // enable cloud backup snapshots
    ///     var testCloudProviderSnapshot = new Mongodbatlas.CloudProviderSnapshot("testCloudProviderSnapshot", new()
    ///     {
    ///         ProjectId = myCluster.ProjectId,
    ///         ClusterName = myCluster.Name,
    ///         Description = "myDescription",
    ///         RetentionInDays = 1,
    ///     });
    /// 
    ///     var testCloudProviderSnapshotRestoreJob = new Mongodbatlas.CloudProviderSnapshotRestoreJob("testCloudProviderSnapshotRestoreJob", new()
    ///     {
    ///         ProjectId = testCloudProviderSnapshot.ProjectId,
    ///         ClusterName = testCloudProviderSnapshot.ClusterName,
    ///         SnapshotId = testCloudProviderSnapshot.SnapshotId,
    ///         DeliveryTypeConfig = new Mongodbatlas.Inputs.CloudProviderSnapshotRestoreJobDeliveryTypeConfigArgs
    ///         {
    ///             Automated = true,
    ///             TargetClusterName = "MyCluster",
    ///             TargetProjectId = "5cf5a45a9ccf6400e60981b6",
    ///         },
    ///     }, new CustomResourceOptions
    ///     {
    ///         DependsOn = new[]
    ///         {
    ///             testCloudProviderSnapshot,
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// ### Example download delivery type.
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using Mongodbatlas = Pulumi.Mongodbatlas;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var myCluster = new Mongodbatlas.Cluster("myCluster", new()
    ///     {
    ///         ProjectId = "5cf5a45a9ccf6400e60981b6",
    ///         DiskSizeGb = 5,
    ///         ProviderName = "AWS",
    ///         ProviderRegionName = "EU_WEST_2",
    ///         ProviderInstanceSizeName = "M10",
    ///         CloudBackup = true,
    ///     });
    /// 
    ///     // enable cloud backup snapshots
    ///     var testCloudProviderSnapshot = new Mongodbatlas.CloudProviderSnapshot("testCloudProviderSnapshot", new()
    ///     {
    ///         ProjectId = myCluster.ProjectId,
    ///         ClusterName = myCluster.Name,
    ///         Description = "myDescription",
    ///         RetentionInDays = 1,
    ///     });
    /// 
    ///     var testCloudProviderSnapshotRestoreJob = new Mongodbatlas.CloudProviderSnapshotRestoreJob("testCloudProviderSnapshotRestoreJob", new()
    ///     {
    ///         ProjectId = testCloudProviderSnapshot.ProjectId,
    ///         ClusterName = testCloudProviderSnapshot.ClusterName,
    ///         SnapshotId = testCloudProviderSnapshot.SnapshotId,
    ///         DeliveryTypeConfig = new Mongodbatlas.Inputs.CloudProviderSnapshotRestoreJobDeliveryTypeConfigArgs
    ///         {
    ///             Download = true,
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Cloud Backup Snapshot Restore Job entries can be imported using project project_id, cluster_name and snapshot_id (Unique identifier of the snapshot), in the format `PROJECTID-CLUSTERNAME-JOBID`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import mongodbatlas:index/cloudProviderSnapshotRestoreJob:CloudProviderSnapshotRestoreJob test 5cf5a45a9ccf6400e60981b6-MyCluster-5d1b654ecf09a24b888f4c79
    /// ```
    /// 
    ///  For more information see[MongoDB Atlas API Reference.](https://docs.atlas.mongodb.com/reference/api/cloud-backup/restore/restores/)
    /// </summary>
    [MongodbatlasResourceType("mongodbatlas:index/cloudProviderSnapshotRestoreJob:CloudProviderSnapshotRestoreJob")]
    public partial class CloudProviderSnapshotRestoreJob : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Indicates whether the restore job was canceled.
        /// </summary>
        [Output("cancelled")]
        public Output<bool> Cancelled { get; private set; } = null!;

        /// <summary>
        /// The name of the Atlas cluster whose snapshot you want to restore.
        /// </summary>
        [Output("clusterName")]
        public Output<string> ClusterName { get; private set; } = null!;

        /// <summary>
        /// UTC ISO 8601 formatted point in time when Atlas created the restore job.
        /// </summary>
        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        /// <summary>
        /// Type of restore job to create. Possible values are: **download** or **automated**, only one must be set it in ``true``.
        /// </summary>
        [Output("deliveryType")]
        public Output<ImmutableDictionary<string, string>?> DeliveryType { get; private set; } = null!;

        /// <summary>
        /// Type of restore job to create. Possible values are: automated and download.
        /// </summary>
        [Output("deliveryTypeConfig")]
        public Output<Outputs.CloudProviderSnapshotRestoreJobDeliveryTypeConfig?> DeliveryTypeConfig { get; private set; } = null!;

        /// <summary>
        /// One or more URLs for the compressed snapshot files for manual download. Only visible if deliveryType is download.
        /// </summary>
        [Output("deliveryUrls")]
        public Output<ImmutableArray<string>> DeliveryUrls { get; private set; } = null!;

        /// <summary>
        /// Indicates whether the restore job expired.
        /// </summary>
        [Output("expired")]
        public Output<bool> Expired { get; private set; } = null!;

        /// <summary>
        /// UTC ISO 8601 formatted point in time when the restore job expires.
        /// </summary>
        [Output("expiresAt")]
        public Output<string> ExpiresAt { get; private set; } = null!;

        /// <summary>
        /// UTC ISO 8601 formatted point in time when the restore job completed.
        /// </summary>
        [Output("finishedAt")]
        public Output<string> FinishedAt { get; private set; } = null!;

        /// <summary>
        /// The unique identifier of the project for the Atlas cluster whose snapshot you want to restore.
        /// </summary>
        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;

        /// <summary>
        /// Unique identifier of the snapshot to restore.
        /// </summary>
        [Output("snapshotId")]
        public Output<string> SnapshotId { get; private set; } = null!;

        /// <summary>
        /// The unique identifier of the restore job.
        /// </summary>
        [Output("snapshotRestoreJobId")]
        public Output<string> SnapshotRestoreJobId { get; private set; } = null!;

        /// <summary>
        /// Timestamp in ISO 8601 date and time format in UTC when the snapshot associated to snapshotId was taken.
        /// </summary>
        [Output("timestamp")]
        public Output<string> Timestamp { get; private set; } = null!;


        /// <summary>
        /// Create a CloudProviderSnapshotRestoreJob resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public CloudProviderSnapshotRestoreJob(string name, CloudProviderSnapshotRestoreJobArgs args, CustomResourceOptions? options = null)
            : base("mongodbatlas:index/cloudProviderSnapshotRestoreJob:CloudProviderSnapshotRestoreJob", name, args ?? new CloudProviderSnapshotRestoreJobArgs(), MakeResourceOptions(options, ""))
        {
        }

        private CloudProviderSnapshotRestoreJob(string name, Input<string> id, CloudProviderSnapshotRestoreJobState? state = null, CustomResourceOptions? options = null)
            : base("mongodbatlas:index/cloudProviderSnapshotRestoreJob:CloudProviderSnapshotRestoreJob", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing CloudProviderSnapshotRestoreJob resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static CloudProviderSnapshotRestoreJob Get(string name, Input<string> id, CloudProviderSnapshotRestoreJobState? state = null, CustomResourceOptions? options = null)
        {
            return new CloudProviderSnapshotRestoreJob(name, id, state, options);
        }
    }

    public sealed class CloudProviderSnapshotRestoreJobArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the Atlas cluster whose snapshot you want to restore.
        /// </summary>
        [Input("clusterName", required: true)]
        public Input<string> ClusterName { get; set; } = null!;

        [Input("deliveryType")]
        private InputMap<string>? _deliveryType;

        /// <summary>
        /// Type of restore job to create. Possible values are: **download** or **automated**, only one must be set it in ``true``.
        /// </summary>
        [Obsolete(@"use delivery_type_config instead")]
        public InputMap<string> DeliveryType
        {
            get => _deliveryType ?? (_deliveryType = new InputMap<string>());
            set => _deliveryType = value;
        }

        /// <summary>
        /// Type of restore job to create. Possible values are: automated and download.
        /// </summary>
        [Input("deliveryTypeConfig")]
        public Input<Inputs.CloudProviderSnapshotRestoreJobDeliveryTypeConfigArgs>? DeliveryTypeConfig { get; set; }

        /// <summary>
        /// The unique identifier of the project for the Atlas cluster whose snapshot you want to restore.
        /// </summary>
        [Input("projectId", required: true)]
        public Input<string> ProjectId { get; set; } = null!;

        /// <summary>
        /// Unique identifier of the snapshot to restore.
        /// </summary>
        [Input("snapshotId", required: true)]
        public Input<string> SnapshotId { get; set; } = null!;

        public CloudProviderSnapshotRestoreJobArgs()
        {
        }
        public static new CloudProviderSnapshotRestoreJobArgs Empty => new CloudProviderSnapshotRestoreJobArgs();
    }

    public sealed class CloudProviderSnapshotRestoreJobState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Indicates whether the restore job was canceled.
        /// </summary>
        [Input("cancelled")]
        public Input<bool>? Cancelled { get; set; }

        /// <summary>
        /// The name of the Atlas cluster whose snapshot you want to restore.
        /// </summary>
        [Input("clusterName")]
        public Input<string>? ClusterName { get; set; }

        /// <summary>
        /// UTC ISO 8601 formatted point in time when Atlas created the restore job.
        /// </summary>
        [Input("createdAt")]
        public Input<string>? CreatedAt { get; set; }

        [Input("deliveryType")]
        private InputMap<string>? _deliveryType;

        /// <summary>
        /// Type of restore job to create. Possible values are: **download** or **automated**, only one must be set it in ``true``.
        /// </summary>
        [Obsolete(@"use delivery_type_config instead")]
        public InputMap<string> DeliveryType
        {
            get => _deliveryType ?? (_deliveryType = new InputMap<string>());
            set => _deliveryType = value;
        }

        /// <summary>
        /// Type of restore job to create. Possible values are: automated and download.
        /// </summary>
        [Input("deliveryTypeConfig")]
        public Input<Inputs.CloudProviderSnapshotRestoreJobDeliveryTypeConfigGetArgs>? DeliveryTypeConfig { get; set; }

        [Input("deliveryUrls")]
        private InputList<string>? _deliveryUrls;

        /// <summary>
        /// One or more URLs for the compressed snapshot files for manual download. Only visible if deliveryType is download.
        /// </summary>
        public InputList<string> DeliveryUrls
        {
            get => _deliveryUrls ?? (_deliveryUrls = new InputList<string>());
            set => _deliveryUrls = value;
        }

        /// <summary>
        /// Indicates whether the restore job expired.
        /// </summary>
        [Input("expired")]
        public Input<bool>? Expired { get; set; }

        /// <summary>
        /// UTC ISO 8601 formatted point in time when the restore job expires.
        /// </summary>
        [Input("expiresAt")]
        public Input<string>? ExpiresAt { get; set; }

        /// <summary>
        /// UTC ISO 8601 formatted point in time when the restore job completed.
        /// </summary>
        [Input("finishedAt")]
        public Input<string>? FinishedAt { get; set; }

        /// <summary>
        /// The unique identifier of the project for the Atlas cluster whose snapshot you want to restore.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// Unique identifier of the snapshot to restore.
        /// </summary>
        [Input("snapshotId")]
        public Input<string>? SnapshotId { get; set; }

        /// <summary>
        /// The unique identifier of the restore job.
        /// </summary>
        [Input("snapshotRestoreJobId")]
        public Input<string>? SnapshotRestoreJobId { get; set; }

        /// <summary>
        /// Timestamp in ISO 8601 date and time format in UTC when the snapshot associated to snapshotId was taken.
        /// </summary>
        [Input("timestamp")]
        public Input<string>? Timestamp { get; set; }

        public CloudProviderSnapshotRestoreJobState()
        {
        }
        public static new CloudProviderSnapshotRestoreJobState Empty => new CloudProviderSnapshotRestoreJobState();
    }
}
