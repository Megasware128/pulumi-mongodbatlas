// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * `mongodbatlas..Auditing` describes a Auditing.
 * 
 * > **NOTE:** Groups and projects are synonymous terms. You may find **group_id** in the official documentation.
 * 
 * 
 * ## Example Usage
 * 
 * 
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as mongodbatlas from "@pulumi/mongodbatlas";
 * 
 * const testMongodbatlasAuditing = new mongodbatlas.Auditing("test", {
 *     auditAuthorizationSuccess: false,
 *     auditFilter: "{ 'atype': 'authenticate', 'param': {   'user': 'auditAdmin',   'db': 'admin',   'mechanism': 'SCRAM-SHA-1' }}",
 *     enabled: true,
 *     projectId: "<project-id>",
 * });
 * const testAuditing = testMongodbatlasAuditing.id.apply(id => mongodbatlas.getAuditing({
 *     projectId: id,
 * }, { async: true }));
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-mongodbatlas/blob/master/website/docs/d/auditing.html.markdown.
 */
export function getAuditing(args: GetAuditingArgs, opts?: pulumi.InvokeOptions): Promise<GetAuditingResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("mongodbatlas:index/getAuditing:getAuditing", {
        "projectId": args.projectId,
    }, opts);
}

/**
 * A collection of arguments for invoking getAuditing.
 */
export interface GetAuditingArgs {
    /**
     * The unique ID for the project to create the database user.
     */
    readonly projectId: string;
}

/**
 * A collection of values returned by getAuditing.
 */
export interface GetAuditingResult {
    /**
     * JSON-formatted audit filter used by the project
     */
    readonly auditAuthorizationSuccess: boolean;
    /**
     * Indicates whether the auditing system captures successful authentication attempts for audit filters using the "atype" : "authCheck" auditing event. For more information, see auditAuthorizationSuccess
     */
    readonly auditFilter: string;
    /**
     * Denotes the configuration method for the audit filter. Possible values are: NONE - auditing not configured for the project.m FILTER_BUILDER - auditing configured via Atlas UI filter builderm FILTER_JSON - auditing configured via Atlas custom filter or API.
     */
    readonly configurationType: string;
    /**
     * Denotes whether or not the project associated with the {GROUP-ID} has database auditing enabled.
     */
    readonly enabled: boolean;
    readonly projectId: string;
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}