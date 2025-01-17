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

__all__ = [
    'GetPrivatelinkEndpointsServiceAdlResult',
    'AwaitableGetPrivatelinkEndpointsServiceAdlResult',
    'get_privatelink_endpoints_service_adl',
    'get_privatelink_endpoints_service_adl_output',
]

@pulumi.output_type
class GetPrivatelinkEndpointsServiceAdlResult:
    """
    A collection of values returned by getPrivatelinkEndpointsServiceAdl.
    """
    def __init__(__self__, id=None, links=None, project_id=None, results=None, total_count=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if links and not isinstance(links, list):
            raise TypeError("Expected argument 'links' to be a list")
        pulumi.set(__self__, "links", links)
        if project_id and not isinstance(project_id, str):
            raise TypeError("Expected argument 'project_id' to be a str")
        pulumi.set(__self__, "project_id", project_id)
        if results and not isinstance(results, list):
            raise TypeError("Expected argument 'results' to be a list")
        pulumi.set(__self__, "results", results)
        if total_count and not isinstance(total_count, int):
            raise TypeError("Expected argument 'total_count' to be a int")
        pulumi.set(__self__, "total_count", total_count)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def links(self) -> Sequence['outputs.GetPrivatelinkEndpointsServiceAdlLinkResult']:
        return pulumi.get(self, "links")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> str:
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter
    def results(self) -> Sequence['outputs.GetPrivatelinkEndpointsServiceAdlResultResult']:
        return pulumi.get(self, "results")

    @property
    @pulumi.getter(name="totalCount")
    def total_count(self) -> int:
        return pulumi.get(self, "total_count")


class AwaitableGetPrivatelinkEndpointsServiceAdlResult(GetPrivatelinkEndpointsServiceAdlResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetPrivatelinkEndpointsServiceAdlResult(
            id=self.id,
            links=self.links,
            project_id=self.project_id,
            results=self.results,
            total_count=self.total_count)


def get_privatelink_endpoints_service_adl(project_id: Optional[str] = None,
                                          opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetPrivatelinkEndpointsServiceAdlResult:
    """
    `privatelink_endpoints_service_adl` Describes the list of all Atlas Data Lake (ADL) and Online Archive PrivateLink endpoints resource.

    > **NOTE:** Groups and projects are synonymous terms. You may find group_id in the official documentation.

    ## Example Usage
    ### Basic
    ```python
    import pulumi
    import pulumi_mongodbatlas as mongodbatlas

    adl_test = mongodbatlas.PrivatelinkEndpointServiceAdl("adlTest",
        project_id="<PROJECT_ID>",
        endpoint_id="<ENDPOINT_ID>",
        comment="Comment for PrivateLink endpoint ADL",
        type="DATA_LAKE",
        provider_name="AWS")
    test = mongodbatlas.get_privatelink_endpoints_service_adl_output(project_id=adl_test.project_id)
    ```


    :param str project_id: The unique ID for the project.
    """
    __args__ = dict()
    __args__['projectId'] = project_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('mongodbatlas:index/getPrivatelinkEndpointsServiceAdl:getPrivatelinkEndpointsServiceAdl', __args__, opts=opts, typ=GetPrivatelinkEndpointsServiceAdlResult).value

    return AwaitableGetPrivatelinkEndpointsServiceAdlResult(
        id=__ret__.id,
        links=__ret__.links,
        project_id=__ret__.project_id,
        results=__ret__.results,
        total_count=__ret__.total_count)


@_utilities.lift_output_func(get_privatelink_endpoints_service_adl)
def get_privatelink_endpoints_service_adl_output(project_id: Optional[pulumi.Input[str]] = None,
                                                 opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetPrivatelinkEndpointsServiceAdlResult]:
    """
    `privatelink_endpoints_service_adl` Describes the list of all Atlas Data Lake (ADL) and Online Archive PrivateLink endpoints resource.

    > **NOTE:** Groups and projects are synonymous terms. You may find group_id in the official documentation.

    ## Example Usage
    ### Basic
    ```python
    import pulumi
    import pulumi_mongodbatlas as mongodbatlas

    adl_test = mongodbatlas.PrivatelinkEndpointServiceAdl("adlTest",
        project_id="<PROJECT_ID>",
        endpoint_id="<ENDPOINT_ID>",
        comment="Comment for PrivateLink endpoint ADL",
        type="DATA_LAKE",
        provider_name="AWS")
    test = mongodbatlas.get_privatelink_endpoints_service_adl_output(project_id=adl_test.project_id)
    ```


    :param str project_id: The unique ID for the project.
    """
    ...
