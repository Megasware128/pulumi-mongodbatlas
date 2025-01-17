// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * `mongodbatlas.CloudProviderAccess` allows you to get a single role for a provider access role setup, currently only AWS is supported.
 *
 * > **NOTE:** Groups and projects are synonymous terms. You may find `groupId` in the official documentation.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as mongodbatlas from "@pulumi/mongodbatlas";
 *
 * const testRole = new mongodbatlas.CloudProviderAccessSetup("testRole", {
 *     projectId: "<PROJECT-ID>",
 *     providerName: "AWS",
 * });
 * const singleSetup = mongodbatlas.getCloudProviderAccessSetupOutput({
 *     projectId: testRole.projectId,
 *     providerName: testRole.providerName,
 *     roleId: testRole.roleId,
 * });
 * ```
 */
export function getCloudProviderAccessSetup(args: GetCloudProviderAccessSetupArgs, opts?: pulumi.InvokeOptions): Promise<GetCloudProviderAccessSetupResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("mongodbatlas:index/getCloudProviderAccessSetup:getCloudProviderAccessSetup", {
        "projectId": args.projectId,
        "providerName": args.providerName,
        "roleId": args.roleId,
    }, opts);
}

/**
 * A collection of arguments for invoking getCloudProviderAccessSetup.
 */
export interface GetCloudProviderAccessSetupArgs {
    /**
     * The unique ID for the project to get all Cloud Provider Access
     */
    projectId: string;
    /**
     * cloud provider name, currently only AWS is supported
     */
    providerName: string;
    /**
     * unique role id among all the aws roles provided by mongodb atlas
     */
    roleId: string;
}

/**
 * A collection of values returned by getCloudProviderAccessSetup.
 */
export interface GetCloudProviderAccessSetupResult {
    /**
     * aws related role information
     */
    readonly aws: {[key: string]: string};
    readonly awsConfigs: outputs.GetCloudProviderAccessSetupAwsConfig[];
    /**
     * Date on which this role was created.
     */
    readonly createdDate: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly projectId: string;
    readonly providerName: string;
    readonly roleId: string;
}

export function getCloudProviderAccessSetupOutput(args: GetCloudProviderAccessSetupOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetCloudProviderAccessSetupResult> {
    return pulumi.output(args).apply(a => getCloudProviderAccessSetup(a, opts))
}

/**
 * A collection of arguments for invoking getCloudProviderAccessSetup.
 */
export interface GetCloudProviderAccessSetupOutputArgs {
    /**
     * The unique ID for the project to get all Cloud Provider Access
     */
    projectId: pulumi.Input<string>;
    /**
     * cloud provider name, currently only AWS is supported
     */
    providerName: pulumi.Input<string>;
    /**
     * unique role id among all the aws roles provided by mongodb atlas
     */
    roleId: pulumi.Input<string>;
}
