// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export class Teams extends pulumi.CustomResource {
    /**
     * Get an existing Teams resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: TeamsState, opts?: pulumi.CustomResourceOptions): Teams {
        return new Teams(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'mongodbatlas:index/teams:Teams';

    /**
     * Returns true if the given object is an instance of Teams.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Teams {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Teams.__pulumiType;
    }

    public readonly name!: pulumi.Output<string>;
    public readonly orgId!: pulumi.Output<string>;
    public /*out*/ readonly teamId!: pulumi.Output<string>;
    public readonly usernames!: pulumi.Output<string[]>;

    /**
     * Create a Teams resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TeamsArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: TeamsArgs | TeamsState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as TeamsState | undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["orgId"] = state ? state.orgId : undefined;
            inputs["teamId"] = state ? state.teamId : undefined;
            inputs["usernames"] = state ? state.usernames : undefined;
        } else {
            const args = argsOrState as TeamsArgs | undefined;
            if (!args || args.orgId === undefined) {
                throw new Error("Missing required property 'orgId'");
            }
            if (!args || args.usernames === undefined) {
                throw new Error("Missing required property 'usernames'");
            }
            inputs["name"] = args ? args.name : undefined;
            inputs["orgId"] = args ? args.orgId : undefined;
            inputs["usernames"] = args ? args.usernames : undefined;
            inputs["teamId"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Teams.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Teams resources.
 */
export interface TeamsState {
    readonly name?: pulumi.Input<string>;
    readonly orgId?: pulumi.Input<string>;
    readonly teamId?: pulumi.Input<string>;
    readonly usernames?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a Teams resource.
 */
export interface TeamsArgs {
    readonly name?: pulumi.Input<string>;
    readonly orgId: pulumi.Input<string>;
    readonly usernames: pulumi.Input<pulumi.Input<string>[]>;
}