// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.mongodbatlas.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.mongodbatlas.outputs.GetCustomDbRolesResult;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetCustomDbRolesInvokeResult {
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    private String projectId;
    /**
     * @return A list where each represents a custom db roles.
     * 
     */
    private List<GetCustomDbRolesResult> results;

    private GetCustomDbRolesInvokeResult() {}
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    public String projectId() {
        return this.projectId;
    }
    /**
     * @return A list where each represents a custom db roles.
     * 
     */
    public List<GetCustomDbRolesResult> results() {
        return this.results;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetCustomDbRolesInvokeResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String id;
        private String projectId;
        private List<GetCustomDbRolesResult> results;
        public Builder() {}
        public Builder(GetCustomDbRolesInvokeResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.id = defaults.id;
    	      this.projectId = defaults.projectId;
    	      this.results = defaults.results;
        }

        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder projectId(String projectId) {
            this.projectId = Objects.requireNonNull(projectId);
            return this;
        }
        @CustomType.Setter
        public Builder results(List<GetCustomDbRolesResult> results) {
            this.results = Objects.requireNonNull(results);
            return this;
        }
        public Builder results(GetCustomDbRolesResult... results) {
            return results(List.of(results));
        }
        public GetCustomDbRolesInvokeResult build() {
            final var o = new GetCustomDbRolesInvokeResult();
            o.id = id;
            o.projectId = projectId;
            o.results = results;
            return o;
        }
    }
}