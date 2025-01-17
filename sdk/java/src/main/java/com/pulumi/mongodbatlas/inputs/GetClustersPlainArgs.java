// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.mongodbatlas.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class GetClustersPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetClustersPlainArgs Empty = new GetClustersPlainArgs();

    /**
     * The unique ID for the project to get the clusters.
     * 
     */
    @Import(name="projectId", required=true)
    private String projectId;

    /**
     * @return The unique ID for the project to get the clusters.
     * 
     */
    public String projectId() {
        return this.projectId;
    }

    private GetClustersPlainArgs() {}

    private GetClustersPlainArgs(GetClustersPlainArgs $) {
        this.projectId = $.projectId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetClustersPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetClustersPlainArgs $;

        public Builder() {
            $ = new GetClustersPlainArgs();
        }

        public Builder(GetClustersPlainArgs defaults) {
            $ = new GetClustersPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param projectId The unique ID for the project to get the clusters.
         * 
         * @return builder
         * 
         */
        public Builder projectId(String projectId) {
            $.projectId = projectId;
            return this;
        }

        public GetClustersPlainArgs build() {
            $.projectId = Objects.requireNonNull($.projectId, "expected parameter 'projectId' to be non-null");
            return $;
        }
    }

}
