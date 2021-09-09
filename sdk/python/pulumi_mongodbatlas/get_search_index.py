# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'GetSearchIndexResult',
    'AwaitableGetSearchIndexResult',
    'get_search_index',
]

@pulumi.output_type
class GetSearchIndexResult:
    """
    A collection of values returned by getSearchIndex.
    """
    def __init__(__self__, analyzer=None, analyzers=None, cluster_name=None, collection_name=None, database=None, id=None, index_id=None, mappings_dynamic=None, mappings_fields=None, name=None, project_id=None, search_analyzer=None, status=None):
        if analyzer and not isinstance(analyzer, str):
            raise TypeError("Expected argument 'analyzer' to be a str")
        pulumi.set(__self__, "analyzer", analyzer)
        if analyzers and not isinstance(analyzers, str):
            raise TypeError("Expected argument 'analyzers' to be a str")
        pulumi.set(__self__, "analyzers", analyzers)
        if cluster_name and not isinstance(cluster_name, str):
            raise TypeError("Expected argument 'cluster_name' to be a str")
        pulumi.set(__self__, "cluster_name", cluster_name)
        if collection_name and not isinstance(collection_name, str):
            raise TypeError("Expected argument 'collection_name' to be a str")
        pulumi.set(__self__, "collection_name", collection_name)
        if database and not isinstance(database, str):
            raise TypeError("Expected argument 'database' to be a str")
        pulumi.set(__self__, "database", database)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if index_id and not isinstance(index_id, str):
            raise TypeError("Expected argument 'index_id' to be a str")
        pulumi.set(__self__, "index_id", index_id)
        if mappings_dynamic and not isinstance(mappings_dynamic, bool):
            raise TypeError("Expected argument 'mappings_dynamic' to be a bool")
        pulumi.set(__self__, "mappings_dynamic", mappings_dynamic)
        if mappings_fields and not isinstance(mappings_fields, str):
            raise TypeError("Expected argument 'mappings_fields' to be a str")
        pulumi.set(__self__, "mappings_fields", mappings_fields)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if project_id and not isinstance(project_id, str):
            raise TypeError("Expected argument 'project_id' to be a str")
        pulumi.set(__self__, "project_id", project_id)
        if search_analyzer and not isinstance(search_analyzer, str):
            raise TypeError("Expected argument 'search_analyzer' to be a str")
        pulumi.set(__self__, "search_analyzer", search_analyzer)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter
    def analyzer(self) -> Optional[str]:
        """
        [Analyzer](https://docs.atlas.mongodb.com/reference/atlas-search/analyzers/#std-label-analyzers-ref) to use when creating the index.
        """
        return pulumi.get(self, "analyzer")

    @property
    @pulumi.getter
    def analyzers(self) -> Optional[str]:
        """
        [Custom analyzers](https://docs.atlas.mongodb.com/reference/atlas-search/analyzers/custom/#std-label-custom-analyzers) to use in this index (this is an array of objects).
        """
        return pulumi.get(self, "analyzers")

    @property
    @pulumi.getter(name="clusterName")
    def cluster_name(self) -> str:
        return pulumi.get(self, "cluster_name")

    @property
    @pulumi.getter(name="collectionName")
    def collection_name(self) -> Optional[str]:
        """
        (Required) Name of the collection the index is on.
        """
        return pulumi.get(self, "collection_name")

    @property
    @pulumi.getter
    def database(self) -> Optional[str]:
        """
        (Required) Name of the database the collection is in.
        """
        return pulumi.get(self, "database")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="indexId")
    def index_id(self) -> str:
        return pulumi.get(self, "index_id")

    @property
    @pulumi.getter(name="mappingsDynamic")
    def mappings_dynamic(self) -> Optional[bool]:
        """
        Flag indicating whether the index uses dynamic or static mappings.
        """
        return pulumi.get(self, "mappings_dynamic")

    @property
    @pulumi.getter(name="mappingsFields")
    def mappings_fields(self) -> Optional[str]:
        """
        Object containing one or more field specifications.
        """
        return pulumi.get(self, "mappings_fields")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        Name of the index.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> str:
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter(name="searchAnalyzer")
    def search_analyzer(self) -> Optional[str]:
        """
        [Analyzer](https://docs.atlas.mongodb.com/reference/atlas-search/analyzers/#std-label-analyzers-ref) to use when searching the index.
        """
        return pulumi.get(self, "search_analyzer")

    @property
    @pulumi.getter
    def status(self) -> str:
        return pulumi.get(self, "status")


class AwaitableGetSearchIndexResult(GetSearchIndexResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSearchIndexResult(
            analyzer=self.analyzer,
            analyzers=self.analyzers,
            cluster_name=self.cluster_name,
            collection_name=self.collection_name,
            database=self.database,
            id=self.id,
            index_id=self.index_id,
            mappings_dynamic=self.mappings_dynamic,
            mappings_fields=self.mappings_fields,
            name=self.name,
            project_id=self.project_id,
            search_analyzer=self.search_analyzer,
            status=self.status)


def get_search_index(analyzer: Optional[str] = None,
                     analyzers: Optional[str] = None,
                     cluster_name: Optional[str] = None,
                     collection_name: Optional[str] = None,
                     database: Optional[str] = None,
                     index_id: Optional[str] = None,
                     mappings_dynamic: Optional[bool] = None,
                     mappings_fields: Optional[str] = None,
                     name: Optional[str] = None,
                     project_id: Optional[str] = None,
                     search_analyzer: Optional[str] = None,
                     status: Optional[str] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSearchIndexResult:
    """
    `SearchIndex` describe a single search indexes. This represents a single search index that have been created.

    > **NOTE:** Groups and projects are synonymous terms. You may find `groupId` in the official documentation.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_mongodbatlas as mongodbatlas

    test = mongodbatlas.get_search_index(cluster_name="<CLUSTER_NAME>",
        index_id="<INDEX_ID",
        project_id="<PROJECT_ID>")
    ```


    :param str analyzer: [Analyzer](https://docs.atlas.mongodb.com/reference/atlas-search/analyzers/#std-label-analyzers-ref) to use when creating the index.
    :param str analyzers: [Custom analyzers](https://docs.atlas.mongodb.com/reference/atlas-search/analyzers/custom/#std-label-custom-analyzers) to use in this index (this is an array of objects).
    :param str cluster_name: The name of the cluster containing the collection with one or more Atlas Search indexes.
    :param str collection_name: (Required) Name of the collection the index is on.
    :param str database: (Required) Name of the database the collection is in.
    :param str index_id: The unique identifier of the Atlas Search index. Use the `getSearchIndexes`datasource to find the IDs of all Atlas Search indexes.
    :param bool mappings_dynamic: Flag indicating whether the index uses dynamic or static mappings.
    :param str mappings_fields: Object containing one or more field specifications.
    :param str name: Name of the index.
    :param str project_id: The unique identifier for the [project](https://docs.atlas.mongodb.com/organizations-projects/#std-label-projects) that contains the specified cluster.
    :param str search_analyzer: [Analyzer](https://docs.atlas.mongodb.com/reference/atlas-search/analyzers/#std-label-analyzers-ref) to use when searching the index.
    """
    __args__ = dict()
    __args__['analyzer'] = analyzer
    __args__['analyzers'] = analyzers
    __args__['clusterName'] = cluster_name
    __args__['collectionName'] = collection_name
    __args__['database'] = database
    __args__['indexId'] = index_id
    __args__['mappingsDynamic'] = mappings_dynamic
    __args__['mappingsFields'] = mappings_fields
    __args__['name'] = name
    __args__['projectId'] = project_id
    __args__['searchAnalyzer'] = search_analyzer
    __args__['status'] = status
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('mongodbatlas:index/getSearchIndex:getSearchIndex', __args__, opts=opts, typ=GetSearchIndexResult).value

    return AwaitableGetSearchIndexResult(
        analyzer=__ret__.analyzer,
        analyzers=__ret__.analyzers,
        cluster_name=__ret__.cluster_name,
        collection_name=__ret__.collection_name,
        database=__ret__.database,
        id=__ret__.id,
        index_id=__ret__.index_id,
        mappings_dynamic=__ret__.mappings_dynamic,
        mappings_fields=__ret__.mappings_fields,
        name=__ret__.name,
        project_id=__ret__.project_id,
        search_analyzer=__ret__.search_analyzer,
        status=__ret__.status)
