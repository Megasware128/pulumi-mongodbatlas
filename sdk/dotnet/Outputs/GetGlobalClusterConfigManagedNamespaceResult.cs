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
    public sealed class GetGlobalClusterConfigManagedNamespaceResult
    {
        /// <summary>
        /// (Required) The name of the collection associated with the managed namespace.
        /// </summary>
        public readonly string Collection;
        /// <summary>
        /// (Required)	The custom shard key for the collection. Global Clusters require a compound shard key consisting of a location field and a user-selected second key, the custom shard key.
        /// </summary>
        public readonly string CustomShardKey;
        /// <summary>
        /// (Required) The name of the database containing the collection.
        /// </summary>
        public readonly string Db;
        /// <summary>
        /// Specifies whether the custom shard key for the collection is [hashed](https://docs.mongodb.com/manual/reference/method/sh.shardCollection/#hashed-shard-keys). If omitted, defaults to `false`. If `false`, Atlas uses [ranged sharding](https://docs.mongodb.com/manual/core/ranged-sharding/). This is only available for Atlas clusters with MongoDB v4.4 and later.
        /// </summary>
        public readonly bool IsCustomShardKeyHashed;
        /// <summary>
        /// Specifies whether the underlying index enforces a unique constraint. If omitted, defaults to false. You cannot specify true when using [hashed shard keys](https://docs.mongodb.com/manual/core/hashed-sharding/#std-label-sharding-hashed).
        /// </summary>
        public readonly bool IsShardKeyUnique;

        [OutputConstructor]
        private GetGlobalClusterConfigManagedNamespaceResult(
            string collection,

            string customShardKey,

            string db,

            bool isCustomShardKeyHashed,

            bool isShardKeyUnique)
        {
            Collection = collection;
            CustomShardKey = customShardKey;
            Db = db;
            IsCustomShardKeyHashed = isCustomShardKeyHashed;
            IsShardKeyUnique = isShardKeyUnique;
        }
    }
}
