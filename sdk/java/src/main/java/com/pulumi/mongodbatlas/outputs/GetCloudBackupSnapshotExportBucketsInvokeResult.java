// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.mongodbatlas.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.mongodbatlas.outputs.GetCloudBackupSnapshotExportBucketsResult;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetCloudBackupSnapshotExportBucketsInvokeResult {
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    private @Nullable Integer itemsPerPage;
    private @Nullable Integer pageNum;
    /**
     * @return The unique identifier of the project for the Atlas cluster.
     * 
     */
    private String projectId;
    /**
     * @return Includes CloudProviderSnapshotExportBucket object for each item detailed in the results array section.
     * 
     */
    private List<GetCloudBackupSnapshotExportBucketsResult> results;
    private Integer totalCount;

    private GetCloudBackupSnapshotExportBucketsInvokeResult() {}
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    public Optional<Integer> itemsPerPage() {
        return Optional.ofNullable(this.itemsPerPage);
    }
    public Optional<Integer> pageNum() {
        return Optional.ofNullable(this.pageNum);
    }
    /**
     * @return The unique identifier of the project for the Atlas cluster.
     * 
     */
    public String projectId() {
        return this.projectId;
    }
    /**
     * @return Includes CloudProviderSnapshotExportBucket object for each item detailed in the results array section.
     * 
     */
    public List<GetCloudBackupSnapshotExportBucketsResult> results() {
        return this.results;
    }
    public Integer totalCount() {
        return this.totalCount;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetCloudBackupSnapshotExportBucketsInvokeResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String id;
        private @Nullable Integer itemsPerPage;
        private @Nullable Integer pageNum;
        private String projectId;
        private List<GetCloudBackupSnapshotExportBucketsResult> results;
        private Integer totalCount;
        public Builder() {}
        public Builder(GetCloudBackupSnapshotExportBucketsInvokeResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.id = defaults.id;
    	      this.itemsPerPage = defaults.itemsPerPage;
    	      this.pageNum = defaults.pageNum;
    	      this.projectId = defaults.projectId;
    	      this.results = defaults.results;
    	      this.totalCount = defaults.totalCount;
        }

        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder itemsPerPage(@Nullable Integer itemsPerPage) {
            this.itemsPerPage = itemsPerPage;
            return this;
        }
        @CustomType.Setter
        public Builder pageNum(@Nullable Integer pageNum) {
            this.pageNum = pageNum;
            return this;
        }
        @CustomType.Setter
        public Builder projectId(String projectId) {
            this.projectId = Objects.requireNonNull(projectId);
            return this;
        }
        @CustomType.Setter
        public Builder results(List<GetCloudBackupSnapshotExportBucketsResult> results) {
            this.results = Objects.requireNonNull(results);
            return this;
        }
        public Builder results(GetCloudBackupSnapshotExportBucketsResult... results) {
            return results(List.of(results));
        }
        @CustomType.Setter
        public Builder totalCount(Integer totalCount) {
            this.totalCount = Objects.requireNonNull(totalCount);
            return this;
        }
        public GetCloudBackupSnapshotExportBucketsInvokeResult build() {
            final var o = new GetCloudBackupSnapshotExportBucketsInvokeResult();
            o.id = id;
            o.itemsPerPage = itemsPerPage;
            o.pageNum = pageNum;
            o.projectId = projectId;
            o.results = results;
            o.totalCount = totalCount;
            return o;
        }
    }
}
