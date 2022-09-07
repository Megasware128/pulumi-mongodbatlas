// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mongodbatlas

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-mongodbatlas/sdk/v3/go/mongodbatlas"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := mongodbatlas.NewThirdPartyIntegration(ctx, "testFlowdock", &mongodbatlas.ThirdPartyIntegrationArgs{
//				ApiToken:  pulumi.String("<API-TOKEN>"),
//				FlowName:  pulumi.String("<FLOW-NAME>"),
//				OrgName:   pulumi.String("<ORG-NAME>"),
//				ProjectId: pulumi.String("<PROJECT-ID>"),
//				Type:      pulumi.String("FLOWDOCK"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// Third-Party Integration Settings can be imported using project ID and the integration type, in the format `project_id`-`type`, e.g.
//
// ```sh
//
//	$ pulumi import mongodbatlas:index/thirdPartyIntegration:ThirdPartyIntegration my_user 1112222b3bf99403840e8934-OPS_GENIE
//
// ```
//
//	See [MongoDB Atlas API](https://docs.atlas.mongodb.com/reference/api/third-party-integration-settings-create/) Documentation for more information.
type ThirdPartyIntegration struct {
	pulumi.CustomResourceState

	// Unique identifier of your New Relic account.
	AccountId pulumi.StringPtrOutput `pulumi:"accountId"`
	// Your API Key.
	ApiKey pulumi.StringPtrOutput `pulumi:"apiKey"`
	// Your API Token.
	ApiToken    pulumi.StringPtrOutput `pulumi:"apiToken"`
	ChannelName pulumi.StringPtrOutput `pulumi:"channelName"`
	// Whether your cluster has Prometheus enabled.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// Your Flowdock Flow name.
	FlowName pulumi.StringPtrOutput `pulumi:"flowName"`
	// Your License Key.
	LicenseKey pulumi.StringPtrOutput `pulumi:"licenseKey"`
	// Your Microsoft Teams incoming webhook URL.
	// * `PROMETHEUS`
	MicrosoftTeamsWebhookUrl pulumi.StringPtrOutput `pulumi:"microsoftTeamsWebhookUrl"`
	// Your Flowdock organization name.
	// * `WEBHOOK`
	OrgName pulumi.StringPtrOutput `pulumi:"orgName"`
	// Your Prometheus password.
	Password pulumi.StringPtrOutput `pulumi:"password"`
	// The unique ID for the project to get all Third-Party service integrations
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// Your Insights Query Key.
	// * `OPS_GENIE`
	ReadToken pulumi.StringPtrOutput `pulumi:"readToken"`
	// Indicates which API URL to use, either US or EU. Opsgenie will use US by default.
	// * `VICTOR_OPS`
	Region pulumi.StringPtrOutput `pulumi:"region"`
	// An optional field for your Routing Key.
	// * `FLOWDOCK`
	RoutingKey pulumi.StringPtrOutput `pulumi:"routingKey"`
	// Your Prometheus protocol scheme configured for requests.
	Scheme pulumi.StringPtrOutput `pulumi:"scheme"`
	// An optional field for your webhook secret.
	// * `MICROSOFT_TEAMS`
	Secret pulumi.StringPtrOutput `pulumi:"secret"`
	// Indicates which service discovery method is used, either file or http.
	ServiceDiscovery pulumi.StringPtrOutput `pulumi:"serviceDiscovery"`
	// Your Service Key.
	// * `DATADOG`
	ServiceKey pulumi.StringPtrOutput `pulumi:"serviceKey"`
	TeamName   pulumi.StringPtrOutput `pulumi:"teamName"`
	// Third-Party Integration Settings type
	// * PAGER_DUTY
	// * DATADOG
	// * NEW_RELIC
	// * OPS_GENIE
	// * VICTOR_OPS
	// * FLOWDOCK
	// * WEBHOOK
	Type pulumi.StringOutput `pulumi:"type"`
	// Your webhook URL.
	Url pulumi.StringPtrOutput `pulumi:"url"`
	// Your Prometheus username.
	UserName pulumi.StringPtrOutput `pulumi:"userName"`
	// Your Insights Insert Key.
	WriteToken pulumi.StringPtrOutput `pulumi:"writeToken"`
}

// NewThirdPartyIntegration registers a new resource with the given unique name, arguments, and options.
func NewThirdPartyIntegration(ctx *pulumi.Context,
	name string, args *ThirdPartyIntegrationArgs, opts ...pulumi.ResourceOption) (*ThirdPartyIntegration, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ProjectId == nil {
		return nil, errors.New("invalid value for required argument 'ProjectId'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	var resource ThirdPartyIntegration
	err := ctx.RegisterResource("mongodbatlas:index/thirdPartyIntegration:ThirdPartyIntegration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetThirdPartyIntegration gets an existing ThirdPartyIntegration resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetThirdPartyIntegration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ThirdPartyIntegrationState, opts ...pulumi.ResourceOption) (*ThirdPartyIntegration, error) {
	var resource ThirdPartyIntegration
	err := ctx.ReadResource("mongodbatlas:index/thirdPartyIntegration:ThirdPartyIntegration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ThirdPartyIntegration resources.
type thirdPartyIntegrationState struct {
	// Unique identifier of your New Relic account.
	AccountId *string `pulumi:"accountId"`
	// Your API Key.
	ApiKey *string `pulumi:"apiKey"`
	// Your API Token.
	ApiToken    *string `pulumi:"apiToken"`
	ChannelName *string `pulumi:"channelName"`
	// Whether your cluster has Prometheus enabled.
	Enabled *bool `pulumi:"enabled"`
	// Your Flowdock Flow name.
	FlowName *string `pulumi:"flowName"`
	// Your License Key.
	LicenseKey *string `pulumi:"licenseKey"`
	// Your Microsoft Teams incoming webhook URL.
	// * `PROMETHEUS`
	MicrosoftTeamsWebhookUrl *string `pulumi:"microsoftTeamsWebhookUrl"`
	// Your Flowdock organization name.
	// * `WEBHOOK`
	OrgName *string `pulumi:"orgName"`
	// Your Prometheus password.
	Password *string `pulumi:"password"`
	// The unique ID for the project to get all Third-Party service integrations
	ProjectId *string `pulumi:"projectId"`
	// Your Insights Query Key.
	// * `OPS_GENIE`
	ReadToken *string `pulumi:"readToken"`
	// Indicates which API URL to use, either US or EU. Opsgenie will use US by default.
	// * `VICTOR_OPS`
	Region *string `pulumi:"region"`
	// An optional field for your Routing Key.
	// * `FLOWDOCK`
	RoutingKey *string `pulumi:"routingKey"`
	// Your Prometheus protocol scheme configured for requests.
	Scheme *string `pulumi:"scheme"`
	// An optional field for your webhook secret.
	// * `MICROSOFT_TEAMS`
	Secret *string `pulumi:"secret"`
	// Indicates which service discovery method is used, either file or http.
	ServiceDiscovery *string `pulumi:"serviceDiscovery"`
	// Your Service Key.
	// * `DATADOG`
	ServiceKey *string `pulumi:"serviceKey"`
	TeamName   *string `pulumi:"teamName"`
	// Third-Party Integration Settings type
	// * PAGER_DUTY
	// * DATADOG
	// * NEW_RELIC
	// * OPS_GENIE
	// * VICTOR_OPS
	// * FLOWDOCK
	// * WEBHOOK
	Type *string `pulumi:"type"`
	// Your webhook URL.
	Url *string `pulumi:"url"`
	// Your Prometheus username.
	UserName *string `pulumi:"userName"`
	// Your Insights Insert Key.
	WriteToken *string `pulumi:"writeToken"`
}

type ThirdPartyIntegrationState struct {
	// Unique identifier of your New Relic account.
	AccountId pulumi.StringPtrInput
	// Your API Key.
	ApiKey pulumi.StringPtrInput
	// Your API Token.
	ApiToken    pulumi.StringPtrInput
	ChannelName pulumi.StringPtrInput
	// Whether your cluster has Prometheus enabled.
	Enabled pulumi.BoolPtrInput
	// Your Flowdock Flow name.
	FlowName pulumi.StringPtrInput
	// Your License Key.
	LicenseKey pulumi.StringPtrInput
	// Your Microsoft Teams incoming webhook URL.
	// * `PROMETHEUS`
	MicrosoftTeamsWebhookUrl pulumi.StringPtrInput
	// Your Flowdock organization name.
	// * `WEBHOOK`
	OrgName pulumi.StringPtrInput
	// Your Prometheus password.
	Password pulumi.StringPtrInput
	// The unique ID for the project to get all Third-Party service integrations
	ProjectId pulumi.StringPtrInput
	// Your Insights Query Key.
	// * `OPS_GENIE`
	ReadToken pulumi.StringPtrInput
	// Indicates which API URL to use, either US or EU. Opsgenie will use US by default.
	// * `VICTOR_OPS`
	Region pulumi.StringPtrInput
	// An optional field for your Routing Key.
	// * `FLOWDOCK`
	RoutingKey pulumi.StringPtrInput
	// Your Prometheus protocol scheme configured for requests.
	Scheme pulumi.StringPtrInput
	// An optional field for your webhook secret.
	// * `MICROSOFT_TEAMS`
	Secret pulumi.StringPtrInput
	// Indicates which service discovery method is used, either file or http.
	ServiceDiscovery pulumi.StringPtrInput
	// Your Service Key.
	// * `DATADOG`
	ServiceKey pulumi.StringPtrInput
	TeamName   pulumi.StringPtrInput
	// Third-Party Integration Settings type
	// * PAGER_DUTY
	// * DATADOG
	// * NEW_RELIC
	// * OPS_GENIE
	// * VICTOR_OPS
	// * FLOWDOCK
	// * WEBHOOK
	Type pulumi.StringPtrInput
	// Your webhook URL.
	Url pulumi.StringPtrInput
	// Your Prometheus username.
	UserName pulumi.StringPtrInput
	// Your Insights Insert Key.
	WriteToken pulumi.StringPtrInput
}

func (ThirdPartyIntegrationState) ElementType() reflect.Type {
	return reflect.TypeOf((*thirdPartyIntegrationState)(nil)).Elem()
}

type thirdPartyIntegrationArgs struct {
	// Unique identifier of your New Relic account.
	AccountId *string `pulumi:"accountId"`
	// Your API Key.
	ApiKey *string `pulumi:"apiKey"`
	// Your API Token.
	ApiToken    *string `pulumi:"apiToken"`
	ChannelName *string `pulumi:"channelName"`
	// Whether your cluster has Prometheus enabled.
	Enabled *bool `pulumi:"enabled"`
	// Your Flowdock Flow name.
	FlowName *string `pulumi:"flowName"`
	// Your License Key.
	LicenseKey *string `pulumi:"licenseKey"`
	// Your Microsoft Teams incoming webhook URL.
	// * `PROMETHEUS`
	MicrosoftTeamsWebhookUrl *string `pulumi:"microsoftTeamsWebhookUrl"`
	// Your Flowdock organization name.
	// * `WEBHOOK`
	OrgName *string `pulumi:"orgName"`
	// Your Prometheus password.
	Password *string `pulumi:"password"`
	// The unique ID for the project to get all Third-Party service integrations
	ProjectId string `pulumi:"projectId"`
	// Your Insights Query Key.
	// * `OPS_GENIE`
	ReadToken *string `pulumi:"readToken"`
	// Indicates which API URL to use, either US or EU. Opsgenie will use US by default.
	// * `VICTOR_OPS`
	Region *string `pulumi:"region"`
	// An optional field for your Routing Key.
	// * `FLOWDOCK`
	RoutingKey *string `pulumi:"routingKey"`
	// Your Prometheus protocol scheme configured for requests.
	Scheme *string `pulumi:"scheme"`
	// An optional field for your webhook secret.
	// * `MICROSOFT_TEAMS`
	Secret *string `pulumi:"secret"`
	// Indicates which service discovery method is used, either file or http.
	ServiceDiscovery *string `pulumi:"serviceDiscovery"`
	// Your Service Key.
	// * `DATADOG`
	ServiceKey *string `pulumi:"serviceKey"`
	TeamName   *string `pulumi:"teamName"`
	// Third-Party Integration Settings type
	// * PAGER_DUTY
	// * DATADOG
	// * NEW_RELIC
	// * OPS_GENIE
	// * VICTOR_OPS
	// * FLOWDOCK
	// * WEBHOOK
	Type string `pulumi:"type"`
	// Your webhook URL.
	Url *string `pulumi:"url"`
	// Your Prometheus username.
	UserName *string `pulumi:"userName"`
	// Your Insights Insert Key.
	WriteToken *string `pulumi:"writeToken"`
}

// The set of arguments for constructing a ThirdPartyIntegration resource.
type ThirdPartyIntegrationArgs struct {
	// Unique identifier of your New Relic account.
	AccountId pulumi.StringPtrInput
	// Your API Key.
	ApiKey pulumi.StringPtrInput
	// Your API Token.
	ApiToken    pulumi.StringPtrInput
	ChannelName pulumi.StringPtrInput
	// Whether your cluster has Prometheus enabled.
	Enabled pulumi.BoolPtrInput
	// Your Flowdock Flow name.
	FlowName pulumi.StringPtrInput
	// Your License Key.
	LicenseKey pulumi.StringPtrInput
	// Your Microsoft Teams incoming webhook URL.
	// * `PROMETHEUS`
	MicrosoftTeamsWebhookUrl pulumi.StringPtrInput
	// Your Flowdock organization name.
	// * `WEBHOOK`
	OrgName pulumi.StringPtrInput
	// Your Prometheus password.
	Password pulumi.StringPtrInput
	// The unique ID for the project to get all Third-Party service integrations
	ProjectId pulumi.StringInput
	// Your Insights Query Key.
	// * `OPS_GENIE`
	ReadToken pulumi.StringPtrInput
	// Indicates which API URL to use, either US or EU. Opsgenie will use US by default.
	// * `VICTOR_OPS`
	Region pulumi.StringPtrInput
	// An optional field for your Routing Key.
	// * `FLOWDOCK`
	RoutingKey pulumi.StringPtrInput
	// Your Prometheus protocol scheme configured for requests.
	Scheme pulumi.StringPtrInput
	// An optional field for your webhook secret.
	// * `MICROSOFT_TEAMS`
	Secret pulumi.StringPtrInput
	// Indicates which service discovery method is used, either file or http.
	ServiceDiscovery pulumi.StringPtrInput
	// Your Service Key.
	// * `DATADOG`
	ServiceKey pulumi.StringPtrInput
	TeamName   pulumi.StringPtrInput
	// Third-Party Integration Settings type
	// * PAGER_DUTY
	// * DATADOG
	// * NEW_RELIC
	// * OPS_GENIE
	// * VICTOR_OPS
	// * FLOWDOCK
	// * WEBHOOK
	Type pulumi.StringInput
	// Your webhook URL.
	Url pulumi.StringPtrInput
	// Your Prometheus username.
	UserName pulumi.StringPtrInput
	// Your Insights Insert Key.
	WriteToken pulumi.StringPtrInput
}

func (ThirdPartyIntegrationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*thirdPartyIntegrationArgs)(nil)).Elem()
}

type ThirdPartyIntegrationInput interface {
	pulumi.Input

	ToThirdPartyIntegrationOutput() ThirdPartyIntegrationOutput
	ToThirdPartyIntegrationOutputWithContext(ctx context.Context) ThirdPartyIntegrationOutput
}

func (*ThirdPartyIntegration) ElementType() reflect.Type {
	return reflect.TypeOf((**ThirdPartyIntegration)(nil)).Elem()
}

func (i *ThirdPartyIntegration) ToThirdPartyIntegrationOutput() ThirdPartyIntegrationOutput {
	return i.ToThirdPartyIntegrationOutputWithContext(context.Background())
}

func (i *ThirdPartyIntegration) ToThirdPartyIntegrationOutputWithContext(ctx context.Context) ThirdPartyIntegrationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ThirdPartyIntegrationOutput)
}

// ThirdPartyIntegrationArrayInput is an input type that accepts ThirdPartyIntegrationArray and ThirdPartyIntegrationArrayOutput values.
// You can construct a concrete instance of `ThirdPartyIntegrationArrayInput` via:
//
//	ThirdPartyIntegrationArray{ ThirdPartyIntegrationArgs{...} }
type ThirdPartyIntegrationArrayInput interface {
	pulumi.Input

	ToThirdPartyIntegrationArrayOutput() ThirdPartyIntegrationArrayOutput
	ToThirdPartyIntegrationArrayOutputWithContext(context.Context) ThirdPartyIntegrationArrayOutput
}

type ThirdPartyIntegrationArray []ThirdPartyIntegrationInput

func (ThirdPartyIntegrationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ThirdPartyIntegration)(nil)).Elem()
}

func (i ThirdPartyIntegrationArray) ToThirdPartyIntegrationArrayOutput() ThirdPartyIntegrationArrayOutput {
	return i.ToThirdPartyIntegrationArrayOutputWithContext(context.Background())
}

func (i ThirdPartyIntegrationArray) ToThirdPartyIntegrationArrayOutputWithContext(ctx context.Context) ThirdPartyIntegrationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ThirdPartyIntegrationArrayOutput)
}

// ThirdPartyIntegrationMapInput is an input type that accepts ThirdPartyIntegrationMap and ThirdPartyIntegrationMapOutput values.
// You can construct a concrete instance of `ThirdPartyIntegrationMapInput` via:
//
//	ThirdPartyIntegrationMap{ "key": ThirdPartyIntegrationArgs{...} }
type ThirdPartyIntegrationMapInput interface {
	pulumi.Input

	ToThirdPartyIntegrationMapOutput() ThirdPartyIntegrationMapOutput
	ToThirdPartyIntegrationMapOutputWithContext(context.Context) ThirdPartyIntegrationMapOutput
}

type ThirdPartyIntegrationMap map[string]ThirdPartyIntegrationInput

func (ThirdPartyIntegrationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ThirdPartyIntegration)(nil)).Elem()
}

func (i ThirdPartyIntegrationMap) ToThirdPartyIntegrationMapOutput() ThirdPartyIntegrationMapOutput {
	return i.ToThirdPartyIntegrationMapOutputWithContext(context.Background())
}

func (i ThirdPartyIntegrationMap) ToThirdPartyIntegrationMapOutputWithContext(ctx context.Context) ThirdPartyIntegrationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ThirdPartyIntegrationMapOutput)
}

type ThirdPartyIntegrationOutput struct{ *pulumi.OutputState }

func (ThirdPartyIntegrationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ThirdPartyIntegration)(nil)).Elem()
}

func (o ThirdPartyIntegrationOutput) ToThirdPartyIntegrationOutput() ThirdPartyIntegrationOutput {
	return o
}

func (o ThirdPartyIntegrationOutput) ToThirdPartyIntegrationOutputWithContext(ctx context.Context) ThirdPartyIntegrationOutput {
	return o
}

// Unique identifier of your New Relic account.
func (o ThirdPartyIntegrationOutput) AccountId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ThirdPartyIntegration) pulumi.StringPtrOutput { return v.AccountId }).(pulumi.StringPtrOutput)
}

// Your API Key.
func (o ThirdPartyIntegrationOutput) ApiKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ThirdPartyIntegration) pulumi.StringPtrOutput { return v.ApiKey }).(pulumi.StringPtrOutput)
}

// Your API Token.
func (o ThirdPartyIntegrationOutput) ApiToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ThirdPartyIntegration) pulumi.StringPtrOutput { return v.ApiToken }).(pulumi.StringPtrOutput)
}

func (o ThirdPartyIntegrationOutput) ChannelName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ThirdPartyIntegration) pulumi.StringPtrOutput { return v.ChannelName }).(pulumi.StringPtrOutput)
}

// Whether your cluster has Prometheus enabled.
func (o ThirdPartyIntegrationOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ThirdPartyIntegration) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// Your Flowdock Flow name.
func (o ThirdPartyIntegrationOutput) FlowName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ThirdPartyIntegration) pulumi.StringPtrOutput { return v.FlowName }).(pulumi.StringPtrOutput)
}

// Your License Key.
func (o ThirdPartyIntegrationOutput) LicenseKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ThirdPartyIntegration) pulumi.StringPtrOutput { return v.LicenseKey }).(pulumi.StringPtrOutput)
}

// Your Microsoft Teams incoming webhook URL.
// * `PROMETHEUS`
func (o ThirdPartyIntegrationOutput) MicrosoftTeamsWebhookUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ThirdPartyIntegration) pulumi.StringPtrOutput { return v.MicrosoftTeamsWebhookUrl }).(pulumi.StringPtrOutput)
}

// Your Flowdock organization name.
// * `WEBHOOK`
func (o ThirdPartyIntegrationOutput) OrgName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ThirdPartyIntegration) pulumi.StringPtrOutput { return v.OrgName }).(pulumi.StringPtrOutput)
}

// Your Prometheus password.
func (o ThirdPartyIntegrationOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ThirdPartyIntegration) pulumi.StringPtrOutput { return v.Password }).(pulumi.StringPtrOutput)
}

// The unique ID for the project to get all Third-Party service integrations
func (o ThirdPartyIntegrationOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *ThirdPartyIntegration) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// Your Insights Query Key.
// * `OPS_GENIE`
func (o ThirdPartyIntegrationOutput) ReadToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ThirdPartyIntegration) pulumi.StringPtrOutput { return v.ReadToken }).(pulumi.StringPtrOutput)
}

// Indicates which API URL to use, either US or EU. Opsgenie will use US by default.
// * `VICTOR_OPS`
func (o ThirdPartyIntegrationOutput) Region() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ThirdPartyIntegration) pulumi.StringPtrOutput { return v.Region }).(pulumi.StringPtrOutput)
}

// An optional field for your Routing Key.
// * `FLOWDOCK`
func (o ThirdPartyIntegrationOutput) RoutingKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ThirdPartyIntegration) pulumi.StringPtrOutput { return v.RoutingKey }).(pulumi.StringPtrOutput)
}

// Your Prometheus protocol scheme configured for requests.
func (o ThirdPartyIntegrationOutput) Scheme() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ThirdPartyIntegration) pulumi.StringPtrOutput { return v.Scheme }).(pulumi.StringPtrOutput)
}

// An optional field for your webhook secret.
// * `MICROSOFT_TEAMS`
func (o ThirdPartyIntegrationOutput) Secret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ThirdPartyIntegration) pulumi.StringPtrOutput { return v.Secret }).(pulumi.StringPtrOutput)
}

// Indicates which service discovery method is used, either file or http.
func (o ThirdPartyIntegrationOutput) ServiceDiscovery() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ThirdPartyIntegration) pulumi.StringPtrOutput { return v.ServiceDiscovery }).(pulumi.StringPtrOutput)
}

// Your Service Key.
// * `DATADOG`
func (o ThirdPartyIntegrationOutput) ServiceKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ThirdPartyIntegration) pulumi.StringPtrOutput { return v.ServiceKey }).(pulumi.StringPtrOutput)
}

func (o ThirdPartyIntegrationOutput) TeamName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ThirdPartyIntegration) pulumi.StringPtrOutput { return v.TeamName }).(pulumi.StringPtrOutput)
}

// Third-Party Integration Settings type
// * PAGER_DUTY
// * DATADOG
// * NEW_RELIC
// * OPS_GENIE
// * VICTOR_OPS
// * FLOWDOCK
// * WEBHOOK
func (o ThirdPartyIntegrationOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ThirdPartyIntegration) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Your webhook URL.
func (o ThirdPartyIntegrationOutput) Url() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ThirdPartyIntegration) pulumi.StringPtrOutput { return v.Url }).(pulumi.StringPtrOutput)
}

// Your Prometheus username.
func (o ThirdPartyIntegrationOutput) UserName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ThirdPartyIntegration) pulumi.StringPtrOutput { return v.UserName }).(pulumi.StringPtrOutput)
}

// Your Insights Insert Key.
func (o ThirdPartyIntegrationOutput) WriteToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ThirdPartyIntegration) pulumi.StringPtrOutput { return v.WriteToken }).(pulumi.StringPtrOutput)
}

type ThirdPartyIntegrationArrayOutput struct{ *pulumi.OutputState }

func (ThirdPartyIntegrationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ThirdPartyIntegration)(nil)).Elem()
}

func (o ThirdPartyIntegrationArrayOutput) ToThirdPartyIntegrationArrayOutput() ThirdPartyIntegrationArrayOutput {
	return o
}

func (o ThirdPartyIntegrationArrayOutput) ToThirdPartyIntegrationArrayOutputWithContext(ctx context.Context) ThirdPartyIntegrationArrayOutput {
	return o
}

func (o ThirdPartyIntegrationArrayOutput) Index(i pulumi.IntInput) ThirdPartyIntegrationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ThirdPartyIntegration {
		return vs[0].([]*ThirdPartyIntegration)[vs[1].(int)]
	}).(ThirdPartyIntegrationOutput)
}

type ThirdPartyIntegrationMapOutput struct{ *pulumi.OutputState }

func (ThirdPartyIntegrationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ThirdPartyIntegration)(nil)).Elem()
}

func (o ThirdPartyIntegrationMapOutput) ToThirdPartyIntegrationMapOutput() ThirdPartyIntegrationMapOutput {
	return o
}

func (o ThirdPartyIntegrationMapOutput) ToThirdPartyIntegrationMapOutputWithContext(ctx context.Context) ThirdPartyIntegrationMapOutput {
	return o
}

func (o ThirdPartyIntegrationMapOutput) MapIndex(k pulumi.StringInput) ThirdPartyIntegrationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ThirdPartyIntegration {
		return vs[0].(map[string]*ThirdPartyIntegration)[vs[1].(string)]
	}).(ThirdPartyIntegrationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ThirdPartyIntegrationInput)(nil)).Elem(), &ThirdPartyIntegration{})
	pulumi.RegisterInputType(reflect.TypeOf((*ThirdPartyIntegrationArrayInput)(nil)).Elem(), ThirdPartyIntegrationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ThirdPartyIntegrationMapInput)(nil)).Elem(), ThirdPartyIntegrationMap{})
	pulumi.RegisterOutputType(ThirdPartyIntegrationOutput{})
	pulumi.RegisterOutputType(ThirdPartyIntegrationArrayOutput{})
	pulumi.RegisterOutputType(ThirdPartyIntegrationMapOutput{})
}
