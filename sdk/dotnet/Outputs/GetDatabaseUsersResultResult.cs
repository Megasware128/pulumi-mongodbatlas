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
    public sealed class GetDatabaseUsersResultResult
    {
        /// <summary>
        /// (Required) Database against which Atlas authenticates the user. A user must provide both a username and authentication database to log into MongoDB.
        /// Possible values include:
        /// * `admin` if `x509_type` and `aws_iam_type` are omitted or NONE.
        /// * `$external` if:
        /// * `x509_type` is MANAGED or CUSTOMER, or
        /// * `aws_iam_type` is USER or ROLE.
        /// </summary>
        public readonly string AuthDatabaseName;
        /// <summary>
        /// The new database user authenticates with AWS IAM credentials. Default is `NONE`, `USER` means user has AWS IAM user credentials, `ROLE` - means user has credentials associated with an AWS IAM role.
        /// </summary>
        public readonly string AwsIamType;
        public readonly ImmutableArray<Outputs.GetDatabaseUsersResultLabelResult> Labels;
        /// <summary>
        /// The unique ID for the project to get all database users.
        /// </summary>
        public readonly string ProjectId;
        /// <summary>
        /// List of user’s roles and the databases / collections on which the roles apply. A role allows the user to perform particular actions on the specified database. A role on the admin database can include privileges that apply to the other databases as well. See Roles below for more details.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetDatabaseUsersResultRoleResult> Roles;
        /// <summary>
        /// Array of clusters and Atlas Data Lakes that this user has access to.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetDatabaseUsersResultScopeResult> Scopes;
        /// <summary>
        /// Username for authenticating to MongoDB.
        /// </summary>
        public readonly string Username;
        /// <summary>
        /// X.509 method by which the provided username is authenticated.
        /// </summary>
        public readonly string X509Type;

        [OutputConstructor]
        private GetDatabaseUsersResultResult(
            string authDatabaseName,

            string awsIamType,

            ImmutableArray<Outputs.GetDatabaseUsersResultLabelResult> labels,

            string projectId,

            ImmutableArray<Outputs.GetDatabaseUsersResultRoleResult> roles,

            ImmutableArray<Outputs.GetDatabaseUsersResultScopeResult> scopes,

            string username,

            string x509Type)
        {
            AuthDatabaseName = authDatabaseName;
            AwsIamType = awsIamType;
            Labels = labels;
            ProjectId = projectId;
            Roles = roles;
            Scopes = scopes;
            Username = username;
            X509Type = x509Type;
        }
    }
}
