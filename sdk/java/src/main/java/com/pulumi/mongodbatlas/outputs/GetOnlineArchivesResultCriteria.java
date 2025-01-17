// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.mongodbatlas.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetOnlineArchivesResultCriteria {
    private String dateField;
    private String dateFormat;
    private Integer expireAfterDays;
    private String query;
    private String type;

    private GetOnlineArchivesResultCriteria() {}
    public String dateField() {
        return this.dateField;
    }
    public String dateFormat() {
        return this.dateFormat;
    }
    public Integer expireAfterDays() {
        return this.expireAfterDays;
    }
    public String query() {
        return this.query;
    }
    public String type() {
        return this.type;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetOnlineArchivesResultCriteria defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String dateField;
        private String dateFormat;
        private Integer expireAfterDays;
        private String query;
        private String type;
        public Builder() {}
        public Builder(GetOnlineArchivesResultCriteria defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.dateField = defaults.dateField;
    	      this.dateFormat = defaults.dateFormat;
    	      this.expireAfterDays = defaults.expireAfterDays;
    	      this.query = defaults.query;
    	      this.type = defaults.type;
        }

        @CustomType.Setter
        public Builder dateField(String dateField) {
            this.dateField = Objects.requireNonNull(dateField);
            return this;
        }
        @CustomType.Setter
        public Builder dateFormat(String dateFormat) {
            this.dateFormat = Objects.requireNonNull(dateFormat);
            return this;
        }
        @CustomType.Setter
        public Builder expireAfterDays(Integer expireAfterDays) {
            this.expireAfterDays = Objects.requireNonNull(expireAfterDays);
            return this;
        }
        @CustomType.Setter
        public Builder query(String query) {
            this.query = Objects.requireNonNull(query);
            return this;
        }
        @CustomType.Setter
        public Builder type(String type) {
            this.type = Objects.requireNonNull(type);
            return this;
        }
        public GetOnlineArchivesResultCriteria build() {
            final var o = new GetOnlineArchivesResultCriteria();
            o.dateField = dateField;
            o.dateFormat = dateFormat;
            o.expireAfterDays = expireAfterDays;
            o.query = query;
            o.type = type;
            return o;
        }
    }
}
