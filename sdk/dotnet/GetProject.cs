// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Mongodbatlas
{
    public static class GetProject
    {
        /// <summary>
        /// `mongodbatlas.Project` describes a MongoDB Atlas Project. This represents a project that has been created.
        /// 
        /// &gt; **NOTE:** Groups and projects are synonymous terms. You may find group_id in the official documentation.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetProjectResult> InvokeAsync(GetProjectArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetProjectResult>("mongodbatlas:index/getProject:getProject", args ?? new GetProjectArgs(), options.WithDefaults());

        /// <summary>
        /// `mongodbatlas.Project` describes a MongoDB Atlas Project. This represents a project that has been created.
        /// 
        /// &gt; **NOTE:** Groups and projects are synonymous terms. You may find group_id in the official documentation.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetProjectResult> Invoke(GetProjectInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetProjectResult>("mongodbatlas:index/getProject:getProject", args ?? new GetProjectInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetProjectArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The unique ID for the project.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        /// <summary>
        /// The unique ID for the project.
        /// </summary>
        [Input("projectId")]
        public string? ProjectId { get; set; }

        public GetProjectArgs()
        {
        }
    }

    public sealed class GetProjectInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The unique ID for the project.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The unique ID for the project.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        public GetProjectInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetProjectResult
    {
        public readonly ImmutableArray<Outputs.GetProjectApiKeyResult> ApiKeys;
        public readonly int ClusterCount;
        public readonly string Created;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The name of the project you want to create. (Cannot be changed via this Provider after creation.)
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// The ID of the organization you want to create the project within.
        /// *`cluster_count` - The number of Atlas clusters deployed in the project.
        /// *`created` - The ISO-8601-formatted timestamp of when Atlas created the project.
        /// * `teams.#.team_id` - The unique identifier of the team you want to associate with the project. The team and project must share the same parent organization.
        /// * `teams.#.role_names` - Each string in the array represents a project role assigned to the team. Every user associated with the team inherits these roles.
        /// The following are valid roles:
        /// * `GROUP_OWNER`
        /// * `GROUP_READ_ONLY`
        /// * `GROUP_DATA_ACCESS_ADMIN`
        /// * `GROUP_DATA_ACCESS_READ_WRITE`
        /// * `GROUP_DATA_ACCESS_READ_ONLY`
        /// * `GROUP_CLUSTER_MANAGER`
        /// * `api_keys.#.api_key_id` - The unique identifier of the programmatic API key you want to associate with the project. The programmatic API key and project must share the same parent organization.
        /// * `api_keys.#.role_names` - Each string in the array represents a project role assigned to the programmatic API key.
        /// The following are valid roles:
        /// * `GROUP_OWNER`
        /// * `GROUP_READ_ONLY`
        /// * `GROUP_DATA_ACCESS_ADMIN`
        /// * `GROUP_DATA_ACCESS_READ_WRITE`
        /// * `GROUP_DATA_ACCESS_READ_ONLY`
        /// * `GROUP_CLUSTER_MANAGER`
        /// </summary>
        public readonly string OrgId;
        public readonly string? ProjectId;
        public readonly ImmutableArray<Outputs.GetProjectTeamResult> Teams;

        [OutputConstructor]
        private GetProjectResult(
            ImmutableArray<Outputs.GetProjectApiKeyResult> apiKeys,

            int clusterCount,

            string created,

            string id,

            string? name,

            string orgId,

            string? projectId,

            ImmutableArray<Outputs.GetProjectTeamResult> teams)
        {
            ApiKeys = apiKeys;
            ClusterCount = clusterCount;
            Created = created;
            Id = id;
            Name = name;
            OrgId = orgId;
            ProjectId = projectId;
            Teams = teams;
        }
    }
}
