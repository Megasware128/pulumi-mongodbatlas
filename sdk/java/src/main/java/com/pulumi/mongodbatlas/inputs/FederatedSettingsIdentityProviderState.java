// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.mongodbatlas.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class FederatedSettingsIdentityProviderState extends com.pulumi.resources.ResourceArgs {

    public static final FederatedSettingsIdentityProviderState Empty = new FederatedSettingsIdentityProviderState();

    /**
     * List that contains the domains associated with the identity provider.
     * 
     */
    @Import(name="associatedDomains")
    private @Nullable Output<List<String>> associatedDomains;

    /**
     * @return List that contains the domains associated with the identity provider.
     * 
     */
    public Optional<Output<List<String>>> associatedDomains() {
        return Optional.ofNullable(this.associatedDomains);
    }

    /**
     * Unique 24-hexadecimal digit string that identifies the federated authentication configuration.
     * 
     */
    @Import(name="federationSettingsId")
    private @Nullable Output<String> federationSettingsId;

    /**
     * @return Unique 24-hexadecimal digit string that identifies the federated authentication configuration.
     * 
     */
    public Optional<Output<String>> federationSettingsId() {
        return Optional.ofNullable(this.federationSettingsId);
    }

    /**
     * Unique string that identifies the issuer of the SAML
     * 
     */
    @Import(name="issuerUri")
    private @Nullable Output<String> issuerUri;

    /**
     * @return Unique string that identifies the issuer of the SAML
     * 
     */
    public Optional<Output<String>> issuerUri() {
        return Optional.ofNullable(this.issuerUri);
    }

    /**
     * Human-readable label that identifies the identity provider.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Human-readable label that identifies the identity provider.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * Unique 20-hexadecimal digit string that identifies the IdP.
     * 
     */
    @Import(name="oktaIdpId")
    private @Nullable Output<String> oktaIdpId;

    /**
     * @return Unique 20-hexadecimal digit string that identifies the IdP.
     * 
     */
    public Optional<Output<String>> oktaIdpId() {
        return Optional.ofNullable(this.oktaIdpId);
    }

    /**
     * SAML Authentication Request Protocol HTTP method binding (POST or REDIRECT) that Federated Authentication uses to send the authentication request. Atlas supports the following binding values:
     * - HTTP POST
     * - HTTP REDIRECT
     * 
     */
    @Import(name="requestBinding")
    private @Nullable Output<String> requestBinding;

    /**
     * @return SAML Authentication Request Protocol HTTP method binding (POST or REDIRECT) that Federated Authentication uses to send the authentication request. Atlas supports the following binding values:
     * - HTTP POST
     * - HTTP REDIRECT
     * 
     */
    public Optional<Output<String>> requestBinding() {
        return Optional.ofNullable(this.requestBinding);
    }

    /**
     * Signature algorithm that Federated Authentication uses to encrypt the identity provider signature.  Valid values include SHA-1 and SHA-256.
     * 
     */
    @Import(name="responseSignatureAlgorithm")
    private @Nullable Output<String> responseSignatureAlgorithm;

    /**
     * @return Signature algorithm that Federated Authentication uses to encrypt the identity provider signature.  Valid values include SHA-1 and SHA-256.
     * 
     */
    public Optional<Output<String>> responseSignatureAlgorithm() {
        return Optional.ofNullable(this.responseSignatureAlgorithm);
    }

    /**
     * Flag that indicates whether the identity provider has SSO debug enabled.
     * 
     */
    @Import(name="ssoDebugEnabled")
    private @Nullable Output<Boolean> ssoDebugEnabled;

    /**
     * @return Flag that indicates whether the identity provider has SSO debug enabled.
     * 
     */
    public Optional<Output<Boolean>> ssoDebugEnabled() {
        return Optional.ofNullable(this.ssoDebugEnabled);
    }

    /**
     * Unique string that identifies the intended audience of the SAML assertion.
     * 
     */
    @Import(name="ssoUrl")
    private @Nullable Output<String> ssoUrl;

    /**
     * @return Unique string that identifies the intended audience of the SAML assertion.
     * 
     */
    public Optional<Output<String>> ssoUrl() {
        return Optional.ofNullable(this.ssoUrl);
    }

    /**
     * String enum that indicates whether the identity provider is active or not. Accepted values are ACTIVE or INACTIVE.
     * 
     */
    @Import(name="status")
    private @Nullable Output<String> status;

    /**
     * @return String enum that indicates whether the identity provider is active or not. Accepted values are ACTIVE or INACTIVE.
     * 
     */
    public Optional<Output<String>> status() {
        return Optional.ofNullable(this.status);
    }

    private FederatedSettingsIdentityProviderState() {}

    private FederatedSettingsIdentityProviderState(FederatedSettingsIdentityProviderState $) {
        this.associatedDomains = $.associatedDomains;
        this.federationSettingsId = $.federationSettingsId;
        this.issuerUri = $.issuerUri;
        this.name = $.name;
        this.oktaIdpId = $.oktaIdpId;
        this.requestBinding = $.requestBinding;
        this.responseSignatureAlgorithm = $.responseSignatureAlgorithm;
        this.ssoDebugEnabled = $.ssoDebugEnabled;
        this.ssoUrl = $.ssoUrl;
        this.status = $.status;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(FederatedSettingsIdentityProviderState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private FederatedSettingsIdentityProviderState $;

        public Builder() {
            $ = new FederatedSettingsIdentityProviderState();
        }

        public Builder(FederatedSettingsIdentityProviderState defaults) {
            $ = new FederatedSettingsIdentityProviderState(Objects.requireNonNull(defaults));
        }

        /**
         * @param associatedDomains List that contains the domains associated with the identity provider.
         * 
         * @return builder
         * 
         */
        public Builder associatedDomains(@Nullable Output<List<String>> associatedDomains) {
            $.associatedDomains = associatedDomains;
            return this;
        }

        /**
         * @param associatedDomains List that contains the domains associated with the identity provider.
         * 
         * @return builder
         * 
         */
        public Builder associatedDomains(List<String> associatedDomains) {
            return associatedDomains(Output.of(associatedDomains));
        }

        /**
         * @param associatedDomains List that contains the domains associated with the identity provider.
         * 
         * @return builder
         * 
         */
        public Builder associatedDomains(String... associatedDomains) {
            return associatedDomains(List.of(associatedDomains));
        }

        /**
         * @param federationSettingsId Unique 24-hexadecimal digit string that identifies the federated authentication configuration.
         * 
         * @return builder
         * 
         */
        public Builder federationSettingsId(@Nullable Output<String> federationSettingsId) {
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
         * @param issuerUri Unique string that identifies the issuer of the SAML
         * 
         * @return builder
         * 
         */
        public Builder issuerUri(@Nullable Output<String> issuerUri) {
            $.issuerUri = issuerUri;
            return this;
        }

        /**
         * @param issuerUri Unique string that identifies the issuer of the SAML
         * 
         * @return builder
         * 
         */
        public Builder issuerUri(String issuerUri) {
            return issuerUri(Output.of(issuerUri));
        }

        /**
         * @param name Human-readable label that identifies the identity provider.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Human-readable label that identifies the identity provider.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param oktaIdpId Unique 20-hexadecimal digit string that identifies the IdP.
         * 
         * @return builder
         * 
         */
        public Builder oktaIdpId(@Nullable Output<String> oktaIdpId) {
            $.oktaIdpId = oktaIdpId;
            return this;
        }

        /**
         * @param oktaIdpId Unique 20-hexadecimal digit string that identifies the IdP.
         * 
         * @return builder
         * 
         */
        public Builder oktaIdpId(String oktaIdpId) {
            return oktaIdpId(Output.of(oktaIdpId));
        }

        /**
         * @param requestBinding SAML Authentication Request Protocol HTTP method binding (POST or REDIRECT) that Federated Authentication uses to send the authentication request. Atlas supports the following binding values:
         * - HTTP POST
         * - HTTP REDIRECT
         * 
         * @return builder
         * 
         */
        public Builder requestBinding(@Nullable Output<String> requestBinding) {
            $.requestBinding = requestBinding;
            return this;
        }

        /**
         * @param requestBinding SAML Authentication Request Protocol HTTP method binding (POST or REDIRECT) that Federated Authentication uses to send the authentication request. Atlas supports the following binding values:
         * - HTTP POST
         * - HTTP REDIRECT
         * 
         * @return builder
         * 
         */
        public Builder requestBinding(String requestBinding) {
            return requestBinding(Output.of(requestBinding));
        }

        /**
         * @param responseSignatureAlgorithm Signature algorithm that Federated Authentication uses to encrypt the identity provider signature.  Valid values include SHA-1 and SHA-256.
         * 
         * @return builder
         * 
         */
        public Builder responseSignatureAlgorithm(@Nullable Output<String> responseSignatureAlgorithm) {
            $.responseSignatureAlgorithm = responseSignatureAlgorithm;
            return this;
        }

        /**
         * @param responseSignatureAlgorithm Signature algorithm that Federated Authentication uses to encrypt the identity provider signature.  Valid values include SHA-1 and SHA-256.
         * 
         * @return builder
         * 
         */
        public Builder responseSignatureAlgorithm(String responseSignatureAlgorithm) {
            return responseSignatureAlgorithm(Output.of(responseSignatureAlgorithm));
        }

        /**
         * @param ssoDebugEnabled Flag that indicates whether the identity provider has SSO debug enabled.
         * 
         * @return builder
         * 
         */
        public Builder ssoDebugEnabled(@Nullable Output<Boolean> ssoDebugEnabled) {
            $.ssoDebugEnabled = ssoDebugEnabled;
            return this;
        }

        /**
         * @param ssoDebugEnabled Flag that indicates whether the identity provider has SSO debug enabled.
         * 
         * @return builder
         * 
         */
        public Builder ssoDebugEnabled(Boolean ssoDebugEnabled) {
            return ssoDebugEnabled(Output.of(ssoDebugEnabled));
        }

        /**
         * @param ssoUrl Unique string that identifies the intended audience of the SAML assertion.
         * 
         * @return builder
         * 
         */
        public Builder ssoUrl(@Nullable Output<String> ssoUrl) {
            $.ssoUrl = ssoUrl;
            return this;
        }

        /**
         * @param ssoUrl Unique string that identifies the intended audience of the SAML assertion.
         * 
         * @return builder
         * 
         */
        public Builder ssoUrl(String ssoUrl) {
            return ssoUrl(Output.of(ssoUrl));
        }

        /**
         * @param status String enum that indicates whether the identity provider is active or not. Accepted values are ACTIVE or INACTIVE.
         * 
         * @return builder
         * 
         */
        public Builder status(@Nullable Output<String> status) {
            $.status = status;
            return this;
        }

        /**
         * @param status String enum that indicates whether the identity provider is active or not. Accepted values are ACTIVE or INACTIVE.
         * 
         * @return builder
         * 
         */
        public Builder status(String status) {
            return status(Output.of(status));
        }

        public FederatedSettingsIdentityProviderState build() {
            return $;
        }
    }

}
