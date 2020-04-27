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
    public sealed class CloudProviderSnapshotRestoreJobDeliveryType
    {
        public readonly bool? Automated;
        public readonly bool? Download;
        /// <summary>
        /// Name of the target Atlas cluster to which the restore job restores the snapshot. Only required if deliveryType is automated.
        /// </summary>
        public readonly string? TargetClusterName;
        public readonly string? TargetProjectId;

        [OutputConstructor]
        private CloudProviderSnapshotRestoreJobDeliveryType(
            bool? automated,

            bool? download,

            string? targetClusterName,

            string? targetProjectId)
        {
            Automated = automated;
            Download = download;
            TargetClusterName = targetClusterName;
            TargetProjectId = targetProjectId;
        }
    }
}