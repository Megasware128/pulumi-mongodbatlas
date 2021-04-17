// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package mongodbatlas

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// `Cluster` describes all Clusters by the provided project_id. The data source requires your Project ID.
//
// > **NOTE:** Groups and projects are synonymous terms. You may find groupId in the official documentation.
//
// > **IMPORTANT:**
// <br> &#8226; Changes to cluster configurations can affect costs. Before making changes, please see [Billing](https://docs.atlas.mongodb.com/billing/).
// <br> &#8226; If your Atlas project contains a custom role that uses actions introduced in a specific MongoDB version, you cannot create a cluster with a MongoDB version less than that version unless you delete the custom role.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-mongodbatlas/sdk/v2/go/mongodbatlas"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		testCluster, err := mongodbatlas.NewCluster(ctx, "testCluster", &mongodbatlas.ClusterArgs{
// 			ProjectId:   pulumi.String("<YOUR-PROJECT-ID>"),
// 			DiskSizeGb:  pulumi.Float64(100),
// 			ClusterType: pulumi.String("REPLICASET"),
// 			ReplicationSpecs: mongodbatlas.ClusterReplicationSpecArray{
// 				&mongodbatlas.ClusterReplicationSpecArgs{
// 					NumShards: pulumi.Int(1),
// 					RegionsConfigs: mongodbatlas.ClusterReplicationSpecRegionsConfigArray{
// 						&mongodbatlas.ClusterReplicationSpecRegionsConfigArgs{
// 							RegionName:     pulumi.String("US_EAST_1"),
// 							ElectableNodes: pulumi.Int(3),
// 							Priority:       pulumi.Int(7),
// 							ReadOnlyNodes:  pulumi.Int(0),
// 						},
// 					},
// 				},
// 			},
// 			ProviderBackupEnabled:    pulumi.Bool(true),
// 			AutoScalingDiskGbEnabled: pulumi.Bool(true),
// 			ProviderName:             pulumi.String("AWS"),
// 			ProviderDiskIops:         pulumi.Int(300),
// 			ProviderVolumeType:       pulumi.String("STANDARD"),
// 			ProviderEncryptEbsVolume: pulumi.Bool(true),
// 			ProviderInstanceSizeName: pulumi.String("M40"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetClusters(ctx *pulumi.Context, args *GetClustersArgs, opts ...pulumi.InvokeOption) (*GetClustersResult, error) {
	var rv GetClustersResult
	err := ctx.Invoke("mongodbatlas:index/getClusters:getClusters", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getClusters.
type GetClustersArgs struct {
	// The unique ID for the project to get the clusters.
	ProjectId string `pulumi:"projectId"`
}

// A collection of values returned by getClusters.
type GetClustersResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id        string `pulumi:"id"`
	ProjectId string `pulumi:"projectId"`
	// A list where each represents a Cluster. See Cluster below for more details.
	Results []GetClustersResultType `pulumi:"results"`
}
