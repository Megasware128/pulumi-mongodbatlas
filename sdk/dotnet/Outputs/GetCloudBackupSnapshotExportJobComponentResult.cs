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
    public sealed class GetCloudBackupSnapshotExportJobComponentResult
    {
        /// <summary>
        /// _Returned for sharded clusters only._ Export job details for each replica set in the sharded cluster.
        /// </summary>
        public readonly string ExportId;
        /// <summary>
        /// _Returned for sharded clusters only._ Unique identifier of the export job for the replica set.
        /// </summary>
        public readonly string ReplicaSetName;

        [OutputConstructor]
        private GetCloudBackupSnapshotExportJobComponentResult(
            string exportId,

            string replicaSetName)
        {
            ExportId = exportId;
            ReplicaSetName = replicaSetName;
        }
    }
}
