// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.mongodbatlas;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;


public final class CustomDnsConfigurationClusterAwsArgs extends com.pulumi.resources.ResourceArgs {

    public static final CustomDnsConfigurationClusterAwsArgs Empty = new CustomDnsConfigurationClusterAwsArgs();

    /**
     * Indicates whether the project&#39;s clusters deployed to AWS use custom DNS. If `true`, the `Get All Clusters` and `Get One Cluster` endpoints return the `connectionStrings.private` and `connectionStrings.privateSrv` fields for clusters deployed to AWS .
     * 
     */
    @Import(name="enabled", required=true)
    private Output<Boolean> enabled;

    /**
     * @return Indicates whether the project&#39;s clusters deployed to AWS use custom DNS. If `true`, the `Get All Clusters` and `Get One Cluster` endpoints return the `connectionStrings.private` and `connectionStrings.privateSrv` fields for clusters deployed to AWS .
     * 
     */
    public Output<Boolean> enabled() {
        return this.enabled;
    }

    /**
     * Required 	Unique identifier for the project.
     * 
     */
    @Import(name="projectId", required=true)
    private Output<String> projectId;

    /**
     * @return Required 	Unique identifier for the project.
     * 
     */
    public Output<String> projectId() {
        return this.projectId;
    }

    private CustomDnsConfigurationClusterAwsArgs() {}

    private CustomDnsConfigurationClusterAwsArgs(CustomDnsConfigurationClusterAwsArgs $) {
        this.enabled = $.enabled;
        this.projectId = $.projectId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(CustomDnsConfigurationClusterAwsArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private CustomDnsConfigurationClusterAwsArgs $;

        public Builder() {
            $ = new CustomDnsConfigurationClusterAwsArgs();
        }

        public Builder(CustomDnsConfigurationClusterAwsArgs defaults) {
            $ = new CustomDnsConfigurationClusterAwsArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param enabled Indicates whether the project&#39;s clusters deployed to AWS use custom DNS. If `true`, the `Get All Clusters` and `Get One Cluster` endpoints return the `connectionStrings.private` and `connectionStrings.privateSrv` fields for clusters deployed to AWS .
         * 
         * @return builder
         * 
         */
        public Builder enabled(Output<Boolean> enabled) {
            $.enabled = enabled;
            return this;
        }

        /**
         * @param enabled Indicates whether the project&#39;s clusters deployed to AWS use custom DNS. If `true`, the `Get All Clusters` and `Get One Cluster` endpoints return the `connectionStrings.private` and `connectionStrings.privateSrv` fields for clusters deployed to AWS .
         * 
         * @return builder
         * 
         */
        public Builder enabled(Boolean enabled) {
            return enabled(Output.of(enabled));
        }

        /**
         * @param projectId Required 	Unique identifier for the project.
         * 
         * @return builder
         * 
         */
        public Builder projectId(Output<String> projectId) {
            $.projectId = projectId;
            return this;
        }

        /**
         * @param projectId Required 	Unique identifier for the project.
         * 
         * @return builder
         * 
         */
        public Builder projectId(String projectId) {
            return projectId(Output.of(projectId));
        }

        public CustomDnsConfigurationClusterAwsArgs build() {
            $.enabled = Objects.requireNonNull($.enabled, "expected parameter 'enabled' to be non-null");
            $.projectId = Objects.requireNonNull($.projectId, "expected parameter 'projectId' to be non-null");
            return $;
        }
    }

}