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

    public readonly apiKeys!: pulumi.Output<outputs.ProjectApiKey[]>;
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
     * It allows users to disable the creation of the default alert settings. By default, this flag is set to true.
     */
    public readonly withDefaultAlertsSettings!: pulumi.Output<boolean | undefined>;

    /**
     * Create a Project resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ProjectArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ProjectArgs | ProjectState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ProjectState | undefined;
            resourceInputs["apiKeys"] = state ? state.apiKeys : undefined;
            resourceInputs["clusterCount"] = state ? state.clusterCount : undefined;
            resourceInputs["created"] = state ? state.created : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["orgId"] = state ? state.orgId : undefined;
            resourceInputs["projectOwnerId"] = state ? state.projectOwnerId : undefined;
            resourceInputs["teams"] = state ? state.teams : undefined;
            resourceInputs["withDefaultAlertsSettings"] = state ? state.withDefaultAlertsSettings : undefined;
        } else {
            const args = argsOrState as ProjectArgs | undefined;
            if ((!args || args.orgId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'orgId'");
            }
            resourceInputs["apiKeys"] = args ? args.apiKeys : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["orgId"] = args ? args.orgId : undefined;
            resourceInputs["projectOwnerId"] = args ? args.projectOwnerId : undefined;
            resourceInputs["teams"] = args ? args.teams : undefined;
            resourceInputs["withDefaultAlertsSettings"] = args ? args.withDefaultAlertsSettings : undefined;
            resourceInputs["clusterCount"] = undefined /*out*/;
            resourceInputs["created"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Project.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Project resources.
 */
export interface ProjectState {
    apiKeys?: pulumi.Input<pulumi.Input<inputs.ProjectApiKey>[]>;
    /**
     * The number of Atlas clusters deployed in the project..
     */
    clusterCount?: pulumi.Input<number>;
    /**
     * The ISO-8601-formatted timestamp of when Atlas created the project..
     */
    created?: pulumi.Input<string>;
    /**
     * The name of the project you want to create. (Cannot be changed via this Provider after creation.)
     */
    name?: pulumi.Input<string>;
    /**
     * The ID of the organization you want to create the project within.
     */
    orgId?: pulumi.Input<string>;
    /**
     * Unique 24-hexadecimal digit string that identifies the Atlas user account to be granted the [Project Owner](https://docs.atlas.mongodb.com/reference/user-roles/#mongodb-authrole-Project-Owner) role on the specified project. If you set this parameter, it overrides the default value of the oldest [Organization Owner](https://docs.atlas.mongodb.com/reference/user-roles/#mongodb-authrole-Organization-Owner).
     */
    projectOwnerId?: pulumi.Input<string>;
    teams?: pulumi.Input<pulumi.Input<inputs.ProjectTeam>[]>;
    /**
     * It allows users to disable the creation of the default alert settings. By default, this flag is set to true.
     */
    withDefaultAlertsSettings?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a Project resource.
 */
export interface ProjectArgs {
    apiKeys?: pulumi.Input<pulumi.Input<inputs.ProjectApiKey>[]>;
    /**
     * The name of the project you want to create. (Cannot be changed via this Provider after creation.)
     */
    name?: pulumi.Input<string>;
    /**
     * The ID of the organization you want to create the project within.
     */
    orgId: pulumi.Input<string>;
    /**
     * Unique 24-hexadecimal digit string that identifies the Atlas user account to be granted the [Project Owner](https://docs.atlas.mongodb.com/reference/user-roles/#mongodb-authrole-Project-Owner) role on the specified project. If you set this parameter, it overrides the default value of the oldest [Organization Owner](https://docs.atlas.mongodb.com/reference/user-roles/#mongodb-authrole-Organization-Owner).
     */
    projectOwnerId?: pulumi.Input<string>;
    teams?: pulumi.Input<pulumi.Input<inputs.ProjectTeam>[]>;
    /**
     * It allows users to disable the creation of the default alert settings. By default, this flag is set to true.
     */
    withDefaultAlertsSettings?: pulumi.Input<boolean>;
}
