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
    public sealed class GetAdvancedClusterLabelResult
    {
        /// <summary>
        /// The key that you want to write.
        /// </summary>
        public readonly string Key;
        /// <summary>
        /// The value that you want to write.
        /// </summary>
        public readonly string Value;

        [OutputConstructor]
        private GetAdvancedClusterLabelResult(
            string key,

            string value)
        {
            Key = key;
            Value = value;
        }
    }
}
