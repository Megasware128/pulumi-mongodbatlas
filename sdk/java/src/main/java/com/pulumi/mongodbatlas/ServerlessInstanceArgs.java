// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.mongodbatlas;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.mongodbatlas.inputs.ServerlessInstanceLinkArgs;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ServerlessInstanceArgs extends com.pulumi.resources.ResourceArgs {

    public static final ServerlessInstanceArgs Empty = new ServerlessInstanceArgs();

    @Import(name="links")
    private @Nullable Output<List<ServerlessInstanceLinkArgs>> links;

    public Optional<Output<List<ServerlessInstanceLinkArgs>>> links() {
        return Optional.ofNullable(this.links);
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
    @Import(name="projectId", required=true)
    private Output<String> projectId;

    /**
     * @return The ID of the organization or project you want to create the serverless instance within.
     * 
     */
    public Output<String> projectId() {
        return this.projectId;
    }

    /**
     * Cloud service provider on which MongoDB Cloud provisioned the serverless instance.
     * 
     */
    @Import(name="providerSettingsBackingProviderName", required=true)
    private Output<String> providerSettingsBackingProviderName;

    /**
     * @return Cloud service provider on which MongoDB Cloud provisioned the serverless instance.
     * 
     */
    public Output<String> providerSettingsBackingProviderName() {
        return this.providerSettingsBackingProviderName;
    }

    /**
     * Cloud service provider that applies to the provisioned the serverless instance.
     * 
     */
    @Import(name="providerSettingsProviderName", required=true)
    private Output<String> providerSettingsProviderName;

    /**
     * @return Cloud service provider that applies to the provisioned the serverless instance.
     * 
     */
    public Output<String> providerSettingsProviderName() {
        return this.providerSettingsProviderName;
    }

    /**
     * Human-readable label that identifies the physical location of your MongoDB serverless instance. The region you choose can affect network latency for clients accessing your databases.
     * 
     */
    @Import(name="providerSettingsRegionName", required=true)
    private Output<String> providerSettingsRegionName;

    /**
     * @return Human-readable label that identifies the physical location of your MongoDB serverless instance. The region you choose can affect network latency for clients accessing your databases.
     * 
     */
    public Output<String> providerSettingsRegionName() {
        return this.providerSettingsRegionName;
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

    private ServerlessInstanceArgs() {}

    private ServerlessInstanceArgs(ServerlessInstanceArgs $) {
        this.links = $.links;
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
    public static Builder builder(ServerlessInstanceArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ServerlessInstanceArgs $;

        public Builder() {
            $ = new ServerlessInstanceArgs();
        }

        public Builder(ServerlessInstanceArgs defaults) {
            $ = new ServerlessInstanceArgs(Objects.requireNonNull(defaults));
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
        public Builder projectId(Output<String> projectId) {
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
        public Builder providerSettingsBackingProviderName(Output<String> providerSettingsBackingProviderName) {
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
        public Builder providerSettingsProviderName(Output<String> providerSettingsProviderName) {
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
        public Builder providerSettingsRegionName(Output<String> providerSettingsRegionName) {
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

        public ServerlessInstanceArgs build() {
            $.projectId = Objects.requireNonNull($.projectId, "expected parameter 'projectId' to be non-null");
            $.providerSettingsBackingProviderName = Objects.requireNonNull($.providerSettingsBackingProviderName, "expected parameter 'providerSettingsBackingProviderName' to be non-null");
            $.providerSettingsProviderName = Objects.requireNonNull($.providerSettingsProviderName, "expected parameter 'providerSettingsProviderName' to be non-null");
            $.providerSettingsRegionName = Objects.requireNonNull($.providerSettingsRegionName, "expected parameter 'providerSettingsRegionName' to be non-null");
            return $;
        }
    }

}
