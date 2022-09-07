// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Mongodbatlas
{
    public static class GetDataLake
    {
        /// <summary>
        /// `mongodbatlas.DataLake` describe a Data Lake.
        /// 
        /// &gt; **NOTE:** Groups and projects are synonymous terms. You may find group_id in the official documentation.
        /// </summary>
        public static Task<GetDataLakeResult> InvokeAsync(GetDataLakeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetDataLakeResult>("mongodbatlas:index/getDataLake:getDataLake", args ?? new GetDataLakeArgs(), options.WithDefaults());

        /// <summary>
        /// `mongodbatlas.DataLake` describe a Data Lake.
        /// 
        /// &gt; **NOTE:** Groups and projects are synonymous terms. You may find group_id in the official documentation.
        /// </summary>
        public static Output<GetDataLakeResult> Invoke(GetDataLakeInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetDataLakeResult>("mongodbatlas:index/getDataLake:getDataLake", args ?? new GetDataLakeInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDataLakeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the data lake.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The unique ID for the project to create a data lake.
        /// </summary>
        [Input("projectId", required: true)]
        public string ProjectId { get; set; } = null!;

        public GetDataLakeArgs()
        {
        }
        public static new GetDataLakeArgs Empty => new GetDataLakeArgs();
    }

    public sealed class GetDataLakeInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the data lake.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The unique ID for the project to create a data lake.
        /// </summary>
        [Input("projectId", required: true)]
        public Input<string> ProjectId { get; set; } = null!;

        public GetDataLakeInvokeArgs()
        {
        }
        public static new GetDataLakeInvokeArgs Empty => new GetDataLakeInvokeArgs();
    }


    [OutputType]
    public sealed class GetDataLakeResult
    {
        /// <summary>
        /// AWS provider of the cloud service where Data Lake can access the S3 Bucket.
        /// * `aws.0.role_id` - Unique identifier of the role that Data Lake can use to access the data stores.
        /// * `aws.0.test_s3_bucket` - Name of the S3 data bucket that the provided role ID is authorized to access.
        /// * `aws.0.role_id` - Unique identifier of the role that Data Lake can use to access the data stores.
        /// * `aws.0.test_s3_bucket` - Name of the S3 data bucket that the provided role ID is authorized to access.
        /// * `aws.0.iam_assumed_role_arn` - Amazon Resource Name (ARN) of the IAM Role that Data Lake assumes when accessing S3 Bucket data stores.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetDataLakeAwResult> Aws;
        /// <summary>
        /// The cloud provider region to which Atlas Data Lake routes client connections for data processing.
        /// * `data_process_region.0.cloud_provider` - Name of the cloud service provider.
        /// * `data_process_region.0.region` -Name of the region to which Data Lake routes client connections for data processing.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetDataLakeDataProcessRegionResult> DataProcessRegions;
        /// <summary>
        /// The list of hostnames assigned to the Atlas Data Lake. Each string in the array is a hostname assigned to the Atlas Data Lake.
        /// </summary>
        public readonly ImmutableArray<string> Hostnames;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Name;
        public readonly string ProjectId;
        /// <summary>
        /// Current state of the Atlas Data Lake:
        /// </summary>
        public readonly string State;
        /// <summary>
        /// Configuration details for mapping each data store to queryable databases and collections.
        /// * `storage_databases.#.name` - Name of the database to which Data Lake maps the data contained in the data store.
        /// * `storage_databases.#.collections` -     Array of objects where each object represents a collection and data sources that map to a [stores](https://docs.mongodb.com/datalake/reference/format/data-lake-configuration#mongodb-datalakeconf-datalakeconf.stores) data store.
        /// * `storage_databases.#.collections.#.name` - Name of the collection.
        /// * `storage_databases.#.collections.#.data_sources` -     Array of objects where each object represents a stores data store to map with the collection.
        /// * `storage_databases.#.collections.#.data_sources.#.store_name` -     Name of a data store to map to the `&lt;collection&gt;`.
        /// * `storage_databases.#.collections.#.data_sources.#.default_format` - Default format that Data Lake assumes if it encounters a file without an extension while searching the storeName.
        /// * `storage_databases.#.collections.#.data_sources.#.path` - Controls how Atlas Data Lake searches for and parses files in the storeName before mapping them to the `&lt;collection&gt;`.
        /// * `storage_databases.#.views` -     Array of objects where each object represents an [aggregation pipeline](https://docs.mongodb.com/manual/core/aggregation-pipeline/#id1) on a collection.
        /// * `storage_databases.#.views.#.name` - Name of the view.
        /// * `storage_databases.#.views.#.source` -  Name of the source collection for the view.
        /// * `storage_databases.#.views.#.pipeline`- Aggregation pipeline stage(s) to apply to the source collection.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetDataLakeStorageDatabaseResult> StorageDatabases;
        /// <summary>
        /// Each object in the array represents a data store. Data Lake uses the storage.databases configuration details to map data in each data store to queryable databases and collections.
        /// * `storage_stores.#.name` - Name of the data store.
        /// * `storage_stores.#.provider` - Defines where the data is stored.
        /// * `storage_stores.#.region` - Name of the AWS region in which the S3 bucket is hosted.
        /// * `storage_stores.#.bucket` - Name of the AWS S3 bucket.
        /// * `storage_stores.#.prefix` - Prefix Data Lake applies when searching for files in the S3 bucket .
        /// * `storage_stores.#.delimiter` - The delimiter that separates `storage_databases.#.collections.#.data_sources.#.path` segments in the data store.
        /// * `storage_stores.#.include_tags` - Determines whether or not to use S3 tags on the files in the given path as additional partition attributes.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetDataLakeStorageStoreResult> StorageStores;

        [OutputConstructor]
        private GetDataLakeResult(
            ImmutableArray<Outputs.GetDataLakeAwResult> aws,

            ImmutableArray<Outputs.GetDataLakeDataProcessRegionResult> dataProcessRegions,

            ImmutableArray<string> hostnames,

            string id,

            string name,

            string projectId,

            string state,

            ImmutableArray<Outputs.GetDataLakeStorageDatabaseResult> storageDatabases,

            ImmutableArray<Outputs.GetDataLakeStorageStoreResult> storageStores)
        {
            Aws = aws;
            DataProcessRegions = dataProcessRegions;
            Hostnames = hostnames;
            Id = id;
            Name = name;
            ProjectId = projectId;
            State = state;
            StorageDatabases = storageDatabases;
            StorageStores = storageStores;
        }
    }
}
