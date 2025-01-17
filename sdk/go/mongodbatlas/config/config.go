// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package config

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

// MongoDB Atlas Base URL
func GetBaseUrl(ctx *pulumi.Context) string {
	return config.Get(ctx, "mongodbatlas:baseUrl")
}

// MongoDB Atlas Programmatic Private Key
func GetPrivateKey(ctx *pulumi.Context) string {
	return config.Get(ctx, "mongodbatlas:privateKey")
}

// MongoDB Atlas Programmatic Public Key
func GetPublicKey(ctx *pulumi.Context) string {
	return config.Get(ctx, "mongodbatlas:publicKey")
}

// MongoDB Realm Base URL
func GetRealmBaseUrl(ctx *pulumi.Context) string {
	return config.Get(ctx, "mongodbatlas:realmBaseUrl")
}
