# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from . import _utilities, _tables
from . import outputs

__all__ = [
    'GetClusterResult',
    'AwaitableGetClusterResult',
    'get_cluster',
]

@pulumi.output_type
class GetClusterResult:
    """
    A collection of values returned by getCluster.
    """
    def __init__(__self__, auto_scaling_compute_enabled=None, auto_scaling_compute_scale_down_enabled=None, auto_scaling_disk_gb_enabled=None, backing_provider_name=None, backup_enabled=None, bi_connector=None, cluster_type=None, connection_strings=None, container_id=None, disk_size_gb=None, encryption_at_rest_provider=None, id=None, labels=None, mongo_db_major_version=None, mongo_db_version=None, mongo_uri=None, mongo_uri_updated=None, mongo_uri_with_options=None, name=None, num_shards=None, paused=None, pit_enabled=None, project_id=None, provider_auto_scaling_compute_max_instance_size=None, provider_auto_scaling_compute_min_instance_size=None, provider_backup_enabled=None, provider_disk_iops=None, provider_disk_type_name=None, provider_encrypt_ebs_volume=None, provider_instance_size_name=None, provider_name=None, provider_region_name=None, provider_volume_type=None, replication_factor=None, replication_specs=None, snapshot_backup_policies=None, srv_address=None, state_name=None):
        if auto_scaling_compute_enabled and not isinstance(auto_scaling_compute_enabled, bool):
            raise TypeError("Expected argument 'auto_scaling_compute_enabled' to be a bool")
        pulumi.set(__self__, "auto_scaling_compute_enabled", auto_scaling_compute_enabled)
        if auto_scaling_compute_scale_down_enabled and not isinstance(auto_scaling_compute_scale_down_enabled, bool):
            raise TypeError("Expected argument 'auto_scaling_compute_scale_down_enabled' to be a bool")
        pulumi.set(__self__, "auto_scaling_compute_scale_down_enabled", auto_scaling_compute_scale_down_enabled)
        if auto_scaling_disk_gb_enabled and not isinstance(auto_scaling_disk_gb_enabled, bool):
            raise TypeError("Expected argument 'auto_scaling_disk_gb_enabled' to be a bool")
        pulumi.set(__self__, "auto_scaling_disk_gb_enabled", auto_scaling_disk_gb_enabled)
        if backing_provider_name and not isinstance(backing_provider_name, str):
            raise TypeError("Expected argument 'backing_provider_name' to be a str")
        pulumi.set(__self__, "backing_provider_name", backing_provider_name)
        if backup_enabled and not isinstance(backup_enabled, bool):
            raise TypeError("Expected argument 'backup_enabled' to be a bool")
        pulumi.set(__self__, "backup_enabled", backup_enabled)
        if bi_connector and not isinstance(bi_connector, dict):
            raise TypeError("Expected argument 'bi_connector' to be a dict")
        pulumi.set(__self__, "bi_connector", bi_connector)
        if cluster_type and not isinstance(cluster_type, str):
            raise TypeError("Expected argument 'cluster_type' to be a str")
        pulumi.set(__self__, "cluster_type", cluster_type)
        if connection_strings and not isinstance(connection_strings, dict):
            raise TypeError("Expected argument 'connection_strings' to be a dict")
        pulumi.set(__self__, "connection_strings", connection_strings)
        if container_id and not isinstance(container_id, str):
            raise TypeError("Expected argument 'container_id' to be a str")
        pulumi.set(__self__, "container_id", container_id)
        if disk_size_gb and not isinstance(disk_size_gb, float):
            raise TypeError("Expected argument 'disk_size_gb' to be a float")
        pulumi.set(__self__, "disk_size_gb", disk_size_gb)
        if encryption_at_rest_provider and not isinstance(encryption_at_rest_provider, str):
            raise TypeError("Expected argument 'encryption_at_rest_provider' to be a str")
        pulumi.set(__self__, "encryption_at_rest_provider", encryption_at_rest_provider)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if labels and not isinstance(labels, list):
            raise TypeError("Expected argument 'labels' to be a list")
        pulumi.set(__self__, "labels", labels)
        if mongo_db_major_version and not isinstance(mongo_db_major_version, str):
            raise TypeError("Expected argument 'mongo_db_major_version' to be a str")
        pulumi.set(__self__, "mongo_db_major_version", mongo_db_major_version)
        if mongo_db_version and not isinstance(mongo_db_version, str):
            raise TypeError("Expected argument 'mongo_db_version' to be a str")
        pulumi.set(__self__, "mongo_db_version", mongo_db_version)
        if mongo_uri and not isinstance(mongo_uri, str):
            raise TypeError("Expected argument 'mongo_uri' to be a str")
        pulumi.set(__self__, "mongo_uri", mongo_uri)
        if mongo_uri_updated and not isinstance(mongo_uri_updated, str):
            raise TypeError("Expected argument 'mongo_uri_updated' to be a str")
        pulumi.set(__self__, "mongo_uri_updated", mongo_uri_updated)
        if mongo_uri_with_options and not isinstance(mongo_uri_with_options, str):
            raise TypeError("Expected argument 'mongo_uri_with_options' to be a str")
        pulumi.set(__self__, "mongo_uri_with_options", mongo_uri_with_options)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if num_shards and not isinstance(num_shards, float):
            raise TypeError("Expected argument 'num_shards' to be a float")
        pulumi.set(__self__, "num_shards", num_shards)
        if paused and not isinstance(paused, bool):
            raise TypeError("Expected argument 'paused' to be a bool")
        pulumi.set(__self__, "paused", paused)
        if pit_enabled and not isinstance(pit_enabled, bool):
            raise TypeError("Expected argument 'pit_enabled' to be a bool")
        pulumi.set(__self__, "pit_enabled", pit_enabled)
        if project_id and not isinstance(project_id, str):
            raise TypeError("Expected argument 'project_id' to be a str")
        pulumi.set(__self__, "project_id", project_id)
        if provider_auto_scaling_compute_max_instance_size and not isinstance(provider_auto_scaling_compute_max_instance_size, str):
            raise TypeError("Expected argument 'provider_auto_scaling_compute_max_instance_size' to be a str")
        pulumi.set(__self__, "provider_auto_scaling_compute_max_instance_size", provider_auto_scaling_compute_max_instance_size)
        if provider_auto_scaling_compute_min_instance_size and not isinstance(provider_auto_scaling_compute_min_instance_size, str):
            raise TypeError("Expected argument 'provider_auto_scaling_compute_min_instance_size' to be a str")
        pulumi.set(__self__, "provider_auto_scaling_compute_min_instance_size", provider_auto_scaling_compute_min_instance_size)
        if provider_backup_enabled and not isinstance(provider_backup_enabled, bool):
            raise TypeError("Expected argument 'provider_backup_enabled' to be a bool")
        pulumi.set(__self__, "provider_backup_enabled", provider_backup_enabled)
        if provider_disk_iops and not isinstance(provider_disk_iops, float):
            raise TypeError("Expected argument 'provider_disk_iops' to be a float")
        pulumi.set(__self__, "provider_disk_iops", provider_disk_iops)
        if provider_disk_type_name and not isinstance(provider_disk_type_name, str):
            raise TypeError("Expected argument 'provider_disk_type_name' to be a str")
        pulumi.set(__self__, "provider_disk_type_name", provider_disk_type_name)
        if provider_encrypt_ebs_volume and not isinstance(provider_encrypt_ebs_volume, bool):
            raise TypeError("Expected argument 'provider_encrypt_ebs_volume' to be a bool")
        pulumi.set(__self__, "provider_encrypt_ebs_volume", provider_encrypt_ebs_volume)
        if provider_instance_size_name and not isinstance(provider_instance_size_name, str):
            raise TypeError("Expected argument 'provider_instance_size_name' to be a str")
        pulumi.set(__self__, "provider_instance_size_name", provider_instance_size_name)
        if provider_name and not isinstance(provider_name, str):
            raise TypeError("Expected argument 'provider_name' to be a str")
        pulumi.set(__self__, "provider_name", provider_name)
        if provider_region_name and not isinstance(provider_region_name, str):
            raise TypeError("Expected argument 'provider_region_name' to be a str")
        pulumi.set(__self__, "provider_region_name", provider_region_name)
        if provider_volume_type and not isinstance(provider_volume_type, str):
            raise TypeError("Expected argument 'provider_volume_type' to be a str")
        pulumi.set(__self__, "provider_volume_type", provider_volume_type)
        if replication_factor and not isinstance(replication_factor, float):
            raise TypeError("Expected argument 'replication_factor' to be a float")
        pulumi.set(__self__, "replication_factor", replication_factor)
        if replication_specs and not isinstance(replication_specs, list):
            raise TypeError("Expected argument 'replication_specs' to be a list")
        pulumi.set(__self__, "replication_specs", replication_specs)
        if snapshot_backup_policies and not isinstance(snapshot_backup_policies, list):
            raise TypeError("Expected argument 'snapshot_backup_policies' to be a list")
        pulumi.set(__self__, "snapshot_backup_policies", snapshot_backup_policies)
        if srv_address and not isinstance(srv_address, str):
            raise TypeError("Expected argument 'srv_address' to be a str")
        pulumi.set(__self__, "srv_address", srv_address)
        if state_name and not isinstance(state_name, str):
            raise TypeError("Expected argument 'state_name' to be a str")
        pulumi.set(__self__, "state_name", state_name)

    @property
    @pulumi.getter(name="autoScalingComputeEnabled")
    def auto_scaling_compute_enabled(self) -> bool:
        """
        (Optional) Specifies whether cluster tier auto-scaling is enabled. The default is false.
        """
        return pulumi.get(self, "auto_scaling_compute_enabled")

    @property
    @pulumi.getter(name="autoScalingComputeScaleDownEnabled")
    def auto_scaling_compute_scale_down_enabled(self) -> bool:
        """
        (Optional) Set to `true` to enable the cluster tier to scale down.
        """
        return pulumi.get(self, "auto_scaling_compute_scale_down_enabled")

    @property
    @pulumi.getter(name="autoScalingDiskGbEnabled")
    def auto_scaling_disk_gb_enabled(self) -> bool:
        """
        Indicates whether disk auto-scaling is enabled.
        """
        return pulumi.get(self, "auto_scaling_disk_gb_enabled")

    @property
    @pulumi.getter(name="backingProviderName")
    def backing_provider_name(self) -> str:
        """
        Indicates Cloud service provider on which the server for a multi-tenant cluster is provisioned.
        """
        return pulumi.get(self, "backing_provider_name")

    @property
    @pulumi.getter(name="backupEnabled")
    def backup_enabled(self) -> bool:
        """
        Legacy Option, Indicates whether Atlas continuous backups are enabled for the cluster.
        """
        return pulumi.get(self, "backup_enabled")

    @property
    @pulumi.getter(name="biConnector")
    def bi_connector(self) -> 'outputs.GetClusterBiConnectorResult':
        """
        Indicates BI Connector for Atlas configuration on this cluster. BI Connector for Atlas is only available for M10+ clusters. See BI Connector below for more details.
        """
        return pulumi.get(self, "bi_connector")

    @property
    @pulumi.getter(name="clusterType")
    def cluster_type(self) -> str:
        """
        Indicates the type of the cluster that you want to modify. You cannot convert a sharded cluster deployment to a replica set deployment.
        """
        return pulumi.get(self, "cluster_type")

    @property
    @pulumi.getter(name="connectionStrings")
    def connection_strings(self) -> 'outputs.GetClusterConnectionStringsResult':
        """
        Set of connection strings that your applications use to connect to this cluster. More info in [Connection-strings](https://docs.mongodb.com/manual/reference/connection-string/). Use the parameters in this object to connect your applications to this cluster. To learn more about the formats of connection strings, see [Connection String Options](https://docs.atlas.mongodb.com/reference/faq/connection-changes/). NOTE: Atlas returns the contents of this object after the cluster is operational, not while it builds the cluster.
        - `connection_strings.standard` -   Public mongodb:// connection string for this cluster.
        - `connection_strings.standard_srv` - Public mongodb+srv:// connection string for this cluster. The mongodb+srv protocol tells the driver to look up the seed list of hosts in DNS. Atlas synchronizes this list with the nodes in a cluster. If the connection string uses this URI format, you don’t need to append the seed list or change the URI if the nodes change. Use this URI format if your driver supports it. If it doesn’t, use connectionStrings.standard.
        - `connection_strings.aws_private_link` -  [Private-endpoint-aware](https://docs.atlas.mongodb.com/security-private-endpoint/#private-endpoint-connection-strings) mongodb://connection strings for each interface VPC endpoint you configured to connect to this cluster. Returned only if you created a AWS PrivateLink connection to this cluster.
        - `connection_strings.aws_private_link_srv` - [Private-endpoint-aware](https://docs.atlas.mongodb.com/security-private-endpoint/#private-endpoint-connection-strings) mongodb+srv://connection strings for each interface VPC endpoint you configured to connect to this cluster. Returned only if you created a AWS PrivateLink connection to this cluster. Use this URI format if your driver supports it. If it doesn’t, use connectionStrings.awsPrivateLink.
        - `connection_strings.private` -   [Network-peering-endpoint-aware](https://docs.atlas.mongodb.com/security-vpc-peering/#vpc-peering) mongodb://connection strings for each interface VPC endpoint you configured to connect to this cluster. Returned only if you created a network peering connection to this cluster.
        - `connection_strings.private_srv` -  [Network-peering-endpoint-aware](https://docs.atlas.mongodb.com/security-vpc-peering/#vpc-peering) mongodb+srv://connection strings for each interface VPC endpoint you configured to connect to this cluster. Returned only if you created a network peering connection to this cluster.
        """
        return pulumi.get(self, "connection_strings")

    @property
    @pulumi.getter(name="containerId")
    def container_id(self) -> str:
        """
        The Network Peering Container ID.
        """
        return pulumi.get(self, "container_id")

    @property
    @pulumi.getter(name="diskSizeGb")
    def disk_size_gb(self) -> float:
        """
        Indicates the size in gigabytes of the server’s root volume (AWS/GCP Only).
        """
        return pulumi.get(self, "disk_size_gb")

    @property
    @pulumi.getter(name="encryptionAtRestProvider")
    def encryption_at_rest_provider(self) -> str:
        """
        Indicates whether Encryption at Rest is enabled or disabled.
        """
        return pulumi.get(self, "encryption_at_rest_provider")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def labels(self) -> List['outputs.GetClusterLabelResult']:
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter(name="mongoDbMajorVersion")
    def mongo_db_major_version(self) -> str:
        """
        Indicates the version of the cluster to deploy.
        """
        return pulumi.get(self, "mongo_db_major_version")

    @property
    @pulumi.getter(name="mongoDbVersion")
    def mongo_db_version(self) -> str:
        """
        Version of MongoDB the cluster runs, in `major-version`.`minor-version` format.
        """
        return pulumi.get(self, "mongo_db_version")

    @property
    @pulumi.getter(name="mongoUri")
    def mongo_uri(self) -> str:
        """
        Base connection string for the cluster. Atlas only displays this field after the cluster is operational, not while it builds the cluster.
        """
        return pulumi.get(self, "mongo_uri")

    @property
    @pulumi.getter(name="mongoUriUpdated")
    def mongo_uri_updated(self) -> str:
        """
        Lists when the connection string was last updated. The connection string changes, for example, if you change a replica set to a sharded cluster.
        """
        return pulumi.get(self, "mongo_uri_updated")

    @property
    @pulumi.getter(name="mongoUriWithOptions")
    def mongo_uri_with_options(self) -> str:
        """
        Describes connection string for connecting to the Atlas cluster. Includes the replicaSet, ssl, and authSource query parameters in the connection string with values appropriate for the cluster.
        """
        return pulumi.get(self, "mongo_uri_with_options")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the current plugin
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="numShards")
    def num_shards(self) -> float:
        """
        Number of shards to deploy in the specified zone.
        """
        return pulumi.get(self, "num_shards")

    @property
    @pulumi.getter
    def paused(self) -> bool:
        """
        Flag that indicates whether the cluster is paused or not.
        """
        return pulumi.get(self, "paused")

    @property
    @pulumi.getter(name="pitEnabled")
    def pit_enabled(self) -> bool:
        """
        Flag that indicates if the cluster uses Continuous Cloud Backup.
        """
        return pulumi.get(self, "pit_enabled")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> str:
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter(name="providerAutoScalingComputeMaxInstanceSize")
    def provider_auto_scaling_compute_max_instance_size(self) -> str:
        """
        (Optional) Maximum instance size to which your cluster can automatically scale.
        """
        return pulumi.get(self, "provider_auto_scaling_compute_max_instance_size")

    @property
    @pulumi.getter(name="providerAutoScalingComputeMinInstanceSize")
    def provider_auto_scaling_compute_min_instance_size(self) -> str:
        """
        (Optional) Minimum instance size to which your cluster can automatically scale.
        """
        return pulumi.get(self, "provider_auto_scaling_compute_min_instance_size")

    @property
    @pulumi.getter(name="providerBackupEnabled")
    def provider_backup_enabled(self) -> bool:
        """
        Flag indicating if the cluster uses Cloud Backup Snapshots for backups.
        """
        return pulumi.get(self, "provider_backup_enabled")

    @property
    @pulumi.getter(name="providerDiskIops")
    def provider_disk_iops(self) -> float:
        """
        Indicates the maximum input/output operations per second (IOPS) the system can perform. The possible values depend on the selected providerSettings.instanceSizeName and diskSizeGB.
        """
        return pulumi.get(self, "provider_disk_iops")

    @property
    @pulumi.getter(name="providerDiskTypeName")
    def provider_disk_type_name(self) -> str:
        """
        Describes Azure disk type of the server’s root volume (Azure Only).
        """
        return pulumi.get(self, "provider_disk_type_name")

    @property
    @pulumi.getter(name="providerEncryptEbsVolume")
    def provider_encrypt_ebs_volume(self) -> bool:
        """
        Indicates whether the Amazon EBS encryption is enabled. This feature encrypts the server’s root volume for both data at rest within the volume and data moving between the volume and the instance.
        """
        return pulumi.get(self, "provider_encrypt_ebs_volume")

    @property
    @pulumi.getter(name="providerInstanceSizeName")
    def provider_instance_size_name(self) -> str:
        """
        Atlas provides different instance sizes, each with a default storage capacity and RAM size.
        """
        return pulumi.get(self, "provider_instance_size_name")

    @property
    @pulumi.getter(name="providerName")
    def provider_name(self) -> str:
        """
        Indicates the cloud service provider on which the servers are provisioned.
        """
        return pulumi.get(self, "provider_name")

    @property
    @pulumi.getter(name="providerRegionName")
    def provider_region_name(self) -> str:
        """
        Indicates Physical location of your MongoDB cluster. The region you choose can affect network latency for clients accessing your databases.  Requires the Atlas Region name, see the reference list for [AWS](https://docs.atlas.mongodb.com/reference/amazon-aws/), [GCP](https://docs.atlas.mongodb.com/reference/google-gcp/), [Azure](https://docs.atlas.mongodb.com/reference/microsoft-azure/).
        """
        return pulumi.get(self, "provider_region_name")

    @property
    @pulumi.getter(name="providerVolumeType")
    def provider_volume_type(self) -> str:
        """
        Indicates the type of the volume. The possible values are: `STANDARD` and `PROVISIONED`.
        """
        return pulumi.get(self, "provider_volume_type")

    @property
    @pulumi.getter(name="replicationFactor")
    def replication_factor(self) -> float:
        """
        Number of replica set members. Each member keeps a copy of your databases, providing high availability and data redundancy. The possible values are 3, 5, or 7. The default value is 3.
        """
        return pulumi.get(self, "replication_factor")

    @property
    @pulumi.getter(name="replicationSpecs")
    def replication_specs(self) -> List['outputs.GetClusterReplicationSpecResult']:
        """
        Configuration for cluster regions.  See Replication Spec below for more details.
        """
        return pulumi.get(self, "replication_specs")

    @property
    @pulumi.getter(name="snapshotBackupPolicies")
    def snapshot_backup_policies(self) -> List['outputs.GetClusterSnapshotBackupPolicyResult']:
        """
        current snapshot schedule and retention settings for the cluster.
        """
        return pulumi.get(self, "snapshot_backup_policies")

    @property
    @pulumi.getter(name="srvAddress")
    def srv_address(self) -> str:
        """
        Connection string for connecting to the Atlas cluster. The +srv modifier forces the connection to use TLS/SSL. See the mongoURI for additional options.
        """
        return pulumi.get(self, "srv_address")

    @property
    @pulumi.getter(name="stateName")
    def state_name(self) -> str:
        """
        Indicates the current state of the cluster. The possible states are:
        - IDLE
        - CREATING
        - UPDATING
        - DELETING
        - DELETED
        - REPAIRING
        """
        return pulumi.get(self, "state_name")


class AwaitableGetClusterResult(GetClusterResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetClusterResult(
            auto_scaling_compute_enabled=self.auto_scaling_compute_enabled,
            auto_scaling_compute_scale_down_enabled=self.auto_scaling_compute_scale_down_enabled,
            auto_scaling_disk_gb_enabled=self.auto_scaling_disk_gb_enabled,
            backing_provider_name=self.backing_provider_name,
            backup_enabled=self.backup_enabled,
            bi_connector=self.bi_connector,
            cluster_type=self.cluster_type,
            connection_strings=self.connection_strings,
            container_id=self.container_id,
            disk_size_gb=self.disk_size_gb,
            encryption_at_rest_provider=self.encryption_at_rest_provider,
            id=self.id,
            labels=self.labels,
            mongo_db_major_version=self.mongo_db_major_version,
            mongo_db_version=self.mongo_db_version,
            mongo_uri=self.mongo_uri,
            mongo_uri_updated=self.mongo_uri_updated,
            mongo_uri_with_options=self.mongo_uri_with_options,
            name=self.name,
            num_shards=self.num_shards,
            paused=self.paused,
            pit_enabled=self.pit_enabled,
            project_id=self.project_id,
            provider_auto_scaling_compute_max_instance_size=self.provider_auto_scaling_compute_max_instance_size,
            provider_auto_scaling_compute_min_instance_size=self.provider_auto_scaling_compute_min_instance_size,
            provider_backup_enabled=self.provider_backup_enabled,
            provider_disk_iops=self.provider_disk_iops,
            provider_disk_type_name=self.provider_disk_type_name,
            provider_encrypt_ebs_volume=self.provider_encrypt_ebs_volume,
            provider_instance_size_name=self.provider_instance_size_name,
            provider_name=self.provider_name,
            provider_region_name=self.provider_region_name,
            provider_volume_type=self.provider_volume_type,
            replication_factor=self.replication_factor,
            replication_specs=self.replication_specs,
            snapshot_backup_policies=self.snapshot_backup_policies,
            srv_address=self.srv_address,
            state_name=self.state_name)


def get_cluster(name: Optional[str] = None,
                project_id: Optional[str] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetClusterResult:
    """
    `Cluster` describes a Cluster. The. The data source requires your Project ID.

    > **NOTE:** Groups and projects are synonymous terms. You may find group_id in the official documentation.

    > **IMPORTANT:**
    <br> &#8226; Changes to cluster configurations can affect costs. Before making changes, please see [Billing](https://docs.atlas.mongodb.com/billing/).
    <br> &#8226; If your Atlas project contains a custom role that uses actions introduced in a specific MongoDB version, you cannot create a cluster with a MongoDB version less than that version unless you delete the custom role.


    :param str name: Name of the cluster as it appears in Atlas. Once the cluster is created, its name cannot be changed.
    :param str project_id: The unique ID for the project to create the database user.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['projectId'] = project_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('mongodbatlas:index/getCluster:getCluster', __args__, opts=opts, typ=GetClusterResult).value

    return AwaitableGetClusterResult(
        auto_scaling_compute_enabled=__ret__.auto_scaling_compute_enabled,
        auto_scaling_compute_scale_down_enabled=__ret__.auto_scaling_compute_scale_down_enabled,
        auto_scaling_disk_gb_enabled=__ret__.auto_scaling_disk_gb_enabled,
        backing_provider_name=__ret__.backing_provider_name,
        backup_enabled=__ret__.backup_enabled,
        bi_connector=__ret__.bi_connector,
        cluster_type=__ret__.cluster_type,
        connection_strings=__ret__.connection_strings,
        container_id=__ret__.container_id,
        disk_size_gb=__ret__.disk_size_gb,
        encryption_at_rest_provider=__ret__.encryption_at_rest_provider,
        id=__ret__.id,
        labels=__ret__.labels,
        mongo_db_major_version=__ret__.mongo_db_major_version,
        mongo_db_version=__ret__.mongo_db_version,
        mongo_uri=__ret__.mongo_uri,
        mongo_uri_updated=__ret__.mongo_uri_updated,
        mongo_uri_with_options=__ret__.mongo_uri_with_options,
        name=__ret__.name,
        num_shards=__ret__.num_shards,
        paused=__ret__.paused,
        pit_enabled=__ret__.pit_enabled,
        project_id=__ret__.project_id,
        provider_auto_scaling_compute_max_instance_size=__ret__.provider_auto_scaling_compute_max_instance_size,
        provider_auto_scaling_compute_min_instance_size=__ret__.provider_auto_scaling_compute_min_instance_size,
        provider_backup_enabled=__ret__.provider_backup_enabled,
        provider_disk_iops=__ret__.provider_disk_iops,
        provider_disk_type_name=__ret__.provider_disk_type_name,
        provider_encrypt_ebs_volume=__ret__.provider_encrypt_ebs_volume,
        provider_instance_size_name=__ret__.provider_instance_size_name,
        provider_name=__ret__.provider_name,
        provider_region_name=__ret__.provider_region_name,
        provider_volume_type=__ret__.provider_volume_type,
        replication_factor=__ret__.replication_factor,
        replication_specs=__ret__.replication_specs,
        snapshot_backup_policies=__ret__.snapshot_backup_policies,
        srv_address=__ret__.srv_address,
        state_name=__ret__.state_name)
