// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Mongodbatlas
{
    /// <summary>
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Mongodbatlas = Pulumi.Mongodbatlas;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var testFlowdock = new Mongodbatlas.ThirdPartyIntegration("testFlowdock", new Mongodbatlas.ThirdPartyIntegrationArgs
    ///         {
    ///             ApiToken = "&lt;API-TOKEN&gt;",
    ///             FlowName = "&lt;FLOW-NAME&gt;",
    ///             OrgName = "&lt;ORG-NAME&gt;",
    ///             ProjectId = "&lt;PROJECT-ID&gt;",
    ///             Type = "FLOWDOCK",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Third-Party Integration Settings can be imported using project ID and the integration type, in the format `project_id`-`type`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import mongodbatlas:index/thirdPartyIntegration:ThirdPartyIntegration my_user 1112222b3bf99403840e8934-OPS_GENIE
    /// ```
    /// 
    ///  See [MongoDB Atlas API](https://docs.atlas.mongodb.com/reference/api/third-party-integration-settings-create/) Documentation for more information.
    /// </summary>
    [MongodbatlasResourceType("mongodbatlas:index/thirdPartyIntegration:ThirdPartyIntegration")]
    public partial class ThirdPartyIntegration : Pulumi.CustomResource
    {
        /// <summary>
        /// Unique identifier of your New Relic account.
        /// </summary>
        [Output("accountId")]
        public Output<string?> AccountId { get; private set; } = null!;

        /// <summary>
        /// Your API Key.
        /// </summary>
        [Output("apiKey")]
        public Output<string?> ApiKey { get; private set; } = null!;

        /// <summary>
        /// Your API Token.
        /// </summary>
        [Output("apiToken")]
        public Output<string?> ApiToken { get; private set; } = null!;

        [Output("channelName")]
        public Output<string?> ChannelName { get; private set; } = null!;

        /// <summary>
        /// Whether your cluster has Prometheus enabled.
        /// </summary>
        [Output("enabled")]
        public Output<bool?> Enabled { get; private set; } = null!;

        /// <summary>
        /// Your Flowdock Flow name.
        /// </summary>
        [Output("flowName")]
        public Output<string?> FlowName { get; private set; } = null!;

        /// <summary>
        /// Your License Key.
        /// </summary>
        [Output("licenseKey")]
        public Output<string?> LicenseKey { get; private set; } = null!;

        /// <summary>
        /// Your Microsoft Teams incoming webhook URL.
        /// * `PROMETHEUS`
        /// </summary>
        [Output("microsoftTeamsWebhookUrl")]
        public Output<string?> MicrosoftTeamsWebhookUrl { get; private set; } = null!;

        /// <summary>
        /// Your Flowdock organization name.
        /// * `WEBHOOK`
        /// </summary>
        [Output("orgName")]
        public Output<string?> OrgName { get; private set; } = null!;

        /// <summary>
        /// Your Prometheus password.
        /// </summary>
        [Output("password")]
        public Output<string?> Password { get; private set; } = null!;

        /// <summary>
        /// The unique ID for the project to get all Third-Party service integrations
        /// </summary>
        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;

        /// <summary>
        /// Your Insights Query Key.
        /// * `OPS_GENIE`
        /// </summary>
        [Output("readToken")]
        public Output<string?> ReadToken { get; private set; } = null!;

        /// <summary>
        /// Indicates which API URL to use, either US or EU. Opsgenie will use US by default.
        /// * `VICTOR_OPS`
        /// </summary>
        [Output("region")]
        public Output<string?> Region { get; private set; } = null!;

        /// <summary>
        /// An optional field for your Routing Key.
        /// * `FLOWDOCK`
        /// </summary>
        [Output("routingKey")]
        public Output<string?> RoutingKey { get; private set; } = null!;

        /// <summary>
        /// Your Prometheus protocol scheme configured for requests.
        /// </summary>
        [Output("scheme")]
        public Output<string?> Scheme { get; private set; } = null!;

        /// <summary>
        /// An optional field for your webhook secret.
        /// * `MICROSOFT_TEAMS`
        /// </summary>
        [Output("secret")]
        public Output<string?> Secret { get; private set; } = null!;

        /// <summary>
        /// Indicates which service discovery method is used, either file or http.
        /// </summary>
        [Output("serviceDiscovery")]
        public Output<string?> ServiceDiscovery { get; private set; } = null!;

        /// <summary>
        /// Your Service Key.
        /// * `DATADOG`
        /// </summary>
        [Output("serviceKey")]
        public Output<string?> ServiceKey { get; private set; } = null!;

        [Output("teamName")]
        public Output<string?> TeamName { get; private set; } = null!;

        /// <summary>
        /// Third-Party Integration Settings type 
        /// * PAGER_DUTY
        /// * DATADOG
        /// * NEW_RELIC
        /// * OPS_GENIE
        /// * VICTOR_OPS
        /// * FLOWDOCK
        /// * WEBHOOK
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// Your webhook URL.
        /// </summary>
        [Output("url")]
        public Output<string?> Url { get; private set; } = null!;

        /// <summary>
        /// Your Prometheus username.
        /// </summary>
        [Output("userName")]
        public Output<string?> UserName { get; private set; } = null!;

        /// <summary>
        /// Your Insights Insert Key.
        /// </summary>
        [Output("writeToken")]
        public Output<string?> WriteToken { get; private set; } = null!;


        /// <summary>
        /// Create a ThirdPartyIntegration resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ThirdPartyIntegration(string name, ThirdPartyIntegrationArgs args, CustomResourceOptions? options = null)
            : base("mongodbatlas:index/thirdPartyIntegration:ThirdPartyIntegration", name, args ?? new ThirdPartyIntegrationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ThirdPartyIntegration(string name, Input<string> id, ThirdPartyIntegrationState? state = null, CustomResourceOptions? options = null)
            : base("mongodbatlas:index/thirdPartyIntegration:ThirdPartyIntegration", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ThirdPartyIntegration resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ThirdPartyIntegration Get(string name, Input<string> id, ThirdPartyIntegrationState? state = null, CustomResourceOptions? options = null)
        {
            return new ThirdPartyIntegration(name, id, state, options);
        }
    }

    public sealed class ThirdPartyIntegrationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Unique identifier of your New Relic account.
        /// </summary>
        [Input("accountId")]
        public Input<string>? AccountId { get; set; }

        /// <summary>
        /// Your API Key.
        /// </summary>
        [Input("apiKey")]
        public Input<string>? ApiKey { get; set; }

        /// <summary>
        /// Your API Token.
        /// </summary>
        [Input("apiToken")]
        public Input<string>? ApiToken { get; set; }

        [Input("channelName")]
        public Input<string>? ChannelName { get; set; }

        /// <summary>
        /// Whether your cluster has Prometheus enabled.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// Your Flowdock Flow name.
        /// </summary>
        [Input("flowName")]
        public Input<string>? FlowName { get; set; }

        /// <summary>
        /// Your License Key.
        /// </summary>
        [Input("licenseKey")]
        public Input<string>? LicenseKey { get; set; }

        /// <summary>
        /// Your Microsoft Teams incoming webhook URL.
        /// * `PROMETHEUS`
        /// </summary>
        [Input("microsoftTeamsWebhookUrl")]
        public Input<string>? MicrosoftTeamsWebhookUrl { get; set; }

        /// <summary>
        /// Your Flowdock organization name.
        /// * `WEBHOOK`
        /// </summary>
        [Input("orgName")]
        public Input<string>? OrgName { get; set; }

        /// <summary>
        /// Your Prometheus password.
        /// </summary>
        [Input("password")]
        public Input<string>? Password { get; set; }

        /// <summary>
        /// The unique ID for the project to get all Third-Party service integrations
        /// </summary>
        [Input("projectId", required: true)]
        public Input<string> ProjectId { get; set; } = null!;

        /// <summary>
        /// Your Insights Query Key.
        /// * `OPS_GENIE`
        /// </summary>
        [Input("readToken")]
        public Input<string>? ReadToken { get; set; }

        /// <summary>
        /// Indicates which API URL to use, either US or EU. Opsgenie will use US by default.
        /// * `VICTOR_OPS`
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// An optional field for your Routing Key.
        /// * `FLOWDOCK`
        /// </summary>
        [Input("routingKey")]
        public Input<string>? RoutingKey { get; set; }

        /// <summary>
        /// Your Prometheus protocol scheme configured for requests.
        /// </summary>
        [Input("scheme")]
        public Input<string>? Scheme { get; set; }

        /// <summary>
        /// An optional field for your webhook secret.
        /// * `MICROSOFT_TEAMS`
        /// </summary>
        [Input("secret")]
        public Input<string>? Secret { get; set; }

        /// <summary>
        /// Indicates which service discovery method is used, either file or http.
        /// </summary>
        [Input("serviceDiscovery")]
        public Input<string>? ServiceDiscovery { get; set; }

        /// <summary>
        /// Your Service Key.
        /// * `DATADOG`
        /// </summary>
        [Input("serviceKey")]
        public Input<string>? ServiceKey { get; set; }

        [Input("teamName")]
        public Input<string>? TeamName { get; set; }

        /// <summary>
        /// Third-Party Integration Settings type 
        /// * PAGER_DUTY
        /// * DATADOG
        /// * NEW_RELIC
        /// * OPS_GENIE
        /// * VICTOR_OPS
        /// * FLOWDOCK
        /// * WEBHOOK
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        /// <summary>
        /// Your webhook URL.
        /// </summary>
        [Input("url")]
        public Input<string>? Url { get; set; }

        /// <summary>
        /// Your Prometheus username.
        /// </summary>
        [Input("userName")]
        public Input<string>? UserName { get; set; }

        /// <summary>
        /// Your Insights Insert Key.
        /// </summary>
        [Input("writeToken")]
        public Input<string>? WriteToken { get; set; }

        public ThirdPartyIntegrationArgs()
        {
        }
    }

    public sealed class ThirdPartyIntegrationState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Unique identifier of your New Relic account.
        /// </summary>
        [Input("accountId")]
        public Input<string>? AccountId { get; set; }

        /// <summary>
        /// Your API Key.
        /// </summary>
        [Input("apiKey")]
        public Input<string>? ApiKey { get; set; }

        /// <summary>
        /// Your API Token.
        /// </summary>
        [Input("apiToken")]
        public Input<string>? ApiToken { get; set; }

        [Input("channelName")]
        public Input<string>? ChannelName { get; set; }

        /// <summary>
        /// Whether your cluster has Prometheus enabled.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// Your Flowdock Flow name.
        /// </summary>
        [Input("flowName")]
        public Input<string>? FlowName { get; set; }

        /// <summary>
        /// Your License Key.
        /// </summary>
        [Input("licenseKey")]
        public Input<string>? LicenseKey { get; set; }

        /// <summary>
        /// Your Microsoft Teams incoming webhook URL.
        /// * `PROMETHEUS`
        /// </summary>
        [Input("microsoftTeamsWebhookUrl")]
        public Input<string>? MicrosoftTeamsWebhookUrl { get; set; }

        /// <summary>
        /// Your Flowdock organization name.
        /// * `WEBHOOK`
        /// </summary>
        [Input("orgName")]
        public Input<string>? OrgName { get; set; }

        /// <summary>
        /// Your Prometheus password.
        /// </summary>
        [Input("password")]
        public Input<string>? Password { get; set; }

        /// <summary>
        /// The unique ID for the project to get all Third-Party service integrations
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// Your Insights Query Key.
        /// * `OPS_GENIE`
        /// </summary>
        [Input("readToken")]
        public Input<string>? ReadToken { get; set; }

        /// <summary>
        /// Indicates which API URL to use, either US or EU. Opsgenie will use US by default.
        /// * `VICTOR_OPS`
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// An optional field for your Routing Key.
        /// * `FLOWDOCK`
        /// </summary>
        [Input("routingKey")]
        public Input<string>? RoutingKey { get; set; }

        /// <summary>
        /// Your Prometheus protocol scheme configured for requests.
        /// </summary>
        [Input("scheme")]
        public Input<string>? Scheme { get; set; }

        /// <summary>
        /// An optional field for your webhook secret.
        /// * `MICROSOFT_TEAMS`
        /// </summary>
        [Input("secret")]
        public Input<string>? Secret { get; set; }

        /// <summary>
        /// Indicates which service discovery method is used, either file or http.
        /// </summary>
        [Input("serviceDiscovery")]
        public Input<string>? ServiceDiscovery { get; set; }

        /// <summary>
        /// Your Service Key.
        /// * `DATADOG`
        /// </summary>
        [Input("serviceKey")]
        public Input<string>? ServiceKey { get; set; }

        [Input("teamName")]
        public Input<string>? TeamName { get; set; }

        /// <summary>
        /// Third-Party Integration Settings type 
        /// * PAGER_DUTY
        /// * DATADOG
        /// * NEW_RELIC
        /// * OPS_GENIE
        /// * VICTOR_OPS
        /// * FLOWDOCK
        /// * WEBHOOK
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// Your webhook URL.
        /// </summary>
        [Input("url")]
        public Input<string>? Url { get; set; }

        /// <summary>
        /// Your Prometheus username.
        /// </summary>
        [Input("userName")]
        public Input<string>? UserName { get; set; }

        /// <summary>
        /// Your Insights Insert Key.
        /// </summary>
        [Input("writeToken")]
        public Input<string>? WriteToken { get; set; }

        public ThirdPartyIntegrationState()
        {
        }
    }
}
