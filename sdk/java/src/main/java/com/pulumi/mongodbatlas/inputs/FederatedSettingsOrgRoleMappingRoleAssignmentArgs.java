// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.mongodbatlas.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class FederatedSettingsOrgRoleMappingRoleAssignmentArgs extends com.pulumi.resources.ResourceArgs {

    public static final FederatedSettingsOrgRoleMappingRoleAssignmentArgs Empty = new FederatedSettingsOrgRoleMappingRoleAssignmentArgs();

    /**
     * Unique identifier of the project to which you want the role mapping to apply.
     * 
     */
    @Import(name="groupId")
    private @Nullable Output<String> groupId;

    /**
     * @return Unique identifier of the project to which you want the role mapping to apply.
     * 
     */
    public Optional<Output<String>> groupId() {
        return Optional.ofNullable(this.groupId);
    }

    /**
     * Unique 24-hexadecimal digit string that identifies the organization that contains your projects.
     * 
     */
    @Import(name="orgId")
    private @Nullable Output<String> orgId;

    /**
     * @return Unique 24-hexadecimal digit string that identifies the organization that contains your projects.
     * 
     */
    public Optional<Output<String>> orgId() {
        return Optional.ofNullable(this.orgId);
    }

    /**
     * Specifies the Roles that are attached to the Role Mapping.
     * 
     */
    @Import(name="roles")
    private @Nullable Output<List<String>> roles;

    /**
     * @return Specifies the Roles that are attached to the Role Mapping.
     * 
     */
    public Optional<Output<List<String>>> roles() {
        return Optional.ofNullable(this.roles);
    }

    private FederatedSettingsOrgRoleMappingRoleAssignmentArgs() {}

    private FederatedSettingsOrgRoleMappingRoleAssignmentArgs(FederatedSettingsOrgRoleMappingRoleAssignmentArgs $) {
        this.groupId = $.groupId;
        this.orgId = $.orgId;
        this.roles = $.roles;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(FederatedSettingsOrgRoleMappingRoleAssignmentArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private FederatedSettingsOrgRoleMappingRoleAssignmentArgs $;

        public Builder() {
            $ = new FederatedSettingsOrgRoleMappingRoleAssignmentArgs();
        }

        public Builder(FederatedSettingsOrgRoleMappingRoleAssignmentArgs defaults) {
            $ = new FederatedSettingsOrgRoleMappingRoleAssignmentArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param groupId Unique identifier of the project to which you want the role mapping to apply.
         * 
         * @return builder
         * 
         */
        public Builder groupId(@Nullable Output<String> groupId) {
            $.groupId = groupId;
            return this;
        }

        /**
         * @param groupId Unique identifier of the project to which you want the role mapping to apply.
         * 
         * @return builder
         * 
         */
        public Builder groupId(String groupId) {
            return groupId(Output.of(groupId));
        }

        /**
         * @param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects.
         * 
         * @return builder
         * 
         */
        public Builder orgId(@Nullable Output<String> orgId) {
            $.orgId = orgId;
            return this;
        }

        /**
         * @param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects.
         * 
         * @return builder
         * 
         */
        public Builder orgId(String orgId) {
            return orgId(Output.of(orgId));
        }

        /**
         * @param roles Specifies the Roles that are attached to the Role Mapping.
         * 
         * @return builder
         * 
         */
        public Builder roles(@Nullable Output<List<String>> roles) {
            $.roles = roles;
            return this;
        }

        /**
         * @param roles Specifies the Roles that are attached to the Role Mapping.
         * 
         * @return builder
         * 
         */
        public Builder roles(List<String> roles) {
            return roles(Output.of(roles));
        }

        /**
         * @param roles Specifies the Roles that are attached to the Role Mapping.
         * 
         * @return builder
         * 
         */
        public Builder roles(String... roles) {
            return roles(List.of(roles));
        }

        public FederatedSettingsOrgRoleMappingRoleAssignmentArgs build() {
            return $;
        }
    }

}
