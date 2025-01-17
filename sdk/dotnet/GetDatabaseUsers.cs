// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Mongodbatlas
{
    public static class GetDatabaseUsers
    {
        /// <summary>
        /// `mongodbatlas.getDatabaseUsers` describe all Database Users. This represents a database user which will be applied to all clusters within the project.
        /// 
        /// Each user has a set of roles that provide access to the project’s databases. User's roles apply to all the clusters in the project: if two clusters have a `products` database and a user has a role granting `read` access on the products database, the user has that access on both clusters.
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
        ///         AuthDatabaseName = "admin",
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
        ///         Labels = new[]
        ///         {
        ///             new Mongodbatlas.Inputs.DatabaseUserLabelArgs
        ///             {
        ///                 Key = "key 1",
        ///                 Value = "value 1",
        ///             },
        ///             new Mongodbatlas.Inputs.DatabaseUserLabelArgs
        ///             {
        ///                 Key = "key 2",
        ///                 Value = "value 2",
        ///             },
        ///         },
        ///     });
        /// 
        ///     var testDatabaseUsers = Mongodbatlas.GetDatabaseUsers.Invoke(new()
        ///     {
        ///         ProjectId = testDatabaseUser.ProjectId,
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetDatabaseUsersResult> InvokeAsync(GetDatabaseUsersArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetDatabaseUsersResult>("mongodbatlas:index/getDatabaseUsers:getDatabaseUsers", args ?? new GetDatabaseUsersArgs(), options.WithDefaults());

        /// <summary>
        /// `mongodbatlas.getDatabaseUsers` describe all Database Users. This represents a database user which will be applied to all clusters within the project.
        /// 
        /// Each user has a set of roles that provide access to the project’s databases. User's roles apply to all the clusters in the project: if two clusters have a `products` database and a user has a role granting `read` access on the products database, the user has that access on both clusters.
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
        ///         AuthDatabaseName = "admin",
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
        ///         Labels = new[]
        ///         {
        ///             new Mongodbatlas.Inputs.DatabaseUserLabelArgs
        ///             {
        ///                 Key = "key 1",
        ///                 Value = "value 1",
        ///             },
        ///             new Mongodbatlas.Inputs.DatabaseUserLabelArgs
        ///             {
        ///                 Key = "key 2",
        ///                 Value = "value 2",
        ///             },
        ///         },
        ///     });
        /// 
        ///     var testDatabaseUsers = Mongodbatlas.GetDatabaseUsers.Invoke(new()
        ///     {
        ///         ProjectId = testDatabaseUser.ProjectId,
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetDatabaseUsersResult> Invoke(GetDatabaseUsersInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetDatabaseUsersResult>("mongodbatlas:index/getDatabaseUsers:getDatabaseUsers", args ?? new GetDatabaseUsersInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDatabaseUsersArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The unique ID for the project to get all database users.
        /// </summary>
        [Input("projectId", required: true)]
        public string ProjectId { get; set; } = null!;

        public GetDatabaseUsersArgs()
        {
        }
        public static new GetDatabaseUsersArgs Empty => new GetDatabaseUsersArgs();
    }

    public sealed class GetDatabaseUsersInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The unique ID for the project to get all database users.
        /// </summary>
        [Input("projectId", required: true)]
        public Input<string> ProjectId { get; set; } = null!;

        public GetDatabaseUsersInvokeArgs()
        {
        }
        public static new GetDatabaseUsersInvokeArgs Empty => new GetDatabaseUsersInvokeArgs();
    }


    [OutputType]
    public sealed class GetDatabaseUsersResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// ID of the Atlas project the user belongs to.
        /// </summary>
        public readonly string ProjectId;
        /// <summary>
        /// A list where each represents a Database user.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetDatabaseUsersResultResult> Results;

        [OutputConstructor]
        private GetDatabaseUsersResult(
            string id,

            string projectId,

            ImmutableArray<Outputs.GetDatabaseUsersResultResult> results)
        {
            Id = id;
            ProjectId = projectId;
            Results = results;
        }
    }
}
