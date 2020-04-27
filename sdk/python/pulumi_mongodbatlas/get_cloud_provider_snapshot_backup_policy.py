# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from . import utilities, tables

class GetCloudProviderSnapshotBackupPolicyResult:
    """
    A collection of values returned by getCloudProviderSnapshotBackupPolicy.
    """
    def __init__(__self__, cluster_id=None, cluster_name=None, id=None, next_snapshot=None, policies=None, project_id=None, reference_hour_of_day=None, reference_minute_of_hour=None, restore_window_days=None, update_snapshots=None):
        if cluster_id and not isinstance(cluster_id, str):
            raise TypeError("Expected argument 'cluster_id' to be a str")
        __self__.cluster_id = cluster_id
        """
        Unique identifier of the Atlas cluster.
        """
        if cluster_name and not isinstance(cluster_name, str):
            raise TypeError("Expected argument 'cluster_name' to be a str")
        __self__.cluster_name = cluster_name
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """
        if next_snapshot and not isinstance(next_snapshot, str):
            raise TypeError("Expected argument 'next_snapshot' to be a str")
        __self__.next_snapshot = next_snapshot
        """
        UTC ISO 8601 formatted point in time when Atlas will take the next snapshot.
        """
        if policies and not isinstance(policies, list):
            raise TypeError("Expected argument 'policies' to be a list")
        __self__.policies = policies
        """
        A list of policy definitions for the cluster.
        * `policies.#.id` - Unique identifier of the backup policy.
        """
        if project_id and not isinstance(project_id, str):
            raise TypeError("Expected argument 'project_id' to be a str")
        __self__.project_id = project_id
        if reference_hour_of_day and not isinstance(reference_hour_of_day, float):
            raise TypeError("Expected argument 'reference_hour_of_day' to be a float")
        __self__.reference_hour_of_day = reference_hour_of_day
        """
        UTC Hour of day between 0 and 23 representing which hour of the day that Atlas takes a snapshot.
        """
        if reference_minute_of_hour and not isinstance(reference_minute_of_hour, float):
            raise TypeError("Expected argument 'reference_minute_of_hour' to be a float")
        __self__.reference_minute_of_hour = reference_minute_of_hour
        """
        UTC Minute of day between 0 and 59 representing which minute of the referenceHourOfDay that Atlas takes the snapshot.
        """
        if restore_window_days and not isinstance(restore_window_days, float):
            raise TypeError("Expected argument 'restore_window_days' to be a float")
        __self__.restore_window_days = restore_window_days
        """
        Specifies a restore window in days for the cloud provider backup to maintain.
        """
        if update_snapshots and not isinstance(update_snapshots, bool):
            raise TypeError("Expected argument 'update_snapshots' to be a bool")
        __self__.update_snapshots = update_snapshots
class AwaitableGetCloudProviderSnapshotBackupPolicyResult(GetCloudProviderSnapshotBackupPolicyResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetCloudProviderSnapshotBackupPolicyResult(
            cluster_id=self.cluster_id,
            cluster_name=self.cluster_name,
            id=self.id,
            next_snapshot=self.next_snapshot,
            policies=self.policies,
            project_id=self.project_id,
            reference_hour_of_day=self.reference_hour_of_day,
            reference_minute_of_hour=self.reference_minute_of_hour,
            restore_window_days=self.restore_window_days,
            update_snapshots=self.update_snapshots)

def get_cloud_provider_snapshot_backup_policy(cluster_name=None,project_id=None,opts=None):
    """
    `.CloudProviderSnapshotBackupPolicy` provides a Cloud Provider Snapshot Backup Policy entry datasource. An Atlas Cloud Provider Snapshot Backup Policy provides the current snapshot schedule and retention settings for the cluster.

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
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('mongodbatlas:index/getCloudProviderSnapshotBackupPolicy:getCloudProviderSnapshotBackupPolicy', __args__, opts=opts).value

    return AwaitableGetCloudProviderSnapshotBackupPolicyResult(
        cluster_id=__ret__.get('clusterId'),
        cluster_name=__ret__.get('clusterName'),
        id=__ret__.get('id'),
        next_snapshot=__ret__.get('nextSnapshot'),
        policies=__ret__.get('policies'),
        project_id=__ret__.get('projectId'),
        reference_hour_of_day=__ret__.get('referenceHourOfDay'),
        reference_minute_of_hour=__ret__.get('referenceMinuteOfHour'),
        restore_window_days=__ret__.get('restoreWindowDays'),
        update_snapshots=__ret__.get('updateSnapshots'))