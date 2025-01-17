// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.mongodbatlas.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetAlertConfigurationMatcher {
    /**
     * @return Name of the field in the target object to match on.
     * 
     */
    private String fieldName;
    /**
     * @return Operator to apply when checking the current metric value against the threshold value.
     * Accepted values are:
     * - `GREATER_THAN`
     * - `LESS_THAN`
     * 
     */
    private String operator;
    /**
     * @return Value to test with the specified operator. If `field_name` is set to TYPE_NAME, you can match on the following values:
     * - `PRIMARY`
     * - `SECONDARY`
     * - `STANDALONE`
     * - `CONFIG`
     * - `MONGOS`
     * 
     */
    private String value;

    private GetAlertConfigurationMatcher() {}
    /**
     * @return Name of the field in the target object to match on.
     * 
     */
    public String fieldName() {
        return this.fieldName;
    }
    /**
     * @return Operator to apply when checking the current metric value against the threshold value.
     * Accepted values are:
     * - `GREATER_THAN`
     * - `LESS_THAN`
     * 
     */
    public String operator() {
        return this.operator;
    }
    /**
     * @return Value to test with the specified operator. If `field_name` is set to TYPE_NAME, you can match on the following values:
     * - `PRIMARY`
     * - `SECONDARY`
     * - `STANDALONE`
     * - `CONFIG`
     * - `MONGOS`
     * 
     */
    public String value() {
        return this.value;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetAlertConfigurationMatcher defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String fieldName;
        private String operator;
        private String value;
        public Builder() {}
        public Builder(GetAlertConfigurationMatcher defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.fieldName = defaults.fieldName;
    	      this.operator = defaults.operator;
    	      this.value = defaults.value;
        }

        @CustomType.Setter
        public Builder fieldName(String fieldName) {
            this.fieldName = Objects.requireNonNull(fieldName);
            return this;
        }
        @CustomType.Setter
        public Builder operator(String operator) {
            this.operator = Objects.requireNonNull(operator);
            return this;
        }
        @CustomType.Setter
        public Builder value(String value) {
            this.value = Objects.requireNonNull(value);
            return this;
        }
        public GetAlertConfigurationMatcher build() {
            final var o = new GetAlertConfigurationMatcher();
            o.fieldName = fieldName;
            o.operator = operator;
            o.value = value;
            return o;
        }
    }
}
