# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from . import utilities, tables

class GlobalClusterConfig(pulumi.CustomResource):
    cluster_name: pulumi.Output[str]
    custom_zone_mapping: pulumi.Output[dict]
    """
    A map of all custom zone mappings defined for the Global Cluster. Atlas automatically maps each location code to the closest geographical zone. Custom zone mappings allow administrators to override these automatic mappings. If your Global Cluster does not have any custom zone mappings, this document is empty.
    """
    custom_zone_mappings: pulumi.Output[list]
    """
    Each element in the list maps one ISO location code to a zone in your Global Cluster. See Custom Zone Mapping below for more details.

      * `location` (`str`) - The ISO location code to which you want to map a zone in your Global Cluster. You can find a list of all supported location codes [here](https://cloud.mongodb.com/static/atlas/country_iso_codes.txt).
      * `zone` (`str`) - The name of the zone in your Global Cluster that you want to map to location.
    """
    managed_namespaces: pulumi.Output[list]
    """
    Add a managed namespaces to a Global Cluster. For more information about managed namespaces, see [Global Clusters](https://docs.atlas.mongodb.com/reference/api/global-clusters/). See Managed Namespace below for more details.

      * `collection` (`str`) - The name of the collection associated with the managed namespace.
      * `customShardKey` (`str`) - The custom shard key for the collection. Global Clusters require a compound shard key consisting of a location field and a user-selected second key, the custom shard key.
      * `db` (`str`) - The name of the database containing the collection.
    """
    project_id: pulumi.Output[str]
    """
    The unique ID for the project to create the database user.
    * `cluster_name - (Required) The name of the Global Cluster.
    """
    def __init__(__self__, resource_name, opts=None, cluster_name=None, custom_zone_mappings=None, managed_namespaces=None, project_id=None, __props__=None, __name__=None, __opts__=None):
        """
        `.GlobalClusterConfig` provides a Global Cluster Configuration resource.


        > **NOTE:** Groups and projects are synonymous terms. You may find group_id in the official documentation.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] custom_zone_mappings: Each element in the list maps one ISO location code to a zone in your Global Cluster. See Custom Zone Mapping below for more details.
        :param pulumi.Input[list] managed_namespaces: Add a managed namespaces to a Global Cluster. For more information about managed namespaces, see [Global Clusters](https://docs.atlas.mongodb.com/reference/api/global-clusters/). See Managed Namespace below for more details.
        :param pulumi.Input[str] project_id: The unique ID for the project to create the database user.
               * `cluster_name - (Required) The name of the Global Cluster.

        The **custom_zone_mappings** object supports the following:

          * `location` (`pulumi.Input[str]`) - The ISO location code to which you want to map a zone in your Global Cluster. You can find a list of all supported location codes [here](https://cloud.mongodb.com/static/atlas/country_iso_codes.txt).
          * `zone` (`pulumi.Input[str]`) - The name of the zone in your Global Cluster that you want to map to location.

        The **managed_namespaces** object supports the following:

          * `collection` (`pulumi.Input[str]`) - The name of the collection associated with the managed namespace.
          * `customShardKey` (`pulumi.Input[str]`) - The custom shard key for the collection. Global Clusters require a compound shard key consisting of a location field and a user-selected second key, the custom shard key.
          * `db` (`pulumi.Input[str]`) - The name of the database containing the collection.
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
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            if cluster_name is None:
                raise TypeError("Missing required property 'cluster_name'")
            __props__['cluster_name'] = cluster_name
            __props__['custom_zone_mappings'] = custom_zone_mappings
            __props__['managed_namespaces'] = managed_namespaces
            if project_id is None:
                raise TypeError("Missing required property 'project_id'")
            __props__['project_id'] = project_id
            __props__['custom_zone_mapping'] = None
        super(GlobalClusterConfig, __self__).__init__(
            'mongodbatlas:index/globalClusterConfig:GlobalClusterConfig',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, cluster_name=None, custom_zone_mapping=None, custom_zone_mappings=None, managed_namespaces=None, project_id=None):
        """
        Get an existing GlobalClusterConfig resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] custom_zone_mapping: A map of all custom zone mappings defined for the Global Cluster. Atlas automatically maps each location code to the closest geographical zone. Custom zone mappings allow administrators to override these automatic mappings. If your Global Cluster does not have any custom zone mappings, this document is empty.
        :param pulumi.Input[list] custom_zone_mappings: Each element in the list maps one ISO location code to a zone in your Global Cluster. See Custom Zone Mapping below for more details.
        :param pulumi.Input[list] managed_namespaces: Add a managed namespaces to a Global Cluster. For more information about managed namespaces, see [Global Clusters](https://docs.atlas.mongodb.com/reference/api/global-clusters/). See Managed Namespace below for more details.
        :param pulumi.Input[str] project_id: The unique ID for the project to create the database user.
               * `cluster_name - (Required) The name of the Global Cluster.

        The **custom_zone_mappings** object supports the following:

          * `location` (`pulumi.Input[str]`) - The ISO location code to which you want to map a zone in your Global Cluster. You can find a list of all supported location codes [here](https://cloud.mongodb.com/static/atlas/country_iso_codes.txt).
          * `zone` (`pulumi.Input[str]`) - The name of the zone in your Global Cluster that you want to map to location.

        The **managed_namespaces** object supports the following:

          * `collection` (`pulumi.Input[str]`) - The name of the collection associated with the managed namespace.
          * `customShardKey` (`pulumi.Input[str]`) - The custom shard key for the collection. Global Clusters require a compound shard key consisting of a location field and a user-selected second key, the custom shard key.
          * `db` (`pulumi.Input[str]`) - The name of the database containing the collection.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["cluster_name"] = cluster_name
        __props__["custom_zone_mapping"] = custom_zone_mapping
        __props__["custom_zone_mappings"] = custom_zone_mappings
        __props__["managed_namespaces"] = managed_namespaces
        __props__["project_id"] = project_id
        return GlobalClusterConfig(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
