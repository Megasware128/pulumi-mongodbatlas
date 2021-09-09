// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package mongodbatlas

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// `SearchIndex` describe a single search indexes. This represents a single search index that have been created.
//
// > **NOTE:** Groups and projects are synonymous terms. You may find `groupId` in the official documentation.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-mongodbatlas/sdk/v3/go/mongodbatlas"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := mongodbatlas.LookupSearchIndex(ctx, &mongodbatlas.LookupSearchIndexArgs{
// 			ClusterName: "<CLUSTER_NAME>",
// 			IndexId:     "<INDEX_ID",
// 			ProjectId:   "<PROJECT_ID>",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupSearchIndex(ctx *pulumi.Context, args *LookupSearchIndexArgs, opts ...pulumi.InvokeOption) (*LookupSearchIndexResult, error) {
	var rv LookupSearchIndexResult
	err := ctx.Invoke("mongodbatlas:index/getSearchIndex:getSearchIndex", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSearchIndex.
type LookupSearchIndexArgs struct {
	// [Analyzer](https://docs.atlas.mongodb.com/reference/atlas-search/analyzers/#std-label-analyzers-ref) to use when creating the index.
	Analyzer *string `pulumi:"analyzer"`
	// [Custom analyzers](https://docs.atlas.mongodb.com/reference/atlas-search/analyzers/custom/#std-label-custom-analyzers) to use in this index (this is an array of objects).
	Analyzers *string `pulumi:"analyzers"`
	// The name of the cluster containing the collection with one or more Atlas Search indexes.
	ClusterName string `pulumi:"clusterName"`
	// (Required) Name of the collection the index is on.
	CollectionName *string `pulumi:"collectionName"`
	// (Required) Name of the database the collection is in.
	Database *string `pulumi:"database"`
	// The unique identifier of the Atlas Search index. Use the `getSearchIndexes`datasource to find the IDs of all Atlas Search indexes.
	IndexId string `pulumi:"indexId"`
	// Flag indicating whether the index uses dynamic or static mappings.
	MappingsDynamic *bool `pulumi:"mappingsDynamic"`
	// Object containing one or more field specifications.
	MappingsFields *string `pulumi:"mappingsFields"`
	// Name of the index.
	Name *string `pulumi:"name"`
	// The unique identifier for the [project](https://docs.atlas.mongodb.com/organizations-projects/#std-label-projects) that contains the specified cluster.
	ProjectId string `pulumi:"projectId"`
	// [Analyzer](https://docs.atlas.mongodb.com/reference/atlas-search/analyzers/#std-label-analyzers-ref) to use when searching the index.
	SearchAnalyzer *string `pulumi:"searchAnalyzer"`
	Status         *string `pulumi:"status"`
}

// A collection of values returned by getSearchIndex.
type LookupSearchIndexResult struct {
	// [Analyzer](https://docs.atlas.mongodb.com/reference/atlas-search/analyzers/#std-label-analyzers-ref) to use when creating the index.
	Analyzer *string `pulumi:"analyzer"`
	// [Custom analyzers](https://docs.atlas.mongodb.com/reference/atlas-search/analyzers/custom/#std-label-custom-analyzers) to use in this index (this is an array of objects).
	Analyzers   *string `pulumi:"analyzers"`
	ClusterName string  `pulumi:"clusterName"`
	// (Required) Name of the collection the index is on.
	CollectionName *string `pulumi:"collectionName"`
	// (Required) Name of the database the collection is in.
	Database *string `pulumi:"database"`
	// The provider-assigned unique ID for this managed resource.
	Id      string `pulumi:"id"`
	IndexId string `pulumi:"indexId"`
	// Flag indicating whether the index uses dynamic or static mappings.
	MappingsDynamic *bool `pulumi:"mappingsDynamic"`
	// Object containing one or more field specifications.
	MappingsFields *string `pulumi:"mappingsFields"`
	// Name of the index.
	Name      *string `pulumi:"name"`
	ProjectId string  `pulumi:"projectId"`
	// [Analyzer](https://docs.atlas.mongodb.com/reference/atlas-search/analyzers/#std-label-analyzers-ref) to use when searching the index.
	SearchAnalyzer *string `pulumi:"searchAnalyzer"`
	Status         string  `pulumi:"status"`
}
