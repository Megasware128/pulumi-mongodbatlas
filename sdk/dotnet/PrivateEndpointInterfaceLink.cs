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
    /// `mongodbatlas.PrivateEndpointInterfaceLink` provides a Private Endpoint Interface Link resource. This represents a Private Endpoint Interface Link, which adds one interface endpoint to a private endpoint connection in an Atlas project.
    /// 
    /// &gt; **IMPORTANT:**You must have one of the following roles to successfully handle the resource:
    ///   * Organization Owner
    ///   * Project Owner
    /// 
    /// &gt; **NOTE:** Groups and projects are synonymous terms. You may find group_id in the official documentation.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// using Mongodbatlas = Pulumi.Mongodbatlas;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var testPrivateEndpoint = new Mongodbatlas.PrivateEndpoint("testPrivateEndpoint", new Mongodbatlas.PrivateEndpointArgs
    ///         {
    ///             ProjectId = "&lt;PROJECT_ID&gt;",
    ///             ProviderName = "AWS",
    ///             Region = "us-east-1",
    ///         });
    ///         var ptfeService = new Aws.Ec2.VpcEndpoint("ptfeService", new Aws.Ec2.VpcEndpointArgs
    ///         {
    ///             SecurityGroupIds = 
    ///             {
    ///                 "sg-3f238186",
    ///             },
    ///             ServiceName = testPrivateEndpoint.EndpointServiceName,
    ///             SubnetIds = 
    ///             {
    ///                 "subnet-de0406d2",
    ///             },
    ///             VpcEndpointType = "Interface",
    ///             VpcId = "vpc-7fc0a543",
    ///         });
    ///         var testPrivateEndpointInterfaceLink = new Mongodbatlas.PrivateEndpointInterfaceLink("testPrivateEndpointInterfaceLink", new Mongodbatlas.PrivateEndpointInterfaceLinkArgs
    ///         {
    ///             InterfaceEndpointId = ptfeService.Id,
    ///             PrivateLinkId = testPrivateEndpoint.PrivateLinkId,
    ///             ProjectId = testPrivateEndpoint.ProjectId,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Private Endpoint Link Connection can be imported using project ID and username, in the format `{project_id}-{private_link_id}-{interface_endpoint_id}`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import mongodbatlas:index/privateEndpointInterfaceLink:PrivateEndpointInterfaceLink test 1112222b3bf99403840e8934-3242342343112-vpce-4242342343
    /// ```
    /// 
    ///  See detailed information for arguments and attributes[MongoDB API Private Endpoint Link Connection](https://docs.atlas.mongodb.com/reference/api/private-endpoint-create-one-interface-endpoint/)
    /// </summary>
    public partial class PrivateEndpointInterfaceLink : Pulumi.CustomResource
    {
        /// <summary>
        /// Status of the interface endpoint.
        /// Returns one of the following values:
        /// </summary>
        [Output("connectionStatus")]
        public Output<string> ConnectionStatus { get; private set; } = null!;

        /// <summary>
        /// Indicates if Atlas received a request to remove the interface endpoint from the private endpoint connection.
        /// </summary>
        [Output("deleteRequested")]
        public Output<bool> DeleteRequested { get; private set; } = null!;

        /// <summary>
        /// Error message pertaining to the interface endpoint. Returns null if there are no errors.
        /// </summary>
        [Output("errorMessage")]
        public Output<string> ErrorMessage { get; private set; } = null!;

        /// <summary>
        /// Unique identifier of the interface endpoint you created in your VPC with the AWS resource.
        /// </summary>
        [Output("interfaceEndpointId")]
        public Output<string> InterfaceEndpointId { get; private set; } = null!;

        /// <summary>
        /// Unique identifier of the AWS PrivateLink connection which is created by `mongodbatlas.PrivateEndpoint` resource.
        /// </summary>
        [Output("privateLinkId")]
        public Output<string> PrivateLinkId { get; private set; } = null!;

        /// <summary>
        /// Unique identifier for the project.
        /// </summary>
        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;


        /// <summary>
        /// Create a PrivateEndpointInterfaceLink resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public PrivateEndpointInterfaceLink(string name, PrivateEndpointInterfaceLinkArgs args, CustomResourceOptions? options = null)
            : base("mongodbatlas:index/privateEndpointInterfaceLink:PrivateEndpointInterfaceLink", name, args ?? new PrivateEndpointInterfaceLinkArgs(), MakeResourceOptions(options, ""))
        {
        }

        private PrivateEndpointInterfaceLink(string name, Input<string> id, PrivateEndpointInterfaceLinkState? state = null, CustomResourceOptions? options = null)
            : base("mongodbatlas:index/privateEndpointInterfaceLink:PrivateEndpointInterfaceLink", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing PrivateEndpointInterfaceLink resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static PrivateEndpointInterfaceLink Get(string name, Input<string> id, PrivateEndpointInterfaceLinkState? state = null, CustomResourceOptions? options = null)
        {
            return new PrivateEndpointInterfaceLink(name, id, state, options);
        }
    }

    public sealed class PrivateEndpointInterfaceLinkArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Unique identifier of the interface endpoint you created in your VPC with the AWS resource.
        /// </summary>
        [Input("interfaceEndpointId", required: true)]
        public Input<string> InterfaceEndpointId { get; set; } = null!;

        /// <summary>
        /// Unique identifier of the AWS PrivateLink connection which is created by `mongodbatlas.PrivateEndpoint` resource.
        /// </summary>
        [Input("privateLinkId", required: true)]
        public Input<string> PrivateLinkId { get; set; } = null!;

        /// <summary>
        /// Unique identifier for the project.
        /// </summary>
        [Input("projectId", required: true)]
        public Input<string> ProjectId { get; set; } = null!;

        public PrivateEndpointInterfaceLinkArgs()
        {
        }
    }

    public sealed class PrivateEndpointInterfaceLinkState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Status of the interface endpoint.
        /// Returns one of the following values:
        /// </summary>
        [Input("connectionStatus")]
        public Input<string>? ConnectionStatus { get; set; }

        /// <summary>
        /// Indicates if Atlas received a request to remove the interface endpoint from the private endpoint connection.
        /// </summary>
        [Input("deleteRequested")]
        public Input<bool>? DeleteRequested { get; set; }

        /// <summary>
        /// Error message pertaining to the interface endpoint. Returns null if there are no errors.
        /// </summary>
        [Input("errorMessage")]
        public Input<string>? ErrorMessage { get; set; }

        /// <summary>
        /// Unique identifier of the interface endpoint you created in your VPC with the AWS resource.
        /// </summary>
        [Input("interfaceEndpointId")]
        public Input<string>? InterfaceEndpointId { get; set; }

        /// <summary>
        /// Unique identifier of the AWS PrivateLink connection which is created by `mongodbatlas.PrivateEndpoint` resource.
        /// </summary>
        [Input("privateLinkId")]
        public Input<string>? PrivateLinkId { get; set; }

        /// <summary>
        /// Unique identifier for the project.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        public PrivateEndpointInterfaceLinkState()
        {
        }
    }
}
