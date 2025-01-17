// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Mongodbatlas.Outputs
{

    [OutputType]
    public sealed class GetAdvancedClustersResultResult
    {
        /// <summary>
        /// Get the advanced configuration options. See Advanced Configuration below for more details.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetAdvancedClustersResultAdvancedConfigurationResult> AdvancedConfigurations;
        public readonly bool BackupEnabled;
        /// <summary>
        /// Configuration settings applied to BI Connector for Atlas on this cluster. See below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetAdvancedClustersResultBiConnectorResult> BiConnectors;
        /// <summary>
        /// Type of the cluster that you want to create.
        /// </summary>
        public readonly string ClusterType;
        /// <summary>
        /// Set of connection strings that your applications use to connect to this cluster. More info in [Connection-strings](https://docs.mongodb.com/manual/reference/connection-string/). Use the parameters in this object to connect your applications to this cluster. To learn more about the formats of connection strings, see [Connection String Options](https://docs.atlas.mongodb.com/reference/faq/connection-changes/). NOTE: Atlas returns the contents of this object after the cluster is operational, not while it builds the cluster.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetAdvancedClustersResultConnectionStringResult> ConnectionStrings;
        public readonly string CreateDate;
        /// <summary>
        /// Capacity, in gigabytes, of the host's root volume.
        /// </summary>
        public readonly double DiskSizeGb;
        /// <summary>
        /// Possible values are AWS, GCP, AZURE or NONE.
        /// </summary>
        public readonly string EncryptionAtRestProvider;
        /// <summary>
        /// Configuration for the collection of key-value pairs that tag and categorize the cluster. See below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetAdvancedClustersResultLabelResult> Labels;
        /// <summary>
        /// Version of the cluster to deploy.
        /// </summary>
        public readonly string MongoDbMajorVersion;
        /// <summary>
        /// Version of MongoDB the cluster runs, in `major-version`.`minor-version` format.
        /// </summary>
        public readonly string MongoDbVersion;
        public readonly string Name;
        /// <summary>
        /// Flag that indicates whether the cluster is paused or not.
        /// </summary>
        public readonly bool Paused;
        /// <summary>
        /// Flag that indicates if the cluster uses Continuous Cloud Backup.
        /// </summary>
        public readonly bool PitEnabled;
        /// <summary>
        /// Configuration for cluster regions and the hardware provisioned in them. See below
        /// </summary>
        public readonly ImmutableArray<Outputs.GetAdvancedClustersResultReplicationSpecResult> ReplicationSpecs;
        /// <summary>
        /// Certificate Authority that MongoDB Atlas clusters use.
        /// </summary>
        public readonly string RootCertType;
        /// <summary>
        /// Current state of the cluster. The possible states are:
        /// </summary>
        public readonly string StateName;
        /// <summary>
        /// Release cadence that Atlas uses for this cluster.
        /// </summary>
        public readonly string VersionReleaseSystem;

        [OutputConstructor]
        private GetAdvancedClustersResultResult(
            ImmutableArray<Outputs.GetAdvancedClustersResultAdvancedConfigurationResult> advancedConfigurations,

            bool backupEnabled,

            ImmutableArray<Outputs.GetAdvancedClustersResultBiConnectorResult> biConnectors,

            string clusterType,

            ImmutableArray<Outputs.GetAdvancedClustersResultConnectionStringResult> connectionStrings,

            string createDate,

            double diskSizeGb,

            string encryptionAtRestProvider,

            ImmutableArray<Outputs.GetAdvancedClustersResultLabelResult> labels,

            string mongoDbMajorVersion,

            string mongoDbVersion,

            string name,

            bool paused,

            bool pitEnabled,

            ImmutableArray<Outputs.GetAdvancedClustersResultReplicationSpecResult> replicationSpecs,

            string rootCertType,

            string stateName,

            string versionReleaseSystem)
        {
            AdvancedConfigurations = advancedConfigurations;
            BackupEnabled = backupEnabled;
            BiConnectors = biConnectors;
            ClusterType = clusterType;
            ConnectionStrings = connectionStrings;
            CreateDate = createDate;
            DiskSizeGb = diskSizeGb;
            EncryptionAtRestProvider = encryptionAtRestProvider;
            Labels = labels;
            MongoDbMajorVersion = mongoDbMajorVersion;
            MongoDbVersion = mongoDbVersion;
            Name = name;
            Paused = paused;
            PitEnabled = pitEnabled;
            ReplicationSpecs = replicationSpecs;
            RootCertType = rootCertType;
            StateName = stateName;
            VersionReleaseSystem = versionReleaseSystem;
        }
    }
}
