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
    public sealed class GetDataLakesResultStorageDatabaseViewResult
    {
        public readonly string Name;
        public readonly string Pipeline;
        public readonly string Source;

        [OutputConstructor]
        private GetDataLakesResultStorageDatabaseViewResult(
            string name,

            string pipeline,

            string source)
        {
            Name = name;
            Pipeline = pipeline;
            Source = source;
        }
    }
}
