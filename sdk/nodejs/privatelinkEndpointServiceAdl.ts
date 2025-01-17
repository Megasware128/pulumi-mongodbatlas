// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * `privatelinkEndpointServiceAdl` Provides an Atlas Data Lake (ADL) and Online Archive PrivateLink endpoint resource.   The same configuration will provide a PrivateLink connection for either Atlas Data Lake or Online Archive.
 *
 * ## Example Usage
 * ### Basic
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as mongodbatlas from "@pulumi/mongodbatlas";
 *
 * const adlTest = new mongodbatlas.PrivatelinkEndpointServiceAdl("adl_test", {
 *     comment: "comments for private link endpoint adl",
 *     endpointId: "<ENDPOINT_ID>",
 *     projectId: "<PROJECT_ID>",
 *     providerName: "AWS",
 *     type: "DATA_LAKE",
 * });
 * ```
 *
 * ## Import
 *
 * ADL privatelink endpoint can be imported using project ID and endpoint ID, in the format `project_id`--`endpoint_id`, e.g.
 *
 * ```sh
 *  $ pulumi import mongodbatlas:index/privatelinkEndpointServiceAdl:PrivatelinkEndpointServiceAdl test 1112222b3bf99403840e8934--vpce-jjg5e24qp93513h03
 * ```
 *
 *  For more information see[MongoDB Atlas API - DataLake](https://docs.mongodb.com/datalake/reference/api/datalakes-api/)
 *
 * and [MongoDB Atlas API - Online Archive](https://docs.atlas.mongodb.com/reference/api/online-archive/) Documentation.
 */
export class PrivatelinkEndpointServiceAdl extends pulumi.CustomResource {
    /**
     * Get an existing PrivatelinkEndpointServiceAdl resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PrivatelinkEndpointServiceAdlState, opts?: pulumi.CustomResourceOptions): PrivatelinkEndpointServiceAdl {
        return new PrivatelinkEndpointServiceAdl(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'mongodbatlas:index/privatelinkEndpointServiceAdl:PrivatelinkEndpointServiceAdl';

    /**
     * Returns true if the given object is an instance of PrivatelinkEndpointServiceAdl.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is PrivatelinkEndpointServiceAdl {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === PrivatelinkEndpointServiceAdl.__pulumiType;
    }

    /**
     * Human-readable string to associate with this private endpoint.
     */
    public readonly comment!: pulumi.Output<string | undefined>;
    /**
     * Unique 22-character alphanumeric string that identifies the private endpoint. Atlas supports AWS private endpoints using the [|aws| PrivateLink](https://aws.amazon.com/privatelink/) feature.
     */
    public readonly endpointId!: pulumi.Output<string>;
    /**
     * Unique 24-digit hexadecimal string that identifies the project.
     */
    public readonly projectId!: pulumi.Output<string>;
    /**
     * Human-readable label that identifies the cloud provider for this endpoint. Atlas supports AWS only. If empty, defaults to AWS.
     */
    public readonly providerName!: pulumi.Output<string>;
    /**
     * Human-readable label that identifies the type of resource to associate with this private endpoint. Atlas supports `DATA_LAKE` only. If empty, defaults to `DATA_LAKE`.
     */
    public readonly type!: pulumi.Output<string>;

    /**
     * Create a PrivatelinkEndpointServiceAdl resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PrivatelinkEndpointServiceAdlArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PrivatelinkEndpointServiceAdlArgs | PrivatelinkEndpointServiceAdlState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as PrivatelinkEndpointServiceAdlState | undefined;
            resourceInputs["comment"] = state ? state.comment : undefined;
            resourceInputs["endpointId"] = state ? state.endpointId : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["providerName"] = state ? state.providerName : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
        } else {
            const args = argsOrState as PrivatelinkEndpointServiceAdlArgs | undefined;
            if ((!args || args.endpointId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'endpointId'");
            }
            if ((!args || args.projectId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'projectId'");
            }
            if ((!args || args.providerName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'providerName'");
            }
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            resourceInputs["comment"] = args ? args.comment : undefined;
            resourceInputs["endpointId"] = args ? args.endpointId : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["providerName"] = args ? args.providerName : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(PrivatelinkEndpointServiceAdl.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering PrivatelinkEndpointServiceAdl resources.
 */
export interface PrivatelinkEndpointServiceAdlState {
    /**
     * Human-readable string to associate with this private endpoint.
     */
    comment?: pulumi.Input<string>;
    /**
     * Unique 22-character alphanumeric string that identifies the private endpoint. Atlas supports AWS private endpoints using the [|aws| PrivateLink](https://aws.amazon.com/privatelink/) feature.
     */
    endpointId?: pulumi.Input<string>;
    /**
     * Unique 24-digit hexadecimal string that identifies the project.
     */
    projectId?: pulumi.Input<string>;
    /**
     * Human-readable label that identifies the cloud provider for this endpoint. Atlas supports AWS only. If empty, defaults to AWS.
     */
    providerName?: pulumi.Input<string>;
    /**
     * Human-readable label that identifies the type of resource to associate with this private endpoint. Atlas supports `DATA_LAKE` only. If empty, defaults to `DATA_LAKE`.
     */
    type?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a PrivatelinkEndpointServiceAdl resource.
 */
export interface PrivatelinkEndpointServiceAdlArgs {
    /**
     * Human-readable string to associate with this private endpoint.
     */
    comment?: pulumi.Input<string>;
    /**
     * Unique 22-character alphanumeric string that identifies the private endpoint. Atlas supports AWS private endpoints using the [|aws| PrivateLink](https://aws.amazon.com/privatelink/) feature.
     */
    endpointId: pulumi.Input<string>;
    /**
     * Unique 24-digit hexadecimal string that identifies the project.
     */
    projectId: pulumi.Input<string>;
    /**
     * Human-readable label that identifies the cloud provider for this endpoint. Atlas supports AWS only. If empty, defaults to AWS.
     */
    providerName: pulumi.Input<string>;
    /**
     * Human-readable label that identifies the type of resource to associate with this private endpoint. Atlas supports `DATA_LAKE` only. If empty, defaults to `DATA_LAKE`.
     */
    type: pulumi.Input<string>;
}
