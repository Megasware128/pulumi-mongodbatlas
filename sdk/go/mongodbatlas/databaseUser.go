// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mongodbatlas

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Example Usage
// ### S
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-mongodbatlas/sdk/v3/go/mongodbatlas"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := mongodbatlas.NewDatabaseUser(ctx, "test", &mongodbatlas.DatabaseUserArgs{
//				AuthDatabaseName: pulumi.String("admin"),
//				Labels: DatabaseUserLabelArray{
//					&DatabaseUserLabelArgs{
//						Key:   pulumi.String("My Key"),
//						Value: pulumi.String("My Value"),
//					},
//				},
//				Password:  pulumi.String("test-acc-password"),
//				ProjectId: pulumi.String("<PROJECT-ID>"),
//				Roles: DatabaseUserRoleArray{
//					&DatabaseUserRoleArgs{
//						DatabaseName: pulumi.String("dbforApp"),
//						RoleName:     pulumi.String("readWrite"),
//					},
//					&DatabaseUserRoleArgs{
//						DatabaseName: pulumi.String("admin"),
//						RoleName:     pulumi.String("readAnyDatabase"),
//					},
//				},
//				Scopes: DatabaseUserScopeArray{
//					&DatabaseUserScopeArgs{
//						Name: pulumi.String("My cluster name"),
//						Type: pulumi.String("CLUSTER"),
//					},
//					&DatabaseUserScopeArgs{
//						Name: pulumi.String("My second cluster name"),
//						Type: pulumi.String("CLUSTER"),
//					},
//				},
//				Username: pulumi.String("test-acc-username"),
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
// ```go
// package main
//
// import (
//
//	"fmt"
//
//	"github.com/pulumi/pulumi-mongodbatlas/sdk/v3/go/mongodbatlas"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := mongodbatlas.NewDatabaseUser(ctx, "test", &mongodbatlas.DatabaseUserArgs{
//				AuthDatabaseName: pulumi.String(fmt.Sprintf("$external")),
//				Labels: DatabaseUserLabelArray{
//					&DatabaseUserLabelArgs{
//						Key:   pulumi.String(fmt.Sprintf("%vs", "%")),
//						Value: pulumi.String(fmt.Sprintf("%vs", "%")),
//					},
//				},
//				ProjectId: pulumi.String("<PROJECT-ID>"),
//				Roles: DatabaseUserRoleArray{
//					&DatabaseUserRoleArgs{
//						DatabaseName: pulumi.String("admin"),
//						RoleName:     pulumi.String("readAnyDatabase"),
//					},
//				},
//				Scopes: DatabaseUserScopeArray{
//					&DatabaseUserScopeArgs{
//						Name: pulumi.String("My cluster name"),
//						Type: pulumi.String("CLUSTER"),
//					},
//				},
//				Username: pulumi.String("test-acc-username"),
//				X509Type: pulumi.String("MANAGED"),
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
// ```go
// package main
//
// import (
//
//	"fmt"
//
//	"github.com/pulumi/pulumi-mongodbatlas/sdk/v3/go/mongodbatlas"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := mongodbatlas.NewDatabaseUser(ctx, "test", &mongodbatlas.DatabaseUserArgs{
//				Username:         pulumi.Any(aws_iam_role.Test.Arn),
//				ProjectId:        pulumi.String("<PROJECT-ID>"),
//				AuthDatabaseName: pulumi.String(fmt.Sprintf("$external")),
//				AwsIamType:       pulumi.String("ROLE"),
//				Roles: DatabaseUserRoleArray{
//					&DatabaseUserRoleArgs{
//						RoleName:     pulumi.String("readAnyDatabase"),
//						DatabaseName: pulumi.String("admin"),
//					},
//				},
//				Labels: DatabaseUserLabelArray{
//					&DatabaseUserLabelArgs{
//						Key:   pulumi.String(fmt.Sprintf("%vs", "%")),
//						Value: pulumi.String(fmt.Sprintf("%vs", "%")),
//					},
//				},
//				Scopes: DatabaseUserScopeArray{
//					&DatabaseUserScopeArgs{
//						Name: pulumi.String("My cluster name"),
//						Type: pulumi.String("CLUSTER"),
//					},
//				},
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
// Database users can be imported using project ID and username, in the format `project_id`-`username`-`auth_database_name`, e.g.
//
// ```sh
//
//	$ pulumi import mongodbatlas:index/databaseUser:DatabaseUser my_user 1112222b3bf99403840e8934-my_user-admin
//
// ```
type DatabaseUser struct {
	pulumi.CustomResourceState

	// Database against which Atlas authenticates the user. A user must provide both a username and authentication database to log into MongoDB.
	// Accepted values include:
	AuthDatabaseName pulumi.StringPtrOutput `pulumi:"authDatabaseName"`
	// If this value is set, the new database user authenticates with AWS IAM credentials. If no value is given, Atlas uses the default value of NONE. The accepted types are:
	AwsIamType pulumi.StringPtrOutput `pulumi:"awsIamType"`
	// Database on which the user has the specified role. A role on the `admin` database can include privileges that apply to the other databases.
	//
	// Deprecated: use auth_database_name instead
	DatabaseName pulumi.StringPtrOutput       `pulumi:"databaseName"`
	Labels       DatabaseUserLabelArrayOutput `pulumi:"labels"`
	// Method by which the provided `username` is authenticated. If no value is given, Atlas uses the default value of `NONE`.
	LdapAuthType pulumi.StringPtrOutput `pulumi:"ldapAuthType"`
	Password     pulumi.StringPtrOutput `pulumi:"password"`
	// The unique ID for the project to create the database user.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// List of user’s roles and the databases / collections on which the roles apply. A role allows the user to perform particular actions on the specified database. A role on the admin database can include privileges that apply to the other databases as well. See Roles below for more details.
	Roles  DatabaseUserRoleArrayOutput  `pulumi:"roles"`
	Scopes DatabaseUserScopeArrayOutput `pulumi:"scopes"`
	// Username for authenticating to MongoDB. USER_ARN or ROLE_ARN if `awsIamType` is USER or ROLE.
	Username pulumi.StringOutput `pulumi:"username"`
	// X.509 method by which the provided username is authenticated. If no value is given, Atlas uses the default value of NONE. The accepted types are:
	X509Type pulumi.StringPtrOutput `pulumi:"x509Type"`
}

// NewDatabaseUser registers a new resource with the given unique name, arguments, and options.
func NewDatabaseUser(ctx *pulumi.Context,
	name string, args *DatabaseUserArgs, opts ...pulumi.ResourceOption) (*DatabaseUser, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ProjectId == nil {
		return nil, errors.New("invalid value for required argument 'ProjectId'")
	}
	if args.Roles == nil {
		return nil, errors.New("invalid value for required argument 'Roles'")
	}
	if args.Username == nil {
		return nil, errors.New("invalid value for required argument 'Username'")
	}
	var resource DatabaseUser
	err := ctx.RegisterResource("mongodbatlas:index/databaseUser:DatabaseUser", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDatabaseUser gets an existing DatabaseUser resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDatabaseUser(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DatabaseUserState, opts ...pulumi.ResourceOption) (*DatabaseUser, error) {
	var resource DatabaseUser
	err := ctx.ReadResource("mongodbatlas:index/databaseUser:DatabaseUser", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DatabaseUser resources.
type databaseUserState struct {
	// Database against which Atlas authenticates the user. A user must provide both a username and authentication database to log into MongoDB.
	// Accepted values include:
	AuthDatabaseName *string `pulumi:"authDatabaseName"`
	// If this value is set, the new database user authenticates with AWS IAM credentials. If no value is given, Atlas uses the default value of NONE. The accepted types are:
	AwsIamType *string `pulumi:"awsIamType"`
	// Database on which the user has the specified role. A role on the `admin` database can include privileges that apply to the other databases.
	//
	// Deprecated: use auth_database_name instead
	DatabaseName *string             `pulumi:"databaseName"`
	Labels       []DatabaseUserLabel `pulumi:"labels"`
	// Method by which the provided `username` is authenticated. If no value is given, Atlas uses the default value of `NONE`.
	LdapAuthType *string `pulumi:"ldapAuthType"`
	Password     *string `pulumi:"password"`
	// The unique ID for the project to create the database user.
	ProjectId *string `pulumi:"projectId"`
	// List of user’s roles and the databases / collections on which the roles apply. A role allows the user to perform particular actions on the specified database. A role on the admin database can include privileges that apply to the other databases as well. See Roles below for more details.
	Roles  []DatabaseUserRole  `pulumi:"roles"`
	Scopes []DatabaseUserScope `pulumi:"scopes"`
	// Username for authenticating to MongoDB. USER_ARN or ROLE_ARN if `awsIamType` is USER or ROLE.
	Username *string `pulumi:"username"`
	// X.509 method by which the provided username is authenticated. If no value is given, Atlas uses the default value of NONE. The accepted types are:
	X509Type *string `pulumi:"x509Type"`
}

type DatabaseUserState struct {
	// Database against which Atlas authenticates the user. A user must provide both a username and authentication database to log into MongoDB.
	// Accepted values include:
	AuthDatabaseName pulumi.StringPtrInput
	// If this value is set, the new database user authenticates with AWS IAM credentials. If no value is given, Atlas uses the default value of NONE. The accepted types are:
	AwsIamType pulumi.StringPtrInput
	// Database on which the user has the specified role. A role on the `admin` database can include privileges that apply to the other databases.
	//
	// Deprecated: use auth_database_name instead
	DatabaseName pulumi.StringPtrInput
	Labels       DatabaseUserLabelArrayInput
	// Method by which the provided `username` is authenticated. If no value is given, Atlas uses the default value of `NONE`.
	LdapAuthType pulumi.StringPtrInput
	Password     pulumi.StringPtrInput
	// The unique ID for the project to create the database user.
	ProjectId pulumi.StringPtrInput
	// List of user’s roles and the databases / collections on which the roles apply. A role allows the user to perform particular actions on the specified database. A role on the admin database can include privileges that apply to the other databases as well. See Roles below for more details.
	Roles  DatabaseUserRoleArrayInput
	Scopes DatabaseUserScopeArrayInput
	// Username for authenticating to MongoDB. USER_ARN or ROLE_ARN if `awsIamType` is USER or ROLE.
	Username pulumi.StringPtrInput
	// X.509 method by which the provided username is authenticated. If no value is given, Atlas uses the default value of NONE. The accepted types are:
	X509Type pulumi.StringPtrInput
}

func (DatabaseUserState) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseUserState)(nil)).Elem()
}

type databaseUserArgs struct {
	// Database against which Atlas authenticates the user. A user must provide both a username and authentication database to log into MongoDB.
	// Accepted values include:
	AuthDatabaseName *string `pulumi:"authDatabaseName"`
	// If this value is set, the new database user authenticates with AWS IAM credentials. If no value is given, Atlas uses the default value of NONE. The accepted types are:
	AwsIamType *string `pulumi:"awsIamType"`
	// Database on which the user has the specified role. A role on the `admin` database can include privileges that apply to the other databases.
	//
	// Deprecated: use auth_database_name instead
	DatabaseName *string             `pulumi:"databaseName"`
	Labels       []DatabaseUserLabel `pulumi:"labels"`
	// Method by which the provided `username` is authenticated. If no value is given, Atlas uses the default value of `NONE`.
	LdapAuthType *string `pulumi:"ldapAuthType"`
	Password     *string `pulumi:"password"`
	// The unique ID for the project to create the database user.
	ProjectId string `pulumi:"projectId"`
	// List of user’s roles and the databases / collections on which the roles apply. A role allows the user to perform particular actions on the specified database. A role on the admin database can include privileges that apply to the other databases as well. See Roles below for more details.
	Roles  []DatabaseUserRole  `pulumi:"roles"`
	Scopes []DatabaseUserScope `pulumi:"scopes"`
	// Username for authenticating to MongoDB. USER_ARN or ROLE_ARN if `awsIamType` is USER or ROLE.
	Username string `pulumi:"username"`
	// X.509 method by which the provided username is authenticated. If no value is given, Atlas uses the default value of NONE. The accepted types are:
	X509Type *string `pulumi:"x509Type"`
}

// The set of arguments for constructing a DatabaseUser resource.
type DatabaseUserArgs struct {
	// Database against which Atlas authenticates the user. A user must provide both a username and authentication database to log into MongoDB.
	// Accepted values include:
	AuthDatabaseName pulumi.StringPtrInput
	// If this value is set, the new database user authenticates with AWS IAM credentials. If no value is given, Atlas uses the default value of NONE. The accepted types are:
	AwsIamType pulumi.StringPtrInput
	// Database on which the user has the specified role. A role on the `admin` database can include privileges that apply to the other databases.
	//
	// Deprecated: use auth_database_name instead
	DatabaseName pulumi.StringPtrInput
	Labels       DatabaseUserLabelArrayInput
	// Method by which the provided `username` is authenticated. If no value is given, Atlas uses the default value of `NONE`.
	LdapAuthType pulumi.StringPtrInput
	Password     pulumi.StringPtrInput
	// The unique ID for the project to create the database user.
	ProjectId pulumi.StringInput
	// List of user’s roles and the databases / collections on which the roles apply. A role allows the user to perform particular actions on the specified database. A role on the admin database can include privileges that apply to the other databases as well. See Roles below for more details.
	Roles  DatabaseUserRoleArrayInput
	Scopes DatabaseUserScopeArrayInput
	// Username for authenticating to MongoDB. USER_ARN or ROLE_ARN if `awsIamType` is USER or ROLE.
	Username pulumi.StringInput
	// X.509 method by which the provided username is authenticated. If no value is given, Atlas uses the default value of NONE. The accepted types are:
	X509Type pulumi.StringPtrInput
}

func (DatabaseUserArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseUserArgs)(nil)).Elem()
}

type DatabaseUserInput interface {
	pulumi.Input

	ToDatabaseUserOutput() DatabaseUserOutput
	ToDatabaseUserOutputWithContext(ctx context.Context) DatabaseUserOutput
}

func (*DatabaseUser) ElementType() reflect.Type {
	return reflect.TypeOf((**DatabaseUser)(nil)).Elem()
}

func (i *DatabaseUser) ToDatabaseUserOutput() DatabaseUserOutput {
	return i.ToDatabaseUserOutputWithContext(context.Background())
}

func (i *DatabaseUser) ToDatabaseUserOutputWithContext(ctx context.Context) DatabaseUserOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseUserOutput)
}

// DatabaseUserArrayInput is an input type that accepts DatabaseUserArray and DatabaseUserArrayOutput values.
// You can construct a concrete instance of `DatabaseUserArrayInput` via:
//
//	DatabaseUserArray{ DatabaseUserArgs{...} }
type DatabaseUserArrayInput interface {
	pulumi.Input

	ToDatabaseUserArrayOutput() DatabaseUserArrayOutput
	ToDatabaseUserArrayOutputWithContext(context.Context) DatabaseUserArrayOutput
}

type DatabaseUserArray []DatabaseUserInput

func (DatabaseUserArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DatabaseUser)(nil)).Elem()
}

func (i DatabaseUserArray) ToDatabaseUserArrayOutput() DatabaseUserArrayOutput {
	return i.ToDatabaseUserArrayOutputWithContext(context.Background())
}

func (i DatabaseUserArray) ToDatabaseUserArrayOutputWithContext(ctx context.Context) DatabaseUserArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseUserArrayOutput)
}

// DatabaseUserMapInput is an input type that accepts DatabaseUserMap and DatabaseUserMapOutput values.
// You can construct a concrete instance of `DatabaseUserMapInput` via:
//
//	DatabaseUserMap{ "key": DatabaseUserArgs{...} }
type DatabaseUserMapInput interface {
	pulumi.Input

	ToDatabaseUserMapOutput() DatabaseUserMapOutput
	ToDatabaseUserMapOutputWithContext(context.Context) DatabaseUserMapOutput
}

type DatabaseUserMap map[string]DatabaseUserInput

func (DatabaseUserMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DatabaseUser)(nil)).Elem()
}

func (i DatabaseUserMap) ToDatabaseUserMapOutput() DatabaseUserMapOutput {
	return i.ToDatabaseUserMapOutputWithContext(context.Background())
}

func (i DatabaseUserMap) ToDatabaseUserMapOutputWithContext(ctx context.Context) DatabaseUserMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseUserMapOutput)
}

type DatabaseUserOutput struct{ *pulumi.OutputState }

func (DatabaseUserOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DatabaseUser)(nil)).Elem()
}

func (o DatabaseUserOutput) ToDatabaseUserOutput() DatabaseUserOutput {
	return o
}

func (o DatabaseUserOutput) ToDatabaseUserOutputWithContext(ctx context.Context) DatabaseUserOutput {
	return o
}

// Database against which Atlas authenticates the user. A user must provide both a username and authentication database to log into MongoDB.
// Accepted values include:
func (o DatabaseUserOutput) AuthDatabaseName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatabaseUser) pulumi.StringPtrOutput { return v.AuthDatabaseName }).(pulumi.StringPtrOutput)
}

// If this value is set, the new database user authenticates with AWS IAM credentials. If no value is given, Atlas uses the default value of NONE. The accepted types are:
func (o DatabaseUserOutput) AwsIamType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatabaseUser) pulumi.StringPtrOutput { return v.AwsIamType }).(pulumi.StringPtrOutput)
}

// Database on which the user has the specified role. A role on the `admin` database can include privileges that apply to the other databases.
//
// Deprecated: use auth_database_name instead
func (o DatabaseUserOutput) DatabaseName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatabaseUser) pulumi.StringPtrOutput { return v.DatabaseName }).(pulumi.StringPtrOutput)
}

func (o DatabaseUserOutput) Labels() DatabaseUserLabelArrayOutput {
	return o.ApplyT(func(v *DatabaseUser) DatabaseUserLabelArrayOutput { return v.Labels }).(DatabaseUserLabelArrayOutput)
}

// Method by which the provided `username` is authenticated. If no value is given, Atlas uses the default value of `NONE`.
func (o DatabaseUserOutput) LdapAuthType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatabaseUser) pulumi.StringPtrOutput { return v.LdapAuthType }).(pulumi.StringPtrOutput)
}

func (o DatabaseUserOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatabaseUser) pulumi.StringPtrOutput { return v.Password }).(pulumi.StringPtrOutput)
}

// The unique ID for the project to create the database user.
func (o DatabaseUserOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseUser) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// List of user’s roles and the databases / collections on which the roles apply. A role allows the user to perform particular actions on the specified database. A role on the admin database can include privileges that apply to the other databases as well. See Roles below for more details.
func (o DatabaseUserOutput) Roles() DatabaseUserRoleArrayOutput {
	return o.ApplyT(func(v *DatabaseUser) DatabaseUserRoleArrayOutput { return v.Roles }).(DatabaseUserRoleArrayOutput)
}

func (o DatabaseUserOutput) Scopes() DatabaseUserScopeArrayOutput {
	return o.ApplyT(func(v *DatabaseUser) DatabaseUserScopeArrayOutput { return v.Scopes }).(DatabaseUserScopeArrayOutput)
}

// Username for authenticating to MongoDB. USER_ARN or ROLE_ARN if `awsIamType` is USER or ROLE.
func (o DatabaseUserOutput) Username() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseUser) pulumi.StringOutput { return v.Username }).(pulumi.StringOutput)
}

// X.509 method by which the provided username is authenticated. If no value is given, Atlas uses the default value of NONE. The accepted types are:
func (o DatabaseUserOutput) X509Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatabaseUser) pulumi.StringPtrOutput { return v.X509Type }).(pulumi.StringPtrOutput)
}

type DatabaseUserArrayOutput struct{ *pulumi.OutputState }

func (DatabaseUserArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DatabaseUser)(nil)).Elem()
}

func (o DatabaseUserArrayOutput) ToDatabaseUserArrayOutput() DatabaseUserArrayOutput {
	return o
}

func (o DatabaseUserArrayOutput) ToDatabaseUserArrayOutputWithContext(ctx context.Context) DatabaseUserArrayOutput {
	return o
}

func (o DatabaseUserArrayOutput) Index(i pulumi.IntInput) DatabaseUserOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DatabaseUser {
		return vs[0].([]*DatabaseUser)[vs[1].(int)]
	}).(DatabaseUserOutput)
}

type DatabaseUserMapOutput struct{ *pulumi.OutputState }

func (DatabaseUserMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DatabaseUser)(nil)).Elem()
}

func (o DatabaseUserMapOutput) ToDatabaseUserMapOutput() DatabaseUserMapOutput {
	return o
}

func (o DatabaseUserMapOutput) ToDatabaseUserMapOutputWithContext(ctx context.Context) DatabaseUserMapOutput {
	return o
}

func (o DatabaseUserMapOutput) MapIndex(k pulumi.StringInput) DatabaseUserOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DatabaseUser {
		return vs[0].(map[string]*DatabaseUser)[vs[1].(string)]
	}).(DatabaseUserOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DatabaseUserInput)(nil)).Elem(), &DatabaseUser{})
	pulumi.RegisterInputType(reflect.TypeOf((*DatabaseUserArrayInput)(nil)).Elem(), DatabaseUserArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DatabaseUserMapInput)(nil)).Elem(), DatabaseUserMap{})
	pulumi.RegisterOutputType(DatabaseUserOutput{})
	pulumi.RegisterOutputType(DatabaseUserArrayOutput{})
	pulumi.RegisterOutputType(DatabaseUserMapOutput{})
}
