// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

export class CloudProviderAccessAuthorization extends pulumi.CustomResource {
    /**
     * Get an existing CloudProviderAccessAuthorization resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: CloudProviderAccessAuthorizationState, opts?: pulumi.CustomResourceOptions): CloudProviderAccessAuthorization {
        return new CloudProviderAccessAuthorization(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'mongodbatlas:index/cloudProviderAccessAuthorization:CloudProviderAccessAuthorization';

    /**
     * Returns true if the given object is an instance of CloudProviderAccessAuthorization.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is CloudProviderAccessAuthorization {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === CloudProviderAccessAuthorization.__pulumiType;
    }

    public /*out*/ readonly authorizedDate!: pulumi.Output<string>;
    public readonly aws!: pulumi.Output<outputs.CloudProviderAccessAuthorizationAws | undefined>;
    public /*out*/ readonly featureUsages!: pulumi.Output<outputs.CloudProviderAccessAuthorizationFeatureUsage[]>;
    public readonly projectId!: pulumi.Output<string>;
    public readonly roleId!: pulumi.Output<string>;

    /**
     * Create a CloudProviderAccessAuthorization resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CloudProviderAccessAuthorizationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: CloudProviderAccessAuthorizationArgs | CloudProviderAccessAuthorizationState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as CloudProviderAccessAuthorizationState | undefined;
            resourceInputs["authorizedDate"] = state ? state.authorizedDate : undefined;
            resourceInputs["aws"] = state ? state.aws : undefined;
            resourceInputs["featureUsages"] = state ? state.featureUsages : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["roleId"] = state ? state.roleId : undefined;
        } else {
            const args = argsOrState as CloudProviderAccessAuthorizationArgs | undefined;
            if ((!args || args.projectId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'projectId'");
            }
            if ((!args || args.roleId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'roleId'");
            }
            resourceInputs["aws"] = args ? args.aws : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["roleId"] = args ? args.roleId : undefined;
            resourceInputs["authorizedDate"] = undefined /*out*/;
            resourceInputs["featureUsages"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(CloudProviderAccessAuthorization.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering CloudProviderAccessAuthorization resources.
 */
export interface CloudProviderAccessAuthorizationState {
    authorizedDate?: pulumi.Input<string>;
    aws?: pulumi.Input<inputs.CloudProviderAccessAuthorizationAws>;
    featureUsages?: pulumi.Input<pulumi.Input<inputs.CloudProviderAccessAuthorizationFeatureUsage>[]>;
    projectId?: pulumi.Input<string>;
    roleId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a CloudProviderAccessAuthorization resource.
 */
export interface CloudProviderAccessAuthorizationArgs {
    aws?: pulumi.Input<inputs.CloudProviderAccessAuthorizationAws>;
    projectId: pulumi.Input<string>;
    roleId: pulumi.Input<string>;
}
