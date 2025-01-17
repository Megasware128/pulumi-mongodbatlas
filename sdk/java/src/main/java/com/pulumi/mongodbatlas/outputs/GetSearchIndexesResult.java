// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.mongodbatlas.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.mongodbatlas.outputs.GetSearchIndexesResultSynonym;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetSearchIndexesResult {
    /**
     * @return [Analyzer](https://docs.atlas.mongodb.com/reference/atlas-search/analyzers/#std-label-analyzers-ref) to use when creating the index.
     * 
     */
    private String analyzer;
    /**
     * @return [Custom analyzers](https://docs.atlas.mongodb.com/reference/atlas-search/analyzers/custom/#std-label-custom-analyzers) to use in this index (this is an array of objects).
     * 
     */
    private @Nullable String analyzers;
    /**
     * @return Name of the cluster containing the collection with one or more Atlas Search indexes.
     * 
     */
    private String clusterName;
    /**
     * @return Name of the collection with one or more Atlas Search indexes.
     * 
     */
    private String collectionName;
    /**
     * @return (Required) Name of the database the collection is in.
     * 
     */
    private String database;
    private String indexId;
    /**
     * @return Flag indicating whether the index uses dynamic or static mappings.
     * 
     */
    private @Nullable Boolean mappingsDynamic;
    /**
     * @return Object containing one or more field specifications.
     * 
     */
    private @Nullable String mappingsFields;
    /**
     * @return Name of the index.
     * 
     */
    private String name;
    /**
     * @return Unique identifier for the [project](https://docs.atlas.mongodb.com/organizations-projects/#std-label-projects) that contains the specified cluster.
     * 
     */
    private String projectId;
    /**
     * @return [Analyzer](https://docs.atlas.mongodb.com/reference/atlas-search/analyzers/#std-label-analyzers-ref) to use when searching the index.
     * 
     */
    private @Nullable String searchAnalyzer;
    private String status;
    /**
     * @return Synonyms mapping definition to use in this index.
     * * `synonyms.#.name` - Name of the [synonym mapping definition](https://docs.atlas.mongodb.com/reference/atlas-search/synonyms/#std-label-synonyms-ref).
     * * `synonyms.#.source_collection` - Name of the source MongoDB collection for the synonyms.
     * * `synonyms.#.analyzer` - Name of the [analyzer](https://docs.atlas.mongodb.com/reference/atlas-search/analyzers/#std-label-analyzers-ref) to use with this synonym mapping.
     * 
     */
    private @Nullable List<GetSearchIndexesResultSynonym> synonyms;

    private GetSearchIndexesResult() {}
    /**
     * @return [Analyzer](https://docs.atlas.mongodb.com/reference/atlas-search/analyzers/#std-label-analyzers-ref) to use when creating the index.
     * 
     */
    public String analyzer() {
        return this.analyzer;
    }
    /**
     * @return [Custom analyzers](https://docs.atlas.mongodb.com/reference/atlas-search/analyzers/custom/#std-label-custom-analyzers) to use in this index (this is an array of objects).
     * 
     */
    public Optional<String> analyzers() {
        return Optional.ofNullable(this.analyzers);
    }
    /**
     * @return Name of the cluster containing the collection with one or more Atlas Search indexes.
     * 
     */
    public String clusterName() {
        return this.clusterName;
    }
    /**
     * @return Name of the collection with one or more Atlas Search indexes.
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
    public String indexId() {
        return this.indexId;
    }
    /**
     * @return Flag indicating whether the index uses dynamic or static mappings.
     * 
     */
    public Optional<Boolean> mappingsDynamic() {
        return Optional.ofNullable(this.mappingsDynamic);
    }
    /**
     * @return Object containing one or more field specifications.
     * 
     */
    public Optional<String> mappingsFields() {
        return Optional.ofNullable(this.mappingsFields);
    }
    /**
     * @return Name of the index.
     * 
     */
    public String name() {
        return this.name;
    }
    /**
     * @return Unique identifier for the [project](https://docs.atlas.mongodb.com/organizations-projects/#std-label-projects) that contains the specified cluster.
     * 
     */
    public String projectId() {
        return this.projectId;
    }
    /**
     * @return [Analyzer](https://docs.atlas.mongodb.com/reference/atlas-search/analyzers/#std-label-analyzers-ref) to use when searching the index.
     * 
     */
    public Optional<String> searchAnalyzer() {
        return Optional.ofNullable(this.searchAnalyzer);
    }
    public String status() {
        return this.status;
    }
    /**
     * @return Synonyms mapping definition to use in this index.
     * * `synonyms.#.name` - Name of the [synonym mapping definition](https://docs.atlas.mongodb.com/reference/atlas-search/synonyms/#std-label-synonyms-ref).
     * * `synonyms.#.source_collection` - Name of the source MongoDB collection for the synonyms.
     * * `synonyms.#.analyzer` - Name of the [analyzer](https://docs.atlas.mongodb.com/reference/atlas-search/analyzers/#std-label-analyzers-ref) to use with this synonym mapping.
     * 
     */
    public List<GetSearchIndexesResultSynonym> synonyms() {
        return this.synonyms == null ? List.of() : this.synonyms;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetSearchIndexesResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String analyzer;
        private @Nullable String analyzers;
        private String clusterName;
        private String collectionName;
        private String database;
        private String indexId;
        private @Nullable Boolean mappingsDynamic;
        private @Nullable String mappingsFields;
        private String name;
        private String projectId;
        private @Nullable String searchAnalyzer;
        private String status;
        private @Nullable List<GetSearchIndexesResultSynonym> synonyms;
        public Builder() {}
        public Builder(GetSearchIndexesResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.analyzer = defaults.analyzer;
    	      this.analyzers = defaults.analyzers;
    	      this.clusterName = defaults.clusterName;
    	      this.collectionName = defaults.collectionName;
    	      this.database = defaults.database;
    	      this.indexId = defaults.indexId;
    	      this.mappingsDynamic = defaults.mappingsDynamic;
    	      this.mappingsFields = defaults.mappingsFields;
    	      this.name = defaults.name;
    	      this.projectId = defaults.projectId;
    	      this.searchAnalyzer = defaults.searchAnalyzer;
    	      this.status = defaults.status;
    	      this.synonyms = defaults.synonyms;
        }

        @CustomType.Setter
        public Builder analyzer(String analyzer) {
            this.analyzer = Objects.requireNonNull(analyzer);
            return this;
        }
        @CustomType.Setter
        public Builder analyzers(@Nullable String analyzers) {
            this.analyzers = analyzers;
            return this;
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
        public Builder indexId(String indexId) {
            this.indexId = Objects.requireNonNull(indexId);
            return this;
        }
        @CustomType.Setter
        public Builder mappingsDynamic(@Nullable Boolean mappingsDynamic) {
            this.mappingsDynamic = mappingsDynamic;
            return this;
        }
        @CustomType.Setter
        public Builder mappingsFields(@Nullable String mappingsFields) {
            this.mappingsFields = mappingsFields;
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            this.name = Objects.requireNonNull(name);
            return this;
        }
        @CustomType.Setter
        public Builder projectId(String projectId) {
            this.projectId = Objects.requireNonNull(projectId);
            return this;
        }
        @CustomType.Setter
        public Builder searchAnalyzer(@Nullable String searchAnalyzer) {
            this.searchAnalyzer = searchAnalyzer;
            return this;
        }
        @CustomType.Setter
        public Builder status(String status) {
            this.status = Objects.requireNonNull(status);
            return this;
        }
        @CustomType.Setter
        public Builder synonyms(@Nullable List<GetSearchIndexesResultSynonym> synonyms) {
            this.synonyms = synonyms;
            return this;
        }
        public Builder synonyms(GetSearchIndexesResultSynonym... synonyms) {
            return synonyms(List.of(synonyms));
        }
        public GetSearchIndexesResult build() {
            final var o = new GetSearchIndexesResult();
            o.analyzer = analyzer;
            o.analyzers = analyzers;
            o.clusterName = clusterName;
            o.collectionName = collectionName;
            o.database = database;
            o.indexId = indexId;
            o.mappingsDynamic = mappingsDynamic;
            o.mappingsFields = mappingsFields;
            o.name = name;
            o.projectId = projectId;
            o.searchAnalyzer = searchAnalyzer;
            o.status = status;
            o.synonyms = synonyms;
            return o;
        }
    }
}
