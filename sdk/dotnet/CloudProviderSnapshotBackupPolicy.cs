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
    /// ## Import
    /// 
    /// Cloud Backup Snapshot Policy entries can be imported using project project_id and cluster_name, in the format `PROJECTID-CLUSTERNAME`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import mongodbatlas:index/cloudProviderSnapshotBackupPolicy:CloudProviderSnapshotBackupPolicy test 5d0f1f73cf09a29120e173cf-MyClusterTest
    /// ```
    /// 
    ///  For more information see[MongoDB Atlas API Reference.](https://docs.atlas.mongodb.com/reference/api/cloud-backup/schedule/modify-one-schedule/)
    /// </summary>
    [MongodbatlasResourceType("mongodbatlas:index/cloudProviderSnapshotBackupPolicy:CloudProviderSnapshotBackupPolicy")]
    public partial class CloudProviderSnapshotBackupPolicy : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Unique identifier of the Atlas cluster.
        /// </summary>
        [Output("clusterId")]
        public Output<string> ClusterId { get; private set; } = null!;

        /// <summary>
        /// The name of the Atlas cluster that contains the snapshot backup policy you want to retrieve.
        /// </summary>
        [Output("clusterName")]
        public Output<string> ClusterName { get; private set; } = null!;

        /// <summary>
        /// Timestamp in the number of seconds that have elapsed since the UNIX epoch when Atlas takes the next snapshot.
        /// </summary>
        [Output("nextSnapshot")]
        public Output<string> NextSnapshot { get; private set; } = null!;

        /// <summary>
        /// Contains a document for each backup policy item in the desired updated backup policy.
        /// * `policies.#.id` - (Required) Unique identifier of the backup policy that you want to update. policies.#.id is a value obtained via the mongodbatlas.Cluster resource. provider_backup_enabled of the mongodbatlas.Cluster resource must be set to true. See the example above for how to refer to the mongodbatlas.Cluster resource for policies.#.id
        /// </summary>
        [Output("policies")]
        public Output<ImmutableArray<Outputs.CloudProviderSnapshotBackupPolicyPolicy>> Policies { get; private set; } = null!;

        /// <summary>
        /// The unique identifier of the project for the Atlas cluster.
        /// </summary>
        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;

        /// <summary>
        /// UTC Hour of day between 0 and 23, inclusive, representing which hour of the day that Atlas takes snapshots for backup policy items.
        /// </summary>
        [Output("referenceHourOfDay")]
        public Output<int> ReferenceHourOfDay { get; private set; } = null!;

        /// <summary>
        /// UTC Minutes after referenceHourOfDay that Atlas takes snapshots for backup policy items. Must be between 0 and 59, inclusive.
        /// </summary>
        [Output("referenceMinuteOfHour")]
        public Output<int> ReferenceMinuteOfHour { get; private set; } = null!;

        /// <summary>
        /// Number of days back in time you can restore to with point-in-time accuracy. Must be a positive, non-zero integer.
        /// </summary>
        [Output("restoreWindowDays")]
        public Output<int> RestoreWindowDays { get; private set; } = null!;

        /// <summary>
        /// Specify true to apply the retention changes in the updated backup policy to snapshots that Atlas took previously.
        /// </summary>
        [Output("updateSnapshots")]
        public Output<bool> UpdateSnapshots { get; private set; } = null!;


        /// <summary>
        /// Create a CloudProviderSnapshotBackupPolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public CloudProviderSnapshotBackupPolicy(string name, CloudProviderSnapshotBackupPolicyArgs args, CustomResourceOptions? options = null)
            : base("mongodbatlas:index/cloudProviderSnapshotBackupPolicy:CloudProviderSnapshotBackupPolicy", name, args ?? new CloudProviderSnapshotBackupPolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private CloudProviderSnapshotBackupPolicy(string name, Input<string> id, CloudProviderSnapshotBackupPolicyState? state = null, CustomResourceOptions? options = null)
            : base("mongodbatlas:index/cloudProviderSnapshotBackupPolicy:CloudProviderSnapshotBackupPolicy", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing CloudProviderSnapshotBackupPolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static CloudProviderSnapshotBackupPolicy Get(string name, Input<string> id, CloudProviderSnapshotBackupPolicyState? state = null, CustomResourceOptions? options = null)
        {
            return new CloudProviderSnapshotBackupPolicy(name, id, state, options);
        }
    }

    public sealed class CloudProviderSnapshotBackupPolicyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the Atlas cluster that contains the snapshot backup policy you want to retrieve.
        /// </summary>
        [Input("clusterName", required: true)]
        public Input<string> ClusterName { get; set; } = null!;

        [Input("policies", required: true)]
        private InputList<Inputs.CloudProviderSnapshotBackupPolicyPolicyArgs>? _policies;

        /// <summary>
        /// Contains a document for each backup policy item in the desired updated backup policy.
        /// * `policies.#.id` - (Required) Unique identifier of the backup policy that you want to update. policies.#.id is a value obtained via the mongodbatlas.Cluster resource. provider_backup_enabled of the mongodbatlas.Cluster resource must be set to true. See the example above for how to refer to the mongodbatlas.Cluster resource for policies.#.id
        /// </summary>
        public InputList<Inputs.CloudProviderSnapshotBackupPolicyPolicyArgs> Policies
        {
            get => _policies ?? (_policies = new InputList<Inputs.CloudProviderSnapshotBackupPolicyPolicyArgs>());
            set => _policies = value;
        }

        /// <summary>
        /// The unique identifier of the project for the Atlas cluster.
        /// </summary>
        [Input("projectId", required: true)]
        public Input<string> ProjectId { get; set; } = null!;

        /// <summary>
        /// UTC Hour of day between 0 and 23, inclusive, representing which hour of the day that Atlas takes snapshots for backup policy items.
        /// </summary>
        [Input("referenceHourOfDay")]
        public Input<int>? ReferenceHourOfDay { get; set; }

        /// <summary>
        /// UTC Minutes after referenceHourOfDay that Atlas takes snapshots for backup policy items. Must be between 0 and 59, inclusive.
        /// </summary>
        [Input("referenceMinuteOfHour")]
        public Input<int>? ReferenceMinuteOfHour { get; set; }

        /// <summary>
        /// Number of days back in time you can restore to with point-in-time accuracy. Must be a positive, non-zero integer.
        /// </summary>
        [Input("restoreWindowDays")]
        public Input<int>? RestoreWindowDays { get; set; }

        /// <summary>
        /// Specify true to apply the retention changes in the updated backup policy to snapshots that Atlas took previously.
        /// </summary>
        [Input("updateSnapshots")]
        public Input<bool>? UpdateSnapshots { get; set; }

        public CloudProviderSnapshotBackupPolicyArgs()
        {
        }
        public static new CloudProviderSnapshotBackupPolicyArgs Empty => new CloudProviderSnapshotBackupPolicyArgs();
    }

    public sealed class CloudProviderSnapshotBackupPolicyState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Unique identifier of the Atlas cluster.
        /// </summary>
        [Input("clusterId")]
        public Input<string>? ClusterId { get; set; }

        /// <summary>
        /// The name of the Atlas cluster that contains the snapshot backup policy you want to retrieve.
        /// </summary>
        [Input("clusterName")]
        public Input<string>? ClusterName { get; set; }

        /// <summary>
        /// Timestamp in the number of seconds that have elapsed since the UNIX epoch when Atlas takes the next snapshot.
        /// </summary>
        [Input("nextSnapshot")]
        public Input<string>? NextSnapshot { get; set; }

        [Input("policies")]
        private InputList<Inputs.CloudProviderSnapshotBackupPolicyPolicyGetArgs>? _policies;

        /// <summary>
        /// Contains a document for each backup policy item in the desired updated backup policy.
        /// * `policies.#.id` - (Required) Unique identifier of the backup policy that you want to update. policies.#.id is a value obtained via the mongodbatlas.Cluster resource. provider_backup_enabled of the mongodbatlas.Cluster resource must be set to true. See the example above for how to refer to the mongodbatlas.Cluster resource for policies.#.id
        /// </summary>
        public InputList<Inputs.CloudProviderSnapshotBackupPolicyPolicyGetArgs> Policies
        {
            get => _policies ?? (_policies = new InputList<Inputs.CloudProviderSnapshotBackupPolicyPolicyGetArgs>());
            set => _policies = value;
        }

        /// <summary>
        /// The unique identifier of the project for the Atlas cluster.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// UTC Hour of day between 0 and 23, inclusive, representing which hour of the day that Atlas takes snapshots for backup policy items.
        /// </summary>
        [Input("referenceHourOfDay")]
        public Input<int>? ReferenceHourOfDay { get; set; }

        /// <summary>
        /// UTC Minutes after referenceHourOfDay that Atlas takes snapshots for backup policy items. Must be between 0 and 59, inclusive.
        /// </summary>
        [Input("referenceMinuteOfHour")]
        public Input<int>? ReferenceMinuteOfHour { get; set; }

        /// <summary>
        /// Number of days back in time you can restore to with point-in-time accuracy. Must be a positive, non-zero integer.
        /// </summary>
        [Input("restoreWindowDays")]
        public Input<int>? RestoreWindowDays { get; set; }

        /// <summary>
        /// Specify true to apply the retention changes in the updated backup policy to snapshots that Atlas took previously.
        /// </summary>
        [Input("updateSnapshots")]
        public Input<bool>? UpdateSnapshots { get; set; }

        public CloudProviderSnapshotBackupPolicyState()
        {
        }
        public static new CloudProviderSnapshotBackupPolicyState Empty => new CloudProviderSnapshotBackupPolicyState();
    }
}
