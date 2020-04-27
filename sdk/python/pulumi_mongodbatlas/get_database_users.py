# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from . import utilities, tables

class GetDatabaseUsersResult:
    """
    A collection of values returned by getDatabaseUsers.
    """
    def __init__(__self__, id=None, project_id=None, results=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """
        if project_id and not isinstance(project_id, str):
            raise TypeError("Expected argument 'project_id' to be a str")
        __self__.project_id = project_id
        """
        ID of the Atlas project the user belongs to.
        """
        if results and not isinstance(results, list):
            raise TypeError("Expected argument 'results' to be a list")
        __self__.results = results
        """
        A list where each represents a Database user.
        """
class AwaitableGetDatabaseUsersResult(GetDatabaseUsersResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDatabaseUsersResult(
            id=self.id,
            project_id=self.project_id,
            results=self.results)

def get_database_users(project_id=None,opts=None):
    """
    `.getDatabaseUsers` describe all Database Users. This represents a database user which will be applied to all clusters within the project.

    Each user has a set of roles that provide access to the project’s databases. User's roles apply to all the clusters in the project: if two clusters have a `products` database and a user has a role granting `read` access on the products database, the user has that access on both clusters.

    > **NOTE:** Groups and projects are synonymous terms. You may find `groupId` in the official documentation.




    :param str project_id: The unique ID for the project to get all database users.
    """
    __args__ = dict()


    __args__['projectId'] = project_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('mongodbatlas:index/getDatabaseUsers:getDatabaseUsers', __args__, opts=opts).value

    return AwaitableGetDatabaseUsersResult(
        id=__ret__.get('id'),
        project_id=__ret__.get('projectId'),
        results=__ret__.get('results'))