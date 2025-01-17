// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.mongodbatlas;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class OrgInvitationArgs extends com.pulumi.resources.ResourceArgs {

    public static final OrgInvitationArgs Empty = new OrgInvitationArgs();

    /**
     * Unique 24-hexadecimal digit string that identifies the organization to which you want to invite a user.
     * 
     */
    @Import(name="orgId", required=true)
    private Output<String> orgId;

    /**
     * @return Unique 24-hexadecimal digit string that identifies the organization to which you want to invite a user.
     * 
     */
    public Output<String> orgId() {
        return this.orgId;
    }

    /**
     * Atlas roles to assign to the invited user. If the user accepts the invitation, Atlas assigns these roles to them. The following options are available:
     * * ORG_OWNER
     * * ORG_GROUP_CREATOR
     * * ORG_BILLING_ADMIN
     * * ORG_READ_ONLY
     * * ORG_MEMBER
     * 
     */
    @Import(name="roles", required=true)
    private Output<List<String>> roles;

    /**
     * @return Atlas roles to assign to the invited user. If the user accepts the invitation, Atlas assigns these roles to them. The following options are available:
     * * ORG_OWNER
     * * ORG_GROUP_CREATOR
     * * ORG_BILLING_ADMIN
     * * ORG_READ_ONLY
     * * ORG_MEMBER
     * 
     */
    public Output<List<String>> roles() {
        return this.roles;
    }

    /**
     * An array of unique 24-hexadecimal digit strings that identify the teams that the user was invited to join.
     * 
     */
    @Import(name="teamsIds")
    private @Nullable Output<List<String>> teamsIds;

    /**
     * @return An array of unique 24-hexadecimal digit strings that identify the teams that the user was invited to join.
     * 
     */
    public Optional<Output<List<String>>> teamsIds() {
        return Optional.ofNullable(this.teamsIds);
    }

    /**
     * Email address of the invited user. This is the address to which Atlas sends the invite. If the user accepts the invitation, they log in to Atlas with this username.
     * 
     */
    @Import(name="username", required=true)
    private Output<String> username;

    /**
     * @return Email address of the invited user. This is the address to which Atlas sends the invite. If the user accepts the invitation, they log in to Atlas with this username.
     * 
     */
    public Output<String> username() {
        return this.username;
    }

    private OrgInvitationArgs() {}

    private OrgInvitationArgs(OrgInvitationArgs $) {
        this.orgId = $.orgId;
        this.roles = $.roles;
        this.teamsIds = $.teamsIds;
        this.username = $.username;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(OrgInvitationArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private OrgInvitationArgs $;

        public Builder() {
            $ = new OrgInvitationArgs();
        }

        public Builder(OrgInvitationArgs defaults) {
            $ = new OrgInvitationArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param orgId Unique 24-hexadecimal digit string that identifies the organization to which you want to invite a user.
         * 
         * @return builder
         * 
         */
        public Builder orgId(Output<String> orgId) {
            $.orgId = orgId;
            return this;
        }

        /**
         * @param orgId Unique 24-hexadecimal digit string that identifies the organization to which you want to invite a user.
         * 
         * @return builder
         * 
         */
        public Builder orgId(String orgId) {
            return orgId(Output.of(orgId));
        }

        /**
         * @param roles Atlas roles to assign to the invited user. If the user accepts the invitation, Atlas assigns these roles to them. The following options are available:
         * * ORG_OWNER
         * * ORG_GROUP_CREATOR
         * * ORG_BILLING_ADMIN
         * * ORG_READ_ONLY
         * * ORG_MEMBER
         * 
         * @return builder
         * 
         */
        public Builder roles(Output<List<String>> roles) {
            $.roles = roles;
            return this;
        }

        /**
         * @param roles Atlas roles to assign to the invited user. If the user accepts the invitation, Atlas assigns these roles to them. The following options are available:
         * * ORG_OWNER
         * * ORG_GROUP_CREATOR
         * * ORG_BILLING_ADMIN
         * * ORG_READ_ONLY
         * * ORG_MEMBER
         * 
         * @return builder
         * 
         */
        public Builder roles(List<String> roles) {
            return roles(Output.of(roles));
        }

        /**
         * @param roles Atlas roles to assign to the invited user. If the user accepts the invitation, Atlas assigns these roles to them. The following options are available:
         * * ORG_OWNER
         * * ORG_GROUP_CREATOR
         * * ORG_BILLING_ADMIN
         * * ORG_READ_ONLY
         * * ORG_MEMBER
         * 
         * @return builder
         * 
         */
        public Builder roles(String... roles) {
            return roles(List.of(roles));
        }

        /**
         * @param teamsIds An array of unique 24-hexadecimal digit strings that identify the teams that the user was invited to join.
         * 
         * @return builder
         * 
         */
        public Builder teamsIds(@Nullable Output<List<String>> teamsIds) {
            $.teamsIds = teamsIds;
            return this;
        }

        /**
         * @param teamsIds An array of unique 24-hexadecimal digit strings that identify the teams that the user was invited to join.
         * 
         * @return builder
         * 
         */
        public Builder teamsIds(List<String> teamsIds) {
            return teamsIds(Output.of(teamsIds));
        }

        /**
         * @param teamsIds An array of unique 24-hexadecimal digit strings that identify the teams that the user was invited to join.
         * 
         * @return builder
         * 
         */
        public Builder teamsIds(String... teamsIds) {
            return teamsIds(List.of(teamsIds));
        }

        /**
         * @param username Email address of the invited user. This is the address to which Atlas sends the invite. If the user accepts the invitation, they log in to Atlas with this username.
         * 
         * @return builder
         * 
         */
        public Builder username(Output<String> username) {
            $.username = username;
            return this;
        }

        /**
         * @param username Email address of the invited user. This is the address to which Atlas sends the invite. If the user accepts the invitation, they log in to Atlas with this username.
         * 
         * @return builder
         * 
         */
        public Builder username(String username) {
            return username(Output.of(username));
        }

        public OrgInvitationArgs build() {
            $.orgId = Objects.requireNonNull($.orgId, "expected parameter 'orgId' to be non-null");
            $.roles = Objects.requireNonNull($.roles, "expected parameter 'roles' to be non-null");
            $.username = Objects.requireNonNull($.username, "expected parameter 'username' to be non-null");
            return $;
        }
    }

}
