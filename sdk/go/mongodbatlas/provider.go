// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package mongodbatlas

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The provider type for the mongodbatlas package. By default, resources use package-wide configuration
// settings, however an explicit `Provider` instance may be created and passed during resource
// construction to achieve fine-grained programmatic control over provider settings. See the
// [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.
type Provider struct {
	pulumi.ProviderResourceState

	// MongoDB Atlas Base URL
	BaseUrl pulumi.StringPtrOutput `pulumi:"baseUrl"`
	// MongoDB Atlas Programmatic Private Key
	PrivateKey pulumi.StringOutput `pulumi:"privateKey"`
	// MongoDB Atlas Programmatic Public Key
	PublicKey pulumi.StringOutput `pulumi:"publicKey"`
	// MongoDB Realm Base URL
	RealmBaseUrl pulumi.StringPtrOutput `pulumi:"realmBaseUrl"`
}

// NewProvider registers a new resource with the given unique name, arguments, and options.
func NewProvider(ctx *pulumi.Context,
	name string, args *ProviderArgs, opts ...pulumi.ResourceOption) (*Provider, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PrivateKey == nil {
		return nil, errors.New("invalid value for required argument 'PrivateKey'")
	}
	if args.PublicKey == nil {
		return nil, errors.New("invalid value for required argument 'PublicKey'")
	}
	var resource Provider
	err := ctx.RegisterResource("pulumi:providers:mongodbatlas", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

type providerArgs struct {
	// MongoDB Atlas Base URL
	BaseUrl *string `pulumi:"baseUrl"`
	// MongoDB Atlas Programmatic Private Key
	PrivateKey string `pulumi:"privateKey"`
	// MongoDB Atlas Programmatic Public Key
	PublicKey string `pulumi:"publicKey"`
	// MongoDB Realm Base URL
	RealmBaseUrl *string `pulumi:"realmBaseUrl"`
}

// The set of arguments for constructing a Provider resource.
type ProviderArgs struct {
	// MongoDB Atlas Base URL
	BaseUrl pulumi.StringPtrInput
	// MongoDB Atlas Programmatic Private Key
	PrivateKey pulumi.StringInput
	// MongoDB Atlas Programmatic Public Key
	PublicKey pulumi.StringInput
	// MongoDB Realm Base URL
	RealmBaseUrl pulumi.StringPtrInput
}

func (ProviderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*providerArgs)(nil)).Elem()
}

type ProviderInput interface {
	pulumi.Input

	ToProviderOutput() ProviderOutput
	ToProviderOutputWithContext(ctx context.Context) ProviderOutput
}

func (*Provider) ElementType() reflect.Type {
	return reflect.TypeOf((**Provider)(nil)).Elem()
}

func (i *Provider) ToProviderOutput() ProviderOutput {
	return i.ToProviderOutputWithContext(context.Background())
}

func (i *Provider) ToProviderOutputWithContext(ctx context.Context) ProviderOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProviderOutput)
}

type ProviderOutput struct{ *pulumi.OutputState }

func (ProviderOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Provider)(nil)).Elem()
}

func (o ProviderOutput) ToProviderOutput() ProviderOutput {
	return o
}

func (o ProviderOutput) ToProviderOutputWithContext(ctx context.Context) ProviderOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProviderInput)(nil)).Elem(), &Provider{})
	pulumi.RegisterOutputType(ProviderOutput{})
}
