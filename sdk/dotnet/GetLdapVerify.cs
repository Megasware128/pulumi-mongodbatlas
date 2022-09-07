// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Mongodbatlas
{
    public static class GetLdapVerify
    {
        /// <summary>
        /// `mongodbatlas.LdapVerify` describes a LDAP Verify.
        /// 
        /// &gt; **NOTE:** Groups and projects are synonymous terms. You may find **group_id** in the official documentation.
        /// </summary>
        public static Task<GetLdapVerifyResult> InvokeAsync(GetLdapVerifyArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetLdapVerifyResult>("mongodbatlas:index/getLdapVerify:getLdapVerify", args ?? new GetLdapVerifyArgs(), options.WithDefaults());

        /// <summary>
        /// `mongodbatlas.LdapVerify` describes a LDAP Verify.
        /// 
        /// &gt; **NOTE:** Groups and projects are synonymous terms. You may find **group_id** in the official documentation.
        /// </summary>
        public static Output<GetLdapVerifyResult> Invoke(GetLdapVerifyInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetLdapVerifyResult>("mongodbatlas:index/getLdapVerify:getLdapVerify", args ?? new GetLdapVerifyInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetLdapVerifyArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Unique identifier for the Atlas project associated with the verification request.
        /// </summary>
        [Input("projectId", required: true)]
        public string ProjectId { get; set; } = null!;

        /// <summary>
        /// Unique identifier of a request to verify an LDAP configuration.
        /// </summary>
        [Input("requestId", required: true)]
        public string RequestId { get; set; } = null!;

        public GetLdapVerifyArgs()
        {
        }
        public static new GetLdapVerifyArgs Empty => new GetLdapVerifyArgs();
    }

    public sealed class GetLdapVerifyInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Unique identifier for the Atlas project associated with the verification request.
        /// </summary>
        [Input("projectId", required: true)]
        public Input<string> ProjectId { get; set; } = null!;

        /// <summary>
        /// Unique identifier of a request to verify an LDAP configuration.
        /// </summary>
        [Input("requestId", required: true)]
        public Input<string> RequestId { get; set; } = null!;

        public GetLdapVerifyInvokeArgs()
        {
        }
        public static new GetLdapVerifyInvokeArgs Empty => new GetLdapVerifyInvokeArgs();
    }


    [OutputType]
    public sealed class GetLdapVerifyResult
    {
        /// <summary>
        /// The user DN that Atlas uses to connect to the LDAP server.
        /// </summary>
        public readonly string BindUsername;
        /// <summary>
        /// (Required) The hostname or IP address of the LDAP server.
        /// </summary>
        public readonly string Hostname;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// One or more links to sub-resources. The relations in the URLs are explained in the Web Linking Specification.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetLdapVerifyLinkResult> Links;
        /// <summary>
        /// LDAP ConfigurationThe port to which the LDAP server listens for client connections.
        /// </summary>
        public readonly int Port;
        public readonly string ProjectId;
        /// <summary>
        /// The unique identifier for the request to verify the LDAP over TLS/SSL configuration.
        /// </summary>
        public readonly string RequestId;
        /// <summary>
        /// The current status of the LDAP over TLS/SSL configuration.
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// Array of validation messages related to the verification of the provided LDAP over TLS/SSL configuration details.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetLdapVerifyValidationResult> Validations;

        [OutputConstructor]
        private GetLdapVerifyResult(
            string bindUsername,

            string hostname,

            string id,

            ImmutableArray<Outputs.GetLdapVerifyLinkResult> links,

            int port,

            string projectId,

            string requestId,

            string status,

            ImmutableArray<Outputs.GetLdapVerifyValidationResult> validations)
        {
            BindUsername = bindUsername;
            Hostname = hostname;
            Id = id;
            Links = links;
            Port = port;
            ProjectId = projectId;
            RequestId = requestId;
            Status = status;
            Validations = validations;
        }
    }
}
