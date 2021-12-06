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
    public static class GetEventTriggers
    {
        /// <summary>
        /// `mongodbatlas.getEventTriggers` describe all Event Triggers.
        /// </summary>
        public static Task<GetEventTriggersResult> InvokeAsync(GetEventTriggersArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetEventTriggersResult>("mongodbatlas:index/getEventTriggers:getEventTriggers", args ?? new GetEventTriggersArgs(), options.WithVersion());

        /// <summary>
        /// `mongodbatlas.getEventTriggers` describe all Event Triggers.
        /// </summary>
        public static Output<GetEventTriggersResult> Invoke(GetEventTriggersInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetEventTriggersResult>("mongodbatlas:index/getEventTriggers:getEventTriggers", args ?? new GetEventTriggersInvokeArgs(), options.WithVersion());
    }


    public sealed class GetEventTriggersArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ObjectID of your application.
        /// </summary>
        [Input("appId", required: true)]
        public string AppId { get; set; } = null!;

        /// <summary>
        /// The unique ID for the project to get all event triggers.
        /// </summary>
        [Input("projectId", required: true)]
        public string ProjectId { get; set; } = null!;

        public GetEventTriggersArgs()
        {
        }
    }

    public sealed class GetEventTriggersInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ObjectID of your application.
        /// </summary>
        [Input("appId", required: true)]
        public Input<string> AppId { get; set; } = null!;

        /// <summary>
        /// The unique ID for the project to get all event triggers.
        /// </summary>
        [Input("projectId", required: true)]
        public Input<string> ProjectId { get; set; } = null!;

        public GetEventTriggersInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetEventTriggersResult
    {
        public readonly string AppId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string ProjectId;
        /// <summary>
        /// A list where each represents a Event Trigger.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetEventTriggersResultResult> Results;

        [OutputConstructor]
        private GetEventTriggersResult(
            string appId,

            string id,

            string projectId,

            ImmutableArray<Outputs.GetEventTriggersResultResult> results)
        {
            AppId = appId;
            Id = id;
            ProjectId = projectId;
            Results = results;
        }
    }
}
