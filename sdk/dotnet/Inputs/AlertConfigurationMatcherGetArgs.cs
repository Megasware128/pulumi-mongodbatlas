// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Mongodbatlas.Inputs
{

    public sealed class AlertConfigurationMatcherGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the field in the target object to match on.
        /// Host alerts support these fields:
        /// - `TYPE_NAME`
        /// - `HOSTNAME`
        /// - `PORT`
        /// - `HOSTNAME_AND_PORT`
        /// - `REPLICA_SET_NAME`
        /// Replica set alerts support these fields:
        /// - `REPLICA_SET_NAME`
        /// - `SHARD_NAME`
        /// - `CLUSTER_NAME`
        /// Sharded cluster alerts support these fields:
        /// - `CLUSTER_NAME`
        /// - `SHARD_NAME`
        /// </summary>
        [Input("fieldName")]
        public Input<string>? FieldName { get; set; }

        /// <summary>
        /// Operator to apply when checking the current metric value against the threshold value.
        /// Accepted values are:
        /// - `GREATER_THAN`
        /// - `LESS_THAN`
        /// </summary>
        [Input("operator")]
        public Input<string>? Operator { get; set; }

        /// <summary>
        /// Value to test with the specified operator. If `field_name` is set to TYPE_NAME, you can match on the following values:
        /// - `PRIMARY`
        /// - `SECONDARY`
        /// - `STANDALONE`
        /// - `CONFIG`
        /// - `MONGOS`
        /// </summary>
        [Input("value")]
        public Input<string>? Value { get; set; }

        public AlertConfigurationMatcherGetArgs()
        {
        }
    }
}