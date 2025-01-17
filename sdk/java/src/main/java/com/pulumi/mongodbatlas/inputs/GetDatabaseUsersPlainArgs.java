// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.mongodbatlas.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class GetDatabaseUsersPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetDatabaseUsersPlainArgs Empty = new GetDatabaseUsersPlainArgs();

    /**
     * The unique ID for the project to get all database users.
     * 
     */
    @Import(name="projectId", required=true)
    private String projectId;

    /**
     * @return The unique ID for the project to get all database users.
     * 
     */
    public String projectId() {
        return this.projectId;
    }

    private GetDatabaseUsersPlainArgs() {}

    private GetDatabaseUsersPlainArgs(GetDatabaseUsersPlainArgs $) {
        this.projectId = $.projectId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetDatabaseUsersPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetDatabaseUsersPlainArgs $;

        public Builder() {
            $ = new GetDatabaseUsersPlainArgs();
        }

        public Builder(GetDatabaseUsersPlainArgs defaults) {
            $ = new GetDatabaseUsersPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param projectId The unique ID for the project to get all database users.
         * 
         * @return builder
         * 
         */
        public Builder projectId(String projectId) {
            $.projectId = projectId;
            return this;
        }

        public GetDatabaseUsersPlainArgs build() {
            $.projectId = Objects.requireNonNull($.projectId, "expected parameter 'projectId' to be non-null");
            return $;
        }
    }

}
