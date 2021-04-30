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
    'GetLdapVerifyResult',
    'AwaitableGetLdapVerifyResult',
    'get_ldap_verify',
]

@pulumi.output_type
class GetLdapVerifyResult:
    """
    A collection of values returned by getLdapVerify.
    """
    def __init__(__self__, bind_username=None, hostname=None, id=None, links=None, port=None, project_id=None, request_id=None, status=None, validations=None):
        if bind_username and not isinstance(bind_username, str):
            raise TypeError("Expected argument 'bind_username' to be a str")
        pulumi.set(__self__, "bind_username", bind_username)
        if hostname and not isinstance(hostname, str):
            raise TypeError("Expected argument 'hostname' to be a str")
        pulumi.set(__self__, "hostname", hostname)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if links and not isinstance(links, list):
            raise TypeError("Expected argument 'links' to be a list")
        pulumi.set(__self__, "links", links)
        if port and not isinstance(port, int):
            raise TypeError("Expected argument 'port' to be a int")
        pulumi.set(__self__, "port", port)
        if project_id and not isinstance(project_id, str):
            raise TypeError("Expected argument 'project_id' to be a str")
        pulumi.set(__self__, "project_id", project_id)
        if request_id and not isinstance(request_id, str):
            raise TypeError("Expected argument 'request_id' to be a str")
        pulumi.set(__self__, "request_id", request_id)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if validations and not isinstance(validations, list):
            raise TypeError("Expected argument 'validations' to be a list")
        pulumi.set(__self__, "validations", validations)

    @property
    @pulumi.getter(name="bindUsername")
    def bind_username(self) -> str:
        """
        The user DN that Atlas uses to connect to the LDAP server.
        """
        return pulumi.get(self, "bind_username")

    @property
    @pulumi.getter
    def hostname(self) -> str:
        """
        (Required) The hostname or IP address of the LDAP server.
        """
        return pulumi.get(self, "hostname")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def links(self) -> Sequence['outputs.GetLdapVerifyLinkResult']:
        """
        One or more links to sub-resources. The relations in the URLs are explained in the Web Linking Specification.
        """
        return pulumi.get(self, "links")

    @property
    @pulumi.getter
    def port(self) -> int:
        """
        LDAP ConfigurationThe port to which the LDAP server listens for client connections.
        """
        return pulumi.get(self, "port")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> str:
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter(name="requestId")
    def request_id(self) -> str:
        """
        The unique identifier for the request to verify the LDAP over TLS/SSL configuration.
        """
        return pulumi.get(self, "request_id")

    @property
    @pulumi.getter
    def status(self) -> str:
        """
        The current status of the LDAP over TLS/SSL configuration.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def validations(self) -> Sequence['outputs.GetLdapVerifyValidationResult']:
        """
        Array of validation messages related to the verification of the provided LDAP over TLS/SSL configuration details.
        """
        return pulumi.get(self, "validations")


class AwaitableGetLdapVerifyResult(GetLdapVerifyResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetLdapVerifyResult(
            bind_username=self.bind_username,
            hostname=self.hostname,
            id=self.id,
            links=self.links,
            port=self.port,
            project_id=self.project_id,
            request_id=self.request_id,
            status=self.status,
            validations=self.validations)


def get_ldap_verify(project_id: Optional[str] = None,
                    request_id: Optional[str] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetLdapVerifyResult:
    """
    `LdapVerify` describes a LDAP Verify.

    > **NOTE:** Groups and projects are synonymous terms. You may find **group_id** in the official documentation.


    :param str project_id: Unique identifier for the Atlas project associated with the verification request.
    :param str request_id: Unique identifier of a request to verify an LDAP configuration.
    """
    __args__ = dict()
    __args__['projectId'] = project_id
    __args__['requestId'] = request_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('mongodbatlas:index/getLdapVerify:getLdapVerify', __args__, opts=opts, typ=GetLdapVerifyResult).value

    return AwaitableGetLdapVerifyResult(
        bind_username=__ret__.bind_username,
        hostname=__ret__.hostname,
        id=__ret__.id,
        links=__ret__.links,
        port=__ret__.port,
        project_id=__ret__.project_id,
        request_id=__ret__.request_id,
        status=__ret__.status,
        validations=__ret__.validations)
