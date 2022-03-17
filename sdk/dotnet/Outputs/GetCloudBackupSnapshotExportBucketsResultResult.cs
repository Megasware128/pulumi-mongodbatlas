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
    public sealed class GetCloudBackupSnapshotExportBucketsResultResult
    {
        /// <summary>
        /// Name of the bucket that the provided role ID is authorized to access. You must also specify the `iam_role_id`.
        /// </summary>
        public readonly string BucketName;
        /// <summary>
        /// Name of the provider of the cloud service where Atlas can access the S3 bucket. Atlas only supports `AWS`.
        /// </summary>
        public readonly string CloudProvider;
        /// <summary>
        /// Unique identifier of the snapshot bucket id.
        /// </summary>
        public readonly string ExportBucketId;
        /// <summary>
        /// Unique identifier of the role that Atlas can use to access the bucket. You must also specify the `bucket_name`.
        /// </summary>
        public readonly string IamRoleId;

        [OutputConstructor]
        private GetCloudBackupSnapshotExportBucketsResultResult(
            string bucketName,

            string cloudProvider,

            string exportBucketId,

            string iamRoleId)
        {
            BucketName = bucketName;
            CloudProvider = cloudProvider;
            ExportBucketId = exportBucketId;
            IamRoleId = iamRoleId;
        }
    }
}
