// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package mongodbatlas

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// `.EncryptionAtRest` Atlas encrypts your data at rest using encrypted storage media.
// Using keys you manage with AWS KMS, Atlas encrypts your data a second time when it writes it to the MongoDB encrypted storage engine.
// You can use the following clouds: AWS CMK, AZURE KEY VAULT and GOOGLE KEY VAULT to encrypt the MongoDB master encryption keys.
//
// > **NOTE:** Groups and projects are synonymous terms. You may find `groupId` in the official documentation.
type EncryptionAtRest struct {
	pulumi.CustomResourceState

	// Specifies AWS KMS configuration details and whether Encryption at Rest is enabled for an Atlas project.
	AwsKms EncryptionAtRestAwsKmsPtrOutput `pulumi:"awsKms"`
	// Specifies Azure Key Vault configuration details and whether Encryption at Rest is enabled for an Atlas project.
	AzureKeyVault EncryptionAtRestAzureKeyVaultPtrOutput `pulumi:"azureKeyVault"`
	// Specifies GCP KMS configuration details and whether Encryption at Rest is enabled for an Atlas project.
	GoogleCloudKms EncryptionAtRestGoogleCloudKmsPtrOutput `pulumi:"googleCloudKms"`
	// The unique identifier for the project.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
}

// NewEncryptionAtRest registers a new resource with the given unique name, arguments, and options.
func NewEncryptionAtRest(ctx *pulumi.Context,
	name string, args *EncryptionAtRestArgs, opts ...pulumi.ResourceOption) (*EncryptionAtRest, error) {
	if args == nil || args.ProjectId == nil {
		return nil, errors.New("missing required argument 'ProjectId'")
	}
	if args == nil {
		args = &EncryptionAtRestArgs{}
	}
	var resource EncryptionAtRest
	err := ctx.RegisterResource("mongodbatlas:index/encryptionAtRest:EncryptionAtRest", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEncryptionAtRest gets an existing EncryptionAtRest resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEncryptionAtRest(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EncryptionAtRestState, opts ...pulumi.ResourceOption) (*EncryptionAtRest, error) {
	var resource EncryptionAtRest
	err := ctx.ReadResource("mongodbatlas:index/encryptionAtRest:EncryptionAtRest", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EncryptionAtRest resources.
type encryptionAtRestState struct {
	// Specifies AWS KMS configuration details and whether Encryption at Rest is enabled for an Atlas project.
	AwsKms *EncryptionAtRestAwsKms `pulumi:"awsKms"`
	// Specifies Azure Key Vault configuration details and whether Encryption at Rest is enabled for an Atlas project.
	AzureKeyVault *EncryptionAtRestAzureKeyVault `pulumi:"azureKeyVault"`
	// Specifies GCP KMS configuration details and whether Encryption at Rest is enabled for an Atlas project.
	GoogleCloudKms *EncryptionAtRestGoogleCloudKms `pulumi:"googleCloudKms"`
	// The unique identifier for the project.
	ProjectId *string `pulumi:"projectId"`
}

type EncryptionAtRestState struct {
	// Specifies AWS KMS configuration details and whether Encryption at Rest is enabled for an Atlas project.
	AwsKms EncryptionAtRestAwsKmsPtrInput
	// Specifies Azure Key Vault configuration details and whether Encryption at Rest is enabled for an Atlas project.
	AzureKeyVault EncryptionAtRestAzureKeyVaultPtrInput
	// Specifies GCP KMS configuration details and whether Encryption at Rest is enabled for an Atlas project.
	GoogleCloudKms EncryptionAtRestGoogleCloudKmsPtrInput
	// The unique identifier for the project.
	ProjectId pulumi.StringPtrInput
}

func (EncryptionAtRestState) ElementType() reflect.Type {
	return reflect.TypeOf((*encryptionAtRestState)(nil)).Elem()
}

type encryptionAtRestArgs struct {
	// Specifies AWS KMS configuration details and whether Encryption at Rest is enabled for an Atlas project.
	AwsKms *EncryptionAtRestAwsKms `pulumi:"awsKms"`
	// Specifies Azure Key Vault configuration details and whether Encryption at Rest is enabled for an Atlas project.
	AzureKeyVault *EncryptionAtRestAzureKeyVault `pulumi:"azureKeyVault"`
	// Specifies GCP KMS configuration details and whether Encryption at Rest is enabled for an Atlas project.
	GoogleCloudKms *EncryptionAtRestGoogleCloudKms `pulumi:"googleCloudKms"`
	// The unique identifier for the project.
	ProjectId string `pulumi:"projectId"`
}

// The set of arguments for constructing a EncryptionAtRest resource.
type EncryptionAtRestArgs struct {
	// Specifies AWS KMS configuration details and whether Encryption at Rest is enabled for an Atlas project.
	AwsKms EncryptionAtRestAwsKmsPtrInput
	// Specifies Azure Key Vault configuration details and whether Encryption at Rest is enabled for an Atlas project.
	AzureKeyVault EncryptionAtRestAzureKeyVaultPtrInput
	// Specifies GCP KMS configuration details and whether Encryption at Rest is enabled for an Atlas project.
	GoogleCloudKms EncryptionAtRestGoogleCloudKmsPtrInput
	// The unique identifier for the project.
	ProjectId pulumi.StringInput
}

func (EncryptionAtRestArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*encryptionAtRestArgs)(nil)).Elem()
}