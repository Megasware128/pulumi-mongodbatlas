// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.mongodbatlas;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.mongodbatlas.CloudBackupSnapshotExportBucketArgs;
import com.pulumi.mongodbatlas.Utilities;
import com.pulumi.mongodbatlas.inputs.CloudBackupSnapshotExportBucketState;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * `mongodbatlas.CloudBackupSnapshotExportBucket` resource allows you to create an export snapshot bucket for the specified project.
 * 
 * &gt; **NOTE:** Groups and projects are synonymous terms. You may find `groupId` in the official documentation.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.mongodbatlas.CloudBackupSnapshotExportBucket;
 * import com.pulumi.mongodbatlas.CloudBackupSnapshotExportBucketArgs;
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
 *         var test = new CloudBackupSnapshotExportBucket(&#34;test&#34;, CloudBackupSnapshotExportBucketArgs.builder()        
 *             .bucketName(&#34;example-bucket&#34;)
 *             .cloudProvider(&#34;AWS&#34;)
 *             .iamRoleId(&#34;{IAM_ROLE_ID}&#34;)
 *             .projectId(&#34;{PROJECT_ID}&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Cloud Backup Snapshot Export Backup entries can be imported using project project_id, and bucket_id (Unique identifier of the snapshot export bucket), in the format `PROJECTID-BUCKETID`, e.g.
 * 
 * ```sh
 *  $ pulumi import mongodbatlas:index/cloudBackupSnapshotExportBucket:CloudBackupSnapshotExportBucket test 5d0f1f73cf09a29120e173cf-5d116d82014b764445b2f9b5
 * ```
 * 
 *  For more information see[MongoDB Atlas API Reference.](https://docs.atlas.mongodb.com/reference/api/cloud-backup/export/create-one-export-bucket/)
 * 
 */
@ResourceType(type="mongodbatlas:index/cloudBackupSnapshotExportBucket:CloudBackupSnapshotExportBucket")
public class CloudBackupSnapshotExportBucket extends com.pulumi.resources.CustomResource {
    /**
     * Name of the bucket that the provided role ID is authorized to access. You must also specify the `iam_role_id`.
     * 
     */
    @Export(name="bucketName", type=String.class, parameters={})
    private Output<String> bucketName;

    /**
     * @return Name of the bucket that the provided role ID is authorized to access. You must also specify the `iam_role_id`.
     * 
     */
    public Output<String> bucketName() {
        return this.bucketName;
    }
    /**
     * Name of the provider of the cloud service where Atlas can access the S3 bucket. Atlas only supports `AWS`.
     * 
     */
    @Export(name="cloudProvider", type=String.class, parameters={})
    private Output<String> cloudProvider;

    /**
     * @return Name of the provider of the cloud service where Atlas can access the S3 bucket. Atlas only supports `AWS`.
     * 
     */
    public Output<String> cloudProvider() {
        return this.cloudProvider;
    }
    /**
     * Unique identifier of the snapshot export bucket.
     * 
     */
    @Export(name="exportBucketId", type=String.class, parameters={})
    private Output<String> exportBucketId;

    /**
     * @return Unique identifier of the snapshot export bucket.
     * 
     */
    public Output<String> exportBucketId() {
        return this.exportBucketId;
    }
    /**
     * Unique identifier of the role that Atlas can use to access the bucket. You must also specify the `bucket_name`.
     * 
     */
    @Export(name="iamRoleId", type=String.class, parameters={})
    private Output<String> iamRoleId;

    /**
     * @return Unique identifier of the role that Atlas can use to access the bucket. You must also specify the `bucket_name`.
     * 
     */
    public Output<String> iamRoleId() {
        return this.iamRoleId;
    }
    /**
     * The unique identifier of the project for the Atlas cluster.
     * 
     */
    @Export(name="projectId", type=String.class, parameters={})
    private Output<String> projectId;

    /**
     * @return The unique identifier of the project for the Atlas cluster.
     * 
     */
    public Output<String> projectId() {
        return this.projectId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public CloudBackupSnapshotExportBucket(String name) {
        this(name, CloudBackupSnapshotExportBucketArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public CloudBackupSnapshotExportBucket(String name, CloudBackupSnapshotExportBucketArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public CloudBackupSnapshotExportBucket(String name, CloudBackupSnapshotExportBucketArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("mongodbatlas:index/cloudBackupSnapshotExportBucket:CloudBackupSnapshotExportBucket", name, args == null ? CloudBackupSnapshotExportBucketArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private CloudBackupSnapshotExportBucket(String name, Output<String> id, @Nullable CloudBackupSnapshotExportBucketState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("mongodbatlas:index/cloudBackupSnapshotExportBucket:CloudBackupSnapshotExportBucket", name, state, makeResourceOptions(options, id));
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
    public static CloudBackupSnapshotExportBucket get(String name, Output<String> id, @Nullable CloudBackupSnapshotExportBucketState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new CloudBackupSnapshotExportBucket(name, id, state, options);
    }
}
