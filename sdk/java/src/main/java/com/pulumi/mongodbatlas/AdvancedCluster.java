// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.mongodbatlas;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.mongodbatlas.AdvancedClusterArgs;
import com.pulumi.mongodbatlas.Utilities;
import com.pulumi.mongodbatlas.inputs.AdvancedClusterState;
import com.pulumi.mongodbatlas.outputs.AdvancedClusterAdvancedConfiguration;
import com.pulumi.mongodbatlas.outputs.AdvancedClusterBiConnector;
import com.pulumi.mongodbatlas.outputs.AdvancedClusterConnectionString;
import com.pulumi.mongodbatlas.outputs.AdvancedClusterLabel;
import com.pulumi.mongodbatlas.outputs.AdvancedClusterReplicationSpec;
import java.lang.Boolean;
import java.lang.Double;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * ## Example Usage
 * ### Example single provider and single region
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.mongodbatlas.AdvancedCluster;
 * import com.pulumi.mongodbatlas.AdvancedClusterArgs;
 * import com.pulumi.mongodbatlas.inputs.AdvancedClusterReplicationSpecArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var test = new AdvancedCluster(&#34;test&#34;, AdvancedClusterArgs.builder()        
 *             .clusterType(&#34;REPLICASET&#34;)
 *             .projectId(&#34;PROJECT ID&#34;)
 *             .replicationSpecs(AdvancedClusterReplicationSpecArgs.builder()
 *                 .regionConfigs(AdvancedClusterReplicationSpecRegionConfigArgs.builder()
 *                     .analyticsSpecs(AdvancedClusterReplicationSpecRegionConfigAnalyticsSpecsArgs.builder()
 *                         .instanceSize(&#34;M10&#34;)
 *                         .nodeCount(1)
 *                         .build())
 *                     .electableSpecs(AdvancedClusterReplicationSpecRegionConfigElectableSpecsArgs.builder()
 *                         .instanceSize(&#34;M10&#34;)
 *                         .nodeCount(3)
 *                         .build())
 *                     .priority(1)
 *                     .providerName(&#34;AWS&#34;)
 *                     .regionName(&#34;US_EAST_1&#34;)
 *                     .build())
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### Example Tenant Cluster
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.mongodbatlas.AdvancedCluster;
 * import com.pulumi.mongodbatlas.AdvancedClusterArgs;
 * import com.pulumi.mongodbatlas.inputs.AdvancedClusterReplicationSpecArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var test = new AdvancedCluster(&#34;test&#34;, AdvancedClusterArgs.builder()        
 *             .clusterType(&#34;REPLICASET&#34;)
 *             .projectId(&#34;PROJECT ID&#34;)
 *             .replicationSpecs(AdvancedClusterReplicationSpecArgs.builder()
 *                 .regionConfigs(AdvancedClusterReplicationSpecRegionConfigArgs.builder()
 *                     .backingProviderName(&#34;AWS&#34;)
 *                     .electableSpecs(AdvancedClusterReplicationSpecRegionConfigElectableSpecsArgs.builder()
 *                         .instanceSize(&#34;M5&#34;)
 *                         .build())
 *                     .priority(1)
 *                     .providerName(&#34;TENANT&#34;)
 *                     .regionName(&#34;US_EAST_1&#34;)
 *                     .build())
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### Example Multicloud.
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.mongodbatlas.AdvancedCluster;
 * import com.pulumi.mongodbatlas.AdvancedClusterArgs;
 * import com.pulumi.mongodbatlas.inputs.AdvancedClusterReplicationSpecArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var test = new AdvancedCluster(&#34;test&#34;, AdvancedClusterArgs.builder()        
 *             .clusterType(&#34;REPLICASET&#34;)
 *             .projectId(&#34;PROJECT ID&#34;)
 *             .replicationSpecs(AdvancedClusterReplicationSpecArgs.builder()
 *                 .regionConfigs(                
 *                     AdvancedClusterReplicationSpecRegionConfigArgs.builder()
 *                         .analyticsSpecs(AdvancedClusterReplicationSpecRegionConfigAnalyticsSpecsArgs.builder()
 *                             .instanceSize(&#34;M10&#34;)
 *                             .nodeCount(1)
 *                             .build())
 *                         .electableSpecs(AdvancedClusterReplicationSpecRegionConfigElectableSpecsArgs.builder()
 *                             .instanceSize(&#34;M10&#34;)
 *                             .nodeCount(3)
 *                             .build())
 *                         .priority(1)
 *                         .providerName(&#34;AWS&#34;)
 *                         .regionName(&#34;US_EAST_1&#34;)
 *                         .build(),
 *                     AdvancedClusterReplicationSpecRegionConfigArgs.builder()
 *                         .electableSpecs(AdvancedClusterReplicationSpecRegionConfigElectableSpecsArgs.builder()
 *                             .instanceSize(&#34;M10&#34;)
 *                             .nodeCount(2)
 *                             .build())
 *                         .priority(6)
 *                         .providerName(&#34;GCP&#34;)
 *                         .regionName(&#34;NORTH_AMERICA_NORTHEAST_1&#34;)
 *                         .build())
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Clusters can be imported using project ID and cluster name, in the format `PROJECTID-CLUSTERNAME`, e.g.
 * 
 * ```sh
 *  $ pulumi import mongodbatlas:index/advancedCluster:AdvancedCluster my_cluster 1112222b3bf99403840e8934-Cluster0
 * ```
 * 
 *  See detailed information for arguments and attributes[MongoDB API Advanced Clusters](https://docs.atlas.mongodb.com/reference/api/cluster-advanced/create-one-cluster-advanced/)
 * 
 */
@ResourceType(type="mongodbatlas:index/advancedCluster:AdvancedCluster")
public class AdvancedCluster extends com.pulumi.resources.CustomResource {
    @Export(name="advancedConfiguration", type=AdvancedClusterAdvancedConfiguration.class, parameters={})
    private Output<AdvancedClusterAdvancedConfiguration> advancedConfiguration;

    public Output<AdvancedClusterAdvancedConfiguration> advancedConfiguration() {
        return this.advancedConfiguration;
    }
    /**
     * Flag that indicates whether the cluster can perform backups.
     * If `true`, the cluster can perform backups. You must set this value to `true` for NVMe clusters.
     * 
     */
    @Export(name="backupEnabled", type=Boolean.class, parameters={})
    private Output<Boolean> backupEnabled;

    /**
     * @return Flag that indicates whether the cluster can perform backups.
     * If `true`, the cluster can perform backups. You must set this value to `true` for NVMe clusters.
     * 
     */
    public Output<Boolean> backupEnabled() {
        return this.backupEnabled;
    }
    /**
     * Configuration settings applied to BI Connector for Atlas on this cluster. The MongoDB Connector for Business Intelligence for Atlas (BI Connector) is only available for M10 and larger clusters. The BI Connector is a powerful tool which provides users SQL-based access to their MongoDB databases. As a result, the BI Connector performs operations which may be CPU and memory intensive. Given the limited hardware resources on M10 and M20 cluster tiers, you may experience performance degradation of the cluster when enabling the BI Connector. If this occurs, upgrade to an M30 or larger cluster or disable the BI Connector. See below.
     * 
     */
    @Export(name="biConnector", type=AdvancedClusterBiConnector.class, parameters={})
    private Output<AdvancedClusterBiConnector> biConnector;

    /**
     * @return Configuration settings applied to BI Connector for Atlas on this cluster. The MongoDB Connector for Business Intelligence for Atlas (BI Connector) is only available for M10 and larger clusters. The BI Connector is a powerful tool which provides users SQL-based access to their MongoDB databases. As a result, the BI Connector performs operations which may be CPU and memory intensive. Given the limited hardware resources on M10 and M20 cluster tiers, you may experience performance degradation of the cluster when enabling the BI Connector. If this occurs, upgrade to an M30 or larger cluster or disable the BI Connector. See below.
     * 
     */
    public Output<AdvancedClusterBiConnector> biConnector() {
        return this.biConnector;
    }
    /**
     * The cluster ID.
     * 
     */
    @Export(name="clusterId", type=String.class, parameters={})
    private Output<String> clusterId;

    /**
     * @return The cluster ID.
     * 
     */
    public Output<String> clusterId() {
        return this.clusterId;
    }
    /**
     * Type of the cluster that you want to create.
     * Accepted values include:
     * 
     */
    @Export(name="clusterType", type=String.class, parameters={})
    private Output<String> clusterType;

    /**
     * @return Type of the cluster that you want to create.
     * Accepted values include:
     * 
     */
    public Output<String> clusterType() {
        return this.clusterType;
    }
    /**
     * Set of connection strings that your applications use to connect to this cluster. More info in [Connection-strings](https://docs.mongodb.com/manual/reference/connection-string/). Use the parameters in this object to connect your applications to this cluster. To learn more about the formats of connection strings, see [Connection String Options](https://docs.atlas.mongodb.com/reference/faq/connection-changes/). NOTE: Atlas returns the contents of this object after the cluster is operational, not while it builds the cluster.
     * 
     */
    @Export(name="connectionStrings", type=List.class, parameters={AdvancedClusterConnectionString.class})
    private Output<List<AdvancedClusterConnectionString>> connectionStrings;

    /**
     * @return Set of connection strings that your applications use to connect to this cluster. More info in [Connection-strings](https://docs.mongodb.com/manual/reference/connection-string/). Use the parameters in this object to connect your applications to this cluster. To learn more about the formats of connection strings, see [Connection String Options](https://docs.atlas.mongodb.com/reference/faq/connection-changes/). NOTE: Atlas returns the contents of this object after the cluster is operational, not while it builds the cluster.
     * 
     */
    public Output<List<AdvancedClusterConnectionString>> connectionStrings() {
        return this.connectionStrings;
    }
    @Export(name="createDate", type=String.class, parameters={})
    private Output<String> createDate;

    public Output<String> createDate() {
        return this.createDate;
    }
    /**
     * Capacity, in gigabytes, of the host&#39;s root volume. Increase this number to add capacity, up to a maximum possible value of 4096 (i.e., 4 TB). This value must be a positive number. You can&#39;t set this value with clusters with local [NVMe SSDs](https://docs.atlas.mongodb.com/cluster-tier/#std-label-nvme-storage). The minimum disk size for dedicated clusters is 10 GB for AWS and GCP. If you specify diskSizeGB with a lower disk size, Atlas defaults to the minimum disk size value. If your cluster includes Azure nodes, this value must correspond to an existing Azure disk type (8, 16, 32, 64, 128, 256, 512, 1024, 2048, or 4095)Atlas calculates storage charges differently depending on whether you choose the default value or a custom value. The maximum value for disk storage cannot exceed 50 times the maximum RAM for the selected cluster. If you require additional storage space beyond this limitation, consider [upgrading your cluster](https://docs.atlas.mongodb.com/scale-cluster/#std-label-scale-cluster-instance) to a higher tier. If your cluster spans cloud service providers, this value defaults to the minimum default of the providers involved.
     * 
     */
    @Export(name="diskSizeGb", type=Double.class, parameters={})
    private Output<Double> diskSizeGb;

    /**
     * @return Capacity, in gigabytes, of the host&#39;s root volume. Increase this number to add capacity, up to a maximum possible value of 4096 (i.e., 4 TB). This value must be a positive number. You can&#39;t set this value with clusters with local [NVMe SSDs](https://docs.atlas.mongodb.com/cluster-tier/#std-label-nvme-storage). The minimum disk size for dedicated clusters is 10 GB for AWS and GCP. If you specify diskSizeGB with a lower disk size, Atlas defaults to the minimum disk size value. If your cluster includes Azure nodes, this value must correspond to an existing Azure disk type (8, 16, 32, 64, 128, 256, 512, 1024, 2048, or 4095)Atlas calculates storage charges differently depending on whether you choose the default value or a custom value. The maximum value for disk storage cannot exceed 50 times the maximum RAM for the selected cluster. If you require additional storage space beyond this limitation, consider [upgrading your cluster](https://docs.atlas.mongodb.com/scale-cluster/#std-label-scale-cluster-instance) to a higher tier. If your cluster spans cloud service providers, this value defaults to the minimum default of the providers involved.
     * 
     */
    public Output<Double> diskSizeGb() {
        return this.diskSizeGb;
    }
    /**
     * Possible values are AWS, GCP, AZURE or NONE.  Only needed if you desire to manage the keys, see [Encryption at Rest using Customer Key Management](https://docs.atlas.mongodb.com/security-kms-encryption/) for complete documentation.  You must configure encryption at rest for the Atlas project before enabling it on any cluster in the project. For Documentation, see [AWS](https://docs.atlas.mongodb.com/security-aws-kms/), [GCP](https://docs.atlas.mongodb.com/security-kms-encryption/) and [Azure](https://docs.atlas.mongodb.com/security-azure-kms/#std-label-security-azure-kms). Requirements are if `replication_specs.#.region_configs.#.&lt;type&gt;Specs.instance_size` is M10 or greater and `backup_enabled` is false or omitted.
     * 
     */
    @Export(name="encryptionAtRestProvider", type=String.class, parameters={})
    private Output<String> encryptionAtRestProvider;

    /**
     * @return Possible values are AWS, GCP, AZURE or NONE.  Only needed if you desire to manage the keys, see [Encryption at Rest using Customer Key Management](https://docs.atlas.mongodb.com/security-kms-encryption/) for complete documentation.  You must configure encryption at rest for the Atlas project before enabling it on any cluster in the project. For Documentation, see [AWS](https://docs.atlas.mongodb.com/security-aws-kms/), [GCP](https://docs.atlas.mongodb.com/security-kms-encryption/) and [Azure](https://docs.atlas.mongodb.com/security-azure-kms/#std-label-security-azure-kms). Requirements are if `replication_specs.#.region_configs.#.&lt;type&gt;Specs.instance_size` is M10 or greater and `backup_enabled` is false or omitted.
     * 
     */
    public Output<String> encryptionAtRestProvider() {
        return this.encryptionAtRestProvider;
    }
    /**
     * Configuration for the collection of key-value pairs that tag and categorize the cluster. See below.
     * 
     */
    @Export(name="labels", type=List.class, parameters={AdvancedClusterLabel.class})
    private Output<List<AdvancedClusterLabel>> labels;

    /**
     * @return Configuration for the collection of key-value pairs that tag and categorize the cluster. See below.
     * 
     */
    public Output<List<AdvancedClusterLabel>> labels() {
        return this.labels;
    }
    /**
     * Version of the cluster to deploy. Atlas supports the following MongoDB versions for M10+ clusters: `4.0`, `4.2`, `4.4`, or `5.0`. If omitted, Atlas deploys a cluster that runs MongoDB 4.4. If `replication_specs#.region_configs#.&lt;type&gt;Specs.instance_size`: `M0`, `M2` or `M5`, Atlas deploys MongoDB 4.4. Atlas always deploys the cluster with the latest stable release of the specified version.  If you set a value to this parameter and set `version_release_system` `CONTINUOUS`, the resource returns an error. Either clear this parameter or set `version_release_system`: `LTS`.
     * 
     */
    @Export(name="mongoDbMajorVersion", type=String.class, parameters={})
    private Output<String> mongoDbMajorVersion;

    /**
     * @return Version of the cluster to deploy. Atlas supports the following MongoDB versions for M10+ clusters: `4.0`, `4.2`, `4.4`, or `5.0`. If omitted, Atlas deploys a cluster that runs MongoDB 4.4. If `replication_specs#.region_configs#.&lt;type&gt;Specs.instance_size`: `M0`, `M2` or `M5`, Atlas deploys MongoDB 4.4. Atlas always deploys the cluster with the latest stable release of the specified version.  If you set a value to this parameter and set `version_release_system` `CONTINUOUS`, the resource returns an error. Either clear this parameter or set `version_release_system`: `LTS`.
     * 
     */
    public Output<String> mongoDbMajorVersion() {
        return this.mongoDbMajorVersion;
    }
    /**
     * Version of MongoDB the cluster runs, in `major-version`.`minor-version` format.
     * 
     */
    @Export(name="mongoDbVersion", type=String.class, parameters={})
    private Output<String> mongoDbVersion;

    /**
     * @return Version of MongoDB the cluster runs, in `major-version`.`minor-version` format.
     * 
     */
    public Output<String> mongoDbVersion() {
        return this.mongoDbVersion;
    }
    /**
     * Name of the cluster as it appears in Atlas. Once the cluster is created, its name cannot be changed.
     * 
     */
    @Export(name="name", type=String.class, parameters={})
    private Output<String> name;

    /**
     * @return Name of the cluster as it appears in Atlas. Once the cluster is created, its name cannot be changed.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    @Export(name="paused", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> paused;

    public Output<Optional<Boolean>> paused() {
        return Codegen.optional(this.paused);
    }
    /**
     * - Flag that indicates if the cluster uses Continuous Cloud Backup.
     * 
     */
    @Export(name="pitEnabled", type=Boolean.class, parameters={})
    private Output<Boolean> pitEnabled;

    /**
     * @return - Flag that indicates if the cluster uses Continuous Cloud Backup.
     * 
     */
    public Output<Boolean> pitEnabled() {
        return this.pitEnabled;
    }
    /**
     * Unique ID for the project to create the database user.
     * 
     */
    @Export(name="projectId", type=String.class, parameters={})
    private Output<String> projectId;

    /**
     * @return Unique ID for the project to create the database user.
     * 
     */
    public Output<String> projectId() {
        return this.projectId;
    }
    /**
     * Configuration for cluster regions and the hardware provisioned in them. See below
     * 
     */
    @Export(name="replicationSpecs", type=List.class, parameters={AdvancedClusterReplicationSpec.class})
    private Output<List<AdvancedClusterReplicationSpec>> replicationSpecs;

    /**
     * @return Configuration for cluster regions and the hardware provisioned in them. See below
     * 
     */
    public Output<List<AdvancedClusterReplicationSpec>> replicationSpecs() {
        return this.replicationSpecs;
    }
    /**
     * - Certificate Authority that MongoDB Atlas clusters use. You can specify ISRGROOTX1 (for ISRG Root X1).
     * 
     */
    @Export(name="rootCertType", type=String.class, parameters={})
    private Output<String> rootCertType;

    /**
     * @return - Certificate Authority that MongoDB Atlas clusters use. You can specify ISRGROOTX1 (for ISRG Root X1).
     * 
     */
    public Output<String> rootCertType() {
        return this.rootCertType;
    }
    /**
     * Current state of the cluster. The possible states are:
     * - IDLE
     * - CREATING
     * - UPDATING
     * - DELETING
     * - DELETED
     * - REPAIRING
     * 
     */
    @Export(name="stateName", type=String.class, parameters={})
    private Output<String> stateName;

    /**
     * @return Current state of the cluster. The possible states are:
     * - IDLE
     * - CREATING
     * - UPDATING
     * - DELETING
     * - DELETED
     * - REPAIRING
     * 
     */
    public Output<String> stateName() {
        return this.stateName;
    }
    /**
     * - Release cadence that Atlas uses for this cluster. This parameter defaults to `LTS`. If you set this field to `CONTINUOUS`, you must omit the `mongo_db_major_version` field. Atlas accepts:
     * - `CONTINUOUS`:  Atlas creates your cluster using the most recent MongoDB release. Atlas automatically updates your cluster to the latest major and rapid MongoDB releases as they become available.
     * - `LTS`: Atlas creates your cluster using the latest patch release of the MongoDB version that you specify in the mongoDBMajorVersion field. Atlas automatically updates your cluster to subsequent patch releases of this MongoDB version. Atlas doesn&#39;t update your cluster to newer rapid or major MongoDB releases as they become available.
     * 
     */
    @Export(name="versionReleaseSystem", type=String.class, parameters={})
    private Output</* @Nullable */ String> versionReleaseSystem;

    /**
     * @return - Release cadence that Atlas uses for this cluster. This parameter defaults to `LTS`. If you set this field to `CONTINUOUS`, you must omit the `mongo_db_major_version` field. Atlas accepts:
     * - `CONTINUOUS`:  Atlas creates your cluster using the most recent MongoDB release. Atlas automatically updates your cluster to the latest major and rapid MongoDB releases as they become available.
     * - `LTS`: Atlas creates your cluster using the latest patch release of the MongoDB version that you specify in the mongoDBMajorVersion field. Atlas automatically updates your cluster to subsequent patch releases of this MongoDB version. Atlas doesn&#39;t update your cluster to newer rapid or major MongoDB releases as they become available.
     * 
     */
    public Output<Optional<String>> versionReleaseSystem() {
        return Codegen.optional(this.versionReleaseSystem);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public AdvancedCluster(String name) {
        this(name, AdvancedClusterArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public AdvancedCluster(String name, AdvancedClusterArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public AdvancedCluster(String name, AdvancedClusterArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("mongodbatlas:index/advancedCluster:AdvancedCluster", name, args == null ? AdvancedClusterArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private AdvancedCluster(String name, Output<String> id, @Nullable AdvancedClusterState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("mongodbatlas:index/advancedCluster:AdvancedCluster", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static AdvancedCluster get(String name, Output<String> id, @Nullable AdvancedClusterState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new AdvancedCluster(name, id, state, options);
    }
}
