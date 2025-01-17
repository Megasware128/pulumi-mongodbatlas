// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Mongodbatlas
{
    [MongodbatlasResourceType("mongodbatlas:index/encryptionAtRest:EncryptionAtRest")]
    public partial class EncryptionAtRest : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Specifies AWS KMS configuration details and whether Encryption at Rest is enabled for an Atlas project.
        /// </summary>
        [Output("awsKms")]
        public Output<ImmutableDictionary<string, string>?> AwsKms { get; private set; } = null!;

        [Output("awsKmsConfig")]
        public Output<Outputs.EncryptionAtRestAwsKmsConfig?> AwsKmsConfig { get; private set; } = null!;

        /// <summary>
        /// Specifies Azure Key Vault configuration details and whether Encryption at Rest is enabled for an Atlas project.
        /// </summary>
        [Output("azureKeyVault")]
        public Output<ImmutableDictionary<string, string>?> AzureKeyVault { get; private set; } = null!;

        [Output("azureKeyVaultConfig")]
        public Output<Outputs.EncryptionAtRestAzureKeyVaultConfig?> AzureKeyVaultConfig { get; private set; } = null!;

        /// <summary>
        /// Specifies GCP KMS configuration details and whether Encryption at Rest is enabled for an Atlas project.
        /// </summary>
        [Output("googleCloudKms")]
        public Output<ImmutableDictionary<string, string>?> GoogleCloudKms { get; private set; } = null!;

        [Output("googleCloudKmsConfig")]
        public Output<Outputs.EncryptionAtRestGoogleCloudKmsConfig?> GoogleCloudKmsConfig { get; private set; } = null!;

        /// <summary>
        /// The unique identifier for the project.
        /// </summary>
        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;


        /// <summary>
        /// Create a EncryptionAtRest resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public EncryptionAtRest(string name, EncryptionAtRestArgs args, CustomResourceOptions? options = null)
            : base("mongodbatlas:index/encryptionAtRest:EncryptionAtRest", name, args ?? new EncryptionAtRestArgs(), MakeResourceOptions(options, ""))
        {
        }

        private EncryptionAtRest(string name, Input<string> id, EncryptionAtRestState? state = null, CustomResourceOptions? options = null)
            : base("mongodbatlas:index/encryptionAtRest:EncryptionAtRest", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing EncryptionAtRest resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static EncryptionAtRest Get(string name, Input<string> id, EncryptionAtRestState? state = null, CustomResourceOptions? options = null)
        {
            return new EncryptionAtRest(name, id, state, options);
        }
    }

    public sealed class EncryptionAtRestArgs : global::Pulumi.ResourceArgs
    {
        [Input("awsKms")]
        private InputMap<string>? _awsKms;

        /// <summary>
        /// Specifies AWS KMS configuration details and whether Encryption at Rest is enabled for an Atlas project.
        /// </summary>
        [Obsolete(@"use aws_kms_config instead")]
        public InputMap<string> AwsKms
        {
            get => _awsKms ?? (_awsKms = new InputMap<string>());
            set => _awsKms = value;
        }

        [Input("awsKmsConfig")]
        public Input<Inputs.EncryptionAtRestAwsKmsConfigArgs>? AwsKmsConfig { get; set; }

        [Input("azureKeyVault")]
        private InputMap<string>? _azureKeyVault;

        /// <summary>
        /// Specifies Azure Key Vault configuration details and whether Encryption at Rest is enabled for an Atlas project.
        /// </summary>
        [Obsolete(@"use azure_key_vault_config instead")]
        public InputMap<string> AzureKeyVault
        {
            get => _azureKeyVault ?? (_azureKeyVault = new InputMap<string>());
            set => _azureKeyVault = value;
        }

        [Input("azureKeyVaultConfig")]
        public Input<Inputs.EncryptionAtRestAzureKeyVaultConfigArgs>? AzureKeyVaultConfig { get; set; }

        [Input("googleCloudKms")]
        private InputMap<string>? _googleCloudKms;

        /// <summary>
        /// Specifies GCP KMS configuration details and whether Encryption at Rest is enabled for an Atlas project.
        /// </summary>
        [Obsolete(@"use google_cloud_kms_config instead")]
        public InputMap<string> GoogleCloudKms
        {
            get => _googleCloudKms ?? (_googleCloudKms = new InputMap<string>());
            set => _googleCloudKms = value;
        }

        [Input("googleCloudKmsConfig")]
        public Input<Inputs.EncryptionAtRestGoogleCloudKmsConfigArgs>? GoogleCloudKmsConfig { get; set; }

        /// <summary>
        /// The unique identifier for the project.
        /// </summary>
        [Input("projectId", required: true)]
        public Input<string> ProjectId { get; set; } = null!;

        public EncryptionAtRestArgs()
        {
        }
        public static new EncryptionAtRestArgs Empty => new EncryptionAtRestArgs();
    }

    public sealed class EncryptionAtRestState : global::Pulumi.ResourceArgs
    {
        [Input("awsKms")]
        private InputMap<string>? _awsKms;

        /// <summary>
        /// Specifies AWS KMS configuration details and whether Encryption at Rest is enabled for an Atlas project.
        /// </summary>
        [Obsolete(@"use aws_kms_config instead")]
        public InputMap<string> AwsKms
        {
            get => _awsKms ?? (_awsKms = new InputMap<string>());
            set => _awsKms = value;
        }

        [Input("awsKmsConfig")]
        public Input<Inputs.EncryptionAtRestAwsKmsConfigGetArgs>? AwsKmsConfig { get; set; }

        [Input("azureKeyVault")]
        private InputMap<string>? _azureKeyVault;

        /// <summary>
        /// Specifies Azure Key Vault configuration details and whether Encryption at Rest is enabled for an Atlas project.
        /// </summary>
        [Obsolete(@"use azure_key_vault_config instead")]
        public InputMap<string> AzureKeyVault
        {
            get => _azureKeyVault ?? (_azureKeyVault = new InputMap<string>());
            set => _azureKeyVault = value;
        }

        [Input("azureKeyVaultConfig")]
        public Input<Inputs.EncryptionAtRestAzureKeyVaultConfigGetArgs>? AzureKeyVaultConfig { get; set; }

        [Input("googleCloudKms")]
        private InputMap<string>? _googleCloudKms;

        /// <summary>
        /// Specifies GCP KMS configuration details and whether Encryption at Rest is enabled for an Atlas project.
        /// </summary>
        [Obsolete(@"use google_cloud_kms_config instead")]
        public InputMap<string> GoogleCloudKms
        {
            get => _googleCloudKms ?? (_googleCloudKms = new InputMap<string>());
            set => _googleCloudKms = value;
        }

        [Input("googleCloudKmsConfig")]
        public Input<Inputs.EncryptionAtRestGoogleCloudKmsConfigGetArgs>? GoogleCloudKmsConfig { get; set; }

        /// <summary>
        /// The unique identifier for the project.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        public EncryptionAtRestState()
        {
        }
        public static new EncryptionAtRestState Empty => new EncryptionAtRestState();
    }
}
