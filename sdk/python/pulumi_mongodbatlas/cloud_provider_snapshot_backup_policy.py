# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from . import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['CloudProviderSnapshotBackupPolicy']


class CloudProviderSnapshotBackupPolicy(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cluster_name: Optional[pulumi.Input[str]] = None,
                 policies: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['CloudProviderSnapshotBackupPolicyPolicyArgs']]]]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 reference_hour_of_day: Optional[pulumi.Input[int]] = None,
                 reference_minute_of_hour: Optional[pulumi.Input[int]] = None,
                 restore_window_days: Optional[pulumi.Input[int]] = None,
                 update_snapshots: Optional[pulumi.Input[bool]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        ## Import

        Cloud Backup Snapshot Policy entries can be imported using project project_id and cluster_name, in the format `PROJECTID-CLUSTERNAME`, e.g.

        ```sh
         $ pulumi import mongodbatlas:index/cloudProviderSnapshotBackupPolicy:CloudProviderSnapshotBackupPolicy test 5d0f1f73cf09a29120e173cf-MyClusterTest
        ```

         For more information see[MongoDB Atlas API Reference.](https://docs.atlas.mongodb.com/reference/api/cloud-backup/schedule/modify-one-schedule/)

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cluster_name: The name of the Atlas cluster that contains the snapshot backup policy you want to retrieve.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['CloudProviderSnapshotBackupPolicyPolicyArgs']]]] policies: Contains a document for each backup policy item in the desired updated backup policy.
               * `policies.#.id` - (Required) Unique identifier of the backup policy that you want to update. policies.#.id is a value obtained via the Cluster resource. provider_backup_enabled of the Cluster resource must be set to true. See the example above for how to refer to the Cluster resource for policies.#.id
        :param pulumi.Input[str] project_id: The unique identifier of the project for the Atlas cluster.
        :param pulumi.Input[int] reference_hour_of_day: UTC Hour of day between 0 and 23, inclusive, representing which hour of the day that Atlas takes snapshots for backup policy items.
        :param pulumi.Input[int] reference_minute_of_hour: UTC Minutes after referenceHourOfDay that Atlas takes snapshots for backup policy items. Must be between 0 and 59, inclusive.
        :param pulumi.Input[int] restore_window_days: Number of days back in time you can restore to with point-in-time accuracy. Must be a positive, non-zero integer.
        :param pulumi.Input[bool] update_snapshots: Specify true to apply the retention changes in the updated backup policy to snapshots that Atlas took previously.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            if cluster_name is None:
                raise TypeError("Missing required property 'cluster_name'")
            __props__['cluster_name'] = cluster_name
            if policies is None:
                raise TypeError("Missing required property 'policies'")
            __props__['policies'] = policies
            if project_id is None:
                raise TypeError("Missing required property 'project_id'")
            __props__['project_id'] = project_id
            __props__['reference_hour_of_day'] = reference_hour_of_day
            __props__['reference_minute_of_hour'] = reference_minute_of_hour
            __props__['restore_window_days'] = restore_window_days
            __props__['update_snapshots'] = update_snapshots
            __props__['cluster_id'] = None
            __props__['next_snapshot'] = None
        super(CloudProviderSnapshotBackupPolicy, __self__).__init__(
            'mongodbatlas:index/cloudProviderSnapshotBackupPolicy:CloudProviderSnapshotBackupPolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            cluster_id: Optional[pulumi.Input[str]] = None,
            cluster_name: Optional[pulumi.Input[str]] = None,
            next_snapshot: Optional[pulumi.Input[str]] = None,
            policies: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['CloudProviderSnapshotBackupPolicyPolicyArgs']]]]] = None,
            project_id: Optional[pulumi.Input[str]] = None,
            reference_hour_of_day: Optional[pulumi.Input[int]] = None,
            reference_minute_of_hour: Optional[pulumi.Input[int]] = None,
            restore_window_days: Optional[pulumi.Input[int]] = None,
            update_snapshots: Optional[pulumi.Input[bool]] = None) -> 'CloudProviderSnapshotBackupPolicy':
        """
        Get an existing CloudProviderSnapshotBackupPolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cluster_id: Unique identifier of the Atlas cluster.
        :param pulumi.Input[str] cluster_name: The name of the Atlas cluster that contains the snapshot backup policy you want to retrieve.
        :param pulumi.Input[str] next_snapshot: Timestamp in the number of seconds that have elapsed since the UNIX epoch when Atlas takes the next snapshot.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['CloudProviderSnapshotBackupPolicyPolicyArgs']]]] policies: Contains a document for each backup policy item in the desired updated backup policy.
               * `policies.#.id` - (Required) Unique identifier of the backup policy that you want to update. policies.#.id is a value obtained via the Cluster resource. provider_backup_enabled of the Cluster resource must be set to true. See the example above for how to refer to the Cluster resource for policies.#.id
        :param pulumi.Input[str] project_id: The unique identifier of the project for the Atlas cluster.
        :param pulumi.Input[int] reference_hour_of_day: UTC Hour of day between 0 and 23, inclusive, representing which hour of the day that Atlas takes snapshots for backup policy items.
        :param pulumi.Input[int] reference_minute_of_hour: UTC Minutes after referenceHourOfDay that Atlas takes snapshots for backup policy items. Must be between 0 and 59, inclusive.
        :param pulumi.Input[int] restore_window_days: Number of days back in time you can restore to with point-in-time accuracy. Must be a positive, non-zero integer.
        :param pulumi.Input[bool] update_snapshots: Specify true to apply the retention changes in the updated backup policy to snapshots that Atlas took previously.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["cluster_id"] = cluster_id
        __props__["cluster_name"] = cluster_name
        __props__["next_snapshot"] = next_snapshot
        __props__["policies"] = policies
        __props__["project_id"] = project_id
        __props__["reference_hour_of_day"] = reference_hour_of_day
        __props__["reference_minute_of_hour"] = reference_minute_of_hour
        __props__["restore_window_days"] = restore_window_days
        __props__["update_snapshots"] = update_snapshots
        return CloudProviderSnapshotBackupPolicy(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> pulumi.Output[str]:
        """
        Unique identifier of the Atlas cluster.
        """
        return pulumi.get(self, "cluster_id")

    @property
    @pulumi.getter(name="clusterName")
    def cluster_name(self) -> pulumi.Output[str]:
        """
        The name of the Atlas cluster that contains the snapshot backup policy you want to retrieve.
        """
        return pulumi.get(self, "cluster_name")

    @property
    @pulumi.getter(name="nextSnapshot")
    def next_snapshot(self) -> pulumi.Output[str]:
        """
        Timestamp in the number of seconds that have elapsed since the UNIX epoch when Atlas takes the next snapshot.
        """
        return pulumi.get(self, "next_snapshot")

    @property
    @pulumi.getter
    def policies(self) -> pulumi.Output[Sequence['outputs.CloudProviderSnapshotBackupPolicyPolicy']]:
        """
        Contains a document for each backup policy item in the desired updated backup policy.
        * `policies.#.id` - (Required) Unique identifier of the backup policy that you want to update. policies.#.id is a value obtained via the Cluster resource. provider_backup_enabled of the Cluster resource must be set to true. See the example above for how to refer to the Cluster resource for policies.#.id
        """
        return pulumi.get(self, "policies")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Output[str]:
        """
        The unique identifier of the project for the Atlas cluster.
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter(name="referenceHourOfDay")
    def reference_hour_of_day(self) -> pulumi.Output[int]:
        """
        UTC Hour of day between 0 and 23, inclusive, representing which hour of the day that Atlas takes snapshots for backup policy items.
        """
        return pulumi.get(self, "reference_hour_of_day")

    @property
    @pulumi.getter(name="referenceMinuteOfHour")
    def reference_minute_of_hour(self) -> pulumi.Output[int]:
        """
        UTC Minutes after referenceHourOfDay that Atlas takes snapshots for backup policy items. Must be between 0 and 59, inclusive.
        """
        return pulumi.get(self, "reference_minute_of_hour")

    @property
    @pulumi.getter(name="restoreWindowDays")
    def restore_window_days(self) -> pulumi.Output[int]:
        """
        Number of days back in time you can restore to with point-in-time accuracy. Must be a positive, non-zero integer.
        """
        return pulumi.get(self, "restore_window_days")

    @property
    @pulumi.getter(name="updateSnapshots")
    def update_snapshots(self) -> pulumi.Output[bool]:
        """
        Specify true to apply the retention changes in the updated backup policy to snapshots that Atlas took previously.
        """
        return pulumi.get(self, "update_snapshots")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

