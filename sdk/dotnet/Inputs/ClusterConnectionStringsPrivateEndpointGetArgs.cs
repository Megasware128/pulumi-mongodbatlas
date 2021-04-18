// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Mongodbatlas.Inputs
{

    public sealed class ClusterConnectionStringsPrivateEndpointGetArgs : Pulumi.ResourceArgs
    {
        [Input("connectionString")]
        public Input<string>? ConnectionString { get; set; }

        [Input("endpoints")]
        private InputList<Inputs.ClusterConnectionStringsPrivateEndpointEndpointGetArgs>? _endpoints;
        public InputList<Inputs.ClusterConnectionStringsPrivateEndpointEndpointGetArgs> Endpoints
        {
            get => _endpoints ?? (_endpoints = new InputList<Inputs.ClusterConnectionStringsPrivateEndpointEndpointGetArgs>());
            set => _endpoints = value;
        }

        [Input("srvConnectionString")]
        public Input<string>? SrvConnectionString { get; set; }

        [Input("type")]
        public Input<string>? Type { get; set; }

        public ClusterConnectionStringsPrivateEndpointGetArgs()
        {
        }
    }
}