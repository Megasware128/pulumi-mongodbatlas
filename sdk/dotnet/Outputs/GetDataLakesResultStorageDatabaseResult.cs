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
    public sealed class GetDataLakesResultStorageDatabaseResult
    {
        public readonly ImmutableArray<Outputs.GetDataLakesResultStorageDatabaseCollectionResult> Collections;
        public readonly int MaxWildcardCollections;
        public readonly string Name;
        public readonly ImmutableArray<Outputs.GetDataLakesResultStorageDatabaseViewResult> Views;

        [OutputConstructor]
        private GetDataLakesResultStorageDatabaseResult(
            ImmutableArray<Outputs.GetDataLakesResultStorageDatabaseCollectionResult> collections,

            int maxWildcardCollections,

            string name,

            ImmutableArray<Outputs.GetDataLakesResultStorageDatabaseViewResult> views)
        {
            Collections = collections;
            MaxWildcardCollections = maxWildcardCollections;
            Name = name;
            Views = views;
        }
    }
}