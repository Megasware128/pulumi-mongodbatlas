// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * `mongodbatlas.getThirdPartyIntegrations` describe all Third-Party Integration Settings. This represents two Third-Party services `PAGER_DUTY` and `FLOWDOCK`
 * applied across the project.
 *
 * > **NOTE:** Groups and projects are synonymous terms. You may find `groupId` in the official documentation.
 */
export function getThirdPartyIntegrations(args: GetThirdPartyIntegrationsArgs, opts?: pulumi.InvokeOptions): Promise<GetThirdPartyIntegrationsResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("mongodbatlas:index/getThirdPartyIntegrations:getThirdPartyIntegrations", {
        "projectId": args.projectId,
    }, opts);
}

/**
 * A collection of arguments for invoking getThirdPartyIntegrations.
 */
export interface GetThirdPartyIntegrationsArgs {
    /**
     * The unique ID for the project to get all Third-Party service integrations
     */
    projectId: string;
}

/**
 * A collection of values returned by getThirdPartyIntegrations.
 */
export interface GetThirdPartyIntegrationsResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * (Required) ID of the Atlas project the Third-Party Service Integration belongs to.
     */
    readonly projectId: string;
    /**
     * A list where each represents a Third-Party service integration.
     */
    readonly results: outputs.GetThirdPartyIntegrationsResult[];
}

export function getThirdPartyIntegrationsOutput(args: GetThirdPartyIntegrationsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetThirdPartyIntegrationsResult> {
    return pulumi.output(args).apply(a => getThirdPartyIntegrations(a, opts))
}

/**
 * A collection of arguments for invoking getThirdPartyIntegrations.
 */
export interface GetThirdPartyIntegrationsOutputArgs {
    /**
     * The unique ID for the project to get all Third-Party service integrations
     */
    projectId: pulumi.Input<string>;
}
