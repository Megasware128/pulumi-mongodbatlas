// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.mongodbatlas.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class GetNetworkContainerPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetNetworkContainerPlainArgs Empty = new GetNetworkContainerPlainArgs();

    /**
     * The Network Peering Container ID.
     * 
     */
    @Import(name="containerId", required=true)
    private String containerId;

    /**
     * @return The Network Peering Container ID.
     * 
     */
    public String containerId() {
        return this.containerId;
    }

    /**
     * The unique ID for the project to create the database user.
     * 
     */
    @Import(name="projectId", required=true)
    private String projectId;

    /**
     * @return The unique ID for the project to create the database user.
     * 
     */
    public String projectId() {
        return this.projectId;
    }

    private GetNetworkContainerPlainArgs() {}

    private GetNetworkContainerPlainArgs(GetNetworkContainerPlainArgs $) {
        this.containerId = $.containerId;
        this.projectId = $.projectId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetNetworkContainerPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetNetworkContainerPlainArgs $;

        public Builder() {
            $ = new GetNetworkContainerPlainArgs();
        }

        public Builder(GetNetworkContainerPlainArgs defaults) {
            $ = new GetNetworkContainerPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param containerId The Network Peering Container ID.
         * 
         * @return builder
         * 
         */
        public Builder containerId(String containerId) {
            $.containerId = containerId;
            return this;
        }

        /**
         * @param projectId The unique ID for the project to create the database user.
         * 
         * @return builder
         * 
         */
        public Builder projectId(String projectId) {
            $.projectId = projectId;
            return this;
        }

        public GetNetworkContainerPlainArgs build() {
            $.containerId = Objects.requireNonNull($.containerId, "expected parameter 'containerId' to be non-null");
            $.projectId = Objects.requireNonNull($.projectId, "expected parameter 'projectId' to be non-null");
            return $;
        }
    }

}
