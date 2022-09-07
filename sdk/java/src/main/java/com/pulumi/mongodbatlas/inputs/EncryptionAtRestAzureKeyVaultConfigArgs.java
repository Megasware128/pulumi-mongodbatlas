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


public final class EncryptionAtRestAzureKeyVaultConfigArgs extends com.pulumi.resources.ResourceArgs {

    public static final EncryptionAtRestAzureKeyVaultConfigArgs Empty = new EncryptionAtRestAzureKeyVaultConfigArgs();

    /**
     * The Azure environment where the Azure account credentials reside. Valid values are the following: AZURE, AZURE_CHINA, AZURE_GERMANY
     * 
     */
    @Import(name="azureEnvironment")
    private @Nullable Output<String> azureEnvironment;

    /**
     * @return The Azure environment where the Azure account credentials reside. Valid values are the following: AZURE, AZURE_CHINA, AZURE_GERMANY
     * 
     */
    public Optional<Output<String>> azureEnvironment() {
        return Optional.ofNullable(this.azureEnvironment);
    }

    /**
     * The client ID, also known as the application ID, for an Azure application associated with the Azure AD tenant.
     * 
     */
    @Import(name="clientId")
    private @Nullable Output<String> clientId;

    /**
     * @return The client ID, also known as the application ID, for an Azure application associated with the Azure AD tenant.
     * 
     */
    public Optional<Output<String>> clientId() {
        return Optional.ofNullable(this.clientId);
    }

    /**
     * Specifies whether Encryption at Rest is enabled for an Atlas project. To disable Encryption at Rest, pass only this parameter with a value of false. When you disable Encryption at Rest, Atlas also removes the configuration details.
     * 
     */
    @Import(name="enabled", required=true)
    private Output<Boolean> enabled;

    /**
     * @return Specifies whether Encryption at Rest is enabled for an Atlas project. To disable Encryption at Rest, pass only this parameter with a value of false. When you disable Encryption at Rest, Atlas also removes the configuration details.
     * 
     */
    public Output<Boolean> enabled() {
        return this.enabled;
    }

    /**
     * The unique identifier of a key in an Azure Key Vault.
     * 
     */
    @Import(name="keyIdentifier")
    private @Nullable Output<String> keyIdentifier;

    /**
     * @return The unique identifier of a key in an Azure Key Vault.
     * 
     */
    public Optional<Output<String>> keyIdentifier() {
        return Optional.ofNullable(this.keyIdentifier);
    }

    /**
     * The name of an Azure Key Vault containing your key.
     * 
     */
    @Import(name="keyVaultName")
    private @Nullable Output<String> keyVaultName;

    /**
     * @return The name of an Azure Key Vault containing your key.
     * 
     */
    public Optional<Output<String>> keyVaultName() {
        return Optional.ofNullable(this.keyVaultName);
    }

    /**
     * The name of the Azure Resource group that contains an Azure Key Vault.
     * 
     */
    @Import(name="resourceGroupName")
    private @Nullable Output<String> resourceGroupName;

    /**
     * @return The name of the Azure Resource group that contains an Azure Key Vault.
     * 
     */
    public Optional<Output<String>> resourceGroupName() {
        return Optional.ofNullable(this.resourceGroupName);
    }

    /**
     * The secret associated with the Azure Key Vault specified by azureKeyVault.tenantID.
     * 
     */
    @Import(name="secret")
    private @Nullable Output<String> secret;

    /**
     * @return The secret associated with the Azure Key Vault specified by azureKeyVault.tenantID.
     * 
     */
    public Optional<Output<String>> secret() {
        return Optional.ofNullable(this.secret);
    }

    /**
     * The unique identifier associated with an Azure subscription.
     * 
     */
    @Import(name="subscriptionId")
    private @Nullable Output<String> subscriptionId;

    /**
     * @return The unique identifier associated with an Azure subscription.
     * 
     */
    public Optional<Output<String>> subscriptionId() {
        return Optional.ofNullable(this.subscriptionId);
    }

    /**
     * The unique identifier for an Azure AD tenant within an Azure subscription.
     * 
     */
    @Import(name="tenantId")
    private @Nullable Output<String> tenantId;

    /**
     * @return The unique identifier for an Azure AD tenant within an Azure subscription.
     * 
     */
    public Optional<Output<String>> tenantId() {
        return Optional.ofNullable(this.tenantId);
    }

    private EncryptionAtRestAzureKeyVaultConfigArgs() {}

    private EncryptionAtRestAzureKeyVaultConfigArgs(EncryptionAtRestAzureKeyVaultConfigArgs $) {
        this.azureEnvironment = $.azureEnvironment;
        this.clientId = $.clientId;
        this.enabled = $.enabled;
        this.keyIdentifier = $.keyIdentifier;
        this.keyVaultName = $.keyVaultName;
        this.resourceGroupName = $.resourceGroupName;
        this.secret = $.secret;
        this.subscriptionId = $.subscriptionId;
        this.tenantId = $.tenantId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(EncryptionAtRestAzureKeyVaultConfigArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private EncryptionAtRestAzureKeyVaultConfigArgs $;

        public Builder() {
            $ = new EncryptionAtRestAzureKeyVaultConfigArgs();
        }

        public Builder(EncryptionAtRestAzureKeyVaultConfigArgs defaults) {
            $ = new EncryptionAtRestAzureKeyVaultConfigArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param azureEnvironment The Azure environment where the Azure account credentials reside. Valid values are the following: AZURE, AZURE_CHINA, AZURE_GERMANY
         * 
         * @return builder
         * 
         */
        public Builder azureEnvironment(@Nullable Output<String> azureEnvironment) {
            $.azureEnvironment = azureEnvironment;
            return this;
        }

        /**
         * @param azureEnvironment The Azure environment where the Azure account credentials reside. Valid values are the following: AZURE, AZURE_CHINA, AZURE_GERMANY
         * 
         * @return builder
         * 
         */
        public Builder azureEnvironment(String azureEnvironment) {
            return azureEnvironment(Output.of(azureEnvironment));
        }

        /**
         * @param clientId The client ID, also known as the application ID, for an Azure application associated with the Azure AD tenant.
         * 
         * @return builder
         * 
         */
        public Builder clientId(@Nullable Output<String> clientId) {
            $.clientId = clientId;
            return this;
        }

        /**
         * @param clientId The client ID, also known as the application ID, for an Azure application associated with the Azure AD tenant.
         * 
         * @return builder
         * 
         */
        public Builder clientId(String clientId) {
            return clientId(Output.of(clientId));
        }

        /**
         * @param enabled Specifies whether Encryption at Rest is enabled for an Atlas project. To disable Encryption at Rest, pass only this parameter with a value of false. When you disable Encryption at Rest, Atlas also removes the configuration details.
         * 
         * @return builder
         * 
         */
        public Builder enabled(Output<Boolean> enabled) {
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
         * @param keyIdentifier The unique identifier of a key in an Azure Key Vault.
         * 
         * @return builder
         * 
         */
        public Builder keyIdentifier(@Nullable Output<String> keyIdentifier) {
            $.keyIdentifier = keyIdentifier;
            return this;
        }

        /**
         * @param keyIdentifier The unique identifier of a key in an Azure Key Vault.
         * 
         * @return builder
         * 
         */
        public Builder keyIdentifier(String keyIdentifier) {
            return keyIdentifier(Output.of(keyIdentifier));
        }

        /**
         * @param keyVaultName The name of an Azure Key Vault containing your key.
         * 
         * @return builder
         * 
         */
        public Builder keyVaultName(@Nullable Output<String> keyVaultName) {
            $.keyVaultName = keyVaultName;
            return this;
        }

        /**
         * @param keyVaultName The name of an Azure Key Vault containing your key.
         * 
         * @return builder
         * 
         */
        public Builder keyVaultName(String keyVaultName) {
            return keyVaultName(Output.of(keyVaultName));
        }

        /**
         * @param resourceGroupName The name of the Azure Resource group that contains an Azure Key Vault.
         * 
         * @return builder
         * 
         */
        public Builder resourceGroupName(@Nullable Output<String> resourceGroupName) {
            $.resourceGroupName = resourceGroupName;
            return this;
        }

        /**
         * @param resourceGroupName The name of the Azure Resource group that contains an Azure Key Vault.
         * 
         * @return builder
         * 
         */
        public Builder resourceGroupName(String resourceGroupName) {
            return resourceGroupName(Output.of(resourceGroupName));
        }

        /**
         * @param secret The secret associated with the Azure Key Vault specified by azureKeyVault.tenantID.
         * 
         * @return builder
         * 
         */
        public Builder secret(@Nullable Output<String> secret) {
            $.secret = secret;
            return this;
        }

        /**
         * @param secret The secret associated with the Azure Key Vault specified by azureKeyVault.tenantID.
         * 
         * @return builder
         * 
         */
        public Builder secret(String secret) {
            return secret(Output.of(secret));
        }

        /**
         * @param subscriptionId The unique identifier associated with an Azure subscription.
         * 
         * @return builder
         * 
         */
        public Builder subscriptionId(@Nullable Output<String> subscriptionId) {
            $.subscriptionId = subscriptionId;
            return this;
        }

        /**
         * @param subscriptionId The unique identifier associated with an Azure subscription.
         * 
         * @return builder
         * 
         */
        public Builder subscriptionId(String subscriptionId) {
            return subscriptionId(Output.of(subscriptionId));
        }

        /**
         * @param tenantId The unique identifier for an Azure AD tenant within an Azure subscription.
         * 
         * @return builder
         * 
         */
        public Builder tenantId(@Nullable Output<String> tenantId) {
            $.tenantId = tenantId;
            return this;
        }

        /**
         * @param tenantId The unique identifier for an Azure AD tenant within an Azure subscription.
         * 
         * @return builder
         * 
         */
        public Builder tenantId(String tenantId) {
            return tenantId(Output.of(tenantId));
        }

        public EncryptionAtRestAzureKeyVaultConfigArgs build() {
            $.enabled = Objects.requireNonNull($.enabled, "expected parameter 'enabled' to be non-null");
            return $;
        }
    }

}
