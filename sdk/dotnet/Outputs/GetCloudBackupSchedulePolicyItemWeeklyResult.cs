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
    public sealed class GetCloudBackupSchedulePolicyItemWeeklyResult
    {
        /// <summary>
        /// Desired frequency of the new backup policy item specified by `frequency_type`.
        /// </summary>
        public readonly int FrequencyInterval;
        /// <summary>
        /// Frequency associated with the backup policy item.
        /// </summary>
        public readonly string FrequencyType;
        /// <summary>
        /// Unique identifier of the backup policy item.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Scope of the backup policy item: days, weeks, or months.
        /// </summary>
        public readonly string RetentionUnit;
        /// <summary>
        /// Value to associate with `retention_unit`.
        /// </summary>
        public readonly int RetentionValue;

        [OutputConstructor]
        private GetCloudBackupSchedulePolicyItemWeeklyResult(
            int frequencyInterval,

            string frequencyType,

            string id,

            string retentionUnit,

            int retentionValue)
        {
            FrequencyInterval = frequencyInterval;
            FrequencyType = frequencyType;
            Id = id;
            RetentionUnit = retentionUnit;
            RetentionValue = retentionValue;
        }
    }
}
