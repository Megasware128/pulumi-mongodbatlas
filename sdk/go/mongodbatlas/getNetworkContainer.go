// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package mongodbatlas

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// `.NetworkContainer` describes a Network Peering Container. The resource requires your Project ID and container ID.
//
// > **IMPORTANT:** This resource creates one Network Peering container into which Atlas can deploy Network Peering connections. An Atlas project can have a maximum of one container for each cloud provider. You must have either the Project Owner or Organization Owner role to successfully call this endpoint.
//
// > **NOTE:** Groups and projects are synonymous terms. You may find **group_id** in the official documentation.
func LookupNetworkContainer(ctx *pulumi.Context, args *LookupNetworkContainerArgs, opts ...pulumi.InvokeOption) (*LookupNetworkContainerResult, error) {
	var rv LookupNetworkContainerResult
	err := ctx.Invoke("mongodbatlas:index/getNetworkContainer:getNetworkContainer", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getNetworkContainer.
type LookupNetworkContainerArgs struct {
	// The Network Peering Container ID.
	ContainerId string `pulumi:"containerId"`
	// The unique ID for the project to create the database user.
	ProjectId string `pulumi:"projectId"`
}

// A collection of values returned by getNetworkContainer.
type LookupNetworkContainerResult struct {
	// CIDR block that Atlas uses for your clusters. Atlas uses the specified CIDR block for all other Network Peering connections created in the project. The Atlas CIDR block must be at least a /24 and at most a /21 in one of the following [private networks](https://tools.ietf.org/html/rfc1918.html#section-3).
	AtlasCidrBlock string `pulumi:"atlasCidrBlock"`
	// Unique identifer of the Azure subscription in which the VNet resides.
	AzureSubscriptionId string `pulumi:"azureSubscriptionId"`
	ContainerId         string `pulumi:"containerId"`
	// Unique identifier of the GCP project in which the Network Peering connection resides.
	GcpProjectId string `pulumi:"gcpProjectId"`
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Name of the Network Peering connection in the Atlas project.
	NetworkName string `pulumi:"networkName"`
	ProjectId   string `pulumi:"projectId"`
	// Cloud provider for this Network Peering connection. If omitted, Atlas sets this parameter to AWS.
	ProviderName string `pulumi:"providerName"`
	// Indicates whether the project has Network Peering connections deployed in the container.
	Provisioned bool `pulumi:"provisioned"`
	// The Atlas Azure region name for where this container will exist.
	Region string `pulumi:"region"`
	// The Atlas AWS region name for where this container will exist.
	RegionName string `pulumi:"regionName"`
	// The name of the Azure VNet. This value is null until you provision an Azure VNet in the container.
	VnetName string `pulumi:"vnetName"`
	// Unique identifier of the project’s VPC.
	VpcId string `pulumi:"vpcId"`
}