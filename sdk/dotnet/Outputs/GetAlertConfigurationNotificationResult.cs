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
    public sealed class GetAlertConfigurationNotificationResult
    {
        /// <summary>
        /// Slack API token. Required for the SLACK notifications type. If the token later becomes invalid, Atlas sends an email to the project owner and eventually removes the token.
        /// </summary>
        public readonly string ApiToken;
        /// <summary>
        /// Slack channel name. Required for the SLACK notifications type.
        /// </summary>
        public readonly string ChannelName;
        /// <summary>
        /// Datadog API Key. Found in the Datadog dashboard. Required for the DATADOG notifications type.
        /// </summary>
        public readonly string DatadogApiKey;
        /// <summary>
        /// Region that indicates which API URL to use. Accepted regions are: `US`, `EU`. The default Datadog region is US.
        /// </summary>
        public readonly string DatadogRegion;
        /// <summary>
        /// Number of minutes to wait after an alert condition is detected before sending out the first notification.
        /// </summary>
        public readonly int DelayMin;
        /// <summary>
        /// Email address to which alert notifications are sent. Required for the EMAIL notifications type.
        /// </summary>
        public readonly string EmailAddress;
        /// <summary>
        /// Flag indicating email notifications should be sent. Atlas returns this value if `type_name` is set  to `ORG`, `GROUP`, or `USER`.
        /// </summary>
        public readonly bool EmailEnabled;
        /// <summary>
        /// Flowdock flow name in lower-case letters. Required for the `FLOWDOCK` notifications type
        /// </summary>
        public readonly string FlowName;
        /// <summary>
        /// The Flowdock personal API token. Required for the `FLOWDOCK` notifications type. If the token later becomes invalid, Atlas sends an email to the project owner and eventually removes the token.
        /// </summary>
        public readonly string FlowdockApiToken;
        /// <summary>
        /// Number of minutes to wait between successive notifications for unacknowledged alerts that are not resolved. The minimum value is 5.
        /// </summary>
        public readonly int IntervalMin;
        /// <summary>
        /// Mobile number to which alert notifications are sent. Required for the SMS notifications type.
        /// </summary>
        public readonly string MobileNumber;
        /// <summary>
        /// Opsgenie API Key. Required for the `OPS_GENIE` notifications type. If the key later becomes invalid, Atlas sends an email to the project owner and eventually removes the token.
        /// </summary>
        public readonly string OpsGenieApiKey;
        /// <summary>
        /// Region that indicates which API URL to use. Accepted regions are: `US` ,`EU`. The default Opsgenie region is US.
        /// </summary>
        public readonly string OpsGenieRegion;
        /// <summary>
        /// Flowdock organization name in lower-case letters. This is the name that appears after www.flowdock.com/app/ in the URL string. Required for the FLOWDOCK notifications type.
        /// </summary>
        public readonly string OrgName;
        /// <summary>
        /// Atlas role in current Project or Organization. Atlas returns this value if you set `type_name` to `ORG` or `GROUP`.
        /// </summary>
        public readonly ImmutableArray<string> Roles;
        /// <summary>
        /// PagerDuty service key. Required for the PAGER_DUTY notifications type. If the key later becomes invalid, Atlas sends an email to the project owner and eventually removes the key.
        /// </summary>
        public readonly string ServiceKey;
        /// <summary>
        /// Flag indicating text notifications should be sent. Atlas returns this value if `type_name` is set to `ORG`, `GROUP`, or `USER`.
        /// </summary>
        public readonly bool SmsEnabled;
        /// <summary>
        /// Unique identifier of a team.
        /// </summary>
        public readonly string TeamId;
        /// <summary>
        /// Label for the team that receives this notification.
        /// </summary>
        public readonly string TeamName;
        /// <summary>
        /// Type of alert notification.
        /// Accepted values are:
        /// - `DATADOG`
        /// - `EMAIL`
        /// - `FLOWDOCK`
        /// - `GROUP` (Project)
        /// - `OPS_GENIE`
        /// - `ORG`
        /// - `PAGER_DUTY`
        /// - `SLACK`
        /// - `SMS`
        /// - `TEAM`
        /// - `USER`
        /// - `VICTOR_OPS`
        /// - `WEBHOOK`
        /// </summary>
        public readonly string TypeName;
        /// <summary>
        /// Name of the Atlas user to which to send notifications. Only a user in the project that owns the alert configuration is allowed here. Required for the `USER` notifications type.
        /// </summary>
        public readonly string Username;
        /// <summary>
        /// VictorOps API key. Required for the `VICTOR_OPS` notifications type. If the key later becomes invalid, Atlas sends an email to the project owner and eventually removes the key.
        /// </summary>
        public readonly string VictorOpsApiKey;
        /// <summary>
        /// VictorOps routing key. Optional for the `VICTOR_OPS` notifications type. If the key later becomes invalid, Atlas sends an email to the project owner and eventually removes the key.
        /// </summary>
        public readonly string VictorOpsRoutingKey;

        [OutputConstructor]
        private GetAlertConfigurationNotificationResult(
            string apiToken,

            string channelName,

            string datadogApiKey,

            string datadogRegion,

            int delayMin,

            string emailAddress,

            bool emailEnabled,

            string flowName,

            string flowdockApiToken,

            int intervalMin,

            string mobileNumber,

            string opsGenieApiKey,

            string opsGenieRegion,

            string orgName,

            ImmutableArray<string> roles,

            string serviceKey,

            bool smsEnabled,

            string teamId,

            string teamName,

            string typeName,

            string username,

            string victorOpsApiKey,

            string victorOpsRoutingKey)
        {
            ApiToken = apiToken;
            ChannelName = channelName;
            DatadogApiKey = datadogApiKey;
            DatadogRegion = datadogRegion;
            DelayMin = delayMin;
            EmailAddress = emailAddress;
            EmailEnabled = emailEnabled;
            FlowName = flowName;
            FlowdockApiToken = flowdockApiToken;
            IntervalMin = intervalMin;
            MobileNumber = mobileNumber;
            OpsGenieApiKey = opsGenieApiKey;
            OpsGenieRegion = opsGenieRegion;
            OrgName = orgName;
            Roles = roles;
            ServiceKey = serviceKey;
            SmsEnabled = smsEnabled;
            TeamId = teamId;
            TeamName = teamName;
            TypeName = typeName;
            Username = username;
            VictorOpsApiKey = victorOpsApiKey;
            VictorOpsRoutingKey = victorOpsRoutingKey;
        }
    }
}
