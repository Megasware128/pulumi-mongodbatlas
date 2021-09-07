// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Mongodbatlas
{
    public static class GetSearchIndexes
    {
        /// <summary>
        /// `mongodbatlas.getSearchIndexes` describe all search indexes. This represents search indexes that have been created.
        /// 
        /// &gt; **NOTE:** Groups and projects are synonymous terms. You may find `groupId` in the official documentation.
        /// </summary>
        public static Task<GetSearchIndexesResult> InvokeAsync(GetSearchIndexesArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetSearchIndexesResult>("mongodbatlas:index/getSearchIndexes:getSearchIndexes", args ?? new GetSearchIndexesArgs(), options.WithVersion());
    }


    public sealed class GetSearchIndexesArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the cluster containing the collection with one or more Atlas Search indexes.
        /// </summary>
        [Input("clusterName", required: true)]
        public string ClusterName { get; set; } = null!;

        /// <summary>
        /// Name of the collection with one or more Atlas Search indexes.
        /// </summary>
        [Input("collectionName", required: true)]
        public string CollectionName { get; set; } = null!;

        /// <summary>
        /// (Required) Name of the database the collection is in.
        /// </summary>
        [Input("database", required: true)]
        public string Database { get; set; } = null!;

        /// <summary>
        /// Number of items that Atlas returns per page, up to a maximum of 500.
        /// </summary>
        [Input("itemsPerPage")]
        public int? ItemsPerPage { get; set; }

        /// <summary>
        /// Page number, starting with one, that Atlas returns of the total number of objects.
        /// </summary>
        [Input("pageNum")]
        public int? PageNum { get; set; }

        /// <summary>
        /// Unique identifier for the [project](https://docs.atlas.mongodb.com/organizations-projects/#std-label-projects) that contains the specified cluster.
        /// </summary>
        [Input("projectId", required: true)]
        public string ProjectId { get; set; } = null!;

        public GetSearchIndexesArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetSearchIndexesResult
    {
        public readonly string ClusterName;
        /// <summary>
        /// (Required) Name of the collection the index is on.
        /// </summary>
        public readonly string CollectionName;
        /// <summary>
        /// (Required) Name of the database the collection is in.
        /// </summary>
        public readonly string Database;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly int? ItemsPerPage;
        public readonly int? PageNum;
        public readonly string ProjectId;
        /// <summary>
        /// A list where each represents a search index.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetSearchIndexesResultResult> Results;
        /// <summary>
        /// Represents the total of the search indexes
        /// </summary>
        public readonly int TotalCount;

        [OutputConstructor]
        private GetSearchIndexesResult(
            string clusterName,

            string collectionName,

            string database,

            string id,

            int? itemsPerPage,

            int? pageNum,

            string projectId,

            ImmutableArray<Outputs.GetSearchIndexesResultResult> results,

            int totalCount)
        {
            ClusterName = clusterName;
            CollectionName = collectionName;
            Database = database;
            Id = id;
            ItemsPerPage = itemsPerPage;
            PageNum = pageNum;
            ProjectId = projectId;
            Results = results;
            TotalCount = totalCount;
        }
    }
}
