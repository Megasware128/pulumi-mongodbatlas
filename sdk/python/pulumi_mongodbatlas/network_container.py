# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from . import _utilities, _tables

__all__ = ['NetworkContainer']


class NetworkContainer(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 atlas_cidr_block: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 provider_name: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 region_name: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        `NetworkContainer` provides a Network Peering Container resource. The resource lets you create, edit and delete network peering containers. The resource requires your Project ID.  Each cloud provider requires slightly different attributes so read the argument reference carefully.

         Network peering container is a general term used to describe any cloud providers' VPC/VNet concept.  Containers only need to be created if the peering connection to the cloud provider will be created before the first cluster that requires the container.  If the cluster has been/will be created first Atlas automatically creates the required container per the "containers per cloud provider" information that follows (in this case you can obtain the container id from the cluster resource attribute `container_id`).

        The following is the maximum number of Network Peering containers per cloud provider:
        <br> &#8226;  GCP -  One container per project.
        <br> &#8226;  AWS and Azure - One container per cloud provider region.

        > **NOTE:** Groups and projects are synonymous terms. You may find **group_id** in the official documentation.

        ## Example Usage
        ### Example with AWS

        ```python
        import pulumi
        import pulumi_mongodbatlas as mongodbatlas

        test = mongodbatlas.NetworkContainer("test",
            atlas_cidr_block="10.8.0.0/21",
            project_id="<YOUR-PROJECT-ID>",
            provider_name="AWS",
            region_name="US_EAST_1")
        ```
        ### Example with GCP

        ```python
        import pulumi
        import pulumi_mongodbatlas as mongodbatlas

        test = mongodbatlas.NetworkContainer("test",
            atlas_cidr_block="10.8.0.0/21",
            project_id="<YOUR-PROJECT-ID>",
            provider_name="GCP")
        ```
        ### Example with Azure

        ```python
        import pulumi
        import pulumi_mongodbatlas as mongodbatlas

        test = mongodbatlas.NetworkContainer("test",
            atlas_cidr_block="10.8.0.0/21",
            project_id="<YOUR-PROJECT-ID>",
            provider_name="AZURE",
            region="US_EAST_2")
        ```

        ## Import

        Clusters can be imported using project ID and network peering container id, in the format `PROJECTID-CONTAINER-ID`, e.g.

        ```sh
         $ pulumi import mongodbatlas:index/networkContainer:NetworkContainer my_container 1112222b3bf99403840e8934-5cbf563d87d9d67253be590a
        ```

         See detailed information for arguments and attributes[MongoDB API Network Peering Container](https://docs.atlas.mongodb.com/reference/api/vpc-create-container/)

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] atlas_cidr_block: CIDR block that Atlas uses for the Network Peering containers in your project.  Atlas uses the specified CIDR block for all other Network Peering connections created in the project. The Atlas CIDR block must be at least a /24 and at most a /21 in one of the following [private networks](https://tools.ietf.org/html/rfc1918.html#section-3):
               * Lower bound: 10.0.0.0 -	Upper bound: 10.255.255.255 -	Prefix: 10/8
               * Lower bound: 172.16.0.0 -	Upper bound:172.31.255.255 -	Prefix:	172.16/12
               * Lower bound: 192.168.0.0 -	Upper bound:192.168.255.255 -	Prefix:	192.168/16
        :param pulumi.Input[str] project_id: Unique identifier for the Atlas project for this Network Peering Container.
        :param pulumi.Input[str] provider_name: Cloud provider for this Network Peering connection.  Accepted values are GCP, AWS, AZURE. If omitted, Atlas sets this parameter to AWS.
        :param pulumi.Input[str] region: Atlas region where the container resides, see the reference list for Atlas Azure region names [Azure](https://docs.atlas.mongodb.com/reference/microsoft-azure/).
        :param pulumi.Input[str] region_name: The Atlas AWS region name for where this container will exist, see the reference list for Atlas AWS region names [AWS](https://docs.atlas.mongodb.com/reference/amazon-aws/).
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            if atlas_cidr_block is None:
                raise TypeError("Missing required property 'atlas_cidr_block'")
            __props__['atlas_cidr_block'] = atlas_cidr_block
            if project_id is None:
                raise TypeError("Missing required property 'project_id'")
            __props__['project_id'] = project_id
            __props__['provider_name'] = provider_name
            __props__['region'] = region
            __props__['region_name'] = region_name
            __props__['azure_subscription_id'] = None
            __props__['container_id'] = None
            __props__['gcp_project_id'] = None
            __props__['network_name'] = None
            __props__['provisioned'] = None
            __props__['vnet_name'] = None
            __props__['vpc_id'] = None
        super(NetworkContainer, __self__).__init__(
            'mongodbatlas:index/networkContainer:NetworkContainer',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            atlas_cidr_block: Optional[pulumi.Input[str]] = None,
            azure_subscription_id: Optional[pulumi.Input[str]] = None,
            container_id: Optional[pulumi.Input[str]] = None,
            gcp_project_id: Optional[pulumi.Input[str]] = None,
            network_name: Optional[pulumi.Input[str]] = None,
            project_id: Optional[pulumi.Input[str]] = None,
            provider_name: Optional[pulumi.Input[str]] = None,
            provisioned: Optional[pulumi.Input[bool]] = None,
            region: Optional[pulumi.Input[str]] = None,
            region_name: Optional[pulumi.Input[str]] = None,
            vnet_name: Optional[pulumi.Input[str]] = None,
            vpc_id: Optional[pulumi.Input[str]] = None) -> 'NetworkContainer':
        """
        Get an existing NetworkContainer resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] atlas_cidr_block: CIDR block that Atlas uses for the Network Peering containers in your project.  Atlas uses the specified CIDR block for all other Network Peering connections created in the project. The Atlas CIDR block must be at least a /24 and at most a /21 in one of the following [private networks](https://tools.ietf.org/html/rfc1918.html#section-3):
               * Lower bound: 10.0.0.0 -	Upper bound: 10.255.255.255 -	Prefix: 10/8
               * Lower bound: 172.16.0.0 -	Upper bound:172.31.255.255 -	Prefix:	172.16/12
               * Lower bound: 192.168.0.0 -	Upper bound:192.168.255.255 -	Prefix:	192.168/16
        :param pulumi.Input[str] azure_subscription_id: Unique identifier of the Azure subscription in which the VNet resides.
               * * `vnet_name` - 	The name of the Azure VNet. Returns null. This value is populated once you create a new network peering connection with the network peering resource.
        :param pulumi.Input[str] container_id: The Network Peering Container ID.
        :param pulumi.Input[str] gcp_project_id: Unique identifier of the GCP project in which the network peer resides. Returns null. This value is populated once you create a new network peering connection with the network peering resource.
        :param pulumi.Input[str] network_name: Unique identifier of the Network Peering connection in the Atlas project. Returns null. This value is populated once you create a new network peering connection with the network peering resource.
               **AZURE ONLY:**
        :param pulumi.Input[str] project_id: Unique identifier for the Atlas project for this Network Peering Container.
        :param pulumi.Input[str] provider_name: Cloud provider for this Network Peering connection.  Accepted values are GCP, AWS, AZURE. If omitted, Atlas sets this parameter to AWS.
        :param pulumi.Input[bool] provisioned: Indicates whether the project has Network Peering connections deployed in the container.
               **AWS ONLY:**
        :param pulumi.Input[str] region: Atlas region where the container resides, see the reference list for Atlas Azure region names [Azure](https://docs.atlas.mongodb.com/reference/microsoft-azure/).
        :param pulumi.Input[str] region_name: The Atlas AWS region name for where this container will exist, see the reference list for Atlas AWS region names [AWS](https://docs.atlas.mongodb.com/reference/amazon-aws/).
        :param pulumi.Input[str] vpc_id: Unique identifier of Atlas' AWS VPC.
               **CGP ONLY:**
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["atlas_cidr_block"] = atlas_cidr_block
        __props__["azure_subscription_id"] = azure_subscription_id
        __props__["container_id"] = container_id
        __props__["gcp_project_id"] = gcp_project_id
        __props__["network_name"] = network_name
        __props__["project_id"] = project_id
        __props__["provider_name"] = provider_name
        __props__["provisioned"] = provisioned
        __props__["region"] = region
        __props__["region_name"] = region_name
        __props__["vnet_name"] = vnet_name
        __props__["vpc_id"] = vpc_id
        return NetworkContainer(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="atlasCidrBlock")
    def atlas_cidr_block(self) -> pulumi.Output[str]:
        """
        CIDR block that Atlas uses for the Network Peering containers in your project.  Atlas uses the specified CIDR block for all other Network Peering connections created in the project. The Atlas CIDR block must be at least a /24 and at most a /21 in one of the following [private networks](https://tools.ietf.org/html/rfc1918.html#section-3):
        * Lower bound: 10.0.0.0 -	Upper bound: 10.255.255.255 -	Prefix: 10/8
        * Lower bound: 172.16.0.0 -	Upper bound:172.31.255.255 -	Prefix:	172.16/12
        * Lower bound: 192.168.0.0 -	Upper bound:192.168.255.255 -	Prefix:	192.168/16
        """
        return pulumi.get(self, "atlas_cidr_block")

    @property
    @pulumi.getter(name="azureSubscriptionId")
    def azure_subscription_id(self) -> pulumi.Output[str]:
        """
        Unique identifier of the Azure subscription in which the VNet resides.
        * * `vnet_name` - 	The name of the Azure VNet. Returns null. This value is populated once you create a new network peering connection with the network peering resource.
        """
        return pulumi.get(self, "azure_subscription_id")

    @property
    @pulumi.getter(name="containerId")
    def container_id(self) -> pulumi.Output[str]:
        """
        The Network Peering Container ID.
        """
        return pulumi.get(self, "container_id")

    @property
    @pulumi.getter(name="gcpProjectId")
    def gcp_project_id(self) -> pulumi.Output[str]:
        """
        Unique identifier of the GCP project in which the network peer resides. Returns null. This value is populated once you create a new network peering connection with the network peering resource.
        """
        return pulumi.get(self, "gcp_project_id")

    @property
    @pulumi.getter(name="networkName")
    def network_name(self) -> pulumi.Output[str]:
        """
        Unique identifier of the Network Peering connection in the Atlas project. Returns null. This value is populated once you create a new network peering connection with the network peering resource.
        **AZURE ONLY:**
        """
        return pulumi.get(self, "network_name")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Output[str]:
        """
        Unique identifier for the Atlas project for this Network Peering Container.
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter(name="providerName")
    def provider_name(self) -> pulumi.Output[Optional[str]]:
        """
        Cloud provider for this Network Peering connection.  Accepted values are GCP, AWS, AZURE. If omitted, Atlas sets this parameter to AWS.
        """
        return pulumi.get(self, "provider_name")

    @property
    @pulumi.getter
    def provisioned(self) -> pulumi.Output[bool]:
        """
        Indicates whether the project has Network Peering connections deployed in the container.
        **AWS ONLY:**
        """
        return pulumi.get(self, "provisioned")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        Atlas region where the container resides, see the reference list for Atlas Azure region names [Azure](https://docs.atlas.mongodb.com/reference/microsoft-azure/).
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="regionName")
    def region_name(self) -> pulumi.Output[str]:
        """
        The Atlas AWS region name for where this container will exist, see the reference list for Atlas AWS region names [AWS](https://docs.atlas.mongodb.com/reference/amazon-aws/).
        """
        return pulumi.get(self, "region_name")

    @property
    @pulumi.getter(name="vnetName")
    def vnet_name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "vnet_name")

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> pulumi.Output[str]:
        """
        Unique identifier of Atlas' AWS VPC.
        **CGP ONLY:**
        """
        return pulumi.get(self, "vpc_id")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

