// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package mongodbatlas

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// `X509AuthenticationDatabaseUser` provides a X509 Authentication Database User resource. The X509AuthenticationDatabaseUser resource lets you manage MongoDB users who authenticate using X.509 certificates. You can manage these X.509 certificates or let Atlas do it for you.
//
// | Management  | Description  |
// |---|---|
// | Atlas  | Atlas manages your Certificate Authority and can generate certificates for your MongoDB users. No additional X.509 configuration is required.  |
// | Customer  |  You must provide a Certificate Authority and generate certificates for your MongoDB users. |
//
// > **NOTE:** Groups and projects are synonymous terms. You may find groupId in the official documentation.
//
// ## Example Usage
//
// ### S
// ### Example Usage: Generate an Atlas-managed X.509 certificate for a MongoDB user
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-mongodbatlas/sdk/v3/go/mongodbatlas"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		user, err := mongodbatlas.NewDatabaseUser(ctx, "user", &mongodbatlas.DatabaseUserArgs{
// 			ProjectId:    pulumi.String("<PROJECT-ID>"),
// 			Username:     pulumi.String("myUsername"),
// 			X509Type:     pulumi.String("MANAGED"),
// 			DatabaseName: pulumi.String(fmt.Sprintf("%v%v", "$", "external")),
// 			Roles: mongodbatlas.DatabaseUserRoleArray{
// 				&mongodbatlas.DatabaseUserRoleArgs{
// 					RoleName:     pulumi.String("atlasAdmin"),
// 					DatabaseName: pulumi.String("admin"),
// 				},
// 			},
// 			Labels: mongodbatlas.DatabaseUserLabelArray{
// 				&mongodbatlas.DatabaseUserLabelArgs{
// 					Key:   pulumi.String("My Key"),
// 					Value: pulumi.String("My Value"),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = mongodbatlas.NewX509AuthenticationDatabaseUser(ctx, "test", &mongodbatlas.X509AuthenticationDatabaseUserArgs{
// 			ProjectId:             user.ProjectId,
// 			Username:              user.Username,
// 			MonthsUntilExpiration: pulumi.Int(2),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Example Usage: Save a customer-managed X.509 configuration for an Atlas project
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-mongodbatlas/sdk/v3/go/mongodbatlas"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := mongodbatlas.NewX509AuthenticationDatabaseUser(ctx, "test", &mongodbatlas.X509AuthenticationDatabaseUserArgs{
// 			CustomerX509Cas: pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v", "  -----BEGIN CERTIFICATE-----\n", "  MIICmTCCAgICCQDZnHzklxsT9TANBgkqhkiG9w0BAQsFADCBkDELMAkGA1UEBhMC\n", "  VVMxDjAMBgNVBAgMBVRleGFzMQ8wDQYDVQQHDAZBdXN0aW4xETAPBgNVBAoMCHRl\n", "  c3QuY29tMQ0wCwYDVQQLDARUZXN0MREwDwYDVQQDDAh0ZXN0LmNvbTErMCkGCSqG\n", "  SIb3DQEJARYcbWVsaXNzYS5wbHVua2V0dEBtb25nb2RiLmNvbTAeFw0yMDAyMDQy\n", "  MDQ2MDFaFw0yMTAyMDMyMDQ2MDFaMIGQMQswCQYDVQQGEwJVUzEOMAwGA1UECAwF\n", "  VGV4YXMxDzANBgNVBAcMBkF1c3RpbjERMA8GA1UECgwIdGVzdC5jb20xDTALBgNV\n", "  BAsMBFRlc3QxETAPBgNVBAMMCHRlc3QuY29tMSswKQYJKoZIhvcNAQkBFhxtZWxp\n", "  c3NhLnBsdW5rZXR0QG1vbmdvZGIuY29tMIGfMA0GCSqGSIb3DQEBAQUAA4GNADCB\n", "  iQKBgQCf1LRqr1zftzdYx2Aj9G76tb0noMPtj6faGLlPji1+m6Rn7RWD9L0ntWAr\n", "  cURxvypa9jZ9MXFzDtLevvd3tHEmfrUT3ukNDX6+Jtc4kWm+Dh2A70Pd+deKZ2/O\n", "  Fh8audEKAESGXnTbeJCeQa1XKlIkjqQHBNwES5h1b9vJtFoLJwIDAQABMA0GCSqG\n", "  SIb3DQEBCwUAA4GBADMUncjEPV/MiZUcVNGmktP6BPmEqMXQWUDpdGW2+Tg2JtUA\n", "  7MMILtepBkFzLO+GlpZxeAlXO0wxiNgEmCRONgh4+t2w3e7a8GFijYQ99FHrAC5A\n", "  iul59bdl18gVqXia1Yeq/iK7Ohfy/Jwd7Hsm530elwkM/ZEkYDjBlZSXYdyz\n", "  -----END CERTIFICATE-----\"\n", "\n")),
// 			ProjectId:       pulumi.String("<PROJECT-ID>"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// X.509 Certificates for a User can be imported using project ID and username, in the format `project_id-username`, e.g.
//
// ```sh
//  $ pulumi import mongodbatlas:index/x509AuthenticationDatabaseUser:X509AuthenticationDatabaseUser test 1112222b3bf99403840e8934-myUsername
// ```
//
//  For more information see[MongoDB Atlas API Reference.](https://docs.atlas.mongodb.com/reference/api/x509-configuration-get-certificates/) Current X.509 Configuration can be imported using project ID, in the format `project_id`, e.g.
//
// ```sh
//  $ pulumi import mongodbatlas:index/x509AuthenticationDatabaseUser:X509AuthenticationDatabaseUser test 1112222b3bf99403840e8934
// ```
//
//  For more information see[MongoDB Atlas API Reference.](https://docs.atlas.mongodb.com/reference/api/x509-configuration-get-certificates/)
type X509AuthenticationDatabaseUser struct {
	pulumi.CustomResourceState

	// Array of objects where each details one unexpired database user certificate.
	Certificates X509AuthenticationDatabaseUserCertificateArrayOutput `pulumi:"certificates"`
	// Contains the last X.509 certificate and private key created for a database user.
	CurrentCertificate pulumi.StringOutput `pulumi:"currentCertificate"`
	// PEM string containing one or more customer CAs for database user authentication.
	CustomerX509Cas pulumi.StringPtrOutput `pulumi:"customerX509Cas"`
	// A number of months that the created certificate is valid for before expiry, up to 24 months. By default is 3.
	MonthsUntilExpiration pulumi.IntPtrOutput `pulumi:"monthsUntilExpiration"`
	// Identifier for the Atlas project associated with the X.509 configuration.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// Username of the database user to create a certificate for.
	Username pulumi.StringPtrOutput `pulumi:"username"`
}

// NewX509AuthenticationDatabaseUser registers a new resource with the given unique name, arguments, and options.
func NewX509AuthenticationDatabaseUser(ctx *pulumi.Context,
	name string, args *X509AuthenticationDatabaseUserArgs, opts ...pulumi.ResourceOption) (*X509AuthenticationDatabaseUser, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ProjectId == nil {
		return nil, errors.New("invalid value for required argument 'ProjectId'")
	}
	var resource X509AuthenticationDatabaseUser
	err := ctx.RegisterResource("mongodbatlas:index/x509AuthenticationDatabaseUser:X509AuthenticationDatabaseUser", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetX509AuthenticationDatabaseUser gets an existing X509AuthenticationDatabaseUser resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetX509AuthenticationDatabaseUser(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *X509AuthenticationDatabaseUserState, opts ...pulumi.ResourceOption) (*X509AuthenticationDatabaseUser, error) {
	var resource X509AuthenticationDatabaseUser
	err := ctx.ReadResource("mongodbatlas:index/x509AuthenticationDatabaseUser:X509AuthenticationDatabaseUser", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering X509AuthenticationDatabaseUser resources.
type x509authenticationDatabaseUserState struct {
	// Array of objects where each details one unexpired database user certificate.
	Certificates []X509AuthenticationDatabaseUserCertificate `pulumi:"certificates"`
	// Contains the last X.509 certificate and private key created for a database user.
	CurrentCertificate *string `pulumi:"currentCertificate"`
	// PEM string containing one or more customer CAs for database user authentication.
	CustomerX509Cas *string `pulumi:"customerX509Cas"`
	// A number of months that the created certificate is valid for before expiry, up to 24 months. By default is 3.
	MonthsUntilExpiration *int `pulumi:"monthsUntilExpiration"`
	// Identifier for the Atlas project associated with the X.509 configuration.
	ProjectId *string `pulumi:"projectId"`
	// Username of the database user to create a certificate for.
	Username *string `pulumi:"username"`
}

type X509AuthenticationDatabaseUserState struct {
	// Array of objects where each details one unexpired database user certificate.
	Certificates X509AuthenticationDatabaseUserCertificateArrayInput
	// Contains the last X.509 certificate and private key created for a database user.
	CurrentCertificate pulumi.StringPtrInput
	// PEM string containing one or more customer CAs for database user authentication.
	CustomerX509Cas pulumi.StringPtrInput
	// A number of months that the created certificate is valid for before expiry, up to 24 months. By default is 3.
	MonthsUntilExpiration pulumi.IntPtrInput
	// Identifier for the Atlas project associated with the X.509 configuration.
	ProjectId pulumi.StringPtrInput
	// Username of the database user to create a certificate for.
	Username pulumi.StringPtrInput
}

func (X509AuthenticationDatabaseUserState) ElementType() reflect.Type {
	return reflect.TypeOf((*x509authenticationDatabaseUserState)(nil)).Elem()
}

type x509authenticationDatabaseUserArgs struct {
	// PEM string containing one or more customer CAs for database user authentication.
	CustomerX509Cas *string `pulumi:"customerX509Cas"`
	// A number of months that the created certificate is valid for before expiry, up to 24 months. By default is 3.
	MonthsUntilExpiration *int `pulumi:"monthsUntilExpiration"`
	// Identifier for the Atlas project associated with the X.509 configuration.
	ProjectId string `pulumi:"projectId"`
	// Username of the database user to create a certificate for.
	Username *string `pulumi:"username"`
}

// The set of arguments for constructing a X509AuthenticationDatabaseUser resource.
type X509AuthenticationDatabaseUserArgs struct {
	// PEM string containing one or more customer CAs for database user authentication.
	CustomerX509Cas pulumi.StringPtrInput
	// A number of months that the created certificate is valid for before expiry, up to 24 months. By default is 3.
	MonthsUntilExpiration pulumi.IntPtrInput
	// Identifier for the Atlas project associated with the X.509 configuration.
	ProjectId pulumi.StringInput
	// Username of the database user to create a certificate for.
	Username pulumi.StringPtrInput
}

func (X509AuthenticationDatabaseUserArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*x509authenticationDatabaseUserArgs)(nil)).Elem()
}

type X509AuthenticationDatabaseUserInput interface {
	pulumi.Input

	ToX509AuthenticationDatabaseUserOutput() X509AuthenticationDatabaseUserOutput
	ToX509AuthenticationDatabaseUserOutputWithContext(ctx context.Context) X509AuthenticationDatabaseUserOutput
}

func (*X509AuthenticationDatabaseUser) ElementType() reflect.Type {
	return reflect.TypeOf((*X509AuthenticationDatabaseUser)(nil))
}

func (i *X509AuthenticationDatabaseUser) ToX509AuthenticationDatabaseUserOutput() X509AuthenticationDatabaseUserOutput {
	return i.ToX509AuthenticationDatabaseUserOutputWithContext(context.Background())
}

func (i *X509AuthenticationDatabaseUser) ToX509AuthenticationDatabaseUserOutputWithContext(ctx context.Context) X509AuthenticationDatabaseUserOutput {
	return pulumi.ToOutputWithContext(ctx, i).(X509AuthenticationDatabaseUserOutput)
}

func (i *X509AuthenticationDatabaseUser) ToX509AuthenticationDatabaseUserPtrOutput() X509AuthenticationDatabaseUserPtrOutput {
	return i.ToX509AuthenticationDatabaseUserPtrOutputWithContext(context.Background())
}

func (i *X509AuthenticationDatabaseUser) ToX509AuthenticationDatabaseUserPtrOutputWithContext(ctx context.Context) X509AuthenticationDatabaseUserPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(X509AuthenticationDatabaseUserPtrOutput)
}

type X509AuthenticationDatabaseUserPtrInput interface {
	pulumi.Input

	ToX509AuthenticationDatabaseUserPtrOutput() X509AuthenticationDatabaseUserPtrOutput
	ToX509AuthenticationDatabaseUserPtrOutputWithContext(ctx context.Context) X509AuthenticationDatabaseUserPtrOutput
}

type x509authenticationDatabaseUserPtrType X509AuthenticationDatabaseUserArgs

func (*x509authenticationDatabaseUserPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**X509AuthenticationDatabaseUser)(nil))
}

func (i *x509authenticationDatabaseUserPtrType) ToX509AuthenticationDatabaseUserPtrOutput() X509AuthenticationDatabaseUserPtrOutput {
	return i.ToX509AuthenticationDatabaseUserPtrOutputWithContext(context.Background())
}

func (i *x509authenticationDatabaseUserPtrType) ToX509AuthenticationDatabaseUserPtrOutputWithContext(ctx context.Context) X509AuthenticationDatabaseUserPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(X509AuthenticationDatabaseUserPtrOutput)
}

// X509AuthenticationDatabaseUserArrayInput is an input type that accepts X509AuthenticationDatabaseUserArray and X509AuthenticationDatabaseUserArrayOutput values.
// You can construct a concrete instance of `X509AuthenticationDatabaseUserArrayInput` via:
//
//          X509AuthenticationDatabaseUserArray{ X509AuthenticationDatabaseUserArgs{...} }
type X509AuthenticationDatabaseUserArrayInput interface {
	pulumi.Input

	ToX509AuthenticationDatabaseUserArrayOutput() X509AuthenticationDatabaseUserArrayOutput
	ToX509AuthenticationDatabaseUserArrayOutputWithContext(context.Context) X509AuthenticationDatabaseUserArrayOutput
}

type X509AuthenticationDatabaseUserArray []X509AuthenticationDatabaseUserInput

func (X509AuthenticationDatabaseUserArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*X509AuthenticationDatabaseUser)(nil))
}

func (i X509AuthenticationDatabaseUserArray) ToX509AuthenticationDatabaseUserArrayOutput() X509AuthenticationDatabaseUserArrayOutput {
	return i.ToX509AuthenticationDatabaseUserArrayOutputWithContext(context.Background())
}

func (i X509AuthenticationDatabaseUserArray) ToX509AuthenticationDatabaseUserArrayOutputWithContext(ctx context.Context) X509AuthenticationDatabaseUserArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(X509AuthenticationDatabaseUserArrayOutput)
}

// X509AuthenticationDatabaseUserMapInput is an input type that accepts X509AuthenticationDatabaseUserMap and X509AuthenticationDatabaseUserMapOutput values.
// You can construct a concrete instance of `X509AuthenticationDatabaseUserMapInput` via:
//
//          X509AuthenticationDatabaseUserMap{ "key": X509AuthenticationDatabaseUserArgs{...} }
type X509AuthenticationDatabaseUserMapInput interface {
	pulumi.Input

	ToX509AuthenticationDatabaseUserMapOutput() X509AuthenticationDatabaseUserMapOutput
	ToX509AuthenticationDatabaseUserMapOutputWithContext(context.Context) X509AuthenticationDatabaseUserMapOutput
}

type X509AuthenticationDatabaseUserMap map[string]X509AuthenticationDatabaseUserInput

func (X509AuthenticationDatabaseUserMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*X509AuthenticationDatabaseUser)(nil))
}

func (i X509AuthenticationDatabaseUserMap) ToX509AuthenticationDatabaseUserMapOutput() X509AuthenticationDatabaseUserMapOutput {
	return i.ToX509AuthenticationDatabaseUserMapOutputWithContext(context.Background())
}

func (i X509AuthenticationDatabaseUserMap) ToX509AuthenticationDatabaseUserMapOutputWithContext(ctx context.Context) X509AuthenticationDatabaseUserMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(X509AuthenticationDatabaseUserMapOutput)
}

type X509AuthenticationDatabaseUserOutput struct {
	*pulumi.OutputState
}

func (X509AuthenticationDatabaseUserOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*X509AuthenticationDatabaseUser)(nil))
}

func (o X509AuthenticationDatabaseUserOutput) ToX509AuthenticationDatabaseUserOutput() X509AuthenticationDatabaseUserOutput {
	return o
}

func (o X509AuthenticationDatabaseUserOutput) ToX509AuthenticationDatabaseUserOutputWithContext(ctx context.Context) X509AuthenticationDatabaseUserOutput {
	return o
}

func (o X509AuthenticationDatabaseUserOutput) ToX509AuthenticationDatabaseUserPtrOutput() X509AuthenticationDatabaseUserPtrOutput {
	return o.ToX509AuthenticationDatabaseUserPtrOutputWithContext(context.Background())
}

func (o X509AuthenticationDatabaseUserOutput) ToX509AuthenticationDatabaseUserPtrOutputWithContext(ctx context.Context) X509AuthenticationDatabaseUserPtrOutput {
	return o.ApplyT(func(v X509AuthenticationDatabaseUser) *X509AuthenticationDatabaseUser {
		return &v
	}).(X509AuthenticationDatabaseUserPtrOutput)
}

type X509AuthenticationDatabaseUserPtrOutput struct {
	*pulumi.OutputState
}

func (X509AuthenticationDatabaseUserPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**X509AuthenticationDatabaseUser)(nil))
}

func (o X509AuthenticationDatabaseUserPtrOutput) ToX509AuthenticationDatabaseUserPtrOutput() X509AuthenticationDatabaseUserPtrOutput {
	return o
}

func (o X509AuthenticationDatabaseUserPtrOutput) ToX509AuthenticationDatabaseUserPtrOutputWithContext(ctx context.Context) X509AuthenticationDatabaseUserPtrOutput {
	return o
}

type X509AuthenticationDatabaseUserArrayOutput struct{ *pulumi.OutputState }

func (X509AuthenticationDatabaseUserArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]X509AuthenticationDatabaseUser)(nil))
}

func (o X509AuthenticationDatabaseUserArrayOutput) ToX509AuthenticationDatabaseUserArrayOutput() X509AuthenticationDatabaseUserArrayOutput {
	return o
}

func (o X509AuthenticationDatabaseUserArrayOutput) ToX509AuthenticationDatabaseUserArrayOutputWithContext(ctx context.Context) X509AuthenticationDatabaseUserArrayOutput {
	return o
}

func (o X509AuthenticationDatabaseUserArrayOutput) Index(i pulumi.IntInput) X509AuthenticationDatabaseUserOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) X509AuthenticationDatabaseUser {
		return vs[0].([]X509AuthenticationDatabaseUser)[vs[1].(int)]
	}).(X509AuthenticationDatabaseUserOutput)
}

type X509AuthenticationDatabaseUserMapOutput struct{ *pulumi.OutputState }

func (X509AuthenticationDatabaseUserMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]X509AuthenticationDatabaseUser)(nil))
}

func (o X509AuthenticationDatabaseUserMapOutput) ToX509AuthenticationDatabaseUserMapOutput() X509AuthenticationDatabaseUserMapOutput {
	return o
}

func (o X509AuthenticationDatabaseUserMapOutput) ToX509AuthenticationDatabaseUserMapOutputWithContext(ctx context.Context) X509AuthenticationDatabaseUserMapOutput {
	return o
}

func (o X509AuthenticationDatabaseUserMapOutput) MapIndex(k pulumi.StringInput) X509AuthenticationDatabaseUserOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) X509AuthenticationDatabaseUser {
		return vs[0].(map[string]X509AuthenticationDatabaseUser)[vs[1].(string)]
	}).(X509AuthenticationDatabaseUserOutput)
}

func init() {
	pulumi.RegisterOutputType(X509AuthenticationDatabaseUserOutput{})
	pulumi.RegisterOutputType(X509AuthenticationDatabaseUserPtrOutput{})
	pulumi.RegisterOutputType(X509AuthenticationDatabaseUserArrayOutput{})
	pulumi.RegisterOutputType(X509AuthenticationDatabaseUserMapOutput{})
}
