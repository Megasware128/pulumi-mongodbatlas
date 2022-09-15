// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.mongodbatlas.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Double;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class AlertConfigurationThresholdConfigArgs extends com.pulumi.resources.ResourceArgs {

    public static final AlertConfigurationThresholdConfigArgs Empty = new AlertConfigurationThresholdConfigArgs();

    /**
     * Operator to apply when checking the current metric value against the threshold value.
     * Accepted values are:
     * - `GREATER_THAN`
     * - `LESS_THAN`
     * 
     */
    @Import(name="operator")
    private @Nullable Output<String> operator;

    /**
     * @return Operator to apply when checking the current metric value against the threshold value.
     * Accepted values are:
     * - `GREATER_THAN`
     * - `LESS_THAN`
     * 
     */
    public Optional<Output<String>> operator() {
        return Optional.ofNullable(this.operator);
    }

    /**
     * Threshold value outside of which an alert will be triggered.
     * 
     */
    @Import(name="threshold")
    private @Nullable Output<Double> threshold;

    /**
     * @return Threshold value outside of which an alert will be triggered.
     * 
     */
    public Optional<Output<Double>> threshold() {
        return Optional.ofNullable(this.threshold);
    }

    /**
     * The units for the threshold value. Depends on the type of metric.
     * Accepted values are:
     * - `RAW`
     * - `BITS`
     * - `BYTES`
     * - `KILOBITS`
     * - `KILOBYTES`
     * - `MEGABITS`
     * - `MEGABYTES`
     * - `GIGABITS`
     * - `GIGABYTES`
     * - `TERABYTES`
     * - `PETABYTES`
     * - `MILLISECONDS`
     * - `SECONDS`
     * - `MINUTES`
     * - `HOURS`
     * - `DAYS`
     * 
     */
    @Import(name="units")
    private @Nullable Output<String> units;

    /**
     * @return The units for the threshold value. Depends on the type of metric.
     * Accepted values are:
     * - `RAW`
     * - `BITS`
     * - `BYTES`
     * - `KILOBITS`
     * - `KILOBYTES`
     * - `MEGABITS`
     * - `MEGABYTES`
     * - `GIGABITS`
     * - `GIGABYTES`
     * - `TERABYTES`
     * - `PETABYTES`
     * - `MILLISECONDS`
     * - `SECONDS`
     * - `MINUTES`
     * - `HOURS`
     * - `DAYS`
     * 
     */
    public Optional<Output<String>> units() {
        return Optional.ofNullable(this.units);
    }

    private AlertConfigurationThresholdConfigArgs() {}

    private AlertConfigurationThresholdConfigArgs(AlertConfigurationThresholdConfigArgs $) {
        this.operator = $.operator;
        this.threshold = $.threshold;
        this.units = $.units;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(AlertConfigurationThresholdConfigArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private AlertConfigurationThresholdConfigArgs $;

        public Builder() {
            $ = new AlertConfigurationThresholdConfigArgs();
        }

        public Builder(AlertConfigurationThresholdConfigArgs defaults) {
            $ = new AlertConfigurationThresholdConfigArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param operator Operator to apply when checking the current metric value against the threshold value.
         * Accepted values are:
         * - `GREATER_THAN`
         * - `LESS_THAN`
         * 
         * @return builder
         * 
         */
        public Builder operator(@Nullable Output<String> operator) {
            $.operator = operator;
            return this;
        }

        /**
         * @param operator Operator to apply when checking the current metric value against the threshold value.
         * Accepted values are:
         * - `GREATER_THAN`
         * - `LESS_THAN`
         * 
         * @return builder
         * 
         */
        public Builder operator(String operator) {
            return operator(Output.of(operator));
        }

        /**
         * @param threshold Threshold value outside of which an alert will be triggered.
         * 
         * @return builder
         * 
         */
        public Builder threshold(@Nullable Output<Double> threshold) {
            $.threshold = threshold;
            return this;
        }

        /**
         * @param threshold Threshold value outside of which an alert will be triggered.
         * 
         * @return builder
         * 
         */
        public Builder threshold(Double threshold) {
            return threshold(Output.of(threshold));
        }

        /**
         * @param units The units for the threshold value. Depends on the type of metric.
         * Accepted values are:
         * - `RAW`
         * - `BITS`
         * - `BYTES`
         * - `KILOBITS`
         * - `KILOBYTES`
         * - `MEGABITS`
         * - `MEGABYTES`
         * - `GIGABITS`
         * - `GIGABYTES`
         * - `TERABYTES`
         * - `PETABYTES`
         * - `MILLISECONDS`
         * - `SECONDS`
         * - `MINUTES`
         * - `HOURS`
         * - `DAYS`
         * 
         * @return builder
         * 
         */
        public Builder units(@Nullable Output<String> units) {
            $.units = units;
            return this;
        }

        /**
         * @param units The units for the threshold value. Depends on the type of metric.
         * Accepted values are:
         * - `RAW`
         * - `BITS`
         * - `BYTES`
         * - `KILOBITS`
         * - `KILOBYTES`
         * - `MEGABITS`
         * - `MEGABYTES`
         * - `GIGABITS`
         * - `GIGABYTES`
         * - `TERABYTES`
         * - `PETABYTES`
         * - `MILLISECONDS`
         * - `SECONDS`
         * - `MINUTES`
         * - `HOURS`
         * - `DAYS`
         * 
         * @return builder
         * 
         */
        public Builder units(String units) {
            return units(Output.of(units));
        }

        public AlertConfigurationThresholdConfigArgs build() {
            return $;
        }
    }

}