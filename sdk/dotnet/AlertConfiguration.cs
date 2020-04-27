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
    /// `mongodbatlas..AlertConfiguration` provides an Alert Configuration resource to define the conditions that trigger an alert and the methods of notification within a MongoDB Atlas project.
    /// 
    /// &gt; **NOTE:** Groups and projects are synonymous terms. You may find `groupId` in the official documentation.
    /// </summary>
    public partial class AlertConfiguration : Pulumi.CustomResource
    {
        /// <summary>
        /// Unique identifier for the alert configuration.
        /// </summary>
        [Output("alertConfigurationId")]
        public Output<string> AlertConfigurationId { get; private set; } = null!;

        /// <summary>
        /// Timestamp in ISO 8601 date and time format in UTC when this alert configuration was created.
        /// </summary>
        [Output("created")]
        public Output<string> Created { get; private set; } = null!;

        /// <summary>
        /// It is not required, but If the attribute is omitted, by default will be false, and the configuration would be disabled. You must set true to enable the configuration.
        /// </summary>
        [Output("enabled")]
        public Output<bool> Enabled { get; private set; } = null!;

        /// <summary>
        /// The type of event that will trigger an alert.
        /// Alert type 	Possible values:
        /// * Host
        /// - `OUTSIDE_METRIC_THRESHOLD`
        /// - `HOST_RESTARTED`
        /// - `HOST_UPGRADED`
        /// - `HOST_NOW_SECONDARY`
        /// - `HOST_NOW_PRIMARY`
        /// * Replica set
        /// - `NO_PRIMARY`
        /// - `TOO_MANY_ELECTIONS`
        /// * Sharded cluster
        /// - `CLUSTER_MONGOS_IS_MISSING`
        /// - `User`
        /// - `JOINED_GROUP`
        /// - `REMOVED_FROM_GROUP`
        /// - `USER_ROLES_CHANGED_AUDIT`
        /// * Project
        /// - `USERS_AWAITING_APPROVAL`
        /// - `USERS_WITHOUT_MULTI_FACTOR_AUTH`
        /// - `GROUP_CREATED`
        /// * Team
        /// - `JOINED_TEAM`
        /// - `REMOVED_FROM_TEAM`
        /// * Organization
        /// - `INVITED_TO_ORG`
        /// - `JOINED_ORG`
        /// * Data Explorer
        /// - `DATA_EXPLORER`
        /// - `DATA_EXPLORER_CRUD`
        /// * Billing
        /// - `CREDIT_CARD_ABOUT_TO_EXPIRE`
        /// - `CHARGE_SUCCEEDED`
        /// - `INVOICE_CLOSED`
        /// </summary>
        [Output("eventType")]
        public Output<string> EventType { get; private set; } = null!;

        [Output("matchers")]
        public Output<ImmutableArray<Outputs.AlertConfigurationMatcher>> Matchers { get; private set; } = null!;

        [Output("metricThreshold")]
        public Output<Outputs.AlertConfigurationMetricThreshold?> MetricThreshold { get; private set; } = null!;

        [Output("notifications")]
        public Output<ImmutableArray<Outputs.AlertConfigurationNotification>> Notifications { get; private set; } = null!;

        /// <summary>
        /// The ID of the project where the alert configuration will create.
        /// </summary>
        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;

        /// <summary>
        /// Timestamp in ISO 8601 date and time format in UTC when this alert configuration was last updated.
        /// </summary>
        [Output("updated")]
        public Output<string> Updated { get; private set; } = null!;


        /// <summary>
        /// Create a AlertConfiguration resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AlertConfiguration(string name, AlertConfigurationArgs args, CustomResourceOptions? options = null)
            : base("mongodbatlas:index/alertConfiguration:AlertConfiguration", name, args ?? new AlertConfigurationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AlertConfiguration(string name, Input<string> id, AlertConfigurationState? state = null, CustomResourceOptions? options = null)
            : base("mongodbatlas:index/alertConfiguration:AlertConfiguration", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AlertConfiguration resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AlertConfiguration Get(string name, Input<string> id, AlertConfigurationState? state = null, CustomResourceOptions? options = null)
        {
            return new AlertConfiguration(name, id, state, options);
        }
    }

    public sealed class AlertConfigurationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// It is not required, but If the attribute is omitted, by default will be false, and the configuration would be disabled. You must set true to enable the configuration.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// The type of event that will trigger an alert.
        /// Alert type 	Possible values:
        /// * Host
        /// - `OUTSIDE_METRIC_THRESHOLD`
        /// - `HOST_RESTARTED`
        /// - `HOST_UPGRADED`
        /// - `HOST_NOW_SECONDARY`
        /// - `HOST_NOW_PRIMARY`
        /// * Replica set
        /// - `NO_PRIMARY`
        /// - `TOO_MANY_ELECTIONS`
        /// * Sharded cluster
        /// - `CLUSTER_MONGOS_IS_MISSING`
        /// - `User`
        /// - `JOINED_GROUP`
        /// - `REMOVED_FROM_GROUP`
        /// - `USER_ROLES_CHANGED_AUDIT`
        /// * Project
        /// - `USERS_AWAITING_APPROVAL`
        /// - `USERS_WITHOUT_MULTI_FACTOR_AUTH`
        /// - `GROUP_CREATED`
        /// * Team
        /// - `JOINED_TEAM`
        /// - `REMOVED_FROM_TEAM`
        /// * Organization
        /// - `INVITED_TO_ORG`
        /// - `JOINED_ORG`
        /// * Data Explorer
        /// - `DATA_EXPLORER`
        /// - `DATA_EXPLORER_CRUD`
        /// * Billing
        /// - `CREDIT_CARD_ABOUT_TO_EXPIRE`
        /// - `CHARGE_SUCCEEDED`
        /// - `INVOICE_CLOSED`
        /// </summary>
        [Input("eventType", required: true)]
        public Input<string> EventType { get; set; } = null!;

        [Input("matchers")]
        private InputList<Inputs.AlertConfigurationMatcherArgs>? _matchers;
        public InputList<Inputs.AlertConfigurationMatcherArgs> Matchers
        {
            get => _matchers ?? (_matchers = new InputList<Inputs.AlertConfigurationMatcherArgs>());
            set => _matchers = value;
        }

        [Input("metricThreshold")]
        public Input<Inputs.AlertConfigurationMetricThresholdArgs>? MetricThreshold { get; set; }

        [Input("notifications", required: true)]
        private InputList<Inputs.AlertConfigurationNotificationArgs>? _notifications;
        public InputList<Inputs.AlertConfigurationNotificationArgs> Notifications
        {
            get => _notifications ?? (_notifications = new InputList<Inputs.AlertConfigurationNotificationArgs>());
            set => _notifications = value;
        }

        /// <summary>
        /// The ID of the project where the alert configuration will create.
        /// </summary>
        [Input("projectId", required: true)]
        public Input<string> ProjectId { get; set; } = null!;

        public AlertConfigurationArgs()
        {
        }
    }

    public sealed class AlertConfigurationState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Unique identifier for the alert configuration.
        /// </summary>
        [Input("alertConfigurationId")]
        public Input<string>? AlertConfigurationId { get; set; }

        /// <summary>
        /// Timestamp in ISO 8601 date and time format in UTC when this alert configuration was created.
        /// </summary>
        [Input("created")]
        public Input<string>? Created { get; set; }

        /// <summary>
        /// It is not required, but If the attribute is omitted, by default will be false, and the configuration would be disabled. You must set true to enable the configuration.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// The type of event that will trigger an alert.
        /// Alert type 	Possible values:
        /// * Host
        /// - `OUTSIDE_METRIC_THRESHOLD`
        /// - `HOST_RESTARTED`
        /// - `HOST_UPGRADED`
        /// - `HOST_NOW_SECONDARY`
        /// - `HOST_NOW_PRIMARY`
        /// * Replica set
        /// - `NO_PRIMARY`
        /// - `TOO_MANY_ELECTIONS`
        /// * Sharded cluster
        /// - `CLUSTER_MONGOS_IS_MISSING`
        /// - `User`
        /// - `JOINED_GROUP`
        /// - `REMOVED_FROM_GROUP`
        /// - `USER_ROLES_CHANGED_AUDIT`
        /// * Project
        /// - `USERS_AWAITING_APPROVAL`
        /// - `USERS_WITHOUT_MULTI_FACTOR_AUTH`
        /// - `GROUP_CREATED`
        /// * Team
        /// - `JOINED_TEAM`
        /// - `REMOVED_FROM_TEAM`
        /// * Organization
        /// - `INVITED_TO_ORG`
        /// - `JOINED_ORG`
        /// * Data Explorer
        /// - `DATA_EXPLORER`
        /// - `DATA_EXPLORER_CRUD`
        /// * Billing
        /// - `CREDIT_CARD_ABOUT_TO_EXPIRE`
        /// - `CHARGE_SUCCEEDED`
        /// - `INVOICE_CLOSED`
        /// </summary>
        [Input("eventType")]
        public Input<string>? EventType { get; set; }

        [Input("matchers")]
        private InputList<Inputs.AlertConfigurationMatcherGetArgs>? _matchers;
        public InputList<Inputs.AlertConfigurationMatcherGetArgs> Matchers
        {
            get => _matchers ?? (_matchers = new InputList<Inputs.AlertConfigurationMatcherGetArgs>());
            set => _matchers = value;
        }

        [Input("metricThreshold")]
        public Input<Inputs.AlertConfigurationMetricThresholdGetArgs>? MetricThreshold { get; set; }

        [Input("notifications")]
        private InputList<Inputs.AlertConfigurationNotificationGetArgs>? _notifications;
        public InputList<Inputs.AlertConfigurationNotificationGetArgs> Notifications
        {
            get => _notifications ?? (_notifications = new InputList<Inputs.AlertConfigurationNotificationGetArgs>());
            set => _notifications = value;
        }

        /// <summary>
        /// The ID of the project where the alert configuration will create.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// Timestamp in ISO 8601 date and time format in UTC when this alert configuration was last updated.
        /// </summary>
        [Input("updated")]
        public Input<string>? Updated { get; set; }

        public AlertConfigurationState()
        {
        }
    }
}