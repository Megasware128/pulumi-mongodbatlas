// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.mongodbatlas.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.mongodbatlas.outputs.GetSearchIndexesResult;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetSearchIndexesInvokeResult {
    private String clusterName;
    /**
     * @return (Required) Name of the collection the index is on.
     * 
     */
    private String collectionName;
    /**
     * @return (Required) Name of the database the collection is in.
     * 
     */
    private String database;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    private @Nullable Integer itemsPerPage;
    private @Nullable Integer pageNum;
    private String projectId;
    /**
     * @return A list where each represents a search index.
     * 
     */
    private List<GetSearchIndexesResult> results;
    /**
     * @return Represents the total of the search indexes
     * 
     */
    private Integer totalCount;

    private GetSearchIndexesInvokeResult() {}
    public String clusterName() {
        return this.clusterName;
    }
    /**
     * @return (Required) Name of the collection the index is on.
     * 
     */
    public String collectionName() {
        return this.collectionName;
    }
    /**
     * @return (Required) Name of the database the collection is in.
     * 
     */
    public String database() {
        return this.database;
    }
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
    public String projectId() {
        return this.projectId;
    }
    /**
     * @return A list where each represents a search index.
     * 
     */
    public List<GetSearchIndexesResult> results() {
        return this.results;
    }
    /**
     * @return Represents the total of the search indexes
     * 
     */
    public Integer totalCount() {
        return this.totalCount;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetSearchIndexesInvokeResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String clusterName;
        private String collectionName;
        private String database;
        private String id;
        private @Nullable Integer itemsPerPage;
        private @Nullable Integer pageNum;
        private String projectId;
        private List<GetSearchIndexesResult> results;
        private Integer totalCount;
        public Builder() {}
        public Builder(GetSearchIndexesInvokeResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.clusterName = defaults.clusterName;
    	      this.collectionName = defaults.collectionName;
    	      this.database = defaults.database;
    	      this.id = defaults.id;
    	      this.itemsPerPage = defaults.itemsPerPage;
    	      this.pageNum = defaults.pageNum;
    	      this.projectId = defaults.projectId;
    	      this.results = defaults.results;
    	      this.totalCount = defaults.totalCount;
        }

        @CustomType.Setter
        public Builder clusterName(String clusterName) {
            this.clusterName = Objects.requireNonNull(clusterName);
            return this;
        }
        @CustomType.Setter
        public Builder collectionName(String collectionName) {
            this.collectionName = Objects.requireNonNull(collectionName);
            return this;
        }
        @CustomType.Setter
        public Builder database(String database) {
            this.database = Objects.requireNonNull(database);
            return this;
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
        public Builder results(List<GetSearchIndexesResult> results) {
            this.results = Objects.requireNonNull(results);
            return this;
        }
        public Builder results(GetSearchIndexesResult... results) {
            return results(List.of(results));
        }
        @CustomType.Setter
        public Builder totalCount(Integer totalCount) {
            this.totalCount = Objects.requireNonNull(totalCount);
            return this;
        }
        public GetSearchIndexesInvokeResult build() {
            final var o = new GetSearchIndexesInvokeResult();
            o.clusterName = clusterName;
            o.collectionName = collectionName;
            o.database = database;
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
