// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Mongodbatlas.Inputs
{

    public sealed class CloudProviderSnapshotBackupPolicyPolicyArgs : global::Pulumi.ResourceArgs
    {
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        [Input("policyItems", required: true)]
        private InputList<Inputs.CloudProviderSnapshotBackupPolicyPolicyPolicyItemArgs>? _policyItems;
        public InputList<Inputs.CloudProviderSnapshotBackupPolicyPolicyPolicyItemArgs> PolicyItems
        {
            get => _policyItems ?? (_policyItems = new InputList<Inputs.CloudProviderSnapshotBackupPolicyPolicyPolicyItemArgs>());
            set => _policyItems = value;
        }

        public CloudProviderSnapshotBackupPolicyPolicyArgs()
        {
        }
        public static new CloudProviderSnapshotBackupPolicyPolicyArgs Empty => new CloudProviderSnapshotBackupPolicyPolicyArgs();
    }
}
