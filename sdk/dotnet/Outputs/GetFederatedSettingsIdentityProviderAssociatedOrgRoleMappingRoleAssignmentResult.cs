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
    public sealed class GetFederatedSettingsIdentityProviderAssociatedOrgRoleMappingRoleAssignmentResult
    {
        /// <summary>
        /// Unique identifier of the project to which you want the role mapping to apply.
        /// </summary>
        public readonly string GroupId;
        /// <summary>
        /// Unique 24-hexadecimal digit string that identifies the organization that contains your projects.
        /// </summary>
        public readonly string OrgId;
        /// <summary>
        /// Specifies the Role that is attached to the Role Mapping.
        /// </summary>
        public readonly string Role;

        [OutputConstructor]
        private GetFederatedSettingsIdentityProviderAssociatedOrgRoleMappingRoleAssignmentResult(
            string groupId,

            string orgId,

            string role)
        {
            GroupId = groupId;
            OrgId = orgId;
            Role = role;
        }
    }
}
