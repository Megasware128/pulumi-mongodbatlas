// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Mongodbatlas
{
    public static class GetCustomDbRoles
    {
        /// <summary>
        /// `mongodbatlas.getCustomDbRoles` describe all Custom DB Roles. This represents a custom db roles.
        /// 
        /// &gt; **NOTE:** Groups and projects are synonymous terms. You may find `groupId` in the official documentation.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Mongodbatlas = Pulumi.Mongodbatlas;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var testDatabaseUser = new Mongodbatlas.DatabaseUser("testDatabaseUser", new()
        ///     {
        ///         Username = "test-acc-username",
        ///         Password = "test-acc-password",
        ///         ProjectId = "&lt;PROJECT-ID&gt;",
        ///         DatabaseName = "admin",
        ///         Roles = new[]
        ///         {
        ///             new Mongodbatlas.Inputs.DatabaseUserRoleArgs
        ///             {
        ///                 RoleName = "readWrite",
        ///                 DatabaseName = "admin",
        ///             },
        ///             new Mongodbatlas.Inputs.DatabaseUserRoleArgs
        ///             {
        ///                 RoleName = "atlasAdmin",
        ///                 DatabaseName = "admin",
        ///             },
        ///         },
        ///     });
        /// 
        ///     var testCustomDbRoles = Mongodbatlas.GetCustomDbRoles.Invoke(new()
        ///     {
        ///         ProjectId = mongodbatlas_custom_db_role.Test.Project_id,
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetCustomDbRolesResult> InvokeAsync(GetCustomDbRolesArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetCustomDbRolesResult>("mongodbatlas:index/getCustomDbRoles:getCustomDbRoles", args ?? new GetCustomDbRolesArgs(), options.WithDefaults());

        /// <summary>
        /// `mongodbatlas.getCustomDbRoles` describe all Custom DB Roles. This represents a custom db roles.
        /// 
        /// &gt; **NOTE:** Groups and projects are synonymous terms. You may find `groupId` in the official documentation.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Mongodbatlas = Pulumi.Mongodbatlas;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var testDatabaseUser = new Mongodbatlas.DatabaseUser("testDatabaseUser", new()
        ///     {
        ///         Username = "test-acc-username",
        ///         Password = "test-acc-password",
        ///         ProjectId = "&lt;PROJECT-ID&gt;",
        ///         DatabaseName = "admin",
        ///         Roles = new[]
        ///         {
        ///             new Mongodbatlas.Inputs.DatabaseUserRoleArgs
        ///             {
        ///                 RoleName = "readWrite",
        ///                 DatabaseName = "admin",
        ///             },
        ///             new Mongodbatlas.Inputs.DatabaseUserRoleArgs
        ///             {
        ///                 RoleName = "atlasAdmin",
        ///                 DatabaseName = "admin",
        ///             },
        ///         },
        ///     });
        /// 
        ///     var testCustomDbRoles = Mongodbatlas.GetCustomDbRoles.Invoke(new()
        ///     {
        ///         ProjectId = mongodbatlas_custom_db_role.Test.Project_id,
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetCustomDbRolesResult> Invoke(GetCustomDbRolesInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetCustomDbRolesResult>("mongodbatlas:index/getCustomDbRoles:getCustomDbRoles", args ?? new GetCustomDbRolesInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetCustomDbRolesArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The unique ID for the project to get all custom db roles.
        /// </summary>
        [Input("projectId", required: true)]
        public string ProjectId { get; set; } = null!;

        public GetCustomDbRolesArgs()
        {
        }
        public static new GetCustomDbRolesArgs Empty => new GetCustomDbRolesArgs();
    }

    public sealed class GetCustomDbRolesInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The unique ID for the project to get all custom db roles.
        /// </summary>
        [Input("projectId", required: true)]
        public Input<string> ProjectId { get; set; } = null!;

        public GetCustomDbRolesInvokeArgs()
        {
        }
        public static new GetCustomDbRolesInvokeArgs Empty => new GetCustomDbRolesInvokeArgs();
    }


    [OutputType]
    public sealed class GetCustomDbRolesResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string ProjectId;
        /// <summary>
        /// A list where each represents a custom db roles.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetCustomDbRolesResultResult> Results;

        [OutputConstructor]
        private GetCustomDbRolesResult(
            string id,

            string projectId,

            ImmutableArray<Outputs.GetCustomDbRolesResultResult> results)
        {
            Id = id;
            ProjectId = projectId;
            Results = results;
        }
    }
}
