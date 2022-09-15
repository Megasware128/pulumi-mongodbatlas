// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.mongodbatlas;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.mongodbatlas.inputs.DataLakeAwsArgs;
import com.pulumi.mongodbatlas.inputs.DataLakeDataProcessRegionArgs;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class DataLakeArgs extends com.pulumi.resources.ResourceArgs {

    public static final DataLakeArgs Empty = new DataLakeArgs();

    /**
     * AWS provider of the cloud service where Data Lake can access the S3 Bucket.
     * * `aws.0.role_id` - (Required) Unique identifier of the role that Data Lake can use to access the data stores. If necessary, use the Atlas [UI](https://docs.atlas.mongodb.com/security/manage-iam-roles/) or [API](https://docs.atlas.mongodb.com/reference/api/cloud-provider-access-get-roles/) to retrieve the role ID. You must also specify the `aws.0.test_s3_bucket`.
     * * `aws.0.test_s3_bucket` - (Required) Name of the S3 data bucket that the provided role ID is authorized to access. You must also specify the `aws.0.role_id`.
     * 
     */
    @Import(name="aws", required=true)
    private Output<DataLakeAwsArgs> aws;

    /**
     * @return AWS provider of the cloud service where Data Lake can access the S3 Bucket.
     * * `aws.0.role_id` - (Required) Unique identifier of the role that Data Lake can use to access the data stores. If necessary, use the Atlas [UI](https://docs.atlas.mongodb.com/security/manage-iam-roles/) or [API](https://docs.atlas.mongodb.com/reference/api/cloud-provider-access-get-roles/) to retrieve the role ID. You must also specify the `aws.0.test_s3_bucket`.
     * * `aws.0.test_s3_bucket` - (Required) Name of the S3 data bucket that the provided role ID is authorized to access. You must also specify the `aws.0.role_id`.
     * 
     */
    public Output<DataLakeAwsArgs> aws() {
        return this.aws;
    }

    /**
     * The cloud provider region to which Atlas Data Lake routes client connections for data processing. Set to `null` to direct Atlas Data Lake to route client connections to the region nearest to the client based on DNS resolution.
     * * `data_process_region.0.cloud_provider` - (Required) Name of the cloud service provider. Atlas Data Lake only supports AWS.
     * * `data_process_region.0.region` - (Required). Name of the region to which Data Lake routes client connections for data processing. Atlas Data Lake only supports the following regions:
     * 
     */
    @Import(name="dataProcessRegion")
    private @Nullable Output<DataLakeDataProcessRegionArgs> dataProcessRegion;

    /**
     * @return The cloud provider region to which Atlas Data Lake routes client connections for data processing. Set to `null` to direct Atlas Data Lake to route client connections to the region nearest to the client based on DNS resolution.
     * * `data_process_region.0.cloud_provider` - (Required) Name of the cloud service provider. Atlas Data Lake only supports AWS.
     * * `data_process_region.0.region` - (Required). Name of the region to which Data Lake routes client connections for data processing. Atlas Data Lake only supports the following regions:
     * 
     */
    public Optional<Output<DataLakeDataProcessRegionArgs>> dataProcessRegion() {
        return Optional.ofNullable(this.dataProcessRegion);
    }

    /**
     * Name of the Atlas Data Lake.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Name of the Atlas Data Lake.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The unique ID for the project to create a data lake.
     * 
     */
    @Import(name="projectId", required=true)
    private Output<String> projectId;

    /**
     * @return The unique ID for the project to create a data lake.
     * 
     */
    public Output<String> projectId() {
        return this.projectId;
    }

    private DataLakeArgs() {}

    private DataLakeArgs(DataLakeArgs $) {
        this.aws = $.aws;
        this.dataProcessRegion = $.dataProcessRegion;
        this.name = $.name;
        this.projectId = $.projectId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(DataLakeArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private DataLakeArgs $;

        public Builder() {
            $ = new DataLakeArgs();
        }

        public Builder(DataLakeArgs defaults) {
            $ = new DataLakeArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param aws AWS provider of the cloud service where Data Lake can access the S3 Bucket.
         * * `aws.0.role_id` - (Required) Unique identifier of the role that Data Lake can use to access the data stores. If necessary, use the Atlas [UI](https://docs.atlas.mongodb.com/security/manage-iam-roles/) or [API](https://docs.atlas.mongodb.com/reference/api/cloud-provider-access-get-roles/) to retrieve the role ID. You must also specify the `aws.0.test_s3_bucket`.
         * * `aws.0.test_s3_bucket` - (Required) Name of the S3 data bucket that the provided role ID is authorized to access. You must also specify the `aws.0.role_id`.
         * 
         * @return builder
         * 
         */
        public Builder aws(Output<DataLakeAwsArgs> aws) {
            $.aws = aws;
            return this;
        }

        /**
         * @param aws AWS provider of the cloud service where Data Lake can access the S3 Bucket.
         * * `aws.0.role_id` - (Required) Unique identifier of the role that Data Lake can use to access the data stores. If necessary, use the Atlas [UI](https://docs.atlas.mongodb.com/security/manage-iam-roles/) or [API](https://docs.atlas.mongodb.com/reference/api/cloud-provider-access-get-roles/) to retrieve the role ID. You must also specify the `aws.0.test_s3_bucket`.
         * * `aws.0.test_s3_bucket` - (Required) Name of the S3 data bucket that the provided role ID is authorized to access. You must also specify the `aws.0.role_id`.
         * 
         * @return builder
         * 
         */
        public Builder aws(DataLakeAwsArgs aws) {
            return aws(Output.of(aws));
        }

        /**
         * @param dataProcessRegion The cloud provider region to which Atlas Data Lake routes client connections for data processing. Set to `null` to direct Atlas Data Lake to route client connections to the region nearest to the client based on DNS resolution.
         * * `data_process_region.0.cloud_provider` - (Required) Name of the cloud service provider. Atlas Data Lake only supports AWS.
         * * `data_process_region.0.region` - (Required). Name of the region to which Data Lake routes client connections for data processing. Atlas Data Lake only supports the following regions:
         * 
         * @return builder
         * 
         */
        public Builder dataProcessRegion(@Nullable Output<DataLakeDataProcessRegionArgs> dataProcessRegion) {
            $.dataProcessRegion = dataProcessRegion;
            return this;
        }

        /**
         * @param dataProcessRegion The cloud provider region to which Atlas Data Lake routes client connections for data processing. Set to `null` to direct Atlas Data Lake to route client connections to the region nearest to the client based on DNS resolution.
         * * `data_process_region.0.cloud_provider` - (Required) Name of the cloud service provider. Atlas Data Lake only supports AWS.
         * * `data_process_region.0.region` - (Required). Name of the region to which Data Lake routes client connections for data processing. Atlas Data Lake only supports the following regions:
         * 
         * @return builder
         * 
         */
        public Builder dataProcessRegion(DataLakeDataProcessRegionArgs dataProcessRegion) {
            return dataProcessRegion(Output.of(dataProcessRegion));
        }

        /**
         * @param name Name of the Atlas Data Lake.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Name of the Atlas Data Lake.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param projectId The unique ID for the project to create a data lake.
         * 
         * @return builder
         * 
         */
        public Builder projectId(Output<String> projectId) {
            $.projectId = projectId;
            return this;
        }

        /**
         * @param projectId The unique ID for the project to create a data lake.
         * 
         * @return builder
         * 
         */
        public Builder projectId(String projectId) {
            return projectId(Output.of(projectId));
        }

        public DataLakeArgs build() {
            $.aws = Objects.requireNonNull($.aws, "expected parameter 'aws' to be non-null");
            $.projectId = Objects.requireNonNull($.projectId, "expected parameter 'projectId' to be non-null");
            return $;
        }
    }

}