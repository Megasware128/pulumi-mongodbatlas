// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.mongodbatlas.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetProjectsResultApiKey {
    private String apiKeyId;
    private List<String> roleNames;

    private GetProjectsResultApiKey() {}
    public String apiKeyId() {
        return this.apiKeyId;
    }
    public List<String> roleNames() {
        return this.roleNames;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetProjectsResultApiKey defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String apiKeyId;
        private List<String> roleNames;
        public Builder() {}
        public Builder(GetProjectsResultApiKey defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.apiKeyId = defaults.apiKeyId;
    	      this.roleNames = defaults.roleNames;
        }

        @CustomType.Setter
        public Builder apiKeyId(String apiKeyId) {
            this.apiKeyId = Objects.requireNonNull(apiKeyId);
            return this;
        }
        @CustomType.Setter
        public Builder roleNames(List<String> roleNames) {
            this.roleNames = Objects.requireNonNull(roleNames);
            return this;
        }
        public Builder roleNames(String... roleNames) {
            return roleNames(List.of(roleNames));
        }
        public GetProjectsResultApiKey build() {
            final var o = new GetProjectsResultApiKey();
            o.apiKeyId = apiKeyId;
            o.roleNames = roleNames;
            return o;
        }
    }
}
