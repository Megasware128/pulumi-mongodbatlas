// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.mongodbatlas;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.mongodbatlas.ThirdPartyIntegrationArgs;
import com.pulumi.mongodbatlas.Utilities;
import com.pulumi.mongodbatlas.inputs.ThirdPartyIntegrationState;
import java.lang.Boolean;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.mongodbatlas.ThirdPartyIntegration;
 * import com.pulumi.mongodbatlas.ThirdPartyIntegrationArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var testFlowdock = new ThirdPartyIntegration(&#34;testFlowdock&#34;, ThirdPartyIntegrationArgs.builder()        
 *             .apiToken(&#34;&lt;API-TOKEN&gt;&#34;)
 *             .flowName(&#34;&lt;FLOW-NAME&gt;&#34;)
 *             .orgName(&#34;&lt;ORG-NAME&gt;&#34;)
 *             .projectId(&#34;&lt;PROJECT-ID&gt;&#34;)
 *             .type(&#34;FLOWDOCK&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Third-Party Integration Settings can be imported using project ID and the integration type, in the format `project_id`-`type`, e.g.
 * 
 * ```sh
 *  $ pulumi import mongodbatlas:index/thirdPartyIntegration:ThirdPartyIntegration my_user 1112222b3bf99403840e8934-OPS_GENIE
 * ```
 * 
 *  See [MongoDB Atlas API](https://docs.atlas.mongodb.com/reference/api/third-party-integration-settings-create/) Documentation for more information.
 * 
 */
@ResourceType(type="mongodbatlas:index/thirdPartyIntegration:ThirdPartyIntegration")
public class ThirdPartyIntegration extends com.pulumi.resources.CustomResource {
    /**
     * Unique identifier of your New Relic account.
     * 
     */
    @Export(name="accountId", type=String.class, parameters={})
    private Output</* @Nullable */ String> accountId;

    /**
     * @return Unique identifier of your New Relic account.
     * 
     */
    public Output<Optional<String>> accountId() {
        return Codegen.optional(this.accountId);
    }
    /**
     * Your API Key.
     * 
     */
    @Export(name="apiKey", type=String.class, parameters={})
    private Output</* @Nullable */ String> apiKey;

    /**
     * @return Your API Key.
     * 
     */
    public Output<Optional<String>> apiKey() {
        return Codegen.optional(this.apiKey);
    }
    /**
     * Your API Token.
     * 
     */
    @Export(name="apiToken", type=String.class, parameters={})
    private Output</* @Nullable */ String> apiToken;

    /**
     * @return Your API Token.
     * 
     */
    public Output<Optional<String>> apiToken() {
        return Codegen.optional(this.apiToken);
    }
    @Export(name="channelName", type=String.class, parameters={})
    private Output</* @Nullable */ String> channelName;

    public Output<Optional<String>> channelName() {
        return Codegen.optional(this.channelName);
    }
    /**
     * Whether your cluster has Prometheus enabled.
     * 
     */
    @Export(name="enabled", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> enabled;

    /**
     * @return Whether your cluster has Prometheus enabled.
     * 
     */
    public Output<Optional<Boolean>> enabled() {
        return Codegen.optional(this.enabled);
    }
    /**
     * Your Flowdock Flow name.
     * 
     */
    @Export(name="flowName", type=String.class, parameters={})
    private Output</* @Nullable */ String> flowName;

    /**
     * @return Your Flowdock Flow name.
     * 
     */
    public Output<Optional<String>> flowName() {
        return Codegen.optional(this.flowName);
    }
    /**
     * Your License Key.
     * 
     */
    @Export(name="licenseKey", type=String.class, parameters={})
    private Output</* @Nullable */ String> licenseKey;

    /**
     * @return Your License Key.
     * 
     */
    public Output<Optional<String>> licenseKey() {
        return Codegen.optional(this.licenseKey);
    }
    /**
     * Your Microsoft Teams incoming webhook URL.
     * * `PROMETHEUS`
     * 
     */
    @Export(name="microsoftTeamsWebhookUrl", type=String.class, parameters={})
    private Output</* @Nullable */ String> microsoftTeamsWebhookUrl;

    /**
     * @return Your Microsoft Teams incoming webhook URL.
     * * `PROMETHEUS`
     * 
     */
    public Output<Optional<String>> microsoftTeamsWebhookUrl() {
        return Codegen.optional(this.microsoftTeamsWebhookUrl);
    }
    /**
     * Your Flowdock organization name.
     * * `WEBHOOK`
     * 
     */
    @Export(name="orgName", type=String.class, parameters={})
    private Output</* @Nullable */ String> orgName;

    /**
     * @return Your Flowdock organization name.
     * * `WEBHOOK`
     * 
     */
    public Output<Optional<String>> orgName() {
        return Codegen.optional(this.orgName);
    }
    /**
     * Your Prometheus password.
     * 
     */
    @Export(name="password", type=String.class, parameters={})
    private Output</* @Nullable */ String> password;

    /**
     * @return Your Prometheus password.
     * 
     */
    public Output<Optional<String>> password() {
        return Codegen.optional(this.password);
    }
    /**
     * The unique ID for the project to get all Third-Party service integrations
     * 
     */
    @Export(name="projectId", type=String.class, parameters={})
    private Output<String> projectId;

    /**
     * @return The unique ID for the project to get all Third-Party service integrations
     * 
     */
    public Output<String> projectId() {
        return this.projectId;
    }
    /**
     * Your Insights Query Key.
     * * `OPS_GENIE`
     * 
     */
    @Export(name="readToken", type=String.class, parameters={})
    private Output</* @Nullable */ String> readToken;

    /**
     * @return Your Insights Query Key.
     * * `OPS_GENIE`
     * 
     */
    public Output<Optional<String>> readToken() {
        return Codegen.optional(this.readToken);
    }
    /**
     * Indicates which API URL to use, either US or EU. Opsgenie will use US by default.
     * * `VICTOR_OPS`
     * 
     */
    @Export(name="region", type=String.class, parameters={})
    private Output</* @Nullable */ String> region;

    /**
     * @return Indicates which API URL to use, either US or EU. Opsgenie will use US by default.
     * * `VICTOR_OPS`
     * 
     */
    public Output<Optional<String>> region() {
        return Codegen.optional(this.region);
    }
    /**
     * An optional field for your Routing Key.
     * * `FLOWDOCK`
     * 
     */
    @Export(name="routingKey", type=String.class, parameters={})
    private Output</* @Nullable */ String> routingKey;

    /**
     * @return An optional field for your Routing Key.
     * * `FLOWDOCK`
     * 
     */
    public Output<Optional<String>> routingKey() {
        return Codegen.optional(this.routingKey);
    }
    /**
     * Your Prometheus protocol scheme configured for requests.
     * 
     */
    @Export(name="scheme", type=String.class, parameters={})
    private Output</* @Nullable */ String> scheme;

    /**
     * @return Your Prometheus protocol scheme configured for requests.
     * 
     */
    public Output<Optional<String>> scheme() {
        return Codegen.optional(this.scheme);
    }
    /**
     * An optional field for your webhook secret.
     * * `MICROSOFT_TEAMS`
     * 
     */
    @Export(name="secret", type=String.class, parameters={})
    private Output</* @Nullable */ String> secret;

    /**
     * @return An optional field for your webhook secret.
     * * `MICROSOFT_TEAMS`
     * 
     */
    public Output<Optional<String>> secret() {
        return Codegen.optional(this.secret);
    }
    /**
     * Indicates which service discovery method is used, either file or http.
     * 
     */
    @Export(name="serviceDiscovery", type=String.class, parameters={})
    private Output</* @Nullable */ String> serviceDiscovery;

    /**
     * @return Indicates which service discovery method is used, either file or http.
     * 
     */
    public Output<Optional<String>> serviceDiscovery() {
        return Codegen.optional(this.serviceDiscovery);
    }
    /**
     * Your Service Key.
     * * `DATADOG`
     * 
     */
    @Export(name="serviceKey", type=String.class, parameters={})
    private Output</* @Nullable */ String> serviceKey;

    /**
     * @return Your Service Key.
     * * `DATADOG`
     * 
     */
    public Output<Optional<String>> serviceKey() {
        return Codegen.optional(this.serviceKey);
    }
    @Export(name="teamName", type=String.class, parameters={})
    private Output</* @Nullable */ String> teamName;

    public Output<Optional<String>> teamName() {
        return Codegen.optional(this.teamName);
    }
    /**
     * Third-Party Integration Settings type
     * * PAGER_DUTY
     * * DATADOG
     * * NEW_RELIC
     * * OPS_GENIE
     * * VICTOR_OPS
     * * FLOWDOCK
     * * WEBHOOK
     * 
     */
    @Export(name="type", type=String.class, parameters={})
    private Output<String> type;

    /**
     * @return Third-Party Integration Settings type
     * * PAGER_DUTY
     * * DATADOG
     * * NEW_RELIC
     * * OPS_GENIE
     * * VICTOR_OPS
     * * FLOWDOCK
     * * WEBHOOK
     * 
     */
    public Output<String> type() {
        return this.type;
    }
    /**
     * Your webhook URL.
     * 
     */
    @Export(name="url", type=String.class, parameters={})
    private Output</* @Nullable */ String> url;

    /**
     * @return Your webhook URL.
     * 
     */
    public Output<Optional<String>> url() {
        return Codegen.optional(this.url);
    }
    /**
     * Your Prometheus username.
     * 
     */
    @Export(name="userName", type=String.class, parameters={})
    private Output</* @Nullable */ String> userName;

    /**
     * @return Your Prometheus username.
     * 
     */
    public Output<Optional<String>> userName() {
        return Codegen.optional(this.userName);
    }
    /**
     * Your Insights Insert Key.
     * 
     */
    @Export(name="writeToken", type=String.class, parameters={})
    private Output</* @Nullable */ String> writeToken;

    /**
     * @return Your Insights Insert Key.
     * 
     */
    public Output<Optional<String>> writeToken() {
        return Codegen.optional(this.writeToken);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ThirdPartyIntegration(String name) {
        this(name, ThirdPartyIntegrationArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ThirdPartyIntegration(String name, ThirdPartyIntegrationArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ThirdPartyIntegration(String name, ThirdPartyIntegrationArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("mongodbatlas:index/thirdPartyIntegration:ThirdPartyIntegration", name, args == null ? ThirdPartyIntegrationArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private ThirdPartyIntegration(String name, Output<String> id, @Nullable ThirdPartyIntegrationState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("mongodbatlas:index/thirdPartyIntegration:ThirdPartyIntegration", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static ThirdPartyIntegration get(String name, Output<String> id, @Nullable ThirdPartyIntegrationState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ThirdPartyIntegration(name, id, state, options);
    }
}
