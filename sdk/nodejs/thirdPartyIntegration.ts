// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * ## Import
 *
 * Third-Party Integration Settings can be imported using project ID and the integration type, in the format `project_id`-`type`, e.g.
 *
 * ```sh
 *  $ pulumi import mongodbatlas:index/thirdPartyIntegration:ThirdPartyIntegration my_user 1112222b3bf99403840e8934-OPS_GENIE
 * ```
 *
 *  See [MongoDB Atlas API](https://docs.atlas.mongodb.com/reference/api/third-party-integration-settings-create/) Documentation for more information.
 */
export class ThirdPartyIntegration extends pulumi.CustomResource {
    /**
     * Get an existing ThirdPartyIntegration resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ThirdPartyIntegrationState, opts?: pulumi.CustomResourceOptions): ThirdPartyIntegration {
        return new ThirdPartyIntegration(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'mongodbatlas:index/thirdPartyIntegration:ThirdPartyIntegration';

    /**
     * Returns true if the given object is an instance of ThirdPartyIntegration.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ThirdPartyIntegration {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ThirdPartyIntegration.__pulumiType;
    }

    /**
     * Unique identifier of your New Relic account.
     */
    public readonly accountId!: pulumi.Output<string | undefined>;
    /**
     * Your API Key.
     */
    public readonly apiKey!: pulumi.Output<string | undefined>;
    /**
     * Your API Token.
     */
    public readonly apiToken!: pulumi.Output<string | undefined>;
    public readonly channelName!: pulumi.Output<string | undefined>;
    /**
     * Your Flowdock Flow name.
     */
    public readonly flowName!: pulumi.Output<string | undefined>;
    /**
     * Your License Key.
     */
    public readonly licenseKey!: pulumi.Output<string | undefined>;
    /**
     * Your Flowdock organization name.
     * * `WEBHOOK`
     */
    public readonly orgName!: pulumi.Output<string | undefined>;
    /**
     * The unique ID for the project to get all Third-Party service integrations
     */
    public readonly projectId!: pulumi.Output<string>;
    /**
     * Your Insights Query Key.
     * * `OPS_GENIE`
     */
    public readonly readToken!: pulumi.Output<string | undefined>;
    /**
     * Indicates which API URL to use, either US or EU. Opsgenie will use US by default.
     * * `VICTOR_OPS`
     */
    public readonly region!: pulumi.Output<string | undefined>;
    /**
     * An optional field for your Routing Key.
     * * `FLOWDOCK`
     */
    public readonly routingKey!: pulumi.Output<string | undefined>;
    /**
     * An optional field for your webhook secret.
     */
    public readonly secret!: pulumi.Output<string | undefined>;
    /**
     * Your Service Key.
     * * `DATADOG`
     */
    public readonly serviceKey!: pulumi.Output<string | undefined>;
    public readonly teamName!: pulumi.Output<string | undefined>;
    /**
     * Third-Party Integration Settings type 
     * * PAGER_DUTY
     * * DATADOG
     * * NEW_RELIC
     * * OPS_GENIE
     * * VICTOR_OPS
     * * FLOWDOCK
     * * WEBHOOK
     */
    public readonly type!: pulumi.Output<string>;
    /**
     * Your webhook URL.
     */
    public readonly url!: pulumi.Output<string | undefined>;
    /**
     * Your Insights Insert Key.
     */
    public readonly writeToken!: pulumi.Output<string | undefined>;

    /**
     * Create a ThirdPartyIntegration resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ThirdPartyIntegrationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ThirdPartyIntegrationArgs | ThirdPartyIntegrationState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ThirdPartyIntegrationState | undefined;
            inputs["accountId"] = state ? state.accountId : undefined;
            inputs["apiKey"] = state ? state.apiKey : undefined;
            inputs["apiToken"] = state ? state.apiToken : undefined;
            inputs["channelName"] = state ? state.channelName : undefined;
            inputs["flowName"] = state ? state.flowName : undefined;
            inputs["licenseKey"] = state ? state.licenseKey : undefined;
            inputs["orgName"] = state ? state.orgName : undefined;
            inputs["projectId"] = state ? state.projectId : undefined;
            inputs["readToken"] = state ? state.readToken : undefined;
            inputs["region"] = state ? state.region : undefined;
            inputs["routingKey"] = state ? state.routingKey : undefined;
            inputs["secret"] = state ? state.secret : undefined;
            inputs["serviceKey"] = state ? state.serviceKey : undefined;
            inputs["teamName"] = state ? state.teamName : undefined;
            inputs["type"] = state ? state.type : undefined;
            inputs["url"] = state ? state.url : undefined;
            inputs["writeToken"] = state ? state.writeToken : undefined;
        } else {
            const args = argsOrState as ThirdPartyIntegrationArgs | undefined;
            if ((!args || args.projectId === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'projectId'");
            }
            if ((!args || args.type === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'type'");
            }
            inputs["accountId"] = args ? args.accountId : undefined;
            inputs["apiKey"] = args ? args.apiKey : undefined;
            inputs["apiToken"] = args ? args.apiToken : undefined;
            inputs["channelName"] = args ? args.channelName : undefined;
            inputs["flowName"] = args ? args.flowName : undefined;
            inputs["licenseKey"] = args ? args.licenseKey : undefined;
            inputs["orgName"] = args ? args.orgName : undefined;
            inputs["projectId"] = args ? args.projectId : undefined;
            inputs["readToken"] = args ? args.readToken : undefined;
            inputs["region"] = args ? args.region : undefined;
            inputs["routingKey"] = args ? args.routingKey : undefined;
            inputs["secret"] = args ? args.secret : undefined;
            inputs["serviceKey"] = args ? args.serviceKey : undefined;
            inputs["teamName"] = args ? args.teamName : undefined;
            inputs["type"] = args ? args.type : undefined;
            inputs["url"] = args ? args.url : undefined;
            inputs["writeToken"] = args ? args.writeToken : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(ThirdPartyIntegration.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ThirdPartyIntegration resources.
 */
export interface ThirdPartyIntegrationState {
    /**
     * Unique identifier of your New Relic account.
     */
    readonly accountId?: pulumi.Input<string>;
    /**
     * Your API Key.
     */
    readonly apiKey?: pulumi.Input<string>;
    /**
     * Your API Token.
     */
    readonly apiToken?: pulumi.Input<string>;
    readonly channelName?: pulumi.Input<string>;
    /**
     * Your Flowdock Flow name.
     */
    readonly flowName?: pulumi.Input<string>;
    /**
     * Your License Key.
     */
    readonly licenseKey?: pulumi.Input<string>;
    /**
     * Your Flowdock organization name.
     * * `WEBHOOK`
     */
    readonly orgName?: pulumi.Input<string>;
    /**
     * The unique ID for the project to get all Third-Party service integrations
     */
    readonly projectId?: pulumi.Input<string>;
    /**
     * Your Insights Query Key.
     * * `OPS_GENIE`
     */
    readonly readToken?: pulumi.Input<string>;
    /**
     * Indicates which API URL to use, either US or EU. Opsgenie will use US by default.
     * * `VICTOR_OPS`
     */
    readonly region?: pulumi.Input<string>;
    /**
     * An optional field for your Routing Key.
     * * `FLOWDOCK`
     */
    readonly routingKey?: pulumi.Input<string>;
    /**
     * An optional field for your webhook secret.
     */
    readonly secret?: pulumi.Input<string>;
    /**
     * Your Service Key.
     * * `DATADOG`
     */
    readonly serviceKey?: pulumi.Input<string>;
    readonly teamName?: pulumi.Input<string>;
    /**
     * Third-Party Integration Settings type 
     * * PAGER_DUTY
     * * DATADOG
     * * NEW_RELIC
     * * OPS_GENIE
     * * VICTOR_OPS
     * * FLOWDOCK
     * * WEBHOOK
     */
    readonly type?: pulumi.Input<string>;
    /**
     * Your webhook URL.
     */
    readonly url?: pulumi.Input<string>;
    /**
     * Your Insights Insert Key.
     */
    readonly writeToken?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ThirdPartyIntegration resource.
 */
export interface ThirdPartyIntegrationArgs {
    /**
     * Unique identifier of your New Relic account.
     */
    readonly accountId?: pulumi.Input<string>;
    /**
     * Your API Key.
     */
    readonly apiKey?: pulumi.Input<string>;
    /**
     * Your API Token.
     */
    readonly apiToken?: pulumi.Input<string>;
    readonly channelName?: pulumi.Input<string>;
    /**
     * Your Flowdock Flow name.
     */
    readonly flowName?: pulumi.Input<string>;
    /**
     * Your License Key.
     */
    readonly licenseKey?: pulumi.Input<string>;
    /**
     * Your Flowdock organization name.
     * * `WEBHOOK`
     */
    readonly orgName?: pulumi.Input<string>;
    /**
     * The unique ID for the project to get all Third-Party service integrations
     */
    readonly projectId: pulumi.Input<string>;
    /**
     * Your Insights Query Key.
     * * `OPS_GENIE`
     */
    readonly readToken?: pulumi.Input<string>;
    /**
     * Indicates which API URL to use, either US or EU. Opsgenie will use US by default.
     * * `VICTOR_OPS`
     */
    readonly region?: pulumi.Input<string>;
    /**
     * An optional field for your Routing Key.
     * * `FLOWDOCK`
     */
    readonly routingKey?: pulumi.Input<string>;
    /**
     * An optional field for your webhook secret.
     */
    readonly secret?: pulumi.Input<string>;
    /**
     * Your Service Key.
     * * `DATADOG`
     */
    readonly serviceKey?: pulumi.Input<string>;
    readonly teamName?: pulumi.Input<string>;
    /**
     * Third-Party Integration Settings type 
     * * PAGER_DUTY
     * * DATADOG
     * * NEW_RELIC
     * * OPS_GENIE
     * * VICTOR_OPS
     * * FLOWDOCK
     * * WEBHOOK
     */
    readonly type: pulumi.Input<string>;
    /**
     * Your webhook URL.
     */
    readonly url?: pulumi.Input<string>;
    /**
     * Your Insights Insert Key.
     */
    readonly writeToken?: pulumi.Input<string>;
}
