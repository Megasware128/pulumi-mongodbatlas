// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.mongodbatlas.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class GetAlertConfigurationPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetAlertConfigurationPlainArgs Empty = new GetAlertConfigurationPlainArgs();

    /**
     * Unique identifier for the alert configuration.
     * 
     */
    @Import(name="alertConfigurationId", required=true)
    private String alertConfigurationId;

    /**
     * @return Unique identifier for the alert configuration.
     * 
     */
    public String alertConfigurationId() {
        return this.alertConfigurationId;
    }

    /**
     * The ID of the project where the alert configuration will create.
     * 
     */
    @Import(name="projectId", required=true)
    private String projectId;

    /**
     * @return The ID of the project where the alert configuration will create.
     * 
     */
    public String projectId() {
        return this.projectId;
    }

    private GetAlertConfigurationPlainArgs() {}

    private GetAlertConfigurationPlainArgs(GetAlertConfigurationPlainArgs $) {
        this.alertConfigurationId = $.alertConfigurationId;
        this.projectId = $.projectId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetAlertConfigurationPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetAlertConfigurationPlainArgs $;

        public Builder() {
            $ = new GetAlertConfigurationPlainArgs();
        }

        public Builder(GetAlertConfigurationPlainArgs defaults) {
            $ = new GetAlertConfigurationPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param alertConfigurationId Unique identifier for the alert configuration.
         * 
         * @return builder
         * 
         */
        public Builder alertConfigurationId(String alertConfigurationId) {
            $.alertConfigurationId = alertConfigurationId;
            return this;
        }

        /**
         * @param projectId The ID of the project where the alert configuration will create.
         * 
         * @return builder
         * 
         */
        public Builder projectId(String projectId) {
            $.projectId = projectId;
            return this;
        }

        public GetAlertConfigurationPlainArgs build() {
            $.alertConfigurationId = Objects.requireNonNull($.alertConfigurationId, "expected parameter 'alertConfigurationId' to be non-null");
            $.projectId = Objects.requireNonNull($.projectId, "expected parameter 'projectId' to be non-null");
            return $;
        }
    }

}