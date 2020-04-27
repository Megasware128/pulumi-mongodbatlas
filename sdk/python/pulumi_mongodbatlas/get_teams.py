# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from . import utilities, tables

class GetTeamsResult:
    """
    A collection of values returned by getTeams.
    """
    def __init__(__self__, id=None, name=None, org_id=None, team_id=None, usernames=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
        if org_id and not isinstance(org_id, str):
            raise TypeError("Expected argument 'org_id' to be a str")
        __self__.org_id = org_id
        if team_id and not isinstance(team_id, str):
            raise TypeError("Expected argument 'team_id' to be a str")
        __self__.team_id = team_id
        if usernames and not isinstance(usernames, list):
            raise TypeError("Expected argument 'usernames' to be a list")
        __self__.usernames = usernames
class AwaitableGetTeamsResult(GetTeamsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetTeamsResult(
            id=self.id,
            name=self.name,
            org_id=self.org_id,
            team_id=self.team_id,
            usernames=self.usernames)

def get_teams(org_id=None,team_id=None,opts=None):
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()


    __args__['orgId'] = org_id
    __args__['teamId'] = team_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('mongodbatlas:index/getTeams:getTeams', __args__, opts=opts).value

    return AwaitableGetTeamsResult(
        id=__ret__.get('id'),
        name=__ret__.get('name'),
        org_id=__ret__.get('orgId'),
        team_id=__ret__.get('teamId'),
        usernames=__ret__.get('usernames'))