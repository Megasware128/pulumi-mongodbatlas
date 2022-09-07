// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.mongodbatlas.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.mongodbatlas.outputs.GetAdvancedClusterAdvancedConfiguration;
import com.pulumi.mongodbatlas.outputs.GetAdvancedClusterBiConnector;
import com.pulumi.mongodbatlas.outputs.GetAdvancedClusterConnectionString;
import com.pulumi.mongodbatlas.outputs.GetAdvancedClusterLabel;
import com.pulumi.mongodbatlas.outputs.GetAdvancedClusterReplicationSpec;
import java.lang.Boolean;
import java.lang.Double;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetAdvancedClusterResult {
    /**
     * @return Get the advanced configuration options. See Advanced Configuration below for more details.
     * 
     */
    private List<GetAdvancedClusterAdvancedConfiguration> advancedConfigurations;
    private Boolean backupEnabled;
    /**
     * @return Configuration settings applied to BI Connector for Atlas on this cluster. See below.
     * 
     */
    private List<GetAdvancedClusterBiConnector> biConnectors;
    /**
     * @return Type of the cluster that you want to create.
     * 
     */
    private String clusterType;
    /**
     * @return Set of connection strings that your applications use to connect to this cluster. More info in [Connection-strings](https://docs.mongodb.com/manual/reference/connection-string/). Use the parameters in this object to connect your applications to this cluster. To learn more about the formats of connection strings, see [Connection String Options](https://docs.atlas.mongodb.com/reference/faq/connection-changes/). NOTE: Atlas returns the contents of this object after the cluster is operational, not while it builds the cluster.
     * 
     */
    private List<GetAdvancedClusterConnectionString> connectionStrings;
    private String createDate;
    /**
     * @return Capacity, in gigabytes, of the host&#39;s root volume.
     * 
     */
    private Double diskSizeGb;
    /**
     * @return Possible values are AWS, GCP, AZURE or NONE.
     * 
     */
    private String encryptionAtRestProvider;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return Configuration for the collection of key-value pairs that tag and categorize the cluster. See below.
     * 
     */
    private List<GetAdvancedClusterLabel> labels;
    /**
     * @return Version of the cluster to deploy.
     * 
     */
    private String mongoDbMajorVersion;
    /**
     * @return Version of MongoDB the cluster runs, in `major-version`.`minor-version` format.
     * 
     */
    private String mongoDbVersion;
    private String name;
    /**
     * @return Flag that indicates whether the cluster is paused or not.
     * 
     */
    private Boolean paused;
    /**
     * @return Flag that indicates if the cluster uses Continuous Cloud Backup.
     * 
     */
    private Boolean pitEnabled;
    private String projectId;
    /**
     * @return Configuration for cluster regions and the hardware provisioned in them. See below
     * 
     */
    private List<GetAdvancedClusterReplicationSpec> replicationSpecs;
    /**
     * @return Certificate Authority that MongoDB Atlas clusters use.
     * 
     */
    private String rootCertType;
    /**
     * @return Current state of the cluster. The possible states are:
     * 
     */
    private String stateName;
    /**
     * @return Release cadence that Atlas uses for this cluster.
     * 
     */
    private String versionReleaseSystem;

    private GetAdvancedClusterResult() {}
    /**
     * @return Get the advanced configuration options. See Advanced Configuration below for more details.
     * 
     */
    public List<GetAdvancedClusterAdvancedConfiguration> advancedConfigurations() {
        return this.advancedConfigurations;
    }
    public Boolean backupEnabled() {
        return this.backupEnabled;
    }
    /**
     * @return Configuration settings applied to BI Connector for Atlas on this cluster. See below.
     * 
     */
    public List<GetAdvancedClusterBiConnector> biConnectors() {
        return this.biConnectors;
    }
    /**
     * @return Type of the cluster that you want to create.
     * 
     */
    public String clusterType() {
        return this.clusterType;
    }
    /**
     * @return Set of connection strings that your applications use to connect to this cluster. More info in [Connection-strings](https://docs.mongodb.com/manual/reference/connection-string/). Use the parameters in this object to connect your applications to this cluster. To learn more about the formats of connection strings, see [Connection String Options](https://docs.atlas.mongodb.com/reference/faq/connection-changes/). NOTE: Atlas returns the contents of this object after the cluster is operational, not while it builds the cluster.
     * 
     */
    public List<GetAdvancedClusterConnectionString> connectionStrings() {
        return this.connectionStrings;
    }
    public String createDate() {
        return this.createDate;
    }
    /**
     * @return Capacity, in gigabytes, of the host&#39;s root volume.
     * 
     */
    public Double diskSizeGb() {
        return this.diskSizeGb;
    }
    /**
     * @return Possible values are AWS, GCP, AZURE or NONE.
     * 
     */
    public String encryptionAtRestProvider() {
        return this.encryptionAtRestProvider;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return Configuration for the collection of key-value pairs that tag and categorize the cluster. See below.
     * 
     */
    public List<GetAdvancedClusterLabel> labels() {
        return this.labels;
    }
    /**
     * @return Version of the cluster to deploy.
     * 
     */
    public String mongoDbMajorVersion() {
        return this.mongoDbMajorVersion;
    }
    /**
     * @return Version of MongoDB the cluster runs, in `major-version`.`minor-version` format.
     * 
     */
    public String mongoDbVersion() {
        return this.mongoDbVersion;
    }
    public String name() {
        return this.name;
    }
    /**
     * @return Flag that indicates whether the cluster is paused or not.
     * 
     */
    public Boolean paused() {
        return this.paused;
    }
    /**
     * @return Flag that indicates if the cluster uses Continuous Cloud Backup.
     * 
     */
    public Boolean pitEnabled() {
        return this.pitEnabled;
    }
    public String projectId() {
        return this.projectId;
    }
    /**
     * @return Configuration for cluster regions and the hardware provisioned in them. See below
     * 
     */
    public List<GetAdvancedClusterReplicationSpec> replicationSpecs() {
        return this.replicationSpecs;
    }
    /**
     * @return Certificate Authority that MongoDB Atlas clusters use.
     * 
     */
    public String rootCertType() {
        return this.rootCertType;
    }
    /**
     * @return Current state of the cluster. The possible states are:
     * 
     */
    public String stateName() {
        return this.stateName;
    }
    /**
     * @return Release cadence that Atlas uses for this cluster.
     * 
     */
    public String versionReleaseSystem() {
        return this.versionReleaseSystem;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetAdvancedClusterResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private List<GetAdvancedClusterAdvancedConfiguration> advancedConfigurations;
        private Boolean backupEnabled;
        private List<GetAdvancedClusterBiConnector> biConnectors;
        private String clusterType;
        private List<GetAdvancedClusterConnectionString> connectionStrings;
        private String createDate;
        private Double diskSizeGb;
        private String encryptionAtRestProvider;
        private String id;
        private List<GetAdvancedClusterLabel> labels;
        private String mongoDbMajorVersion;
        private String mongoDbVersion;
        private String name;
        private Boolean paused;
        private Boolean pitEnabled;
        private String projectId;
        private List<GetAdvancedClusterReplicationSpec> replicationSpecs;
        private String rootCertType;
        private String stateName;
        private String versionReleaseSystem;
        public Builder() {}
        public Builder(GetAdvancedClusterResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.advancedConfigurations = defaults.advancedConfigurations;
    	      this.backupEnabled = defaults.backupEnabled;
    	      this.biConnectors = defaults.biConnectors;
    	      this.clusterType = defaults.clusterType;
    	      this.connectionStrings = defaults.connectionStrings;
    	      this.createDate = defaults.createDate;
    	      this.diskSizeGb = defaults.diskSizeGb;
    	      this.encryptionAtRestProvider = defaults.encryptionAtRestProvider;
    	      this.id = defaults.id;
    	      this.labels = defaults.labels;
    	      this.mongoDbMajorVersion = defaults.mongoDbMajorVersion;
    	      this.mongoDbVersion = defaults.mongoDbVersion;
    	      this.name = defaults.name;
    	      this.paused = defaults.paused;
    	      this.pitEnabled = defaults.pitEnabled;
    	      this.projectId = defaults.projectId;
    	      this.replicationSpecs = defaults.replicationSpecs;
    	      this.rootCertType = defaults.rootCertType;
    	      this.stateName = defaults.stateName;
    	      this.versionReleaseSystem = defaults.versionReleaseSystem;
        }

        @CustomType.Setter
        public Builder advancedConfigurations(List<GetAdvancedClusterAdvancedConfiguration> advancedConfigurations) {
            this.advancedConfigurations = Objects.requireNonNull(advancedConfigurations);
            return this;
        }
        public Builder advancedConfigurations(GetAdvancedClusterAdvancedConfiguration... advancedConfigurations) {
            return advancedConfigurations(List.of(advancedConfigurations));
        }
        @CustomType.Setter
        public Builder backupEnabled(Boolean backupEnabled) {
            this.backupEnabled = Objects.requireNonNull(backupEnabled);
            return this;
        }
        @CustomType.Setter
        public Builder biConnectors(List<GetAdvancedClusterBiConnector> biConnectors) {
            this.biConnectors = Objects.requireNonNull(biConnectors);
            return this;
        }
        public Builder biConnectors(GetAdvancedClusterBiConnector... biConnectors) {
            return biConnectors(List.of(biConnectors));
        }
        @CustomType.Setter
        public Builder clusterType(String clusterType) {
            this.clusterType = Objects.requireNonNull(clusterType);
            return this;
        }
        @CustomType.Setter
        public Builder connectionStrings(List<GetAdvancedClusterConnectionString> connectionStrings) {
            this.connectionStrings = Objects.requireNonNull(connectionStrings);
            return this;
        }
        public Builder connectionStrings(GetAdvancedClusterConnectionString... connectionStrings) {
            return connectionStrings(List.of(connectionStrings));
        }
        @CustomType.Setter
        public Builder createDate(String createDate) {
            this.createDate = Objects.requireNonNull(createDate);
            return this;
        }
        @CustomType.Setter
        public Builder diskSizeGb(Double diskSizeGb) {
            this.diskSizeGb = Objects.requireNonNull(diskSizeGb);
            return this;
        }
        @CustomType.Setter
        public Builder encryptionAtRestProvider(String encryptionAtRestProvider) {
            this.encryptionAtRestProvider = Objects.requireNonNull(encryptionAtRestProvider);
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder labels(List<GetAdvancedClusterLabel> labels) {
            this.labels = Objects.requireNonNull(labels);
            return this;
        }
        public Builder labels(GetAdvancedClusterLabel... labels) {
            return labels(List.of(labels));
        }
        @CustomType.Setter
        public Builder mongoDbMajorVersion(String mongoDbMajorVersion) {
            this.mongoDbMajorVersion = Objects.requireNonNull(mongoDbMajorVersion);
            return this;
        }
        @CustomType.Setter
        public Builder mongoDbVersion(String mongoDbVersion) {
            this.mongoDbVersion = Objects.requireNonNull(mongoDbVersion);
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            this.name = Objects.requireNonNull(name);
            return this;
        }
        @CustomType.Setter
        public Builder paused(Boolean paused) {
            this.paused = Objects.requireNonNull(paused);
            return this;
        }
        @CustomType.Setter
        public Builder pitEnabled(Boolean pitEnabled) {
            this.pitEnabled = Objects.requireNonNull(pitEnabled);
            return this;
        }
        @CustomType.Setter
        public Builder projectId(String projectId) {
            this.projectId = Objects.requireNonNull(projectId);
            return this;
        }
        @CustomType.Setter
        public Builder replicationSpecs(List<GetAdvancedClusterReplicationSpec> replicationSpecs) {
            this.replicationSpecs = Objects.requireNonNull(replicationSpecs);
            return this;
        }
        public Builder replicationSpecs(GetAdvancedClusterReplicationSpec... replicationSpecs) {
            return replicationSpecs(List.of(replicationSpecs));
        }
        @CustomType.Setter
        public Builder rootCertType(String rootCertType) {
            this.rootCertType = Objects.requireNonNull(rootCertType);
            return this;
        }
        @CustomType.Setter
        public Builder stateName(String stateName) {
            this.stateName = Objects.requireNonNull(stateName);
            return this;
        }
        @CustomType.Setter
        public Builder versionReleaseSystem(String versionReleaseSystem) {
            this.versionReleaseSystem = Objects.requireNonNull(versionReleaseSystem);
            return this;
        }
        public GetAdvancedClusterResult build() {
            final var o = new GetAdvancedClusterResult();
            o.advancedConfigurations = advancedConfigurations;
            o.backupEnabled = backupEnabled;
            o.biConnectors = biConnectors;
            o.clusterType = clusterType;
            o.connectionStrings = connectionStrings;
            o.createDate = createDate;
            o.diskSizeGb = diskSizeGb;
            o.encryptionAtRestProvider = encryptionAtRestProvider;
            o.id = id;
            o.labels = labels;
            o.mongoDbMajorVersion = mongoDbMajorVersion;
            o.mongoDbVersion = mongoDbVersion;
            o.name = name;
            o.paused = paused;
            o.pitEnabled = pitEnabled;
            o.projectId = projectId;
            o.replicationSpecs = replicationSpecs;
            o.rootCertType = rootCertType;
            o.stateName = stateName;
            o.versionReleaseSystem = versionReleaseSystem;
            return o;
        }
    }
}
