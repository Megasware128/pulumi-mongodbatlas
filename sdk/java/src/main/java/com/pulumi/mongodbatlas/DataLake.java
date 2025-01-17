// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.mongodbatlas;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.mongodbatlas.DataLakeArgs;
import com.pulumi.mongodbatlas.Utilities;
import com.pulumi.mongodbatlas.inputs.DataLakeState;
import com.pulumi.mongodbatlas.outputs.DataLakeAws;
import com.pulumi.mongodbatlas.outputs.DataLakeDataProcessRegion;
import com.pulumi.mongodbatlas.outputs.DataLakeStorageDatabase;
import com.pulumi.mongodbatlas.outputs.DataLakeStorageStore;
import java.lang.String;
import java.util.List;
import javax.annotation.Nullable;

/**
 * ## Example Usage
 * ### S
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.mongodbatlas.Project;
 * import com.pulumi.mongodbatlas.ProjectArgs;
 * import com.pulumi.mongodbatlas.CloudProviderAccess;
 * import com.pulumi.mongodbatlas.CloudProviderAccessArgs;
 * import com.pulumi.mongodbatlas.DataLake;
 * import com.pulumi.mongodbatlas.DataLakeArgs;
 * import com.pulumi.mongodbatlas.inputs.DataLakeAwsArgs;
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
 *         var testProject = new Project(&#34;testProject&#34;, ProjectArgs.builder()        
 *             .orgId(&#34;ORGANIZATION ID&#34;)
 *             .build());
 * 
 *         var testCloudProviderAccess = new CloudProviderAccess(&#34;testCloudProviderAccess&#34;, CloudProviderAccessArgs.builder()        
 *             .projectId(testProject.id())
 *             .providerName(&#34;AWS&#34;)
 *             .iamAssumedRoleArn(&#34;AWS ROLE ID&#34;)
 *             .build());
 * 
 *         var basicDs = new DataLake(&#34;basicDs&#34;, DataLakeArgs.builder()        
 *             .projectId(testProject.id())
 *             .aws(DataLakeAwsArgs.builder()
 *                 .roleId(testCloudProviderAccess.roleId())
 *                 .testS3Bucket(&#34;TEST S3 BUCKET NAME&#34;)
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Data Lake can be imported using project ID, name of the data lake and name of the AWS s3 bucket, in the format `project_id`--`name`--`aws_test_s3_bucket`, e.g.
 * 
 * ```sh
 *  $ pulumi import mongodbatlas:index/dataLake:DataLake example 1112222b3bf99403840e8934--test-data-lake--s3-test
 * ```
 * 
 *  See [MongoDB Atlas API](https://docs.mongodb.com/datalake/reference/api/dataLakes-create-one-tenant) Documentation for more information.
 * 
 */
@ResourceType(type="mongodbatlas:index/dataLake:DataLake")
public class DataLake extends com.pulumi.resources.CustomResource {
    /**
     * AWS provider of the cloud service where Data Lake can access the S3 Bucket.
     * * `aws.0.role_id` - (Required) Unique identifier of the role that Data Lake can use to access the data stores. If necessary, use the Atlas [UI](https://docs.atlas.mongodb.com/security/manage-iam-roles/) or [API](https://docs.atlas.mongodb.com/reference/api/cloud-provider-access-get-roles/) to retrieve the role ID. You must also specify the `aws.0.test_s3_bucket`.
     * * `aws.0.test_s3_bucket` - (Required) Name of the S3 data bucket that the provided role ID is authorized to access. You must also specify the `aws.0.role_id`.
     * 
     */
    @Export(name="aws", type=DataLakeAws.class, parameters={})
    private Output<DataLakeAws> aws;

    /**
     * @return AWS provider of the cloud service where Data Lake can access the S3 Bucket.
     * * `aws.0.role_id` - (Required) Unique identifier of the role that Data Lake can use to access the data stores. If necessary, use the Atlas [UI](https://docs.atlas.mongodb.com/security/manage-iam-roles/) or [API](https://docs.atlas.mongodb.com/reference/api/cloud-provider-access-get-roles/) to retrieve the role ID. You must also specify the `aws.0.test_s3_bucket`.
     * * `aws.0.test_s3_bucket` - (Required) Name of the S3 data bucket that the provided role ID is authorized to access. You must also specify the `aws.0.role_id`.
     * 
     */
    public Output<DataLakeAws> aws() {
        return this.aws;
    }
    /**
     * The cloud provider region to which Atlas Data Lake routes client connections for data processing. Set to `null` to direct Atlas Data Lake to route client connections to the region nearest to the client based on DNS resolution.
     * * `data_process_region.0.cloud_provider` - (Required) Name of the cloud service provider. Atlas Data Lake only supports AWS.
     * * `data_process_region.0.region` - (Required). Name of the region to which Data Lake routes client connections for data processing. Atlas Data Lake only supports the following regions:
     * 
     */
    @Export(name="dataProcessRegion", type=DataLakeDataProcessRegion.class, parameters={})
    private Output<DataLakeDataProcessRegion> dataProcessRegion;

    /**
     * @return The cloud provider region to which Atlas Data Lake routes client connections for data processing. Set to `null` to direct Atlas Data Lake to route client connections to the region nearest to the client based on DNS resolution.
     * * `data_process_region.0.cloud_provider` - (Required) Name of the cloud service provider. Atlas Data Lake only supports AWS.
     * * `data_process_region.0.region` - (Required). Name of the region to which Data Lake routes client connections for data processing. Atlas Data Lake only supports the following regions:
     * 
     */
    public Output<DataLakeDataProcessRegion> dataProcessRegion() {
        return this.dataProcessRegion;
    }
    /**
     * The list of hostnames assigned to the Atlas Data Lake. Each string in the array is a hostname assigned to the Atlas Data Lake.
     * 
     */
    @Export(name="hostnames", type=List.class, parameters={String.class})
    private Output<List<String>> hostnames;

    /**
     * @return The list of hostnames assigned to the Atlas Data Lake. Each string in the array is a hostname assigned to the Atlas Data Lake.
     * 
     */
    public Output<List<String>> hostnames() {
        return this.hostnames;
    }
    /**
     * Name of the Atlas Data Lake.
     * 
     */
    @Export(name="name", type=String.class, parameters={})
    private Output<String> name;

    /**
     * @return Name of the Atlas Data Lake.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The unique ID for the project to create a data lake.
     * 
     */
    @Export(name="projectId", type=String.class, parameters={})
    private Output<String> projectId;

    /**
     * @return The unique ID for the project to create a data lake.
     * 
     */
    public Output<String> projectId() {
        return this.projectId;
    }
    /**
     * Current state of the Atlas Data Lake:
     * 
     */
    @Export(name="state", type=String.class, parameters={})
    private Output<String> state;

    /**
     * @return Current state of the Atlas Data Lake:
     * 
     */
    public Output<String> state() {
        return this.state;
    }
    /**
     * Configuration details for mapping each data store to queryable databases and collections. For complete documentation on this object and its nested fields, see [databases](https://docs.mongodb.com/datalake/reference/format/data-lake-configuration#std-label-datalake-databases-reference). An empty object indicates that the Data Lake has no mapping configuration for any data store.
     * * `storage_databases.#.name` - Name of the database to which Data Lake maps the data contained in the data store.
     * * `storage_databases.#.collections` -     Array of objects where each object represents a collection and data sources that map to a [stores](https://docs.mongodb.com/datalake/reference/format/data-lake-configuration#mongodb-datalakeconf-datalakeconf.stores) data store.
     * * `storage_databases.#.collections.#.name` - Name of the collection.
     * * `storage_databases.#.collections.#.data_sources` -     Array of objects where each object represents a stores data store to map with the collection.
     * * `storage_databases.#.collections.#.data_sources.#.store_name` -     Name of a data store to map to the `&lt;collection&gt;`. Must match the name of an object in the stores array.
     * * `storage_databases.#.collections.#.data_sources.#.default_format` - Default format that Data Lake assumes if it encounters a file without an extension while searching the storeName.
     * * `storage_databases.#.collections.#.data_sources.#.path` - Controls how Atlas Data Lake searches for and parses files in the storeName before mapping them to the `&lt;collection&gt;`.
     * * `storage_databases.#.views` -     Array of objects where each object represents an [aggregation pipeline](https://docs.mongodb.com/manual/core/aggregation-pipeline/#id1) on a collection. To learn more about views, see [Views](https://docs.mongodb.com/manual/core/views/).
     * * `storage_databases.#.views.#.name` - Name of the view.
     * * `storage_databases.#.views.#.source` -  Name of the source collection for the view.
     * * `storage_databases.#.views.#.pipeline`- Aggregation pipeline stage(s) to apply to the source collection.
     * 
     */
    @Export(name="storageDatabases", type=List.class, parameters={DataLakeStorageDatabase.class})
    private Output<List<DataLakeStorageDatabase>> storageDatabases;

    /**
     * @return Configuration details for mapping each data store to queryable databases and collections. For complete documentation on this object and its nested fields, see [databases](https://docs.mongodb.com/datalake/reference/format/data-lake-configuration#std-label-datalake-databases-reference). An empty object indicates that the Data Lake has no mapping configuration for any data store.
     * * `storage_databases.#.name` - Name of the database to which Data Lake maps the data contained in the data store.
     * * `storage_databases.#.collections` -     Array of objects where each object represents a collection and data sources that map to a [stores](https://docs.mongodb.com/datalake/reference/format/data-lake-configuration#mongodb-datalakeconf-datalakeconf.stores) data store.
     * * `storage_databases.#.collections.#.name` - Name of the collection.
     * * `storage_databases.#.collections.#.data_sources` -     Array of objects where each object represents a stores data store to map with the collection.
     * * `storage_databases.#.collections.#.data_sources.#.store_name` -     Name of a data store to map to the `&lt;collection&gt;`. Must match the name of an object in the stores array.
     * * `storage_databases.#.collections.#.data_sources.#.default_format` - Default format that Data Lake assumes if it encounters a file without an extension while searching the storeName.
     * * `storage_databases.#.collections.#.data_sources.#.path` - Controls how Atlas Data Lake searches for and parses files in the storeName before mapping them to the `&lt;collection&gt;`.
     * * `storage_databases.#.views` -     Array of objects where each object represents an [aggregation pipeline](https://docs.mongodb.com/manual/core/aggregation-pipeline/#id1) on a collection. To learn more about views, see [Views](https://docs.mongodb.com/manual/core/views/).
     * * `storage_databases.#.views.#.name` - Name of the view.
     * * `storage_databases.#.views.#.source` -  Name of the source collection for the view.
     * * `storage_databases.#.views.#.pipeline`- Aggregation pipeline stage(s) to apply to the source collection.
     * 
     */
    public Output<List<DataLakeStorageDatabase>> storageDatabases() {
        return this.storageDatabases;
    }
    /**
     * Each object in the array represents a data store. Data Lake uses the storage.databases configuration details to map data in each data store to queryable databases and collections. For complete documentation on this object and its nested fields, see [stores](https://docs.mongodb.com/datalake/reference/format/data-lake-configuration#std-label-datalake-stores-reference). An empty object indicates that the Data Lake has no configured data stores.
     * * `storage_stores.#.name` - Name of the data store.
     * * `storage_stores.#.provider` - Defines where the data is stored.
     * * `storage_stores.#.region` - Name of the AWS region in which the S3 bucket is hosted.
     * * `storage_stores.#.bucket` - Name of the AWS S3 bucket.
     * * `storage_stores.#.prefix` - Prefix Data Lake applies when searching for files in the S3 bucket .
     * * `storage_stores.#.delimiter` - The delimiter that separates `storage_databases.#.collections.#.data_sources.#.path` segments in the data store.
     * * `storage_stores.#.include_tags` - Determines whether or not to use S3 tags on the files in the given path as additional partition attributes.
     * 
     */
    @Export(name="storageStores", type=List.class, parameters={DataLakeStorageStore.class})
    private Output<List<DataLakeStorageStore>> storageStores;

    /**
     * @return Each object in the array represents a data store. Data Lake uses the storage.databases configuration details to map data in each data store to queryable databases and collections. For complete documentation on this object and its nested fields, see [stores](https://docs.mongodb.com/datalake/reference/format/data-lake-configuration#std-label-datalake-stores-reference). An empty object indicates that the Data Lake has no configured data stores.
     * * `storage_stores.#.name` - Name of the data store.
     * * `storage_stores.#.provider` - Defines where the data is stored.
     * * `storage_stores.#.region` - Name of the AWS region in which the S3 bucket is hosted.
     * * `storage_stores.#.bucket` - Name of the AWS S3 bucket.
     * * `storage_stores.#.prefix` - Prefix Data Lake applies when searching for files in the S3 bucket .
     * * `storage_stores.#.delimiter` - The delimiter that separates `storage_databases.#.collections.#.data_sources.#.path` segments in the data store.
     * * `storage_stores.#.include_tags` - Determines whether or not to use S3 tags on the files in the given path as additional partition attributes.
     * 
     */
    public Output<List<DataLakeStorageStore>> storageStores() {
        return this.storageStores;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public DataLake(String name) {
        this(name, DataLakeArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public DataLake(String name, DataLakeArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public DataLake(String name, DataLakeArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("mongodbatlas:index/dataLake:DataLake", name, args == null ? DataLakeArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private DataLake(String name, Output<String> id, @Nullable DataLakeState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("mongodbatlas:index/dataLake:DataLake", name, state, makeResourceOptions(options, id));
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
    public static DataLake get(String name, Output<String> id, @Nullable DataLakeState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new DataLake(name, id, state, options);
    }
}
