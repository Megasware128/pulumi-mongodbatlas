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
    public sealed class GetCloudProviderAccessAwsIamRoleFeatureUsageResult
    {
        public readonly string FeatureId;
        public readonly string FeatureType;

        [OutputConstructor]
        private GetCloudProviderAccessAwsIamRoleFeatureUsageResult(
            string featureId,

            string featureType)
        {
            FeatureId = featureId;
            FeatureType = featureType;
        }
    }
}
