// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.mongodbatlas.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class GetCloudProviderSnapshotPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetCloudProviderSnapshotPlainArgs Empty = new GetCloudProviderSnapshotPlainArgs();

    /**
     * The name of the Atlas cluster that contains the snapshot you want to retrieve.
     * 
     */
    @Import(name="clusterName", required=true)
    private String clusterName;

    /**
     * @return The name of the Atlas cluster that contains the snapshot you want to retrieve.
     * 
     */
    public String clusterName() {
        return this.clusterName;
    }

    @Import(name="projectId", required=true)
    private String projectId;

    public String projectId() {
        return this.projectId;
    }

    /**
     * The unique identifier of the snapshot you want to retrieve.
     * 
     */
    @Import(name="snapshotId", required=true)
    private String snapshotId;

    /**
     * @return The unique identifier of the snapshot you want to retrieve.
     * 
     */
    public String snapshotId() {
        return this.snapshotId;
    }

    private GetCloudProviderSnapshotPlainArgs() {}

    private GetCloudProviderSnapshotPlainArgs(GetCloudProviderSnapshotPlainArgs $) {
        this.clusterName = $.clusterName;
        this.projectId = $.projectId;
        this.snapshotId = $.snapshotId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetCloudProviderSnapshotPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetCloudProviderSnapshotPlainArgs $;

        public Builder() {
            $ = new GetCloudProviderSnapshotPlainArgs();
        }

        public Builder(GetCloudProviderSnapshotPlainArgs defaults) {
            $ = new GetCloudProviderSnapshotPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param clusterName The name of the Atlas cluster that contains the snapshot you want to retrieve.
         * 
         * @return builder
         * 
         */
        public Builder clusterName(String clusterName) {
            $.clusterName = clusterName;
            return this;
        }

        public Builder projectId(String projectId) {
            $.projectId = projectId;
            return this;
        }

        /**
         * @param snapshotId The unique identifier of the snapshot you want to retrieve.
         * 
         * @return builder
         * 
         */
        public Builder snapshotId(String snapshotId) {
            $.snapshotId = snapshotId;
            return this;
        }

        public GetCloudProviderSnapshotPlainArgs build() {
            $.clusterName = Objects.requireNonNull($.clusterName, "expected parameter 'clusterName' to be non-null");
            $.projectId = Objects.requireNonNull($.projectId, "expected parameter 'projectId' to be non-null");
            $.snapshotId = Objects.requireNonNull($.snapshotId, "expected parameter 'snapshotId' to be non-null");
            return $;
        }
    }

}
