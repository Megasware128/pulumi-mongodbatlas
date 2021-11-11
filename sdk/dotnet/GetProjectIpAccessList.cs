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
    public static class GetProjectIpAccessList
    {
        /// <summary>
        /// `mongodbatlas.ProjectIpAccessList` describes an IP Access List entry resource. The access list grants access from IPs, CIDRs or AWS Security Groups (if VPC Peering is enabled) to clusters within the Project.
        /// 
        /// &gt; **NOTE:** Groups and projects are synonymous terms. You may find `groupId` in the official documentation.
        /// 
        /// &gt; **IMPORTANT:**
        /// When you remove an entry from the access list, existing connections from the removed address(es) may remain open for a variable amount of time. How much time passes before Atlas closes the connection depends on several factors, including how the connection was established, the particular behavior of the application or driver using the address, and the connection protocol (e.g., TCP or UDP). This is particularly important to consider when changing an existing IP address or CIDR block as they cannot be updated via the Provider (comments can however), hence a change will force the destruction and recreation of entries.   
        /// 
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetProjectIpAccessListResult> InvokeAsync(GetProjectIpAccessListArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetProjectIpAccessListResult>("mongodbatlas:index/getProjectIpAccessList:getProjectIpAccessList", args ?? new GetProjectIpAccessListArgs(), options.WithVersion());

        /// <summary>
        /// `mongodbatlas.ProjectIpAccessList` describes an IP Access List entry resource. The access list grants access from IPs, CIDRs or AWS Security Groups (if VPC Peering is enabled) to clusters within the Project.
        /// 
        /// &gt; **NOTE:** Groups and projects are synonymous terms. You may find `groupId` in the official documentation.
        /// 
        /// &gt; **IMPORTANT:**
        /// When you remove an entry from the access list, existing connections from the removed address(es) may remain open for a variable amount of time. How much time passes before Atlas closes the connection depends on several factors, including how the connection was established, the particular behavior of the application or driver using the address, and the connection protocol (e.g., TCP or UDP). This is particularly important to consider when changing an existing IP address or CIDR block as they cannot be updated via the Provider (comments can however), hence a change will force the destruction and recreation of entries.   
        /// 
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetProjectIpAccessListResult> Invoke(GetProjectIpAccessListInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetProjectIpAccessListResult>("mongodbatlas:index/getProjectIpAccessList:getProjectIpAccessList", args ?? new GetProjectIpAccessListInvokeArgs(), options.WithVersion());
    }


    public sealed class GetProjectIpAccessListArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Unique identifier of the AWS security group to add to the access list.
        /// </summary>
        [Input("awsSecurityGroup")]
        public string? AwsSecurityGroup { get; set; }

        /// <summary>
        /// Range of IP addresses in CIDR notation to be added to the access list.
        /// </summary>
        [Input("cidrBlock")]
        public string? CidrBlock { get; set; }

        /// <summary>
        /// Single IP address to be added to the access list.
        /// </summary>
        [Input("ipAddress")]
        public string? IpAddress { get; set; }

        /// <summary>
        /// Unique identifier for the project to which you want to add one or more access list entries.
        /// </summary>
        [Input("projectId", required: true)]
        public string ProjectId { get; set; } = null!;

        public GetProjectIpAccessListArgs()
        {
        }
    }

    public sealed class GetProjectIpAccessListInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Unique identifier of the AWS security group to add to the access list.
        /// </summary>
        [Input("awsSecurityGroup")]
        public Input<string>? AwsSecurityGroup { get; set; }

        /// <summary>
        /// Range of IP addresses in CIDR notation to be added to the access list.
        /// </summary>
        [Input("cidrBlock")]
        public Input<string>? CidrBlock { get; set; }

        /// <summary>
        /// Single IP address to be added to the access list.
        /// </summary>
        [Input("ipAddress")]
        public Input<string>? IpAddress { get; set; }

        /// <summary>
        /// Unique identifier for the project to which you want to add one or more access list entries.
        /// </summary>
        [Input("projectId", required: true)]
        public Input<string> ProjectId { get; set; } = null!;

        public GetProjectIpAccessListInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetProjectIpAccessListResult
    {
        public readonly string AwsSecurityGroup;
        public readonly string CidrBlock;
        /// <summary>
        /// Comment to add to the access list entry.
        /// </summary>
        public readonly string Comment;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string IpAddress;
        public readonly string ProjectId;

        [OutputConstructor]
        private GetProjectIpAccessListResult(
            string awsSecurityGroup,

            string cidrBlock,

            string comment,

            string id,

            string ipAddress,

            string projectId)
        {
            AwsSecurityGroup = awsSecurityGroup;
            CidrBlock = cidrBlock;
            Comment = comment;
            Id = id;
            IpAddress = ipAddress;
            ProjectId = projectId;
        }
    }
}
