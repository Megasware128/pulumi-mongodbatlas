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

__all__ = ['LdapVerifyArgs', 'LdapVerify']

@pulumi.input_type
class LdapVerifyArgs:
    def __init__(__self__, *,
                 bind_password: pulumi.Input[str],
                 bind_username: pulumi.Input[str],
                 hostname: pulumi.Input[str],
                 port: pulumi.Input[int],
                 project_id: pulumi.Input[str],
                 authz_query_template: Optional[pulumi.Input[str]] = None,
                 ca_certificate: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a LdapVerify resource.
        :param pulumi.Input[str] bind_password: The password used to authenticate the `bind_username`.
        :param pulumi.Input[str] bind_username: The user DN that Atlas uses to connect to the LDAP server. Must be the full DN, such as `CN=BindUser,CN=Users,DC=myldapserver,DC=mycompany,DC=com`.
        :param pulumi.Input[str] hostname: The hostname or IP address of the LDAP server. The server must be visible to the internet or connected to your Atlas cluster with VPC Peering.
        :param pulumi.Input[int] port: The port to which the LDAP server listens for client connections. Default: `636`
        :param pulumi.Input[str] project_id: The unique ID for the project to configure LDAP.
        :param pulumi.Input[str] authz_query_template: An LDAP query template that Atlas executes to obtain the LDAP groups to which the authenticated user belongs. Used only for user authorization. Use the {USER} placeholder in the URL to substitute the authenticated username. The query is relative to the host specified with hostname. The formatting for the query must conform to RFC4515 and RFC 4516. If you do not provide a query template, Atlas attempts to use the default value: `{USER}?memberOf?base`.
        :param pulumi.Input[str] ca_certificate: CA certificate used to verify the identify of the LDAP server. Self-signed certificates are allowed.
        """
        pulumi.set(__self__, "bind_password", bind_password)
        pulumi.set(__self__, "bind_username", bind_username)
        pulumi.set(__self__, "hostname", hostname)
        pulumi.set(__self__, "port", port)
        pulumi.set(__self__, "project_id", project_id)
        if authz_query_template is not None:
            pulumi.set(__self__, "authz_query_template", authz_query_template)
        if ca_certificate is not None:
            pulumi.set(__self__, "ca_certificate", ca_certificate)

    @property
    @pulumi.getter(name="bindPassword")
    def bind_password(self) -> pulumi.Input[str]:
        """
        The password used to authenticate the `bind_username`.
        """
        return pulumi.get(self, "bind_password")

    @bind_password.setter
    def bind_password(self, value: pulumi.Input[str]):
        pulumi.set(self, "bind_password", value)

    @property
    @pulumi.getter(name="bindUsername")
    def bind_username(self) -> pulumi.Input[str]:
        """
        The user DN that Atlas uses to connect to the LDAP server. Must be the full DN, such as `CN=BindUser,CN=Users,DC=myldapserver,DC=mycompany,DC=com`.
        """
        return pulumi.get(self, "bind_username")

    @bind_username.setter
    def bind_username(self, value: pulumi.Input[str]):
        pulumi.set(self, "bind_username", value)

    @property
    @pulumi.getter
    def hostname(self) -> pulumi.Input[str]:
        """
        The hostname or IP address of the LDAP server. The server must be visible to the internet or connected to your Atlas cluster with VPC Peering.
        """
        return pulumi.get(self, "hostname")

    @hostname.setter
    def hostname(self, value: pulumi.Input[str]):
        pulumi.set(self, "hostname", value)

    @property
    @pulumi.getter
    def port(self) -> pulumi.Input[int]:
        """
        The port to which the LDAP server listens for client connections. Default: `636`
        """
        return pulumi.get(self, "port")

    @port.setter
    def port(self, value: pulumi.Input[int]):
        pulumi.set(self, "port", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Input[str]:
        """
        The unique ID for the project to configure LDAP.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter(name="authzQueryTemplate")
    def authz_query_template(self) -> Optional[pulumi.Input[str]]:
        """
        An LDAP query template that Atlas executes to obtain the LDAP groups to which the authenticated user belongs. Used only for user authorization. Use the {USER} placeholder in the URL to substitute the authenticated username. The query is relative to the host specified with hostname. The formatting for the query must conform to RFC4515 and RFC 4516. If you do not provide a query template, Atlas attempts to use the default value: `{USER}?memberOf?base`.
        """
        return pulumi.get(self, "authz_query_template")

    @authz_query_template.setter
    def authz_query_template(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "authz_query_template", value)

    @property
    @pulumi.getter(name="caCertificate")
    def ca_certificate(self) -> Optional[pulumi.Input[str]]:
        """
        CA certificate used to verify the identify of the LDAP server. Self-signed certificates are allowed.
        """
        return pulumi.get(self, "ca_certificate")

    @ca_certificate.setter
    def ca_certificate(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ca_certificate", value)


@pulumi.input_type
class _LdapVerifyState:
    def __init__(__self__, *,
                 authz_query_template: Optional[pulumi.Input[str]] = None,
                 bind_password: Optional[pulumi.Input[str]] = None,
                 bind_username: Optional[pulumi.Input[str]] = None,
                 ca_certificate: Optional[pulumi.Input[str]] = None,
                 hostname: Optional[pulumi.Input[str]] = None,
                 links: Optional[pulumi.Input[Sequence[pulumi.Input['LdapVerifyLinkArgs']]]] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 request_id: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 validations: Optional[pulumi.Input[Sequence[pulumi.Input['LdapVerifyValidationArgs']]]] = None):
        """
        Input properties used for looking up and filtering LdapVerify resources.
        :param pulumi.Input[str] authz_query_template: An LDAP query template that Atlas executes to obtain the LDAP groups to which the authenticated user belongs. Used only for user authorization. Use the {USER} placeholder in the URL to substitute the authenticated username. The query is relative to the host specified with hostname. The formatting for the query must conform to RFC4515 and RFC 4516. If you do not provide a query template, Atlas attempts to use the default value: `{USER}?memberOf?base`.
        :param pulumi.Input[str] bind_password: The password used to authenticate the `bind_username`.
        :param pulumi.Input[str] bind_username: The user DN that Atlas uses to connect to the LDAP server. Must be the full DN, such as `CN=BindUser,CN=Users,DC=myldapserver,DC=mycompany,DC=com`.
        :param pulumi.Input[str] ca_certificate: CA certificate used to verify the identify of the LDAP server. Self-signed certificates are allowed.
        :param pulumi.Input[str] hostname: The hostname or IP address of the LDAP server. The server must be visible to the internet or connected to your Atlas cluster with VPC Peering.
        :param pulumi.Input[Sequence[pulumi.Input['LdapVerifyLinkArgs']]] links: One or more links to sub-resources. The relations in the URLs are explained in the Web Linking Specification.
        :param pulumi.Input[int] port: The port to which the LDAP server listens for client connections. Default: `636`
        :param pulumi.Input[str] project_id: The unique ID for the project to configure LDAP.
        :param pulumi.Input[str] request_id: The unique identifier for the request to verify the LDAP over TLS/SSL configuration.
        :param pulumi.Input[str] status: The current status of the LDAP over TLS/SSL configuration. One of the following values: `PENDING`, `SUCCESS`, and `FAILED`.
        :param pulumi.Input[Sequence[pulumi.Input['LdapVerifyValidationArgs']]] validations: Array of validation messages related to the verification of the provided LDAP over TLS/SSL configuration details. The array contains a document for each test that Atlas runs. Atlas stops running tests after the first failure. The following return values can be seen here: [Values](https://docs.atlas.mongodb.com/reference/api/ldaps-configuration-request-verification)
        """
        if authz_query_template is not None:
            pulumi.set(__self__, "authz_query_template", authz_query_template)
        if bind_password is not None:
            pulumi.set(__self__, "bind_password", bind_password)
        if bind_username is not None:
            pulumi.set(__self__, "bind_username", bind_username)
        if ca_certificate is not None:
            pulumi.set(__self__, "ca_certificate", ca_certificate)
        if hostname is not None:
            pulumi.set(__self__, "hostname", hostname)
        if links is not None:
            pulumi.set(__self__, "links", links)
        if port is not None:
            pulumi.set(__self__, "port", port)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if request_id is not None:
            pulumi.set(__self__, "request_id", request_id)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if validations is not None:
            pulumi.set(__self__, "validations", validations)

    @property
    @pulumi.getter(name="authzQueryTemplate")
    def authz_query_template(self) -> Optional[pulumi.Input[str]]:
        """
        An LDAP query template that Atlas executes to obtain the LDAP groups to which the authenticated user belongs. Used only for user authorization. Use the {USER} placeholder in the URL to substitute the authenticated username. The query is relative to the host specified with hostname. The formatting for the query must conform to RFC4515 and RFC 4516. If you do not provide a query template, Atlas attempts to use the default value: `{USER}?memberOf?base`.
        """
        return pulumi.get(self, "authz_query_template")

    @authz_query_template.setter
    def authz_query_template(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "authz_query_template", value)

    @property
    @pulumi.getter(name="bindPassword")
    def bind_password(self) -> Optional[pulumi.Input[str]]:
        """
        The password used to authenticate the `bind_username`.
        """
        return pulumi.get(self, "bind_password")

    @bind_password.setter
    def bind_password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "bind_password", value)

    @property
    @pulumi.getter(name="bindUsername")
    def bind_username(self) -> Optional[pulumi.Input[str]]:
        """
        The user DN that Atlas uses to connect to the LDAP server. Must be the full DN, such as `CN=BindUser,CN=Users,DC=myldapserver,DC=mycompany,DC=com`.
        """
        return pulumi.get(self, "bind_username")

    @bind_username.setter
    def bind_username(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "bind_username", value)

    @property
    @pulumi.getter(name="caCertificate")
    def ca_certificate(self) -> Optional[pulumi.Input[str]]:
        """
        CA certificate used to verify the identify of the LDAP server. Self-signed certificates are allowed.
        """
        return pulumi.get(self, "ca_certificate")

    @ca_certificate.setter
    def ca_certificate(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ca_certificate", value)

    @property
    @pulumi.getter
    def hostname(self) -> Optional[pulumi.Input[str]]:
        """
        The hostname or IP address of the LDAP server. The server must be visible to the internet or connected to your Atlas cluster with VPC Peering.
        """
        return pulumi.get(self, "hostname")

    @hostname.setter
    def hostname(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "hostname", value)

    @property
    @pulumi.getter
    def links(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['LdapVerifyLinkArgs']]]]:
        """
        One or more links to sub-resources. The relations in the URLs are explained in the Web Linking Specification.
        """
        return pulumi.get(self, "links")

    @links.setter
    def links(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['LdapVerifyLinkArgs']]]]):
        pulumi.set(self, "links", value)

    @property
    @pulumi.getter
    def port(self) -> Optional[pulumi.Input[int]]:
        """
        The port to which the LDAP server listens for client connections. Default: `636`
        """
        return pulumi.get(self, "port")

    @port.setter
    def port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "port", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[str]]:
        """
        The unique ID for the project to configure LDAP.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter(name="requestId")
    def request_id(self) -> Optional[pulumi.Input[str]]:
        """
        The unique identifier for the request to verify the LDAP over TLS/SSL configuration.
        """
        return pulumi.get(self, "request_id")

    @request_id.setter
    def request_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "request_id", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        The current status of the LDAP over TLS/SSL configuration. One of the following values: `PENDING`, `SUCCESS`, and `FAILED`.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter
    def validations(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['LdapVerifyValidationArgs']]]]:
        """
        Array of validation messages related to the verification of the provided LDAP over TLS/SSL configuration details. The array contains a document for each test that Atlas runs. Atlas stops running tests after the first failure. The following return values can be seen here: [Values](https://docs.atlas.mongodb.com/reference/api/ldaps-configuration-request-verification)
        """
        return pulumi.get(self, "validations")

    @validations.setter
    def validations(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['LdapVerifyValidationArgs']]]]):
        pulumi.set(self, "validations", value)


class LdapVerify(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 authz_query_template: Optional[pulumi.Input[str]] = None,
                 bind_password: Optional[pulumi.Input[str]] = None,
                 bind_username: Optional[pulumi.Input[str]] = None,
                 ca_certificate: Optional[pulumi.Input[str]] = None,
                 hostname: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        `LdapVerify` provides an LDAP Verify resource. This allows a a verification of an LDAP configuration over TLS for an Atlas project. Atlas retains only the most recent request for each project.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_mongodbatlas as mongodbatlas

        test_project = mongodbatlas.Project("testProject", org_id="ORG ID")
        test_cluster = mongodbatlas.Cluster("testCluster",
            project_id=test_project.id,
            disk_size_gb=5,
            provider_name="AWS",
            provider_region_name="US_EAST_2",
            provider_instance_size_name="M10",
            cloud_backup=True)
        #enable cloud provider snapshots
        test_ldap_verify = mongodbatlas.LdapVerify("testLdapVerify",
            project_id=test_project.id,
            hostname="HOSTNAME",
            port=636,
            bind_username="USERNAME",
            bind_password="PASSWORD",
            opts=pulumi.ResourceOptions(depends_on=[test_cluster]))
        ```

        ## Import

        LDAP Configuration must be imported using project ID and request ID, e.g.

        ```sh
         $ pulumi import mongodbatlas:index/ldapVerify:LdapVerify test 5d09d6a59ccf6445652a444a-5d09d6a59ccf6445652a444a
        ```

         For more information see[MongoDB Atlas API Reference.](https://docs.atlas.mongodb.com/reference/api/ldaps-configuration-request-verification)

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] authz_query_template: An LDAP query template that Atlas executes to obtain the LDAP groups to which the authenticated user belongs. Used only for user authorization. Use the {USER} placeholder in the URL to substitute the authenticated username. The query is relative to the host specified with hostname. The formatting for the query must conform to RFC4515 and RFC 4516. If you do not provide a query template, Atlas attempts to use the default value: `{USER}?memberOf?base`.
        :param pulumi.Input[str] bind_password: The password used to authenticate the `bind_username`.
        :param pulumi.Input[str] bind_username: The user DN that Atlas uses to connect to the LDAP server. Must be the full DN, such as `CN=BindUser,CN=Users,DC=myldapserver,DC=mycompany,DC=com`.
        :param pulumi.Input[str] ca_certificate: CA certificate used to verify the identify of the LDAP server. Self-signed certificates are allowed.
        :param pulumi.Input[str] hostname: The hostname or IP address of the LDAP server. The server must be visible to the internet or connected to your Atlas cluster with VPC Peering.
        :param pulumi.Input[int] port: The port to which the LDAP server listens for client connections. Default: `636`
        :param pulumi.Input[str] project_id: The unique ID for the project to configure LDAP.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: LdapVerifyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        `LdapVerify` provides an LDAP Verify resource. This allows a a verification of an LDAP configuration over TLS for an Atlas project. Atlas retains only the most recent request for each project.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_mongodbatlas as mongodbatlas

        test_project = mongodbatlas.Project("testProject", org_id="ORG ID")
        test_cluster = mongodbatlas.Cluster("testCluster",
            project_id=test_project.id,
            disk_size_gb=5,
            provider_name="AWS",
            provider_region_name="US_EAST_2",
            provider_instance_size_name="M10",
            cloud_backup=True)
        #enable cloud provider snapshots
        test_ldap_verify = mongodbatlas.LdapVerify("testLdapVerify",
            project_id=test_project.id,
            hostname="HOSTNAME",
            port=636,
            bind_username="USERNAME",
            bind_password="PASSWORD",
            opts=pulumi.ResourceOptions(depends_on=[test_cluster]))
        ```

        ## Import

        LDAP Configuration must be imported using project ID and request ID, e.g.

        ```sh
         $ pulumi import mongodbatlas:index/ldapVerify:LdapVerify test 5d09d6a59ccf6445652a444a-5d09d6a59ccf6445652a444a
        ```

         For more information see[MongoDB Atlas API Reference.](https://docs.atlas.mongodb.com/reference/api/ldaps-configuration-request-verification)

        :param str resource_name: The name of the resource.
        :param LdapVerifyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(LdapVerifyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 authz_query_template: Optional[pulumi.Input[str]] = None,
                 bind_password: Optional[pulumi.Input[str]] = None,
                 bind_username: Optional[pulumi.Input[str]] = None,
                 ca_certificate: Optional[pulumi.Input[str]] = None,
                 hostname: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = LdapVerifyArgs.__new__(LdapVerifyArgs)

            __props__.__dict__["authz_query_template"] = authz_query_template
            if bind_password is None and not opts.urn:
                raise TypeError("Missing required property 'bind_password'")
            __props__.__dict__["bind_password"] = bind_password
            if bind_username is None and not opts.urn:
                raise TypeError("Missing required property 'bind_username'")
            __props__.__dict__["bind_username"] = bind_username
            __props__.__dict__["ca_certificate"] = ca_certificate
            if hostname is None and not opts.urn:
                raise TypeError("Missing required property 'hostname'")
            __props__.__dict__["hostname"] = hostname
            if port is None and not opts.urn:
                raise TypeError("Missing required property 'port'")
            __props__.__dict__["port"] = port
            if project_id is None and not opts.urn:
                raise TypeError("Missing required property 'project_id'")
            __props__.__dict__["project_id"] = project_id
            __props__.__dict__["links"] = None
            __props__.__dict__["request_id"] = None
            __props__.__dict__["status"] = None
            __props__.__dict__["validations"] = None
        super(LdapVerify, __self__).__init__(
            'mongodbatlas:index/ldapVerify:LdapVerify',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            authz_query_template: Optional[pulumi.Input[str]] = None,
            bind_password: Optional[pulumi.Input[str]] = None,
            bind_username: Optional[pulumi.Input[str]] = None,
            ca_certificate: Optional[pulumi.Input[str]] = None,
            hostname: Optional[pulumi.Input[str]] = None,
            links: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['LdapVerifyLinkArgs']]]]] = None,
            port: Optional[pulumi.Input[int]] = None,
            project_id: Optional[pulumi.Input[str]] = None,
            request_id: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None,
            validations: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['LdapVerifyValidationArgs']]]]] = None) -> 'LdapVerify':
        """
        Get an existing LdapVerify resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] authz_query_template: An LDAP query template that Atlas executes to obtain the LDAP groups to which the authenticated user belongs. Used only for user authorization. Use the {USER} placeholder in the URL to substitute the authenticated username. The query is relative to the host specified with hostname. The formatting for the query must conform to RFC4515 and RFC 4516. If you do not provide a query template, Atlas attempts to use the default value: `{USER}?memberOf?base`.
        :param pulumi.Input[str] bind_password: The password used to authenticate the `bind_username`.
        :param pulumi.Input[str] bind_username: The user DN that Atlas uses to connect to the LDAP server. Must be the full DN, such as `CN=BindUser,CN=Users,DC=myldapserver,DC=mycompany,DC=com`.
        :param pulumi.Input[str] ca_certificate: CA certificate used to verify the identify of the LDAP server. Self-signed certificates are allowed.
        :param pulumi.Input[str] hostname: The hostname or IP address of the LDAP server. The server must be visible to the internet or connected to your Atlas cluster with VPC Peering.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['LdapVerifyLinkArgs']]]] links: One or more links to sub-resources. The relations in the URLs are explained in the Web Linking Specification.
        :param pulumi.Input[int] port: The port to which the LDAP server listens for client connections. Default: `636`
        :param pulumi.Input[str] project_id: The unique ID for the project to configure LDAP.
        :param pulumi.Input[str] request_id: The unique identifier for the request to verify the LDAP over TLS/SSL configuration.
        :param pulumi.Input[str] status: The current status of the LDAP over TLS/SSL configuration. One of the following values: `PENDING`, `SUCCESS`, and `FAILED`.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['LdapVerifyValidationArgs']]]] validations: Array of validation messages related to the verification of the provided LDAP over TLS/SSL configuration details. The array contains a document for each test that Atlas runs. Atlas stops running tests after the first failure. The following return values can be seen here: [Values](https://docs.atlas.mongodb.com/reference/api/ldaps-configuration-request-verification)
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _LdapVerifyState.__new__(_LdapVerifyState)

        __props__.__dict__["authz_query_template"] = authz_query_template
        __props__.__dict__["bind_password"] = bind_password
        __props__.__dict__["bind_username"] = bind_username
        __props__.__dict__["ca_certificate"] = ca_certificate
        __props__.__dict__["hostname"] = hostname
        __props__.__dict__["links"] = links
        __props__.__dict__["port"] = port
        __props__.__dict__["project_id"] = project_id
        __props__.__dict__["request_id"] = request_id
        __props__.__dict__["status"] = status
        __props__.__dict__["validations"] = validations
        return LdapVerify(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="authzQueryTemplate")
    def authz_query_template(self) -> pulumi.Output[str]:
        """
        An LDAP query template that Atlas executes to obtain the LDAP groups to which the authenticated user belongs. Used only for user authorization. Use the {USER} placeholder in the URL to substitute the authenticated username. The query is relative to the host specified with hostname. The formatting for the query must conform to RFC4515 and RFC 4516. If you do not provide a query template, Atlas attempts to use the default value: `{USER}?memberOf?base`.
        """
        return pulumi.get(self, "authz_query_template")

    @property
    @pulumi.getter(name="bindPassword")
    def bind_password(self) -> pulumi.Output[str]:
        """
        The password used to authenticate the `bind_username`.
        """
        return pulumi.get(self, "bind_password")

    @property
    @pulumi.getter(name="bindUsername")
    def bind_username(self) -> pulumi.Output[str]:
        """
        The user DN that Atlas uses to connect to the LDAP server. Must be the full DN, such as `CN=BindUser,CN=Users,DC=myldapserver,DC=mycompany,DC=com`.
        """
        return pulumi.get(self, "bind_username")

    @property
    @pulumi.getter(name="caCertificate")
    def ca_certificate(self) -> pulumi.Output[str]:
        """
        CA certificate used to verify the identify of the LDAP server. Self-signed certificates are allowed.
        """
        return pulumi.get(self, "ca_certificate")

    @property
    @pulumi.getter
    def hostname(self) -> pulumi.Output[str]:
        """
        The hostname or IP address of the LDAP server. The server must be visible to the internet or connected to your Atlas cluster with VPC Peering.
        """
        return pulumi.get(self, "hostname")

    @property
    @pulumi.getter
    def links(self) -> pulumi.Output[Sequence['outputs.LdapVerifyLink']]:
        """
        One or more links to sub-resources. The relations in the URLs are explained in the Web Linking Specification.
        """
        return pulumi.get(self, "links")

    @property
    @pulumi.getter
    def port(self) -> pulumi.Output[int]:
        """
        The port to which the LDAP server listens for client connections. Default: `636`
        """
        return pulumi.get(self, "port")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Output[str]:
        """
        The unique ID for the project to configure LDAP.
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter(name="requestId")
    def request_id(self) -> pulumi.Output[str]:
        """
        The unique identifier for the request to verify the LDAP over TLS/SSL configuration.
        """
        return pulumi.get(self, "request_id")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        The current status of the LDAP over TLS/SSL configuration. One of the following values: `PENDING`, `SUCCESS`, and `FAILED`.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def validations(self) -> pulumi.Output[Sequence['outputs.LdapVerifyValidation']]:
        """
        Array of validation messages related to the verification of the provided LDAP over TLS/SSL configuration details. The array contains a document for each test that Atlas runs. Atlas stops running tests after the first failure. The following return values can be seen here: [Values](https://docs.atlas.mongodb.com/reference/api/ldaps-configuration-request-verification)
        """
        return pulumi.get(self, "validations")

