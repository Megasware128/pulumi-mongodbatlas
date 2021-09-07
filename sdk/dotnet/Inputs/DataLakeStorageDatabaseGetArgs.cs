// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Mongodbatlas.Inputs
{

    public sealed class DataLakeStorageDatabaseGetArgs : Pulumi.ResourceArgs
    {
        [Input("collections")]
        private InputList<Inputs.DataLakeStorageDatabaseCollectionGetArgs>? _collections;
        public InputList<Inputs.DataLakeStorageDatabaseCollectionGetArgs> Collections
        {
            get => _collections ?? (_collections = new InputList<Inputs.DataLakeStorageDatabaseCollectionGetArgs>());
            set => _collections = value;
        }

        [Input("maxWildcardCollections")]
        public Input<int>? MaxWildcardCollections { get; set; }

        /// <summary>
        /// Name of the Atlas Data Lake.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("views")]
        private InputList<Inputs.DataLakeStorageDatabaseViewGetArgs>? _views;
        public InputList<Inputs.DataLakeStorageDatabaseViewGetArgs> Views
        {
            get => _views ?? (_views = new InputList<Inputs.DataLakeStorageDatabaseViewGetArgs>());
            set => _views = value;
        }

        public DataLakeStorageDatabaseGetArgs()
        {
        }
    }
}
