// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.mongodbatlas.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class GetServerlessInstanceLink extends com.pulumi.resources.InvokeArgs {

    public static final GetServerlessInstanceLink Empty = new GetServerlessInstanceLink();

    @Import(name="href", required=true)
    private String href;

    public String href() {
        return this.href;
    }

    @Import(name="rel", required=true)
    private String rel;

    public String rel() {
        return this.rel;
    }

    private GetServerlessInstanceLink() {}

    private GetServerlessInstanceLink(GetServerlessInstanceLink $) {
        this.href = $.href;
        this.rel = $.rel;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetServerlessInstanceLink defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetServerlessInstanceLink $;

        public Builder() {
            $ = new GetServerlessInstanceLink();
        }

        public Builder(GetServerlessInstanceLink defaults) {
            $ = new GetServerlessInstanceLink(Objects.requireNonNull(defaults));
        }

        public Builder href(String href) {
            $.href = href;
            return this;
        }

        public Builder rel(String rel) {
            $.rel = rel;
            return this;
        }

        public GetServerlessInstanceLink build() {
            $.href = Objects.requireNonNull($.href, "expected parameter 'href' to be non-null");
            $.rel = Objects.requireNonNull($.rel, "expected parameter 'rel' to be non-null");
            return $;
        }
    }

}
