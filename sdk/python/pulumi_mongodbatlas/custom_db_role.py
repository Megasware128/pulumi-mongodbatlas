# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from . import utilities, tables

class CustomDbRole(pulumi.CustomResource):
    actions: pulumi.Output[list]
    inherited_roles: pulumi.Output[list]
    project_id: pulumi.Output[str]
    """
    The unique ID for the project to create the database user.
    """
    role_name: pulumi.Output[str]
    """
    Name of the inherited role. This can either be another custom role or a built-in role.
    """
    def __init__(__self__, resource_name, opts=None, actions=None, inherited_roles=None, project_id=None, role_name=None, __props__=None, __name__=None, __opts__=None):
        """
        `.CustomDbRole` provides a Custom DB Role resource. The customDBRoles resource lets you retrieve, create and modify the custom MongoDB roles in your cluster. Use custom MongoDB roles to specify custom sets of actions which cannot be described by the built-in Atlas database user privileges.

        > **IMPORTANT** Custom roles cannot use actions unavailable to any cluster version in your project. Custom roles are defined at the project level, and must be compatible with each MongoDB version used by your project’s clusters. If you have a cluster in your project with MongoDB 3.4, you cannot create a custom role that uses actions introduced in MongoDB 3.6, such as useUUID.


        > **NOTE:** Groups and projects are synonymous terms. You may find group_id in the official documentation.

        ## Example Usage



        ```python
        import pulumi
        import pulumi_mongodbatlas as mongodbatlas

        test_role = mongodbatlas.CustomDbRole("testRole",
            actions=[
                {
                    "action": "UPDATE",
                    "resources": [{
                        "collectionName": "",
                        "databaseName": "anyDatabase",
                    }],
                },
                {
                    "action": "INSERT",
                    "resources": [{
                        "collectionName": "",
                        "databaseName": "anyDatabase",
                    }],
                },
                {
                    "action": "REMOVE",
                    "resources": [{
                        "collectionName": "",
                        "databaseName": "anyDatabase",
                    }],
                },
            ],
            project_id="<PROJECT-ID>",
            role_name="myCustomRole")
        ```

        ## Example Usage with inherited roles

        ```python
        import pulumi
        import pulumi_mongodbatlas as mongodbatlas

        inherited_role_one = mongodbatlas.CustomDbRole("inheritedRoleOne",
            actions=[{
                "action": "INSERT",
                "resources": [{
                    "collectionName": "",
                    "databaseName": "anyDatabase",
                }],
            }],
            project_id="<PROJECT-ID>",
            role_name="insertRole")
        inherited_role_two = mongodbatlas.CustomDbRole("inheritedRoleTwo",
            actions=[{
                "action": "SERVER_STATUS",
                "resources": [{
                    "cluster": True,
                }],
            }],
            project_id=inherited_role_one.project_id,
            role_name="statusServerRole")
        test_role = mongodbatlas.CustomDbRole("testRole",
            actions=[
                {
                    "action": "UPDATE",
                    "resources": [{
                        "collectionName": "",
                        "databaseName": "anyDatabase",
                    }],
                },
                {
                    "action": "REMOVE",
                    "resources": [{
                        "collectionName": "",
                        "databaseName": "anyDatabase",
                    }],
                },
            ],
            inherited_roles=[
                {
                    "databaseName": "admin",
                    "roleName": inherited_role_one.role_name,
                },
                {
                    "databaseName": "admin",
                    "roleName": inherited_role_two.role_name,
                },
            ],
            project_id=inherited_role_one.project_id,
            role_name="myCustomRole")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] project_id: The unique ID for the project to create the database user.
        :param pulumi.Input[str] role_name: Name of the inherited role. This can either be another custom role or a built-in role.

        The **actions** object supports the following:

          * `action` (`pulumi.Input[str]`) - Name of the privilege action. For a complete list of actions available in the Atlas API, see [Custom Role Actions](https://docs.atlas.mongodb.com/reference/api/custom-role-actions)
            > **Note**: The privilege actions available to the Custom Roles API resource represent a subset of the privilege actions available in the Atlas Custom Roles UI.
          * `resources` (`pulumi.Input[list]`) - Contains information on where the action is granted. Each object in the array either indicates a database and collection on which the action is granted, or indicates that the action is granted on the cluster resource.
            * `cluster` (`pulumi.Input[bool]`)
            * `collectionName` (`pulumi.Input[str]`)
            * `database_name` (`pulumi.Input[str]`) - Database on which the inherited role is granted.

        The **inherited_roles** object supports the following:

          * `database_name` (`pulumi.Input[str]`) - Database on which the inherited role is granted.
          * `role_name` (`pulumi.Input[str]`) - Name of the inherited role. This can either be another custom role or a built-in role.
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

            if actions is None:
                raise TypeError("Missing required property 'actions'")
            __props__['actions'] = actions
            __props__['inherited_roles'] = inherited_roles
            if project_id is None:
                raise TypeError("Missing required property 'project_id'")
            __props__['project_id'] = project_id
            if role_name is None:
                raise TypeError("Missing required property 'role_name'")
            __props__['role_name'] = role_name
        super(CustomDbRole, __self__).__init__(
            'mongodbatlas:index/customDbRole:CustomDbRole',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, actions=None, inherited_roles=None, project_id=None, role_name=None):
        """
        Get an existing CustomDbRole resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] project_id: The unique ID for the project to create the database user.
        :param pulumi.Input[str] role_name: Name of the inherited role. This can either be another custom role or a built-in role.

        The **actions** object supports the following:

          * `action` (`pulumi.Input[str]`) - Name of the privilege action. For a complete list of actions available in the Atlas API, see [Custom Role Actions](https://docs.atlas.mongodb.com/reference/api/custom-role-actions)
            > **Note**: The privilege actions available to the Custom Roles API resource represent a subset of the privilege actions available in the Atlas Custom Roles UI.
          * `resources` (`pulumi.Input[list]`) - Contains information on where the action is granted. Each object in the array either indicates a database and collection on which the action is granted, or indicates that the action is granted on the cluster resource.
            * `cluster` (`pulumi.Input[bool]`)
            * `collectionName` (`pulumi.Input[str]`)
            * `database_name` (`pulumi.Input[str]`) - Database on which the inherited role is granted.

        The **inherited_roles** object supports the following:

          * `database_name` (`pulumi.Input[str]`) - Database on which the inherited role is granted.
          * `role_name` (`pulumi.Input[str]`) - Name of the inherited role. This can either be another custom role or a built-in role.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["actions"] = actions
        __props__["inherited_roles"] = inherited_roles
        __props__["project_id"] = project_id
        __props__["role_name"] = role_name
        return CustomDbRole(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

