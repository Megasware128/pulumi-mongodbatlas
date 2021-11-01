// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * ## Import
 *
 * Project must be imported using project ID, e.g.
 *
 * ```sh
 *  $ pulumi import mongodbatlas:index/project:Project my_project 5d09d6a59ccf6445652a444a
 * ```
 *
 *  For more information see[MongoDB Atlas API Reference.](https://docs.atlas.mongodb.com/reference/api/projects/) - [and MongoDB Atlas API - Teams](https://docs.atlas.mongodb.com/reference/api/teams/) Documentation for more information.
 */
export class Project extends pulumi.CustomResource {
    /**
     * Get an existing Project resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ProjectState, opts?: pulumi.CustomResourceOptions): Project {
        return new Project(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'mongodbatlas:index/project:Project';

    /**
     * Returns true if the given object is an instance of Project.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Project {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Project.__pulumiType;
    }

    /**
     * The number of Atlas clusters deployed in the project..
     */
    public /*out*/ readonly clusterCount!: pulumi.Output<number>;
    /**
     * The ISO-8601-formatted timestamp of when Atlas created the project..
     */
    public /*out*/ readonly created!: pulumi.Output<string>;
    /**
     * The name of the project you want to create. (Cannot be changed via this Provider after creation.)
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The ID of the organization you want to create the project within.
     */
    public readonly orgId!: pulumi.Output<string>;
    /**
     * Unique 24-hexadecimal digit string that identifies the Atlas user account to be granted the [Project Owner](https://docs.atlas.mongodb.com/reference/user-roles/#mongodb-authrole-Project-Owner) role on the specified project. If you set this parameter, it overrides the default value of the oldest [Organization Owner](https://docs.atlas.mongodb.com/reference/user-roles/#mongodb-authrole-Organization-Owner).
     */
    public readonly projectOwnerId!: pulumi.Output<string | undefined>;
    public readonly teams!: pulumi.Output<outputs.ProjectTeam[] | undefined>;

    /**
     * Create a Project resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ProjectArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ProjectArgs | ProjectState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ProjectState | undefined;
            inputs["clusterCount"] = state ? state.clusterCount : undefined;
            inputs["created"] = state ? state.created : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["orgId"] = state ? state.orgId : undefined;
            inputs["projectOwnerId"] = state ? state.projectOwnerId : undefined;
            inputs["teams"] = state ? state.teams : undefined;
        } else {
            const args = argsOrState as ProjectArgs | undefined;
            if ((!args || args.orgId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'orgId'");
            }
            inputs["name"] = args ? args.name : undefined;
            inputs["orgId"] = args ? args.orgId : undefined;
            inputs["projectOwnerId"] = args ? args.projectOwnerId : undefined;
            inputs["teams"] = args ? args.teams : undefined;
            inputs["clusterCount"] = undefined /*out*/;
            inputs["created"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Project.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Project resources.
 */
export interface ProjectState {
    /**
     * The number of Atlas clusters deployed in the project..
     */
    readonly clusterCount?: pulumi.Input<number>;
    /**
     * The ISO-8601-formatted timestamp of when Atlas created the project..
     */
    readonly created?: pulumi.Input<string>;
    /**
     * The name of the project you want to create. (Cannot be changed via this Provider after creation.)
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The ID of the organization you want to create the project within.
     */
    readonly orgId?: pulumi.Input<string>;
    /**
     * Unique 24-hexadecimal digit string that identifies the Atlas user account to be granted the [Project Owner](https://docs.atlas.mongodb.com/reference/user-roles/#mongodb-authrole-Project-Owner) role on the specified project. If you set this parameter, it overrides the default value of the oldest [Organization Owner](https://docs.atlas.mongodb.com/reference/user-roles/#mongodb-authrole-Organization-Owner).
     */
    readonly projectOwnerId?: pulumi.Input<string>;
    readonly teams?: pulumi.Input<pulumi.Input<inputs.ProjectTeam>[]>;
}

/**
 * The set of arguments for constructing a Project resource.
 */
export interface ProjectArgs {
    /**
     * The name of the project you want to create. (Cannot be changed via this Provider after creation.)
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The ID of the organization you want to create the project within.
     */
    readonly orgId: pulumi.Input<string>;
    /**
     * Unique 24-hexadecimal digit string that identifies the Atlas user account to be granted the [Project Owner](https://docs.atlas.mongodb.com/reference/user-roles/#mongodb-authrole-Project-Owner) role on the specified project. If you set this parameter, it overrides the default value of the oldest [Organization Owner](https://docs.atlas.mongodb.com/reference/user-roles/#mongodb-authrole-Organization-Owner).
     */
    readonly projectOwnerId?: pulumi.Input<string>;
    readonly teams?: pulumi.Input<pulumi.Input<inputs.ProjectTeam>[]>;
}
