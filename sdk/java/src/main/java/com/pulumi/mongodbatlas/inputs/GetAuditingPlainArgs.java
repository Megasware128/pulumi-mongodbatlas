// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.mongodbatlas.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class GetAuditingPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetAuditingPlainArgs Empty = new GetAuditingPlainArgs();

    /**
     * The unique ID for the project to create the database user.
     * 
     */
    @Import(name="projectId", required=true)
    private String projectId;

    /**
     * @return The unique ID for the project to create the database user.
     * 
     */
    public String projectId() {
        return this.projectId;
    }

    private GetAuditingPlainArgs() {}

    private GetAuditingPlainArgs(GetAuditingPlainArgs $) {
        this.projectId = $.projectId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetAuditingPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetAuditingPlainArgs $;

        public Builder() {
            $ = new GetAuditingPlainArgs();
        }

        public Builder(GetAuditingPlainArgs defaults) {
            $ = new GetAuditingPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param projectId The unique ID for the project to create the database user.
         * 
         * @return builder
         * 
         */
        public Builder projectId(String projectId) {
            $.projectId = projectId;
            return this;
        }

        public GetAuditingPlainArgs build() {
            $.projectId = Objects.requireNonNull($.projectId, "expected parameter 'projectId' to be non-null");
            return $;
        }
    }

}
