// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.mongodbatlas;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;


public final class CloudBackupSnapshotArgs extends com.pulumi.resources.ResourceArgs {

    public static final CloudBackupSnapshotArgs Empty = new CloudBackupSnapshotArgs();

    /**
     * The name of the Atlas cluster that contains the snapshots you want to retrieve.
     * 
     */
    @Import(name="clusterName", required=true)
    private Output<String> clusterName;

    /**
     * @return The name of the Atlas cluster that contains the snapshots you want to retrieve.
     * 
     */
    public Output<String> clusterName() {
        return this.clusterName;
    }

    /**
     * Description of the on-demand snapshot.
     * 
     */
    @Import(name="description", required=true)
    private Output<String> description;

    /**
     * @return Description of the on-demand snapshot.
     * 
     */
    public Output<String> description() {
        return this.description;
    }

    /**
     * The unique identifier of the project for the Atlas cluster.
     * 
     */
    @Import(name="projectId", required=true)
    private Output<String> projectId;

    /**
     * @return The unique identifier of the project for the Atlas cluster.
     * 
     */
    public Output<String> projectId() {
        return this.projectId;
    }

    /**
     * The number of days that Atlas should retain the on-demand snapshot. Must be at least 1.
     * 
     */
    @Import(name="retentionInDays", required=true)
    private Output<Integer> retentionInDays;

    /**
     * @return The number of days that Atlas should retain the on-demand snapshot. Must be at least 1.
     * 
     */
    public Output<Integer> retentionInDays() {
        return this.retentionInDays;
    }

    private CloudBackupSnapshotArgs() {}

    private CloudBackupSnapshotArgs(CloudBackupSnapshotArgs $) {
        this.clusterName = $.clusterName;
        this.description = $.description;
        this.projectId = $.projectId;
        this.retentionInDays = $.retentionInDays;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(CloudBackupSnapshotArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private CloudBackupSnapshotArgs $;

        public Builder() {
            $ = new CloudBackupSnapshotArgs();
        }

        public Builder(CloudBackupSnapshotArgs defaults) {
            $ = new CloudBackupSnapshotArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param clusterName The name of the Atlas cluster that contains the snapshots you want to retrieve.
         * 
         * @return builder
         * 
         */
        public Builder clusterName(Output<String> clusterName) {
            $.clusterName = clusterName;
            return this;
        }

        /**
         * @param clusterName The name of the Atlas cluster that contains the snapshots you want to retrieve.
         * 
         * @return builder
         * 
         */
        public Builder clusterName(String clusterName) {
            return clusterName(Output.of(clusterName));
        }

        /**
         * @param description Description of the on-demand snapshot.
         * 
         * @return builder
         * 
         */
        public Builder description(Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description Description of the on-demand snapshot.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param projectId The unique identifier of the project for the Atlas cluster.
         * 
         * @return builder
         * 
         */
        public Builder projectId(Output<String> projectId) {
            $.projectId = projectId;
            return this;
        }

        /**
         * @param projectId The unique identifier of the project for the Atlas cluster.
         * 
         * @return builder
         * 
         */
        public Builder projectId(String projectId) {
            return projectId(Output.of(projectId));
        }

        /**
         * @param retentionInDays The number of days that Atlas should retain the on-demand snapshot. Must be at least 1.
         * 
         * @return builder
         * 
         */
        public Builder retentionInDays(Output<Integer> retentionInDays) {
            $.retentionInDays = retentionInDays;
            return this;
        }

        /**
         * @param retentionInDays The number of days that Atlas should retain the on-demand snapshot. Must be at least 1.
         * 
         * @return builder
         * 
         */
        public Builder retentionInDays(Integer retentionInDays) {
            return retentionInDays(Output.of(retentionInDays));
        }

        public CloudBackupSnapshotArgs build() {
            $.clusterName = Objects.requireNonNull($.clusterName, "expected parameter 'clusterName' to be non-null");
            $.description = Objects.requireNonNull($.description, "expected parameter 'description' to be non-null");
            $.projectId = Objects.requireNonNull($.projectId, "expected parameter 'projectId' to be non-null");
            $.retentionInDays = Objects.requireNonNull($.retentionInDays, "expected parameter 'retentionInDays' to be non-null");
            return $;
        }
    }

}
