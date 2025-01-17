// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.mongodbatlas.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;


public final class CloudProviderSnapshotBackupPolicyPolicyPolicyItemArgs extends com.pulumi.resources.ResourceArgs {

    public static final CloudProviderSnapshotBackupPolicyPolicyPolicyItemArgs Empty = new CloudProviderSnapshotBackupPolicyPolicyPolicyItemArgs();

    @Import(name="frequencyInterval", required=true)
    private Output<Integer> frequencyInterval;

    public Output<Integer> frequencyInterval() {
        return this.frequencyInterval;
    }

    @Import(name="frequencyType", required=true)
    private Output<String> frequencyType;

    public Output<String> frequencyType() {
        return this.frequencyType;
    }

    @Import(name="id", required=true)
    private Output<String> id;

    public Output<String> id() {
        return this.id;
    }

    @Import(name="retentionUnit", required=true)
    private Output<String> retentionUnit;

    public Output<String> retentionUnit() {
        return this.retentionUnit;
    }

    @Import(name="retentionValue", required=true)
    private Output<Integer> retentionValue;

    public Output<Integer> retentionValue() {
        return this.retentionValue;
    }

    private CloudProviderSnapshotBackupPolicyPolicyPolicyItemArgs() {}

    private CloudProviderSnapshotBackupPolicyPolicyPolicyItemArgs(CloudProviderSnapshotBackupPolicyPolicyPolicyItemArgs $) {
        this.frequencyInterval = $.frequencyInterval;
        this.frequencyType = $.frequencyType;
        this.id = $.id;
        this.retentionUnit = $.retentionUnit;
        this.retentionValue = $.retentionValue;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(CloudProviderSnapshotBackupPolicyPolicyPolicyItemArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private CloudProviderSnapshotBackupPolicyPolicyPolicyItemArgs $;

        public Builder() {
            $ = new CloudProviderSnapshotBackupPolicyPolicyPolicyItemArgs();
        }

        public Builder(CloudProviderSnapshotBackupPolicyPolicyPolicyItemArgs defaults) {
            $ = new CloudProviderSnapshotBackupPolicyPolicyPolicyItemArgs(Objects.requireNonNull(defaults));
        }

        public Builder frequencyInterval(Output<Integer> frequencyInterval) {
            $.frequencyInterval = frequencyInterval;
            return this;
        }

        public Builder frequencyInterval(Integer frequencyInterval) {
            return frequencyInterval(Output.of(frequencyInterval));
        }

        public Builder frequencyType(Output<String> frequencyType) {
            $.frequencyType = frequencyType;
            return this;
        }

        public Builder frequencyType(String frequencyType) {
            return frequencyType(Output.of(frequencyType));
        }

        public Builder id(Output<String> id) {
            $.id = id;
            return this;
        }

        public Builder id(String id) {
            return id(Output.of(id));
        }

        public Builder retentionUnit(Output<String> retentionUnit) {
            $.retentionUnit = retentionUnit;
            return this;
        }

        public Builder retentionUnit(String retentionUnit) {
            return retentionUnit(Output.of(retentionUnit));
        }

        public Builder retentionValue(Output<Integer> retentionValue) {
            $.retentionValue = retentionValue;
            return this;
        }

        public Builder retentionValue(Integer retentionValue) {
            return retentionValue(Output.of(retentionValue));
        }

        public CloudProviderSnapshotBackupPolicyPolicyPolicyItemArgs build() {
            $.frequencyInterval = Objects.requireNonNull($.frequencyInterval, "expected parameter 'frequencyInterval' to be non-null");
            $.frequencyType = Objects.requireNonNull($.frequencyType, "expected parameter 'frequencyType' to be non-null");
            $.id = Objects.requireNonNull($.id, "expected parameter 'id' to be non-null");
            $.retentionUnit = Objects.requireNonNull($.retentionUnit, "expected parameter 'retentionUnit' to be non-null");
            $.retentionValue = Objects.requireNonNull($.retentionValue, "expected parameter 'retentionValue' to be non-null");
            return $;
        }
    }

}
