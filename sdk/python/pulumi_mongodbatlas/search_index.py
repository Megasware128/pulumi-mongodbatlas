# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['SearchIndexArgs', 'SearchIndex']

@pulumi.input_type
class SearchIndexArgs:
    def __init__(__self__, *,
                 analyzer: pulumi.Input[str],
                 cluster_name: pulumi.Input[str],
                 collection_name: pulumi.Input[str],
                 database: pulumi.Input[str],
                 project_id: pulumi.Input[str],
                 analyzers: Optional[pulumi.Input[str]] = None,
                 mappings_dynamic: Optional[pulumi.Input[bool]] = None,
                 mappings_fields: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 search_analyzer: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a SearchIndex resource.
        :param pulumi.Input[str] analyzer: [Analyzer](https://docs.atlas.mongodb.com/reference/atlas-search/analyzers/#std-label-analyzers-ref) to use when creating the index. Defaults to [lucene.standard](https://docs.atlas.mongodb.com/reference/atlas-search/analyzers/standard/#std-label-ref-standard-analyzer)
        :param pulumi.Input[str] cluster_name: The name of the cluster where you want to create the search index within.
        :param pulumi.Input[str] collection_name: Name of the collection the index is on.
        :param pulumi.Input[str] database: Name of the database the collection is in.
        :param pulumi.Input[str] project_id: The ID of the organization or project you want to create the search index within.
        :param pulumi.Input[str] analyzers: [Custom analyzers](https://docs.atlas.mongodb.com/reference/atlas-search/analyzers/custom/#std-label-custom-analyzers) to use in this index. This is an array of JSON objects.
        :param pulumi.Input[bool] mappings_dynamic: Indicates whether the index uses dynamic or static mapping. For dynamic mapping, set the value to `true`. For static mapping, specify the fields to index using `mappings_fields`
        :param pulumi.Input[str] mappings_fields: attribute is required when `mappings_dynamic` is true. This field needs to be a JSON string in order to be decoded correctly.
        :param pulumi.Input[str] name: Name of the custom analyzer. Names must be unique within an index, and may **not** start with any of the following strings:
               * `lucene`
               * `builtin`
               * `mongodb`
        :param pulumi.Input[str] search_analyzer: [Analyzer](https://docs.atlas.mongodb.com/reference/atlas-search/analyzers/#std-label-analyzers-ref) to use when searching the index. Defaults to [lucene.standard](https://docs.atlas.mongodb.com/reference/atlas-search/analyzers/standard/#std-label-ref-standard-analyzer)
        """
        pulumi.set(__self__, "analyzer", analyzer)
        pulumi.set(__self__, "cluster_name", cluster_name)
        pulumi.set(__self__, "collection_name", collection_name)
        pulumi.set(__self__, "database", database)
        pulumi.set(__self__, "project_id", project_id)
        if analyzers is not None:
            pulumi.set(__self__, "analyzers", analyzers)
        if mappings_dynamic is not None:
            pulumi.set(__self__, "mappings_dynamic", mappings_dynamic)
        if mappings_fields is not None:
            pulumi.set(__self__, "mappings_fields", mappings_fields)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if search_analyzer is not None:
            pulumi.set(__self__, "search_analyzer", search_analyzer)
        if status is not None:
            pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter
    def analyzer(self) -> pulumi.Input[str]:
        """
        [Analyzer](https://docs.atlas.mongodb.com/reference/atlas-search/analyzers/#std-label-analyzers-ref) to use when creating the index. Defaults to [lucene.standard](https://docs.atlas.mongodb.com/reference/atlas-search/analyzers/standard/#std-label-ref-standard-analyzer)
        """
        return pulumi.get(self, "analyzer")

    @analyzer.setter
    def analyzer(self, value: pulumi.Input[str]):
        pulumi.set(self, "analyzer", value)

    @property
    @pulumi.getter(name="clusterName")
    def cluster_name(self) -> pulumi.Input[str]:
        """
        The name of the cluster where you want to create the search index within.
        """
        return pulumi.get(self, "cluster_name")

    @cluster_name.setter
    def cluster_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "cluster_name", value)

    @property
    @pulumi.getter(name="collectionName")
    def collection_name(self) -> pulumi.Input[str]:
        """
        Name of the collection the index is on.
        """
        return pulumi.get(self, "collection_name")

    @collection_name.setter
    def collection_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "collection_name", value)

    @property
    @pulumi.getter
    def database(self) -> pulumi.Input[str]:
        """
        Name of the database the collection is in.
        """
        return pulumi.get(self, "database")

    @database.setter
    def database(self, value: pulumi.Input[str]):
        pulumi.set(self, "database", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Input[str]:
        """
        The ID of the organization or project you want to create the search index within.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter
    def analyzers(self) -> Optional[pulumi.Input[str]]:
        """
        [Custom analyzers](https://docs.atlas.mongodb.com/reference/atlas-search/analyzers/custom/#std-label-custom-analyzers) to use in this index. This is an array of JSON objects.
        """
        return pulumi.get(self, "analyzers")

    @analyzers.setter
    def analyzers(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "analyzers", value)

    @property
    @pulumi.getter(name="mappingsDynamic")
    def mappings_dynamic(self) -> Optional[pulumi.Input[bool]]:
        """
        Indicates whether the index uses dynamic or static mapping. For dynamic mapping, set the value to `true`. For static mapping, specify the fields to index using `mappings_fields`
        """
        return pulumi.get(self, "mappings_dynamic")

    @mappings_dynamic.setter
    def mappings_dynamic(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "mappings_dynamic", value)

    @property
    @pulumi.getter(name="mappingsFields")
    def mappings_fields(self) -> Optional[pulumi.Input[str]]:
        """
        attribute is required when `mappings_dynamic` is true. This field needs to be a JSON string in order to be decoded correctly.
        """
        return pulumi.get(self, "mappings_fields")

    @mappings_fields.setter
    def mappings_fields(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "mappings_fields", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the custom analyzer. Names must be unique within an index, and may **not** start with any of the following strings:
        * `lucene`
        * `builtin`
        * `mongodb`
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="searchAnalyzer")
    def search_analyzer(self) -> Optional[pulumi.Input[str]]:
        """
        [Analyzer](https://docs.atlas.mongodb.com/reference/atlas-search/analyzers/#std-label-analyzers-ref) to use when searching the index. Defaults to [lucene.standard](https://docs.atlas.mongodb.com/reference/atlas-search/analyzers/standard/#std-label-ref-standard-analyzer)
        """
        return pulumi.get(self, "search_analyzer")

    @search_analyzer.setter
    def search_analyzer(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "search_analyzer", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)


@pulumi.input_type
class _SearchIndexState:
    def __init__(__self__, *,
                 analyzer: Optional[pulumi.Input[str]] = None,
                 analyzers: Optional[pulumi.Input[str]] = None,
                 cluster_name: Optional[pulumi.Input[str]] = None,
                 collection_name: Optional[pulumi.Input[str]] = None,
                 database: Optional[pulumi.Input[str]] = None,
                 index_id: Optional[pulumi.Input[str]] = None,
                 mappings_dynamic: Optional[pulumi.Input[bool]] = None,
                 mappings_fields: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 search_analyzer: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering SearchIndex resources.
        :param pulumi.Input[str] analyzer: [Analyzer](https://docs.atlas.mongodb.com/reference/atlas-search/analyzers/#std-label-analyzers-ref) to use when creating the index. Defaults to [lucene.standard](https://docs.atlas.mongodb.com/reference/atlas-search/analyzers/standard/#std-label-ref-standard-analyzer)
        :param pulumi.Input[str] analyzers: [Custom analyzers](https://docs.atlas.mongodb.com/reference/atlas-search/analyzers/custom/#std-label-custom-analyzers) to use in this index. This is an array of JSON objects.
        :param pulumi.Input[str] cluster_name: The name of the cluster where you want to create the search index within.
        :param pulumi.Input[str] collection_name: Name of the collection the index is on.
        :param pulumi.Input[str] database: Name of the database the collection is in.
        :param pulumi.Input[bool] mappings_dynamic: Indicates whether the index uses dynamic or static mapping. For dynamic mapping, set the value to `true`. For static mapping, specify the fields to index using `mappings_fields`
        :param pulumi.Input[str] mappings_fields: attribute is required when `mappings_dynamic` is true. This field needs to be a JSON string in order to be decoded correctly.
        :param pulumi.Input[str] name: Name of the custom analyzer. Names must be unique within an index, and may **not** start with any of the following strings:
               * `lucene`
               * `builtin`
               * `mongodb`
        :param pulumi.Input[str] project_id: The ID of the organization or project you want to create the search index within.
        :param pulumi.Input[str] search_analyzer: [Analyzer](https://docs.atlas.mongodb.com/reference/atlas-search/analyzers/#std-label-analyzers-ref) to use when searching the index. Defaults to [lucene.standard](https://docs.atlas.mongodb.com/reference/atlas-search/analyzers/standard/#std-label-ref-standard-analyzer)
        """
        if analyzer is not None:
            pulumi.set(__self__, "analyzer", analyzer)
        if analyzers is not None:
            pulumi.set(__self__, "analyzers", analyzers)
        if cluster_name is not None:
            pulumi.set(__self__, "cluster_name", cluster_name)
        if collection_name is not None:
            pulumi.set(__self__, "collection_name", collection_name)
        if database is not None:
            pulumi.set(__self__, "database", database)
        if index_id is not None:
            pulumi.set(__self__, "index_id", index_id)
        if mappings_dynamic is not None:
            pulumi.set(__self__, "mappings_dynamic", mappings_dynamic)
        if mappings_fields is not None:
            pulumi.set(__self__, "mappings_fields", mappings_fields)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if search_analyzer is not None:
            pulumi.set(__self__, "search_analyzer", search_analyzer)
        if status is not None:
            pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter
    def analyzer(self) -> Optional[pulumi.Input[str]]:
        """
        [Analyzer](https://docs.atlas.mongodb.com/reference/atlas-search/analyzers/#std-label-analyzers-ref) to use when creating the index. Defaults to [lucene.standard](https://docs.atlas.mongodb.com/reference/atlas-search/analyzers/standard/#std-label-ref-standard-analyzer)
        """
        return pulumi.get(self, "analyzer")

    @analyzer.setter
    def analyzer(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "analyzer", value)

    @property
    @pulumi.getter
    def analyzers(self) -> Optional[pulumi.Input[str]]:
        """
        [Custom analyzers](https://docs.atlas.mongodb.com/reference/atlas-search/analyzers/custom/#std-label-custom-analyzers) to use in this index. This is an array of JSON objects.
        """
        return pulumi.get(self, "analyzers")

    @analyzers.setter
    def analyzers(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "analyzers", value)

    @property
    @pulumi.getter(name="clusterName")
    def cluster_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the cluster where you want to create the search index within.
        """
        return pulumi.get(self, "cluster_name")

    @cluster_name.setter
    def cluster_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cluster_name", value)

    @property
    @pulumi.getter(name="collectionName")
    def collection_name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the collection the index is on.
        """
        return pulumi.get(self, "collection_name")

    @collection_name.setter
    def collection_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "collection_name", value)

    @property
    @pulumi.getter
    def database(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the database the collection is in.
        """
        return pulumi.get(self, "database")

    @database.setter
    def database(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "database", value)

    @property
    @pulumi.getter(name="indexId")
    def index_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "index_id")

    @index_id.setter
    def index_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "index_id", value)

    @property
    @pulumi.getter(name="mappingsDynamic")
    def mappings_dynamic(self) -> Optional[pulumi.Input[bool]]:
        """
        Indicates whether the index uses dynamic or static mapping. For dynamic mapping, set the value to `true`. For static mapping, specify the fields to index using `mappings_fields`
        """
        return pulumi.get(self, "mappings_dynamic")

    @mappings_dynamic.setter
    def mappings_dynamic(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "mappings_dynamic", value)

    @property
    @pulumi.getter(name="mappingsFields")
    def mappings_fields(self) -> Optional[pulumi.Input[str]]:
        """
        attribute is required when `mappings_dynamic` is true. This field needs to be a JSON string in order to be decoded correctly.
        """
        return pulumi.get(self, "mappings_fields")

    @mappings_fields.setter
    def mappings_fields(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "mappings_fields", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the custom analyzer. Names must be unique within an index, and may **not** start with any of the following strings:
        * `lucene`
        * `builtin`
        * `mongodb`
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the organization or project you want to create the search index within.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter(name="searchAnalyzer")
    def search_analyzer(self) -> Optional[pulumi.Input[str]]:
        """
        [Analyzer](https://docs.atlas.mongodb.com/reference/atlas-search/analyzers/#std-label-analyzers-ref) to use when searching the index. Defaults to [lucene.standard](https://docs.atlas.mongodb.com/reference/atlas-search/analyzers/standard/#std-label-ref-standard-analyzer)
        """
        return pulumi.get(self, "search_analyzer")

    @search_analyzer.setter
    def search_analyzer(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "search_analyzer", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)


class SearchIndex(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 analyzer: Optional[pulumi.Input[str]] = None,
                 analyzers: Optional[pulumi.Input[str]] = None,
                 cluster_name: Optional[pulumi.Input[str]] = None,
                 collection_name: Optional[pulumi.Input[str]] = None,
                 database: Optional[pulumi.Input[str]] = None,
                 mappings_dynamic: Optional[pulumi.Input[bool]] = None,
                 mappings_fields: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 search_analyzer: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        `SearchIndex` provides a Search Index resource. This allows indexes to be created.

        ## Example Usage
        ### Basic
        ```python
        import pulumi
        import pulumi_mongodbatlas as mongodbatlas

        test = mongodbatlas.SearchIndex("test",
            analyzer="lucene.standard",
            cluster_name="<CLUSTER_NAME>",
            collection_name="collection_test",
            database="database_test",
            mappings_dynamic=True,
            project_id="<PROJECT_ID>",
            search_analyzer="lucene.standard")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] analyzer: [Analyzer](https://docs.atlas.mongodb.com/reference/atlas-search/analyzers/#std-label-analyzers-ref) to use when creating the index. Defaults to [lucene.standard](https://docs.atlas.mongodb.com/reference/atlas-search/analyzers/standard/#std-label-ref-standard-analyzer)
        :param pulumi.Input[str] analyzers: [Custom analyzers](https://docs.atlas.mongodb.com/reference/atlas-search/analyzers/custom/#std-label-custom-analyzers) to use in this index. This is an array of JSON objects.
        :param pulumi.Input[str] cluster_name: The name of the cluster where you want to create the search index within.
        :param pulumi.Input[str] collection_name: Name of the collection the index is on.
        :param pulumi.Input[str] database: Name of the database the collection is in.
        :param pulumi.Input[bool] mappings_dynamic: Indicates whether the index uses dynamic or static mapping. For dynamic mapping, set the value to `true`. For static mapping, specify the fields to index using `mappings_fields`
        :param pulumi.Input[str] mappings_fields: attribute is required when `mappings_dynamic` is true. This field needs to be a JSON string in order to be decoded correctly.
        :param pulumi.Input[str] name: Name of the custom analyzer. Names must be unique within an index, and may **not** start with any of the following strings:
               * `lucene`
               * `builtin`
               * `mongodb`
        :param pulumi.Input[str] project_id: The ID of the organization or project you want to create the search index within.
        :param pulumi.Input[str] search_analyzer: [Analyzer](https://docs.atlas.mongodb.com/reference/atlas-search/analyzers/#std-label-analyzers-ref) to use when searching the index. Defaults to [lucene.standard](https://docs.atlas.mongodb.com/reference/atlas-search/analyzers/standard/#std-label-ref-standard-analyzer)
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SearchIndexArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        `SearchIndex` provides a Search Index resource. This allows indexes to be created.

        ## Example Usage
        ### Basic
        ```python
        import pulumi
        import pulumi_mongodbatlas as mongodbatlas

        test = mongodbatlas.SearchIndex("test",
            analyzer="lucene.standard",
            cluster_name="<CLUSTER_NAME>",
            collection_name="collection_test",
            database="database_test",
            mappings_dynamic=True,
            project_id="<PROJECT_ID>",
            search_analyzer="lucene.standard")
        ```

        :param str resource_name: The name of the resource.
        :param SearchIndexArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SearchIndexArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 analyzer: Optional[pulumi.Input[str]] = None,
                 analyzers: Optional[pulumi.Input[str]] = None,
                 cluster_name: Optional[pulumi.Input[str]] = None,
                 collection_name: Optional[pulumi.Input[str]] = None,
                 database: Optional[pulumi.Input[str]] = None,
                 mappings_dynamic: Optional[pulumi.Input[bool]] = None,
                 mappings_fields: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 search_analyzer: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SearchIndexArgs.__new__(SearchIndexArgs)

            if analyzer is None and not opts.urn:
                raise TypeError("Missing required property 'analyzer'")
            __props__.__dict__["analyzer"] = analyzer
            __props__.__dict__["analyzers"] = analyzers
            if cluster_name is None and not opts.urn:
                raise TypeError("Missing required property 'cluster_name'")
            __props__.__dict__["cluster_name"] = cluster_name
            if collection_name is None and not opts.urn:
                raise TypeError("Missing required property 'collection_name'")
            __props__.__dict__["collection_name"] = collection_name
            if database is None and not opts.urn:
                raise TypeError("Missing required property 'database'")
            __props__.__dict__["database"] = database
            __props__.__dict__["mappings_dynamic"] = mappings_dynamic
            __props__.__dict__["mappings_fields"] = mappings_fields
            __props__.__dict__["name"] = name
            if project_id is None and not opts.urn:
                raise TypeError("Missing required property 'project_id'")
            __props__.__dict__["project_id"] = project_id
            __props__.__dict__["search_analyzer"] = search_analyzer
            __props__.__dict__["status"] = status
            __props__.__dict__["index_id"] = None
        super(SearchIndex, __self__).__init__(
            'mongodbatlas:index/searchIndex:SearchIndex',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            analyzer: Optional[pulumi.Input[str]] = None,
            analyzers: Optional[pulumi.Input[str]] = None,
            cluster_name: Optional[pulumi.Input[str]] = None,
            collection_name: Optional[pulumi.Input[str]] = None,
            database: Optional[pulumi.Input[str]] = None,
            index_id: Optional[pulumi.Input[str]] = None,
            mappings_dynamic: Optional[pulumi.Input[bool]] = None,
            mappings_fields: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            project_id: Optional[pulumi.Input[str]] = None,
            search_analyzer: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None) -> 'SearchIndex':
        """
        Get an existing SearchIndex resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] analyzer: [Analyzer](https://docs.atlas.mongodb.com/reference/atlas-search/analyzers/#std-label-analyzers-ref) to use when creating the index. Defaults to [lucene.standard](https://docs.atlas.mongodb.com/reference/atlas-search/analyzers/standard/#std-label-ref-standard-analyzer)
        :param pulumi.Input[str] analyzers: [Custom analyzers](https://docs.atlas.mongodb.com/reference/atlas-search/analyzers/custom/#std-label-custom-analyzers) to use in this index. This is an array of JSON objects.
        :param pulumi.Input[str] cluster_name: The name of the cluster where you want to create the search index within.
        :param pulumi.Input[str] collection_name: Name of the collection the index is on.
        :param pulumi.Input[str] database: Name of the database the collection is in.
        :param pulumi.Input[bool] mappings_dynamic: Indicates whether the index uses dynamic or static mapping. For dynamic mapping, set the value to `true`. For static mapping, specify the fields to index using `mappings_fields`
        :param pulumi.Input[str] mappings_fields: attribute is required when `mappings_dynamic` is true. This field needs to be a JSON string in order to be decoded correctly.
        :param pulumi.Input[str] name: Name of the custom analyzer. Names must be unique within an index, and may **not** start with any of the following strings:
               * `lucene`
               * `builtin`
               * `mongodb`
        :param pulumi.Input[str] project_id: The ID of the organization or project you want to create the search index within.
        :param pulumi.Input[str] search_analyzer: [Analyzer](https://docs.atlas.mongodb.com/reference/atlas-search/analyzers/#std-label-analyzers-ref) to use when searching the index. Defaults to [lucene.standard](https://docs.atlas.mongodb.com/reference/atlas-search/analyzers/standard/#std-label-ref-standard-analyzer)
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SearchIndexState.__new__(_SearchIndexState)

        __props__.__dict__["analyzer"] = analyzer
        __props__.__dict__["analyzers"] = analyzers
        __props__.__dict__["cluster_name"] = cluster_name
        __props__.__dict__["collection_name"] = collection_name
        __props__.__dict__["database"] = database
        __props__.__dict__["index_id"] = index_id
        __props__.__dict__["mappings_dynamic"] = mappings_dynamic
        __props__.__dict__["mappings_fields"] = mappings_fields
        __props__.__dict__["name"] = name
        __props__.__dict__["project_id"] = project_id
        __props__.__dict__["search_analyzer"] = search_analyzer
        __props__.__dict__["status"] = status
        return SearchIndex(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def analyzer(self) -> pulumi.Output[str]:
        """
        [Analyzer](https://docs.atlas.mongodb.com/reference/atlas-search/analyzers/#std-label-analyzers-ref) to use when creating the index. Defaults to [lucene.standard](https://docs.atlas.mongodb.com/reference/atlas-search/analyzers/standard/#std-label-ref-standard-analyzer)
        """
        return pulumi.get(self, "analyzer")

    @property
    @pulumi.getter
    def analyzers(self) -> pulumi.Output[Optional[str]]:
        """
        [Custom analyzers](https://docs.atlas.mongodb.com/reference/atlas-search/analyzers/custom/#std-label-custom-analyzers) to use in this index. This is an array of JSON objects.
        """
        return pulumi.get(self, "analyzers")

    @property
    @pulumi.getter(name="clusterName")
    def cluster_name(self) -> pulumi.Output[str]:
        """
        The name of the cluster where you want to create the search index within.
        """
        return pulumi.get(self, "cluster_name")

    @property
    @pulumi.getter(name="collectionName")
    def collection_name(self) -> pulumi.Output[str]:
        """
        Name of the collection the index is on.
        """
        return pulumi.get(self, "collection_name")

    @property
    @pulumi.getter
    def database(self) -> pulumi.Output[str]:
        """
        Name of the database the collection is in.
        """
        return pulumi.get(self, "database")

    @property
    @pulumi.getter(name="indexId")
    def index_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "index_id")

    @property
    @pulumi.getter(name="mappingsDynamic")
    def mappings_dynamic(self) -> pulumi.Output[Optional[bool]]:
        """
        Indicates whether the index uses dynamic or static mapping. For dynamic mapping, set the value to `true`. For static mapping, specify the fields to index using `mappings_fields`
        """
        return pulumi.get(self, "mappings_dynamic")

    @property
    @pulumi.getter(name="mappingsFields")
    def mappings_fields(self) -> pulumi.Output[Optional[str]]:
        """
        attribute is required when `mappings_dynamic` is true. This field needs to be a JSON string in order to be decoded correctly.
        """
        return pulumi.get(self, "mappings_fields")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the custom analyzer. Names must be unique within an index, and may **not** start with any of the following strings:
        * `lucene`
        * `builtin`
        * `mongodb`
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Output[str]:
        """
        The ID of the organization or project you want to create the search index within.
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter(name="searchAnalyzer")
    def search_analyzer(self) -> pulumi.Output[Optional[str]]:
        """
        [Analyzer](https://docs.atlas.mongodb.com/reference/atlas-search/analyzers/#std-label-analyzers-ref) to use when searching the index. Defaults to [lucene.standard](https://docs.atlas.mongodb.com/reference/atlas-search/analyzers/standard/#std-label-ref-standard-analyzer)
        """
        return pulumi.get(self, "search_analyzer")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        return pulumi.get(self, "status")

