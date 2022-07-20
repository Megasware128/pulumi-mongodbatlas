// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * `privateEndpointRegionalMode` describe a Private Endpoint Regional Mode. This represents a Private Endpoint Regional Mode Connection that wants to retrieve settings of an Atlas project.
 *
 * > **NOTE:** Groups and projects are synonymous terms. You may find groupId in the official documentation.
 */
export function getPrivateEndpointRegionalMode(args: GetPrivateEndpointRegionalModeArgs, opts?: pulumi.InvokeOptions): Promise<GetPrivateEndpointRegionalModeResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("mongodbatlas:index/getPrivateEndpointRegionalMode:getPrivateEndpointRegionalMode", {
        "enabled": args.enabled,
        "projectId": args.projectId,
    }, opts);
}

/**
 * A collection of arguments for invoking getPrivateEndpointRegionalMode.
 */
export interface GetPrivateEndpointRegionalModeArgs {
    /**
     * Flag that indicates whether the regionalized private endpoitn setting is enabled for the project.
     */
    enabled?: boolean;
    /**
     * Unique identifier for the project.
     */
    projectId: string;
}

/**
 * A collection of values returned by getPrivateEndpointRegionalMode.
 */
export interface GetPrivateEndpointRegionalModeResult {
    readonly enabled?: boolean;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly projectId: string;
}

export function getPrivateEndpointRegionalModeOutput(args: GetPrivateEndpointRegionalModeOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetPrivateEndpointRegionalModeResult> {
    return pulumi.output(args).apply(a => getPrivateEndpointRegionalMode(a, opts))
}

/**
 * A collection of arguments for invoking getPrivateEndpointRegionalMode.
 */
export interface GetPrivateEndpointRegionalModeOutputArgs {
    /**
     * Flag that indicates whether the regionalized private endpoitn setting is enabled for the project.
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * Unique identifier for the project.
     */
    projectId: pulumi.Input<string>;
}
