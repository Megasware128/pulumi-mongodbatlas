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
    public sealed class DataLakeStorageStore
    {
        public readonly ImmutableArray<string> AdditionalStorageClasses;
        public readonly string? Bucket;
        public readonly string? Delimiter;
        public readonly bool? IncludeTags;
        /// <summary>
        /// Name of the Atlas Data Lake.
        /// </summary>
        public readonly string? Name;
        public readonly string? Prefix;
        public readonly string? Provider;
        public readonly string? Region;

        [OutputConstructor]
        private DataLakeStorageStore(
            ImmutableArray<string> additionalStorageClasses,

            string? bucket,

            string? delimiter,

            bool? includeTags,

            string? name,

            string? prefix,

            string? provider,

            string? region)
        {
            AdditionalStorageClasses = additionalStorageClasses;
            Bucket = bucket;
            Delimiter = delimiter;
            IncludeTags = includeTags;
            Name = name;
            Prefix = prefix;
            Provider = provider;
            Region = region;
        }
    }
}
