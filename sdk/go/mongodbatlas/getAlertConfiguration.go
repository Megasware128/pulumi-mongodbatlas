// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mongodbatlas

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// `AlertConfiguration` describes an Alert Configuration.
//
// > **NOTE:** Groups and projects are synonymous terms. You may find **group_id** in the official documentation.
func LookupAlertConfiguration(ctx *pulumi.Context, args *LookupAlertConfigurationArgs, opts ...pulumi.InvokeOption) (*LookupAlertConfigurationResult, error) {
	var rv LookupAlertConfigurationResult
	err := ctx.Invoke("mongodbatlas:index/getAlertConfiguration:getAlertConfiguration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAlertConfiguration.
type LookupAlertConfigurationArgs struct {
	// Unique identifier for the alert configuration.
	AlertConfigurationId string `pulumi:"alertConfigurationId"`
	// The ID of the project where the alert configuration will create.
	ProjectId string `pulumi:"projectId"`
}

// A collection of values returned by getAlertConfiguration.
type LookupAlertConfigurationResult struct {
	AlertConfigurationId string `pulumi:"alertConfigurationId"`
	// Timestamp in ISO 8601 date and time format in UTC when this alert configuration was created.
	Created string `pulumi:"created"`
	// If set to true, the alert configuration is enabled. If enabled is not exported it is set to false.
	Enabled bool `pulumi:"enabled"`
	// The type of event that will trigger an alert.
	EventType string `pulumi:"eventType"`
	// The provider-assigned unique ID for this managed resource.
	Id                     string                                       `pulumi:"id"`
	Matchers               []GetAlertConfigurationMatcher               `pulumi:"matchers"`
	MetricThreshold        map[string]string                            `pulumi:"metricThreshold"`
	MetricThresholdConfigs []GetAlertConfigurationMetricThresholdConfig `pulumi:"metricThresholdConfigs"`
	Notifications          []GetAlertConfigurationNotification          `pulumi:"notifications"`
	ProjectId              string                                       `pulumi:"projectId"`
	// Threshold value outside of which an alert will be triggered.
	Threshold        map[string]string                      `pulumi:"threshold"`
	ThresholdConfigs []GetAlertConfigurationThresholdConfig `pulumi:"thresholdConfigs"`
	// Timestamp in ISO 8601 date and time format in UTC when this alert configuration was last updated.
	Updated string `pulumi:"updated"`
}

func LookupAlertConfigurationOutput(ctx *pulumi.Context, args LookupAlertConfigurationOutputArgs, opts ...pulumi.InvokeOption) LookupAlertConfigurationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAlertConfigurationResult, error) {
			args := v.(LookupAlertConfigurationArgs)
			r, err := LookupAlertConfiguration(ctx, &args, opts...)
			var s LookupAlertConfigurationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAlertConfigurationResultOutput)
}

// A collection of arguments for invoking getAlertConfiguration.
type LookupAlertConfigurationOutputArgs struct {
	// Unique identifier for the alert configuration.
	AlertConfigurationId pulumi.StringInput `pulumi:"alertConfigurationId"`
	// The ID of the project where the alert configuration will create.
	ProjectId pulumi.StringInput `pulumi:"projectId"`
}

func (LookupAlertConfigurationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAlertConfigurationArgs)(nil)).Elem()
}

// A collection of values returned by getAlertConfiguration.
type LookupAlertConfigurationResultOutput struct{ *pulumi.OutputState }

func (LookupAlertConfigurationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAlertConfigurationResult)(nil)).Elem()
}

func (o LookupAlertConfigurationResultOutput) ToLookupAlertConfigurationResultOutput() LookupAlertConfigurationResultOutput {
	return o
}

func (o LookupAlertConfigurationResultOutput) ToLookupAlertConfigurationResultOutputWithContext(ctx context.Context) LookupAlertConfigurationResultOutput {
	return o
}

func (o LookupAlertConfigurationResultOutput) AlertConfigurationId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAlertConfigurationResult) string { return v.AlertConfigurationId }).(pulumi.StringOutput)
}

// Timestamp in ISO 8601 date and time format in UTC when this alert configuration was created.
func (o LookupAlertConfigurationResultOutput) Created() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAlertConfigurationResult) string { return v.Created }).(pulumi.StringOutput)
}

// If set to true, the alert configuration is enabled. If enabled is not exported it is set to false.
func (o LookupAlertConfigurationResultOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupAlertConfigurationResult) bool { return v.Enabled }).(pulumi.BoolOutput)
}

// The type of event that will trigger an alert.
func (o LookupAlertConfigurationResultOutput) EventType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAlertConfigurationResult) string { return v.EventType }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupAlertConfigurationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAlertConfigurationResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupAlertConfigurationResultOutput) Matchers() GetAlertConfigurationMatcherArrayOutput {
	return o.ApplyT(func(v LookupAlertConfigurationResult) []GetAlertConfigurationMatcher { return v.Matchers }).(GetAlertConfigurationMatcherArrayOutput)
}

func (o LookupAlertConfigurationResultOutput) MetricThreshold() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupAlertConfigurationResult) map[string]string { return v.MetricThreshold }).(pulumi.StringMapOutput)
}

func (o LookupAlertConfigurationResultOutput) MetricThresholdConfigs() GetAlertConfigurationMetricThresholdConfigArrayOutput {
	return o.ApplyT(func(v LookupAlertConfigurationResult) []GetAlertConfigurationMetricThresholdConfig {
		return v.MetricThresholdConfigs
	}).(GetAlertConfigurationMetricThresholdConfigArrayOutput)
}

func (o LookupAlertConfigurationResultOutput) Notifications() GetAlertConfigurationNotificationArrayOutput {
	return o.ApplyT(func(v LookupAlertConfigurationResult) []GetAlertConfigurationNotification { return v.Notifications }).(GetAlertConfigurationNotificationArrayOutput)
}

func (o LookupAlertConfigurationResultOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAlertConfigurationResult) string { return v.ProjectId }).(pulumi.StringOutput)
}

// Threshold value outside of which an alert will be triggered.
func (o LookupAlertConfigurationResultOutput) Threshold() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupAlertConfigurationResult) map[string]string { return v.Threshold }).(pulumi.StringMapOutput)
}

func (o LookupAlertConfigurationResultOutput) ThresholdConfigs() GetAlertConfigurationThresholdConfigArrayOutput {
	return o.ApplyT(func(v LookupAlertConfigurationResult) []GetAlertConfigurationThresholdConfig {
		return v.ThresholdConfigs
	}).(GetAlertConfigurationThresholdConfigArrayOutput)
}

// Timestamp in ISO 8601 date and time format in UTC when this alert configuration was last updated.
func (o LookupAlertConfigurationResultOutput) Updated() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAlertConfigurationResult) string { return v.Updated }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAlertConfigurationResultOutput{})
}
