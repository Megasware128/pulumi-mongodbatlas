// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package mongodbatlas

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// `getEventTriggers` describe all Event Triggers.
func GetEventTriggers(ctx *pulumi.Context, args *GetEventTriggersArgs, opts ...pulumi.InvokeOption) (*GetEventTriggersResult, error) {
	var rv GetEventTriggersResult
	err := ctx.Invoke("mongodbatlas:index/getEventTriggers:getEventTriggers", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getEventTriggers.
type GetEventTriggersArgs struct {
	// The ObjectID of your application.
	AppId string `pulumi:"appId"`
	// The unique ID for the project to get all event triggers.
	ProjectId string `pulumi:"projectId"`
}

// A collection of values returned by getEventTriggers.
type GetEventTriggersResult struct {
	AppId string `pulumi:"appId"`
	// The provider-assigned unique ID for this managed resource.
	Id        string `pulumi:"id"`
	ProjectId string `pulumi:"projectId"`
	// A list where each represents a Event Trigger.
	Results []GetEventTriggersResultType `pulumi:"results"`
}