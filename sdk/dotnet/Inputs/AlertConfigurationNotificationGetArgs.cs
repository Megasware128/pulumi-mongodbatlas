// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Mongodbatlas.Inputs
{

    public sealed class AlertConfigurationNotificationGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Slack API token. Required for the SLACK notifications type. If the token later becomes invalid, Atlas sends an email to the project owner and eventually removes the token.
        /// </summary>
        [Input("apiToken")]
        public Input<string>? ApiToken { get; set; }

        /// <summary>
        /// Slack channel name. Required for the SLACK notifications type.
        /// </summary>
        [Input("channelName")]
        public Input<string>? ChannelName { get; set; }

        /// <summary>
        /// Datadog API Key. Found in the Datadog dashboard. Required for the DATADOG notifications type.
        /// </summary>
        [Input("datadogApiKey")]
        public Input<string>? DatadogApiKey { get; set; }

        /// <summary>
        /// Region that indicates which API URL to use. Accepted regions are: `US`, `EU`. The default Datadog region is US.
        /// </summary>
        [Input("datadogRegion")]
        public Input<string>? DatadogRegion { get; set; }

        /// <summary>
        /// Number of minutes to wait after an alert condition is detected before sending out the first notification.
        /// </summary>
        [Input("delayMin")]
        public Input<int>? DelayMin { get; set; }

        /// <summary>
        /// Email address to which alert notifications are sent. Required for the EMAIL notifications type.
        /// </summary>
        [Input("emailAddress")]
        public Input<string>? EmailAddress { get; set; }

        /// <summary>
        /// Flag indicating if email notifications should be sent. Configurable for `ORG`, `GROUP`, and `USER` notifications types.
        /// </summary>
        [Input("emailEnabled")]
        public Input<bool>? EmailEnabled { get; set; }

        /// <summary>
        /// Flowdock flow name in lower-case letters. Required for the `FLOWDOCK` notifications type
        /// </summary>
        [Input("flowName")]
        public Input<string>? FlowName { get; set; }

        /// <summary>
        /// The Flowdock personal API token. Required for the `FLOWDOCK` notifications type. If the token later becomes invalid, Atlas sends an email to the project owner and eventually removes the token.
        /// </summary>
        [Input("flowdockApiToken")]
        public Input<string>? FlowdockApiToken { get; set; }

        /// <summary>
        /// Number of minutes to wait between successive notifications for unacknowledged alerts that are not resolved. The minimum value is 5.
        /// </summary>
        [Input("intervalMin")]
        public Input<int>? IntervalMin { get; set; }

        /// <summary>
        /// Mobile number to which alert notifications are sent. Required for the SMS notifications type.
        /// </summary>
        [Input("mobileNumber")]
        public Input<string>? MobileNumber { get; set; }

        /// <summary>
        /// Opsgenie API Key. Required for the `OPS_GENIE` notifications type. If the key later becomes invalid, Atlas sends an email to the project owner and eventually removes the token.
        /// </summary>
        [Input("opsGenieApiKey")]
        public Input<string>? OpsGenieApiKey { get; set; }

        /// <summary>
        /// Region that indicates which API URL to use. Accepted regions are: `US` ,`EU`. The default Opsgenie region is US.
        /// </summary>
        [Input("opsGenieRegion")]
        public Input<string>? OpsGenieRegion { get; set; }

        /// <summary>
        /// Flowdock organization name in lower-case letters. This is the name that appears after www.flowdock.com/app/ in the URL string. Required for the FLOWDOCK notifications type.
        /// </summary>
        [Input("orgName")]
        public Input<string>? OrgName { get; set; }

        [Input("roles")]
        private InputList<string>? _roles;
        public InputList<string> Roles
        {
            get => _roles ?? (_roles = new InputList<string>());
            set => _roles = value;
        }

        /// <summary>
        /// PagerDuty service key. Required for the PAGER_DUTY notifications type. If the key later becomes invalid, Atlas sends an email to the project owner and eventually removes the key.
        /// </summary>
        [Input("serviceKey")]
        public Input<string>? ServiceKey { get; set; }

        /// <summary>
        /// Flag indicating if text message notifications should be sent. Configurable for `ORG`, `GROUP`, and `USER` notifications types.
        /// </summary>
        [Input("smsEnabled")]
        public Input<bool>? SmsEnabled { get; set; }

        /// <summary>
        /// Unique identifier of a team.
        /// </summary>
        [Input("teamId")]
        public Input<string>? TeamId { get; set; }

        /// <summary>
        /// Type of alert notification.
        /// Accepted values are:
        /// - `DATADOG`
        /// - `EMAIL`
        /// - `FLOWDOCK`
        /// </summary>
        [Input("typeName")]
        public Input<string>? TypeName { get; set; }

        /// <summary>
        /// Name of the Atlas user to which to send notifications. Only a user in the project that owns the alert configuration is allowed here. Required for the `USER` notifications type.
        /// </summary>
        [Input("username")]
        public Input<string>? Username { get; set; }

        /// <summary>
        /// VictorOps API key. Required for the `VICTOR_OPS` notifications type. If the key later becomes invalid, Atlas sends an email to the project owner and eventually removes the key.
        /// </summary>
        [Input("victorOpsApiKey")]
        public Input<string>? VictorOpsApiKey { get; set; }

        /// <summary>
        /// VictorOps routing key. Optional for the `VICTOR_OPS` notifications type. If the key later becomes invalid, Atlas sends an email to the project owner and eventually removes the key.
        /// </summary>
        [Input("victorOpsRoutingKey")]
        public Input<string>? VictorOpsRoutingKey { get; set; }

        public AlertConfigurationNotificationGetArgs()
        {
        }
    }
}
