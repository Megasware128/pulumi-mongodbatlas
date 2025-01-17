// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.mongodbatlas;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class PrivateLinkEndpointArgs extends com.pulumi.resources.ResourceArgs {

    public static final PrivateLinkEndpointArgs Empty = new PrivateLinkEndpointArgs();

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

    /**
     * Name of the cloud provider for which you want to create the private endpoint service. Atlas accepts `AWS`, `AZURE` or `GCP`.
     * 
     */
    @Import(name="providerName", required=true)
    private Output<String> providerName;

    /**
     * @return Name of the cloud provider for which you want to create the private endpoint service. Atlas accepts `AWS`, `AZURE` or `GCP`.
     * 
     */
    public Output<String> providerName() {
        return this.providerName;
    }

    /**
     * Cloud provider region in which you want to create the private endpoint connection.
     * Accepted values are: [AWS regions](https://docs.atlas.mongodb.com/reference/amazon-aws/#amazon-aws), [AZURE regions](https://docs.atlas.mongodb.com/reference/microsoft-azure/#microsoft-azure) and [GCP regions](https://docs.atlas.mongodb.com/reference/google-gcp/#std-label-google-gcp)
     * 
     */
    @Import(name="region", required=true)
    private Output<String> region;

    /**
     * @return Cloud provider region in which you want to create the private endpoint connection.
     * Accepted values are: [AWS regions](https://docs.atlas.mongodb.com/reference/amazon-aws/#amazon-aws), [AZURE regions](https://docs.atlas.mongodb.com/reference/microsoft-azure/#microsoft-azure) and [GCP regions](https://docs.atlas.mongodb.com/reference/google-gcp/#std-label-google-gcp)
     * 
     */
    public Output<String> region() {
        return this.region;
    }

    private PrivateLinkEndpointArgs() {}

    private PrivateLinkEndpointArgs(PrivateLinkEndpointArgs $) {
        this.projectId = $.projectId;
        this.providerName = $.providerName;
        this.region = $.region;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(PrivateLinkEndpointArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private PrivateLinkEndpointArgs $;

        public Builder() {
            $ = new PrivateLinkEndpointArgs();
        }

        public Builder(PrivateLinkEndpointArgs defaults) {
            $ = new PrivateLinkEndpointArgs(Objects.requireNonNull(defaults));
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

        /**
         * @param providerName Name of the cloud provider for which you want to create the private endpoint service. Atlas accepts `AWS`, `AZURE` or `GCP`.
         * 
         * @return builder
         * 
         */
        public Builder providerName(Output<String> providerName) {
            $.providerName = providerName;
            return this;
        }

        /**
         * @param providerName Name of the cloud provider for which you want to create the private endpoint service. Atlas accepts `AWS`, `AZURE` or `GCP`.
         * 
         * @return builder
         * 
         */
        public Builder providerName(String providerName) {
            return providerName(Output.of(providerName));
        }

        /**
         * @param region Cloud provider region in which you want to create the private endpoint connection.
         * Accepted values are: [AWS regions](https://docs.atlas.mongodb.com/reference/amazon-aws/#amazon-aws), [AZURE regions](https://docs.atlas.mongodb.com/reference/microsoft-azure/#microsoft-azure) and [GCP regions](https://docs.atlas.mongodb.com/reference/google-gcp/#std-label-google-gcp)
         * 
         * @return builder
         * 
         */
        public Builder region(Output<String> region) {
            $.region = region;
            return this;
        }

        /**
         * @param region Cloud provider region in which you want to create the private endpoint connection.
         * Accepted values are: [AWS regions](https://docs.atlas.mongodb.com/reference/amazon-aws/#amazon-aws), [AZURE regions](https://docs.atlas.mongodb.com/reference/microsoft-azure/#microsoft-azure) and [GCP regions](https://docs.atlas.mongodb.com/reference/google-gcp/#std-label-google-gcp)
         * 
         * @return builder
         * 
         */
        public Builder region(String region) {
            return region(Output.of(region));
        }

        public PrivateLinkEndpointArgs build() {
            $.projectId = Objects.requireNonNull($.projectId, "expected parameter 'projectId' to be non-null");
            $.providerName = Objects.requireNonNull($.providerName, "expected parameter 'providerName' to be non-null");
            $.region = Objects.requireNonNull($.region, "expected parameter 'region' to be non-null");
            return $;
        }
    }

}
