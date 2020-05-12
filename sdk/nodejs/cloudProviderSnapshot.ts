// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * `mongodbatlas..CloudProviderSnapshot` provides a resource to take a cloud provider snapshot on demand.
 * On-demand snapshots happen immediately, unlike scheduled snapshots which occur at regular intervals. If there is already an on-demand snapshot with a status of queued or inProgress, you must wait until Atlas has completed the on-demand snapshot before taking another.
 * 
 * > **NOTE:** Groups and projects are synonymous terms. You may find `groupId` in the official documentation.
 * 
 * ## Example Usage
 * 
 * 
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as mongodbatlas from "@pulumi/mongodbatlas";
 * 
 * const myCluster = new mongodbatlas.Cluster("myCluster", {
 *     projectId: "5cf5a45a9ccf6400e60981b6",
 *     diskSizeGb: 5,
 *     providerName: "AWS",
 *     providerRegionName: "EU_WEST_2",
 *     providerInstanceSizeName: "M10",
 *     providerBackupEnabled: true,
 *     providerDiskIops: 100,
 *     providerEncryptEbsVolume: false,
 * });
 * const testCloudProviderSnapshot = new mongodbatlas.CloudProviderSnapshot("testCloudProviderSnapshot", {
 *     projectId: myCluster.projectId,
 *     clusterName: myCluster.name,
 *     description: "myDescription",
 *     retentionInDays: 1,
 * });
 * const testCloudProviderSnapshotRestoreJob = new mongodbatlas.CloudProviderSnapshotRestoreJob("testCloudProviderSnapshotRestoreJob", {
 *     projectId: testCloudProviderSnapshot.projectId,
 *     clusterName: testCloudProviderSnapshot.clusterName,
 *     snapshotId: testCloudProviderSnapshot.snapshotId,
 *     deliveryType: {
 *         download: true,
 *     },
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-mongodbatlas/blob/master/website/docs/r/cloud_provider_snapshot.html.markdown.
 */
export class CloudProviderSnapshot extends pulumi.CustomResource {
    /**
     * Get an existing CloudProviderSnapshot resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: CloudProviderSnapshotState, opts?: pulumi.CustomResourceOptions): CloudProviderSnapshot {
        return new CloudProviderSnapshot(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'mongodbatlas:index/cloudProviderSnapshot:CloudProviderSnapshot';

    /**
     * Returns true if the given object is an instance of CloudProviderSnapshot.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is CloudProviderSnapshot {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === CloudProviderSnapshot.__pulumiType;
    }

    /**
     * The name of the Atlas cluster that contains the snapshots you want to retrieve.
     */
    public readonly clusterName!: pulumi.Output<string>;
    /**
     * UTC ISO 8601 formatted point in time when Atlas took the snapshot.
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * Description of the on-demand snapshot.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * UTC ISO 8601 formatted point in time when Atlas will delete the snapshot.
     */
    public /*out*/ readonly expiresAt!: pulumi.Output<string>;
    /**
     * Unique ID of the AWS KMS Customer Master Key used to encrypt the snapshot. Only visible for clusters using Encryption at Rest via Customer KMS.
     */
    public /*out*/ readonly masterKeyUuid!: pulumi.Output<string>;
    /**
     * Version of the MongoDB server.
     */
    public /*out*/ readonly mongodVersion!: pulumi.Output<string>;
    /**
     * The unique identifier of the project for the Atlas cluster.
     */
    public readonly projectId!: pulumi.Output<string>;
    /**
     * The number of days that Atlas should retain the on-demand snapshot. Must be at least 1.
     */
    public readonly retentionInDays!: pulumi.Output<number>;
    /**
     * Unique identifier of the snapshot.
     */
    public /*out*/ readonly snapshotId!: pulumi.Output<string>;
    /**
     * Specified the type of snapshot. Valid values are onDemand and scheduled.
     */
    public /*out*/ readonly snapshotType!: pulumi.Output<string>;
    /**
     * Current status of the snapshot. One of the following values will be returned: queued, inProgress, completed, failed.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * Specifies the size of the snapshot in bytes.
     */
    public /*out*/ readonly storageSizeBytes!: pulumi.Output<number>;
    /**
     * Specifies the type of cluster: replicaSet or shardedCluster.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a CloudProviderSnapshot resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CloudProviderSnapshotArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: CloudProviderSnapshotArgs | CloudProviderSnapshotState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as CloudProviderSnapshotState | undefined;
            inputs["clusterName"] = state ? state.clusterName : undefined;
            inputs["createdAt"] = state ? state.createdAt : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["expiresAt"] = state ? state.expiresAt : undefined;
            inputs["masterKeyUuid"] = state ? state.masterKeyUuid : undefined;
            inputs["mongodVersion"] = state ? state.mongodVersion : undefined;
            inputs["projectId"] = state ? state.projectId : undefined;
            inputs["retentionInDays"] = state ? state.retentionInDays : undefined;
            inputs["snapshotId"] = state ? state.snapshotId : undefined;
            inputs["snapshotType"] = state ? state.snapshotType : undefined;
            inputs["status"] = state ? state.status : undefined;
            inputs["storageSizeBytes"] = state ? state.storageSizeBytes : undefined;
            inputs["type"] = state ? state.type : undefined;
        } else {
            const args = argsOrState as CloudProviderSnapshotArgs | undefined;
            if (!args || args.clusterName === undefined) {
                throw new Error("Missing required property 'clusterName'");
            }
            if (!args || args.description === undefined) {
                throw new Error("Missing required property 'description'");
            }
            if (!args || args.projectId === undefined) {
                throw new Error("Missing required property 'projectId'");
            }
            if (!args || args.retentionInDays === undefined) {
                throw new Error("Missing required property 'retentionInDays'");
            }
            inputs["clusterName"] = args ? args.clusterName : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["projectId"] = args ? args.projectId : undefined;
            inputs["retentionInDays"] = args ? args.retentionInDays : undefined;
            inputs["createdAt"] = undefined /*out*/;
            inputs["expiresAt"] = undefined /*out*/;
            inputs["masterKeyUuid"] = undefined /*out*/;
            inputs["mongodVersion"] = undefined /*out*/;
            inputs["snapshotId"] = undefined /*out*/;
            inputs["snapshotType"] = undefined /*out*/;
            inputs["status"] = undefined /*out*/;
            inputs["storageSizeBytes"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(CloudProviderSnapshot.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering CloudProviderSnapshot resources.
 */
export interface CloudProviderSnapshotState {
    /**
     * The name of the Atlas cluster that contains the snapshots you want to retrieve.
     */
    readonly clusterName?: pulumi.Input<string>;
    /**
     * UTC ISO 8601 formatted point in time when Atlas took the snapshot.
     */
    readonly createdAt?: pulumi.Input<string>;
    /**
     * Description of the on-demand snapshot.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * UTC ISO 8601 formatted point in time when Atlas will delete the snapshot.
     */
    readonly expiresAt?: pulumi.Input<string>;
    /**
     * Unique ID of the AWS KMS Customer Master Key used to encrypt the snapshot. Only visible for clusters using Encryption at Rest via Customer KMS.
     */
    readonly masterKeyUuid?: pulumi.Input<string>;
    /**
     * Version of the MongoDB server.
     */
    readonly mongodVersion?: pulumi.Input<string>;
    /**
     * The unique identifier of the project for the Atlas cluster.
     */
    readonly projectId?: pulumi.Input<string>;
    /**
     * The number of days that Atlas should retain the on-demand snapshot. Must be at least 1.
     */
    readonly retentionInDays?: pulumi.Input<number>;
    /**
     * Unique identifier of the snapshot.
     */
    readonly snapshotId?: pulumi.Input<string>;
    /**
     * Specified the type of snapshot. Valid values are onDemand and scheduled.
     */
    readonly snapshotType?: pulumi.Input<string>;
    /**
     * Current status of the snapshot. One of the following values will be returned: queued, inProgress, completed, failed.
     */
    readonly status?: pulumi.Input<string>;
    /**
     * Specifies the size of the snapshot in bytes.
     */
    readonly storageSizeBytes?: pulumi.Input<number>;
    /**
     * Specifies the type of cluster: replicaSet or shardedCluster.
     */
    readonly type?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a CloudProviderSnapshot resource.
 */
export interface CloudProviderSnapshotArgs {
    /**
     * The name of the Atlas cluster that contains the snapshots you want to retrieve.
     */
    readonly clusterName: pulumi.Input<string>;
    /**
     * Description of the on-demand snapshot.
     */
    readonly description: pulumi.Input<string>;
    /**
     * The unique identifier of the project for the Atlas cluster.
     */
    readonly projectId: pulumi.Input<string>;
    /**
     * The number of days that Atlas should retain the on-demand snapshot. Must be at least 1.
     */
    readonly retentionInDays: pulumi.Input<number>;
}
