// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Mongodbatlas.Inputs
{

    public sealed class CloudProviderSnapshotBackupPolicyPolicyPolicyItemArgs : Pulumi.ResourceArgs
    {
        [Input("frequencyInterval", required: true)]
        public Input<int> FrequencyInterval { get; set; } = null!;

        [Input("frequencyType", required: true)]
        public Input<string> FrequencyType { get; set; } = null!;

        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        [Input("retentionUnit", required: true)]
        public Input<string> RetentionUnit { get; set; } = null!;

        [Input("retentionValue", required: true)]
        public Input<int> RetentionValue { get; set; } = null!;

        public CloudProviderSnapshotBackupPolicyPolicyPolicyItemArgs()
        {
        }
    }
}