// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.mongodbatlas.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class GetLdapVerifyPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetLdapVerifyPlainArgs Empty = new GetLdapVerifyPlainArgs();

    /**
     * Unique identifier for the Atlas project associated with the verification request.
     * 
     */
    @Import(name="projectId", required=true)
    private String projectId;

    /**
     * @return Unique identifier for the Atlas project associated with the verification request.
     * 
     */
    public String projectId() {
        return this.projectId;
    }

    /**
     * Unique identifier of a request to verify an LDAP configuration.
     * 
     */
    @Import(name="requestId", required=true)
    private String requestId;

    /**
     * @return Unique identifier of a request to verify an LDAP configuration.
     * 
     */
    public String requestId() {
        return this.requestId;
    }

    private GetLdapVerifyPlainArgs() {}

    private GetLdapVerifyPlainArgs(GetLdapVerifyPlainArgs $) {
        this.projectId = $.projectId;
        this.requestId = $.requestId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetLdapVerifyPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetLdapVerifyPlainArgs $;

        public Builder() {
            $ = new GetLdapVerifyPlainArgs();
        }

        public Builder(GetLdapVerifyPlainArgs defaults) {
            $ = new GetLdapVerifyPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param projectId Unique identifier for the Atlas project associated with the verification request.
         * 
         * @return builder
         * 
         */
        public Builder projectId(String projectId) {
            $.projectId = projectId;
            return this;
        }

        /**
         * @param requestId Unique identifier of a request to verify an LDAP configuration.
         * 
         * @return builder
         * 
         */
        public Builder requestId(String requestId) {
            $.requestId = requestId;
            return this;
        }

        public GetLdapVerifyPlainArgs build() {
            $.projectId = Objects.requireNonNull($.projectId, "expected parameter 'projectId' to be non-null");
            $.requestId = Objects.requireNonNull($.requestId, "expected parameter 'requestId' to be non-null");
            return $;
        }
    }

}
