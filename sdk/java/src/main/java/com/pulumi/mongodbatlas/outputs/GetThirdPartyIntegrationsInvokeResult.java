// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.mongodbatlas.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.mongodbatlas.outputs.GetThirdPartyIntegrationsResult;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetThirdPartyIntegrationsInvokeResult {
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return (Required) ID of the Atlas project the Third-Party Service Integration belongs to.
     * 
     */
    private String projectId;
    /**
     * @return A list where each represents a Third-Party service integration.
     * 
     */
    private List<GetThirdPartyIntegrationsResult> results;

    private GetThirdPartyIntegrationsInvokeResult() {}
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return (Required) ID of the Atlas project the Third-Party Service Integration belongs to.
     * 
     */
    public String projectId() {
        return this.projectId;
    }
    /**
     * @return A list where each represents a Third-Party service integration.
     * 
     */
    public List<GetThirdPartyIntegrationsResult> results() {
        return this.results;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetThirdPartyIntegrationsInvokeResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String id;
        private String projectId;
        private List<GetThirdPartyIntegrationsResult> results;
        public Builder() {}
        public Builder(GetThirdPartyIntegrationsInvokeResult defaults) {
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
        public Builder results(List<GetThirdPartyIntegrationsResult> results) {
            this.results = Objects.requireNonNull(results);
            return this;
        }
        public Builder results(GetThirdPartyIntegrationsResult... results) {
            return results(List.of(results));
        }
        public GetThirdPartyIntegrationsInvokeResult build() {
            final var o = new GetThirdPartyIntegrationsInvokeResult();
            o.id = id;
            o.projectId = projectId;
            o.results = results;
            return o;
        }
    }
}
