// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mongodbatlas

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// `Project` describes a MongoDB Atlas Project. This represents a project that has been created.
//
// > **NOTE:** Groups and projects are synonymous terms. You may find groupId in the official documentation.
//
// ## Example Usage
func LookupProject(ctx *pulumi.Context, args *LookupProjectArgs, opts ...pulumi.InvokeOption) (*LookupProjectResult, error) {
	var rv LookupProjectResult
	err := ctx.Invoke("mongodbatlas:index/getProject:getProject", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getProject.
type LookupProjectArgs struct {
	// The unique ID for the project.
	Name *string `pulumi:"name"`
	// The unique ID for the project.
	ProjectId *string `pulumi:"projectId"`
}

// A collection of values returned by getProject.
type LookupProjectResult struct {
	ApiKeys      []GetProjectApiKey `pulumi:"apiKeys"`
	ClusterCount int                `pulumi:"clusterCount"`
	Created      string             `pulumi:"created"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Flag that indicates whether to enable statistics in [cluster metrics](https://www.mongodb.com/docs/atlas/monitor-cluster-metrics/) collection for the project.
	IsCollectDatabaseSpecificsStatisticsEnabled bool `pulumi:"isCollectDatabaseSpecificsStatisticsEnabled"`
	// Flag that indicates whether to enable Data Explorer for the project. If enabled, you can query your database with an easy to use interface.
	IsDataExplorerEnabled bool `pulumi:"isDataExplorerEnabled"`
	// Flag that indicates whether to enable Performance Advisor and Profiler for the project. If enabled, you can analyze database logs to recommend performance improvements.
	IsPerformanceAdvisorEnabled bool `pulumi:"isPerformanceAdvisorEnabled"`
	// Flag that indicates whether to enable Real Time Performance Panel for the project. If enabled, you can see real time metrics from your MongoDB database.
	IsRealtimePerformancePanelEnabled bool `pulumi:"isRealtimePerformancePanelEnabled"`
	// Flag that indicates whether to enable Schema Advisor for the project. If enabled, you receive customized recommendations to optimize your data model and enhance performance. Disable this setting to disable schema suggestions in the [Performance Advisor](https://www.mongodb.com/docs/atlas/performance-advisor/#std-label-performance-advisor) and the [Data Explorer](https://www.mongodb.com/docs/atlas/atlas-ui/#std-label-atlas-ui).
	IsSchemaAdvisorEnabled bool `pulumi:"isSchemaAdvisorEnabled"`
	// The name of the project you want to create. (Cannot be changed via this Provider after creation.)
	Name *string `pulumi:"name"`
	// The ID of the organization you want to create the project within.
	// *`clusterCount` - The number of Atlas clusters deployed in the project.
	// *`created` - The ISO-8601-formatted timestamp of when Atlas created the project.
	// * `teams.#.team_id` - The unique identifier of the team you want to associate with the project. The team and project must share the same parent organization.
	// * `teams.#.role_names` - Each string in the array represents a project role assigned to the team. Every user associated with the team inherits these roles.
	//   The following are valid roles:
	// * `GROUP_OWNER`
	// * `GROUP_READ_ONLY`
	// * `GROUP_DATA_ACCESS_ADMIN`
	// * `GROUP_DATA_ACCESS_READ_WRITE`
	// * `GROUP_DATA_ACCESS_READ_ONLY`
	// * `GROUP_CLUSTER_MANAGER`
	// * `api_keys.#.api_key_id` - The unique identifier of the programmatic API key you want to associate with the project. The programmatic API key and project must share the same parent organization.
	// * `api_keys.#.role_names` - Each string in the array represents a project role assigned to the programmatic API key.
	//   The following are valid roles:
	// * `GROUP_OWNER`
	// * `GROUP_READ_ONLY`
	// * `GROUP_DATA_ACCESS_ADMIN`
	// * `GROUP_DATA_ACCESS_READ_WRITE`
	// * `GROUP_DATA_ACCESS_READ_ONLY`
	// * `GROUP_CLUSTER_MANAGER`
	OrgId     string           `pulumi:"orgId"`
	ProjectId *string          `pulumi:"projectId"`
	Teams     []GetProjectTeam `pulumi:"teams"`
}

func LookupProjectOutput(ctx *pulumi.Context, args LookupProjectOutputArgs, opts ...pulumi.InvokeOption) LookupProjectResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupProjectResult, error) {
			args := v.(LookupProjectArgs)
			r, err := LookupProject(ctx, &args, opts...)
			var s LookupProjectResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupProjectResultOutput)
}

// A collection of arguments for invoking getProject.
type LookupProjectOutputArgs struct {
	// The unique ID for the project.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// The unique ID for the project.
	ProjectId pulumi.StringPtrInput `pulumi:"projectId"`
}

func (LookupProjectOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProjectArgs)(nil)).Elem()
}

// A collection of values returned by getProject.
type LookupProjectResultOutput struct{ *pulumi.OutputState }

func (LookupProjectResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProjectResult)(nil)).Elem()
}

func (o LookupProjectResultOutput) ToLookupProjectResultOutput() LookupProjectResultOutput {
	return o
}

func (o LookupProjectResultOutput) ToLookupProjectResultOutputWithContext(ctx context.Context) LookupProjectResultOutput {
	return o
}

func (o LookupProjectResultOutput) ApiKeys() GetProjectApiKeyArrayOutput {
	return o.ApplyT(func(v LookupProjectResult) []GetProjectApiKey { return v.ApiKeys }).(GetProjectApiKeyArrayOutput)
}

func (o LookupProjectResultOutput) ClusterCount() pulumi.IntOutput {
	return o.ApplyT(func(v LookupProjectResult) int { return v.ClusterCount }).(pulumi.IntOutput)
}

func (o LookupProjectResultOutput) Created() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.Created }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupProjectResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.Id }).(pulumi.StringOutput)
}

// Flag that indicates whether to enable statistics in [cluster metrics](https://www.mongodb.com/docs/atlas/monitor-cluster-metrics/) collection for the project.
func (o LookupProjectResultOutput) IsCollectDatabaseSpecificsStatisticsEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupProjectResult) bool { return v.IsCollectDatabaseSpecificsStatisticsEnabled }).(pulumi.BoolOutput)
}

// Flag that indicates whether to enable Data Explorer for the project. If enabled, you can query your database with an easy to use interface.
func (o LookupProjectResultOutput) IsDataExplorerEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupProjectResult) bool { return v.IsDataExplorerEnabled }).(pulumi.BoolOutput)
}

// Flag that indicates whether to enable Performance Advisor and Profiler for the project. If enabled, you can analyze database logs to recommend performance improvements.
func (o LookupProjectResultOutput) IsPerformanceAdvisorEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupProjectResult) bool { return v.IsPerformanceAdvisorEnabled }).(pulumi.BoolOutput)
}

// Flag that indicates whether to enable Real Time Performance Panel for the project. If enabled, you can see real time metrics from your MongoDB database.
func (o LookupProjectResultOutput) IsRealtimePerformancePanelEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupProjectResult) bool { return v.IsRealtimePerformancePanelEnabled }).(pulumi.BoolOutput)
}

// Flag that indicates whether to enable Schema Advisor for the project. If enabled, you receive customized recommendations to optimize your data model and enhance performance. Disable this setting to disable schema suggestions in the [Performance Advisor](https://www.mongodb.com/docs/atlas/performance-advisor/#std-label-performance-advisor) and the [Data Explorer](https://www.mongodb.com/docs/atlas/atlas-ui/#std-label-atlas-ui).
func (o LookupProjectResultOutput) IsSchemaAdvisorEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupProjectResult) bool { return v.IsSchemaAdvisorEnabled }).(pulumi.BoolOutput)
}

// The name of the project you want to create. (Cannot be changed via this Provider after creation.)
func (o LookupProjectResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupProjectResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// The ID of the organization you want to create the project within.
// *`clusterCount` - The number of Atlas clusters deployed in the project.
// *`created` - The ISO-8601-formatted timestamp of when Atlas created the project.
//   - `teams.#.team_id` - The unique identifier of the team you want to associate with the project. The team and project must share the same parent organization.
//   - `teams.#.role_names` - Each string in the array represents a project role assigned to the team. Every user associated with the team inherits these roles.
//     The following are valid roles:
//   - `GROUP_OWNER`
//   - `GROUP_READ_ONLY`
//   - `GROUP_DATA_ACCESS_ADMIN`
//   - `GROUP_DATA_ACCESS_READ_WRITE`
//   - `GROUP_DATA_ACCESS_READ_ONLY`
//   - `GROUP_CLUSTER_MANAGER`
//   - `api_keys.#.api_key_id` - The unique identifier of the programmatic API key you want to associate with the project. The programmatic API key and project must share the same parent organization.
//   - `api_keys.#.role_names` - Each string in the array represents a project role assigned to the programmatic API key.
//     The following are valid roles:
//   - `GROUP_OWNER`
//   - `GROUP_READ_ONLY`
//   - `GROUP_DATA_ACCESS_ADMIN`
//   - `GROUP_DATA_ACCESS_READ_WRITE`
//   - `GROUP_DATA_ACCESS_READ_ONLY`
//   - `GROUP_CLUSTER_MANAGER`
func (o LookupProjectResultOutput) OrgId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.OrgId }).(pulumi.StringOutput)
}

func (o LookupProjectResultOutput) ProjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupProjectResult) *string { return v.ProjectId }).(pulumi.StringPtrOutput)
}

func (o LookupProjectResultOutput) Teams() GetProjectTeamArrayOutput {
	return o.ApplyT(func(v LookupProjectResult) []GetProjectTeam { return v.Teams }).(GetProjectTeamArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupProjectResultOutput{})
}
