// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi.Utilities;

namespace Pulumi.Mongodbatlas
{
    public static class GetLdapConfiguration
    {
        /// <summary>
        /// `mongodbatlas.LdapConfiguration` describes a LDAP Configuration.
        /// 
        /// &gt; **NOTE:** Groups and projects are synonymous terms. You may find **group_id** in the official documentation.
        /// </summary>
        public static Task<GetLdapConfigurationResult> InvokeAsync(GetLdapConfigurationArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetLdapConfigurationResult>("mongodbatlas:index/getLdapConfiguration:getLdapConfiguration", args ?? new GetLdapConfigurationArgs(), options.WithVersion());

        /// <summary>
        /// `mongodbatlas.LdapConfiguration` describes a LDAP Configuration.
        /// 
        /// &gt; **NOTE:** Groups and projects are synonymous terms. You may find **group_id** in the official documentation.
        /// </summary>
        public static Output<GetLdapConfigurationResult> Invoke(GetLdapConfigurationInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetLdapConfigurationResult>("mongodbatlas:index/getLdapConfiguration:getLdapConfiguration", args ?? new GetLdapConfigurationInvokeArgs(), options.WithVersion());
    }


    public sealed class GetLdapConfigurationArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Identifier for the Atlas project associated with the LDAP over TLS/SSL configuration.
        /// </summary>
        [Input("projectId", required: true)]
        public string ProjectId { get; set; } = null!;

        public GetLdapConfigurationArgs()
        {
        }
    }

    public sealed class GetLdapConfigurationInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Identifier for the Atlas project associated with the LDAP over TLS/SSL configuration.
        /// </summary>
        [Input("projectId", required: true)]
        public Input<string> ProjectId { get; set; } = null!;

        public GetLdapConfigurationInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetLdapConfigurationResult
    {
        /// <summary>
        /// Specifies whether user authentication with LDAP is enabled.
        /// </summary>
        public readonly bool AuthenticationEnabled;
        /// <summary>
        /// Specifies whether user authorization with LDAP is enabled.
        /// </summary>
        public readonly bool AuthorizationEnabled;
        /// <summary>
        /// An LDAP query template that Atlas executes to obtain the LDAP groups to which the authenticated user belongs.
        /// </summary>
        public readonly string AuthzQueryTemplate;
        /// <summary>
        /// The password used to authenticate the `bind_username`.
        /// </summary>
        public readonly string BindPassword;
        /// <summary>
        /// The user DN that Atlas uses to connect to the LDAP server.
        /// </summary>
        public readonly string BindUsername;
        /// <summary>
        /// CA certificate used to verify the identify of the LDAP server.
        /// </summary>
        public readonly string CaCertificate;
        /// <summary>
        /// (Required) The hostname or IP address of the LDAP server.
        /// </summary>
        public readonly string Hostname;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The port to which the LDAP server listens for client connections.
        /// </summary>
        public readonly int Port;
        public readonly string ProjectId;
        /// <summary>
        /// Maps an LDAP username for authentication to an LDAP Distinguished Name (DN).
        /// * `user_to_dn_mapping.0.match` - A regular expression to match against a provided LDAP username.
        /// * `user_to_dn_mapping.0.substitution` - An LDAP Distinguished Name (DN) formatting template that converts the LDAP name matched by the `match` regular expression into an LDAP Distinguished Name.
        /// * `user_to_dn_mapping.0.ldap_query` - An LDAP query formatting template that inserts the LDAP name matched by the `match` regular expression into an LDAP query URI as specified by RFC 4515 and RFC 4516.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetLdapConfigurationUserToDnMappingResult> UserToDnMappings;

        [OutputConstructor]
        private GetLdapConfigurationResult(
            bool authenticationEnabled,

            bool authorizationEnabled,

            string authzQueryTemplate,

            string bindPassword,

            string bindUsername,

            string caCertificate,

            string hostname,

            string id,

            int port,

            string projectId,

            ImmutableArray<Outputs.GetLdapConfigurationUserToDnMappingResult> userToDnMappings)
        {
            AuthenticationEnabled = authenticationEnabled;
            AuthorizationEnabled = authorizationEnabled;
            AuthzQueryTemplate = authzQueryTemplate;
            BindPassword = bindPassword;
            BindUsername = bindUsername;
            CaCertificate = caCertificate;
            Hostname = hostname;
            Id = id;
            Port = port;
            ProjectId = projectId;
            UserToDnMappings = userToDnMappings;
        }
    }
}
