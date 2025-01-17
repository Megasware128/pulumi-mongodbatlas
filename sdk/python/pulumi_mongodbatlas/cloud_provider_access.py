# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs
from ._inputs import *

__all__ = ['CloudProviderAccessArgs', 'CloudProviderAccess']

@pulumi.input_type
class CloudProviderAccessArgs:
    def __init__(__self__, *,
                 project_id: pulumi.Input[str],
                 provider_name: pulumi.Input[str],
                 iam_assumed_role_arn: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a CloudProviderAccess resource.
        :param pulumi.Input[str] project_id: The unique ID for the project
        :param pulumi.Input[str] provider_name: The cloud provider for which to create a new role. Currently only AWS is supported.
        :param pulumi.Input[str] iam_assumed_role_arn: ARN of the IAM Role that Atlas assumes when accessing resources in your AWS account. This value is required after the creation (register of the role) as part of [Set Up Unified AWS Access](https://docs.atlas.mongodb.com/security/set-up-unified-aws-access/#set-up-unified-aws-access).
        """
        pulumi.set(__self__, "project_id", project_id)
        pulumi.set(__self__, "provider_name", provider_name)
        if iam_assumed_role_arn is not None:
            pulumi.set(__self__, "iam_assumed_role_arn", iam_assumed_role_arn)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Input[str]:
        """
        The unique ID for the project
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter(name="providerName")
    def provider_name(self) -> pulumi.Input[str]:
        """
        The cloud provider for which to create a new role. Currently only AWS is supported.
        """
        return pulumi.get(self, "provider_name")

    @provider_name.setter
    def provider_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "provider_name", value)

    @property
    @pulumi.getter(name="iamAssumedRoleArn")
    def iam_assumed_role_arn(self) -> Optional[pulumi.Input[str]]:
        """
        ARN of the IAM Role that Atlas assumes when accessing resources in your AWS account. This value is required after the creation (register of the role) as part of [Set Up Unified AWS Access](https://docs.atlas.mongodb.com/security/set-up-unified-aws-access/#set-up-unified-aws-access).
        """
        return pulumi.get(self, "iam_assumed_role_arn")

    @iam_assumed_role_arn.setter
    def iam_assumed_role_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "iam_assumed_role_arn", value)


@pulumi.input_type
class _CloudProviderAccessState:
    def __init__(__self__, *,
                 atlas_assumed_role_external_id: Optional[pulumi.Input[str]] = None,
                 atlas_aws_account_arn: Optional[pulumi.Input[str]] = None,
                 authorized_date: Optional[pulumi.Input[str]] = None,
                 created_date: Optional[pulumi.Input[str]] = None,
                 feature_usages: Optional[pulumi.Input[Sequence[pulumi.Input['CloudProviderAccessFeatureUsageArgs']]]] = None,
                 iam_assumed_role_arn: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 provider_name: Optional[pulumi.Input[str]] = None,
                 role_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering CloudProviderAccess resources.
        :param pulumi.Input[str] atlas_assumed_role_external_id: Unique external ID Atlas uses when assuming the IAM role in your AWS account.
        :param pulumi.Input[str] atlas_aws_account_arn: ARN associated with the Atlas AWS account used to assume IAM roles in your AWS account.
        :param pulumi.Input[str] authorized_date: Date on which this role was authorized.
        :param pulumi.Input[str] created_date: Date on which this role was created.
        :param pulumi.Input[Sequence[pulumi.Input['CloudProviderAccessFeatureUsageArgs']]] feature_usages: Atlas features this AWS IAM role is linked to.
        :param pulumi.Input[str] iam_assumed_role_arn: ARN of the IAM Role that Atlas assumes when accessing resources in your AWS account. This value is required after the creation (register of the role) as part of [Set Up Unified AWS Access](https://docs.atlas.mongodb.com/security/set-up-unified-aws-access/#set-up-unified-aws-access).
        :param pulumi.Input[str] project_id: The unique ID for the project
        :param pulumi.Input[str] provider_name: The cloud provider for which to create a new role. Currently only AWS is supported.
        :param pulumi.Input[str] role_id: Unique ID of this role returned by mongodb atlas api
        """
        if atlas_assumed_role_external_id is not None:
            pulumi.set(__self__, "atlas_assumed_role_external_id", atlas_assumed_role_external_id)
        if atlas_aws_account_arn is not None:
            pulumi.set(__self__, "atlas_aws_account_arn", atlas_aws_account_arn)
        if authorized_date is not None:
            pulumi.set(__self__, "authorized_date", authorized_date)
        if created_date is not None:
            pulumi.set(__self__, "created_date", created_date)
        if feature_usages is not None:
            pulumi.set(__self__, "feature_usages", feature_usages)
        if iam_assumed_role_arn is not None:
            pulumi.set(__self__, "iam_assumed_role_arn", iam_assumed_role_arn)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if provider_name is not None:
            pulumi.set(__self__, "provider_name", provider_name)
        if role_id is not None:
            pulumi.set(__self__, "role_id", role_id)

    @property
    @pulumi.getter(name="atlasAssumedRoleExternalId")
    def atlas_assumed_role_external_id(self) -> Optional[pulumi.Input[str]]:
        """
        Unique external ID Atlas uses when assuming the IAM role in your AWS account.
        """
        return pulumi.get(self, "atlas_assumed_role_external_id")

    @atlas_assumed_role_external_id.setter
    def atlas_assumed_role_external_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "atlas_assumed_role_external_id", value)

    @property
    @pulumi.getter(name="atlasAwsAccountArn")
    def atlas_aws_account_arn(self) -> Optional[pulumi.Input[str]]:
        """
        ARN associated with the Atlas AWS account used to assume IAM roles in your AWS account.
        """
        return pulumi.get(self, "atlas_aws_account_arn")

    @atlas_aws_account_arn.setter
    def atlas_aws_account_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "atlas_aws_account_arn", value)

    @property
    @pulumi.getter(name="authorizedDate")
    def authorized_date(self) -> Optional[pulumi.Input[str]]:
        """
        Date on which this role was authorized.
        """
        return pulumi.get(self, "authorized_date")

    @authorized_date.setter
    def authorized_date(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "authorized_date", value)

    @property
    @pulumi.getter(name="createdDate")
    def created_date(self) -> Optional[pulumi.Input[str]]:
        """
        Date on which this role was created.
        """
        return pulumi.get(self, "created_date")

    @created_date.setter
    def created_date(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "created_date", value)

    @property
    @pulumi.getter(name="featureUsages")
    def feature_usages(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['CloudProviderAccessFeatureUsageArgs']]]]:
        """
        Atlas features this AWS IAM role is linked to.
        """
        return pulumi.get(self, "feature_usages")

    @feature_usages.setter
    def feature_usages(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['CloudProviderAccessFeatureUsageArgs']]]]):
        pulumi.set(self, "feature_usages", value)

    @property
    @pulumi.getter(name="iamAssumedRoleArn")
    def iam_assumed_role_arn(self) -> Optional[pulumi.Input[str]]:
        """
        ARN of the IAM Role that Atlas assumes when accessing resources in your AWS account. This value is required after the creation (register of the role) as part of [Set Up Unified AWS Access](https://docs.atlas.mongodb.com/security/set-up-unified-aws-access/#set-up-unified-aws-access).
        """
        return pulumi.get(self, "iam_assumed_role_arn")

    @iam_assumed_role_arn.setter
    def iam_assumed_role_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "iam_assumed_role_arn", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[str]]:
        """
        The unique ID for the project
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter(name="providerName")
    def provider_name(self) -> Optional[pulumi.Input[str]]:
        """
        The cloud provider for which to create a new role. Currently only AWS is supported.
        """
        return pulumi.get(self, "provider_name")

    @provider_name.setter
    def provider_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "provider_name", value)

    @property
    @pulumi.getter(name="roleId")
    def role_id(self) -> Optional[pulumi.Input[str]]:
        """
        Unique ID of this role returned by mongodb atlas api
        """
        return pulumi.get(self, "role_id")

    @role_id.setter
    def role_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "role_id", value)


class CloudProviderAccess(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 iam_assumed_role_arn: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 provider_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        ## Import

        The Cloud Provider Access resource can be imported using project ID and the provider name and mongodbatlas role id, in the format `project_id`-`provider_name`-`role_id`, e.g.

        ```sh
         $ pulumi import mongodbatlas:index/cloudProviderAccess:CloudProviderAccess my_role 1112222b3bf99403840e8934-AWS-5fc17d476f7a33224f5b224e
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] iam_assumed_role_arn: ARN of the IAM Role that Atlas assumes when accessing resources in your AWS account. This value is required after the creation (register of the role) as part of [Set Up Unified AWS Access](https://docs.atlas.mongodb.com/security/set-up-unified-aws-access/#set-up-unified-aws-access).
        :param pulumi.Input[str] project_id: The unique ID for the project
        :param pulumi.Input[str] provider_name: The cloud provider for which to create a new role. Currently only AWS is supported.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: CloudProviderAccessArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Import

        The Cloud Provider Access resource can be imported using project ID and the provider name and mongodbatlas role id, in the format `project_id`-`provider_name`-`role_id`, e.g.

        ```sh
         $ pulumi import mongodbatlas:index/cloudProviderAccess:CloudProviderAccess my_role 1112222b3bf99403840e8934-AWS-5fc17d476f7a33224f5b224e
        ```

        :param str resource_name: The name of the resource.
        :param CloudProviderAccessArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(CloudProviderAccessArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 iam_assumed_role_arn: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 provider_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = CloudProviderAccessArgs.__new__(CloudProviderAccessArgs)

            __props__.__dict__["iam_assumed_role_arn"] = iam_assumed_role_arn
            if project_id is None and not opts.urn:
                raise TypeError("Missing required property 'project_id'")
            __props__.__dict__["project_id"] = project_id
            if provider_name is None and not opts.urn:
                raise TypeError("Missing required property 'provider_name'")
            __props__.__dict__["provider_name"] = provider_name
            __props__.__dict__["atlas_assumed_role_external_id"] = None
            __props__.__dict__["atlas_aws_account_arn"] = None
            __props__.__dict__["authorized_date"] = None
            __props__.__dict__["created_date"] = None
            __props__.__dict__["feature_usages"] = None
            __props__.__dict__["role_id"] = None
        super(CloudProviderAccess, __self__).__init__(
            'mongodbatlas:index/cloudProviderAccess:CloudProviderAccess',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            atlas_assumed_role_external_id: Optional[pulumi.Input[str]] = None,
            atlas_aws_account_arn: Optional[pulumi.Input[str]] = None,
            authorized_date: Optional[pulumi.Input[str]] = None,
            created_date: Optional[pulumi.Input[str]] = None,
            feature_usages: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['CloudProviderAccessFeatureUsageArgs']]]]] = None,
            iam_assumed_role_arn: Optional[pulumi.Input[str]] = None,
            project_id: Optional[pulumi.Input[str]] = None,
            provider_name: Optional[pulumi.Input[str]] = None,
            role_id: Optional[pulumi.Input[str]] = None) -> 'CloudProviderAccess':
        """
        Get an existing CloudProviderAccess resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] atlas_assumed_role_external_id: Unique external ID Atlas uses when assuming the IAM role in your AWS account.
        :param pulumi.Input[str] atlas_aws_account_arn: ARN associated with the Atlas AWS account used to assume IAM roles in your AWS account.
        :param pulumi.Input[str] authorized_date: Date on which this role was authorized.
        :param pulumi.Input[str] created_date: Date on which this role was created.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['CloudProviderAccessFeatureUsageArgs']]]] feature_usages: Atlas features this AWS IAM role is linked to.
        :param pulumi.Input[str] iam_assumed_role_arn: ARN of the IAM Role that Atlas assumes when accessing resources in your AWS account. This value is required after the creation (register of the role) as part of [Set Up Unified AWS Access](https://docs.atlas.mongodb.com/security/set-up-unified-aws-access/#set-up-unified-aws-access).
        :param pulumi.Input[str] project_id: The unique ID for the project
        :param pulumi.Input[str] provider_name: The cloud provider for which to create a new role. Currently only AWS is supported.
        :param pulumi.Input[str] role_id: Unique ID of this role returned by mongodb atlas api
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _CloudProviderAccessState.__new__(_CloudProviderAccessState)

        __props__.__dict__["atlas_assumed_role_external_id"] = atlas_assumed_role_external_id
        __props__.__dict__["atlas_aws_account_arn"] = atlas_aws_account_arn
        __props__.__dict__["authorized_date"] = authorized_date
        __props__.__dict__["created_date"] = created_date
        __props__.__dict__["feature_usages"] = feature_usages
        __props__.__dict__["iam_assumed_role_arn"] = iam_assumed_role_arn
        __props__.__dict__["project_id"] = project_id
        __props__.__dict__["provider_name"] = provider_name
        __props__.__dict__["role_id"] = role_id
        return CloudProviderAccess(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="atlasAssumedRoleExternalId")
    def atlas_assumed_role_external_id(self) -> pulumi.Output[str]:
        """
        Unique external ID Atlas uses when assuming the IAM role in your AWS account.
        """
        return pulumi.get(self, "atlas_assumed_role_external_id")

    @property
    @pulumi.getter(name="atlasAwsAccountArn")
    def atlas_aws_account_arn(self) -> pulumi.Output[str]:
        """
        ARN associated with the Atlas AWS account used to assume IAM roles in your AWS account.
        """
        return pulumi.get(self, "atlas_aws_account_arn")

    @property
    @pulumi.getter(name="authorizedDate")
    def authorized_date(self) -> pulumi.Output[str]:
        """
        Date on which this role was authorized.
        """
        return pulumi.get(self, "authorized_date")

    @property
    @pulumi.getter(name="createdDate")
    def created_date(self) -> pulumi.Output[str]:
        """
        Date on which this role was created.
        """
        return pulumi.get(self, "created_date")

    @property
    @pulumi.getter(name="featureUsages")
    def feature_usages(self) -> pulumi.Output[Sequence['outputs.CloudProviderAccessFeatureUsage']]:
        """
        Atlas features this AWS IAM role is linked to.
        """
        return pulumi.get(self, "feature_usages")

    @property
    @pulumi.getter(name="iamAssumedRoleArn")
    def iam_assumed_role_arn(self) -> pulumi.Output[Optional[str]]:
        """
        ARN of the IAM Role that Atlas assumes when accessing resources in your AWS account. This value is required after the creation (register of the role) as part of [Set Up Unified AWS Access](https://docs.atlas.mongodb.com/security/set-up-unified-aws-access/#set-up-unified-aws-access).
        """
        return pulumi.get(self, "iam_assumed_role_arn")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Output[str]:
        """
        The unique ID for the project
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter(name="providerName")
    def provider_name(self) -> pulumi.Output[str]:
        """
        The cloud provider for which to create a new role. Currently only AWS is supported.
        """
        return pulumi.get(self, "provider_name")

    @property
    @pulumi.getter(name="roleId")
    def role_id(self) -> pulumi.Output[str]:
        """
        Unique ID of this role returned by mongodb atlas api
        """
        return pulumi.get(self, "role_id")

