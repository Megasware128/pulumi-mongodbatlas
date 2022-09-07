// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.mongodbatlas.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class GetCloudBackupSnapshotExportJobPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetCloudBackupSnapshotExportJobPlainArgs Empty = new GetCloudBackupSnapshotExportJobPlainArgs();

    /**
     * Name of the Atlas cluster whose export job you want to retrieve.
     * 
     */
    @Import(name="clusterName", required=true)
    private String clusterName;

    /**
     * @return Name of the Atlas cluster whose export job you want to retrieve.
     * 
     */
    public String clusterName() {
        return this.clusterName;
    }

    /**
     * -(Required) Unique identifier of the export job to retrieve.
     * 
     */
    @Import(name="exportJobId", required=true)
    private String exportJobId;

    /**
     * @return -(Required) Unique identifier of the export job to retrieve.
     * 
     */
    public String exportJobId() {
        return this.exportJobId;
    }

    @Import(name="id", required=true)
    private String id;

    public String id() {
        return this.id;
    }

    /**
     * Unique 24-hexadecimal digit string that identifies the project which contains the Atlas cluster whose snapshot you want to retrieve.
     * 
     */
    @Import(name="projectId", required=true)
    private String projectId;

    /**
     * @return Unique 24-hexadecimal digit string that identifies the project which contains the Atlas cluster whose snapshot you want to retrieve.
     * 
     */
    public String projectId() {
        return this.projectId;
    }

    private GetCloudBackupSnapshotExportJobPlainArgs() {}

    private GetCloudBackupSnapshotExportJobPlainArgs(GetCloudBackupSnapshotExportJobPlainArgs $) {
        this.clusterName = $.clusterName;
        this.exportJobId = $.exportJobId;
        this.id = $.id;
        this.projectId = $.projectId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetCloudBackupSnapshotExportJobPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetCloudBackupSnapshotExportJobPlainArgs $;

        public Builder() {
            $ = new GetCloudBackupSnapshotExportJobPlainArgs();
        }

        public Builder(GetCloudBackupSnapshotExportJobPlainArgs defaults) {
            $ = new GetCloudBackupSnapshotExportJobPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param clusterName Name of the Atlas cluster whose export job you want to retrieve.
         * 
         * @return builder
         * 
         */
        public Builder clusterName(String clusterName) {
            $.clusterName = clusterName;
            return this;
        }

        /**
         * @param exportJobId -(Required) Unique identifier of the export job to retrieve.
         * 
         * @return builder
         * 
         */
        public Builder exportJobId(String exportJobId) {
            $.exportJobId = exportJobId;
            return this;
        }

        public Builder id(String id) {
            $.id = id;
            return this;
        }

        /**
         * @param projectId Unique 24-hexadecimal digit string that identifies the project which contains the Atlas cluster whose snapshot you want to retrieve.
         * 
         * @return builder
         * 
         */
        public Builder projectId(String projectId) {
            $.projectId = projectId;
            return this;
        }

        public GetCloudBackupSnapshotExportJobPlainArgs build() {
            $.clusterName = Objects.requireNonNull($.clusterName, "expected parameter 'clusterName' to be non-null");
            $.exportJobId = Objects.requireNonNull($.exportJobId, "expected parameter 'exportJobId' to be non-null");
            $.id = Objects.requireNonNull($.id, "expected parameter 'id' to be non-null");
            $.projectId = Objects.requireNonNull($.projectId, "expected parameter 'projectId' to be non-null");
            return $;
        }
    }

}
