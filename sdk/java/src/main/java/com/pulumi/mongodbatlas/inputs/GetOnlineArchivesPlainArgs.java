// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.mongodbatlas.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class GetOnlineArchivesPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetOnlineArchivesPlainArgs Empty = new GetOnlineArchivesPlainArgs();

    /**
     * Name of the cluster that contains the collection.
     * 
     */
    @Import(name="clusterName", required=true)
    private String clusterName;

    /**
     * @return Name of the cluster that contains the collection.
     * 
     */
    public String clusterName() {
        return this.clusterName;
    }

    /**
     * The unique ID for the project.
     * 
     */
    @Import(name="projectId", required=true)
    private String projectId;

    /**
     * @return The unique ID for the project.
     * 
     */
    public String projectId() {
        return this.projectId;
    }

    private GetOnlineArchivesPlainArgs() {}

    private GetOnlineArchivesPlainArgs(GetOnlineArchivesPlainArgs $) {
        this.clusterName = $.clusterName;
        this.projectId = $.projectId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetOnlineArchivesPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetOnlineArchivesPlainArgs $;

        public Builder() {
            $ = new GetOnlineArchivesPlainArgs();
        }

        public Builder(GetOnlineArchivesPlainArgs defaults) {
            $ = new GetOnlineArchivesPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param clusterName Name of the cluster that contains the collection.
         * 
         * @return builder
         * 
         */
        public Builder clusterName(String clusterName) {
            $.clusterName = clusterName;
            return this;
        }

        /**
         * @param projectId The unique ID for the project.
         * 
         * @return builder
         * 
         */
        public Builder projectId(String projectId) {
            $.projectId = projectId;
            return this;
        }

        public GetOnlineArchivesPlainArgs build() {
            $.clusterName = Objects.requireNonNull($.clusterName, "expected parameter 'clusterName' to be non-null");
            $.projectId = Objects.requireNonNull($.projectId, "expected parameter 'projectId' to be non-null");
            return $;
        }
    }

}