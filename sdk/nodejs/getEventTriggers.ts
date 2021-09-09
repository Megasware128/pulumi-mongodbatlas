// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * `mongodbatlas.getEventTriggers` describe all Event Triggers.
 */
export function getEventTriggers(args: GetEventTriggersArgs, opts?: pulumi.InvokeOptions): Promise<GetEventTriggersResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("mongodbatlas:index/getEventTriggers:getEventTriggers", {
        "appId": args.appId,
        "projectId": args.projectId,
    }, opts);
}

/**
 * A collection of arguments for invoking getEventTriggers.
 */
export interface GetEventTriggersArgs {
    /**
     * The ObjectID of your application.
     */
    readonly appId: string;
    /**
     * The unique ID for the project to get all event triggers.
     */
    readonly projectId: string;
}

/**
 * A collection of values returned by getEventTriggers.
 */
export interface GetEventTriggersResult {
    readonly appId: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly projectId: string;
    /**
     * A list where each represents a Event Trigger.
     */
    readonly results: outputs.GetEventTriggersResult[];
}