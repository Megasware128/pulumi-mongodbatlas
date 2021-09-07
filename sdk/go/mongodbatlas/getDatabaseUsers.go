// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package mongodbatlas

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// `getDatabaseUsers` describe all Database Users. This represents a database user which will be applied to all clusters within the project.
//
// Each user has a set of roles that provide access to the project’s databases. User's roles apply to all the clusters in the project: if two clusters have a `products` database and a user has a role granting `read` access on the products database, the user has that access on both clusters.
//
// > **NOTE:** Groups and projects are synonymous terms. You may find `groupId` in the official documentation.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-mongodbatlas/sdk/v3/go/mongodbatlas"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		testDatabaseUser, err := mongodbatlas.NewDatabaseUser(ctx, "testDatabaseUser", &mongodbatlas.DatabaseUserArgs{
// 			Username:         pulumi.String("test-acc-username"),
// 			Password:         pulumi.String("test-acc-password"),
// 			ProjectId:        pulumi.String("<PROJECT-ID>"),
// 			AuthDatabaseName: pulumi.String("admin"),
// 			Roles: mongodbatlas.DatabaseUserRoleArray{
// 				&mongodbatlas.DatabaseUserRoleArgs{
// 					RoleName:     pulumi.String("readWrite"),
// 					DatabaseName: pulumi.String("admin"),
// 				},
// 				&mongodbatlas.DatabaseUserRoleArgs{
// 					RoleName:     pulumi.String("atlasAdmin"),
// 					DatabaseName: pulumi.String("admin"),
// 				},
// 			},
// 			Labels: mongodbatlas.DatabaseUserLabelArray{
// 				&mongodbatlas.DatabaseUserLabelArgs{
// 					Key:   pulumi.String("key 1"),
// 					Value: pulumi.String("value 1"),
// 				},
// 				&mongodbatlas.DatabaseUserLabelArgs{
// 					Key:   pulumi.String("key 2"),
// 					Value: pulumi.String("value 2"),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetDatabaseUsers(ctx *pulumi.Context, args *GetDatabaseUsersArgs, opts ...pulumi.InvokeOption) (*GetDatabaseUsersResult, error) {
	var rv GetDatabaseUsersResult
	err := ctx.Invoke("mongodbatlas:index/getDatabaseUsers:getDatabaseUsers", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDatabaseUsers.
type GetDatabaseUsersArgs struct {
	// The unique ID for the project to get all database users.
	ProjectId string `pulumi:"projectId"`
}

// A collection of values returned by getDatabaseUsers.
type GetDatabaseUsersResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// ID of the Atlas project the user belongs to.
	ProjectId string `pulumi:"projectId"`
	// A list where each represents a Database user.
	Results []GetDatabaseUsersResultType `pulumi:"results"`
}
