// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.mongodbatlas;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.mongodbatlas.TeamsArgs;
import com.pulumi.mongodbatlas.Utilities;
import com.pulumi.mongodbatlas.inputs.TeamsState;
import java.lang.String;
import java.util.List;
import javax.annotation.Nullable;

@ResourceType(type="mongodbatlas:index/teams:Teams")
public class Teams extends com.pulumi.resources.CustomResource {
    @Export(name="name", type=String.class, parameters={})
    private Output<String> name;

    public Output<String> name() {
        return this.name;
    }
    @Export(name="orgId", type=String.class, parameters={})
    private Output<String> orgId;

    public Output<String> orgId() {
        return this.orgId;
    }
    @Export(name="teamId", type=String.class, parameters={})
    private Output<String> teamId;

    public Output<String> teamId() {
        return this.teamId;
    }
    @Export(name="usernames", type=List.class, parameters={String.class})
    private Output<List<String>> usernames;

    public Output<List<String>> usernames() {
        return this.usernames;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Teams(String name) {
        this(name, TeamsArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Teams(String name, TeamsArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Teams(String name, TeamsArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("mongodbatlas:index/teams:Teams", name, args == null ? TeamsArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Teams(String name, Output<String> id, @Nullable TeamsState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("mongodbatlas:index/teams:Teams", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static Teams get(String name, Output<String> id, @Nullable TeamsState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Teams(name, id, state, options);
    }
}