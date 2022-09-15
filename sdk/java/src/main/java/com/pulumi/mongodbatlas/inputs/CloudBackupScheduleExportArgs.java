// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.mongodbatlas.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class CloudBackupScheduleExportArgs extends com.pulumi.resources.ResourceArgs {

    public static final CloudBackupScheduleExportArgs Empty = new CloudBackupScheduleExportArgs();

    /**
     * Unique identifier of the mongodbatlas.CloudBackupSnapshotExportBucket export_bucket_id value.
     * 
     */
    @Import(name="exportBucketId")
    private @Nullable Output<String> exportBucketId;

    /**
     * @return Unique identifier of the mongodbatlas.CloudBackupSnapshotExportBucket export_bucket_id value.
     * 
     */
    public Optional<Output<String>> exportBucketId() {
        return Optional.ofNullable(this.exportBucketId);
    }

    /**
     * Frequency associated with the export snapshot item.
     * 
     */
    @Import(name="frequencyType")
    private @Nullable Output<String> frequencyType;

    /**
     * @return Frequency associated with the export snapshot item.
     * 
     */
    public Optional<Output<String>> frequencyType() {
        return Optional.ofNullable(this.frequencyType);
    }

    private CloudBackupScheduleExportArgs() {}

    private CloudBackupScheduleExportArgs(CloudBackupScheduleExportArgs $) {
        this.exportBucketId = $.exportBucketId;
        this.frequencyType = $.frequencyType;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(CloudBackupScheduleExportArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private CloudBackupScheduleExportArgs $;

        public Builder() {
            $ = new CloudBackupScheduleExportArgs();
        }

        public Builder(CloudBackupScheduleExportArgs defaults) {
            $ = new CloudBackupScheduleExportArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param exportBucketId Unique identifier of the mongodbatlas.CloudBackupSnapshotExportBucket export_bucket_id value.
         * 
         * @return builder
         * 
         */
        public Builder exportBucketId(@Nullable Output<String> exportBucketId) {
            $.exportBucketId = exportBucketId;
            return this;
        }

        /**
         * @param exportBucketId Unique identifier of the mongodbatlas.CloudBackupSnapshotExportBucket export_bucket_id value.
         * 
         * @return builder
         * 
         */
        public Builder exportBucketId(String exportBucketId) {
            return exportBucketId(Output.of(exportBucketId));
        }

        /**
         * @param frequencyType Frequency associated with the export snapshot item.
         * 
         * @return builder
         * 
         */
        public Builder frequencyType(@Nullable Output<String> frequencyType) {
            $.frequencyType = frequencyType;
            return this;
        }

        /**
         * @param frequencyType Frequency associated with the export snapshot item.
         * 
         * @return builder
         * 
         */
        public Builder frequencyType(String frequencyType) {
            return frequencyType(Output.of(frequencyType));
        }

        public CloudBackupScheduleExportArgs build() {
            return $;
        }
    }

}