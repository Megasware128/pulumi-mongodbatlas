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
    public sealed class GetAdvancedClusterBiConnectorResult
    {
        /// <summary>
        /// Specifies whether or not BI Connector for Atlas is enabled on the cluster.l
        /// </summary>
        public readonly bool Enabled;
        /// <summary>
        /// Specifies the read preference to be used by BI Connector for Atlas on the cluster. Each BI Connector for Atlas read preference contains a distinct combination of [readPreference](https://docs.mongodb.com/manual/core/read-preference/) and [readPreferenceTags](https://docs.mongodb.com/manual/core/read-preference/#tag-sets) options. For details on BI Connector for Atlas read preferences, refer to the [BI Connector Read Preferences Table](https://docs.atlas.mongodb.com/tutorial/create-global-writes-cluster/#bic-read-preferences).
        /// </summary>
        public readonly string ReadPreference;

        [OutputConstructor]
        private GetAdvancedClusterBiConnectorResult(
            bool enabled,

            string readPreference)
        {
            Enabled = enabled;
            ReadPreference = readPreference;
        }
    }
}
