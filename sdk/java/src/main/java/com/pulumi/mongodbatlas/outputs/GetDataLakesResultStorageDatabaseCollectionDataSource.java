// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.mongodbatlas.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetDataLakesResultStorageDatabaseCollectionDataSource {
    private String defaultFormat;
    private String path;
    private String storeName;

    private GetDataLakesResultStorageDatabaseCollectionDataSource() {}
    public String defaultFormat() {
        return this.defaultFormat;
    }
    public String path() {
        return this.path;
    }
    public String storeName() {
        return this.storeName;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetDataLakesResultStorageDatabaseCollectionDataSource defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String defaultFormat;
        private String path;
        private String storeName;
        public Builder() {}
        public Builder(GetDataLakesResultStorageDatabaseCollectionDataSource defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.defaultFormat = defaults.defaultFormat;
    	      this.path = defaults.path;
    	      this.storeName = defaults.storeName;
        }

        @CustomType.Setter
        public Builder defaultFormat(String defaultFormat) {
            this.defaultFormat = Objects.requireNonNull(defaultFormat);
            return this;
        }
        @CustomType.Setter
        public Builder path(String path) {
            this.path = Objects.requireNonNull(path);
            return this;
        }
        @CustomType.Setter
        public Builder storeName(String storeName) {
            this.storeName = Objects.requireNonNull(storeName);
            return this;
        }
        public GetDataLakesResultStorageDatabaseCollectionDataSource build() {
            final var o = new GetDataLakesResultStorageDatabaseCollectionDataSource();
            o.defaultFormat = defaultFormat;
            o.path = path;
            o.storeName = storeName;
            return o;
        }
    }
}