// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.mongodbatlas.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class X509AuthenticationDatabaseUserCertificateArgs extends com.pulumi.resources.ResourceArgs {

    public static final X509AuthenticationDatabaseUserCertificateArgs Empty = new X509AuthenticationDatabaseUserCertificateArgs();

    @Import(name="createdAt")
    private @Nullable Output<String> createdAt;

    public Optional<Output<String>> createdAt() {
        return Optional.ofNullable(this.createdAt);
    }

    @Import(name="groupId")
    private @Nullable Output<String> groupId;

    public Optional<Output<String>> groupId() {
        return Optional.ofNullable(this.groupId);
    }

    @Import(name="notAfter")
    private @Nullable Output<String> notAfter;

    public Optional<Output<String>> notAfter() {
        return Optional.ofNullable(this.notAfter);
    }

    @Import(name="subject")
    private @Nullable Output<String> subject;

    public Optional<Output<String>> subject() {
        return Optional.ofNullable(this.subject);
    }

    private X509AuthenticationDatabaseUserCertificateArgs() {}

    private X509AuthenticationDatabaseUserCertificateArgs(X509AuthenticationDatabaseUserCertificateArgs $) {
        this.createdAt = $.createdAt;
        this.groupId = $.groupId;
        this.notAfter = $.notAfter;
        this.subject = $.subject;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(X509AuthenticationDatabaseUserCertificateArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private X509AuthenticationDatabaseUserCertificateArgs $;

        public Builder() {
            $ = new X509AuthenticationDatabaseUserCertificateArgs();
        }

        public Builder(X509AuthenticationDatabaseUserCertificateArgs defaults) {
            $ = new X509AuthenticationDatabaseUserCertificateArgs(Objects.requireNonNull(defaults));
        }

        public Builder createdAt(@Nullable Output<String> createdAt) {
            $.createdAt = createdAt;
            return this;
        }

        public Builder createdAt(String createdAt) {
            return createdAt(Output.of(createdAt));
        }

        public Builder groupId(@Nullable Output<String> groupId) {
            $.groupId = groupId;
            return this;
        }

        public Builder groupId(String groupId) {
            return groupId(Output.of(groupId));
        }

        public Builder notAfter(@Nullable Output<String> notAfter) {
            $.notAfter = notAfter;
            return this;
        }

        public Builder notAfter(String notAfter) {
            return notAfter(Output.of(notAfter));
        }

        public Builder subject(@Nullable Output<String> subject) {
            $.subject = subject;
            return this;
        }

        public Builder subject(String subject) {
            return subject(Output.of(subject));
        }

        public X509AuthenticationDatabaseUserCertificateArgs build() {
            return $;
        }
    }

}
