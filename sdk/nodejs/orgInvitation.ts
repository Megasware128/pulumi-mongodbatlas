// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * ## Example Usage
 * ### S
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as mongodbatlas from "@pulumi/mongodbatlas";
 *
 * const test0 = new mongodbatlas.OrgInvitation("test0", {
 *     orgId: "<ORG-ID>",
 *     roles: ["ORG_OWNER"],
 *     username: "test0-acc-username",
 * });
 * ```
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as mongodbatlas from "@pulumi/mongodbatlas";
 *
 * const test0 = new mongodbatlas.OrgInvitation("test0", {
 *     orgId: "<ORG-ID>",
 *     roles: [
 *         "ORG_MEMBER",
 *         "ORG_BILLING_ADMIN",
 *     ],
 *     username: "test0-acc-username",
 * });
 * ```
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as mongodbatlas from "@pulumi/mongodbatlas";
 *
 * const test1 = new mongodbatlas.OrgInvitation("test1", {
 *     orgId: "<ORG-ID>",
 *     roles: ["ORG_MEMBER"],
 *     teamsIds: [
 *         "<TEAM-0-ID>",
 *         "<TEAM-1-ID>",
 *     ],
 *     username: "test1-acc-username",
 * });
 * ```
 *
 * ## Import
 *
 * ~> **IMPORTANT:** An organization invitation can **not** be imported once it has been accepted. Import a user's invitation to an organization by separating the `org_id` and the `username` with a hyphen
 *
 * ```sh
 *  $ pulumi import mongodbatlas:index/orgInvitation:OrgInvitation my_user 1112222b3bf99403840e8934-my_user@mongodb.com
 * ```
 */
export class OrgInvitation extends pulumi.CustomResource {
    /**
     * Get an existing OrgInvitation resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: OrgInvitationState, opts?: pulumi.CustomResourceOptions): OrgInvitation {
        return new OrgInvitation(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'mongodbatlas:index/orgInvitation:OrgInvitation';

    /**
     * Returns true if the given object is an instance of OrgInvitation.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is OrgInvitation {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === OrgInvitation.__pulumiType;
    }

    /**
     * Timestamp in ISO 8601 date and time format in UTC when Atlas sent the invitation.
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * Timestamp in ISO 8601 date and time format in UTC when the invitation expires. Users have 30 days to accept an invitation.
     */
    public /*out*/ readonly expiresAt!: pulumi.Output<string>;
    /**
     * Unique 24-hexadecimal digit string that identifies the invitation in Atlas.
     */
    public /*out*/ readonly invitationId!: pulumi.Output<string>;
    /**
     * Atlas user who invited `username` to the organization.
     */
    public /*out*/ readonly inviterUsername!: pulumi.Output<string>;
    /**
     * Unique 24-hexadecimal digit string that identifies the organization to which you want to invite a user.
     */
    public readonly orgId!: pulumi.Output<string>;
    /**
     * Atlas roles to assign to the invited user. If the user accepts the invitation, Atlas assigns these roles to them. The following options are available:
     * * ORG_OWNER
     * * ORG_GROUP_CREATOR
     * * ORG_BILLING_ADMIN
     * * ORG_READ_ONLY
     * * ORG_MEMBER
     */
    public readonly roles!: pulumi.Output<string[]>;
    /**
     * An array of unique 24-hexadecimal digit strings that identify the teams that the user was invited to join.
     */
    public readonly teamsIds!: pulumi.Output<string[] | undefined>;
    /**
     * Email address of the invited user. This is the address to which Atlas sends the invite. If the user accepts the invitation, they log in to Atlas with this username.
     */
    public readonly username!: pulumi.Output<string>;

    /**
     * Create a OrgInvitation resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: OrgInvitationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: OrgInvitationArgs | OrgInvitationState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as OrgInvitationState | undefined;
            resourceInputs["createdAt"] = state ? state.createdAt : undefined;
            resourceInputs["expiresAt"] = state ? state.expiresAt : undefined;
            resourceInputs["invitationId"] = state ? state.invitationId : undefined;
            resourceInputs["inviterUsername"] = state ? state.inviterUsername : undefined;
            resourceInputs["orgId"] = state ? state.orgId : undefined;
            resourceInputs["roles"] = state ? state.roles : undefined;
            resourceInputs["teamsIds"] = state ? state.teamsIds : undefined;
            resourceInputs["username"] = state ? state.username : undefined;
        } else {
            const args = argsOrState as OrgInvitationArgs | undefined;
            if ((!args || args.orgId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'orgId'");
            }
            if ((!args || args.roles === undefined) && !opts.urn) {
                throw new Error("Missing required property 'roles'");
            }
            if ((!args || args.username === undefined) && !opts.urn) {
                throw new Error("Missing required property 'username'");
            }
            resourceInputs["orgId"] = args ? args.orgId : undefined;
            resourceInputs["roles"] = args ? args.roles : undefined;
            resourceInputs["teamsIds"] = args ? args.teamsIds : undefined;
            resourceInputs["username"] = args ? args.username : undefined;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["expiresAt"] = undefined /*out*/;
            resourceInputs["invitationId"] = undefined /*out*/;
            resourceInputs["inviterUsername"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(OrgInvitation.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering OrgInvitation resources.
 */
export interface OrgInvitationState {
    /**
     * Timestamp in ISO 8601 date and time format in UTC when Atlas sent the invitation.
     */
    createdAt?: pulumi.Input<string>;
    /**
     * Timestamp in ISO 8601 date and time format in UTC when the invitation expires. Users have 30 days to accept an invitation.
     */
    expiresAt?: pulumi.Input<string>;
    /**
     * Unique 24-hexadecimal digit string that identifies the invitation in Atlas.
     */
    invitationId?: pulumi.Input<string>;
    /**
     * Atlas user who invited `username` to the organization.
     */
    inviterUsername?: pulumi.Input<string>;
    /**
     * Unique 24-hexadecimal digit string that identifies the organization to which you want to invite a user.
     */
    orgId?: pulumi.Input<string>;
    /**
     * Atlas roles to assign to the invited user. If the user accepts the invitation, Atlas assigns these roles to them. The following options are available:
     * * ORG_OWNER
     * * ORG_GROUP_CREATOR
     * * ORG_BILLING_ADMIN
     * * ORG_READ_ONLY
     * * ORG_MEMBER
     */
    roles?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * An array of unique 24-hexadecimal digit strings that identify the teams that the user was invited to join.
     */
    teamsIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Email address of the invited user. This is the address to which Atlas sends the invite. If the user accepts the invitation, they log in to Atlas with this username.
     */
    username?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a OrgInvitation resource.
 */
export interface OrgInvitationArgs {
    /**
     * Unique 24-hexadecimal digit string that identifies the organization to which you want to invite a user.
     */
    orgId: pulumi.Input<string>;
    /**
     * Atlas roles to assign to the invited user. If the user accepts the invitation, Atlas assigns these roles to them. The following options are available:
     * * ORG_OWNER
     * * ORG_GROUP_CREATOR
     * * ORG_BILLING_ADMIN
     * * ORG_READ_ONLY
     * * ORG_MEMBER
     */
    roles: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * An array of unique 24-hexadecimal digit strings that identify the teams that the user was invited to join.
     */
    teamsIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Email address of the invited user. This is the address to which Atlas sends the invite. If the user accepts the invitation, they log in to Atlas with this username.
     */
    username: pulumi.Input<string>;
}
