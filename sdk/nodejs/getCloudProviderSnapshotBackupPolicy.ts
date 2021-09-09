// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * **WARNING:** This data source is deprecated, use `mongodbatlas.CloudBackupSchedule`
 *
 * `mongodbatlas.CloudProviderSnapshotBackupPolicy` provides a Cloud Backup Snapshot Backup Policy datasource. An Atlas Cloud Backup Snapshot Policy provides the current snapshot schedule and retention settings for the cluster.
 *
 * > **NOTE:** Groups and projects are synonymous terms. You may find `groupId` in the official documentation.
 */
export function getCloudProviderSnapshotBackupPolicy(args: GetCloudProviderSnapshotBackupPolicyArgs, opts?: pulumi.InvokeOptions): Promise<GetCloudProviderSnapshotBackupPolicyResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("mongodbatlas:index/getCloudProviderSnapshotBackupPolicy:getCloudProviderSnapshotBackupPolicy", {
        "clusterName": args.clusterName,
        "projectId": args.projectId,
    }, opts);
}

/**
 * A collection of arguments for invoking getCloudProviderSnapshotBackupPolicy.
 */
export interface GetCloudProviderSnapshotBackupPolicyArgs {
    /**
     * The name of the Atlas cluster that contains the snapshots backup policy you want to retrieve.
     */
    readonly clusterName: string;
    /**
     * The unique identifier of the project for the Atlas cluster.
     */
    readonly projectId: string;
}

/**
 * A collection of values returned by getCloudProviderSnapshotBackupPolicy.
 */
export interface GetCloudProviderSnapshotBackupPolicyResult {
    /**
     * Unique identifier of the Atlas cluster.
     */
    readonly clusterId: string;
    readonly clusterName: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * UTC ISO 8601 formatted point in time when Atlas will take the next snapshot.
     */
    readonly nextSnapshot: string;
    /**
     * A list of policy definitions for the cluster.
     * * `policies.#.id` - Unique identifier of the backup policy.
     */
    readonly policies: outputs.GetCloudProviderSnapshotBackupPolicyPolicy[];
    readonly projectId: string;
    /**
     * UTC Hour of day between 0 and 23 representing which hour of the day that Atlas takes a snapshot.
     */
    readonly referenceHourOfDay: number;
    /**
     * UTC Minute of day between 0 and 59 representing which minute of the referenceHourOfDay that Atlas takes the snapshot.
     */
    readonly referenceMinuteOfHour: number;
    /**
     * Specifies a restore window in days for cloud backup to maintain.
     */
    readonly restoreWindowDays: number;
    readonly updateSnapshots: boolean;
}
