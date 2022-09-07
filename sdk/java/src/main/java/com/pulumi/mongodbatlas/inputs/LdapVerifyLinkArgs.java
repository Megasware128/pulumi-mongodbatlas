// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.mongodbatlas.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class LdapVerifyLinkArgs extends com.pulumi.resources.ResourceArgs {

    public static final LdapVerifyLinkArgs Empty = new LdapVerifyLinkArgs();

    @Import(name="href")
    private @Nullable Output<String> href;

    public Optional<Output<String>> href() {
        return Optional.ofNullable(this.href);
    }

    @Import(name="rel")
    private @Nullable Output<String> rel;

    public Optional<Output<String>> rel() {
        return Optional.ofNullable(this.rel);
    }

    private LdapVerifyLinkArgs() {}

    private LdapVerifyLinkArgs(LdapVerifyLinkArgs $) {
        this.href = $.href;
        this.rel = $.rel;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(LdapVerifyLinkArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private LdapVerifyLinkArgs $;

        public Builder() {
            $ = new LdapVerifyLinkArgs();
        }

        public Builder(LdapVerifyLinkArgs defaults) {
            $ = new LdapVerifyLinkArgs(Objects.requireNonNull(defaults));
        }

        public Builder href(@Nullable Output<String> href) {
            $.href = href;
            return this;
        }

        public Builder href(String href) {
            return href(Output.of(href));
        }

        public Builder rel(@Nullable Output<String> rel) {
            $.rel = rel;
            return this;
        }

        public Builder rel(String rel) {
            return rel(Output.of(rel));
        }

        public LdapVerifyLinkArgs build() {
            return $;
        }
    }

}
