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
    public sealed class CloudProviderSnapshotBackupPolicyPolicy
    {
        public readonly string Id;
        public readonly ImmutableArray<Outputs.CloudProviderSnapshotBackupPolicyPolicyPolicyItem> PolicyItems;

        [OutputConstructor]
        private CloudProviderSnapshotBackupPolicyPolicy(
            string id,

            ImmutableArray<Outputs.CloudProviderSnapshotBackupPolicyPolicyPolicyItem> policyItems)
        {
            Id = id;
            PolicyItems = policyItems;
        }
    }
}