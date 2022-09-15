// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.mongodbatlas.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetCloudBackupSnapshotsArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetCloudBackupSnapshotsArgs Empty = new GetCloudBackupSnapshotsArgs();

    /**
     * The name of the Atlas cluster that contains the snapshot you want to retrieve.
     * 
     */
    @Import(name="clusterName", required=true)
    private Output<String> clusterName;

    /**
     * @return The name of the Atlas cluster that contains the snapshot you want to retrieve.
     * 
     */
    public Output<String> clusterName() {
        return this.clusterName;
    }

    /**
     * Number of items to return per page, up to a maximum of 500. Defaults to `100`.
     * 
     */
    @Import(name="itemsPerPage")
    private @Nullable Output<Integer> itemsPerPage;

    /**
     * @return Number of items to return per page, up to a maximum of 500. Defaults to `100`.
     * 
     */
    public Optional<Output<Integer>> itemsPerPage() {
        return Optional.ofNullable(this.itemsPerPage);
    }

    /**
     * The page to return. Defaults to `1`.
     * 
     */
    @Import(name="pageNum")
    private @Nullable Output<Integer> pageNum;

    /**
     * @return The page to return. Defaults to `1`.
     * 
     */
    public Optional<Output<Integer>> pageNum() {
        return Optional.ofNullable(this.pageNum);
    }

    @Import(name="projectId", required=true)
    private Output<String> projectId;

    public Output<String> projectId() {
        return this.projectId;
    }

    private GetCloudBackupSnapshotsArgs() {}

    private GetCloudBackupSnapshotsArgs(GetCloudBackupSnapshotsArgs $) {
        this.clusterName = $.clusterName;
        this.itemsPerPage = $.itemsPerPage;
        this.pageNum = $.pageNum;
        this.projectId = $.projectId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetCloudBackupSnapshotsArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetCloudBackupSnapshotsArgs $;

        public Builder() {
            $ = new GetCloudBackupSnapshotsArgs();
        }

        public Builder(GetCloudBackupSnapshotsArgs defaults) {
            $ = new GetCloudBackupSnapshotsArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param clusterName The name of the Atlas cluster that contains the snapshot you want to retrieve.
         * 
         * @return builder
         * 
         */
        public Builder clusterName(Output<String> clusterName) {
            $.clusterName = clusterName;
            return this;
        }

        /**
         * @param clusterName The name of the Atlas cluster that contains the snapshot you want to retrieve.
         * 
         * @return builder
         * 
         */
        public Builder clusterName(String clusterName) {
            return clusterName(Output.of(clusterName));
        }

        /**
         * @param itemsPerPage Number of items to return per page, up to a maximum of 500. Defaults to `100`.
         * 
         * @return builder
         * 
         */
        public Builder itemsPerPage(@Nullable Output<Integer> itemsPerPage) {
            $.itemsPerPage = itemsPerPage;
            return this;
        }

        /**
         * @param itemsPerPage Number of items to return per page, up to a maximum of 500. Defaults to `100`.
         * 
         * @return builder
         * 
         */
        public Builder itemsPerPage(Integer itemsPerPage) {
            return itemsPerPage(Output.of(itemsPerPage));
        }

        /**
         * @param pageNum The page to return. Defaults to `1`.
         * 
         * @return builder
         * 
         */
        public Builder pageNum(@Nullable Output<Integer> pageNum) {
            $.pageNum = pageNum;
            return this;
        }

        /**
         * @param pageNum The page to return. Defaults to `1`.
         * 
         * @return builder
         * 
         */
        public Builder pageNum(Integer pageNum) {
            return pageNum(Output.of(pageNum));
        }

        public Builder projectId(Output<String> projectId) {
            $.projectId = projectId;
            return this;
        }

        public Builder projectId(String projectId) {
            return projectId(Output.of(projectId));
        }

        public GetCloudBackupSnapshotsArgs build() {
            $.clusterName = Objects.requireNonNull($.clusterName, "expected parameter 'clusterName' to be non-null");
            $.projectId = Objects.requireNonNull($.projectId, "expected parameter 'projectId' to be non-null");
            return $;
        }
    }

}