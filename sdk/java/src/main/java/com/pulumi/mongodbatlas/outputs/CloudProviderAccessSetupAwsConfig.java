// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.mongodbatlas.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class CloudProviderAccessSetupAwsConfig {
    private @Nullable String atlasAssumedRoleExternalId;
    private @Nullable String atlasAwsAccountArn;

    private CloudProviderAccessSetupAwsConfig() {}
    public Optional<String> atlasAssumedRoleExternalId() {
        return Optional.ofNullable(this.atlasAssumedRoleExternalId);
    }
    public Optional<String> atlasAwsAccountArn() {
        return Optional.ofNullable(this.atlasAwsAccountArn);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(CloudProviderAccessSetupAwsConfig defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String atlasAssumedRoleExternalId;
        private @Nullable String atlasAwsAccountArn;
        public Builder() {}
        public Builder(CloudProviderAccessSetupAwsConfig defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.atlasAssumedRoleExternalId = defaults.atlasAssumedRoleExternalId;
    	      this.atlasAwsAccountArn = defaults.atlasAwsAccountArn;
        }

        @CustomType.Setter
        public Builder atlasAssumedRoleExternalId(@Nullable String atlasAssumedRoleExternalId) {
            this.atlasAssumedRoleExternalId = atlasAssumedRoleExternalId;
            return this;
        }
        @CustomType.Setter
        public Builder atlasAwsAccountArn(@Nullable String atlasAwsAccountArn) {
            this.atlasAwsAccountArn = atlasAwsAccountArn;
            return this;
        }
        public CloudProviderAccessSetupAwsConfig build() {
            final var o = new CloudProviderAccessSetupAwsConfig();
            o.atlasAssumedRoleExternalId = atlasAssumedRoleExternalId;
            o.atlasAwsAccountArn = atlasAwsAccountArn;
            return o;
        }
    }
}
