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
    public sealed class GetFederatedSettingsIdentityProvidersResultAssociatedOrgResult
    {
        /// <summary>
        /// List that contains the approved domains from which organization users can log in.
        /// </summary>
        public readonly ImmutableArray<string> DomainAllowLists;
        /// <summary>
        /// Flag that indicates whether domain restriction is enabled for the connected organization.
        /// </summary>
        public readonly bool DomainRestrictionEnabled;
        /// <summary>
        /// Unique 24-hexadecimal digit string that identifies the federated authentication configuration.
        /// </summary>
        public readonly string IdentityProviderId;
        /// <summary>
        /// Unique 24-hexadecimal digit string that identifies the organization that contains your projects.
        /// </summary>
        public readonly string OrgId;
        /// <summary>
        /// List that contains the default roles granted to users who authenticate through the IdP in a connected organization. If you provide a postAuthRoleGrants field in the request, the array that you provide replaces the current postAuthRoleGrants.
        /// </summary>
        public readonly ImmutableArray<string> PostAuthRoleGrants;
        public readonly ImmutableArray<Outputs.GetFederatedSettingsIdentityProvidersResultAssociatedOrgRoleMappingResult> RoleMappings;
        public readonly ImmutableArray<Outputs.GetFederatedSettingsIdentityProvidersResultAssociatedOrgUserConflictResult> UserConflicts;

        [OutputConstructor]
        private GetFederatedSettingsIdentityProvidersResultAssociatedOrgResult(
            ImmutableArray<string> domainAllowLists,

            bool domainRestrictionEnabled,

            string identityProviderId,

            string orgId,

            ImmutableArray<string> postAuthRoleGrants,

            ImmutableArray<Outputs.GetFederatedSettingsIdentityProvidersResultAssociatedOrgRoleMappingResult> roleMappings,

            ImmutableArray<Outputs.GetFederatedSettingsIdentityProvidersResultAssociatedOrgUserConflictResult> userConflicts)
        {
            DomainAllowLists = domainAllowLists;
            DomainRestrictionEnabled = domainRestrictionEnabled;
            IdentityProviderId = identityProviderId;
            OrgId = orgId;
            PostAuthRoleGrants = postAuthRoleGrants;
            RoleMappings = roleMappings;
            UserConflicts = userConflicts;
        }
    }
}
