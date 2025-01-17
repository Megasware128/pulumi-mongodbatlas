// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * `mongodbatlas.getCloudBackupSnapshotExportJobs` datasource allows you to retrieve all the buckets for the specified project.
 *
 * > **NOTE:** Groups and projects are synonymous terms. You may find `groupId` in the official documentation.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as mongodbatlas from "@pulumi/mongodbatlas";
 *
 * const testCloudBackupSnapshotExportBucket = new mongodbatlas.CloudBackupSnapshotExportBucket("testCloudBackupSnapshotExportBucket", {
 *     projectId: "{PROJECT_ID}",
 *     iamRoleId: "{IAM_ROLE_ID}",
 *     bucketName: "example_bucket",
 *     cloudProvider: "AWS",
 * });
 * const testCloudBackupSnapshotExportJob = new mongodbatlas.CloudBackupSnapshotExportJob("testCloudBackupSnapshotExportJob", {
 *     projectId: "{PROJECT_ID}",
 *     clusterName: "{CLUSTER_NAME}",
 *     snapshotId: "{SNAPSHOT_ID}",
 *     exportBucketId: testCloudBackupSnapshotExportBucket.exportBucketId,
 *     customDatas: [{
 *         key: "exported by",
 *         value: "myName",
 *     }],
 * });
 * const testCloudBackupSnapshotExportJobs = mongodbatlas.getCloudBackupSnapshotExportJobs({
 *     projectId: "{PROJECT_ID}",
 *     clusterName: "{CLUSTER_NAME}",
 * });
 * ```
 */
export function getCloudBackupSnapshotExportJobs(args: GetCloudBackupSnapshotExportJobsArgs, opts?: pulumi.InvokeOptions): Promise<GetCloudBackupSnapshotExportJobsResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("mongodbatlas:index/getCloudBackupSnapshotExportJobs:getCloudBackupSnapshotExportJobs", {
        "clusterName": args.clusterName,
        "itemsPerPage": args.itemsPerPage,
        "pageNum": args.pageNum,
        "projectId": args.projectId,
    }, opts);
}

/**
 * A collection of arguments for invoking getCloudBackupSnapshotExportJobs.
 */
export interface GetCloudBackupSnapshotExportJobsArgs {
    /**
     * Name of the Atlas cluster whose export job you want to retrieve.
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
 * A collection of values returned by getCloudBackupSnapshotExportJobs.
 */
export interface GetCloudBackupSnapshotExportJobsResult {
    readonly clusterName: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly itemsPerPage?: number;
    readonly pageNum?: number;
    /**
     * The unique identifier of the project for the Atlas cluster.
     */
    readonly projectId: string;
    /**
     * Includes CloudProviderSnapshotExportJob object for each item detailed in the results array section.
     */
    readonly results: outputs.GetCloudBackupSnapshotExportJobsResult[];
    readonly totalCount: number;
}

export function getCloudBackupSnapshotExportJobsOutput(args: GetCloudBackupSnapshotExportJobsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetCloudBackupSnapshotExportJobsResult> {
    return pulumi.output(args).apply(a => getCloudBackupSnapshotExportJobs(a, opts))
}

/**
 * A collection of arguments for invoking getCloudBackupSnapshotExportJobs.
 */
export interface GetCloudBackupSnapshotExportJobsOutputArgs {
    /**
     * Name of the Atlas cluster whose export job you want to retrieve.
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
