// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Mongodbatlas
{
    /// <summary>
    /// `mongodbatlas.PrivateEndpoint` provides a Private Endpoint resource. This represents a Private Endpoint Connection that can be created in an Atlas project.
    /// 
    /// &gt; **IMPORTANT:**You must have one of the following roles to successfully handle the resource:
    ///   * Organization Owner
    ///   * Project Owner
    /// 
    /// &gt; **NOTE:** Groups and projects are synonymous terms. You may find group_id in the official documentation.
    /// 
    /// &gt; **NOTE:** A network container is created for a private endpoint to reside in if one does not yet exist in the project.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Mongodbatlas = Pulumi.Mongodbatlas;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var test = new Mongodbatlas.PrivateEndpoint("test", new Mongodbatlas.PrivateEndpointArgs
    ///         {
    ///             ProjectId = "&lt;PROJECT-ID&gt;",
    ///             ProviderName = "AWS",
    ///             Region = "us-east-1",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Private Endpoint Connection can be imported using project ID and username, in the format `{project_id}-{private_link_id}`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import mongodbatlas:index/privateEndpoint:PrivateEndpoint test 1112222b3bf99403840e8934-3242342343112
    /// ```
    /// 
    ///  See detailed information for arguments and attributes[MongoDB API Private Endpoint Connection](https://docs.atlas.mongodb.com/reference/api/private-endpoint-create-one-private-endpoint-connection/)
    /// </summary>
    public partial class PrivateEndpoint : Pulumi.CustomResource
    {
        /// <summary>
        /// Name of the PrivateLink endpoint service in AWS. Returns null while the endpoint service is being created.
        /// </summary>
        [Output("endpointServiceName")]
        public Output<string> EndpointServiceName { get; private set; } = null!;

        /// <summary>
        /// Error message pertaining to the AWS PrivateLink connection. Returns null if there are no errors.
        /// </summary>
        [Output("errorMessage")]
        public Output<string> ErrorMessage { get; private set; } = null!;

        /// <summary>
        /// Unique identifiers of the interface endpoints in your VPC that you added to the AWS PrivateLink connection.
        /// </summary>
        [Output("interfaceEndpoints")]
        public Output<ImmutableArray<string>> InterfaceEndpoints { get; private set; } = null!;

        /// <summary>
        /// Unique identifier of the AWS PrivateLink connection.
        /// </summary>
        [Output("privateLinkId")]
        public Output<string> PrivateLinkId { get; private set; } = null!;

        /// <summary>
        /// Required 	Unique identifier for the project.
        /// </summary>
        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;

        [Output("providerName")]
        public Output<string> ProviderName { get; private set; } = null!;

        /// <summary>
        /// Cloud provider region in which you want to create the private endpoint connection.
        /// Accepted values are:
        /// * `us-east-1`
        /// * `us-east-2`
        /// * `us-west-1`
        /// * `us-west-2`
        /// * `ca-central-1`
        /// * `sa-east-1`
        /// * `eu-north-1`
        /// * `eu-west-1`
        /// * `eu-west-2`
        /// * `eu-west-3`
        /// * `eu-central-1`
        /// * `me-south-1`
        /// * `ap-northeast-1`
        /// * `ap-northeast-2`
        /// * `ap-south-1`
        /// * `ap-southeast-1`
        /// * `ap-southeast-2`
        /// * `ap-east-1`
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// Status of the AWS PrivateLink connection.
        /// Returns one of the following values:
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;


        /// <summary>
        /// Create a PrivateEndpoint resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public PrivateEndpoint(string name, PrivateEndpointArgs args, CustomResourceOptions? options = null)
            : base("mongodbatlas:index/privateEndpoint:PrivateEndpoint", name, args ?? new PrivateEndpointArgs(), MakeResourceOptions(options, ""))
        {
        }

        private PrivateEndpoint(string name, Input<string> id, PrivateEndpointState? state = null, CustomResourceOptions? options = null)
            : base("mongodbatlas:index/privateEndpoint:PrivateEndpoint", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing PrivateEndpoint resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static PrivateEndpoint Get(string name, Input<string> id, PrivateEndpointState? state = null, CustomResourceOptions? options = null)
        {
            return new PrivateEndpoint(name, id, state, options);
        }
    }

    public sealed class PrivateEndpointArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Required 	Unique identifier for the project.
        /// </summary>
        [Input("projectId", required: true)]
        public Input<string> ProjectId { get; set; } = null!;

        [Input("providerName", required: true)]
        public Input<string> ProviderName { get; set; } = null!;

        /// <summary>
        /// Cloud provider region in which you want to create the private endpoint connection.
        /// Accepted values are:
        /// * `us-east-1`
        /// * `us-east-2`
        /// * `us-west-1`
        /// * `us-west-2`
        /// * `ca-central-1`
        /// * `sa-east-1`
        /// * `eu-north-1`
        /// * `eu-west-1`
        /// * `eu-west-2`
        /// * `eu-west-3`
        /// * `eu-central-1`
        /// * `me-south-1`
        /// * `ap-northeast-1`
        /// * `ap-northeast-2`
        /// * `ap-south-1`
        /// * `ap-southeast-1`
        /// * `ap-southeast-2`
        /// * `ap-east-1`
        /// </summary>
        [Input("region", required: true)]
        public Input<string> Region { get; set; } = null!;

        public PrivateEndpointArgs()
        {
        }
    }

    public sealed class PrivateEndpointState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the PrivateLink endpoint service in AWS. Returns null while the endpoint service is being created.
        /// </summary>
        [Input("endpointServiceName")]
        public Input<string>? EndpointServiceName { get; set; }

        /// <summary>
        /// Error message pertaining to the AWS PrivateLink connection. Returns null if there are no errors.
        /// </summary>
        [Input("errorMessage")]
        public Input<string>? ErrorMessage { get; set; }

        [Input("interfaceEndpoints")]
        private InputList<string>? _interfaceEndpoints;

        /// <summary>
        /// Unique identifiers of the interface endpoints in your VPC that you added to the AWS PrivateLink connection.
        /// </summary>
        public InputList<string> InterfaceEndpoints
        {
            get => _interfaceEndpoints ?? (_interfaceEndpoints = new InputList<string>());
            set => _interfaceEndpoints = value;
        }

        /// <summary>
        /// Unique identifier of the AWS PrivateLink connection.
        /// </summary>
        [Input("privateLinkId")]
        public Input<string>? PrivateLinkId { get; set; }

        /// <summary>
        /// Required 	Unique identifier for the project.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        [Input("providerName")]
        public Input<string>? ProviderName { get; set; }

        /// <summary>
        /// Cloud provider region in which you want to create the private endpoint connection.
        /// Accepted values are:
        /// * `us-east-1`
        /// * `us-east-2`
        /// * `us-west-1`
        /// * `us-west-2`
        /// * `ca-central-1`
        /// * `sa-east-1`
        /// * `eu-north-1`
        /// * `eu-west-1`
        /// * `eu-west-2`
        /// * `eu-west-3`
        /// * `eu-central-1`
        /// * `me-south-1`
        /// * `ap-northeast-1`
        /// * `ap-northeast-2`
        /// * `ap-south-1`
        /// * `ap-southeast-1`
        /// * `ap-southeast-2`
        /// * `ap-east-1`
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// Status of the AWS PrivateLink connection.
        /// Returns one of the following values:
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public PrivateEndpointState()
        {
        }
    }
}
