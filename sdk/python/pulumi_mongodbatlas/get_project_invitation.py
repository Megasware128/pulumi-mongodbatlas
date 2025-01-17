# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'GetProjectInvitationResult',
    'AwaitableGetProjectInvitationResult',
    'get_project_invitation',
    'get_project_invitation_output',
]

@pulumi.output_type
class GetProjectInvitationResult:
    """
    A collection of values returned by getProjectInvitation.
    """
    def __init__(__self__, created_at=None, expires_at=None, id=None, invitation_id=None, inviter_username=None, project_id=None, roles=None, username=None):
        if created_at and not isinstance(created_at, str):
            raise TypeError("Expected argument 'created_at' to be a str")
        pulumi.set(__self__, "created_at", created_at)
        if expires_at and not isinstance(expires_at, str):
            raise TypeError("Expected argument 'expires_at' to be a str")
        pulumi.set(__self__, "expires_at", expires_at)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if invitation_id and not isinstance(invitation_id, str):
            raise TypeError("Expected argument 'invitation_id' to be a str")
        pulumi.set(__self__, "invitation_id", invitation_id)
        if inviter_username and not isinstance(inviter_username, str):
            raise TypeError("Expected argument 'inviter_username' to be a str")
        pulumi.set(__self__, "inviter_username", inviter_username)
        if project_id and not isinstance(project_id, str):
            raise TypeError("Expected argument 'project_id' to be a str")
        pulumi.set(__self__, "project_id", project_id)
        if roles and not isinstance(roles, list):
            raise TypeError("Expected argument 'roles' to be a list")
        pulumi.set(__self__, "roles", roles)
        if username and not isinstance(username, str):
            raise TypeError("Expected argument 'username' to be a str")
        pulumi.set(__self__, "username", username)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> str:
        """
        Timestamp in ISO 8601 date and time format in UTC when Atlas sent the invitation.
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter(name="expiresAt")
    def expires_at(self) -> str:
        """
        Timestamp in ISO 8601 date and time format in UTC when the invitation expires. Users have 30 days to accept an invitation.
        """
        return pulumi.get(self, "expires_at")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="invitationId")
    def invitation_id(self) -> str:
        return pulumi.get(self, "invitation_id")

    @property
    @pulumi.getter(name="inviterUsername")
    def inviter_username(self) -> str:
        """
        Atlas user who invited `username` to the project.
        """
        return pulumi.get(self, "inviter_username")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> str:
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter
    def roles(self) -> Sequence[str]:
        """
        Atlas roles to assign to the invited user. If the user accepts the invitation, Atlas assigns these roles to them. The following options are available:
        * GROUP_OWNER
        * GROUP_CLUSTER_MANAGER
        * GROUP_READ_ONLY
        * GROUP_DATA_ACCESS_ADMIN
        * GROUP_DATA_ACCESS_READ_WRITE
        * GROUP_DATA_ACCESS_READ_ONLY
        """
        return pulumi.get(self, "roles")

    @property
    @pulumi.getter
    def username(self) -> str:
        return pulumi.get(self, "username")


class AwaitableGetProjectInvitationResult(GetProjectInvitationResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetProjectInvitationResult(
            created_at=self.created_at,
            expires_at=self.expires_at,
            id=self.id,
            invitation_id=self.invitation_id,
            inviter_username=self.inviter_username,
            project_id=self.project_id,
            roles=self.roles,
            username=self.username)


def get_project_invitation(invitation_id: Optional[str] = None,
                           project_id: Optional[str] = None,
                           username: Optional[str] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetProjectInvitationResult:
    """
    `ProjectInvitation` describes an invitation to a user to join an Atlas project.

    > **NOTE:** Groups and projects are synonymous terms. You may find GROUP-ID in the official documentation.

    ## Example Usage


    :param str invitation_id: Unique 24-hexadecimal digit string that identifies the invitation in Atlas.
    :param str project_id: Unique 24-hexadecimal digit string that identifies the project to which you invited the user.
    :param str username: Email address of the invited user. This is the address to which Atlas sends the invite. If the user accepts the invitation, they log in to Atlas with this username.
    """
    __args__ = dict()
    __args__['invitationId'] = invitation_id
    __args__['projectId'] = project_id
    __args__['username'] = username
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('mongodbatlas:index/getProjectInvitation:getProjectInvitation', __args__, opts=opts, typ=GetProjectInvitationResult).value

    return AwaitableGetProjectInvitationResult(
        created_at=__ret__.created_at,
        expires_at=__ret__.expires_at,
        id=__ret__.id,
        invitation_id=__ret__.invitation_id,
        inviter_username=__ret__.inviter_username,
        project_id=__ret__.project_id,
        roles=__ret__.roles,
        username=__ret__.username)


@_utilities.lift_output_func(get_project_invitation)
def get_project_invitation_output(invitation_id: Optional[pulumi.Input[str]] = None,
                                  project_id: Optional[pulumi.Input[str]] = None,
                                  username: Optional[pulumi.Input[str]] = None,
                                  opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetProjectInvitationResult]:
    """
    `ProjectInvitation` describes an invitation to a user to join an Atlas project.

    > **NOTE:** Groups and projects are synonymous terms. You may find GROUP-ID in the official documentation.

    ## Example Usage


    :param str invitation_id: Unique 24-hexadecimal digit string that identifies the invitation in Atlas.
    :param str project_id: Unique 24-hexadecimal digit string that identifies the project to which you invited the user.
    :param str username: Email address of the invited user. This is the address to which Atlas sends the invite. If the user accepts the invitation, they log in to Atlas with this username.
    """
    ...
