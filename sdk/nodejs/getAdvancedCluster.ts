// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * ## --
 *
 * layout: "mongodbatlas"
 * page_title: "MongoDB Atlas: advancedCluster"
 * sidebar_current: "docs-mongodbatlas-datasource-advanced-cluster"
 * description: |-
 *     Describe an Advanced Cluster.
 * ---
 *
 * # Data Source: mongodbatlas.AdvancedCluster
 *
 * `mongodbatlas.AdvancedCluster` describes an Advanced Cluster. The data source requires your Project ID.
 *
 * > **NOTE:** Groups and projects are synonymous terms. You may find groupId in the official documentation.
 *
 * > **IMPORTANT:**
 * <br> &#8226; Changes to cluster configurations can affect costs. Before making changes, please see [Billing](https://docs.atlas.mongodb.com/billing/).
 * <br> &#8226; If your Atlas project contains a custom role that uses actions introduced in a specific MongoDB version, you cannot create a cluster with a MongoDB version less than that version unless you delete the custom role.
 */
export function getAdvancedCluster(args: GetAdvancedClusterArgs, opts?: pulumi.InvokeOptions): Promise<GetAdvancedClusterResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("mongodbatlas:index/getAdvancedCluster:getAdvancedCluster", {
        "name": args.name,
        "pitEnabled": args.pitEnabled,
        "projectId": args.projectId,
    }, opts);
}

/**
 * A collection of arguments for invoking getAdvancedCluster.
 */
export interface GetAdvancedClusterArgs {
    /**
     * Name of the cluster as it appears in Atlas. Once the cluster is created, its name cannot be changed.
     */
    name: string;
    /**
     * Flag that indicates if the cluster uses Continuous Cloud Backup.
     */
    pitEnabled?: boolean;
    /**
     * The unique ID for the project to create the database user.
     */
    projectId: string;
}

/**
 * A collection of values returned by getAdvancedCluster.
 */
export interface GetAdvancedClusterResult {
    /**
     * Get the advanced configuration options. See Advanced Configuration below for more details.
     */
    readonly advancedConfigurations: outputs.GetAdvancedClusterAdvancedConfiguration[];
    readonly backupEnabled: boolean;
    /**
     * Configuration settings applied to BI Connector for Atlas on this cluster. See below.
     */
    readonly biConnectors: outputs.GetAdvancedClusterBiConnector[];
    /**
     * Type of the cluster that you want to create.
     */
    readonly clusterType: string;
    /**
     * Set of connection strings that your applications use to connect to this cluster. More info in [Connection-strings](https://docs.mongodb.com/manual/reference/connection-string/). Use the parameters in this object to connect your applications to this cluster. To learn more about the formats of connection strings, see [Connection String Options](https://docs.atlas.mongodb.com/reference/faq/connection-changes/). NOTE: Atlas returns the contents of this object after the cluster is operational, not while it builds the cluster.
     */
    readonly connectionStrings: outputs.GetAdvancedClusterConnectionString[];
    readonly createDate: string;
    /**
     * Capacity, in gigabytes, of the host's root volume.
     */
    readonly diskSizeGb: number;
    /**
     * Possible values are AWS, GCP, AZURE or NONE.
     */
    readonly encryptionAtRestProvider: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Configuration for the collection of key-value pairs that tag and categorize the cluster. See below.
     */
    readonly labels: outputs.GetAdvancedClusterLabel[];
    /**
     * Version of the cluster to deploy.
     */
    readonly mongoDbMajorVersion: string;
    /**
     * Version of MongoDB the cluster runs, in `major-version`.`minor-version` format.
     */
    readonly mongoDbVersion: string;
    readonly name: string;
    /**
     * Flag that indicates whether the cluster is paused or not.
     */
    readonly paused: boolean;
    /**
     * Flag that indicates if the cluster uses Continuous Cloud Backup.
     */
    readonly pitEnabled: boolean;
    readonly projectId: string;
    /**
     * Configuration for cluster regions and the hardware provisioned in them. See below
     */
    readonly replicationSpecs: outputs.GetAdvancedClusterReplicationSpec[];
    /**
     * Certificate Authority that MongoDB Atlas clusters use.
     */
    readonly rootCertType: string;
    /**
     * Current state of the cluster. The possible states are:
     */
    readonly stateName: string;
    /**
     * Release cadence that Atlas uses for this cluster.
     */
    readonly versionReleaseSystem: string;
}

export function getAdvancedClusterOutput(args: GetAdvancedClusterOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetAdvancedClusterResult> {
    return pulumi.output(args).apply(a => getAdvancedCluster(a, opts))
}

/**
 * A collection of arguments for invoking getAdvancedCluster.
 */
export interface GetAdvancedClusterOutputArgs {
    /**
     * Name of the cluster as it appears in Atlas. Once the cluster is created, its name cannot be changed.
     */
    name: pulumi.Input<string>;
    /**
     * Flag that indicates if the cluster uses Continuous Cloud Backup.
     */
    pitEnabled?: pulumi.Input<boolean>;
    /**
     * The unique ID for the project to create the database user.
     */
    projectId: pulumi.Input<string>;
}
