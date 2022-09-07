// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.mongodbatlas;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.mongodbatlas.ProjectArgs;
import com.pulumi.mongodbatlas.Utilities;
import com.pulumi.mongodbatlas.inputs.ProjectState;
import com.pulumi.mongodbatlas.outputs.ProjectApiKey;
import com.pulumi.mongodbatlas.outputs.ProjectTeam;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
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
 * import com.pulumi.mongodbatlas.Project;
 * import com.pulumi.mongodbatlas.ProjectArgs;
 * import com.pulumi.mongodbatlas.inputs.ProjectApiKeyArgs;
 * import com.pulumi.mongodbatlas.inputs.ProjectTeamArgs;
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
 *         var test = new Project(&#34;test&#34;, ProjectArgs.builder()        
 *             .apiKeys(ProjectApiKeyArgs.builder()
 *                 .apiKeyId(&#34;61003b299dda8d54a9d7d10c&#34;)
 *                 .roleNames(&#34;GROUP_READ_ONLY&#34;)
 *                 .build())
 *             .isCollectDatabaseSpecificsStatisticsEnabled(true)
 *             .isDataExplorerEnabled(true)
 *             .isPerformanceAdvisorEnabled(true)
 *             .isRealtimePerformancePanelEnabled(true)
 *             .isSchemaAdvisorEnabled(true)
 *             .orgId(&#34;&lt;ORG_ID&gt;&#34;)
 *             .projectOwnerId(&#34;&lt;OWNER_ACCOUNT_ID&gt;&#34;)
 *             .teams(            
 *                 ProjectTeamArgs.builder()
 *                     .roleNames(&#34;GROUP_OWNER&#34;)
 *                     .teamId(&#34;5e0fa8c99ccf641c722fe645&#34;)
 *                     .build(),
 *                 ProjectTeamArgs.builder()
 *                     .roleNames(                    
 *                         &#34;GROUP_READ_ONLY&#34;,
 *                         &#34;GROUP_DATA_ACCESS_READ_WRITE&#34;)
 *                     .teamId(&#34;5e1dd7b4f2a30ba80a70cd4rw&#34;)
 *                     .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Project must be imported using project ID, e.g.
 * 
 * ```sh
 *  $ pulumi import mongodbatlas:index/project:Project my_project 5d09d6a59ccf6445652a444a
 * ```
 * 
 *  For more information see[MongoDB Atlas API Reference.](https://docs.atlas.mongodb.com/reference/api/projects/) - [and MongoDB Atlas API - Teams](https://docs.atlas.mongodb.com/reference/api/teams/) Documentation for more information.
 * 
 */
@ResourceType(type="mongodbatlas:index/project:Project")
public class Project extends com.pulumi.resources.CustomResource {
    @Export(name="apiKeys", type=List.class, parameters={ProjectApiKey.class})
    private Output<List<ProjectApiKey>> apiKeys;

    public Output<List<ProjectApiKey>> apiKeys() {
        return this.apiKeys;
    }
    /**
     * The number of Atlas clusters deployed in the project..
     * 
     */
    @Export(name="clusterCount", type=Integer.class, parameters={})
    private Output<Integer> clusterCount;

    /**
     * @return The number of Atlas clusters deployed in the project..
     * 
     */
    public Output<Integer> clusterCount() {
        return this.clusterCount;
    }
    /**
     * The ISO-8601-formatted timestamp of when Atlas created the project..
     * 
     */
    @Export(name="created", type=String.class, parameters={})
    private Output<String> created;

    /**
     * @return The ISO-8601-formatted timestamp of when Atlas created the project..
     * 
     */
    public Output<String> created() {
        return this.created;
    }
    /**
     * Flag that indicates whether to enable statistics in [cluster metrics](https://www.mongodb.com/docs/atlas/monitor-cluster-metrics/) collection for the project.
     * 
     */
    @Export(name="isCollectDatabaseSpecificsStatisticsEnabled", type=Boolean.class, parameters={})
    private Output<Boolean> isCollectDatabaseSpecificsStatisticsEnabled;

    /**
     * @return Flag that indicates whether to enable statistics in [cluster metrics](https://www.mongodb.com/docs/atlas/monitor-cluster-metrics/) collection for the project.
     * 
     */
    public Output<Boolean> isCollectDatabaseSpecificsStatisticsEnabled() {
        return this.isCollectDatabaseSpecificsStatisticsEnabled;
    }
    /**
     * Flag that indicates whether to enable Data Explorer for the project. If enabled, you can query your database with an easy to use interface.  When Data Explorer is disabled, you cannot terminate slow operations from the [Real-Time Performance Panel](https://www.mongodb.com/docs/atlas/real-time-performance-panel/#std-label-real-time-metrics-status-tab) or create indexes from the [Performance Advisor](https://www.mongodb.com/docs/atlas/performance-advisor/#std-label-performance-advisor). You can still view Performance Advisor recommendations, but you must create those indexes from [mongosh](https://www.mongodb.com/docs/mongodb-shell/#mongodb-binary-bin.mongosh).
     * 
     */
    @Export(name="isDataExplorerEnabled", type=Boolean.class, parameters={})
    private Output<Boolean> isDataExplorerEnabled;

    /**
     * @return Flag that indicates whether to enable Data Explorer for the project. If enabled, you can query your database with an easy to use interface.  When Data Explorer is disabled, you cannot terminate slow operations from the [Real-Time Performance Panel](https://www.mongodb.com/docs/atlas/real-time-performance-panel/#std-label-real-time-metrics-status-tab) or create indexes from the [Performance Advisor](https://www.mongodb.com/docs/atlas/performance-advisor/#std-label-performance-advisor). You can still view Performance Advisor recommendations, but you must create those indexes from [mongosh](https://www.mongodb.com/docs/mongodb-shell/#mongodb-binary-bin.mongosh).
     * 
     */
    public Output<Boolean> isDataExplorerEnabled() {
        return this.isDataExplorerEnabled;
    }
    /**
     * Flag that indicates whether to enable Performance Advisor and Profiler for the project. If enabled, you can analyze database logs to recommend performance improvements.
     * 
     */
    @Export(name="isPerformanceAdvisorEnabled", type=Boolean.class, parameters={})
    private Output<Boolean> isPerformanceAdvisorEnabled;

    /**
     * @return Flag that indicates whether to enable Performance Advisor and Profiler for the project. If enabled, you can analyze database logs to recommend performance improvements.
     * 
     */
    public Output<Boolean> isPerformanceAdvisorEnabled() {
        return this.isPerformanceAdvisorEnabled;
    }
    /**
     * Flag that indicates whether to enable Real Time Performance Panel for the project. If enabled, you can see real time metrics from your MongoDB database.
     * 
     */
    @Export(name="isRealtimePerformancePanelEnabled", type=Boolean.class, parameters={})
    private Output<Boolean> isRealtimePerformancePanelEnabled;

    /**
     * @return Flag that indicates whether to enable Real Time Performance Panel for the project. If enabled, you can see real time metrics from your MongoDB database.
     * 
     */
    public Output<Boolean> isRealtimePerformancePanelEnabled() {
        return this.isRealtimePerformancePanelEnabled;
    }
    /**
     * Flag that indicates whether to enable Schema Advisor for the project. If enabled, you receive customized recommendations to optimize your data model and enhance performance. Disable this setting to disable schema suggestions in the [Performance Advisor](https://www.mongodb.com/docs/atlas/performance-advisor/#std-label-performance-advisor) and the [Data Explorer](https://www.mongodb.com/docs/atlas/atlas-ui/#std-label-atlas-ui).
     * 
     */
    @Export(name="isSchemaAdvisorEnabled", type=Boolean.class, parameters={})
    private Output<Boolean> isSchemaAdvisorEnabled;

    /**
     * @return Flag that indicates whether to enable Schema Advisor for the project. If enabled, you receive customized recommendations to optimize your data model and enhance performance. Disable this setting to disable schema suggestions in the [Performance Advisor](https://www.mongodb.com/docs/atlas/performance-advisor/#std-label-performance-advisor) and the [Data Explorer](https://www.mongodb.com/docs/atlas/atlas-ui/#std-label-atlas-ui).
     * 
     */
    public Output<Boolean> isSchemaAdvisorEnabled() {
        return this.isSchemaAdvisorEnabled;
    }
    /**
     * The name of the project you want to create. (Cannot be changed via this Provider after creation.)
     * 
     */
    @Export(name="name", type=String.class, parameters={})
    private Output<String> name;

    /**
     * @return The name of the project you want to create. (Cannot be changed via this Provider after creation.)
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The ID of the organization you want to create the project within.
     * 
     */
    @Export(name="orgId", type=String.class, parameters={})
    private Output<String> orgId;

    /**
     * @return The ID of the organization you want to create the project within.
     * 
     */
    public Output<String> orgId() {
        return this.orgId;
    }
    /**
     * Unique 24-hexadecimal digit string that identifies the Atlas user account to be granted the [Project Owner](https://docs.atlas.mongodb.com/reference/user-roles/#mongodb-authrole-Project-Owner) role on the specified project. If you set this parameter, it overrides the default value of the oldest [Organization Owner](https://docs.atlas.mongodb.com/reference/user-roles/#mongodb-authrole-Organization-Owner).
     * 
     */
    @Export(name="projectOwnerId", type=String.class, parameters={})
    private Output</* @Nullable */ String> projectOwnerId;

    /**
     * @return Unique 24-hexadecimal digit string that identifies the Atlas user account to be granted the [Project Owner](https://docs.atlas.mongodb.com/reference/user-roles/#mongodb-authrole-Project-Owner) role on the specified project. If you set this parameter, it overrides the default value of the oldest [Organization Owner](https://docs.atlas.mongodb.com/reference/user-roles/#mongodb-authrole-Organization-Owner).
     * 
     */
    public Output<Optional<String>> projectOwnerId() {
        return Codegen.optional(this.projectOwnerId);
    }
    @Export(name="teams", type=List.class, parameters={ProjectTeam.class})
    private Output</* @Nullable */ List<ProjectTeam>> teams;

    public Output<Optional<List<ProjectTeam>>> teams() {
        return Codegen.optional(this.teams);
    }
    /**
     * It allows users to disable the creation of the default alert settings. By default, this flag is set to true.
     * 
     */
    @Export(name="withDefaultAlertsSettings", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> withDefaultAlertsSettings;

    /**
     * @return It allows users to disable the creation of the default alert settings. By default, this flag is set to true.
     * 
     */
    public Output<Optional<Boolean>> withDefaultAlertsSettings() {
        return Codegen.optional(this.withDefaultAlertsSettings);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Project(String name) {
        this(name, ProjectArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Project(String name, ProjectArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Project(String name, ProjectArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("mongodbatlas:index/project:Project", name, args == null ? ProjectArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Project(String name, Output<String> id, @Nullable ProjectState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("mongodbatlas:index/project:Project", name, state, makeResourceOptions(options, id));
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
    public static Project get(String name, Output<String> id, @Nullable ProjectState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Project(name, id, state, options);
    }
}
