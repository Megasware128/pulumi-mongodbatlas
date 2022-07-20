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
    public sealed class GetCloudBackupScheduleExportResult
    {
        /// <summary>
        /// Unique identifier of the mongodbatlas.CloudBackupSnapshotExportBucket export_bucket_id value.
        /// </summary>
        public readonly string ExportBucketId;
        /// <summary>
        /// Frequency associated with the backup policy item.
        /// </summary>
        public readonly string FrequencyType;

        [OutputConstructor]
        private GetCloudBackupScheduleExportResult(
            string exportBucketId,

            string frequencyType)
        {
            ExportBucketId = exportBucketId;
            FrequencyType = frequencyType;
        }
    }
}
