// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package mongodbatlas

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// `.Cluster` describes a Cluster. The. The data source requires your Project ID.
//
// > **NOTE:** Groups and projects are synonymous terms. You may find groupId in the official documentation.
//
// > **IMPORTANT:**
// <br> &#8226; Changes to cluster configurations can affect costs. Before making changes, please see [Billing](https://docs.atlas.mongodb.com/billing/).
// <br> &#8226; If your Atlas project contains a custom role that uses actions introduced in a specific MongoDB version, you cannot create a cluster with a MongoDB version less than that version unless you delete the custom role.
func LookupCluster(ctx *pulumi.Context, args *LookupClusterArgs, opts ...pulumi.InvokeOption) (*LookupClusterResult, error) {
	var rv LookupClusterResult
	err := ctx.Invoke("mongodbatlas:index/getCluster:getCluster", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCluster.
type LookupClusterArgs struct {
	// Name of the cluster as it appears in Atlas. Once the cluster is created, its name cannot be changed.
	Name string `pulumi:"name"`
	// The unique ID for the project to create the database user.
	ProjectId string `pulumi:"projectId"`
}

// A collection of values returned by getCluster.
type LookupClusterResult struct {
	// Indicates whether disk auto-scaling is enabled.
	AutoScalingDiskGbEnabled bool `pulumi:"autoScalingDiskGbEnabled"`
	// Indicates Cloud service provider on which the server for a multi-tenant cluster is provisioned.
	BackingProviderName string `pulumi:"backingProviderName"`
	// Legacy Option, Indicates whether Atlas continuous backups are enabled for the cluster.
	BackupEnabled bool `pulumi:"backupEnabled"`
	// Indicates BI Connector for Atlas configuration on this cluster. BI Connector for Atlas is only available for M10+ clusters. See BI Connector below for more details.
	BiConnector GetClusterBiConnector `pulumi:"biConnector"`
	// Indicates the type of the cluster that you want to modify. You cannot convert a sharded cluster deployment to a replica set deployment.
	ClusterType string `pulumi:"clusterType"`
	// Set of connection strings that your applications use to connect to this cluster. More info in [Connection-strings](https://docs.mongodb.com/manual/reference/connection-string/). Use the parameters in this object to connect your applications to this cluster. To learn more about the formats of connection strings, see [Connection String Options](https://docs.atlas.mongodb.com/reference/faq/connection-changes/). NOTE: Atlas returns the contents of this object after the cluster is operational, not while it builds the cluster.
	// - `connection_strings.standard` -   Public mongodb:// connection string for this cluster.
	// - `connection_strings.standard_srv` - Public mongodb+srv:// connection string for this cluster. The mongodb+srv protocol tells the driver to look up the seed list of hosts in DNS. Atlas synchronizes this list with the nodes in a cluster. If the connection string uses this URI format, you don’t need to append the seed list or change the URI if the nodes change. Use this URI format if your driver supports it. If it doesn’t, use connectionStrings.standard.
	// - `connection_strings.aws_private_link` -  [Private-endpoint-aware](https://docs.atlas.mongodb.com/security-private-endpoint/#private-endpoint-connection-strings) mongodb://connection strings for each interface VPC endpoint you configured to connect to this cluster. Returned only if you created a AWS PrivateLink connection to this cluster.
	// - `connection_strings.aws_private_link_srv` - [Private-endpoint-aware](https://docs.atlas.mongodb.com/security-private-endpoint/#private-endpoint-connection-strings) mongodb+srv://connection strings for each interface VPC endpoint you configured to connect to this cluster. Returned only if you created a AWS PrivateLink connection to this cluster. Use this URI format if your driver supports it. If it doesn’t, use connectionStrings.awsPrivateLink.
	// - `connection_strings.private` -   [Network-peering-endpoint-aware](https://docs.atlas.mongodb.com/security-vpc-peering/#vpc-peering) mongodb://connection strings for each interface VPC endpoint you configured to connect to this cluster. Returned only if you created a network peering connection to this cluster.
	// - `connection_strings.private_srv` -  [Network-peering-endpoint-aware](https://docs.atlas.mongodb.com/security-vpc-peering/#vpc-peering) mongodb+srv://connection strings for each interface VPC endpoint you configured to connect to this cluster. Returned only if you created a network peering connection to this cluster.
	ConnectionStrings GetClusterConnectionStrings `pulumi:"connectionStrings"`
	// Indicates the size in gigabytes of the server’s root volume (AWS/GCP Only).
	DiskSizeGb float64 `pulumi:"diskSizeGb"`
	// Indicates whether Encryption at Rest is enabled or disabled.
	EncryptionAtRestProvider string `pulumi:"encryptionAtRestProvider"`
	// id is the provider-assigned unique ID for this managed resource.
	Id     string            `pulumi:"id"`
	Labels []GetClusterLabel `pulumi:"labels"`
	// Indicates the version of the cluster to deploy.
	MongoDbMajorVersion string `pulumi:"mongoDbMajorVersion"`
	// Version of MongoDB the cluster runs, in `major-version`.`minor-version` format.
	MongoDbVersion string `pulumi:"mongoDbVersion"`
	// Base connection string for the cluster. Atlas only displays this field after the cluster is operational, not while it builds the cluster.
	MongoUri string `pulumi:"mongoUri"`
	// Lists when the connection string was last updated. The connection string changes, for example, if you change a replica set to a sharded cluster.
	MongoUriUpdated string `pulumi:"mongoUriUpdated"`
	// Describes connection string for connecting to the Atlas cluster. Includes the replicaSet, ssl, and authSource query parameters in the connection string with values appropriate for the cluster.
	MongoUriWithOptions string `pulumi:"mongoUriWithOptions"`
	// The name of the current plugin
	Name string `pulumi:"name"`
	// Number of shards to deploy in the specified zone.
	NumShards int `pulumi:"numShards"`
	// Flag that indicates whether the cluster is paused or not.
	Paused bool `pulumi:"paused"`
	// Flag that indicates if the cluster uses Point-in-Time backups.
	PitEnabled bool   `pulumi:"pitEnabled"`
	ProjectId  string `pulumi:"projectId"`
	// Flag indicating if the cluster uses Cloud Provider Snapshots for backups.
	ProviderBackupEnabled bool `pulumi:"providerBackupEnabled"`
	// Indicates the maximum input/output operations per second (IOPS) the system can perform. The possible values depend on the selected providerSettings.instanceSizeName and diskSizeGB.
	ProviderDiskIops int `pulumi:"providerDiskIops"`
	// Describes Azure disk type of the server’s root volume (Azure Only).
	ProviderDiskTypeName string `pulumi:"providerDiskTypeName"`
	// Indicates whether the Amazon EBS encryption is enabled. This feature encrypts the server’s root volume for both data at rest within the volume and data moving between the volume and the instance.
	ProviderEncryptEbsVolume bool `pulumi:"providerEncryptEbsVolume"`
	// Atlas provides different instance sizes, each with a default storage capacity and RAM size.
	ProviderInstanceSizeName string `pulumi:"providerInstanceSizeName"`
	// Indicates the cloud service provider on which the servers are provisioned.
	ProviderName string `pulumi:"providerName"`
	// Indicates Physical location of your MongoDB cluster. The region you choose can affect network latency for clients accessing your databases.  Requires the Atlas Region name, see the reference list for [AWS](https://docs.atlas.mongodb.com/reference/amazon-aws/), [GCP](https://docs.atlas.mongodb.com/reference/google-gcp/), [Azure](https://docs.atlas.mongodb.com/reference/microsoft-azure/).
	ProviderRegionName string `pulumi:"providerRegionName"`
	// Indicates the type of the volume. The possible values are: `STANDARD` and `PROVISIONED`.
	ProviderVolumeType string `pulumi:"providerVolumeType"`
	// Number of replica set members. Each member keeps a copy of your databases, providing high availability and data redundancy. The possible values are 3, 5, or 7. The default value is 3.
	ReplicationFactor int `pulumi:"replicationFactor"`
	// Configuration for cluster regions.  See Replication Spec below for more details.
	ReplicationSpecs []GetClusterReplicationSpec `pulumi:"replicationSpecs"`
	// current snapshot schedule and retention settings for the cluster.
	SnapshotBackupPolicies []GetClusterSnapshotBackupPolicy `pulumi:"snapshotBackupPolicies"`
	// Connection string for connecting to the Atlas cluster. The +srv modifier forces the connection to use TLS/SSL. See the mongoURI for additional options.
	SrvAddress string `pulumi:"srvAddress"`
	// Indicates the current state of the cluster. The possible states are:
	// - IDLE
	// - CREATING
	// - UPDATING
	// - DELETING
	// - DELETED
	// - REPAIRING
	StateName string `pulumi:"stateName"`
}