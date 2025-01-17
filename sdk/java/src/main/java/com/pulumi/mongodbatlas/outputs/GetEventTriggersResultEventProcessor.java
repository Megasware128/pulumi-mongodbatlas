// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.mongodbatlas.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.mongodbatlas.outputs.GetEventTriggersResultEventProcessorAwsEventbridge;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetEventTriggersResultEventProcessor {
    private List<GetEventTriggersResultEventProcessorAwsEventbridge> awsEventbridges;

    private GetEventTriggersResultEventProcessor() {}
    public List<GetEventTriggersResultEventProcessorAwsEventbridge> awsEventbridges() {
        return this.awsEventbridges;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetEventTriggersResultEventProcessor defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private List<GetEventTriggersResultEventProcessorAwsEventbridge> awsEventbridges;
        public Builder() {}
        public Builder(GetEventTriggersResultEventProcessor defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.awsEventbridges = defaults.awsEventbridges;
        }

        @CustomType.Setter
        public Builder awsEventbridges(List<GetEventTriggersResultEventProcessorAwsEventbridge> awsEventbridges) {
            this.awsEventbridges = Objects.requireNonNull(awsEventbridges);
            return this;
        }
        public Builder awsEventbridges(GetEventTriggersResultEventProcessorAwsEventbridge... awsEventbridges) {
            return awsEventbridges(List.of(awsEventbridges));
        }
        public GetEventTriggersResultEventProcessor build() {
            final var o = new GetEventTriggersResultEventProcessor();
            o.awsEventbridges = awsEventbridges;
            return o;
        }
    }
}
