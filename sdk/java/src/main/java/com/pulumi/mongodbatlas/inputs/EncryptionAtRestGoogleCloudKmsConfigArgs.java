// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.mongodbatlas.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class EncryptionAtRestGoogleCloudKmsConfigArgs extends com.pulumi.resources.ResourceArgs {

    public static final EncryptionAtRestGoogleCloudKmsConfigArgs Empty = new EncryptionAtRestGoogleCloudKmsConfigArgs();

    /**
     * Specifies whether Encryption at Rest is enabled for an Atlas project. To disable Encryption at Rest, pass only this parameter with a value of false. When you disable Encryption at Rest, Atlas also removes the configuration details.
     * 
     */
    @Import(name="enabled")
    private @Nullable Output<Boolean> enabled;

    /**
     * @return Specifies whether Encryption at Rest is enabled for an Atlas project. To disable Encryption at Rest, pass only this parameter with a value of false. When you disable Encryption at Rest, Atlas also removes the configuration details.
     * 
     */
    public Optional<Output<Boolean>> enabled() {
        return Optional.ofNullable(this.enabled);
    }

    /**
     * The Key Version Resource ID from your GCP account.
     * 
     */
    @Import(name="keyVersionResourceId")
    private @Nullable Output<String> keyVersionResourceId;

    /**
     * @return The Key Version Resource ID from your GCP account.
     * 
     */
    public Optional<Output<String>> keyVersionResourceId() {
        return Optional.ofNullable(this.keyVersionResourceId);
    }

    /**
     * String-formatted JSON object containing GCP KMS credentials from your GCP account.
     * 
     */
    @Import(name="serviceAccountKey")
    private @Nullable Output<String> serviceAccountKey;

    /**
     * @return String-formatted JSON object containing GCP KMS credentials from your GCP account.
     * 
     */
    public Optional<Output<String>> serviceAccountKey() {
        return Optional.ofNullable(this.serviceAccountKey);
    }

    private EncryptionAtRestGoogleCloudKmsConfigArgs() {}

    private EncryptionAtRestGoogleCloudKmsConfigArgs(EncryptionAtRestGoogleCloudKmsConfigArgs $) {
        this.enabled = $.enabled;
        this.keyVersionResourceId = $.keyVersionResourceId;
        this.serviceAccountKey = $.serviceAccountKey;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(EncryptionAtRestGoogleCloudKmsConfigArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private EncryptionAtRestGoogleCloudKmsConfigArgs $;

        public Builder() {
            $ = new EncryptionAtRestGoogleCloudKmsConfigArgs();
        }

        public Builder(EncryptionAtRestGoogleCloudKmsConfigArgs defaults) {
            $ = new EncryptionAtRestGoogleCloudKmsConfigArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param enabled Specifies whether Encryption at Rest is enabled for an Atlas project. To disable Encryption at Rest, pass only this parameter with a value of false. When you disable Encryption at Rest, Atlas also removes the configuration details.
         * 
         * @return builder
         * 
         */
        public Builder enabled(@Nullable Output<Boolean> enabled) {
            $.enabled = enabled;
            return this;
        }

        /**
         * @param enabled Specifies whether Encryption at Rest is enabled for an Atlas project. To disable Encryption at Rest, pass only this parameter with a value of false. When you disable Encryption at Rest, Atlas also removes the configuration details.
         * 
         * @return builder
         * 
         */
        public Builder enabled(Boolean enabled) {
            return enabled(Output.of(enabled));
        }

        /**
         * @param keyVersionResourceId The Key Version Resource ID from your GCP account.
         * 
         * @return builder
         * 
         */
        public Builder keyVersionResourceId(@Nullable Output<String> keyVersionResourceId) {
            $.keyVersionResourceId = keyVersionResourceId;
            return this;
        }

        /**
         * @param keyVersionResourceId The Key Version Resource ID from your GCP account.
         * 
         * @return builder
         * 
         */
        public Builder keyVersionResourceId(String keyVersionResourceId) {
            return keyVersionResourceId(Output.of(keyVersionResourceId));
        }

        /**
         * @param serviceAccountKey String-formatted JSON object containing GCP KMS credentials from your GCP account.
         * 
         * @return builder
         * 
         */
        public Builder serviceAccountKey(@Nullable Output<String> serviceAccountKey) {
            $.serviceAccountKey = serviceAccountKey;
            return this;
        }

        /**
         * @param serviceAccountKey String-formatted JSON object containing GCP KMS credentials from your GCP account.
         * 
         * @return builder
         * 
         */
        public Builder serviceAccountKey(String serviceAccountKey) {
            return serviceAccountKey(Output.of(serviceAccountKey));
        }

        public EncryptionAtRestGoogleCloudKmsConfigArgs build() {
            return $;
        }
    }

}
