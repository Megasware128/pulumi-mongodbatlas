// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.mongodbatlas.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.mongodbatlas.inputs.CloudProviderSnapshotRestoreJobDeliveryTypeConfigArgs;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class CloudProviderSnapshotRestoreJobState extends com.pulumi.resources.ResourceArgs {

    public static final CloudProviderSnapshotRestoreJobState Empty = new CloudProviderSnapshotRestoreJobState();

    /**
     * Indicates whether the restore job was canceled.
     * 
     */
    @Import(name="cancelled")
    private @Nullable Output<Boolean> cancelled;

    /**
     * @return Indicates whether the restore job was canceled.
     * 
     */
    public Optional<Output<Boolean>> cancelled() {
        return Optional.ofNullable(this.cancelled);
    }

    /**
     * The name of the Atlas cluster whose snapshot you want to restore.
     * 
     */
    @Import(name="clusterName")
    private @Nullable Output<String> clusterName;

    /**
     * @return The name of the Atlas cluster whose snapshot you want to restore.
     * 
     */
    public Optional<Output<String>> clusterName() {
        return Optional.ofNullable(this.clusterName);
    }

    /**
     * UTC ISO 8601 formatted point in time when Atlas created the restore job.
     * 
     */
    @Import(name="createdAt")
    private @Nullable Output<String> createdAt;

    /**
     * @return UTC ISO 8601 formatted point in time when Atlas created the restore job.
     * 
     */
    public Optional<Output<String>> createdAt() {
        return Optional.ofNullable(this.createdAt);
    }

    /**
     * Type of restore job to create. Possible values are: **download** or **automated**, only one must be set it in ``true``.
     * 
     * @deprecated
     * use delivery_type_config instead
     * 
     */
    @Deprecated /* use delivery_type_config instead */
    @Import(name="deliveryType")
    private @Nullable Output<Map<String,String>> deliveryType;

    /**
     * @return Type of restore job to create. Possible values are: **download** or **automated**, only one must be set it in ``true``.
     * 
     * @deprecated
     * use delivery_type_config instead
     * 
     */
    @Deprecated /* use delivery_type_config instead */
    public Optional<Output<Map<String,String>>> deliveryType() {
        return Optional.ofNullable(this.deliveryType);
    }

    /**
     * Type of restore job to create. Possible values are: automated and download.
     * 
     */
    @Import(name="deliveryTypeConfig")
    private @Nullable Output<CloudProviderSnapshotRestoreJobDeliveryTypeConfigArgs> deliveryTypeConfig;

    /**
     * @return Type of restore job to create. Possible values are: automated and download.
     * 
     */
    public Optional<Output<CloudProviderSnapshotRestoreJobDeliveryTypeConfigArgs>> deliveryTypeConfig() {
        return Optional.ofNullable(this.deliveryTypeConfig);
    }

    /**
     * One or more URLs for the compressed snapshot files for manual download. Only visible if deliveryType is download.
     * 
     */
    @Import(name="deliveryUrls")
    private @Nullable Output<List<String>> deliveryUrls;

    /**
     * @return One or more URLs for the compressed snapshot files for manual download. Only visible if deliveryType is download.
     * 
     */
    public Optional<Output<List<String>>> deliveryUrls() {
        return Optional.ofNullable(this.deliveryUrls);
    }

    /**
     * Indicates whether the restore job expired.
     * 
     */
    @Import(name="expired")
    private @Nullable Output<Boolean> expired;

    /**
     * @return Indicates whether the restore job expired.
     * 
     */
    public Optional<Output<Boolean>> expired() {
        return Optional.ofNullable(this.expired);
    }

    /**
     * UTC ISO 8601 formatted point in time when the restore job expires.
     * 
     */
    @Import(name="expiresAt")
    private @Nullable Output<String> expiresAt;

    /**
     * @return UTC ISO 8601 formatted point in time when the restore job expires.
     * 
     */
    public Optional<Output<String>> expiresAt() {
        return Optional.ofNullable(this.expiresAt);
    }

    /**
     * UTC ISO 8601 formatted point in time when the restore job completed.
     * 
     */
    @Import(name="finishedAt")
    private @Nullable Output<String> finishedAt;

    /**
     * @return UTC ISO 8601 formatted point in time when the restore job completed.
     * 
     */
    public Optional<Output<String>> finishedAt() {
        return Optional.ofNullable(this.finishedAt);
    }

    /**
     * The unique identifier of the project for the Atlas cluster whose snapshot you want to restore.
     * 
     */
    @Import(name="projectId")
    private @Nullable Output<String> projectId;

    /**
     * @return The unique identifier of the project for the Atlas cluster whose snapshot you want to restore.
     * 
     */
    public Optional<Output<String>> projectId() {
        return Optional.ofNullable(this.projectId);
    }

    /**
     * Unique identifier of the snapshot to restore.
     * 
     */
    @Import(name="snapshotId")
    private @Nullable Output<String> snapshotId;

    /**
     * @return Unique identifier of the snapshot to restore.
     * 
     */
    public Optional<Output<String>> snapshotId() {
        return Optional.ofNullable(this.snapshotId);
    }

    /**
     * The unique identifier of the restore job.
     * 
     */
    @Import(name="snapshotRestoreJobId")
    private @Nullable Output<String> snapshotRestoreJobId;

    /**
     * @return The unique identifier of the restore job.
     * 
     */
    public Optional<Output<String>> snapshotRestoreJobId() {
        return Optional.ofNullable(this.snapshotRestoreJobId);
    }

    /**
     * Timestamp in ISO 8601 date and time format in UTC when the snapshot associated to snapshotId was taken.
     * 
     */
    @Import(name="timestamp")
    private @Nullable Output<String> timestamp;

    /**
     * @return Timestamp in ISO 8601 date and time format in UTC when the snapshot associated to snapshotId was taken.
     * 
     */
    public Optional<Output<String>> timestamp() {
        return Optional.ofNullable(this.timestamp);
    }

    private CloudProviderSnapshotRestoreJobState() {}

    private CloudProviderSnapshotRestoreJobState(CloudProviderSnapshotRestoreJobState $) {
        this.cancelled = $.cancelled;
        this.clusterName = $.clusterName;
        this.createdAt = $.createdAt;
        this.deliveryType = $.deliveryType;
        this.deliveryTypeConfig = $.deliveryTypeConfig;
        this.deliveryUrls = $.deliveryUrls;
        this.expired = $.expired;
        this.expiresAt = $.expiresAt;
        this.finishedAt = $.finishedAt;
        this.projectId = $.projectId;
        this.snapshotId = $.snapshotId;
        this.snapshotRestoreJobId = $.snapshotRestoreJobId;
        this.timestamp = $.timestamp;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(CloudProviderSnapshotRestoreJobState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private CloudProviderSnapshotRestoreJobState $;

        public Builder() {
            $ = new CloudProviderSnapshotRestoreJobState();
        }

        public Builder(CloudProviderSnapshotRestoreJobState defaults) {
            $ = new CloudProviderSnapshotRestoreJobState(Objects.requireNonNull(defaults));
        }

        /**
         * @param cancelled Indicates whether the restore job was canceled.
         * 
         * @return builder
         * 
         */
        public Builder cancelled(@Nullable Output<Boolean> cancelled) {
            $.cancelled = cancelled;
            return this;
        }

        /**
         * @param cancelled Indicates whether the restore job was canceled.
         * 
         * @return builder
         * 
         */
        public Builder cancelled(Boolean cancelled) {
            return cancelled(Output.of(cancelled));
        }

        /**
         * @param clusterName The name of the Atlas cluster whose snapshot you want to restore.
         * 
         * @return builder
         * 
         */
        public Builder clusterName(@Nullable Output<String> clusterName) {
            $.clusterName = clusterName;
            return this;
        }

        /**
         * @param clusterName The name of the Atlas cluster whose snapshot you want to restore.
         * 
         * @return builder
         * 
         */
        public Builder clusterName(String clusterName) {
            return clusterName(Output.of(clusterName));
        }

        /**
         * @param createdAt UTC ISO 8601 formatted point in time when Atlas created the restore job.
         * 
         * @return builder
         * 
         */
        public Builder createdAt(@Nullable Output<String> createdAt) {
            $.createdAt = createdAt;
            return this;
        }

        /**
         * @param createdAt UTC ISO 8601 formatted point in time when Atlas created the restore job.
         * 
         * @return builder
         * 
         */
        public Builder createdAt(String createdAt) {
            return createdAt(Output.of(createdAt));
        }

        /**
         * @param deliveryType Type of restore job to create. Possible values are: **download** or **automated**, only one must be set it in ``true``.
         * 
         * @return builder
         * 
         * @deprecated
         * use delivery_type_config instead
         * 
         */
        @Deprecated /* use delivery_type_config instead */
        public Builder deliveryType(@Nullable Output<Map<String,String>> deliveryType) {
            $.deliveryType = deliveryType;
            return this;
        }

        /**
         * @param deliveryType Type of restore job to create. Possible values are: **download** or **automated**, only one must be set it in ``true``.
         * 
         * @return builder
         * 
         * @deprecated
         * use delivery_type_config instead
         * 
         */
        @Deprecated /* use delivery_type_config instead */
        public Builder deliveryType(Map<String,String> deliveryType) {
            return deliveryType(Output.of(deliveryType));
        }

        /**
         * @param deliveryTypeConfig Type of restore job to create. Possible values are: automated and download.
         * 
         * @return builder
         * 
         */
        public Builder deliveryTypeConfig(@Nullable Output<CloudProviderSnapshotRestoreJobDeliveryTypeConfigArgs> deliveryTypeConfig) {
            $.deliveryTypeConfig = deliveryTypeConfig;
            return this;
        }

        /**
         * @param deliveryTypeConfig Type of restore job to create. Possible values are: automated and download.
         * 
         * @return builder
         * 
         */
        public Builder deliveryTypeConfig(CloudProviderSnapshotRestoreJobDeliveryTypeConfigArgs deliveryTypeConfig) {
            return deliveryTypeConfig(Output.of(deliveryTypeConfig));
        }

        /**
         * @param deliveryUrls One or more URLs for the compressed snapshot files for manual download. Only visible if deliveryType is download.
         * 
         * @return builder
         * 
         */
        public Builder deliveryUrls(@Nullable Output<List<String>> deliveryUrls) {
            $.deliveryUrls = deliveryUrls;
            return this;
        }

        /**
         * @param deliveryUrls One or more URLs for the compressed snapshot files for manual download. Only visible if deliveryType is download.
         * 
         * @return builder
         * 
         */
        public Builder deliveryUrls(List<String> deliveryUrls) {
            return deliveryUrls(Output.of(deliveryUrls));
        }

        /**
         * @param deliveryUrls One or more URLs for the compressed snapshot files for manual download. Only visible if deliveryType is download.
         * 
         * @return builder
         * 
         */
        public Builder deliveryUrls(String... deliveryUrls) {
            return deliveryUrls(List.of(deliveryUrls));
        }

        /**
         * @param expired Indicates whether the restore job expired.
         * 
         * @return builder
         * 
         */
        public Builder expired(@Nullable Output<Boolean> expired) {
            $.expired = expired;
            return this;
        }

        /**
         * @param expired Indicates whether the restore job expired.
         * 
         * @return builder
         * 
         */
        public Builder expired(Boolean expired) {
            return expired(Output.of(expired));
        }

        /**
         * @param expiresAt UTC ISO 8601 formatted point in time when the restore job expires.
         * 
         * @return builder
         * 
         */
        public Builder expiresAt(@Nullable Output<String> expiresAt) {
            $.expiresAt = expiresAt;
            return this;
        }

        /**
         * @param expiresAt UTC ISO 8601 formatted point in time when the restore job expires.
         * 
         * @return builder
         * 
         */
        public Builder expiresAt(String expiresAt) {
            return expiresAt(Output.of(expiresAt));
        }

        /**
         * @param finishedAt UTC ISO 8601 formatted point in time when the restore job completed.
         * 
         * @return builder
         * 
         */
        public Builder finishedAt(@Nullable Output<String> finishedAt) {
            $.finishedAt = finishedAt;
            return this;
        }

        /**
         * @param finishedAt UTC ISO 8601 formatted point in time when the restore job completed.
         * 
         * @return builder
         * 
         */
        public Builder finishedAt(String finishedAt) {
            return finishedAt(Output.of(finishedAt));
        }

        /**
         * @param projectId The unique identifier of the project for the Atlas cluster whose snapshot you want to restore.
         * 
         * @return builder
         * 
         */
        public Builder projectId(@Nullable Output<String> projectId) {
            $.projectId = projectId;
            return this;
        }

        /**
         * @param projectId The unique identifier of the project for the Atlas cluster whose snapshot you want to restore.
         * 
         * @return builder
         * 
         */
        public Builder projectId(String projectId) {
            return projectId(Output.of(projectId));
        }

        /**
         * @param snapshotId Unique identifier of the snapshot to restore.
         * 
         * @return builder
         * 
         */
        public Builder snapshotId(@Nullable Output<String> snapshotId) {
            $.snapshotId = snapshotId;
            return this;
        }

        /**
         * @param snapshotId Unique identifier of the snapshot to restore.
         * 
         * @return builder
         * 
         */
        public Builder snapshotId(String snapshotId) {
            return snapshotId(Output.of(snapshotId));
        }

        /**
         * @param snapshotRestoreJobId The unique identifier of the restore job.
         * 
         * @return builder
         * 
         */
        public Builder snapshotRestoreJobId(@Nullable Output<String> snapshotRestoreJobId) {
            $.snapshotRestoreJobId = snapshotRestoreJobId;
            return this;
        }

        /**
         * @param snapshotRestoreJobId The unique identifier of the restore job.
         * 
         * @return builder
         * 
         */
        public Builder snapshotRestoreJobId(String snapshotRestoreJobId) {
            return snapshotRestoreJobId(Output.of(snapshotRestoreJobId));
        }

        /**
         * @param timestamp Timestamp in ISO 8601 date and time format in UTC when the snapshot associated to snapshotId was taken.
         * 
         * @return builder
         * 
         */
        public Builder timestamp(@Nullable Output<String> timestamp) {
            $.timestamp = timestamp;
            return this;
        }

        /**
         * @param timestamp Timestamp in ISO 8601 date and time format in UTC when the snapshot associated to snapshotId was taken.
         * 
         * @return builder
         * 
         */
        public Builder timestamp(String timestamp) {
            return timestamp(Output.of(timestamp));
        }

        public CloudProviderSnapshotRestoreJobState build() {
            return $;
        }
    }

}
