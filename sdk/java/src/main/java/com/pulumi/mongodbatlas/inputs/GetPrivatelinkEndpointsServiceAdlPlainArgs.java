// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.mongodbatlas.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class GetPrivatelinkEndpointsServiceAdlPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetPrivatelinkEndpointsServiceAdlPlainArgs Empty = new GetPrivatelinkEndpointsServiceAdlPlainArgs();

    /**
     * The unique ID for the project.
     * 
     */
    @Import(name="projectId", required=true)
    private String projectId;

    /**
     * @return The unique ID for the project.
     * 
     */
    public String projectId() {
        return this.projectId;
    }

    private GetPrivatelinkEndpointsServiceAdlPlainArgs() {}

    private GetPrivatelinkEndpointsServiceAdlPlainArgs(GetPrivatelinkEndpointsServiceAdlPlainArgs $) {
        this.projectId = $.projectId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetPrivatelinkEndpointsServiceAdlPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetPrivatelinkEndpointsServiceAdlPlainArgs $;

        public Builder() {
            $ = new GetPrivatelinkEndpointsServiceAdlPlainArgs();
        }

        public Builder(GetPrivatelinkEndpointsServiceAdlPlainArgs defaults) {
            $ = new GetPrivatelinkEndpointsServiceAdlPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param projectId The unique ID for the project.
         * 
         * @return builder
         * 
         */
        public Builder projectId(String projectId) {
            $.projectId = projectId;
            return this;
        }

        public GetPrivatelinkEndpointsServiceAdlPlainArgs build() {
            $.projectId = Objects.requireNonNull($.projectId, "expected parameter 'projectId' to be non-null");
            return $;
        }
    }

}
