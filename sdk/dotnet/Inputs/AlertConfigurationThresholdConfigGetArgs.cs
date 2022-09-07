// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Mongodbatlas.Inputs
{

    public sealed class AlertConfigurationThresholdConfigGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Operator to apply when checking the current metric value against the threshold value.
        /// Accepted values are:
        /// - `GREATER_THAN`
        /// - `LESS_THAN`
        /// </summary>
        [Input("operator")]
        public Input<string>? Operator { get; set; }

        /// <summary>
        /// Threshold value outside of which an alert will be triggered.
        /// </summary>
        [Input("threshold")]
        public Input<double>? Threshold { get; set; }

        /// <summary>
        /// The units for the threshold value. Depends on the type of metric.
        /// Accepted values are:
        /// - `RAW`
        /// - `BITS`
        /// - `BYTES`
        /// - `KILOBITS`
        /// - `KILOBYTES`
        /// - `MEGABITS`
        /// - `MEGABYTES`
        /// - `GIGABITS`
        /// - `GIGABYTES`
        /// - `TERABYTES`
        /// - `PETABYTES`
        /// - `MILLISECONDS`
        /// - `SECONDS`
        /// - `MINUTES`
        /// - `HOURS`
        /// - `DAYS`
        /// </summary>
        [Input("units")]
        public Input<string>? Units { get; set; }

        public AlertConfigurationThresholdConfigGetArgs()
        {
        }
        public static new AlertConfigurationThresholdConfigGetArgs Empty => new AlertConfigurationThresholdConfigGetArgs();
    }
}
