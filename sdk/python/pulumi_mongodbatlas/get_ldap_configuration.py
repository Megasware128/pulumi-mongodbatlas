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
    'GetLdapConfigurationResult',
    'AwaitableGetLdapConfigurationResult',
    'get_ldap_configuration',
    'get_ldap_configuration_output',
]

@pulumi.output_type
class GetLdapConfigurationResult:
    """
    A collection of values returned by getLdapConfiguration.
    """
    def __init__(__self__, authentication_enabled=None, authorization_enabled=None, authz_query_template=None, bind_password=None, bind_username=None, ca_certificate=None, hostname=None, id=None, port=None, project_id=None, user_to_dn_mappings=None):
        if authentication_enabled and not isinstance(authentication_enabled, bool):
            raise TypeError("Expected argument 'authentication_enabled' to be a bool")
        pulumi.set(__self__, "authentication_enabled", authentication_enabled)
        if authorization_enabled and not isinstance(authorization_enabled, bool):
            raise TypeError("Expected argument 'authorization_enabled' to be a bool")
        pulumi.set(__self__, "authorization_enabled", authorization_enabled)
        if authz_query_template and not isinstance(authz_query_template, str):
            raise TypeError("Expected argument 'authz_query_template' to be a str")
        pulumi.set(__self__, "authz_query_template", authz_query_template)
        if bind_password and not isinstance(bind_password, str):
            raise TypeError("Expected argument 'bind_password' to be a str")
        pulumi.set(__self__, "bind_password", bind_password)
        if bind_username and not isinstance(bind_username, str):
            raise TypeError("Expected argument 'bind_username' to be a str")
        pulumi.set(__self__, "bind_username", bind_username)
        if ca_certificate and not isinstance(ca_certificate, str):
            raise TypeError("Expected argument 'ca_certificate' to be a str")
        pulumi.set(__self__, "ca_certificate", ca_certificate)
        if hostname and not isinstance(hostname, str):
            raise TypeError("Expected argument 'hostname' to be a str")
        pulumi.set(__self__, "hostname", hostname)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if port and not isinstance(port, int):
            raise TypeError("Expected argument 'port' to be a int")
        pulumi.set(__self__, "port", port)
        if project_id and not isinstance(project_id, str):
            raise TypeError("Expected argument 'project_id' to be a str")
        pulumi.set(__self__, "project_id", project_id)
        if user_to_dn_mappings and not isinstance(user_to_dn_mappings, list):
            raise TypeError("Expected argument 'user_to_dn_mappings' to be a list")
        pulumi.set(__self__, "user_to_dn_mappings", user_to_dn_mappings)

    @property
    @pulumi.getter(name="authenticationEnabled")
    def authentication_enabled(self) -> bool:
        """
        Specifies whether user authentication with LDAP is enabled.
        """
        return pulumi.get(self, "authentication_enabled")

    @property
    @pulumi.getter(name="authorizationEnabled")
    def authorization_enabled(self) -> bool:
        """
        Specifies whether user authorization with LDAP is enabled.
        """
        return pulumi.get(self, "authorization_enabled")

    @property
    @pulumi.getter(name="authzQueryTemplate")
    def authz_query_template(self) -> str:
        """
        An LDAP query template that Atlas executes to obtain the LDAP groups to which the authenticated user belongs.
        """
        return pulumi.get(self, "authz_query_template")

    @property
    @pulumi.getter(name="bindPassword")
    def bind_password(self) -> str:
        """
        The password used to authenticate the `bind_username`.
        """
        return pulumi.get(self, "bind_password")

    @property
    @pulumi.getter(name="bindUsername")
    def bind_username(self) -> str:
        """
        The user DN that Atlas uses to connect to the LDAP server.
        """
        return pulumi.get(self, "bind_username")

    @property
    @pulumi.getter(name="caCertificate")
    def ca_certificate(self) -> str:
        """
        CA certificate used to verify the identify of the LDAP server.
        """
        return pulumi.get(self, "ca_certificate")

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
    def port(self) -> int:
        """
        The port to which the LDAP server listens for client connections.
        """
        return pulumi.get(self, "port")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> str:
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter(name="userToDnMappings")
    def user_to_dn_mappings(self) -> Sequence['outputs.GetLdapConfigurationUserToDnMappingResult']:
        """
        Maps an LDAP username for authentication to an LDAP Distinguished Name (DN).
        * `user_to_dn_mapping.0.match` - A regular expression to match against a provided LDAP username.
        * `user_to_dn_mapping.0.substitution` - An LDAP Distinguished Name (DN) formatting template that converts the LDAP name matched by the `match` regular expression into an LDAP Distinguished Name.
        * `user_to_dn_mapping.0.ldap_query` - An LDAP query formatting template that inserts the LDAP name matched by the `match` regular expression into an LDAP query URI as specified by RFC 4515 and RFC 4516.
        """
        return pulumi.get(self, "user_to_dn_mappings")


class AwaitableGetLdapConfigurationResult(GetLdapConfigurationResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetLdapConfigurationResult(
            authentication_enabled=self.authentication_enabled,
            authorization_enabled=self.authorization_enabled,
            authz_query_template=self.authz_query_template,
            bind_password=self.bind_password,
            bind_username=self.bind_username,
            ca_certificate=self.ca_certificate,
            hostname=self.hostname,
            id=self.id,
            port=self.port,
            project_id=self.project_id,
            user_to_dn_mappings=self.user_to_dn_mappings)


def get_ldap_configuration(project_id: Optional[str] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetLdapConfigurationResult:
    """
    `LdapConfiguration` describes a LDAP Configuration.

    > **NOTE:** Groups and projects are synonymous terms. You may find **group_id** in the official documentation.


    :param str project_id: Identifier for the Atlas project associated with the LDAP over TLS/SSL configuration.
    """
    __args__ = dict()
    __args__['projectId'] = project_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('mongodbatlas:index/getLdapConfiguration:getLdapConfiguration', __args__, opts=opts, typ=GetLdapConfigurationResult).value

    return AwaitableGetLdapConfigurationResult(
        authentication_enabled=__ret__.authentication_enabled,
        authorization_enabled=__ret__.authorization_enabled,
        authz_query_template=__ret__.authz_query_template,
        bind_password=__ret__.bind_password,
        bind_username=__ret__.bind_username,
        ca_certificate=__ret__.ca_certificate,
        hostname=__ret__.hostname,
        id=__ret__.id,
        port=__ret__.port,
        project_id=__ret__.project_id,
        user_to_dn_mappings=__ret__.user_to_dn_mappings)


@_utilities.lift_output_func(get_ldap_configuration)
def get_ldap_configuration_output(project_id: Optional[pulumi.Input[str]] = None,
                                  opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetLdapConfigurationResult]:
    """
    `LdapConfiguration` describes a LDAP Configuration.

    > **NOTE:** Groups and projects are synonymous terms. You may find **group_id** in the official documentation.


    :param str project_id: Identifier for the Atlas project associated with the LDAP over TLS/SSL configuration.
    """
    ...
