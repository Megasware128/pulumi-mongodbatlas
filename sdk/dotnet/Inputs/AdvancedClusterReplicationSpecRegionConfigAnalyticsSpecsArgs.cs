// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Mongodbatlas.Inputs
{

    public sealed class AdvancedClusterReplicationSpecRegionConfigAnalyticsSpecsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Target throughput (IOPS) desired for AWS storage attached to your cluster. Set only if you selected AWS as your cloud service provider. You can't set this parameter for a multi-cloud cluster.
        /// </summary>
        [Input("diskIops")]
        public Input<int>? DiskIops { get; set; }

        /// <summary>
        /// Type of storage you want to attach to your AWS-provisioned cluster. Set only if you selected AWS as your cloud service provider. You can't set this parameter for a multi-cloud cluster. Valid values are:
        /// </summary>
        [Input("ebsVolumeType")]
        public Input<string>? EbsVolumeType { get; set; }

        /// <summary>
        /// Hardware specification for the instance sizes in this region. Each instance size has a default storage and memory capacity. The instance size you select applies to all the data-bearing hosts in your instance size.
        /// </summary>
        [Input("instanceSize", required: true)]
        public Input<string> InstanceSize { get; set; } = null!;

        /// <summary>
        /// Number of read-only nodes for Atlas to deploy to the region. Read-only nodes can never become the [primary](https://docs.atlas.mongodb.com/reference/glossary/#std-term-primary), but can enable local reads.
        /// </summary>
        [Input("nodeCount")]
        public Input<int>? NodeCount { get; set; }

        public AdvancedClusterReplicationSpecRegionConfigAnalyticsSpecsArgs()
        {
        }
    }
}