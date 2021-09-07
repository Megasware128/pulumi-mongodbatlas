# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs

__all__ = [
    'GetCloudProviderAccessSetupResult',
    'AwaitableGetCloudProviderAccessSetupResult',
    'get_cloud_provider_access_setup',
]

@pulumi.output_type
class GetCloudProviderAccessSetupResult:
    """
    A collection of values returned by getCloudProviderAccessSetup.
    """
    def __init__(__self__, aws=None, aws_configs=None, created_date=None, id=None, project_id=None, provider_name=None, role_id=None):
        if aws and not isinstance(aws, dict):
            raise TypeError("Expected argument 'aws' to be a dict")
        pulumi.set(__self__, "aws", aws)
        if aws_configs and not isinstance(aws_configs, list):
            raise TypeError("Expected argument 'aws_configs' to be a list")
        pulumi.set(__self__, "aws_configs", aws_configs)
        if created_date and not isinstance(created_date, str):
            raise TypeError("Expected argument 'created_date' to be a str")
        pulumi.set(__self__, "created_date", created_date)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if project_id and not isinstance(project_id, str):
            raise TypeError("Expected argument 'project_id' to be a str")
        pulumi.set(__self__, "project_id", project_id)
        if provider_name and not isinstance(provider_name, str):
            raise TypeError("Expected argument 'provider_name' to be a str")
        pulumi.set(__self__, "provider_name", provider_name)
        if role_id and not isinstance(role_id, str):
            raise TypeError("Expected argument 'role_id' to be a str")
        pulumi.set(__self__, "role_id", role_id)

    @property
    @pulumi.getter
    def aws(self) -> Mapping[str, str]:
        """
        aws related role information
        """
        return pulumi.get(self, "aws")

    @property
    @pulumi.getter(name="awsConfigs")
    def aws_configs(self) -> Sequence['outputs.GetCloudProviderAccessSetupAwsConfigResult']:
        return pulumi.get(self, "aws_configs")

    @property
    @pulumi.getter(name="createdDate")
    def created_date(self) -> str:
        """
        Date on which this role was created.
        """
        return pulumi.get(self, "created_date")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> str:
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter(name="providerName")
    def provider_name(self) -> str:
        return pulumi.get(self, "provider_name")

    @property
    @pulumi.getter(name="roleId")
    def role_id(self) -> str:
        return pulumi.get(self, "role_id")


class AwaitableGetCloudProviderAccessSetupResult(GetCloudProviderAccessSetupResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetCloudProviderAccessSetupResult(
            aws=self.aws,
            aws_configs=self.aws_configs,
            created_date=self.created_date,
            id=self.id,
            project_id=self.project_id,
            provider_name=self.provider_name,
            role_id=self.role_id)


def get_cloud_provider_access_setup(project_id: Optional[str] = None,
                                    provider_name: Optional[str] = None,
                                    role_id: Optional[str] = None,
                                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetCloudProviderAccessSetupResult:
    """
    `CloudProviderAccess` allows you to get a single role for a provider access role setup, currently only AWS is supported.

    > **NOTE:** Groups and projects are synonymous terms. You may find `groupId` in the official documentation.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_mongodbatlas as mongodbatlas

    test_role = mongodbatlas.CloudProviderAccessSetup("testRole",
        project_id="<PROJECT-ID>",
        provider_name="AWS")
    single_setup = pulumi.Output.all(test_role.project_id, test_role.provider_name, test_role.role_id).apply(lambda project_id, provider_name, role_id: mongodbatlas.get_cloud_provider_access_setup(project_id=project_id,
        provider_name=provider_name,
        role_id=role_id))
    ```


    :param str project_id: The unique ID for the project to get all Cloud Provider Access
    :param str provider_name: cloud provider name, currently only AWS is supported
    :param str role_id: unique role id among all the aws roles provided by mongodb atlas
    """
    __args__ = dict()
    __args__['projectId'] = project_id
    __args__['providerName'] = provider_name
    __args__['roleId'] = role_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('mongodbatlas:index/getCloudProviderAccessSetup:getCloudProviderAccessSetup', __args__, opts=opts, typ=GetCloudProviderAccessSetupResult).value

    return AwaitableGetCloudProviderAccessSetupResult(
        aws=__ret__.aws,
        aws_configs=__ret__.aws_configs,
        created_date=__ret__.created_date,
        id=__ret__.id,
        project_id=__ret__.project_id,
        provider_name=__ret__.provider_name,
        role_id=__ret__.role_id)
