// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.mongodbatlas.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetAuditingResult {
    /**
     * @return JSON-formatted audit filter used by the project
     * 
     */
    private Boolean auditAuthorizationSuccess;
    /**
     * @return Indicates whether the auditing system captures successful authentication attempts for audit filters using the &#34;atype&#34; : &#34;authCheck&#34; auditing event. For more information, see auditAuthorizationSuccess
     * 
     */
    private String auditFilter;
    /**
     * @return Denotes the configuration method for the audit filter. Possible values are: NONE - auditing not configured for the project.m FILTER_BUILDER - auditing configured via Atlas UI filter builderm FILTER_JSON - auditing configured via Atlas custom filter or API.
     * 
     */
    private String configurationType;
    /**
     * @return Denotes whether or not the project associated with the {GROUP-ID} has database auditing enabled.
     * 
     */
    private Boolean enabled;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    private String projectId;

    private GetAuditingResult() {}
    /**
     * @return JSON-formatted audit filter used by the project
     * 
     */
    public Boolean auditAuthorizationSuccess() {
        return this.auditAuthorizationSuccess;
    }
    /**
     * @return Indicates whether the auditing system captures successful authentication attempts for audit filters using the &#34;atype&#34; : &#34;authCheck&#34; auditing event. For more information, see auditAuthorizationSuccess
     * 
     */
    public String auditFilter() {
        return this.auditFilter;
    }
    /**
     * @return Denotes the configuration method for the audit filter. Possible values are: NONE - auditing not configured for the project.m FILTER_BUILDER - auditing configured via Atlas UI filter builderm FILTER_JSON - auditing configured via Atlas custom filter or API.
     * 
     */
    public String configurationType() {
        return this.configurationType;
    }
    /**
     * @return Denotes whether or not the project associated with the {GROUP-ID} has database auditing enabled.
     * 
     */
    public Boolean enabled() {
        return this.enabled;
    }
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

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetAuditingResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private Boolean auditAuthorizationSuccess;
        private String auditFilter;
        private String configurationType;
        private Boolean enabled;
        private String id;
        private String projectId;
        public Builder() {}
        public Builder(GetAuditingResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.auditAuthorizationSuccess = defaults.auditAuthorizationSuccess;
    	      this.auditFilter = defaults.auditFilter;
    	      this.configurationType = defaults.configurationType;
    	      this.enabled = defaults.enabled;
    	      this.id = defaults.id;
    	      this.projectId = defaults.projectId;
        }

        @CustomType.Setter
        public Builder auditAuthorizationSuccess(Boolean auditAuthorizationSuccess) {
            this.auditAuthorizationSuccess = Objects.requireNonNull(auditAuthorizationSuccess);
            return this;
        }
        @CustomType.Setter
        public Builder auditFilter(String auditFilter) {
            this.auditFilter = Objects.requireNonNull(auditFilter);
            return this;
        }
        @CustomType.Setter
        public Builder configurationType(String configurationType) {
            this.configurationType = Objects.requireNonNull(configurationType);
            return this;
        }
        @CustomType.Setter
        public Builder enabled(Boolean enabled) {
            this.enabled = Objects.requireNonNull(enabled);
            return this;
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
        public GetAuditingResult build() {
            final var o = new GetAuditingResult();
            o.auditAuthorizationSuccess = auditAuthorizationSuccess;
            o.auditFilter = auditFilter;
            o.configurationType = configurationType;
            o.enabled = enabled;
            o.id = id;
            o.projectId = projectId;
            return o;
        }
    }
}
