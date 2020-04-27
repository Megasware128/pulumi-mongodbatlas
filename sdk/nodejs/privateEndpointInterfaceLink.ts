// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * `mongodbatlas..PrivateEndpointInterfaceLink` provides a Private Endpoint Interface Link resource. This represents a Private Endpoint Interface Link, which adds one interface endpoint to a private endpoint connection in an Atlas project.
 * 
 * > **IMPORTANT:**You must have one of the following roles to successfully handle the resource:
 *   * Organization Owner
 *   * Project Owner
 * 
 * > **NOTE:** Groups and projects are synonymous terms. You may find groupId in the official documentation.
 * 
 * 
 * ## Example Usage
 * 
 * 
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * import * as mongodbatlas from "@pulumi/mongodbatlas";
 * 
 * const testPrivateEndpoint = new mongodbatlas.PrivateEndpoint("test", {
 *     projectId: "<PROJECT_ID>",
 *     providerName: "AWS",
 *     region: "us-east-1",
 * });
 * const ptfeService = new aws.ec2.VpcEndpoint("ptfeService", {
 *     securityGroupIds: ["sg-3f238186"],
 *     serviceName: testPrivateEndpoint.endpointServiceName,
 *     subnetIds: ["subnet-de0406d2"],
 *     vpcEndpointType: "Interface",
 *     vpcId: "vpc-7fc0a543",
 * });
 * const testPrivateEndpointInterfaceLink = new mongodbatlas.PrivateEndpointInterfaceLink("test", {
 *     interfaceEndpointId: ptfeService.id,
 *     privateLinkId: testPrivateEndpoint.privateLinkId,
 *     projectId: testPrivateEndpoint.projectId,
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-mongodbatlas/blob/master/website/docs/r/private_endpoint_interface_link.html.markdown.
 */
export class PrivateEndpointInterfaceLink extends pulumi.CustomResource {
    /**
     * Get an existing PrivateEndpointInterfaceLink resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PrivateEndpointInterfaceLinkState, opts?: pulumi.CustomResourceOptions): PrivateEndpointInterfaceLink {
        return new PrivateEndpointInterfaceLink(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'mongodbatlas:index/privateEndpointInterfaceLink:PrivateEndpointInterfaceLink';

    /**
     * Returns true if the given object is an instance of PrivateEndpointInterfaceLink.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is PrivateEndpointInterfaceLink {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === PrivateEndpointInterfaceLink.__pulumiType;
    }

    /**
     * Status of the interface endpoint.
     * Returns one of the following values:
     */
    public /*out*/ readonly connectionStatus!: pulumi.Output<string>;
    /**
     * Indicates if Atlas received a request to remove the interface endpoint from the private endpoint connection.
     */
    public /*out*/ readonly deleteRequested!: pulumi.Output<boolean>;
    /**
     * Error message pertaining to the interface endpoint. Returns null if there are no errors.
     */
    public /*out*/ readonly errorMessage!: pulumi.Output<string>;
    /**
     * Unique identifier of the interface endpoint you created in your VPC with the AWS resource.
     */
    public readonly interfaceEndpointId!: pulumi.Output<string>;
    /**
     * Unique identifier of the AWS PrivateLink connection which is created by `mongodbatlas..PrivateEndpoint` resource.
     */
    public readonly privateLinkId!: pulumi.Output<string>;
    /**
     * Unique identifier for the project.
     */
    public readonly projectId!: pulumi.Output<string>;

    /**
     * Create a PrivateEndpointInterfaceLink resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PrivateEndpointInterfaceLinkArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PrivateEndpointInterfaceLinkArgs | PrivateEndpointInterfaceLinkState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as PrivateEndpointInterfaceLinkState | undefined;
            inputs["connectionStatus"] = state ? state.connectionStatus : undefined;
            inputs["deleteRequested"] = state ? state.deleteRequested : undefined;
            inputs["errorMessage"] = state ? state.errorMessage : undefined;
            inputs["interfaceEndpointId"] = state ? state.interfaceEndpointId : undefined;
            inputs["privateLinkId"] = state ? state.privateLinkId : undefined;
            inputs["projectId"] = state ? state.projectId : undefined;
        } else {
            const args = argsOrState as PrivateEndpointInterfaceLinkArgs | undefined;
            if (!args || args.interfaceEndpointId === undefined) {
                throw new Error("Missing required property 'interfaceEndpointId'");
            }
            if (!args || args.privateLinkId === undefined) {
                throw new Error("Missing required property 'privateLinkId'");
            }
            if (!args || args.projectId === undefined) {
                throw new Error("Missing required property 'projectId'");
            }
            inputs["interfaceEndpointId"] = args ? args.interfaceEndpointId : undefined;
            inputs["privateLinkId"] = args ? args.privateLinkId : undefined;
            inputs["projectId"] = args ? args.projectId : undefined;
            inputs["connectionStatus"] = undefined /*out*/;
            inputs["deleteRequested"] = undefined /*out*/;
            inputs["errorMessage"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(PrivateEndpointInterfaceLink.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering PrivateEndpointInterfaceLink resources.
 */
export interface PrivateEndpointInterfaceLinkState {
    /**
     * Status of the interface endpoint.
     * Returns one of the following values:
     */
    readonly connectionStatus?: pulumi.Input<string>;
    /**
     * Indicates if Atlas received a request to remove the interface endpoint from the private endpoint connection.
     */
    readonly deleteRequested?: pulumi.Input<boolean>;
    /**
     * Error message pertaining to the interface endpoint. Returns null if there are no errors.
     */
    readonly errorMessage?: pulumi.Input<string>;
    /**
     * Unique identifier of the interface endpoint you created in your VPC with the AWS resource.
     */
    readonly interfaceEndpointId?: pulumi.Input<string>;
    /**
     * Unique identifier of the AWS PrivateLink connection which is created by `mongodbatlas..PrivateEndpoint` resource.
     */
    readonly privateLinkId?: pulumi.Input<string>;
    /**
     * Unique identifier for the project.
     */
    readonly projectId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a PrivateEndpointInterfaceLink resource.
 */
export interface PrivateEndpointInterfaceLinkArgs {
    /**
     * Unique identifier of the interface endpoint you created in your VPC with the AWS resource.
     */
    readonly interfaceEndpointId: pulumi.Input<string>;
    /**
     * Unique identifier of the AWS PrivateLink connection which is created by `mongodbatlas..PrivateEndpoint` resource.
     */
    readonly privateLinkId: pulumi.Input<string>;
    /**
     * Unique identifier for the project.
     */
    readonly projectId: pulumi.Input<string>;
}