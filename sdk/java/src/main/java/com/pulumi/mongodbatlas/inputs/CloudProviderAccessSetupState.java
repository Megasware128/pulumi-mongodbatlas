// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.mongodbatlas.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.mongodbatlas.inputs.CloudProviderAccessSetupAwsConfigArgs;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class CloudProviderAccessSetupState extends com.pulumi.resources.ResourceArgs {

    public static final CloudProviderAccessSetupState Empty = new CloudProviderAccessSetupState();

    /**
     * @deprecated
     * use aws_config instead
     * 
     */
    @Deprecated /* use aws_config instead */
    @Import(name="aws")
    private @Nullable Output<Map<String,String>> aws;

    /**
     * @deprecated
     * use aws_config instead
     * 
     */
    @Deprecated /* use aws_config instead */
    public Optional<Output<Map<String,String>>> aws() {
        return Optional.ofNullable(this.aws);
    }

    @Import(name="awsConfigs")
    private @Nullable Output<List<CloudProviderAccessSetupAwsConfigArgs>> awsConfigs;

    public Optional<Output<List<CloudProviderAccessSetupAwsConfigArgs>>> awsConfigs() {
        return Optional.ofNullable(this.awsConfigs);
    }

    @Import(name="createdDate")
    private @Nullable Output<String> createdDate;

    public Optional<Output<String>> createdDate() {
        return Optional.ofNullable(this.createdDate);
    }

    @Import(name="projectId")
    private @Nullable Output<String> projectId;

    public Optional<Output<String>> projectId() {
        return Optional.ofNullable(this.projectId);
    }

    @Import(name="providerName")
    private @Nullable Output<String> providerName;

    public Optional<Output<String>> providerName() {
        return Optional.ofNullable(this.providerName);
    }

    @Import(name="roleId")
    private @Nullable Output<String> roleId;

    public Optional<Output<String>> roleId() {
        return Optional.ofNullable(this.roleId);
    }

    private CloudProviderAccessSetupState() {}

    private CloudProviderAccessSetupState(CloudProviderAccessSetupState $) {
        this.aws = $.aws;
        this.awsConfigs = $.awsConfigs;
        this.createdDate = $.createdDate;
        this.projectId = $.projectId;
        this.providerName = $.providerName;
        this.roleId = $.roleId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(CloudProviderAccessSetupState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private CloudProviderAccessSetupState $;

        public Builder() {
            $ = new CloudProviderAccessSetupState();
        }

        public Builder(CloudProviderAccessSetupState defaults) {
            $ = new CloudProviderAccessSetupState(Objects.requireNonNull(defaults));
        }

        /**
         * @return builder
         * 
         * @deprecated
         * use aws_config instead
         * 
         */
        @Deprecated /* use aws_config instead */
        public Builder aws(@Nullable Output<Map<String,String>> aws) {
            $.aws = aws;
            return this;
        }

        /**
         * @return builder
         * 
         * @deprecated
         * use aws_config instead
         * 
         */
        @Deprecated /* use aws_config instead */
        public Builder aws(Map<String,String> aws) {
            return aws(Output.of(aws));
        }

        public Builder awsConfigs(@Nullable Output<List<CloudProviderAccessSetupAwsConfigArgs>> awsConfigs) {
            $.awsConfigs = awsConfigs;
            return this;
        }

        public Builder awsConfigs(List<CloudProviderAccessSetupAwsConfigArgs> awsConfigs) {
            return awsConfigs(Output.of(awsConfigs));
        }

        public Builder awsConfigs(CloudProviderAccessSetupAwsConfigArgs... awsConfigs) {
            return awsConfigs(List.of(awsConfigs));
        }

        public Builder createdDate(@Nullable Output<String> createdDate) {
            $.createdDate = createdDate;
            return this;
        }

        public Builder createdDate(String createdDate) {
            return createdDate(Output.of(createdDate));
        }

        public Builder projectId(@Nullable Output<String> projectId) {
            $.projectId = projectId;
            return this;
        }

        public Builder projectId(String projectId) {
            return projectId(Output.of(projectId));
        }

        public Builder providerName(@Nullable Output<String> providerName) {
            $.providerName = providerName;
            return this;
        }

        public Builder providerName(String providerName) {
            return providerName(Output.of(providerName));
        }

        public Builder roleId(@Nullable Output<String> roleId) {
            $.roleId = roleId;
            return this;
        }

        public Builder roleId(String roleId) {
            return roleId(Output.of(roleId));
        }

        public CloudProviderAccessSetupState build() {
            return $;
        }
    }

}