// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.mongodbatlas.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.mongodbatlas.inputs.ServerlessInstanceLinkArgs;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ServerlessInstanceState extends com.pulumi.resources.ResourceArgs {

    public static final ServerlessInstanceState Empty = new ServerlessInstanceState();

    /**
     * Public `mongodb+srv://` connection string that you can use to connect to this serverless instance.
     * 
     */
    @Import(name="connectionStringsStandardSrv")
    private @Nullable Output<String> connectionStringsStandardSrv;

    /**
     * @return Public `mongodb+srv://` connection string that you can use to connect to this serverless instance.
     * 
     */
    public Optional<Output<String>> connectionStringsStandardSrv() {
        return Optional.ofNullable(this.connectionStringsStandardSrv);
    }

    /**
     * Timestamp that indicates when MongoDB Cloud created the serverless instance. The timestamp displays in the ISO 8601 date and time format in UTC.
     * 
     */
    @Import(name="createDate")
    private @Nullable Output<String> createDate;

    /**
     * @return Timestamp that indicates when MongoDB Cloud created the serverless instance. The timestamp displays in the ISO 8601 date and time format in UTC.
     * 
     */
    public Optional<Output<String>> createDate() {
        return Optional.ofNullable(this.createDate);
    }

    @Import(name="links")
    private @Nullable Output<List<ServerlessInstanceLinkArgs>> links;

    public Optional<Output<List<ServerlessInstanceLinkArgs>>> links() {
        return Optional.ofNullable(this.links);
    }

    /**
     * Version of MongoDB that the serverless instance runs, in `&lt;major version&gt;`.`&lt;minor version&gt;` format.
     * 
     */
    @Import(name="mongoDbVersion")
    private @Nullable Output<String> mongoDbVersion;

    /**
     * @return Version of MongoDB that the serverless instance runs, in `&lt;major version&gt;`.`&lt;minor version&gt;` format.
     * 
     */
    public Optional<Output<String>> mongoDbVersion() {
        return Optional.ofNullable(this.mongoDbVersion);
    }

    /**
     * Human-readable label that identifies the serverless instance.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Human-readable label that identifies the serverless instance.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The ID of the organization or project you want to create the serverless instance within.
     * 
     */
    @Import(name="projectId")
    private @Nullable Output<String> projectId;

    /**
     * @return The ID of the organization or project you want to create the serverless instance within.
     * 
     */
    public Optional<Output<String>> projectId() {
        return Optional.ofNullable(this.projectId);
    }

    /**
     * Cloud service provider on which MongoDB Cloud provisioned the serverless instance.
     * 
     */
    @Import(name="providerSettingsBackingProviderName")
    private @Nullable Output<String> providerSettingsBackingProviderName;

    /**
     * @return Cloud service provider on which MongoDB Cloud provisioned the serverless instance.
     * 
     */
    public Optional<Output<String>> providerSettingsBackingProviderName() {
        return Optional.ofNullable(this.providerSettingsBackingProviderName);
    }

    /**
     * Cloud service provider that applies to the provisioned the serverless instance.
     * 
     */
    @Import(name="providerSettingsProviderName")
    private @Nullable Output<String> providerSettingsProviderName;

    /**
     * @return Cloud service provider that applies to the provisioned the serverless instance.
     * 
     */
    public Optional<Output<String>> providerSettingsProviderName() {
        return Optional.ofNullable(this.providerSettingsProviderName);
    }

    /**
     * Human-readable label that identifies the physical location of your MongoDB serverless instance. The region you choose can affect network latency for clients accessing your databases.
     * 
     */
    @Import(name="providerSettingsRegionName")
    private @Nullable Output<String> providerSettingsRegionName;

    /**
     * @return Human-readable label that identifies the physical location of your MongoDB serverless instance. The region you choose can affect network latency for clients accessing your databases.
     * 
     */
    public Optional<Output<String>> providerSettingsRegionName() {
        return Optional.ofNullable(this.providerSettingsRegionName);
    }

    /**
     * Stage of deployment of this serverless instance when the resource made its request.
     * 
     */
    @Import(name="stateName")
    private @Nullable Output<String> stateName;

    /**
     * @return Stage of deployment of this serverless instance when the resource made its request.
     * 
     */
    public Optional<Output<String>> stateName() {
        return Optional.ofNullable(this.stateName);
    }

    private ServerlessInstanceState() {}

    private ServerlessInstanceState(ServerlessInstanceState $) {
        this.connectionStringsStandardSrv = $.connectionStringsStandardSrv;
        this.createDate = $.createDate;
        this.links = $.links;
        this.mongoDbVersion = $.mongoDbVersion;
        this.name = $.name;
        this.projectId = $.projectId;
        this.providerSettingsBackingProviderName = $.providerSettingsBackingProviderName;
        this.providerSettingsProviderName = $.providerSettingsProviderName;
        this.providerSettingsRegionName = $.providerSettingsRegionName;
        this.stateName = $.stateName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ServerlessInstanceState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ServerlessInstanceState $;

        public Builder() {
            $ = new ServerlessInstanceState();
        }

        public Builder(ServerlessInstanceState defaults) {
            $ = new ServerlessInstanceState(Objects.requireNonNull(defaults));
        }

        /**
         * @param connectionStringsStandardSrv Public `mongodb+srv://` connection string that you can use to connect to this serverless instance.
         * 
         * @return builder
         * 
         */
        public Builder connectionStringsStandardSrv(@Nullable Output<String> connectionStringsStandardSrv) {
            $.connectionStringsStandardSrv = connectionStringsStandardSrv;
            return this;
        }

        /**
         * @param connectionStringsStandardSrv Public `mongodb+srv://` connection string that you can use to connect to this serverless instance.
         * 
         * @return builder
         * 
         */
        public Builder connectionStringsStandardSrv(String connectionStringsStandardSrv) {
            return connectionStringsStandardSrv(Output.of(connectionStringsStandardSrv));
        }

        /**
         * @param createDate Timestamp that indicates when MongoDB Cloud created the serverless instance. The timestamp displays in the ISO 8601 date and time format in UTC.
         * 
         * @return builder
         * 
         */
        public Builder createDate(@Nullable Output<String> createDate) {
            $.createDate = createDate;
            return this;
        }

        /**
         * @param createDate Timestamp that indicates when MongoDB Cloud created the serverless instance. The timestamp displays in the ISO 8601 date and time format in UTC.
         * 
         * @return builder
         * 
         */
        public Builder createDate(String createDate) {
            return createDate(Output.of(createDate));
        }

        public Builder links(@Nullable Output<List<ServerlessInstanceLinkArgs>> links) {
            $.links = links;
            return this;
        }

        public Builder links(List<ServerlessInstanceLinkArgs> links) {
            return links(Output.of(links));
        }

        public Builder links(ServerlessInstanceLinkArgs... links) {
            return links(List.of(links));
        }

        /**
         * @param mongoDbVersion Version of MongoDB that the serverless instance runs, in `&lt;major version&gt;`.`&lt;minor version&gt;` format.
         * 
         * @return builder
         * 
         */
        public Builder mongoDbVersion(@Nullable Output<String> mongoDbVersion) {
            $.mongoDbVersion = mongoDbVersion;
            return this;
        }

        /**
         * @param mongoDbVersion Version of MongoDB that the serverless instance runs, in `&lt;major version&gt;`.`&lt;minor version&gt;` format.
         * 
         * @return builder
         * 
         */
        public Builder mongoDbVersion(String mongoDbVersion) {
            return mongoDbVersion(Output.of(mongoDbVersion));
        }

        /**
         * @param name Human-readable label that identifies the serverless instance.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Human-readable label that identifies the serverless instance.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param projectId The ID of the organization or project you want to create the serverless instance within.
         * 
         * @return builder
         * 
         */
        public Builder projectId(@Nullable Output<String> projectId) {
            $.projectId = projectId;
            return this;
        }

        /**
         * @param projectId The ID of the organization or project you want to create the serverless instance within.
         * 
         * @return builder
         * 
         */
        public Builder projectId(String projectId) {
            return projectId(Output.of(projectId));
        }

        /**
         * @param providerSettingsBackingProviderName Cloud service provider on which MongoDB Cloud provisioned the serverless instance.
         * 
         * @return builder
         * 
         */
        public Builder providerSettingsBackingProviderName(@Nullable Output<String> providerSettingsBackingProviderName) {
            $.providerSettingsBackingProviderName = providerSettingsBackingProviderName;
            return this;
        }

        /**
         * @param providerSettingsBackingProviderName Cloud service provider on which MongoDB Cloud provisioned the serverless instance.
         * 
         * @return builder
         * 
         */
        public Builder providerSettingsBackingProviderName(String providerSettingsBackingProviderName) {
            return providerSettingsBackingProviderName(Output.of(providerSettingsBackingProviderName));
        }

        /**
         * @param providerSettingsProviderName Cloud service provider that applies to the provisioned the serverless instance.
         * 
         * @return builder
         * 
         */
        public Builder providerSettingsProviderName(@Nullable Output<String> providerSettingsProviderName) {
            $.providerSettingsProviderName = providerSettingsProviderName;
            return this;
        }

        /**
         * @param providerSettingsProviderName Cloud service provider that applies to the provisioned the serverless instance.
         * 
         * @return builder
         * 
         */
        public Builder providerSettingsProviderName(String providerSettingsProviderName) {
            return providerSettingsProviderName(Output.of(providerSettingsProviderName));
        }

        /**
         * @param providerSettingsRegionName Human-readable label that identifies the physical location of your MongoDB serverless instance. The region you choose can affect network latency for clients accessing your databases.
         * 
         * @return builder
         * 
         */
        public Builder providerSettingsRegionName(@Nullable Output<String> providerSettingsRegionName) {
            $.providerSettingsRegionName = providerSettingsRegionName;
            return this;
        }

        /**
         * @param providerSettingsRegionName Human-readable label that identifies the physical location of your MongoDB serverless instance. The region you choose can affect network latency for clients accessing your databases.
         * 
         * @return builder
         * 
         */
        public Builder providerSettingsRegionName(String providerSettingsRegionName) {
            return providerSettingsRegionName(Output.of(providerSettingsRegionName));
        }

        /**
         * @param stateName Stage of deployment of this serverless instance when the resource made its request.
         * 
         * @return builder
         * 
         */
        public Builder stateName(@Nullable Output<String> stateName) {
            $.stateName = stateName;
            return this;
        }

        /**
         * @param stateName Stage of deployment of this serverless instance when the resource made its request.
         * 
         * @return builder
         * 
         */
        public Builder stateName(String stateName) {
            return stateName(Output.of(stateName));
        }

        public ServerlessInstanceState build() {
            return $;
        }
    }

}
