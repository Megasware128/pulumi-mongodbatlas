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
    public sealed class GetOnlineArchivesResultCriteriaResult
    {
        public readonly string DateField;
        public readonly string DateFormat;
        public readonly int ExpireAfterDays;
        public readonly string Query;
        public readonly string Type;

        [OutputConstructor]
        private GetOnlineArchivesResultCriteriaResult(
            string dateField,

            string dateFormat,

            int expireAfterDays,

            string query,

            string type)
        {
            DateField = dateField;
            DateFormat = dateFormat;
            ExpireAfterDays = expireAfterDays;
            Query = query;
            Type = type;
        }
    }
}
