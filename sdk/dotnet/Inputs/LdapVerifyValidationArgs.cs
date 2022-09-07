// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Mongodbatlas.Inputs
{

    public sealed class LdapVerifyValidationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The current status of the LDAP over TLS/SSL configuration. One of the following values: `PENDING`, `SUCCESS`, and `FAILED`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        [Input("validationType")]
        public Input<string>? ValidationType { get; set; }

        public LdapVerifyValidationArgs()
        {
        }
        public static new LdapVerifyValidationArgs Empty => new LdapVerifyValidationArgs();
    }
}
