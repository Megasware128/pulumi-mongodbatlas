// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.mongodbatlas.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class GetCloudBackupSnapshotArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetCloudBackupSnapshotArgs Empty = new GetCloudBackupSnapshotArgs();

    /**
     * The name of the Atlas cluster that contains the snapshot you want to retrieve.
     * 
     */
    @Import(name="clusterName", required=true)
    private Output<String> clusterName;

    /**
     * @return The name of the Atlas cluster that contains the snapshot you want to retrieve.
     * 
     */
    public Output<String> clusterName() {
        return this.clusterName;
    }

    @Import(name="projectId", required=true)
    private Output<String> projectId;

    public Output<String> projectId() {
        return this.projectId;
    }

    /**
     * The unique identifier of the snapshot you want to retrieve.
     * 
     */
    @Import(name="snapshotId", required=true)
    private Output<String> snapshotId;

    /**
     * @return The unique identifier of the snapshot you want to retrieve.
     * 
     */
    public Output<String> snapshotId() {
        return this.snapshotId;
    }

    private GetCloudBackupSnapshotArgs() {}

    private GetCloudBackupSnapshotArgs(GetCloudBackupSnapshotArgs $) {
        this.clusterName = $.clusterName;
        this.projectId = $.projectId;
        this.snapshotId = $.snapshotId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetCloudBackupSnapshotArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetCloudBackupSnapshotArgs $;

        public Builder() {
            $ = new GetCloudBackupSnapshotArgs();
        }

        public Builder(GetCloudBackupSnapshotArgs defaults) {
            $ = new GetCloudBackupSnapshotArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param clusterName The name of the Atlas cluster that contains the snapshot you want to retrieve.
         * 
         * @return builder
         * 
         */
        public Builder clusterName(Output<String> clusterName) {
            $.clusterName = clusterName;
            return this;
        }

        /**
         * @param clusterName The name of the Atlas cluster that contains the snapshot you want to retrieve.
         * 
         * @return builder
         * 
         */
        public Builder clusterName(String clusterName) {
            return clusterName(Output.of(clusterName));
        }

        public Builder projectId(Output<String> projectId) {
            $.projectId = projectId;
            return this;
        }

        public Builder projectId(String projectId) {
            return projectId(Output.of(projectId));
        }

        /**
         * @param snapshotId The unique identifier of the snapshot you want to retrieve.
         * 
         * @return builder
         * 
         */
        public Builder snapshotId(Output<String> snapshotId) {
            $.snapshotId = snapshotId;
            return this;
        }

        /**
         * @param snapshotId The unique identifier of the snapshot you want to retrieve.
         * 
         * @return builder
         * 
         */
        public Builder snapshotId(String snapshotId) {
            return snapshotId(Output.of(snapshotId));
        }

        public GetCloudBackupSnapshotArgs build() {
            $.clusterName = Objects.requireNonNull($.clusterName, "expected parameter 'clusterName' to be non-null");
            $.projectId = Objects.requireNonNull($.projectId, "expected parameter 'projectId' to be non-null");
            $.snapshotId = Objects.requireNonNull($.snapshotId, "expected parameter 'snapshotId' to be non-null");
            return $;
        }
    }

}
