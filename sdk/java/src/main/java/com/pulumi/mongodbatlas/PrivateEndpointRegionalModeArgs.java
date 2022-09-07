// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.mongodbatlas;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class PrivateEndpointRegionalModeArgs extends com.pulumi.resources.ResourceArgs {

    public static final PrivateEndpointRegionalModeArgs Empty = new PrivateEndpointRegionalModeArgs();

    /**
     * Flag that indicates whether the regionalized private endpoint setting is enabled for the project.   Set this value to true to create more than one private endpoint in a cloud provider region to connect to multi-region and global Atlas sharded clusters. You can enable this setting only if your Atlas project contains no replica sets. You can&#39;t disable this setting if you have:
     * * More than one private endpoint in more than one region, or
     * * More than one private endpoint in one region and one private endpoint in one or more regions.
     *   You can create only sharded clusters when you enable the regionalized private endpoint setting. You can&#39;t create replica sets.
     * 
     */
    @Import(name="enabled")
    private @Nullable Output<Boolean> enabled;

    /**
     * @return Flag that indicates whether the regionalized private endpoint setting is enabled for the project.   Set this value to true to create more than one private endpoint in a cloud provider region to connect to multi-region and global Atlas sharded clusters. You can enable this setting only if your Atlas project contains no replica sets. You can&#39;t disable this setting if you have:
     * * More than one private endpoint in more than one region, or
     * * More than one private endpoint in one region and one private endpoint in one or more regions.
     *   You can create only sharded clusters when you enable the regionalized private endpoint setting. You can&#39;t create replica sets.
     * 
     */
    public Optional<Output<Boolean>> enabled() {
        return Optional.ofNullable(this.enabled);
    }

    /**
     * Unique identifier for the project.
     * 
     */
    @Import(name="projectId", required=true)
    private Output<String> projectId;

    /**
     * @return Unique identifier for the project.
     * 
     */
    public Output<String> projectId() {
        return this.projectId;
    }

    private PrivateEndpointRegionalModeArgs() {}

    private PrivateEndpointRegionalModeArgs(PrivateEndpointRegionalModeArgs $) {
        this.enabled = $.enabled;
        this.projectId = $.projectId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(PrivateEndpointRegionalModeArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private PrivateEndpointRegionalModeArgs $;

        public Builder() {
            $ = new PrivateEndpointRegionalModeArgs();
        }

        public Builder(PrivateEndpointRegionalModeArgs defaults) {
            $ = new PrivateEndpointRegionalModeArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param enabled Flag that indicates whether the regionalized private endpoint setting is enabled for the project.   Set this value to true to create more than one private endpoint in a cloud provider region to connect to multi-region and global Atlas sharded clusters. You can enable this setting only if your Atlas project contains no replica sets. You can&#39;t disable this setting if you have:
         * * More than one private endpoint in more than one region, or
         * * More than one private endpoint in one region and one private endpoint in one or more regions.
         *   You can create only sharded clusters when you enable the regionalized private endpoint setting. You can&#39;t create replica sets.
         * 
         * @return builder
         * 
         */
        public Builder enabled(@Nullable Output<Boolean> enabled) {
            $.enabled = enabled;
            return this;
        }

        /**
         * @param enabled Flag that indicates whether the regionalized private endpoint setting is enabled for the project.   Set this value to true to create more than one private endpoint in a cloud provider region to connect to multi-region and global Atlas sharded clusters. You can enable this setting only if your Atlas project contains no replica sets. You can&#39;t disable this setting if you have:
         * * More than one private endpoint in more than one region, or
         * * More than one private endpoint in one region and one private endpoint in one or more regions.
         *   You can create only sharded clusters when you enable the regionalized private endpoint setting. You can&#39;t create replica sets.
         * 
         * @return builder
         * 
         */
        public Builder enabled(Boolean enabled) {
            return enabled(Output.of(enabled));
        }

        /**
         * @param projectId Unique identifier for the project.
         * 
         * @return builder
         * 
         */
        public Builder projectId(Output<String> projectId) {
            $.projectId = projectId;
            return this;
        }

        /**
         * @param projectId Unique identifier for the project.
         * 
         * @return builder
         * 
         */
        public Builder projectId(String projectId) {
            return projectId(Output.of(projectId));
        }

        public PrivateEndpointRegionalModeArgs build() {
            $.projectId = Objects.requireNonNull($.projectId, "expected parameter 'projectId' to be non-null");
            return $;
        }
    }

}
