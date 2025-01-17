// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.mongodbatlas.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class GetFederatedSettingsIdentityProviderArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetFederatedSettingsIdentityProviderArgs Empty = new GetFederatedSettingsIdentityProviderArgs();

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
     * Unique 20-hexadecimal digit string that identifies the IdP.
     * 
     */
    @Import(name="identityProviderId", required=true)
    private Output<String> identityProviderId;

    /**
     * @return Unique 20-hexadecimal digit string that identifies the IdP.
     * 
     */
    public Output<String> identityProviderId() {
        return this.identityProviderId;
    }

    private GetFederatedSettingsIdentityProviderArgs() {}

    private GetFederatedSettingsIdentityProviderArgs(GetFederatedSettingsIdentityProviderArgs $) {
        this.federationSettingsId = $.federationSettingsId;
        this.identityProviderId = $.identityProviderId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetFederatedSettingsIdentityProviderArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetFederatedSettingsIdentityProviderArgs $;

        public Builder() {
            $ = new GetFederatedSettingsIdentityProviderArgs();
        }

        public Builder(GetFederatedSettingsIdentityProviderArgs defaults) {
            $ = new GetFederatedSettingsIdentityProviderArgs(Objects.requireNonNull(defaults));
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
         * @param identityProviderId Unique 20-hexadecimal digit string that identifies the IdP.
         * 
         * @return builder
         * 
         */
        public Builder identityProviderId(Output<String> identityProviderId) {
            $.identityProviderId = identityProviderId;
            return this;
        }

        /**
         * @param identityProviderId Unique 20-hexadecimal digit string that identifies the IdP.
         * 
         * @return builder
         * 
         */
        public Builder identityProviderId(String identityProviderId) {
            return identityProviderId(Output.of(identityProviderId));
        }

        public GetFederatedSettingsIdentityProviderArgs build() {
            $.federationSettingsId = Objects.requireNonNull($.federationSettingsId, "expected parameter 'federationSettingsId' to be non-null");
            $.identityProviderId = Objects.requireNonNull($.identityProviderId, "expected parameter 'identityProviderId' to be non-null");
            return $;
        }
    }

}
