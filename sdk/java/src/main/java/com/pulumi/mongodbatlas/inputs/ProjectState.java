// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.mongodbatlas.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.mongodbatlas.inputs.ProjectApiKeyArgs;
import com.pulumi.mongodbatlas.inputs.ProjectTeamArgs;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ProjectState extends com.pulumi.resources.ResourceArgs {

    public static final ProjectState Empty = new ProjectState();

    @Import(name="apiKeys")
    private @Nullable Output<List<ProjectApiKeyArgs>> apiKeys;

    public Optional<Output<List<ProjectApiKeyArgs>>> apiKeys() {
        return Optional.ofNullable(this.apiKeys);
    }

    /**
     * The number of Atlas clusters deployed in the project..
     * 
     */
    @Import(name="clusterCount")
    private @Nullable Output<Integer> clusterCount;

    /**
     * @return The number of Atlas clusters deployed in the project..
     * 
     */
    public Optional<Output<Integer>> clusterCount() {
        return Optional.ofNullable(this.clusterCount);
    }

    /**
     * The ISO-8601-formatted timestamp of when Atlas created the project..
     * 
     */
    @Import(name="created")
    private @Nullable Output<String> created;

    /**
     * @return The ISO-8601-formatted timestamp of when Atlas created the project..
     * 
     */
    public Optional<Output<String>> created() {
        return Optional.ofNullable(this.created);
    }

    /**
     * Flag that indicates whether to enable statistics in [cluster metrics](https://www.mongodb.com/docs/atlas/monitor-cluster-metrics/) collection for the project.
     * 
     */
    @Import(name="isCollectDatabaseSpecificsStatisticsEnabled")
    private @Nullable Output<Boolean> isCollectDatabaseSpecificsStatisticsEnabled;

    /**
     * @return Flag that indicates whether to enable statistics in [cluster metrics](https://www.mongodb.com/docs/atlas/monitor-cluster-metrics/) collection for the project.
     * 
     */
    public Optional<Output<Boolean>> isCollectDatabaseSpecificsStatisticsEnabled() {
        return Optional.ofNullable(this.isCollectDatabaseSpecificsStatisticsEnabled);
    }

    /**
     * Flag that indicates whether to enable Data Explorer for the project. If enabled, you can query your database with an easy to use interface.  When Data Explorer is disabled, you cannot terminate slow operations from the [Real-Time Performance Panel](https://www.mongodb.com/docs/atlas/real-time-performance-panel/#std-label-real-time-metrics-status-tab) or create indexes from the [Performance Advisor](https://www.mongodb.com/docs/atlas/performance-advisor/#std-label-performance-advisor). You can still view Performance Advisor recommendations, but you must create those indexes from [mongosh](https://www.mongodb.com/docs/mongodb-shell/#mongodb-binary-bin.mongosh).
     * 
     */
    @Import(name="isDataExplorerEnabled")
    private @Nullable Output<Boolean> isDataExplorerEnabled;

    /**
     * @return Flag that indicates whether to enable Data Explorer for the project. If enabled, you can query your database with an easy to use interface.  When Data Explorer is disabled, you cannot terminate slow operations from the [Real-Time Performance Panel](https://www.mongodb.com/docs/atlas/real-time-performance-panel/#std-label-real-time-metrics-status-tab) or create indexes from the [Performance Advisor](https://www.mongodb.com/docs/atlas/performance-advisor/#std-label-performance-advisor). You can still view Performance Advisor recommendations, but you must create those indexes from [mongosh](https://www.mongodb.com/docs/mongodb-shell/#mongodb-binary-bin.mongosh).
     * 
     */
    public Optional<Output<Boolean>> isDataExplorerEnabled() {
        return Optional.ofNullable(this.isDataExplorerEnabled);
    }

    /**
     * Flag that indicates whether to enable Performance Advisor and Profiler for the project. If enabled, you can analyze database logs to recommend performance improvements.
     * 
     */
    @Import(name="isPerformanceAdvisorEnabled")
    private @Nullable Output<Boolean> isPerformanceAdvisorEnabled;

    /**
     * @return Flag that indicates whether to enable Performance Advisor and Profiler for the project. If enabled, you can analyze database logs to recommend performance improvements.
     * 
     */
    public Optional<Output<Boolean>> isPerformanceAdvisorEnabled() {
        return Optional.ofNullable(this.isPerformanceAdvisorEnabled);
    }

    /**
     * Flag that indicates whether to enable Real Time Performance Panel for the project. If enabled, you can see real time metrics from your MongoDB database.
     * 
     */
    @Import(name="isRealtimePerformancePanelEnabled")
    private @Nullable Output<Boolean> isRealtimePerformancePanelEnabled;

    /**
     * @return Flag that indicates whether to enable Real Time Performance Panel for the project. If enabled, you can see real time metrics from your MongoDB database.
     * 
     */
    public Optional<Output<Boolean>> isRealtimePerformancePanelEnabled() {
        return Optional.ofNullable(this.isRealtimePerformancePanelEnabled);
    }

    /**
     * Flag that indicates whether to enable Schema Advisor for the project. If enabled, you receive customized recommendations to optimize your data model and enhance performance. Disable this setting to disable schema suggestions in the [Performance Advisor](https://www.mongodb.com/docs/atlas/performance-advisor/#std-label-performance-advisor) and the [Data Explorer](https://www.mongodb.com/docs/atlas/atlas-ui/#std-label-atlas-ui).
     * 
     */
    @Import(name="isSchemaAdvisorEnabled")
    private @Nullable Output<Boolean> isSchemaAdvisorEnabled;

    /**
     * @return Flag that indicates whether to enable Schema Advisor for the project. If enabled, you receive customized recommendations to optimize your data model and enhance performance. Disable this setting to disable schema suggestions in the [Performance Advisor](https://www.mongodb.com/docs/atlas/performance-advisor/#std-label-performance-advisor) and the [Data Explorer](https://www.mongodb.com/docs/atlas/atlas-ui/#std-label-atlas-ui).
     * 
     */
    public Optional<Output<Boolean>> isSchemaAdvisorEnabled() {
        return Optional.ofNullable(this.isSchemaAdvisorEnabled);
    }

    /**
     * The name of the project you want to create. (Cannot be changed via this Provider after creation.)
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name of the project you want to create. (Cannot be changed via this Provider after creation.)
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The ID of the organization you want to create the project within.
     * 
     */
    @Import(name="orgId")
    private @Nullable Output<String> orgId;

    /**
     * @return The ID of the organization you want to create the project within.
     * 
     */
    public Optional<Output<String>> orgId() {
        return Optional.ofNullable(this.orgId);
    }

    /**
     * Unique 24-hexadecimal digit string that identifies the Atlas user account to be granted the [Project Owner](https://docs.atlas.mongodb.com/reference/user-roles/#mongodb-authrole-Project-Owner) role on the specified project. If you set this parameter, it overrides the default value of the oldest [Organization Owner](https://docs.atlas.mongodb.com/reference/user-roles/#mongodb-authrole-Organization-Owner).
     * 
     */
    @Import(name="projectOwnerId")
    private @Nullable Output<String> projectOwnerId;

    /**
     * @return Unique 24-hexadecimal digit string that identifies the Atlas user account to be granted the [Project Owner](https://docs.atlas.mongodb.com/reference/user-roles/#mongodb-authrole-Project-Owner) role on the specified project. If you set this parameter, it overrides the default value of the oldest [Organization Owner](https://docs.atlas.mongodb.com/reference/user-roles/#mongodb-authrole-Organization-Owner).
     * 
     */
    public Optional<Output<String>> projectOwnerId() {
        return Optional.ofNullable(this.projectOwnerId);
    }

    @Import(name="teams")
    private @Nullable Output<List<ProjectTeamArgs>> teams;

    public Optional<Output<List<ProjectTeamArgs>>> teams() {
        return Optional.ofNullable(this.teams);
    }

    /**
     * It allows users to disable the creation of the default alert settings. By default, this flag is set to true.
     * 
     */
    @Import(name="withDefaultAlertsSettings")
    private @Nullable Output<Boolean> withDefaultAlertsSettings;

    /**
     * @return It allows users to disable the creation of the default alert settings. By default, this flag is set to true.
     * 
     */
    public Optional<Output<Boolean>> withDefaultAlertsSettings() {
        return Optional.ofNullable(this.withDefaultAlertsSettings);
    }

    private ProjectState() {}

    private ProjectState(ProjectState $) {
        this.apiKeys = $.apiKeys;
        this.clusterCount = $.clusterCount;
        this.created = $.created;
        this.isCollectDatabaseSpecificsStatisticsEnabled = $.isCollectDatabaseSpecificsStatisticsEnabled;
        this.isDataExplorerEnabled = $.isDataExplorerEnabled;
        this.isPerformanceAdvisorEnabled = $.isPerformanceAdvisorEnabled;
        this.isRealtimePerformancePanelEnabled = $.isRealtimePerformancePanelEnabled;
        this.isSchemaAdvisorEnabled = $.isSchemaAdvisorEnabled;
        this.name = $.name;
        this.orgId = $.orgId;
        this.projectOwnerId = $.projectOwnerId;
        this.teams = $.teams;
        this.withDefaultAlertsSettings = $.withDefaultAlertsSettings;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ProjectState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ProjectState $;

        public Builder() {
            $ = new ProjectState();
        }

        public Builder(ProjectState defaults) {
            $ = new ProjectState(Objects.requireNonNull(defaults));
        }

        public Builder apiKeys(@Nullable Output<List<ProjectApiKeyArgs>> apiKeys) {
            $.apiKeys = apiKeys;
            return this;
        }

        public Builder apiKeys(List<ProjectApiKeyArgs> apiKeys) {
            return apiKeys(Output.of(apiKeys));
        }

        public Builder apiKeys(ProjectApiKeyArgs... apiKeys) {
            return apiKeys(List.of(apiKeys));
        }

        /**
         * @param clusterCount The number of Atlas clusters deployed in the project..
         * 
         * @return builder
         * 
         */
        public Builder clusterCount(@Nullable Output<Integer> clusterCount) {
            $.clusterCount = clusterCount;
            return this;
        }

        /**
         * @param clusterCount The number of Atlas clusters deployed in the project..
         * 
         * @return builder
         * 
         */
        public Builder clusterCount(Integer clusterCount) {
            return clusterCount(Output.of(clusterCount));
        }

        /**
         * @param created The ISO-8601-formatted timestamp of when Atlas created the project..
         * 
         * @return builder
         * 
         */
        public Builder created(@Nullable Output<String> created) {
            $.created = created;
            return this;
        }

        /**
         * @param created The ISO-8601-formatted timestamp of when Atlas created the project..
         * 
         * @return builder
         * 
         */
        public Builder created(String created) {
            return created(Output.of(created));
        }

        /**
         * @param isCollectDatabaseSpecificsStatisticsEnabled Flag that indicates whether to enable statistics in [cluster metrics](https://www.mongodb.com/docs/atlas/monitor-cluster-metrics/) collection for the project.
         * 
         * @return builder
         * 
         */
        public Builder isCollectDatabaseSpecificsStatisticsEnabled(@Nullable Output<Boolean> isCollectDatabaseSpecificsStatisticsEnabled) {
            $.isCollectDatabaseSpecificsStatisticsEnabled = isCollectDatabaseSpecificsStatisticsEnabled;
            return this;
        }

        /**
         * @param isCollectDatabaseSpecificsStatisticsEnabled Flag that indicates whether to enable statistics in [cluster metrics](https://www.mongodb.com/docs/atlas/monitor-cluster-metrics/) collection for the project.
         * 
         * @return builder
         * 
         */
        public Builder isCollectDatabaseSpecificsStatisticsEnabled(Boolean isCollectDatabaseSpecificsStatisticsEnabled) {
            return isCollectDatabaseSpecificsStatisticsEnabled(Output.of(isCollectDatabaseSpecificsStatisticsEnabled));
        }

        /**
         * @param isDataExplorerEnabled Flag that indicates whether to enable Data Explorer for the project. If enabled, you can query your database with an easy to use interface.  When Data Explorer is disabled, you cannot terminate slow operations from the [Real-Time Performance Panel](https://www.mongodb.com/docs/atlas/real-time-performance-panel/#std-label-real-time-metrics-status-tab) or create indexes from the [Performance Advisor](https://www.mongodb.com/docs/atlas/performance-advisor/#std-label-performance-advisor). You can still view Performance Advisor recommendations, but you must create those indexes from [mongosh](https://www.mongodb.com/docs/mongodb-shell/#mongodb-binary-bin.mongosh).
         * 
         * @return builder
         * 
         */
        public Builder isDataExplorerEnabled(@Nullable Output<Boolean> isDataExplorerEnabled) {
            $.isDataExplorerEnabled = isDataExplorerEnabled;
            return this;
        }

        /**
         * @param isDataExplorerEnabled Flag that indicates whether to enable Data Explorer for the project. If enabled, you can query your database with an easy to use interface.  When Data Explorer is disabled, you cannot terminate slow operations from the [Real-Time Performance Panel](https://www.mongodb.com/docs/atlas/real-time-performance-panel/#std-label-real-time-metrics-status-tab) or create indexes from the [Performance Advisor](https://www.mongodb.com/docs/atlas/performance-advisor/#std-label-performance-advisor). You can still view Performance Advisor recommendations, but you must create those indexes from [mongosh](https://www.mongodb.com/docs/mongodb-shell/#mongodb-binary-bin.mongosh).
         * 
         * @return builder
         * 
         */
        public Builder isDataExplorerEnabled(Boolean isDataExplorerEnabled) {
            return isDataExplorerEnabled(Output.of(isDataExplorerEnabled));
        }

        /**
         * @param isPerformanceAdvisorEnabled Flag that indicates whether to enable Performance Advisor and Profiler for the project. If enabled, you can analyze database logs to recommend performance improvements.
         * 
         * @return builder
         * 
         */
        public Builder isPerformanceAdvisorEnabled(@Nullable Output<Boolean> isPerformanceAdvisorEnabled) {
            $.isPerformanceAdvisorEnabled = isPerformanceAdvisorEnabled;
            return this;
        }

        /**
         * @param isPerformanceAdvisorEnabled Flag that indicates whether to enable Performance Advisor and Profiler for the project. If enabled, you can analyze database logs to recommend performance improvements.
         * 
         * @return builder
         * 
         */
        public Builder isPerformanceAdvisorEnabled(Boolean isPerformanceAdvisorEnabled) {
            return isPerformanceAdvisorEnabled(Output.of(isPerformanceAdvisorEnabled));
        }

        /**
         * @param isRealtimePerformancePanelEnabled Flag that indicates whether to enable Real Time Performance Panel for the project. If enabled, you can see real time metrics from your MongoDB database.
         * 
         * @return builder
         * 
         */
        public Builder isRealtimePerformancePanelEnabled(@Nullable Output<Boolean> isRealtimePerformancePanelEnabled) {
            $.isRealtimePerformancePanelEnabled = isRealtimePerformancePanelEnabled;
            return this;
        }

        /**
         * @param isRealtimePerformancePanelEnabled Flag that indicates whether to enable Real Time Performance Panel for the project. If enabled, you can see real time metrics from your MongoDB database.
         * 
         * @return builder
         * 
         */
        public Builder isRealtimePerformancePanelEnabled(Boolean isRealtimePerformancePanelEnabled) {
            return isRealtimePerformancePanelEnabled(Output.of(isRealtimePerformancePanelEnabled));
        }

        /**
         * @param isSchemaAdvisorEnabled Flag that indicates whether to enable Schema Advisor for the project. If enabled, you receive customized recommendations to optimize your data model and enhance performance. Disable this setting to disable schema suggestions in the [Performance Advisor](https://www.mongodb.com/docs/atlas/performance-advisor/#std-label-performance-advisor) and the [Data Explorer](https://www.mongodb.com/docs/atlas/atlas-ui/#std-label-atlas-ui).
         * 
         * @return builder
         * 
         */
        public Builder isSchemaAdvisorEnabled(@Nullable Output<Boolean> isSchemaAdvisorEnabled) {
            $.isSchemaAdvisorEnabled = isSchemaAdvisorEnabled;
            return this;
        }

        /**
         * @param isSchemaAdvisorEnabled Flag that indicates whether to enable Schema Advisor for the project. If enabled, you receive customized recommendations to optimize your data model and enhance performance. Disable this setting to disable schema suggestions in the [Performance Advisor](https://www.mongodb.com/docs/atlas/performance-advisor/#std-label-performance-advisor) and the [Data Explorer](https://www.mongodb.com/docs/atlas/atlas-ui/#std-label-atlas-ui).
         * 
         * @return builder
         * 
         */
        public Builder isSchemaAdvisorEnabled(Boolean isSchemaAdvisorEnabled) {
            return isSchemaAdvisorEnabled(Output.of(isSchemaAdvisorEnabled));
        }

        /**
         * @param name The name of the project you want to create. (Cannot be changed via this Provider after creation.)
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the project you want to create. (Cannot be changed via this Provider after creation.)
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param orgId The ID of the organization you want to create the project within.
         * 
         * @return builder
         * 
         */
        public Builder orgId(@Nullable Output<String> orgId) {
            $.orgId = orgId;
            return this;
        }

        /**
         * @param orgId The ID of the organization you want to create the project within.
         * 
         * @return builder
         * 
         */
        public Builder orgId(String orgId) {
            return orgId(Output.of(orgId));
        }

        /**
         * @param projectOwnerId Unique 24-hexadecimal digit string that identifies the Atlas user account to be granted the [Project Owner](https://docs.atlas.mongodb.com/reference/user-roles/#mongodb-authrole-Project-Owner) role on the specified project. If you set this parameter, it overrides the default value of the oldest [Organization Owner](https://docs.atlas.mongodb.com/reference/user-roles/#mongodb-authrole-Organization-Owner).
         * 
         * @return builder
         * 
         */
        public Builder projectOwnerId(@Nullable Output<String> projectOwnerId) {
            $.projectOwnerId = projectOwnerId;
            return this;
        }

        /**
         * @param projectOwnerId Unique 24-hexadecimal digit string that identifies the Atlas user account to be granted the [Project Owner](https://docs.atlas.mongodb.com/reference/user-roles/#mongodb-authrole-Project-Owner) role on the specified project. If you set this parameter, it overrides the default value of the oldest [Organization Owner](https://docs.atlas.mongodb.com/reference/user-roles/#mongodb-authrole-Organization-Owner).
         * 
         * @return builder
         * 
         */
        public Builder projectOwnerId(String projectOwnerId) {
            return projectOwnerId(Output.of(projectOwnerId));
        }

        public Builder teams(@Nullable Output<List<ProjectTeamArgs>> teams) {
            $.teams = teams;
            return this;
        }

        public Builder teams(List<ProjectTeamArgs> teams) {
            return teams(Output.of(teams));
        }

        public Builder teams(ProjectTeamArgs... teams) {
            return teams(List.of(teams));
        }

        /**
         * @param withDefaultAlertsSettings It allows users to disable the creation of the default alert settings. By default, this flag is set to true.
         * 
         * @return builder
         * 
         */
        public Builder withDefaultAlertsSettings(@Nullable Output<Boolean> withDefaultAlertsSettings) {
            $.withDefaultAlertsSettings = withDefaultAlertsSettings;
            return this;
        }

        /**
         * @param withDefaultAlertsSettings It allows users to disable the creation of the default alert settings. By default, this flag is set to true.
         * 
         * @return builder
         * 
         */
        public Builder withDefaultAlertsSettings(Boolean withDefaultAlertsSettings) {
            return withDefaultAlertsSettings(Output.of(withDefaultAlertsSettings));
        }

        public ProjectState build() {
            return $;
        }
    }

}
