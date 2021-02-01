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
    /// `mongodbatlas.PrivateLinkEndpoint` provides a Private Endpoint resource. This represents a Private Endpoint Service that can be created in an Atlas project.
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
    ///         var test = new Mongodbatlas.PrivateLinkEndpoint("test", new Mongodbatlas.PrivateLinkEndpointArgs
    ///         {
    ///             ProjectId = "&lt;PROJECT-ID&gt;",
    ///             ProviderName = "AWS/AZURE",
    ///             Region = "us-east-1",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Private Endpoint Service can be imported using project ID and username, in the format `{project_id}-{private_link_id}-{provider_name}`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import mongodbatlas:index/privateLinkEndpoint:PrivateLinkEndpoint test 1112222b3bf99403840e8934-3242342343112-AWS
    /// ```
    /// 
    ///  See detailed information for arguments and attributes[MongoDB API Private Endpoint Service](https://docs.atlas.mongodb.com/reference/api/private-endpoints-service-create-one//)
    /// </summary>
    [MongodbatlasResourceType("mongodbatlas:index/privateLinkEndpoint:PrivateLinkEndpoint")]
    public partial class PrivateLinkEndpoint : Pulumi.CustomResource
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
        /// All private endpoints that you have added to this Azure Private Link Service.
        /// </summary>
        [Output("privateEndpoints")]
        public Output<ImmutableArray<string>> PrivateEndpoints { get; private set; } = null!;

        /// <summary>
        /// Unique identifier of the AWS PrivateLink connection.
        /// </summary>
        [Output("privateLinkId")]
        public Output<string> PrivateLinkId { get; private set; } = null!;

        /// <summary>
        /// Name of the Azure Private Link Service that Atlas manages.
        /// </summary>
        [Output("privateLinkServiceName")]
        public Output<string> PrivateLinkServiceName { get; private set; } = null!;

        /// <summary>
        /// Resource ID of the Azure Private Link Service that Atlas manages.
        /// Returns one of the following values:
        /// </summary>
        [Output("privateLinkServiceResourceId")]
        public Output<string> PrivateLinkServiceResourceId { get; private set; } = null!;

        /// <summary>
        /// Required 	Unique identifier for the project.
        /// </summary>
        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;

        [Output("providerName")]
        public Output<string> ProviderName { get; private set; } = null!;

        /// <summary>
        /// Cloud provider region in which you want to create the private endpoint connection.
        /// Accepted values are: [AWS regions](https://docs.atlas.mongodb.com/reference/amazon-aws/#amazon-aws) and [AZURE regions](https://docs.atlas.mongodb.com/reference/microsoft-azure/#microsoft-azure)
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// Status of the Private Link Service.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;


        /// <summary>
        /// Create a PrivateLinkEndpoint resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public PrivateLinkEndpoint(string name, PrivateLinkEndpointArgs args, CustomResourceOptions? options = null)
            : base("mongodbatlas:index/privateLinkEndpoint:PrivateLinkEndpoint", name, args ?? new PrivateLinkEndpointArgs(), MakeResourceOptions(options, ""))
        {
        }

        private PrivateLinkEndpoint(string name, Input<string> id, PrivateLinkEndpointState? state = null, CustomResourceOptions? options = null)
            : base("mongodbatlas:index/privateLinkEndpoint:PrivateLinkEndpoint", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing PrivateLinkEndpoint resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static PrivateLinkEndpoint Get(string name, Input<string> id, PrivateLinkEndpointState? state = null, CustomResourceOptions? options = null)
        {
            return new PrivateLinkEndpoint(name, id, state, options);
        }
    }

    public sealed class PrivateLinkEndpointArgs : Pulumi.ResourceArgs
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
        /// Accepted values are: [AWS regions](https://docs.atlas.mongodb.com/reference/amazon-aws/#amazon-aws) and [AZURE regions](https://docs.atlas.mongodb.com/reference/microsoft-azure/#microsoft-azure)
        /// </summary>
        [Input("region", required: true)]
        public Input<string> Region { get; set; } = null!;

        public PrivateLinkEndpointArgs()
        {
        }
    }

    public sealed class PrivateLinkEndpointState : Pulumi.ResourceArgs
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

        [Input("privateEndpoints")]
        private InputList<string>? _privateEndpoints;

        /// <summary>
        /// All private endpoints that you have added to this Azure Private Link Service.
        /// </summary>
        public InputList<string> PrivateEndpoints
        {
            get => _privateEndpoints ?? (_privateEndpoints = new InputList<string>());
            set => _privateEndpoints = value;
        }

        /// <summary>
        /// Unique identifier of the AWS PrivateLink connection.
        /// </summary>
        [Input("privateLinkId")]
        public Input<string>? PrivateLinkId { get; set; }

        /// <summary>
        /// Name of the Azure Private Link Service that Atlas manages.
        /// </summary>
        [Input("privateLinkServiceName")]
        public Input<string>? PrivateLinkServiceName { get; set; }

        /// <summary>
        /// Resource ID of the Azure Private Link Service that Atlas manages.
        /// Returns one of the following values:
        /// </summary>
        [Input("privateLinkServiceResourceId")]
        public Input<string>? PrivateLinkServiceResourceId { get; set; }

        /// <summary>
        /// Required 	Unique identifier for the project.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        [Input("providerName")]
        public Input<string>? ProviderName { get; set; }

        /// <summary>
        /// Cloud provider region in which you want to create the private endpoint connection.
        /// Accepted values are: [AWS regions](https://docs.atlas.mongodb.com/reference/amazon-aws/#amazon-aws) and [AZURE regions](https://docs.atlas.mongodb.com/reference/microsoft-azure/#microsoft-azure)
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// Status of the Private Link Service.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public PrivateLinkEndpointState()
        {
        }
    }
}