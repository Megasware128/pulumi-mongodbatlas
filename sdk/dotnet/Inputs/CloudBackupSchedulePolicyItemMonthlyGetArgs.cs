// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Mongodbatlas.Inputs
{

    public sealed class CloudBackupSchedulePolicyItemMonthlyGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Desired frequency of the new backup policy item specified by `frequency_type`.
        /// </summary>
        [Input("frequencyInterval", required: true)]
        public Input<int> FrequencyInterval { get; set; } = null!;

        /// <summary>
        /// Frequency associated with the export snapshot item.
        /// </summary>
        [Input("frequencyType")]
        public Input<string>? FrequencyType { get; set; }

        /// <summary>
        /// Scope of the backup policy item: days, weeks, or months.
        /// </summary>
        [Input("retentionUnit", required: true)]
        public Input<string> RetentionUnit { get; set; } = null!;

        /// <summary>
        /// Value to associate with `retention_unit`.
        /// </summary>
        [Input("retentionValue", required: true)]
        public Input<int> RetentionValue { get; set; } = null!;

        public CloudBackupSchedulePolicyItemMonthlyGetArgs()
        {
        }
        public static new CloudBackupSchedulePolicyItemMonthlyGetArgs Empty => new CloudBackupSchedulePolicyItemMonthlyGetArgs();
    }
}
