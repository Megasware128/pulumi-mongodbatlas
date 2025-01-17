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
    'GetDataLakeResult',
    'AwaitableGetDataLakeResult',
    'get_data_lake',
    'get_data_lake_output',
]

@pulumi.output_type
class GetDataLakeResult:
    """
    A collection of values returned by getDataLake.
    """
    def __init__(__self__, aws=None, data_process_regions=None, hostnames=None, id=None, name=None, project_id=None, state=None, storage_databases=None, storage_stores=None):
        if aws and not isinstance(aws, list):
            raise TypeError("Expected argument 'aws' to be a list")
        pulumi.set(__self__, "aws", aws)
        if data_process_regions and not isinstance(data_process_regions, list):
            raise TypeError("Expected argument 'data_process_regions' to be a list")
        pulumi.set(__self__, "data_process_regions", data_process_regions)
        if hostnames and not isinstance(hostnames, list):
            raise TypeError("Expected argument 'hostnames' to be a list")
        pulumi.set(__self__, "hostnames", hostnames)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if project_id and not isinstance(project_id, str):
            raise TypeError("Expected argument 'project_id' to be a str")
        pulumi.set(__self__, "project_id", project_id)
        if state and not isinstance(state, str):
            raise TypeError("Expected argument 'state' to be a str")
        pulumi.set(__self__, "state", state)
        if storage_databases and not isinstance(storage_databases, list):
            raise TypeError("Expected argument 'storage_databases' to be a list")
        pulumi.set(__self__, "storage_databases", storage_databases)
        if storage_stores and not isinstance(storage_stores, list):
            raise TypeError("Expected argument 'storage_stores' to be a list")
        pulumi.set(__self__, "storage_stores", storage_stores)

    @property
    @pulumi.getter
    def aws(self) -> Sequence['outputs.GetDataLakeAwResult']:
        """
        AWS provider of the cloud service where Data Lake can access the S3 Bucket.
        * `aws.0.role_id` - Unique identifier of the role that Data Lake can use to access the data stores.
        * `aws.0.test_s3_bucket` - Name of the S3 data bucket that the provided role ID is authorized to access.
        * `aws.0.role_id` - Unique identifier of the role that Data Lake can use to access the data stores.
        * `aws.0.test_s3_bucket` - Name of the S3 data bucket that the provided role ID is authorized to access.
        * `aws.0.iam_assumed_role_arn` - Amazon Resource Name (ARN) of the IAM Role that Data Lake assumes when accessing S3 Bucket data stores.
        """
        return pulumi.get(self, "aws")

    @property
    @pulumi.getter(name="dataProcessRegions")
    def data_process_regions(self) -> Sequence['outputs.GetDataLakeDataProcessRegionResult']:
        """
        The cloud provider region to which Atlas Data Lake routes client connections for data processing.
        * `data_process_region.0.cloud_provider` - Name of the cloud service provider.
        * `data_process_region.0.region` -Name of the region to which Data Lake routes client connections for data processing.
        """
        return pulumi.get(self, "data_process_regions")

    @property
    @pulumi.getter
    def hostnames(self) -> Sequence[str]:
        """
        The list of hostnames assigned to the Atlas Data Lake. Each string in the array is a hostname assigned to the Atlas Data Lake.
        """
        return pulumi.get(self, "hostnames")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> str:
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter
    def state(self) -> str:
        """
        Current state of the Atlas Data Lake:
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter(name="storageDatabases")
    def storage_databases(self) -> Sequence['outputs.GetDataLakeStorageDatabaseResult']:
        """
        Configuration details for mapping each data store to queryable databases and collections.
        * `storage_databases.#.name` - Name of the database to which Data Lake maps the data contained in the data store.
        * `storage_databases.#.collections` -     Array of objects where each object represents a collection and data sources that map to a [stores](https://docs.mongodb.com/datalake/reference/format/data-lake-configuration#mongodb-datalakeconf-datalakeconf.stores) data store.
        * `storage_databases.#.collections.#.name` - Name of the collection.
        * `storage_databases.#.collections.#.data_sources` -     Array of objects where each object represents a stores data store to map with the collection.
        * `storage_databases.#.collections.#.data_sources.#.store_name` -     Name of a data store to map to the `<collection>`.
        * `storage_databases.#.collections.#.data_sources.#.default_format` - Default format that Data Lake assumes if it encounters a file without an extension while searching the storeName.
        * `storage_databases.#.collections.#.data_sources.#.path` - Controls how Atlas Data Lake searches for and parses files in the storeName before mapping them to the `<collection>`.
        * `storage_databases.#.views` -     Array of objects where each object represents an [aggregation pipeline](https://docs.mongodb.com/manual/core/aggregation-pipeline/#id1) on a collection.
        * `storage_databases.#.views.#.name` - Name of the view.
        * `storage_databases.#.views.#.source` -  Name of the source collection for the view.
        * `storage_databases.#.views.#.pipeline`- Aggregation pipeline stage(s) to apply to the source collection.
        """
        return pulumi.get(self, "storage_databases")

    @property
    @pulumi.getter(name="storageStores")
    def storage_stores(self) -> Sequence['outputs.GetDataLakeStorageStoreResult']:
        """
        Each object in the array represents a data store. Data Lake uses the storage.databases configuration details to map data in each data store to queryable databases and collections.
        * `storage_stores.#.name` - Name of the data store.
        * `storage_stores.#.provider` - Defines where the data is stored.
        * `storage_stores.#.region` - Name of the AWS region in which the S3 bucket is hosted.
        * `storage_stores.#.bucket` - Name of the AWS S3 bucket.
        * `storage_stores.#.prefix` - Prefix Data Lake applies when searching for files in the S3 bucket .
        * `storage_stores.#.delimiter` - The delimiter that separates `storage_databases.#.collections.#.data_sources.#.path` segments in the data store.
        * `storage_stores.#.include_tags` - Determines whether or not to use S3 tags on the files in the given path as additional partition attributes.
        """
        return pulumi.get(self, "storage_stores")


class AwaitableGetDataLakeResult(GetDataLakeResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDataLakeResult(
            aws=self.aws,
            data_process_regions=self.data_process_regions,
            hostnames=self.hostnames,
            id=self.id,
            name=self.name,
            project_id=self.project_id,
            state=self.state,
            storage_databases=self.storage_databases,
            storage_stores=self.storage_stores)


def get_data_lake(name: Optional[str] = None,
                  project_id: Optional[str] = None,
                  opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDataLakeResult:
    """
    `DataLake` describe a Data Lake.

    > **NOTE:** Groups and projects are synonymous terms. You may find group_id in the official documentation.


    :param str name: Name of the data lake.
    :param str project_id: The unique ID for the project to create a data lake.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['projectId'] = project_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('mongodbatlas:index/getDataLake:getDataLake', __args__, opts=opts, typ=GetDataLakeResult).value

    return AwaitableGetDataLakeResult(
        aws=__ret__.aws,
        data_process_regions=__ret__.data_process_regions,
        hostnames=__ret__.hostnames,
        id=__ret__.id,
        name=__ret__.name,
        project_id=__ret__.project_id,
        state=__ret__.state,
        storage_databases=__ret__.storage_databases,
        storage_stores=__ret__.storage_stores)


@_utilities.lift_output_func(get_data_lake)
def get_data_lake_output(name: Optional[pulumi.Input[str]] = None,
                         project_id: Optional[pulumi.Input[str]] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetDataLakeResult]:
    """
    `DataLake` describe a Data Lake.

    > **NOTE:** Groups and projects are synonymous terms. You may find group_id in the official documentation.


    :param str name: Name of the data lake.
    :param str project_id: The unique ID for the project to create a data lake.
    """
    ...
