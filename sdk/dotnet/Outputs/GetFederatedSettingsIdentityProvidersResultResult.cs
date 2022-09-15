// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Mongodbatlas.Outputs
{

    [OutputType]
    public sealed class GetFederatedSettingsIdentityProvidersResultResult
    {
        /// <summary>
        /// Assertion consumer service URL to which the IdP sends the SAML response.
        /// </summary>
        public readonly string AcsUrl;
        /// <summary>
        /// List that contains the configured domains from which users can log in for this IdP.
        /// </summary>
        public readonly ImmutableArray<string> AssociatedDomains;
        /// <summary>
        /// List that contains the configured domains from which users can log in for this IdP.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetFederatedSettingsIdentityProvidersResultAssociatedOrgResult> AssociatedOrgs;
        /// <summary>
        /// Identifier for the intended audience of the SAML Assertion.
        /// </summary>
        public readonly string AudienceUri;
        /// <summary>
        /// Human-readable label that identifies the IdP.
        /// </summary>
        public readonly string DisplayName;
        /// <summary>
        /// Identifier for the issuer of the SAML Assertion.
        /// </summary>
        public readonly string IssuerUri;
        public readonly string OktaIdpId;
        public readonly ImmutableArray<Outputs.GetFederatedSettingsIdentityProvidersResultPemFileInfoResult> PemFileInfos;
        /// <summary>
        /// SAML Authentication Request Protocol binding used to send the AuthNRequest. Atlas supports the following binding values:
        /// - HTTP POST
        /// - HTTP REDIRECT
        /// </summary>
        public readonly string RequestBinding;
        /// <summary>
        /// Algorithm used to encrypt the IdP signature. Atlas supports the following signature algorithm values:
        /// - SHA-1
        /// - SHA-256
        /// </summary>
        public readonly string ResponseSignatureAlgorithm;
        /// <summary>
        /// Flag that indicates whether the IdP has enabled Bypass SAML Mode. Enabling this mode generates a URL that allows you bypass SAML and login to your organizations at any point. You can authenticate with this special URL only when Bypass Mode is enabled. Set this parameter to true during testing. This keeps you from getting locked out of MongoDB.
        /// </summary>
        public readonly bool SsoDebugEnabled;
        /// <summary>
        /// URL of the receiver of the SAML AuthNRequest.
        /// </summary>
        public readonly string SsoUrl;
        /// <summary>
        /// Label that indicates whether the identity provider is active. The IdP is Inactive until you map at least one domain to the IdP.
        /// </summary>
        public readonly string Status;

        [OutputConstructor]
        private GetFederatedSettingsIdentityProvidersResultResult(
            string acsUrl,

            ImmutableArray<string> associatedDomains,

            ImmutableArray<Outputs.GetFederatedSettingsIdentityProvidersResultAssociatedOrgResult> associatedOrgs,

            string audienceUri,

            string displayName,

            string issuerUri,

            string oktaIdpId,

            ImmutableArray<Outputs.GetFederatedSettingsIdentityProvidersResultPemFileInfoResult> pemFileInfos,

            string requestBinding,

            string responseSignatureAlgorithm,

            bool ssoDebugEnabled,

            string ssoUrl,

            string status)
        {
            AcsUrl = acsUrl;
            AssociatedDomains = associatedDomains;
            AssociatedOrgs = associatedOrgs;
            AudienceUri = audienceUri;
            DisplayName = displayName;
            IssuerUri = issuerUri;
            OktaIdpId = oktaIdpId;
            PemFileInfos = pemFileInfos;
            RequestBinding = requestBinding;
            ResponseSignatureAlgorithm = responseSignatureAlgorithm;
            SsoDebugEnabled = ssoDebugEnabled;
            SsoUrl = ssoUrl;
            Status = status;
        }
    }
}