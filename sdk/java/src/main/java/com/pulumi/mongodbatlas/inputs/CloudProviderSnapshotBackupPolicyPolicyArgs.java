// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.mongodbatlas.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.mongodbatlas.inputs.CloudProviderSnapshotBackupPolicyPolicyPolicyItemArgs;
import java.util.List;
import java.util.Objects;


public final class CloudProviderSnapshotBackupPolicyPolicyArgs extends com.pulumi.resources.ResourceArgs {

    public static final CloudProviderSnapshotBackupPolicyPolicyArgs Empty = new CloudProviderSnapshotBackupPolicyPolicyArgs();

    @Import(name="policyItems", required=true)
    private Output<List<CloudProviderSnapshotBackupPolicyPolicyPolicyItemArgs>> policyItems;

    public Output<List<CloudProviderSnapshotBackupPolicyPolicyPolicyItemArgs>> policyItems() {
        return this.policyItems;
    }

    private CloudProviderSnapshotBackupPolicyPolicyArgs() {}

    private CloudProviderSnapshotBackupPolicyPolicyArgs(CloudProviderSnapshotBackupPolicyPolicyArgs $) {
        this.policyItems = $.policyItems;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(CloudProviderSnapshotBackupPolicyPolicyArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private CloudProviderSnapshotBackupPolicyPolicyArgs $;

        public Builder() {
            $ = new CloudProviderSnapshotBackupPolicyPolicyArgs();
        }

        public Builder(CloudProviderSnapshotBackupPolicyPolicyArgs defaults) {
            $ = new CloudProviderSnapshotBackupPolicyPolicyArgs(Objects.requireNonNull(defaults));
        }

        public Builder policyItems(Output<List<CloudProviderSnapshotBackupPolicyPolicyPolicyItemArgs>> policyItems) {
            $.policyItems = policyItems;
            return this;
        }

        public Builder policyItems(List<CloudProviderSnapshotBackupPolicyPolicyPolicyItemArgs> policyItems) {
            return policyItems(Output.of(policyItems));
        }

        public Builder policyItems(CloudProviderSnapshotBackupPolicyPolicyPolicyItemArgs... policyItems) {
            return policyItems(List.of(policyItems));
        }

        public CloudProviderSnapshotBackupPolicyPolicyArgs build() {
            $.policyItems = Objects.requireNonNull($.policyItems, "expected parameter 'policyItems' to be non-null");
            return $;
        }
    }

}
