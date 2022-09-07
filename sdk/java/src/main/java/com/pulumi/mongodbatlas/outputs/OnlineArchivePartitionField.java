// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.mongodbatlas.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class OnlineArchivePartitionField {
    private String fieldName;
    private @Nullable String fieldType;
    private Integer order;

    private OnlineArchivePartitionField() {}
    public String fieldName() {
        return this.fieldName;
    }
    public Optional<String> fieldType() {
        return Optional.ofNullable(this.fieldType);
    }
    public Integer order() {
        return this.order;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(OnlineArchivePartitionField defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String fieldName;
        private @Nullable String fieldType;
        private Integer order;
        public Builder() {}
        public Builder(OnlineArchivePartitionField defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.fieldName = defaults.fieldName;
    	      this.fieldType = defaults.fieldType;
    	      this.order = defaults.order;
        }

        @CustomType.Setter
        public Builder fieldName(String fieldName) {
            this.fieldName = Objects.requireNonNull(fieldName);
            return this;
        }
        @CustomType.Setter
        public Builder fieldType(@Nullable String fieldType) {
            this.fieldType = fieldType;
            return this;
        }
        @CustomType.Setter
        public Builder order(Integer order) {
            this.order = Objects.requireNonNull(order);
            return this;
        }
        public OnlineArchivePartitionField build() {
            final var o = new OnlineArchivePartitionField();
            o.fieldName = fieldName;
            o.fieldType = fieldType;
            o.order = order;
            return o;
        }
    }
}
