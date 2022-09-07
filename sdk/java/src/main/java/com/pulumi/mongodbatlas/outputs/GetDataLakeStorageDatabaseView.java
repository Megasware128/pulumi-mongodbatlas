// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.mongodbatlas.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetDataLakeStorageDatabaseView {
    /**
     * @return Name of the data lake.
     * 
     */
    private String name;
    private String pipeline;
    private String source;

    private GetDataLakeStorageDatabaseView() {}
    /**
     * @return Name of the data lake.
     * 
     */
    public String name() {
        return this.name;
    }
    public String pipeline() {
        return this.pipeline;
    }
    public String source() {
        return this.source;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetDataLakeStorageDatabaseView defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String name;
        private String pipeline;
        private String source;
        public Builder() {}
        public Builder(GetDataLakeStorageDatabaseView defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.name = defaults.name;
    	      this.pipeline = defaults.pipeline;
    	      this.source = defaults.source;
        }

        @CustomType.Setter
        public Builder name(String name) {
            this.name = Objects.requireNonNull(name);
            return this;
        }
        @CustomType.Setter
        public Builder pipeline(String pipeline) {
            this.pipeline = Objects.requireNonNull(pipeline);
            return this;
        }
        @CustomType.Setter
        public Builder source(String source) {
            this.source = Objects.requireNonNull(source);
            return this;
        }
        public GetDataLakeStorageDatabaseView build() {
            final var o = new GetDataLakeStorageDatabaseView();
            o.name = name;
            o.pipeline = pipeline;
            o.source = source;
            return o;
        }
    }
}
