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
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as CloudProviderAccessSetupState | undefined;
            inputs["aws"] = state ? state.aws : undefined;
            inputs["awsConfigs"] = state ? state.awsConfigs : undefined;
            inputs["createdDate"] = state ? state.createdDate : undefined;
            inputs["projectId"] = state ? state.projectId : undefined;
            inputs["providerName"] = state ? state.providerName : undefined;
            inputs["roleId"] = state ? state.roleId : undefined;
        } else {
            const args = argsOrState as CloudProviderAccessSetupArgs | undefined;
            if ((!args || args.projectId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'projectId'");
            }
            if ((!args || args.providerName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'providerName'");
            }
            inputs["projectId"] = args ? args.projectId : undefined;
            inputs["providerName"] = args ? args.providerName : undefined;
            inputs["aws"] = undefined /*out*/;
            inputs["awsConfigs"] = undefined /*out*/;
            inputs["createdDate"] = undefined /*out*/;
            inputs["roleId"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(CloudProviderAccessSetup.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering CloudProviderAccessSetup resources.
 */
export interface CloudProviderAccessSetupState {
    /**
     * @deprecated use aws_config instead
     */
    readonly aws?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    readonly awsConfigs?: pulumi.Input<pulumi.Input<inputs.CloudProviderAccessSetupAwsConfig>[]>;
    readonly createdDate?: pulumi.Input<string>;
    readonly projectId?: pulumi.Input<string>;
    readonly providerName?: pulumi.Input<string>;
    readonly roleId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a CloudProviderAccessSetup resource.
 */
export interface CloudProviderAccessSetupArgs {
    readonly projectId: pulumi.Input<string>;
    readonly providerName: pulumi.Input<string>;
}