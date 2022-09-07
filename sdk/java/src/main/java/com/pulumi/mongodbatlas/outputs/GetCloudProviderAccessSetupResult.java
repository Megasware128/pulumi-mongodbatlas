// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.mongodbatlas.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.mongodbatlas.outputs.GetCloudProviderAccessSetupAwsConfig;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;

@CustomType
public final class GetCloudProviderAccessSetupResult {
    /**
     * @return aws related role information
     * 
     */
    private Map<String,String> aws;
    private List<GetCloudProviderAccessSetupAwsConfig> awsConfigs;
    /**
     * @return Date on which this role was created.
     * 
     */
    private String createdDate;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    private String projectId;
    private String providerName;
    private String roleId;

    private GetCloudProviderAccessSetupResult() {}
    /**
     * @return aws related role information
     * 
     */
    public Map<String,String> aws() {
        return this.aws;
    }
    public List<GetCloudProviderAccessSetupAwsConfig> awsConfigs() {
        return this.awsConfigs;
    }
    /**
     * @return Date on which this role was created.
     * 
     */
    public String createdDate() {
        return this.createdDate;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    public String projectId() {
        return this.projectId;
    }
    public String providerName() {
        return this.providerName;
    }
    public String roleId() {
        return this.roleId;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetCloudProviderAccessSetupResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private Map<String,String> aws;
        private List<GetCloudProviderAccessSetupAwsConfig> awsConfigs;
        private String createdDate;
        private String id;
        private String projectId;
        private String providerName;
        private String roleId;
        public Builder() {}
        public Builder(GetCloudProviderAccessSetupResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.aws = defaults.aws;
    	      this.awsConfigs = defaults.awsConfigs;
    	      this.createdDate = defaults.createdDate;
    	      this.id = defaults.id;
    	      this.projectId = defaults.projectId;
    	      this.providerName = defaults.providerName;
    	      this.roleId = defaults.roleId;
        }

        @CustomType.Setter
        public Builder aws(Map<String,String> aws) {
            this.aws = Objects.requireNonNull(aws);
            return this;
        }
        @CustomType.Setter
        public Builder awsConfigs(List<GetCloudProviderAccessSetupAwsConfig> awsConfigs) {
            this.awsConfigs = Objects.requireNonNull(awsConfigs);
            return this;
        }
        public Builder awsConfigs(GetCloudProviderAccessSetupAwsConfig... awsConfigs) {
            return awsConfigs(List.of(awsConfigs));
        }
        @CustomType.Setter
        public Builder createdDate(String createdDate) {
            this.createdDate = Objects.requireNonNull(createdDate);
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder projectId(String projectId) {
            this.projectId = Objects.requireNonNull(projectId);
            return this;
        }
        @CustomType.Setter
        public Builder providerName(String providerName) {
            this.providerName = Objects.requireNonNull(providerName);
            return this;
        }
        @CustomType.Setter
        public Builder roleId(String roleId) {
            this.roleId = Objects.requireNonNull(roleId);
            return this;
        }
        public GetCloudProviderAccessSetupResult build() {
            final var o = new GetCloudProviderAccessSetupResult();
            o.aws = aws;
            o.awsConfigs = awsConfigs;
            o.createdDate = createdDate;
            o.id = id;
            o.projectId = projectId;
            o.providerName = providerName;
            o.roleId = roleId;
            return o;
        }
    }
}
