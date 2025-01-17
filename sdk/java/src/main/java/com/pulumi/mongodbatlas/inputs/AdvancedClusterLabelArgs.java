// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.mongodbatlas.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class AdvancedClusterLabelArgs extends com.pulumi.resources.ResourceArgs {

    public static final AdvancedClusterLabelArgs Empty = new AdvancedClusterLabelArgs();

    /**
     * The key that you want to write.
     * 
     */
    @Import(name="key")
    private @Nullable Output<String> key;

    /**
     * @return The key that you want to write.
     * 
     */
    public Optional<Output<String>> key() {
        return Optional.ofNullable(this.key);
    }

    /**
     * The value that you want to write.
     * 
     */
    @Import(name="value")
    private @Nullable Output<String> value;

    /**
     * @return The value that you want to write.
     * 
     */
    public Optional<Output<String>> value() {
        return Optional.ofNullable(this.value);
    }

    private AdvancedClusterLabelArgs() {}

    private AdvancedClusterLabelArgs(AdvancedClusterLabelArgs $) {
        this.key = $.key;
        this.value = $.value;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(AdvancedClusterLabelArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private AdvancedClusterLabelArgs $;

        public Builder() {
            $ = new AdvancedClusterLabelArgs();
        }

        public Builder(AdvancedClusterLabelArgs defaults) {
            $ = new AdvancedClusterLabelArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param key The key that you want to write.
         * 
         * @return builder
         * 
         */
        public Builder key(@Nullable Output<String> key) {
            $.key = key;
            return this;
        }

        /**
         * @param key The key that you want to write.
         * 
         * @return builder
         * 
         */
        public Builder key(String key) {
            return key(Output.of(key));
        }

        /**
         * @param value The value that you want to write.
         * 
         * @return builder
         * 
         */
        public Builder value(@Nullable Output<String> value) {
            $.value = value;
            return this;
        }

        /**
         * @param value The value that you want to write.
         * 
         * @return builder
         * 
         */
        public Builder value(String value) {
            return value(Output.of(value));
        }

        public AdvancedClusterLabelArgs build() {
            return $;
        }
    }

}
