// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.mongodbatlas.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class GetThirdPartyIntegrationsArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetThirdPartyIntegrationsArgs Empty = new GetThirdPartyIntegrationsArgs();

    /**
     * The unique ID for the project to get all Third-Party service integrations
     * 
     */
    @Import(name="projectId", required=true)
    private Output<String> projectId;

    /**
     * @return The unique ID for the project to get all Third-Party service integrations
     * 
     */
    public Output<String> projectId() {
        return this.projectId;
    }

    private GetThirdPartyIntegrationsArgs() {}

    private GetThirdPartyIntegrationsArgs(GetThirdPartyIntegrationsArgs $) {
        this.projectId = $.projectId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetThirdPartyIntegrationsArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetThirdPartyIntegrationsArgs $;

        public Builder() {
            $ = new GetThirdPartyIntegrationsArgs();
        }

        public Builder(GetThirdPartyIntegrationsArgs defaults) {
            $ = new GetThirdPartyIntegrationsArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param projectId The unique ID for the project to get all Third-Party service integrations
         * 
         * @return builder
         * 
         */
        public Builder projectId(Output<String> projectId) {
            $.projectId = projectId;
            return this;
        }

        /**
         * @param projectId The unique ID for the project to get all Third-Party service integrations
         * 
         * @return builder
         * 
         */
        public Builder projectId(String projectId) {
            return projectId(Output.of(projectId));
        }

        public GetThirdPartyIntegrationsArgs build() {
            $.projectId = Objects.requireNonNull($.projectId, "expected parameter 'projectId' to be non-null");
            return $;
        }
    }

}
