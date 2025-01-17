// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mongodbatlas

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// `FederatedSettingsOrgRoleMapping` provides an Federated Settings Org Role Mapping datasource. Atlas Cloud Federated Settings Org Role Mapping provides federated settings outputs for the configured Org Role Mapping.
//
// ## Example Usage
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
//			orgGroupRoleMappingImport, err := mongodbatlas.NewFederatedSettingsOrgRoleMapping(ctx, "orgGroupRoleMappingImport", &mongodbatlas.FederatedSettingsOrgRoleMappingArgs{
//				FederationSettingsId: pulumi.Any(data.Mongodbatlas_federated_settings.Federated_settings.Id),
//				OrgId:                pulumi.String("627a9683e7f7f7ff7fe306f14"),
//				ExternalGroupName:    pulumi.String("myGrouptest"),
//				RoleAssignments: FederatedSettingsOrgRoleMappingRoleAssignmentArray{
//					&FederatedSettingsOrgRoleMappingRoleAssignmentArgs{
//						OrgId: pulumi.String("627a9683e7f7f7ff7fe306f14"),
//						Roles: pulumi.StringArray{
//							pulumi.String("ORG_MEMBER"),
//							pulumi.String("ORG_GROUP_CREATOR"),
//							pulumi.String("ORG_BILLING_ADMIN"),
//						},
//					},
//					&FederatedSettingsOrgRoleMappingRoleAssignmentArgs{
//						GroupId: pulumi.String("628aa20db7f7f7f98b81b8"),
//						Roles: pulumi.StringArray{
//							pulumi.String("GROUP_OWNER"),
//							pulumi.String("GROUP_DATA_ACCESS_ADMIN"),
//							pulumi.String("GROUP_SEARCH_INDEX_EDITOR"),
//							pulumi.String("GROUP_DATA_ACCESS_READ_ONLY"),
//						},
//					},
//					&FederatedSettingsOrgRoleMappingRoleAssignmentArgs{
//						GroupId: pulumi.String("62b477f7f7f7f5e741489c"),
//						Roles: pulumi.StringArray{
//							pulumi.String("GROUP_OWNER"),
//							pulumi.String("GROUP_DATA_ACCESS_ADMIN"),
//							pulumi.String("GROUP_SEARCH_INDEX_EDITOR"),
//							pulumi.String("GROUP_DATA_ACCESS_READ_ONLY"),
//							pulumi.String("GROUP_DATA_ACCESS_READ_WRITE"),
//						},
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_ = mongodbatlas.LookupFederatedSettingsOrgRoleMappingOutput(ctx, GetFederatedSettingsOrgRoleMappingOutputArgs{
//				FederationSettingsId: orgGroupRoleMappingImport.ID(),
//				OrgId:                pulumi.String("627a9683e7f7f7ff7fe306f14"),
//				RoleMappingId:        pulumi.String("627a9673e7f7f7ff7fe306f14"),
//			}, nil)
//			return nil
//		})
//	}
//
// ```
func LookupFederatedSettingsOrgRoleMapping(ctx *pulumi.Context, args *LookupFederatedSettingsOrgRoleMappingArgs, opts ...pulumi.InvokeOption) (*LookupFederatedSettingsOrgRoleMappingResult, error) {
	var rv LookupFederatedSettingsOrgRoleMappingResult
	err := ctx.Invoke("mongodbatlas:index/getFederatedSettingsOrgRoleMapping:getFederatedSettingsOrgRoleMapping", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getFederatedSettingsOrgRoleMapping.
type LookupFederatedSettingsOrgRoleMappingArgs struct {
	// Unique 24-hexadecimal digit string that identifies the federated authentication configuration.
	FederationSettingsId string `pulumi:"federationSettingsId"`
	// Unique 24-hexadecimal digit string that identifies the organization that contains your projects.
	OrgId         string `pulumi:"orgId"`
	RoleMappingId string `pulumi:"roleMappingId"`
}

// A collection of values returned by getFederatedSettingsOrgRoleMapping.
type LookupFederatedSettingsOrgRoleMappingResult struct {
	// Unique human-readable label that identifies the identity provider group to which this role mapping applies.
	ExternalGroupName    string `pulumi:"externalGroupName"`
	FederationSettingsId string `pulumi:"federationSettingsId"`
	// Unique 24-hexadecimal digit string that identifies this role mapping.
	Id    string `pulumi:"id"`
	OrgId string `pulumi:"orgId"`
	// Atlas roles and the unique identifiers of the groups and organizations associated with each role.
	RoleAssignments []GetFederatedSettingsOrgRoleMappingRoleAssignment `pulumi:"roleAssignments"`
	RoleMappingId   string                                             `pulumi:"roleMappingId"`
}

func LookupFederatedSettingsOrgRoleMappingOutput(ctx *pulumi.Context, args LookupFederatedSettingsOrgRoleMappingOutputArgs, opts ...pulumi.InvokeOption) LookupFederatedSettingsOrgRoleMappingResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupFederatedSettingsOrgRoleMappingResult, error) {
			args := v.(LookupFederatedSettingsOrgRoleMappingArgs)
			r, err := LookupFederatedSettingsOrgRoleMapping(ctx, &args, opts...)
			var s LookupFederatedSettingsOrgRoleMappingResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupFederatedSettingsOrgRoleMappingResultOutput)
}

// A collection of arguments for invoking getFederatedSettingsOrgRoleMapping.
type LookupFederatedSettingsOrgRoleMappingOutputArgs struct {
	// Unique 24-hexadecimal digit string that identifies the federated authentication configuration.
	FederationSettingsId pulumi.StringInput `pulumi:"federationSettingsId"`
	// Unique 24-hexadecimal digit string that identifies the organization that contains your projects.
	OrgId         pulumi.StringInput `pulumi:"orgId"`
	RoleMappingId pulumi.StringInput `pulumi:"roleMappingId"`
}

func (LookupFederatedSettingsOrgRoleMappingOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFederatedSettingsOrgRoleMappingArgs)(nil)).Elem()
}

// A collection of values returned by getFederatedSettingsOrgRoleMapping.
type LookupFederatedSettingsOrgRoleMappingResultOutput struct{ *pulumi.OutputState }

func (LookupFederatedSettingsOrgRoleMappingResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFederatedSettingsOrgRoleMappingResult)(nil)).Elem()
}

func (o LookupFederatedSettingsOrgRoleMappingResultOutput) ToLookupFederatedSettingsOrgRoleMappingResultOutput() LookupFederatedSettingsOrgRoleMappingResultOutput {
	return o
}

func (o LookupFederatedSettingsOrgRoleMappingResultOutput) ToLookupFederatedSettingsOrgRoleMappingResultOutputWithContext(ctx context.Context) LookupFederatedSettingsOrgRoleMappingResultOutput {
	return o
}

// Unique human-readable label that identifies the identity provider group to which this role mapping applies.
func (o LookupFederatedSettingsOrgRoleMappingResultOutput) ExternalGroupName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFederatedSettingsOrgRoleMappingResult) string { return v.ExternalGroupName }).(pulumi.StringOutput)
}

func (o LookupFederatedSettingsOrgRoleMappingResultOutput) FederationSettingsId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFederatedSettingsOrgRoleMappingResult) string { return v.FederationSettingsId }).(pulumi.StringOutput)
}

// Unique 24-hexadecimal digit string that identifies this role mapping.
func (o LookupFederatedSettingsOrgRoleMappingResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFederatedSettingsOrgRoleMappingResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupFederatedSettingsOrgRoleMappingResultOutput) OrgId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFederatedSettingsOrgRoleMappingResult) string { return v.OrgId }).(pulumi.StringOutput)
}

// Atlas roles and the unique identifiers of the groups and organizations associated with each role.
func (o LookupFederatedSettingsOrgRoleMappingResultOutput) RoleAssignments() GetFederatedSettingsOrgRoleMappingRoleAssignmentArrayOutput {
	return o.ApplyT(func(v LookupFederatedSettingsOrgRoleMappingResult) []GetFederatedSettingsOrgRoleMappingRoleAssignment {
		return v.RoleAssignments
	}).(GetFederatedSettingsOrgRoleMappingRoleAssignmentArrayOutput)
}

func (o LookupFederatedSettingsOrgRoleMappingResultOutput) RoleMappingId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFederatedSettingsOrgRoleMappingResult) string { return v.RoleMappingId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFederatedSettingsOrgRoleMappingResultOutput{})
}
