// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.mongodbatlas.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class GetAdvancedClustersArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetAdvancedClustersArgs Empty = new GetAdvancedClustersArgs();

    /**
     * The unique ID for the project to get the clusters.
     * 
     */
    @Import(name="projectId", required=true)
    private Output<String> projectId;

    /**
     * @return The unique ID for the project to get the clusters.
     * 
     */
    public Output<String> projectId() {
        return this.projectId;
    }

    private GetAdvancedClustersArgs() {}

    private GetAdvancedClustersArgs(GetAdvancedClustersArgs $) {
        this.projectId = $.projectId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetAdvancedClustersArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetAdvancedClustersArgs $;

        public Builder() {
            $ = new GetAdvancedClustersArgs();
        }

        public Builder(GetAdvancedClustersArgs defaults) {
            $ = new GetAdvancedClustersArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param projectId The unique ID for the project to get the clusters.
         * 
         * @return builder
         * 
         */
        public Builder projectId(Output<String> projectId) {
            $.projectId = projectId;
            return this;
        }

        /**
         * @param projectId The unique ID for the project to get the clusters.
         * 
         * @return builder
         * 
         */
        public Builder projectId(String projectId) {
            return projectId(Output.of(projectId));
        }

        public GetAdvancedClustersArgs build() {
            $.projectId = Objects.requireNonNull($.projectId, "expected parameter 'projectId' to be non-null");
            return $;
        }
    }

}
