# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'GetPrivateEndpointRegionalModeResult',
    'AwaitableGetPrivateEndpointRegionalModeResult',
    'get_private_endpoint_regional_mode',
    'get_private_endpoint_regional_mode_output',
]

@pulumi.output_type
class GetPrivateEndpointRegionalModeResult:
    """
    A collection of values returned by getPrivateEndpointRegionalMode.
    """
    def __init__(__self__, enabled=None, id=None, project_id=None):
        if enabled and not isinstance(enabled, bool):
            raise TypeError("Expected argument 'enabled' to be a bool")
        pulumi.set(__self__, "enabled", enabled)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if project_id and not isinstance(project_id, str):
            raise TypeError("Expected argument 'project_id' to be a str")
        pulumi.set(__self__, "project_id", project_id)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[bool]:
        return pulumi.get(self, "enabled")

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


class AwaitableGetPrivateEndpointRegionalModeResult(GetPrivateEndpointRegionalModeResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetPrivateEndpointRegionalModeResult(
            enabled=self.enabled,
            id=self.id,
            project_id=self.project_id)


def get_private_endpoint_regional_mode(enabled: Optional[bool] = None,
                                       project_id: Optional[str] = None,
                                       opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetPrivateEndpointRegionalModeResult:
    """
    `private_endpoint_regional_mode` describe a Private Endpoint Regional Mode. This represents a Private Endpoint Regional Mode Connection that wants to retrieve settings of an Atlas project.

    > **NOTE:** Groups and projects are synonymous terms. You may find group_id in the official documentation.


    :param bool enabled: Flag that indicates whether the regionalized private endpoitn setting is enabled for the project.
    :param str project_id: Unique identifier for the project.
    """
    __args__ = dict()
    __args__['enabled'] = enabled
    __args__['projectId'] = project_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('mongodbatlas:index/getPrivateEndpointRegionalMode:getPrivateEndpointRegionalMode', __args__, opts=opts, typ=GetPrivateEndpointRegionalModeResult).value

    return AwaitableGetPrivateEndpointRegionalModeResult(
        enabled=__ret__.enabled,
        id=__ret__.id,
        project_id=__ret__.project_id)


@_utilities.lift_output_func(get_private_endpoint_regional_mode)
def get_private_endpoint_regional_mode_output(enabled: Optional[pulumi.Input[Optional[bool]]] = None,
                                              project_id: Optional[pulumi.Input[str]] = None,
                                              opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetPrivateEndpointRegionalModeResult]:
    """
    `private_endpoint_regional_mode` describe a Private Endpoint Regional Mode. This represents a Private Endpoint Regional Mode Connection that wants to retrieve settings of an Atlas project.

    > **NOTE:** Groups and projects are synonymous terms. You may find group_id in the official documentation.


    :param bool enabled: Flag that indicates whether the regionalized private endpoitn setting is enabled for the project.
    :param str project_id: Unique identifier for the project.
    """
    ...
