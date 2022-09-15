// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.mongodbatlas.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class CloudBackupSnapshotExportBucketState extends com.pulumi.resources.ResourceArgs {

    public static final CloudBackupSnapshotExportBucketState Empty = new CloudBackupSnapshotExportBucketState();

    /**
     * Name of the bucket that the provided role ID is authorized to access. You must also specify the `iam_role_id`.
     * 
     */
    @Import(name="bucketName")
    private @Nullable Output<String> bucketName;

    /**
     * @return Name of the bucket that the provided role ID is authorized to access. You must also specify the `iam_role_id`.
     * 
     */
    public Optional<Output<String>> bucketName() {
        return Optional.ofNullable(this.bucketName);
    }

    /**
     * Name of the provider of the cloud service where Atlas can access the S3 bucket. Atlas only supports `AWS`.
     * 
     */
    @Import(name="cloudProvider")
    private @Nullable Output<String> cloudProvider;

    /**
     * @return Name of the provider of the cloud service where Atlas can access the S3 bucket. Atlas only supports `AWS`.
     * 
     */
    public Optional<Output<String>> cloudProvider() {
        return Optional.ofNullable(this.cloudProvider);
    }

    /**
     * Unique identifier of the snapshot export bucket.
     * 
     */
    @Import(name="exportBucketId")
    private @Nullable Output<String> exportBucketId;

    /**
     * @return Unique identifier of the snapshot export bucket.
     * 
     */
    public Optional<Output<String>> exportBucketId() {
        return Optional.ofNullable(this.exportBucketId);
    }

    /**
     * Unique identifier of the role that Atlas can use to access the bucket. You must also specify the `bucket_name`.
     * 
     */
    @Import(name="iamRoleId")
    private @Nullable Output<String> iamRoleId;

    /**
     * @return Unique identifier of the role that Atlas can use to access the bucket. You must also specify the `bucket_name`.
     * 
     */
    public Optional<Output<String>> iamRoleId() {
        return Optional.ofNullable(this.iamRoleId);
    }

    /**
     * The unique identifier of the project for the Atlas cluster.
     * 
     */
    @Import(name="projectId")
    private @Nullable Output<String> projectId;

    /**
     * @return The unique identifier of the project for the Atlas cluster.
     * 
     */
    public Optional<Output<String>> projectId() {
        return Optional.ofNullable(this.projectId);
    }

    private CloudBackupSnapshotExportBucketState() {}

    private CloudBackupSnapshotExportBucketState(CloudBackupSnapshotExportBucketState $) {
        this.bucketName = $.bucketName;
        this.cloudProvider = $.cloudProvider;
        this.exportBucketId = $.exportBucketId;
        this.iamRoleId = $.iamRoleId;
        this.projectId = $.projectId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(CloudBackupSnapshotExportBucketState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private CloudBackupSnapshotExportBucketState $;

        public Builder() {
            $ = new CloudBackupSnapshotExportBucketState();
        }

        public Builder(CloudBackupSnapshotExportBucketState defaults) {
            $ = new CloudBackupSnapshotExportBucketState(Objects.requireNonNull(defaults));
        }

        /**
         * @param bucketName Name of the bucket that the provided role ID is authorized to access. You must also specify the `iam_role_id`.
         * 
         * @return builder
         * 
         */
        public Builder bucketName(@Nullable Output<String> bucketName) {
            $.bucketName = bucketName;
            return this;
        }

        /**
         * @param bucketName Name of the bucket that the provided role ID is authorized to access. You must also specify the `iam_role_id`.
         * 
         * @return builder
         * 
         */
        public Builder bucketName(String bucketName) {
            return bucketName(Output.of(bucketName));
        }

        /**
         * @param cloudProvider Name of the provider of the cloud service where Atlas can access the S3 bucket. Atlas only supports `AWS`.
         * 
         * @return builder
         * 
         */
        public Builder cloudProvider(@Nullable Output<String> cloudProvider) {
            $.cloudProvider = cloudProvider;
            return this;
        }

        /**
         * @param cloudProvider Name of the provider of the cloud service where Atlas can access the S3 bucket. Atlas only supports `AWS`.
         * 
         * @return builder
         * 
         */
        public Builder cloudProvider(String cloudProvider) {
            return cloudProvider(Output.of(cloudProvider));
        }

        /**
         * @param exportBucketId Unique identifier of the snapshot export bucket.
         * 
         * @return builder
         * 
         */
        public Builder exportBucketId(@Nullable Output<String> exportBucketId) {
            $.exportBucketId = exportBucketId;
            return this;
        }

        /**
         * @param exportBucketId Unique identifier of the snapshot export bucket.
         * 
         * @return builder
         * 
         */
        public Builder exportBucketId(String exportBucketId) {
            return exportBucketId(Output.of(exportBucketId));
        }

        /**
         * @param iamRoleId Unique identifier of the role that Atlas can use to access the bucket. You must also specify the `bucket_name`.
         * 
         * @return builder
         * 
         */
        public Builder iamRoleId(@Nullable Output<String> iamRoleId) {
            $.iamRoleId = iamRoleId;
            return this;
        }

        /**
         * @param iamRoleId Unique identifier of the role that Atlas can use to access the bucket. You must also specify the `bucket_name`.
         * 
         * @return builder
         * 
         */
        public Builder iamRoleId(String iamRoleId) {
            return iamRoleId(Output.of(iamRoleId));
        }

        /**
         * @param projectId The unique identifier of the project for the Atlas cluster.
         * 
         * @return builder
         * 
         */
        public Builder projectId(@Nullable Output<String> projectId) {
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

        public CloudBackupSnapshotExportBucketState build() {
            return $;
        }
    }

}