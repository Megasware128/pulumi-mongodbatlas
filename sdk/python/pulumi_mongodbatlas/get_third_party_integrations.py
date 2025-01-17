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
    'GetThirdPartyIntegrationsResult',
    'AwaitableGetThirdPartyIntegrationsResult',
    'get_third_party_integrations',
    'get_third_party_integrations_output',
]

@pulumi.output_type
class GetThirdPartyIntegrationsResult:
    """
    A collection of values returned by getThirdPartyIntegrations.
    """
    def __init__(__self__, id=None, project_id=None, results=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if project_id and not isinstance(project_id, str):
            raise TypeError("Expected argument 'project_id' to be a str")
        pulumi.set(__self__, "project_id", project_id)
        if results and not isinstance(results, list):
            raise TypeError("Expected argument 'results' to be a list")
        pulumi.set(__self__, "results", results)

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
        """
        (Required) ID of the Atlas project the Third-Party Service Integration belongs to.
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter
    def results(self) -> Sequence['outputs.GetThirdPartyIntegrationsResultResult']:
        """
        A list where each represents a Third-Party service integration.
        """
        return pulumi.get(self, "results")


class AwaitableGetThirdPartyIntegrationsResult(GetThirdPartyIntegrationsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetThirdPartyIntegrationsResult(
            id=self.id,
            project_id=self.project_id,
            results=self.results)


def get_third_party_integrations(project_id: Optional[str] = None,
                                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetThirdPartyIntegrationsResult:
    """
    `get_third_party_integrations` describe all Third-Party Integration Settings. This represents two Third-Party services `PAGER_DUTY` and `FLOWDOCK`
    applied across the project.

    > **NOTE:** Groups and projects are synonymous terms. You may find `groupId` in the official documentation.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_mongodbatlas as mongodbatlas

    test_pager_duty = mongodbatlas.ThirdPartyIntegration("testPagerDuty",
        project_id="<PROJECT-ID>",
        type="PAGER_DUTY",
        service_key="<PAGER-DUTY-SERVICE-KEY>")
    test_flowdock = mongodbatlas.ThirdPartyIntegration("testFlowdock",
        project_id="<PROJECT-ID>",
        type="FLOWDOCK",
        flow_name="<FLOW-NAME>",
        api_token="<API-TOKEN>",
        org_name="<ORG-NAME>")
    test = mongodbatlas.get_third_party_integrations_output(project_id=test_pager_duty.project_id)
    ```


    :param str project_id: The unique ID for the project to get all Third-Party service integrations
    """
    __args__ = dict()
    __args__['projectId'] = project_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('mongodbatlas:index/getThirdPartyIntegrations:getThirdPartyIntegrations', __args__, opts=opts, typ=GetThirdPartyIntegrationsResult).value

    return AwaitableGetThirdPartyIntegrationsResult(
        id=__ret__.id,
        project_id=__ret__.project_id,
        results=__ret__.results)


@_utilities.lift_output_func(get_third_party_integrations)
def get_third_party_integrations_output(project_id: Optional[pulumi.Input[str]] = None,
                                        opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetThirdPartyIntegrationsResult]:
    """
    `get_third_party_integrations` describe all Third-Party Integration Settings. This represents two Third-Party services `PAGER_DUTY` and `FLOWDOCK`
    applied across the project.

    > **NOTE:** Groups and projects are synonymous terms. You may find `groupId` in the official documentation.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_mongodbatlas as mongodbatlas

    test_pager_duty = mongodbatlas.ThirdPartyIntegration("testPagerDuty",
        project_id="<PROJECT-ID>",
        type="PAGER_DUTY",
        service_key="<PAGER-DUTY-SERVICE-KEY>")
    test_flowdock = mongodbatlas.ThirdPartyIntegration("testFlowdock",
        project_id="<PROJECT-ID>",
        type="FLOWDOCK",
        flow_name="<FLOW-NAME>",
        api_token="<API-TOKEN>",
        org_name="<ORG-NAME>")
    test = mongodbatlas.get_third_party_integrations_output(project_id=test_pager_duty.project_id)
    ```


    :param str project_id: The unique ID for the project to get all Third-Party service integrations
    """
    ...
