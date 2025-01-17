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
    'GetDatabaseUserResult',
    'AwaitableGetDatabaseUserResult',
    'get_database_user',
    'get_database_user_output',
]

@pulumi.output_type
class GetDatabaseUserResult:
    """
    A collection of values returned by getDatabaseUser.
    """
    def __init__(__self__, auth_database_name=None, aws_iam_type=None, database_name=None, id=None, labels=None, ldap_auth_type=None, project_id=None, roles=None, scopes=None, username=None, x509_type=None):
        if auth_database_name and not isinstance(auth_database_name, str):
            raise TypeError("Expected argument 'auth_database_name' to be a str")
        pulumi.set(__self__, "auth_database_name", auth_database_name)
        if aws_iam_type and not isinstance(aws_iam_type, str):
            raise TypeError("Expected argument 'aws_iam_type' to be a str")
        pulumi.set(__self__, "aws_iam_type", aws_iam_type)
        if database_name and not isinstance(database_name, str):
            raise TypeError("Expected argument 'database_name' to be a str")
        if database_name is not None:
            warnings.warn("""use auth_database_name instead""", DeprecationWarning)
            pulumi.log.warn("""database_name is deprecated: use auth_database_name instead""")

        pulumi.set(__self__, "database_name", database_name)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if labels and not isinstance(labels, list):
            raise TypeError("Expected argument 'labels' to be a list")
        pulumi.set(__self__, "labels", labels)
        if ldap_auth_type and not isinstance(ldap_auth_type, str):
            raise TypeError("Expected argument 'ldap_auth_type' to be a str")
        pulumi.set(__self__, "ldap_auth_type", ldap_auth_type)
        if project_id and not isinstance(project_id, str):
            raise TypeError("Expected argument 'project_id' to be a str")
        pulumi.set(__self__, "project_id", project_id)
        if roles and not isinstance(roles, list):
            raise TypeError("Expected argument 'roles' to be a list")
        pulumi.set(__self__, "roles", roles)
        if scopes and not isinstance(scopes, list):
            raise TypeError("Expected argument 'scopes' to be a list")
        pulumi.set(__self__, "scopes", scopes)
        if username and not isinstance(username, str):
            raise TypeError("Expected argument 'username' to be a str")
        pulumi.set(__self__, "username", username)
        if x509_type and not isinstance(x509_type, str):
            raise TypeError("Expected argument 'x509_type' to be a str")
        pulumi.set(__self__, "x509_type", x509_type)

    @property
    @pulumi.getter(name="authDatabaseName")
    def auth_database_name(self) -> Optional[str]:
        return pulumi.get(self, "auth_database_name")

    @property
    @pulumi.getter(name="awsIamType")
    def aws_iam_type(self) -> str:
        """
        The new database user authenticates with AWS IAM credentials. Default is `NONE`, `USER` means user has AWS IAM user credentials, `ROLE` - means user has credentials associated with an AWS IAM role.
        """
        return pulumi.get(self, "aws_iam_type")

    @property
    @pulumi.getter(name="databaseName")
    def database_name(self) -> Optional[str]:
        """
        Database on which the user has the specified role. A role on the `admin` database can include privileges that apply to the other databases.
        """
        return pulumi.get(self, "database_name")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def labels(self) -> Sequence['outputs.GetDatabaseUserLabelResult']:
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter(name="ldapAuthType")
    def ldap_auth_type(self) -> str:
        """
        Method by which the provided username is authenticated. Default is `NONE`. Other valid values are: `USER`, `GROUP`.
        """
        return pulumi.get(self, "ldap_auth_type")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> str:
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter
    def roles(self) -> Sequence['outputs.GetDatabaseUserRoleResult']:
        """
        List of user’s roles and the databases / collections on which the roles apply. A role allows the user to perform particular actions on the specified database. A role on the admin database can include privileges that apply to the other databases as well. See Roles below for more details.
        """
        return pulumi.get(self, "roles")

    @property
    @pulumi.getter
    def scopes(self) -> Sequence['outputs.GetDatabaseUserScopeResult']:
        """
        Array of clusters and Atlas Data Lakes that this user has access to.
        """
        return pulumi.get(self, "scopes")

    @property
    @pulumi.getter
    def username(self) -> str:
        return pulumi.get(self, "username")

    @property
    @pulumi.getter(name="x509Type")
    def x509_type(self) -> str:
        """
        X.509 method by which the provided username is authenticated.
        """
        return pulumi.get(self, "x509_type")


class AwaitableGetDatabaseUserResult(GetDatabaseUserResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDatabaseUserResult(
            auth_database_name=self.auth_database_name,
            aws_iam_type=self.aws_iam_type,
            database_name=self.database_name,
            id=self.id,
            labels=self.labels,
            ldap_auth_type=self.ldap_auth_type,
            project_id=self.project_id,
            roles=self.roles,
            scopes=self.scopes,
            username=self.username,
            x509_type=self.x509_type)


def get_database_user(auth_database_name: Optional[str] = None,
                      database_name: Optional[str] = None,
                      project_id: Optional[str] = None,
                      username: Optional[str] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDatabaseUserResult:
    """
    `DatabaseUser` describe a Database User. This represents a database user which will be applied to all clusters within the project.

    Each user has a set of roles that provide access to the project’s databases. User's roles apply to all the clusters in the project: if two clusters have a `products` database and a user has a role granting `read` access on the products database, the user has that access on both clusters.

    > **NOTE:** Groups and projects are synonymous terms. You may find group_id in the official documentation.


    :param str auth_database_name: The user’s authentication database. A user must provide both a username and authentication database to log into MongoDB. In Atlas deployments of MongoDB, the authentication database is almost always the admin database, for X509 it is $external.
    :param str database_name: Database on which the user has the specified role. A role on the `admin` database can include privileges that apply to the other databases.
    :param str project_id: The unique ID for the project to create the database user.
    :param str username: Username for authenticating to MongoDB.
    """
    __args__ = dict()
    __args__['authDatabaseName'] = auth_database_name
    __args__['databaseName'] = database_name
    __args__['projectId'] = project_id
    __args__['username'] = username
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('mongodbatlas:index/getDatabaseUser:getDatabaseUser', __args__, opts=opts, typ=GetDatabaseUserResult).value

    return AwaitableGetDatabaseUserResult(
        auth_database_name=__ret__.auth_database_name,
        aws_iam_type=__ret__.aws_iam_type,
        database_name=__ret__.database_name,
        id=__ret__.id,
        labels=__ret__.labels,
        ldap_auth_type=__ret__.ldap_auth_type,
        project_id=__ret__.project_id,
        roles=__ret__.roles,
        scopes=__ret__.scopes,
        username=__ret__.username,
        x509_type=__ret__.x509_type)


@_utilities.lift_output_func(get_database_user)
def get_database_user_output(auth_database_name: Optional[pulumi.Input[Optional[str]]] = None,
                             database_name: Optional[pulumi.Input[Optional[str]]] = None,
                             project_id: Optional[pulumi.Input[str]] = None,
                             username: Optional[pulumi.Input[str]] = None,
                             opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetDatabaseUserResult]:
    """
    `DatabaseUser` describe a Database User. This represents a database user which will be applied to all clusters within the project.

    Each user has a set of roles that provide access to the project’s databases. User's roles apply to all the clusters in the project: if two clusters have a `products` database and a user has a role granting `read` access on the products database, the user has that access on both clusters.

    > **NOTE:** Groups and projects are synonymous terms. You may find group_id in the official documentation.


    :param str auth_database_name: The user’s authentication database. A user must provide both a username and authentication database to log into MongoDB. In Atlas deployments of MongoDB, the authentication database is almost always the admin database, for X509 it is $external.
    :param str database_name: Database on which the user has the specified role. A role on the `admin` database can include privileges that apply to the other databases.
    :param str project_id: The unique ID for the project to create the database user.
    :param str username: Username for authenticating to MongoDB.
    """
    ...
