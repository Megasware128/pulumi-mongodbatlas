// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mongodbatlas

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// `FederatedSettingsIdentityProvider` provides an Atlas federated settings identity provider resource provides a subset of settings to be maintained post import of the existing resource.
// ## Example Usage
//
// > **IMPORTANT** You **MUST** import this resource before you can manage it with this provider.
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-mongodbatlas/sdk/go/mongodbatlas"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := mongodbatlas.NewFederatedSettingsIdentityProvider(ctx, "identityProvider", &mongodbatlas.FederatedSettingsIdentityProviderArgs{
//				AssociatedDomains: pulumi.StringArray{
//					pulumi.String("yourdomain.com"),
//				},
//				FederationSettingsId:       pulumi.String("627a9687f7f7f7f774de306f14"),
//				IssuerUri:                  pulumi.String("http://www.okta.com/exk17q7f7f7f7fp50h8"),
//				RequestBinding:             pulumi.String("HTTP-POST"),
//				ResponseSignatureAlgorithm: pulumi.String("SHA-256"),
//				SsoDebugEnabled:            pulumi.Bool(true),
//				SsoUrl:                     pulumi.String("https://mysso.oktapreview.com/app/mysso_terraformtestsso/exk17q7f7f7f7f50h8/sso/saml"),
//				Status:                     pulumi.String("ACTIVE"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// Identity Provider **must** be imported before using federation_settings_id-idp_id, e.g.
//
// ```sh
//
//	$ pulumi import mongodbatlas:index/federatedSettingsIdentityProvider:FederatedSettingsIdentityProvider identity_provider 6287a663c660f52b1c441c6c-0oad4fas87jL5Xnk1297
//
// ```
//
//	For more information see[MongoDB Atlas API Reference.](https://www.mongodb.com/docs/atlas/reference/api/federation-configuration/)
type FederatedSettingsIdentityProvider struct {
	pulumi.CustomResourceState

	// List that contains the domains associated with the identity provider.
	AssociatedDomains pulumi.StringArrayOutput `pulumi:"associatedDomains"`
	// Unique 24-hexadecimal digit string that identifies the federated authentication configuration.
	FederationSettingsId pulumi.StringOutput `pulumi:"federationSettingsId"`
	// Unique string that identifies the issuer of the SAML
	IssuerUri pulumi.StringOutput `pulumi:"issuerUri"`
	// Human-readable label that identifies the identity provider.
	Name pulumi.StringOutput `pulumi:"name"`
	// Unique 20-hexadecimal digit string that identifies the IdP.
	OktaIdpId pulumi.StringOutput `pulumi:"oktaIdpId"`
	// SAML Authentication Request Protocol HTTP method binding (POST or REDIRECT) that Federated Authentication uses to send the authentication request. Atlas supports the following binding values:
	// - HTTP POST
	// - HTTP REDIRECT
	RequestBinding pulumi.StringOutput `pulumi:"requestBinding"`
	// Signature algorithm that Federated Authentication uses to encrypt the identity provider signature.  Valid values include SHA-1 and SHA-256.
	ResponseSignatureAlgorithm pulumi.StringOutput `pulumi:"responseSignatureAlgorithm"`
	// Flag that indicates whether the identity provider has SSO debug enabled.
	SsoDebugEnabled pulumi.BoolOutput `pulumi:"ssoDebugEnabled"`
	// Unique string that identifies the intended audience of the SAML assertion.
	SsoUrl pulumi.StringOutput `pulumi:"ssoUrl"`
	// String enum that indicates whether the identity provider is active or not. Accepted values are ACTIVE or INACTIVE.
	Status pulumi.StringOutput `pulumi:"status"`
}

// NewFederatedSettingsIdentityProvider registers a new resource with the given unique name, arguments, and options.
func NewFederatedSettingsIdentityProvider(ctx *pulumi.Context,
	name string, args *FederatedSettingsIdentityProviderArgs, opts ...pulumi.ResourceOption) (*FederatedSettingsIdentityProvider, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.FederationSettingsId == nil {
		return nil, errors.New("invalid value for required argument 'FederationSettingsId'")
	}
	if args.IssuerUri == nil {
		return nil, errors.New("invalid value for required argument 'IssuerUri'")
	}
	if args.RequestBinding == nil {
		return nil, errors.New("invalid value for required argument 'RequestBinding'")
	}
	if args.ResponseSignatureAlgorithm == nil {
		return nil, errors.New("invalid value for required argument 'ResponseSignatureAlgorithm'")
	}
	if args.SsoDebugEnabled == nil {
		return nil, errors.New("invalid value for required argument 'SsoDebugEnabled'")
	}
	if args.SsoUrl == nil {
		return nil, errors.New("invalid value for required argument 'SsoUrl'")
	}
	if args.Status == nil {
		return nil, errors.New("invalid value for required argument 'Status'")
	}
	var resource FederatedSettingsIdentityProvider
	err := ctx.RegisterResource("mongodbatlas:index/federatedSettingsIdentityProvider:FederatedSettingsIdentityProvider", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFederatedSettingsIdentityProvider gets an existing FederatedSettingsIdentityProvider resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFederatedSettingsIdentityProvider(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FederatedSettingsIdentityProviderState, opts ...pulumi.ResourceOption) (*FederatedSettingsIdentityProvider, error) {
	var resource FederatedSettingsIdentityProvider
	err := ctx.ReadResource("mongodbatlas:index/federatedSettingsIdentityProvider:FederatedSettingsIdentityProvider", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FederatedSettingsIdentityProvider resources.
type federatedSettingsIdentityProviderState struct {
	// List that contains the domains associated with the identity provider.
	AssociatedDomains []string `pulumi:"associatedDomains"`
	// Unique 24-hexadecimal digit string that identifies the federated authentication configuration.
	FederationSettingsId *string `pulumi:"federationSettingsId"`
	// Unique string that identifies the issuer of the SAML
	IssuerUri *string `pulumi:"issuerUri"`
	// Human-readable label that identifies the identity provider.
	Name *string `pulumi:"name"`
	// Unique 20-hexadecimal digit string that identifies the IdP.
	OktaIdpId *string `pulumi:"oktaIdpId"`
	// SAML Authentication Request Protocol HTTP method binding (POST or REDIRECT) that Federated Authentication uses to send the authentication request. Atlas supports the following binding values:
	// - HTTP POST
	// - HTTP REDIRECT
	RequestBinding *string `pulumi:"requestBinding"`
	// Signature algorithm that Federated Authentication uses to encrypt the identity provider signature.  Valid values include SHA-1 and SHA-256.
	ResponseSignatureAlgorithm *string `pulumi:"responseSignatureAlgorithm"`
	// Flag that indicates whether the identity provider has SSO debug enabled.
	SsoDebugEnabled *bool `pulumi:"ssoDebugEnabled"`
	// Unique string that identifies the intended audience of the SAML assertion.
	SsoUrl *string `pulumi:"ssoUrl"`
	// String enum that indicates whether the identity provider is active or not. Accepted values are ACTIVE or INACTIVE.
	Status *string `pulumi:"status"`
}

type FederatedSettingsIdentityProviderState struct {
	// List that contains the domains associated with the identity provider.
	AssociatedDomains pulumi.StringArrayInput
	// Unique 24-hexadecimal digit string that identifies the federated authentication configuration.
	FederationSettingsId pulumi.StringPtrInput
	// Unique string that identifies the issuer of the SAML
	IssuerUri pulumi.StringPtrInput
	// Human-readable label that identifies the identity provider.
	Name pulumi.StringPtrInput
	// Unique 20-hexadecimal digit string that identifies the IdP.
	OktaIdpId pulumi.StringPtrInput
	// SAML Authentication Request Protocol HTTP method binding (POST or REDIRECT) that Federated Authentication uses to send the authentication request. Atlas supports the following binding values:
	// - HTTP POST
	// - HTTP REDIRECT
	RequestBinding pulumi.StringPtrInput
	// Signature algorithm that Federated Authentication uses to encrypt the identity provider signature.  Valid values include SHA-1 and SHA-256.
	ResponseSignatureAlgorithm pulumi.StringPtrInput
	// Flag that indicates whether the identity provider has SSO debug enabled.
	SsoDebugEnabled pulumi.BoolPtrInput
	// Unique string that identifies the intended audience of the SAML assertion.
	SsoUrl pulumi.StringPtrInput
	// String enum that indicates whether the identity provider is active or not. Accepted values are ACTIVE or INACTIVE.
	Status pulumi.StringPtrInput
}

func (FederatedSettingsIdentityProviderState) ElementType() reflect.Type {
	return reflect.TypeOf((*federatedSettingsIdentityProviderState)(nil)).Elem()
}

type federatedSettingsIdentityProviderArgs struct {
	// List that contains the domains associated with the identity provider.
	AssociatedDomains []string `pulumi:"associatedDomains"`
	// Unique 24-hexadecimal digit string that identifies the federated authentication configuration.
	FederationSettingsId string `pulumi:"federationSettingsId"`
	// Unique string that identifies the issuer of the SAML
	IssuerUri string `pulumi:"issuerUri"`
	// Human-readable label that identifies the identity provider.
	Name *string `pulumi:"name"`
	// SAML Authentication Request Protocol HTTP method binding (POST or REDIRECT) that Federated Authentication uses to send the authentication request. Atlas supports the following binding values:
	// - HTTP POST
	// - HTTP REDIRECT
	RequestBinding string `pulumi:"requestBinding"`
	// Signature algorithm that Federated Authentication uses to encrypt the identity provider signature.  Valid values include SHA-1 and SHA-256.
	ResponseSignatureAlgorithm string `pulumi:"responseSignatureAlgorithm"`
	// Flag that indicates whether the identity provider has SSO debug enabled.
	SsoDebugEnabled bool `pulumi:"ssoDebugEnabled"`
	// Unique string that identifies the intended audience of the SAML assertion.
	SsoUrl string `pulumi:"ssoUrl"`
	// String enum that indicates whether the identity provider is active or not. Accepted values are ACTIVE or INACTIVE.
	Status string `pulumi:"status"`
}

// The set of arguments for constructing a FederatedSettingsIdentityProvider resource.
type FederatedSettingsIdentityProviderArgs struct {
	// List that contains the domains associated with the identity provider.
	AssociatedDomains pulumi.StringArrayInput
	// Unique 24-hexadecimal digit string that identifies the federated authentication configuration.
	FederationSettingsId pulumi.StringInput
	// Unique string that identifies the issuer of the SAML
	IssuerUri pulumi.StringInput
	// Human-readable label that identifies the identity provider.
	Name pulumi.StringPtrInput
	// SAML Authentication Request Protocol HTTP method binding (POST or REDIRECT) that Federated Authentication uses to send the authentication request. Atlas supports the following binding values:
	// - HTTP POST
	// - HTTP REDIRECT
	RequestBinding pulumi.StringInput
	// Signature algorithm that Federated Authentication uses to encrypt the identity provider signature.  Valid values include SHA-1 and SHA-256.
	ResponseSignatureAlgorithm pulumi.StringInput
	// Flag that indicates whether the identity provider has SSO debug enabled.
	SsoDebugEnabled pulumi.BoolInput
	// Unique string that identifies the intended audience of the SAML assertion.
	SsoUrl pulumi.StringInput
	// String enum that indicates whether the identity provider is active or not. Accepted values are ACTIVE or INACTIVE.
	Status pulumi.StringInput
}

func (FederatedSettingsIdentityProviderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*federatedSettingsIdentityProviderArgs)(nil)).Elem()
}

type FederatedSettingsIdentityProviderInput interface {
	pulumi.Input

	ToFederatedSettingsIdentityProviderOutput() FederatedSettingsIdentityProviderOutput
	ToFederatedSettingsIdentityProviderOutputWithContext(ctx context.Context) FederatedSettingsIdentityProviderOutput
}

func (*FederatedSettingsIdentityProvider) ElementType() reflect.Type {
	return reflect.TypeOf((**FederatedSettingsIdentityProvider)(nil)).Elem()
}

func (i *FederatedSettingsIdentityProvider) ToFederatedSettingsIdentityProviderOutput() FederatedSettingsIdentityProviderOutput {
	return i.ToFederatedSettingsIdentityProviderOutputWithContext(context.Background())
}

func (i *FederatedSettingsIdentityProvider) ToFederatedSettingsIdentityProviderOutputWithContext(ctx context.Context) FederatedSettingsIdentityProviderOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FederatedSettingsIdentityProviderOutput)
}

// FederatedSettingsIdentityProviderArrayInput is an input type that accepts FederatedSettingsIdentityProviderArray and FederatedSettingsIdentityProviderArrayOutput values.
// You can construct a concrete instance of `FederatedSettingsIdentityProviderArrayInput` via:
//
//	FederatedSettingsIdentityProviderArray{ FederatedSettingsIdentityProviderArgs{...} }
type FederatedSettingsIdentityProviderArrayInput interface {
	pulumi.Input

	ToFederatedSettingsIdentityProviderArrayOutput() FederatedSettingsIdentityProviderArrayOutput
	ToFederatedSettingsIdentityProviderArrayOutputWithContext(context.Context) FederatedSettingsIdentityProviderArrayOutput
}

type FederatedSettingsIdentityProviderArray []FederatedSettingsIdentityProviderInput

func (FederatedSettingsIdentityProviderArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FederatedSettingsIdentityProvider)(nil)).Elem()
}

func (i FederatedSettingsIdentityProviderArray) ToFederatedSettingsIdentityProviderArrayOutput() FederatedSettingsIdentityProviderArrayOutput {
	return i.ToFederatedSettingsIdentityProviderArrayOutputWithContext(context.Background())
}

func (i FederatedSettingsIdentityProviderArray) ToFederatedSettingsIdentityProviderArrayOutputWithContext(ctx context.Context) FederatedSettingsIdentityProviderArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FederatedSettingsIdentityProviderArrayOutput)
}

// FederatedSettingsIdentityProviderMapInput is an input type that accepts FederatedSettingsIdentityProviderMap and FederatedSettingsIdentityProviderMapOutput values.
// You can construct a concrete instance of `FederatedSettingsIdentityProviderMapInput` via:
//
//	FederatedSettingsIdentityProviderMap{ "key": FederatedSettingsIdentityProviderArgs{...} }
type FederatedSettingsIdentityProviderMapInput interface {
	pulumi.Input

	ToFederatedSettingsIdentityProviderMapOutput() FederatedSettingsIdentityProviderMapOutput
	ToFederatedSettingsIdentityProviderMapOutputWithContext(context.Context) FederatedSettingsIdentityProviderMapOutput
}

type FederatedSettingsIdentityProviderMap map[string]FederatedSettingsIdentityProviderInput

func (FederatedSettingsIdentityProviderMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FederatedSettingsIdentityProvider)(nil)).Elem()
}

func (i FederatedSettingsIdentityProviderMap) ToFederatedSettingsIdentityProviderMapOutput() FederatedSettingsIdentityProviderMapOutput {
	return i.ToFederatedSettingsIdentityProviderMapOutputWithContext(context.Background())
}

func (i FederatedSettingsIdentityProviderMap) ToFederatedSettingsIdentityProviderMapOutputWithContext(ctx context.Context) FederatedSettingsIdentityProviderMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FederatedSettingsIdentityProviderMapOutput)
}

type FederatedSettingsIdentityProviderOutput struct{ *pulumi.OutputState }

func (FederatedSettingsIdentityProviderOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FederatedSettingsIdentityProvider)(nil)).Elem()
}

func (o FederatedSettingsIdentityProviderOutput) ToFederatedSettingsIdentityProviderOutput() FederatedSettingsIdentityProviderOutput {
	return o
}

func (o FederatedSettingsIdentityProviderOutput) ToFederatedSettingsIdentityProviderOutputWithContext(ctx context.Context) FederatedSettingsIdentityProviderOutput {
	return o
}

// List that contains the domains associated with the identity provider.
func (o FederatedSettingsIdentityProviderOutput) AssociatedDomains() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *FederatedSettingsIdentityProvider) pulumi.StringArrayOutput { return v.AssociatedDomains }).(pulumi.StringArrayOutput)
}

// Unique 24-hexadecimal digit string that identifies the federated authentication configuration.
func (o FederatedSettingsIdentityProviderOutput) FederationSettingsId() pulumi.StringOutput {
	return o.ApplyT(func(v *FederatedSettingsIdentityProvider) pulumi.StringOutput { return v.FederationSettingsId }).(pulumi.StringOutput)
}

// Unique string that identifies the issuer of the SAML
func (o FederatedSettingsIdentityProviderOutput) IssuerUri() pulumi.StringOutput {
	return o.ApplyT(func(v *FederatedSettingsIdentityProvider) pulumi.StringOutput { return v.IssuerUri }).(pulumi.StringOutput)
}

// Human-readable label that identifies the identity provider.
func (o FederatedSettingsIdentityProviderOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *FederatedSettingsIdentityProvider) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Unique 20-hexadecimal digit string that identifies the IdP.
func (o FederatedSettingsIdentityProviderOutput) OktaIdpId() pulumi.StringOutput {
	return o.ApplyT(func(v *FederatedSettingsIdentityProvider) pulumi.StringOutput { return v.OktaIdpId }).(pulumi.StringOutput)
}

// SAML Authentication Request Protocol HTTP method binding (POST or REDIRECT) that Federated Authentication uses to send the authentication request. Atlas supports the following binding values:
// - HTTP POST
// - HTTP REDIRECT
func (o FederatedSettingsIdentityProviderOutput) RequestBinding() pulumi.StringOutput {
	return o.ApplyT(func(v *FederatedSettingsIdentityProvider) pulumi.StringOutput { return v.RequestBinding }).(pulumi.StringOutput)
}

// Signature algorithm that Federated Authentication uses to encrypt the identity provider signature.  Valid values include SHA-1 and SHA-256.
func (o FederatedSettingsIdentityProviderOutput) ResponseSignatureAlgorithm() pulumi.StringOutput {
	return o.ApplyT(func(v *FederatedSettingsIdentityProvider) pulumi.StringOutput { return v.ResponseSignatureAlgorithm }).(pulumi.StringOutput)
}

// Flag that indicates whether the identity provider has SSO debug enabled.
func (o FederatedSettingsIdentityProviderOutput) SsoDebugEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *FederatedSettingsIdentityProvider) pulumi.BoolOutput { return v.SsoDebugEnabled }).(pulumi.BoolOutput)
}

// Unique string that identifies the intended audience of the SAML assertion.
func (o FederatedSettingsIdentityProviderOutput) SsoUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *FederatedSettingsIdentityProvider) pulumi.StringOutput { return v.SsoUrl }).(pulumi.StringOutput)
}

// String enum that indicates whether the identity provider is active or not. Accepted values are ACTIVE or INACTIVE.
func (o FederatedSettingsIdentityProviderOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *FederatedSettingsIdentityProvider) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

type FederatedSettingsIdentityProviderArrayOutput struct{ *pulumi.OutputState }

func (FederatedSettingsIdentityProviderArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FederatedSettingsIdentityProvider)(nil)).Elem()
}

func (o FederatedSettingsIdentityProviderArrayOutput) ToFederatedSettingsIdentityProviderArrayOutput() FederatedSettingsIdentityProviderArrayOutput {
	return o
}

func (o FederatedSettingsIdentityProviderArrayOutput) ToFederatedSettingsIdentityProviderArrayOutputWithContext(ctx context.Context) FederatedSettingsIdentityProviderArrayOutput {
	return o
}

func (o FederatedSettingsIdentityProviderArrayOutput) Index(i pulumi.IntInput) FederatedSettingsIdentityProviderOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FederatedSettingsIdentityProvider {
		return vs[0].([]*FederatedSettingsIdentityProvider)[vs[1].(int)]
	}).(FederatedSettingsIdentityProviderOutput)
}

type FederatedSettingsIdentityProviderMapOutput struct{ *pulumi.OutputState }

func (FederatedSettingsIdentityProviderMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FederatedSettingsIdentityProvider)(nil)).Elem()
}

func (o FederatedSettingsIdentityProviderMapOutput) ToFederatedSettingsIdentityProviderMapOutput() FederatedSettingsIdentityProviderMapOutput {
	return o
}

func (o FederatedSettingsIdentityProviderMapOutput) ToFederatedSettingsIdentityProviderMapOutputWithContext(ctx context.Context) FederatedSettingsIdentityProviderMapOutput {
	return o
}

func (o FederatedSettingsIdentityProviderMapOutput) MapIndex(k pulumi.StringInput) FederatedSettingsIdentityProviderOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FederatedSettingsIdentityProvider {
		return vs[0].(map[string]*FederatedSettingsIdentityProvider)[vs[1].(string)]
	}).(FederatedSettingsIdentityProviderOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FederatedSettingsIdentityProviderInput)(nil)).Elem(), &FederatedSettingsIdentityProvider{})
	pulumi.RegisterInputType(reflect.TypeOf((*FederatedSettingsIdentityProviderArrayInput)(nil)).Elem(), FederatedSettingsIdentityProviderArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FederatedSettingsIdentityProviderMapInput)(nil)).Elem(), FederatedSettingsIdentityProviderMap{})
	pulumi.RegisterOutputType(FederatedSettingsIdentityProviderOutput{})
	pulumi.RegisterOutputType(FederatedSettingsIdentityProviderArrayOutput{})
	pulumi.RegisterOutputType(FederatedSettingsIdentityProviderMapOutput{})
}
