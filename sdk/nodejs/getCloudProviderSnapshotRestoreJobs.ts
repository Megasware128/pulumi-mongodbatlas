// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * **WARNING:** This datasource is deprecated, use `mongodbatlasCloudBackupSnapshotsRestoreJobs`
 *
 * `mongodbatlas.getCloudProviderSnapshotRestoreJobs` provides a Cloud Backup Snapshot Restore Jobs datasource. Gets all the cloud backup snapshot restore jobs for the specified cluster.
 *
 * > **NOTE:** Groups and projects are synonymous terms. You may find `groupId` in the official documentation.
 *
 * ## Example Usage
 *
 * First create a snapshot of the desired cluster. Then request that snapshot be restored in an automated fashion to the designated cluster and project.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as mongodbatlas from "@pulumi/mongodbatlas";
 *
 * const testCloudProviderSnapshot = new mongodbatlas.CloudProviderSnapshot("testCloudProviderSnapshot", {
 *     projectId: "5cf5a45a9ccf6400e60981b6",
 *     clusterName: "MyCluster",
 *     description: "MyDescription",
 *     retentionInDays: 1,
 * });
 * const testCloudProviderSnapshotRestoreJob = new mongodbatlas.CloudProviderSnapshotRestoreJob("testCloudProviderSnapshotRestoreJob", {
 *     projectId: "5cf5a45a9ccf6400e60981b6",
 *     clusterName: "MyCluster",
 *     snapshotId: testCloudProviderSnapshot.id,
 *     deliveryType: [{
 *         automated: true,
 *         targetClusterName: "MyCluster",
 *         targetProjectId: "5cf5a45a9ccf6400e60981b6",
 *     }],
 * });
 * const testCloudProviderSnapshotRestoreJobs = pulumi.all([testCloudProviderSnapshotRestoreJob.projectId, testCloudProviderSnapshotRestoreJob.clusterName]).apply(([projectId, clusterName]) => mongodbatlas.getCloudProviderSnapshotRestoreJobsOutput({
 *     projectId: projectId,
 *     clusterName: clusterName,
 *     pageNum: 1,
 *     itemsPerPage: 5,
 * }));
 * ```
 */
export function getCloudProviderSnapshotRestoreJobs(args: GetCloudProviderSnapshotRestoreJobsArgs, opts?: pulumi.InvokeOptions): Promise<GetCloudProviderSnapshotRestoreJobsResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("mongodbatlas:index/getCloudProviderSnapshotRestoreJobs:getCloudProviderSnapshotRestoreJobs", {
        "clusterName": args.clusterName,
        "itemsPerPage": args.itemsPerPage,
        "pageNum": args.pageNum,
        "projectId": args.projectId,
    }, opts);
}

/**
 * A collection of arguments for invoking getCloudProviderSnapshotRestoreJobs.
 */
export interface GetCloudProviderSnapshotRestoreJobsArgs {
    /**
     * The name of the Atlas cluster for which you want to retrieve restore jobs.
     */
    clusterName: string;
    /**
     * Number of items to return per page, up to a maximum of 500. Defaults to `100`.
     */
    itemsPerPage?: number;
    /**
     * The page to return. Defaults to `1`.
     */
    pageNum?: number;
    /**
     * The unique identifier of the project for the Atlas cluster.
     */
    projectId: string;
}

/**
 * A collection of values returned by getCloudProviderSnapshotRestoreJobs.
 */
export interface GetCloudProviderSnapshotRestoreJobsResult {
    readonly clusterName: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly itemsPerPage?: number;
    readonly pageNum?: number;
    readonly projectId: string;
    /**
     * Includes cloudProviderSnapshotRestoreJob object for each item detailed in the results array section.
     */
    readonly results: outputs.GetCloudProviderSnapshotRestoreJobsResult[];
    readonly totalCount: number;
}

export function getCloudProviderSnapshotRestoreJobsOutput(args: GetCloudProviderSnapshotRestoreJobsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetCloudProviderSnapshotRestoreJobsResult> {
    return pulumi.output(args).apply(a => getCloudProviderSnapshotRestoreJobs(a, opts))
}

/**
 * A collection of arguments for invoking getCloudProviderSnapshotRestoreJobs.
 */
export interface GetCloudProviderSnapshotRestoreJobsOutputArgs {
    /**
     * The name of the Atlas cluster for which you want to retrieve restore jobs.
     */
    clusterName: pulumi.Input<string>;
    /**
     * Number of items to return per page, up to a maximum of 500. Defaults to `100`.
     */
    itemsPerPage?: pulumi.Input<number>;
    /**
     * The page to return. Defaults to `1`.
     */
    pageNum?: pulumi.Input<number>;
    /**
     * The unique identifier of the project for the Atlas cluster.
     */
    projectId: pulumi.Input<string>;
}
