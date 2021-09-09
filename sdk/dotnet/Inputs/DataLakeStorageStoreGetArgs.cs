// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Mongodbatlas.Inputs
{

    public sealed class DataLakeStorageStoreGetArgs : Pulumi.ResourceArgs
    {
        [Input("additionalStorageClasses")]
        private InputList<string>? _additionalStorageClasses;
        public InputList<string> AdditionalStorageClasses
        {
            get => _additionalStorageClasses ?? (_additionalStorageClasses = new InputList<string>());
            set => _additionalStorageClasses = value;
        }

        [Input("bucket")]
        public Input<string>? Bucket { get; set; }

        [Input("delimiter")]
        public Input<string>? Delimiter { get; set; }

        [Input("includeTags")]
        public Input<bool>? IncludeTags { get; set; }

        /// <summary>
        /// Name of the Atlas Data Lake.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("prefix")]
        public Input<string>? Prefix { get; set; }

        [Input("provider")]
        public Input<string>? Provider { get; set; }

        [Input("region")]
        public Input<string>? Region { get; set; }

        public DataLakeStorageStoreGetArgs()
        {
        }
    }
}