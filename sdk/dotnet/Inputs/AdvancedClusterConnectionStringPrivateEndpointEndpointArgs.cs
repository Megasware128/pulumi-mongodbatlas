// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Mongodbatlas.Inputs
{

    public sealed class AdvancedClusterConnectionStringPrivateEndpointEndpointArgs : global::Pulumi.ResourceArgs
    {
        [Input("endpointId")]
        public Input<string>? EndpointId { get; set; }

        /// <summary>
        /// Cloud service provider on which the servers are provisioned.
        /// The possible values are:
        /// </summary>
        [Input("providerName")]
        public Input<string>? ProviderName { get; set; }

        [Input("region")]
        public Input<string>? Region { get; set; }

        public AdvancedClusterConnectionStringPrivateEndpointEndpointArgs()
        {
        }
        public static new AdvancedClusterConnectionStringPrivateEndpointEndpointArgs Empty => new AdvancedClusterConnectionStringPrivateEndpointEndpointArgs();
    }
}
