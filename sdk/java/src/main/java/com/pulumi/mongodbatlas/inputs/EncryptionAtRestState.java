// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.mongodbatlas.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.mongodbatlas.inputs.EncryptionAtRestAwsKmsConfigArgs;
import com.pulumi.mongodbatlas.inputs.EncryptionAtRestAzureKeyVaultConfigArgs;
import com.pulumi.mongodbatlas.inputs.EncryptionAtRestGoogleCloudKmsConfigArgs;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class EncryptionAtRestState extends com.pulumi.resources.ResourceArgs {

    public static final EncryptionAtRestState Empty = new EncryptionAtRestState();

    /**
     * Specifies AWS KMS configuration details and whether Encryption at Rest is enabled for an Atlas project.
     * 
     * @deprecated
     * use aws_kms_config instead
     * 
     */
    @Deprecated /* use aws_kms_config instead */
    @Import(name="awsKms")
    private @Nullable Output<Map<String,String>> awsKms;

    /**
     * @return Specifies AWS KMS configuration details and whether Encryption at Rest is enabled for an Atlas project.
     * 
     * @deprecated
     * use aws_kms_config instead
     * 
     */
    @Deprecated /* use aws_kms_config instead */
    public Optional<Output<Map<String,String>>> awsKms() {
        return Optional.ofNullable(this.awsKms);
    }

    @Import(name="awsKmsConfig")
    private @Nullable Output<EncryptionAtRestAwsKmsConfigArgs> awsKmsConfig;

    public Optional<Output<EncryptionAtRestAwsKmsConfigArgs>> awsKmsConfig() {
        return Optional.ofNullable(this.awsKmsConfig);
    }

    /**
     * Specifies Azure Key Vault configuration details and whether Encryption at Rest is enabled for an Atlas project.
     * 
     * @deprecated
     * use azure_key_vault_config instead
     * 
     */
    @Deprecated /* use azure_key_vault_config instead */
    @Import(name="azureKeyVault")
    private @Nullable Output<Map<String,String>> azureKeyVault;

    /**
     * @return Specifies Azure Key Vault configuration details and whether Encryption at Rest is enabled for an Atlas project.
     * 
     * @deprecated
     * use azure_key_vault_config instead
     * 
     */
    @Deprecated /* use azure_key_vault_config instead */
    public Optional<Output<Map<String,String>>> azureKeyVault() {
        return Optional.ofNullable(this.azureKeyVault);
    }

    @Import(name="azureKeyVaultConfig")
    private @Nullable Output<EncryptionAtRestAzureKeyVaultConfigArgs> azureKeyVaultConfig;

    public Optional<Output<EncryptionAtRestAzureKeyVaultConfigArgs>> azureKeyVaultConfig() {
        return Optional.ofNullable(this.azureKeyVaultConfig);
    }

    /**
     * Specifies GCP KMS configuration details and whether Encryption at Rest is enabled for an Atlas project.
     * 
     * @deprecated
     * use google_cloud_kms_config instead
     * 
     */
    @Deprecated /* use google_cloud_kms_config instead */
    @Import(name="googleCloudKms")
    private @Nullable Output<Map<String,String>> googleCloudKms;

    /**
     * @return Specifies GCP KMS configuration details and whether Encryption at Rest is enabled for an Atlas project.
     * 
     * @deprecated
     * use google_cloud_kms_config instead
     * 
     */
    @Deprecated /* use google_cloud_kms_config instead */
    public Optional<Output<Map<String,String>>> googleCloudKms() {
        return Optional.ofNullable(this.googleCloudKms);
    }

    @Import(name="googleCloudKmsConfig")
    private @Nullable Output<EncryptionAtRestGoogleCloudKmsConfigArgs> googleCloudKmsConfig;

    public Optional<Output<EncryptionAtRestGoogleCloudKmsConfigArgs>> googleCloudKmsConfig() {
        return Optional.ofNullable(this.googleCloudKmsConfig);
    }

    /**
     * The unique identifier for the project.
     * 
     */
    @Import(name="projectId")
    private @Nullable Output<String> projectId;

    /**
     * @return The unique identifier for the project.
     * 
     */
    public Optional<Output<String>> projectId() {
        return Optional.ofNullable(this.projectId);
    }

    private EncryptionAtRestState() {}

    private EncryptionAtRestState(EncryptionAtRestState $) {
        this.awsKms = $.awsKms;
        this.awsKmsConfig = $.awsKmsConfig;
        this.azureKeyVault = $.azureKeyVault;
        this.azureKeyVaultConfig = $.azureKeyVaultConfig;
        this.googleCloudKms = $.googleCloudKms;
        this.googleCloudKmsConfig = $.googleCloudKmsConfig;
        this.projectId = $.projectId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(EncryptionAtRestState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private EncryptionAtRestState $;

        public Builder() {
            $ = new EncryptionAtRestState();
        }

        public Builder(EncryptionAtRestState defaults) {
            $ = new EncryptionAtRestState(Objects.requireNonNull(defaults));
        }

        /**
         * @param awsKms Specifies AWS KMS configuration details and whether Encryption at Rest is enabled for an Atlas project.
         * 
         * @return builder
         * 
         * @deprecated
         * use aws_kms_config instead
         * 
         */
        @Deprecated /* use aws_kms_config instead */
        public Builder awsKms(@Nullable Output<Map<String,String>> awsKms) {
            $.awsKms = awsKms;
            return this;
        }

        /**
         * @param awsKms Specifies AWS KMS configuration details and whether Encryption at Rest is enabled for an Atlas project.
         * 
         * @return builder
         * 
         * @deprecated
         * use aws_kms_config instead
         * 
         */
        @Deprecated /* use aws_kms_config instead */
        public Builder awsKms(Map<String,String> awsKms) {
            return awsKms(Output.of(awsKms));
        }

        public Builder awsKmsConfig(@Nullable Output<EncryptionAtRestAwsKmsConfigArgs> awsKmsConfig) {
            $.awsKmsConfig = awsKmsConfig;
            return this;
        }

        public Builder awsKmsConfig(EncryptionAtRestAwsKmsConfigArgs awsKmsConfig) {
            return awsKmsConfig(Output.of(awsKmsConfig));
        }

        /**
         * @param azureKeyVault Specifies Azure Key Vault configuration details and whether Encryption at Rest is enabled for an Atlas project.
         * 
         * @return builder
         * 
         * @deprecated
         * use azure_key_vault_config instead
         * 
         */
        @Deprecated /* use azure_key_vault_config instead */
        public Builder azureKeyVault(@Nullable Output<Map<String,String>> azureKeyVault) {
            $.azureKeyVault = azureKeyVault;
            return this;
        }

        /**
         * @param azureKeyVault Specifies Azure Key Vault configuration details and whether Encryption at Rest is enabled for an Atlas project.
         * 
         * @return builder
         * 
         * @deprecated
         * use azure_key_vault_config instead
         * 
         */
        @Deprecated /* use azure_key_vault_config instead */
        public Builder azureKeyVault(Map<String,String> azureKeyVault) {
            return azureKeyVault(Output.of(azureKeyVault));
        }

        public Builder azureKeyVaultConfig(@Nullable Output<EncryptionAtRestAzureKeyVaultConfigArgs> azureKeyVaultConfig) {
            $.azureKeyVaultConfig = azureKeyVaultConfig;
            return this;
        }

        public Builder azureKeyVaultConfig(EncryptionAtRestAzureKeyVaultConfigArgs azureKeyVaultConfig) {
            return azureKeyVaultConfig(Output.of(azureKeyVaultConfig));
        }

        /**
         * @param googleCloudKms Specifies GCP KMS configuration details and whether Encryption at Rest is enabled for an Atlas project.
         * 
         * @return builder
         * 
         * @deprecated
         * use google_cloud_kms_config instead
         * 
         */
        @Deprecated /* use google_cloud_kms_config instead */
        public Builder googleCloudKms(@Nullable Output<Map<String,String>> googleCloudKms) {
            $.googleCloudKms = googleCloudKms;
            return this;
        }

        /**
         * @param googleCloudKms Specifies GCP KMS configuration details and whether Encryption at Rest is enabled for an Atlas project.
         * 
         * @return builder
         * 
         * @deprecated
         * use google_cloud_kms_config instead
         * 
         */
        @Deprecated /* use google_cloud_kms_config instead */
        public Builder googleCloudKms(Map<String,String> googleCloudKms) {
            return googleCloudKms(Output.of(googleCloudKms));
        }

        public Builder googleCloudKmsConfig(@Nullable Output<EncryptionAtRestGoogleCloudKmsConfigArgs> googleCloudKmsConfig) {
            $.googleCloudKmsConfig = googleCloudKmsConfig;
            return this;
        }

        public Builder googleCloudKmsConfig(EncryptionAtRestGoogleCloudKmsConfigArgs googleCloudKmsConfig) {
            return googleCloudKmsConfig(Output.of(googleCloudKmsConfig));
        }

        /**
         * @param projectId The unique identifier for the project.
         * 
         * @return builder
         * 
         */
        public Builder projectId(@Nullable Output<String> projectId) {
            $.projectId = projectId;
            return this;
        }

        /**
         * @param projectId The unique identifier for the project.
         * 
         * @return builder
         * 
         */
        public Builder projectId(String projectId) {
            return projectId(Output.of(projectId));
        }

        public EncryptionAtRestState build() {
            return $;
        }
    }

}
