// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

export class CloudProviderAccessSetup extends pulumi.CustomResource {
    /**
     * Get an existing CloudProviderAccessSetup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: CloudProviderAccessSetupState, opts?: pulumi.CustomResourceOptions): CloudProviderAccessSetup {
        return new CloudProviderAccessSetup(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'mongodbatlas:index/cloudProviderAccessSetup:CloudProviderAccessSetup';

    /**
     * Returns true if the given object is an instance of CloudProviderAccessSetup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is CloudProviderAccessSetup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === CloudProviderAccessSetup.__pulumiType;
    }

    /**
     * @deprecated use aws_config instead
     */
    public /*out*/ readonly aws!: pulumi.Output<{[key: string]: string}>;
    public /*out*/ readonly awsConfigs!: pulumi.Output<outputs.CloudProviderAccessSetupAwsConfig[]>;
    public /*out*/ readonly createdDate!: pulumi.Output<string>;
    public readonly projectId!: pulumi.Output<string>;
    public readonly providerName!: pulumi.Output<string>;
    public /*out*/ readonly roleId!: pulumi.Output<string>;

    /**
     * Create a CloudProviderAccessSetup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CloudProviderAccessSetupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: CloudProviderAccessSetupArgs | CloudProviderAccessSetupState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as CloudProviderAccessSetupState | undefined;
            resourceInputs["aws"] = state ? state.aws : undefined;
            resourceInputs["awsConfigs"] = state ? state.awsConfigs : undefined;
            resourceInputs["createdDate"] = state ? state.createdDate : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["providerName"] = state ? state.providerName : undefined;
            resourceInputs["roleId"] = state ? state.roleId : undefined;
        } else {
            const args = argsOrState as CloudProviderAccessSetupArgs | undefined;
            if ((!args || args.projectId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'projectId'");
            }
            if ((!args || args.providerName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'providerName'");
            }
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["providerName"] = args ? args.providerName : undefined;
            resourceInputs["aws"] = undefined /*out*/;
            resourceInputs["awsConfigs"] = undefined /*out*/;
            resourceInputs["createdDate"] = undefined /*out*/;
            resourceInputs["roleId"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(CloudProviderAccessSetup.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering CloudProviderAccessSetup resources.
 */
export interface CloudProviderAccessSetupState {
    /**
     * @deprecated use aws_config instead
     */
    aws?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    awsConfigs?: pulumi.Input<pulumi.Input<inputs.CloudProviderAccessSetupAwsConfig>[]>;
    createdDate?: pulumi.Input<string>;
    projectId?: pulumi.Input<string>;
    providerName?: pulumi.Input<string>;
    roleId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a CloudProviderAccessSetup resource.
 */
export interface CloudProviderAccessSetupArgs {
    projectId: pulumi.Input<string>;
    providerName: pulumi.Input<string>;
}
