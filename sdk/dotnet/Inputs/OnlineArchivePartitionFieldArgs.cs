// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Mongodbatlas.Inputs
{

    public sealed class OnlineArchivePartitionFieldArgs : Pulumi.ResourceArgs
    {
        [Input("fieldName", required: true)]
        public Input<string> FieldName { get; set; } = null!;

        [Input("fieldType")]
        public Input<string>? FieldType { get; set; }

        [Input("order", required: true)]
        public Input<int> Order { get; set; } = null!;

        public OnlineArchivePartitionFieldArgs()
        {
        }
    }
}