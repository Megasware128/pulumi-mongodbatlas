// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Mongodbatlas.Outputs
{

    [OutputType]
    public sealed class EncryptionAtRestAzureKeyVault
    {
        /// <summary>
        /// The Azure environment where the Azure account credentials reside. Valid values are the following: AZURE, AZURE_CHINA, AZURE_GERMANY
        /// </summary>
        public readonly string AzureEnvironment;
        /// <summary>
        /// The client ID, also known as the application ID, for an Azure application associated with the Azure AD tenant.
        /// </summary>
        public readonly string ClientId;
        /// <summary>
        /// Specifies whether Encryption at Rest is enabled for an Atlas project. To disable Encryption at Rest, pass only this parameter with a value of false. When you disable Encryption at Rest, Atlas also removes the configuration details.
        /// </summary>
        public readonly bool Enabled;
        /// <summary>
        /// The unique identifier of a key in an Azure Key Vault.
        /// </summary>
        public readonly string KeyIdentifier;
        /// <summary>
        /// The name of an Azure Key Vault containing your key.
        /// </summary>
        public readonly string KeyVaultName;
        /// <summary>
        /// The name of the Azure Resource group that contains an Azure Key Vault.
        /// </summary>
        public readonly string ResourceGroupName;
        /// <summary>
        /// The secret associated with the Azure Key Vault specified by azureKeyVault.tenantID.
        /// </summary>
        public readonly string Secret;
        /// <summary>
        /// The unique identifier associated with an Azure subscription.
        /// </summary>
        public readonly string SubscriptionId;
        /// <summary>
        /// The unique identifier for an Azure AD tenant within an Azure subscription.
        /// </summary>
        public readonly string TenantId;

        [OutputConstructor]
        private EncryptionAtRestAzureKeyVault(
            string azureEnvironment,

            string clientId,

            bool enabled,

            string keyIdentifier,

            string keyVaultName,

            string resourceGroupName,

            string secret,

            string subscriptionId,

            string tenantId)
        {
            AzureEnvironment = azureEnvironment;
            ClientId = clientId;
            Enabled = enabled;
            KeyIdentifier = keyIdentifier;
            KeyVaultName = keyVaultName;
            ResourceGroupName = resourceGroupName;
            Secret = secret;
            SubscriptionId = subscriptionId;
            TenantId = tenantId;
        }
    }
}