// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Mongodbatlas.Inputs
{

    public sealed class ClusterSnapshotBackupPolicyPolicyPolicyItemGetArgs : Pulumi.ResourceArgs
    {
        [Input("frequencyInterval")]
        public Input<int>? FrequencyInterval { get; set; }

        [Input("frequencyType")]
        public Input<string>? FrequencyType { get; set; }

        [Input("retentionUnit")]
        public Input<string>? RetentionUnit { get; set; }

        [Input("retentionValue")]
        public Input<int>? RetentionValue { get; set; }

        public ClusterSnapshotBackupPolicyPolicyPolicyItemGetArgs()
        {
        }
    }
}
