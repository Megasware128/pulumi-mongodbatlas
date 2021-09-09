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
    'GetCloudBackupScheduleResult',
    'AwaitableGetCloudBackupScheduleResult',
    'get_cloud_backup_schedule',
]

@pulumi.output_type
class GetCloudBackupScheduleResult:
    """
    A collection of values returned by getCloudBackupSchedule.
    """
    def __init__(__self__, cluster_id=None, cluster_name=None, id=None, id_policy=None, next_snapshot=None, policy_item_dailies=None, policy_item_hourlies=None, policy_item_monthlies=None, policy_item_weeklies=None, project_id=None, reference_hour_of_day=None, reference_minute_of_hour=None, restore_window_days=None):
        if cluster_id and not isinstance(cluster_id, str):
            raise TypeError("Expected argument 'cluster_id' to be a str")
        pulumi.set(__self__, "cluster_id", cluster_id)
        if cluster_name and not isinstance(cluster_name, str):
            raise TypeError("Expected argument 'cluster_name' to be a str")
        pulumi.set(__self__, "cluster_name", cluster_name)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if id_policy and not isinstance(id_policy, str):
            raise TypeError("Expected argument 'id_policy' to be a str")
        pulumi.set(__self__, "id_policy", id_policy)
        if next_snapshot and not isinstance(next_snapshot, str):
            raise TypeError("Expected argument 'next_snapshot' to be a str")
        pulumi.set(__self__, "next_snapshot", next_snapshot)
        if policy_item_dailies and not isinstance(policy_item_dailies, list):
            raise TypeError("Expected argument 'policy_item_dailies' to be a list")
        pulumi.set(__self__, "policy_item_dailies", policy_item_dailies)
        if policy_item_hourlies and not isinstance(policy_item_hourlies, list):
            raise TypeError("Expected argument 'policy_item_hourlies' to be a list")
        pulumi.set(__self__, "policy_item_hourlies", policy_item_hourlies)
        if policy_item_monthlies and not isinstance(policy_item_monthlies, list):
            raise TypeError("Expected argument 'policy_item_monthlies' to be a list")
        pulumi.set(__self__, "policy_item_monthlies", policy_item_monthlies)
        if policy_item_weeklies and not isinstance(policy_item_weeklies, list):
            raise TypeError("Expected argument 'policy_item_weeklies' to be a list")
        pulumi.set(__self__, "policy_item_weeklies", policy_item_weeklies)
        if project_id and not isinstance(project_id, str):
            raise TypeError("Expected argument 'project_id' to be a str")
        pulumi.set(__self__, "project_id", project_id)
        if reference_hour_of_day and not isinstance(reference_hour_of_day, int):
            raise TypeError("Expected argument 'reference_hour_of_day' to be a int")
        pulumi.set(__self__, "reference_hour_of_day", reference_hour_of_day)
        if reference_minute_of_hour and not isinstance(reference_minute_of_hour, int):
            raise TypeError("Expected argument 'reference_minute_of_hour' to be a int")
        pulumi.set(__self__, "reference_minute_of_hour", reference_minute_of_hour)
        if restore_window_days and not isinstance(restore_window_days, int):
            raise TypeError("Expected argument 'restore_window_days' to be a int")
        pulumi.set(__self__, "restore_window_days", restore_window_days)

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> str:
        """
        Unique identifier of the Atlas cluster.
        """
        return pulumi.get(self, "cluster_id")

    @property
    @pulumi.getter(name="clusterName")
    def cluster_name(self) -> str:
        return pulumi.get(self, "cluster_name")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="idPolicy")
    def id_policy(self) -> str:
        """
        Unique identifier of the backup policy.
        """
        return pulumi.get(self, "id_policy")

    @property
    @pulumi.getter(name="nextSnapshot")
    def next_snapshot(self) -> str:
        """
        UTC ISO 8601 formatted point in time when Atlas will take the next snapshot.
        """
        return pulumi.get(self, "next_snapshot")

    @property
    @pulumi.getter(name="policyItemDailies")
    def policy_item_dailies(self) -> Sequence['outputs.GetCloudBackupSchedulePolicyItemDailyResult']:
        """
        Daily policy item
        """
        return pulumi.get(self, "policy_item_dailies")

    @property
    @pulumi.getter(name="policyItemHourlies")
    def policy_item_hourlies(self) -> Sequence['outputs.GetCloudBackupSchedulePolicyItemHourlyResult']:
        """
        Hourly policy item
        """
        return pulumi.get(self, "policy_item_hourlies")

    @property
    @pulumi.getter(name="policyItemMonthlies")
    def policy_item_monthlies(self) -> Sequence['outputs.GetCloudBackupSchedulePolicyItemMonthlyResult']:
        """
        Monthly policy item
        """
        return pulumi.get(self, "policy_item_monthlies")

    @property
    @pulumi.getter(name="policyItemWeeklies")
    def policy_item_weeklies(self) -> Sequence['outputs.GetCloudBackupSchedulePolicyItemWeeklyResult']:
        """
        Weekly policy item
        """
        return pulumi.get(self, "policy_item_weeklies")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> str:
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter(name="referenceHourOfDay")
    def reference_hour_of_day(self) -> int:
        """
        UTC Hour of day between 0 and 23 representing which hour of the day that Atlas takes a snapshot.
        """
        return pulumi.get(self, "reference_hour_of_day")

    @property
    @pulumi.getter(name="referenceMinuteOfHour")
    def reference_minute_of_hour(self) -> int:
        """
        UTC Minute of day between 0 and 59 representing which minute of the `reference_hour_of_day` that Atlas takes the snapshot.
        """
        return pulumi.get(self, "reference_minute_of_hour")

    @property
    @pulumi.getter(name="restoreWindowDays")
    def restore_window_days(self) -> int:
        """
        Specifies a restore window in days for cloud backup to maintain.
        """
        return pulumi.get(self, "restore_window_days")


class AwaitableGetCloudBackupScheduleResult(GetCloudBackupScheduleResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetCloudBackupScheduleResult(
            cluster_id=self.cluster_id,
            cluster_name=self.cluster_name,
            id=self.id,
            id_policy=self.id_policy,
            next_snapshot=self.next_snapshot,
            policy_item_dailies=self.policy_item_dailies,
            policy_item_hourlies=self.policy_item_hourlies,
            policy_item_monthlies=self.policy_item_monthlies,
            policy_item_weeklies=self.policy_item_weeklies,
            project_id=self.project_id,
            reference_hour_of_day=self.reference_hour_of_day,
            reference_minute_of_hour=self.reference_minute_of_hour,
            restore_window_days=self.restore_window_days)


def get_cloud_backup_schedule(cluster_name: Optional[str] = None,
                              project_id: Optional[str] = None,
                              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetCloudBackupScheduleResult:
    """
    `CloudBackupSchedule` provides a Cloud Backup Schedule datasource. An Atlas Cloud Backup Schedule provides the current cloud backup schedule for the cluster.

    > **NOTE:** Groups and projects are synonymous terms. You may find `groupId` in the official documentation.


    :param str cluster_name: The name of the Atlas cluster that contains the snapshots backup policy you want to retrieve.
    :param str project_id: The unique identifier of the project for the Atlas cluster.
    """
    __args__ = dict()
    __args__['clusterName'] = cluster_name
    __args__['projectId'] = project_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('mongodbatlas:index/getCloudBackupSchedule:getCloudBackupSchedule', __args__, opts=opts, typ=GetCloudBackupScheduleResult).value

    return AwaitableGetCloudBackupScheduleResult(
        cluster_id=__ret__.cluster_id,
        cluster_name=__ret__.cluster_name,
        id=__ret__.id,
        id_policy=__ret__.id_policy,
        next_snapshot=__ret__.next_snapshot,
        policy_item_dailies=__ret__.policy_item_dailies,
        policy_item_hourlies=__ret__.policy_item_hourlies,
        policy_item_monthlies=__ret__.policy_item_monthlies,
        policy_item_weeklies=__ret__.policy_item_weeklies,
        project_id=__ret__.project_id,
        reference_hour_of_day=__ret__.reference_hour_of_day,
        reference_minute_of_hour=__ret__.reference_minute_of_hour,
        restore_window_days=__ret__.restore_window_days)
