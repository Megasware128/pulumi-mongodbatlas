// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Mongodbatlas.Inputs
{

    public sealed class FederatedSettingsOrgRoleMappingRoleAssignmentArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Unique identifier of the project to which you want the role mapping to apply.
        /// </summary>
        [Input("groupId")]
        public Input<string>? GroupId { get; set; }

        /// <summary>
        /// Unique 24-hexadecimal digit string that identifies the organization that contains your projects.
        /// </summary>
        [Input("orgId")]
        public Input<string>? OrgId { get; set; }

        [Input("roles")]
        private InputList<string>? _roles;

        /// <summary>
        /// Specifies the Roles that are attached to the Role Mapping.
        /// </summary>
        public InputList<string> Roles
        {
            get => _roles ?? (_roles = new InputList<string>());
            set => _roles = value;
        }

        public FederatedSettingsOrgRoleMappingRoleAssignmentArgs()
        {
        }
        public static new FederatedSettingsOrgRoleMappingRoleAssignmentArgs Empty => new FederatedSettingsOrgRoleMappingRoleAssignmentArgs();
    }
}
