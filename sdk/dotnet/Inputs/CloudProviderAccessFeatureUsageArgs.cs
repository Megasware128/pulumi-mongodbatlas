// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Mongodbatlas.Inputs
{

    public sealed class CloudProviderAccessFeatureUsageArgs : global::Pulumi.ResourceArgs
    {
        [Input("featureId")]
        private InputMap<object>? _featureId;
        public InputMap<object> FeatureId
        {
            get => _featureId ?? (_featureId = new InputMap<object>());
            set => _featureId = value;
        }

        [Input("featureType")]
        public Input<string>? FeatureType { get; set; }

        public CloudProviderAccessFeatureUsageArgs()
        {
        }
        public static new CloudProviderAccessFeatureUsageArgs Empty => new CloudProviderAccessFeatureUsageArgs();
    }
}
