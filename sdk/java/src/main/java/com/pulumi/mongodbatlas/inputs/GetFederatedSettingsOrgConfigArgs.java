// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.mongodbatlas.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class GetFederatedSettingsOrgConfigArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetFederatedSettingsOrgConfigArgs Empty = new GetFederatedSettingsOrgConfigArgs();

    /**
     * Unique 24-hexadecimal digit string that identifies the federated authentication configuration.
     * 
     */
    @Import(name="federationSettingsId", required=true)
    private Output<String> federationSettingsId;

    /**
     * @return Unique 24-hexadecimal digit string that identifies the federated authentication configuration.
     * 
     */
    public Output<String> federationSettingsId() {
        return this.federationSettingsId;
    }

    /**
     * Unique 24-hexadecimal digit string that identifies the organization that contains your projects.
     * 
     */
    @Import(name="orgId", required=true)
    private Output<String> orgId;

    /**
     * @return Unique 24-hexadecimal digit string that identifies the organization that contains your projects.
     * 
     */
    public Output<String> orgId() {
        return this.orgId;
    }

    private GetFederatedSettingsOrgConfigArgs() {}

    private GetFederatedSettingsOrgConfigArgs(GetFederatedSettingsOrgConfigArgs $) {
        this.federationSettingsId = $.federationSettingsId;
        this.orgId = $.orgId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetFederatedSettingsOrgConfigArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetFederatedSettingsOrgConfigArgs $;

        public Builder() {
            $ = new GetFederatedSettingsOrgConfigArgs();
        }

        public Builder(GetFederatedSettingsOrgConfigArgs defaults) {
            $ = new GetFederatedSettingsOrgConfigArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param federationSettingsId Unique 24-hexadecimal digit string that identifies the federated authentication configuration.
         * 
         * @return builder
         * 
         */
        public Builder federationSettingsId(Output<String> federationSettingsId) {
            $.federationSettingsId = federationSettingsId;
            return this;
        }

        /**
         * @param federationSettingsId Unique 24-hexadecimal digit string that identifies the federated authentication configuration.
         * 
         * @return builder
         * 
         */
        public Builder federationSettingsId(String federationSettingsId) {
            return federationSettingsId(Output.of(federationSettingsId));
        }

        /**
         * @param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects.
         * 
         * @return builder
         * 
         */
        public Builder orgId(Output<String> orgId) {
            $.orgId = orgId;
            return this;
        }

        /**
         * @param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects.
         * 
         * @return builder
         * 
         */
        public Builder orgId(String orgId) {
            return orgId(Output.of(orgId));
        }

        public GetFederatedSettingsOrgConfigArgs build() {
            $.federationSettingsId = Objects.requireNonNull($.federationSettingsId, "expected parameter 'federationSettingsId' to be non-null");
            $.orgId = Objects.requireNonNull($.orgId, "expected parameter 'orgId' to be non-null");
            return $;
        }
    }

}
