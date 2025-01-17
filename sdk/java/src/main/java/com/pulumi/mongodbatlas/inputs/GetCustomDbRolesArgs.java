// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.mongodbatlas.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class GetCustomDbRolesArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetCustomDbRolesArgs Empty = new GetCustomDbRolesArgs();

    /**
     * The unique ID for the project to get all custom db roles.
     * 
     */
    @Import(name="projectId", required=true)
    private Output<String> projectId;

    /**
     * @return The unique ID for the project to get all custom db roles.
     * 
     */
    public Output<String> projectId() {
        return this.projectId;
    }

    private GetCustomDbRolesArgs() {}

    private GetCustomDbRolesArgs(GetCustomDbRolesArgs $) {
        this.projectId = $.projectId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetCustomDbRolesArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetCustomDbRolesArgs $;

        public Builder() {
            $ = new GetCustomDbRolesArgs();
        }

        public Builder(GetCustomDbRolesArgs defaults) {
            $ = new GetCustomDbRolesArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param projectId The unique ID for the project to get all custom db roles.
         * 
         * @return builder
         * 
         */
        public Builder projectId(Output<String> projectId) {
            $.projectId = projectId;
            return this;
        }

        /**
         * @param projectId The unique ID for the project to get all custom db roles.
         * 
         * @return builder
         * 
         */
        public Builder projectId(String projectId) {
            return projectId(Output.of(projectId));
        }

        public GetCustomDbRolesArgs build() {
            $.projectId = Objects.requireNonNull($.projectId, "expected parameter 'projectId' to be non-null");
            return $;
        }
    }

}
