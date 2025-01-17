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
    'GetOnlineArchiveResult',
    'AwaitableGetOnlineArchiveResult',
    'get_online_archive',
    'get_online_archive_output',
]

@pulumi.output_type
class GetOnlineArchiveResult:
    """
    A collection of values returned by getOnlineArchive.
    """
    def __init__(__self__, archive_id=None, cluster_name=None, coll_name=None, criterias=None, db_name=None, id=None, partition_fields=None, paused=None, project_id=None, state=None):
        if archive_id and not isinstance(archive_id, str):
            raise TypeError("Expected argument 'archive_id' to be a str")
        pulumi.set(__self__, "archive_id", archive_id)
        if cluster_name and not isinstance(cluster_name, str):
            raise TypeError("Expected argument 'cluster_name' to be a str")
        pulumi.set(__self__, "cluster_name", cluster_name)
        if coll_name and not isinstance(coll_name, str):
            raise TypeError("Expected argument 'coll_name' to be a str")
        pulumi.set(__self__, "coll_name", coll_name)
        if criterias and not isinstance(criterias, list):
            raise TypeError("Expected argument 'criterias' to be a list")
        pulumi.set(__self__, "criterias", criterias)
        if db_name and not isinstance(db_name, str):
            raise TypeError("Expected argument 'db_name' to be a str")
        pulumi.set(__self__, "db_name", db_name)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if partition_fields and not isinstance(partition_fields, list):
            raise TypeError("Expected argument 'partition_fields' to be a list")
        pulumi.set(__self__, "partition_fields", partition_fields)
        if paused and not isinstance(paused, bool):
            raise TypeError("Expected argument 'paused' to be a bool")
        pulumi.set(__self__, "paused", paused)
        if project_id and not isinstance(project_id, str):
            raise TypeError("Expected argument 'project_id' to be a str")
        pulumi.set(__self__, "project_id", project_id)
        if state and not isinstance(state, str):
            raise TypeError("Expected argument 'state' to be a str")
        pulumi.set(__self__, "state", state)

    @property
    @pulumi.getter(name="archiveId")
    def archive_id(self) -> str:
        return pulumi.get(self, "archive_id")

    @property
    @pulumi.getter(name="clusterName")
    def cluster_name(self) -> str:
        return pulumi.get(self, "cluster_name")

    @property
    @pulumi.getter(name="collName")
    def coll_name(self) -> str:
        return pulumi.get(self, "coll_name")

    @property
    @pulumi.getter
    def criterias(self) -> Sequence['outputs.GetOnlineArchiveCriteriaResult']:
        return pulumi.get(self, "criterias")

    @property
    @pulumi.getter(name="dbName")
    def db_name(self) -> str:
        return pulumi.get(self, "db_name")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="partitionFields")
    def partition_fields(self) -> Sequence['outputs.GetOnlineArchivePartitionFieldResult']:
        return pulumi.get(self, "partition_fields")

    @property
    @pulumi.getter
    def paused(self) -> bool:
        return pulumi.get(self, "paused")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> str:
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter
    def state(self) -> str:
        return pulumi.get(self, "state")


class AwaitableGetOnlineArchiveResult(GetOnlineArchiveResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetOnlineArchiveResult(
            archive_id=self.archive_id,
            cluster_name=self.cluster_name,
            coll_name=self.coll_name,
            criterias=self.criterias,
            db_name=self.db_name,
            id=self.id,
            partition_fields=self.partition_fields,
            paused=self.paused,
            project_id=self.project_id,
            state=self.state)


def get_online_archive(archive_id: Optional[str] = None,
                       cluster_name: Optional[str] = None,
                       project_id: Optional[str] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetOnlineArchiveResult:
    """
    `OnlineArchive` describes an Online Archive

    > **NOTE:** Groups and projects are synonymous terms. You may find group_id in the official documentation.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_mongodbatlas as mongodbatlas

    test = mongodbatlas.get_online_archive(project_id=var["project_id"],
        cluster_name=var["cluster_name"],
        archive_id="5ebad3c1fe9c0ab8d37d61e1")
    ```
    ## Attributes reference

    * `db_name`          -  Name of the database that contains the collection.
    * `coll_name`        -  Name of the collection.
    * `criteria`         -  Criteria to use for archiving data.
    * `criteria.type`          - Type of criteria (DATE, CUSTOM)
    * `criteria.date_field`    - Name of an already indexed date field from the documents. Data is archived when the current date is greater than the value of the date field specified here plus the number of days specified via the `expire_after_days` parameter.
    * `criteria.date_format`   - the date format. Valid values:  ISODATE (default), EPOCH_SECONDS, EPOCH_MILLIS, EPOCH_NANOSECONDS
    * `criteria.expire_after_days` - Number of days that specifies the age limit for the data in the live Atlas cluster.
    * `criteria.query` - JSON query to use to select documents for archiving. Only for `CUSTOM` type
    * `partition_fields` -  Fields to use to partition data.
    * `partition_fields.field_name` - Name of the field. To specify a nested field, use the dot notation.
    * `partition_fields.order` - Position of the field in the partition. Value can be: 0,1,2
      By default, the date field specified in the criteria.dateField parameter is in the first position of the partition.
    * `partitio_fields.field_type` - Type of the partition field
    * `state`    - Status of the online archive. Valid values are: Pending, Archiving, Idle, Pausing, Paused, Orphaned and Deleted

    See [MongoDB Atlas API](https://docs.atlas.mongodb.com/reference/api/online-archive-get-one/) Documentation for more information.


    :param str archive_id: ID of the online archive.
    :param str cluster_name: Name of the cluster that contains the collection.
    :param str project_id: The unique ID for the project.
    """
    __args__ = dict()
    __args__['archiveId'] = archive_id
    __args__['clusterName'] = cluster_name
    __args__['projectId'] = project_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('mongodbatlas:index/getOnlineArchive:getOnlineArchive', __args__, opts=opts, typ=GetOnlineArchiveResult).value

    return AwaitableGetOnlineArchiveResult(
        archive_id=__ret__.archive_id,
        cluster_name=__ret__.cluster_name,
        coll_name=__ret__.coll_name,
        criterias=__ret__.criterias,
        db_name=__ret__.db_name,
        id=__ret__.id,
        partition_fields=__ret__.partition_fields,
        paused=__ret__.paused,
        project_id=__ret__.project_id,
        state=__ret__.state)


@_utilities.lift_output_func(get_online_archive)
def get_online_archive_output(archive_id: Optional[pulumi.Input[str]] = None,
                              cluster_name: Optional[pulumi.Input[str]] = None,
                              project_id: Optional[pulumi.Input[str]] = None,
                              opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetOnlineArchiveResult]:
    """
    `OnlineArchive` describes an Online Archive

    > **NOTE:** Groups and projects are synonymous terms. You may find group_id in the official documentation.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_mongodbatlas as mongodbatlas

    test = mongodbatlas.get_online_archive(project_id=var["project_id"],
        cluster_name=var["cluster_name"],
        archive_id="5ebad3c1fe9c0ab8d37d61e1")
    ```
    ## Attributes reference

    * `db_name`          -  Name of the database that contains the collection.
    * `coll_name`        -  Name of the collection.
    * `criteria`         -  Criteria to use for archiving data.
    * `criteria.type`          - Type of criteria (DATE, CUSTOM)
    * `criteria.date_field`    - Name of an already indexed date field from the documents. Data is archived when the current date is greater than the value of the date field specified here plus the number of days specified via the `expire_after_days` parameter.
    * `criteria.date_format`   - the date format. Valid values:  ISODATE (default), EPOCH_SECONDS, EPOCH_MILLIS, EPOCH_NANOSECONDS
    * `criteria.expire_after_days` - Number of days that specifies the age limit for the data in the live Atlas cluster.
    * `criteria.query` - JSON query to use to select documents for archiving. Only for `CUSTOM` type
    * `partition_fields` -  Fields to use to partition data.
    * `partition_fields.field_name` - Name of the field. To specify a nested field, use the dot notation.
    * `partition_fields.order` - Position of the field in the partition. Value can be: 0,1,2
      By default, the date field specified in the criteria.dateField parameter is in the first position of the partition.
    * `partitio_fields.field_type` - Type of the partition field
    * `state`    - Status of the online archive. Valid values are: Pending, Archiving, Idle, Pausing, Paused, Orphaned and Deleted

    See [MongoDB Atlas API](https://docs.atlas.mongodb.com/reference/api/online-archive-get-one/) Documentation for more information.


    :param str archive_id: ID of the online archive.
    :param str cluster_name: Name of the cluster that contains the collection.
    :param str project_id: The unique ID for the project.
    """
    ...
