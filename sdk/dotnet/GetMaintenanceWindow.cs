// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi.Utilities;

namespace Pulumi.Mongodbatlas
{
    public static class GetMaintenanceWindow
    {
        /// <summary>
        /// `mongodbatlas.MaintenanceWindow` provides a Maintenance Window entry datasource. Gets information regarding the configured maintenance window for a MongoDB Atlas project.
        /// 
        /// &gt; **NOTE:** Groups and projects are synonymous terms. You may find `groupId` in the official documentation.
        /// </summary>
        public static Task<GetMaintenanceWindowResult> InvokeAsync(GetMaintenanceWindowArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetMaintenanceWindowResult>("mongodbatlas:index/getMaintenanceWindow:getMaintenanceWindow", args ?? new GetMaintenanceWindowArgs(), options.WithVersion());

        /// <summary>
        /// `mongodbatlas.MaintenanceWindow` provides a Maintenance Window entry datasource. Gets information regarding the configured maintenance window for a MongoDB Atlas project.
        /// 
        /// &gt; **NOTE:** Groups and projects are synonymous terms. You may find `groupId` in the official documentation.
        /// </summary>
        public static Output<GetMaintenanceWindowResult> Invoke(GetMaintenanceWindowInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetMaintenanceWindowResult>("mongodbatlas:index/getMaintenanceWindow:getMaintenanceWindow", args ?? new GetMaintenanceWindowInvokeArgs(), options.WithVersion());
    }


    public sealed class GetMaintenanceWindowArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The unique identifier of the project for the Maintenance Window.
        /// </summary>
        [Input("projectId", required: true)]
        public string ProjectId { get; set; } = null!;

        public GetMaintenanceWindowArgs()
        {
        }
    }

    public sealed class GetMaintenanceWindowInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The unique identifier of the project for the Maintenance Window.
        /// </summary>
        [Input("projectId", required: true)]
        public Input<string> ProjectId { get; set; } = null!;

        public GetMaintenanceWindowInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetMaintenanceWindowResult
    {
        /// <summary>
        /// Flag that indicates whether you want to defer all maintenance windows one week they would be triggered.
        /// For more information see: [MongoDB Atlas API Reference.](https://docs.atlas.mongodb.com/reference/api/maintenance-windows/)
        /// </summary>
        public readonly bool AutoDeferOnceEnabled;
        /// <summary>
        /// Day of the week when you would like the maintenance window to start as a 1-based integer: S=1, M=2, T=3, W=4, T=5, F=6, S=7.
        /// </summary>
        public readonly int DayOfWeek;
        /// <summary>
        /// Hour of the day when you would like the maintenance window to start. This parameter uses the 24-hour clock, where midnight is 0, noon is 12  (Time zone is UTC).
        /// </summary>
        public readonly int HourOfDay;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Number of times the current maintenance event for this project has been deferred, you can set a maximum of 2 deferrals.
        /// </summary>
        public readonly int NumberOfDeferrals;
        public readonly string ProjectId;
        /// <summary>
        /// Flag indicating whether project maintenance has been directed to start immediately. If you request that maintenance begin immediately, this field returns true from the time the request was made until the time the maintenance event completes.
        /// </summary>
        public readonly bool StartAsap;

        [OutputConstructor]
        private GetMaintenanceWindowResult(
            bool autoDeferOnceEnabled,

            int dayOfWeek,

            int hourOfDay,

            string id,

            int numberOfDeferrals,

            string projectId,

            bool startAsap)
        {
            AutoDeferOnceEnabled = autoDeferOnceEnabled;
            DayOfWeek = dayOfWeek;
            HourOfDay = hourOfDay;
            Id = id;
            NumberOfDeferrals = numberOfDeferrals;
            ProjectId = projectId;
            StartAsap = startAsap;
        }
    }
}
