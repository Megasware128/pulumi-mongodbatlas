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
    public sealed class GetFederatedSettingsOrgConfigRoleMappingResult
    {
        /// <summary>
        /// Unique human-readable label that identifies the identity provider group to which this role mapping applies.
        /// </summary>
        public readonly string ExternalGroupName;
        /// <summary>
        /// Unique 24-hexadecimal digit string that identifies this role mapping.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Atlas roles and the unique identifiers of the groups and organizations associated with each role.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetFederatedSettingsOrgConfigRoleMappingRoleAssignmentResult> RoleAssignments;

        [OutputConstructor]
        private GetFederatedSettingsOrgConfigRoleMappingResult(
            string externalGroupName,

            string id,

            ImmutableArray<Outputs.GetFederatedSettingsOrgConfigRoleMappingRoleAssignmentResult> roleAssignments)
        {
            ExternalGroupName = externalGroupName;
            Id = id;
            RoleAssignments = roleAssignments;
        }
    }
}
