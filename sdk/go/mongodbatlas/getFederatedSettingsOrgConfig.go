// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mongodbatlas

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// `FederatedSettingsOrgConfig` provides an Federated Settings Identity Providers datasource. Atlas Cloud Federated Settings Organizational configuration provides federated settings outputs for the configured Organizational configuration.
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
//			orgConnections, err := mongodbatlas.NewFederatedSettingsOrgConfig(ctx, "orgConnections", &mongodbatlas.FederatedSettingsOrgConfigArgs{
//				FederationSettingsId:     pulumi.String("627a9687f7f7f7f774de306f14"),
//				OrgId:                    pulumi.String("627a9683ea7ff7f74de306f14"),
//				DomainRestrictionEnabled: pulumi.Bool(false),
//				DomainAllowLists: pulumi.StringArray{
//					pulumi.String("mydomain.com"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_ = mongodbatlas.LookupFederatedSettingsOrgConfigOutput(ctx, GetFederatedSettingsOrgConfigOutputArgs{
//				FederationSettingsId: orgConnections.ID(),
//				OrgId:                pulumi.String("627a9683ea7ff7f74de306f14"),
//			}, nil)
//			return nil
//		})
//	}
//
// ```
func LookupFederatedSettingsOrgConfig(ctx *pulumi.Context, args *LookupFederatedSettingsOrgConfigArgs, opts ...pulumi.InvokeOption) (*LookupFederatedSettingsOrgConfigResult, error) {
	var rv LookupFederatedSettingsOrgConfigResult
	err := ctx.Invoke("mongodbatlas:index/getFederatedSettingsOrgConfig:getFederatedSettingsOrgConfig", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getFederatedSettingsOrgConfig.
type LookupFederatedSettingsOrgConfigArgs struct {
	// Unique 24-hexadecimal digit string that identifies the federated authentication configuration.
	FederationSettingsId string `pulumi:"federationSettingsId"`
	// Unique 24-hexadecimal digit string that identifies the organization that contains your projects.
	OrgId string `pulumi:"orgId"`
}

// A collection of values returned by getFederatedSettingsOrgConfig.
type LookupFederatedSettingsOrgConfigResult struct {
	// List that contains the approved domains from which organization users can log in.  Note: If the organization uses an identity provider,  `domainAllowList` includes: any SSO domains associated with organization's identity provider and any custom domains associated with the specific organization.
	DomainAllowLists []string `pulumi:"domainAllowLists"`
	// Flag that indicates whether domain restriction is enabled for the connected organization.  User Conflicts returns null when `domainRestrictionEnabled` is false.
	DomainRestrictionEnabled bool `pulumi:"domainRestrictionEnabled"`
	// Unique 24-hexadecimal digit string that identifies the federated authentication configuration.
	FederationSettingsId string `pulumi:"federationSettingsId"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Unique 24-hexadecimal digit string that identifies the federated authentication configuration.
	IdentityProviderId string `pulumi:"identityProviderId"`
	OrgId              string `pulumi:"orgId"`
	// List that contains the default [roles](https://www.mongodb.com/docs/atlas/reference/user-roles/#std-label-organization-roles) granted to users who authenticate through the IdP in a connected organization. If you provide a postAuthRoleGrants field in the request, the array that you provide replaces the current postAuthRoleGrants.
	PostAuthRoleGrants []string                                    `pulumi:"postAuthRoleGrants"`
	RoleMappings       []GetFederatedSettingsOrgConfigRoleMapping  `pulumi:"roleMappings"`
	UserConflicts      []GetFederatedSettingsOrgConfigUserConflict `pulumi:"userConflicts"`
}

func LookupFederatedSettingsOrgConfigOutput(ctx *pulumi.Context, args LookupFederatedSettingsOrgConfigOutputArgs, opts ...pulumi.InvokeOption) LookupFederatedSettingsOrgConfigResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupFederatedSettingsOrgConfigResult, error) {
			args := v.(LookupFederatedSettingsOrgConfigArgs)
			r, err := LookupFederatedSettingsOrgConfig(ctx, &args, opts...)
			var s LookupFederatedSettingsOrgConfigResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupFederatedSettingsOrgConfigResultOutput)
}

// A collection of arguments for invoking getFederatedSettingsOrgConfig.
type LookupFederatedSettingsOrgConfigOutputArgs struct {
	// Unique 24-hexadecimal digit string that identifies the federated authentication configuration.
	FederationSettingsId pulumi.StringInput `pulumi:"federationSettingsId"`
	// Unique 24-hexadecimal digit string that identifies the organization that contains your projects.
	OrgId pulumi.StringInput `pulumi:"orgId"`
}

func (LookupFederatedSettingsOrgConfigOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFederatedSettingsOrgConfigArgs)(nil)).Elem()
}

// A collection of values returned by getFederatedSettingsOrgConfig.
type LookupFederatedSettingsOrgConfigResultOutput struct{ *pulumi.OutputState }

func (LookupFederatedSettingsOrgConfigResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFederatedSettingsOrgConfigResult)(nil)).Elem()
}

func (o LookupFederatedSettingsOrgConfigResultOutput) ToLookupFederatedSettingsOrgConfigResultOutput() LookupFederatedSettingsOrgConfigResultOutput {
	return o
}

func (o LookupFederatedSettingsOrgConfigResultOutput) ToLookupFederatedSettingsOrgConfigResultOutputWithContext(ctx context.Context) LookupFederatedSettingsOrgConfigResultOutput {
	return o
}

// List that contains the approved domains from which organization users can log in.  Note: If the organization uses an identity provider,  `domainAllowList` includes: any SSO domains associated with organization's identity provider and any custom domains associated with the specific organization.
func (o LookupFederatedSettingsOrgConfigResultOutput) DomainAllowLists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupFederatedSettingsOrgConfigResult) []string { return v.DomainAllowLists }).(pulumi.StringArrayOutput)
}

// Flag that indicates whether domain restriction is enabled for the connected organization.  User Conflicts returns null when `domainRestrictionEnabled` is false.
func (o LookupFederatedSettingsOrgConfigResultOutput) DomainRestrictionEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupFederatedSettingsOrgConfigResult) bool { return v.DomainRestrictionEnabled }).(pulumi.BoolOutput)
}

// Unique 24-hexadecimal digit string that identifies the federated authentication configuration.
func (o LookupFederatedSettingsOrgConfigResultOutput) FederationSettingsId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFederatedSettingsOrgConfigResult) string { return v.FederationSettingsId }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupFederatedSettingsOrgConfigResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFederatedSettingsOrgConfigResult) string { return v.Id }).(pulumi.StringOutput)
}

// Unique 24-hexadecimal digit string that identifies the federated authentication configuration.
func (o LookupFederatedSettingsOrgConfigResultOutput) IdentityProviderId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFederatedSettingsOrgConfigResult) string { return v.IdentityProviderId }).(pulumi.StringOutput)
}

func (o LookupFederatedSettingsOrgConfigResultOutput) OrgId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFederatedSettingsOrgConfigResult) string { return v.OrgId }).(pulumi.StringOutput)
}

// List that contains the default [roles](https://www.mongodb.com/docs/atlas/reference/user-roles/#std-label-organization-roles) granted to users who authenticate through the IdP in a connected organization. If you provide a postAuthRoleGrants field in the request, the array that you provide replaces the current postAuthRoleGrants.
func (o LookupFederatedSettingsOrgConfigResultOutput) PostAuthRoleGrants() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupFederatedSettingsOrgConfigResult) []string { return v.PostAuthRoleGrants }).(pulumi.StringArrayOutput)
}

func (o LookupFederatedSettingsOrgConfigResultOutput) RoleMappings() GetFederatedSettingsOrgConfigRoleMappingArrayOutput {
	return o.ApplyT(func(v LookupFederatedSettingsOrgConfigResult) []GetFederatedSettingsOrgConfigRoleMapping {
		return v.RoleMappings
	}).(GetFederatedSettingsOrgConfigRoleMappingArrayOutput)
}

func (o LookupFederatedSettingsOrgConfigResultOutput) UserConflicts() GetFederatedSettingsOrgConfigUserConflictArrayOutput {
	return o.ApplyT(func(v LookupFederatedSettingsOrgConfigResult) []GetFederatedSettingsOrgConfigUserConflict {
		return v.UserConflicts
	}).(GetFederatedSettingsOrgConfigUserConflictArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFederatedSettingsOrgConfigResultOutput{})
}
