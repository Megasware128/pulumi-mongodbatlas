// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * `mongodbatlas.NetworkContainer` provides a Network Peering Container resource. The resource lets you create, edit and delete network peering containers. The resource requires your Project ID.  Each cloud provider requires slightly different attributes so read the argument reference carefully.
 *
 *  Network peering container is a general term used to describe any cloud providers' VPC/VNet concept.  Containers only need to be created if the peering connection to the cloud provider will be created before the first cluster that requires the container.  If the cluster has been/will be created first Atlas automatically creates the required container per the "containers per cloud provider" information that follows (in this case you can obtain the container id from the cluster resource attribute `containerId`).
 *
 * The following is the maximum number of Network Peering containers per cloud provider:
 * <br> &#8226;  GCP -  One container per project.
 * <br> &#8226;  AWS and Azure - One container per cloud provider region.
 *
 * > **NOTE:** Groups and projects are synonymous terms. You may find **group_id** in the official documentation.
 *
 * ## Example Usage
 * ### Example with AWS
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as mongodbatlas from "@pulumi/mongodbatlas";
 *
 * const test = new mongodbatlas.NetworkContainer("test", {
 *     atlasCidrBlock: "10.8.0.0/21",
 *     projectId: "<YOUR-PROJECT-ID>",
 *     providerName: "AWS",
 *     regionName: "US_EAST_1",
 * });
 * ```
 * ### Example with GCP
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as mongodbatlas from "@pulumi/mongodbatlas";
 *
 * const test = new mongodbatlas.NetworkContainer("test", {
 *     atlasCidrBlock: "10.8.0.0/21",
 *     projectId: "<YOUR-PROJECT-ID>",
 *     providerName: "GCP",
 * });
 * ```
 * ### Example with Azure
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as mongodbatlas from "@pulumi/mongodbatlas";
 *
 * const test = new mongodbatlas.NetworkContainer("test", {
 *     atlasCidrBlock: "10.8.0.0/21",
 *     projectId: "<YOUR-PROJECT-ID>",
 *     providerName: "AZURE",
 *     region: "US_EAST_2",
 * });
 * ```
 *
 * ## Import
 *
 * Clusters can be imported using project ID and network peering container id, in the format `PROJECTID-CONTAINER-ID`, e.g.
 *
 * ```sh
 *  $ pulumi import mongodbatlas:index/networkContainer:NetworkContainer my_container 1112222b3bf99403840e8934-5cbf563d87d9d67253be590a
 * ```
 *
 *  See detailed information for arguments and attributes[MongoDB API Network Peering Container](https://docs.atlas.mongodb.com/reference/api/vpc-create-container/)
 */
export class NetworkContainer extends pulumi.CustomResource {
    /**
     * Get an existing NetworkContainer resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: NetworkContainerState, opts?: pulumi.CustomResourceOptions): NetworkContainer {
        return new NetworkContainer(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'mongodbatlas:index/networkContainer:NetworkContainer';

    /**
     * Returns true if the given object is an instance of NetworkContainer.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is NetworkContainer {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === NetworkContainer.__pulumiType;
    }

    /**
     * CIDR block that Atlas uses for the Network Peering containers in your project.  Atlas uses the specified CIDR block for all other Network Peering connections created in the project. The Atlas CIDR block must be at least a /24 and at most a /21 in one of the following [private networks](https://tools.ietf.org/html/rfc1918.html#section-3):
     * * Lower bound: 10.0.0.0 -	Upper bound: 10.255.255.255 -	Prefix: 10/8
     * * Lower bound: 172.16.0.0 -	Upper bound:172.31.255.255 -	Prefix:	172.16/12
     * * Lower bound: 192.168.0.0 -	Upper bound:192.168.255.255 -	Prefix:	192.168/16
     */
    public readonly atlasCidrBlock!: pulumi.Output<string>;
    /**
     * Unique identifier of the Azure subscription in which the VNet resides.
     * * * `vnetName` - 	The name of the Azure VNet. Returns null. This value is populated once you create a new network peering connection with the network peering resource.
     */
    public /*out*/ readonly azureSubscriptionId!: pulumi.Output<string>;
    /**
     * The Network Peering Container ID.
     */
    public /*out*/ readonly containerId!: pulumi.Output<string>;
    /**
     * Unique identifier of the GCP project in which the network peer resides. Returns null. This value is populated once you create a new network peering connection with the network peering resource.
     */
    public /*out*/ readonly gcpProjectId!: pulumi.Output<string>;
    /**
     * Unique identifier of the Network Peering connection in the Atlas project. Returns null. This value is populated once you create a new network peering connection with the network peering resource.
     * **AZURE ONLY:**
     */
    public /*out*/ readonly networkName!: pulumi.Output<string>;
    /**
     * Unique identifier for the Atlas project for this Network Peering Container.
     */
    public readonly projectId!: pulumi.Output<string>;
    /**
     * Cloud provider for this Network Peering connection.  Accepted values are GCP, AWS, AZURE. If omitted, Atlas sets this parameter to AWS.
     */
    public readonly providerName!: pulumi.Output<string | undefined>;
    /**
     * Indicates whether the project has Network Peering connections deployed in the container.
     * **AWS ONLY:**
     */
    public /*out*/ readonly provisioned!: pulumi.Output<boolean>;
    /**
     * Atlas region where the container resides, see the reference list for Atlas Azure region names [Azure](https://docs.atlas.mongodb.com/reference/microsoft-azure/).
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * The Atlas AWS region name for where this container will exist, see the reference list for Atlas AWS region names [AWS](https://docs.atlas.mongodb.com/reference/amazon-aws/).
     */
    public readonly regionName!: pulumi.Output<string>;
    public /*out*/ readonly vnetName!: pulumi.Output<string>;
    /**
     * Unique identifier of Atlas' AWS VPC.
     * **CGP ONLY:**
     */
    public /*out*/ readonly vpcId!: pulumi.Output<string>;

    /**
     * Create a NetworkContainer resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NetworkContainerArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: NetworkContainerArgs | NetworkContainerState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as NetworkContainerState | undefined;
            inputs["atlasCidrBlock"] = state ? state.atlasCidrBlock : undefined;
            inputs["azureSubscriptionId"] = state ? state.azureSubscriptionId : undefined;
            inputs["containerId"] = state ? state.containerId : undefined;
            inputs["gcpProjectId"] = state ? state.gcpProjectId : undefined;
            inputs["networkName"] = state ? state.networkName : undefined;
            inputs["projectId"] = state ? state.projectId : undefined;
            inputs["providerName"] = state ? state.providerName : undefined;
            inputs["provisioned"] = state ? state.provisioned : undefined;
            inputs["region"] = state ? state.region : undefined;
            inputs["regionName"] = state ? state.regionName : undefined;
            inputs["vnetName"] = state ? state.vnetName : undefined;
            inputs["vpcId"] = state ? state.vpcId : undefined;
        } else {
            const args = argsOrState as NetworkContainerArgs | undefined;
            if ((!args || args.atlasCidrBlock === undefined) && !opts.urn) {
                throw new Error("Missing required property 'atlasCidrBlock'");
            }
            if ((!args || args.projectId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'projectId'");
            }
            inputs["atlasCidrBlock"] = args ? args.atlasCidrBlock : undefined;
            inputs["projectId"] = args ? args.projectId : undefined;
            inputs["providerName"] = args ? args.providerName : undefined;
            inputs["region"] = args ? args.region : undefined;
            inputs["regionName"] = args ? args.regionName : undefined;
            inputs["azureSubscriptionId"] = undefined /*out*/;
            inputs["containerId"] = undefined /*out*/;
            inputs["gcpProjectId"] = undefined /*out*/;
            inputs["networkName"] = undefined /*out*/;
            inputs["provisioned"] = undefined /*out*/;
            inputs["vnetName"] = undefined /*out*/;
            inputs["vpcId"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(NetworkContainer.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering NetworkContainer resources.
 */
export interface NetworkContainerState {
    /**
     * CIDR block that Atlas uses for the Network Peering containers in your project.  Atlas uses the specified CIDR block for all other Network Peering connections created in the project. The Atlas CIDR block must be at least a /24 and at most a /21 in one of the following [private networks](https://tools.ietf.org/html/rfc1918.html#section-3):
     * * Lower bound: 10.0.0.0 -	Upper bound: 10.255.255.255 -	Prefix: 10/8
     * * Lower bound: 172.16.0.0 -	Upper bound:172.31.255.255 -	Prefix:	172.16/12
     * * Lower bound: 192.168.0.0 -	Upper bound:192.168.255.255 -	Prefix:	192.168/16
     */
    readonly atlasCidrBlock?: pulumi.Input<string>;
    /**
     * Unique identifier of the Azure subscription in which the VNet resides.
     * * * `vnetName` - 	The name of the Azure VNet. Returns null. This value is populated once you create a new network peering connection with the network peering resource.
     */
    readonly azureSubscriptionId?: pulumi.Input<string>;
    /**
     * The Network Peering Container ID.
     */
    readonly containerId?: pulumi.Input<string>;
    /**
     * Unique identifier of the GCP project in which the network peer resides. Returns null. This value is populated once you create a new network peering connection with the network peering resource.
     */
    readonly gcpProjectId?: pulumi.Input<string>;
    /**
     * Unique identifier of the Network Peering connection in the Atlas project. Returns null. This value is populated once you create a new network peering connection with the network peering resource.
     * **AZURE ONLY:**
     */
    readonly networkName?: pulumi.Input<string>;
    /**
     * Unique identifier for the Atlas project for this Network Peering Container.
     */
    readonly projectId?: pulumi.Input<string>;
    /**
     * Cloud provider for this Network Peering connection.  Accepted values are GCP, AWS, AZURE. If omitted, Atlas sets this parameter to AWS.
     */
    readonly providerName?: pulumi.Input<string>;
    /**
     * Indicates whether the project has Network Peering connections deployed in the container.
     * **AWS ONLY:**
     */
    readonly provisioned?: pulumi.Input<boolean>;
    /**
     * Atlas region where the container resides, see the reference list for Atlas Azure region names [Azure](https://docs.atlas.mongodb.com/reference/microsoft-azure/).
     */
    readonly region?: pulumi.Input<string>;
    /**
     * The Atlas AWS region name for where this container will exist, see the reference list for Atlas AWS region names [AWS](https://docs.atlas.mongodb.com/reference/amazon-aws/).
     */
    readonly regionName?: pulumi.Input<string>;
    readonly vnetName?: pulumi.Input<string>;
    /**
     * Unique identifier of Atlas' AWS VPC.
     * **CGP ONLY:**
     */
    readonly vpcId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a NetworkContainer resource.
 */
export interface NetworkContainerArgs {
    /**
     * CIDR block that Atlas uses for the Network Peering containers in your project.  Atlas uses the specified CIDR block for all other Network Peering connections created in the project. The Atlas CIDR block must be at least a /24 and at most a /21 in one of the following [private networks](https://tools.ietf.org/html/rfc1918.html#section-3):
     * * Lower bound: 10.0.0.0 -	Upper bound: 10.255.255.255 -	Prefix: 10/8
     * * Lower bound: 172.16.0.0 -	Upper bound:172.31.255.255 -	Prefix:	172.16/12
     * * Lower bound: 192.168.0.0 -	Upper bound:192.168.255.255 -	Prefix:	192.168/16
     */
    readonly atlasCidrBlock: pulumi.Input<string>;
    /**
     * Unique identifier for the Atlas project for this Network Peering Container.
     */
    readonly projectId: pulumi.Input<string>;
    /**
     * Cloud provider for this Network Peering connection.  Accepted values are GCP, AWS, AZURE. If omitted, Atlas sets this parameter to AWS.
     */
    readonly providerName?: pulumi.Input<string>;
    /**
     * Atlas region where the container resides, see the reference list for Atlas Azure region names [Azure](https://docs.atlas.mongodb.com/reference/microsoft-azure/).
     */
    readonly region?: pulumi.Input<string>;
    /**
     * The Atlas AWS region name for where this container will exist, see the reference list for Atlas AWS region names [AWS](https://docs.atlas.mongodb.com/reference/amazon-aws/).
     */
    readonly regionName?: pulumi.Input<string>;
}
