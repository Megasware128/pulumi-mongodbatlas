// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.mongodbatlas;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.mongodbatlas.inputs.CloudProviderSnapshotBackupPolicyPolicyArgs;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class CloudProviderSnapshotBackupPolicyArgs extends com.pulumi.resources.ResourceArgs {

    public static final CloudProviderSnapshotBackupPolicyArgs Empty = new CloudProviderSnapshotBackupPolicyArgs();

    /**
     * The name of the Atlas cluster that contains the snapshot backup policy you want to retrieve.
     * 
     */
    @Import(name="clusterName", required=true)
    private Output<String> clusterName;

    /**
     * @return The name of the Atlas cluster that contains the snapshot backup policy you want to retrieve.
     * 
     */
    public Output<String> clusterName() {
        return this.clusterName;
    }

    /**
     * Contains a document for each backup policy item in the desired updated backup policy.
     * * `policies.#.id` - (Required) Unique identifier of the backup policy that you want to update. policies.#.id is a value obtained via the mongodbatlas.Cluster resource. provider_backup_enabled of the mongodbatlas.Cluster resource must be set to true. See the example above for how to refer to the mongodbatlas.Cluster resource for policies.#.id
     * 
     */
    @Import(name="policies", required=true)
    private Output<List<CloudProviderSnapshotBackupPolicyPolicyArgs>> policies;

    /**
     * @return Contains a document for each backup policy item in the desired updated backup policy.
     * * `policies.#.id` - (Required) Unique identifier of the backup policy that you want to update. policies.#.id is a value obtained via the mongodbatlas.Cluster resource. provider_backup_enabled of the mongodbatlas.Cluster resource must be set to true. See the example above for how to refer to the mongodbatlas.Cluster resource for policies.#.id
     * 
     */
    public Output<List<CloudProviderSnapshotBackupPolicyPolicyArgs>> policies() {
        return this.policies;
    }

    /**
     * The unique identifier of the project for the Atlas cluster.
     * 
     */
    @Import(name="projectId", required=true)
    private Output<String> projectId;

    /**
     * @return The unique identifier of the project for the Atlas cluster.
     * 
     */
    public Output<String> projectId() {
        return this.projectId;
    }

    /**
     * UTC Hour of day between 0 and 23, inclusive, representing which hour of the day that Atlas takes snapshots for backup policy items.
     * 
     */
    @Import(name="referenceHourOfDay")
    private @Nullable Output<Integer> referenceHourOfDay;

    /**
     * @return UTC Hour of day between 0 and 23, inclusive, representing which hour of the day that Atlas takes snapshots for backup policy items.
     * 
     */
    public Optional<Output<Integer>> referenceHourOfDay() {
        return Optional.ofNullable(this.referenceHourOfDay);
    }

    /**
     * UTC Minutes after referenceHourOfDay that Atlas takes snapshots for backup policy items. Must be between 0 and 59, inclusive.
     * 
     */
    @Import(name="referenceMinuteOfHour")
    private @Nullable Output<Integer> referenceMinuteOfHour;

    /**
     * @return UTC Minutes after referenceHourOfDay that Atlas takes snapshots for backup policy items. Must be between 0 and 59, inclusive.
     * 
     */
    public Optional<Output<Integer>> referenceMinuteOfHour() {
        return Optional.ofNullable(this.referenceMinuteOfHour);
    }

    /**
     * Number of days back in time you can restore to with point-in-time accuracy. Must be a positive, non-zero integer.
     * 
     */
    @Import(name="restoreWindowDays")
    private @Nullable Output<Integer> restoreWindowDays;

    /**
     * @return Number of days back in time you can restore to with point-in-time accuracy. Must be a positive, non-zero integer.
     * 
     */
    public Optional<Output<Integer>> restoreWindowDays() {
        return Optional.ofNullable(this.restoreWindowDays);
    }

    /**
     * Specify true to apply the retention changes in the updated backup policy to snapshots that Atlas took previously.
     * 
     */
    @Import(name="updateSnapshots")
    private @Nullable Output<Boolean> updateSnapshots;

    /**
     * @return Specify true to apply the retention changes in the updated backup policy to snapshots that Atlas took previously.
     * 
     */
    public Optional<Output<Boolean>> updateSnapshots() {
        return Optional.ofNullable(this.updateSnapshots);
    }

    private CloudProviderSnapshotBackupPolicyArgs() {}

    private CloudProviderSnapshotBackupPolicyArgs(CloudProviderSnapshotBackupPolicyArgs $) {
        this.clusterName = $.clusterName;
        this.policies = $.policies;
        this.projectId = $.projectId;
        this.referenceHourOfDay = $.referenceHourOfDay;
        this.referenceMinuteOfHour = $.referenceMinuteOfHour;
        this.restoreWindowDays = $.restoreWindowDays;
        this.updateSnapshots = $.updateSnapshots;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(CloudProviderSnapshotBackupPolicyArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private CloudProviderSnapshotBackupPolicyArgs $;

        public Builder() {
            $ = new CloudProviderSnapshotBackupPolicyArgs();
        }

        public Builder(CloudProviderSnapshotBackupPolicyArgs defaults) {
            $ = new CloudProviderSnapshotBackupPolicyArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param clusterName The name of the Atlas cluster that contains the snapshot backup policy you want to retrieve.
         * 
         * @return builder
         * 
         */
        public Builder clusterName(Output<String> clusterName) {
            $.clusterName = clusterName;
            return this;
        }

        /**
         * @param clusterName The name of the Atlas cluster that contains the snapshot backup policy you want to retrieve.
         * 
         * @return builder
         * 
         */
        public Builder clusterName(String clusterName) {
            return clusterName(Output.of(clusterName));
        }

        /**
         * @param policies Contains a document for each backup policy item in the desired updated backup policy.
         * * `policies.#.id` - (Required) Unique identifier of the backup policy that you want to update. policies.#.id is a value obtained via the mongodbatlas.Cluster resource. provider_backup_enabled of the mongodbatlas.Cluster resource must be set to true. See the example above for how to refer to the mongodbatlas.Cluster resource for policies.#.id
         * 
         * @return builder
         * 
         */
        public Builder policies(Output<List<CloudProviderSnapshotBackupPolicyPolicyArgs>> policies) {
            $.policies = policies;
            return this;
        }

        /**
         * @param policies Contains a document for each backup policy item in the desired updated backup policy.
         * * `policies.#.id` - (Required) Unique identifier of the backup policy that you want to update. policies.#.id is a value obtained via the mongodbatlas.Cluster resource. provider_backup_enabled of the mongodbatlas.Cluster resource must be set to true. See the example above for how to refer to the mongodbatlas.Cluster resource for policies.#.id
         * 
         * @return builder
         * 
         */
        public Builder policies(List<CloudProviderSnapshotBackupPolicyPolicyArgs> policies) {
            return policies(Output.of(policies));
        }

        /**
         * @param policies Contains a document for each backup policy item in the desired updated backup policy.
         * * `policies.#.id` - (Required) Unique identifier of the backup policy that you want to update. policies.#.id is a value obtained via the mongodbatlas.Cluster resource. provider_backup_enabled of the mongodbatlas.Cluster resource must be set to true. See the example above for how to refer to the mongodbatlas.Cluster resource for policies.#.id
         * 
         * @return builder
         * 
         */
        public Builder policies(CloudProviderSnapshotBackupPolicyPolicyArgs... policies) {
            return policies(List.of(policies));
        }

        /**
         * @param projectId The unique identifier of the project for the Atlas cluster.
         * 
         * @return builder
         * 
         */
        public Builder projectId(Output<String> projectId) {
            $.projectId = projectId;
            return this;
        }

        /**
         * @param projectId The unique identifier of the project for the Atlas cluster.
         * 
         * @return builder
         * 
         */
        public Builder projectId(String projectId) {
            return projectId(Output.of(projectId));
        }

        /**
         * @param referenceHourOfDay UTC Hour of day between 0 and 23, inclusive, representing which hour of the day that Atlas takes snapshots for backup policy items.
         * 
         * @return builder
         * 
         */
        public Builder referenceHourOfDay(@Nullable Output<Integer> referenceHourOfDay) {
            $.referenceHourOfDay = referenceHourOfDay;
            return this;
        }

        /**
         * @param referenceHourOfDay UTC Hour of day between 0 and 23, inclusive, representing which hour of the day that Atlas takes snapshots for backup policy items.
         * 
         * @return builder
         * 
         */
        public Builder referenceHourOfDay(Integer referenceHourOfDay) {
            return referenceHourOfDay(Output.of(referenceHourOfDay));
        }

        /**
         * @param referenceMinuteOfHour UTC Minutes after referenceHourOfDay that Atlas takes snapshots for backup policy items. Must be between 0 and 59, inclusive.
         * 
         * @return builder
         * 
         */
        public Builder referenceMinuteOfHour(@Nullable Output<Integer> referenceMinuteOfHour) {
            $.referenceMinuteOfHour = referenceMinuteOfHour;
            return this;
        }

        /**
         * @param referenceMinuteOfHour UTC Minutes after referenceHourOfDay that Atlas takes snapshots for backup policy items. Must be between 0 and 59, inclusive.
         * 
         * @return builder
         * 
         */
        public Builder referenceMinuteOfHour(Integer referenceMinuteOfHour) {
            return referenceMinuteOfHour(Output.of(referenceMinuteOfHour));
        }

        /**
         * @param restoreWindowDays Number of days back in time you can restore to with point-in-time accuracy. Must be a positive, non-zero integer.
         * 
         * @return builder
         * 
         */
        public Builder restoreWindowDays(@Nullable Output<Integer> restoreWindowDays) {
            $.restoreWindowDays = restoreWindowDays;
            return this;
        }

        /**
         * @param restoreWindowDays Number of days back in time you can restore to with point-in-time accuracy. Must be a positive, non-zero integer.
         * 
         * @return builder
         * 
         */
        public Builder restoreWindowDays(Integer restoreWindowDays) {
            return restoreWindowDays(Output.of(restoreWindowDays));
        }

        /**
         * @param updateSnapshots Specify true to apply the retention changes in the updated backup policy to snapshots that Atlas took previously.
         * 
         * @return builder
         * 
         */
        public Builder updateSnapshots(@Nullable Output<Boolean> updateSnapshots) {
            $.updateSnapshots = updateSnapshots;
            return this;
        }

        /**
         * @param updateSnapshots Specify true to apply the retention changes in the updated backup policy to snapshots that Atlas took previously.
         * 
         * @return builder
         * 
         */
        public Builder updateSnapshots(Boolean updateSnapshots) {
            return updateSnapshots(Output.of(updateSnapshots));
        }

        public CloudProviderSnapshotBackupPolicyArgs build() {
            $.clusterName = Objects.requireNonNull($.clusterName, "expected parameter 'clusterName' to be non-null");
            $.policies = Objects.requireNonNull($.policies, "expected parameter 'policies' to be non-null");
            $.projectId = Objects.requireNonNull($.projectId, "expected parameter 'projectId' to be non-null");
            return $;
        }
    }

}
