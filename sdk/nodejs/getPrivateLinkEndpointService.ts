// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * `mongodbatlas.PrivateLinkEndpointService` describe a Private Endpoint Link. This represents a Private Endpoint Link Connection that wants to retrieve details in an Atlas project.
 *
 * > **NOTE:** Groups and projects are synonymous terms. You may find groupId in the official documentation.
 */
export function getPrivateLinkEndpointService(args: GetPrivateLinkEndpointServiceArgs, opts?: pulumi.InvokeOptions): Promise<GetPrivateLinkEndpointServiceResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("mongodbatlas:index/getPrivateLinkEndpointService:getPrivateLinkEndpointService", {
        "endpointServiceId": args.endpointServiceId,
        "privateLinkId": args.privateLinkId,
        "projectId": args.projectId,
        "providerName": args.providerName,
    }, opts);
}

/**
 * A collection of arguments for invoking getPrivateLinkEndpointService.
 */
export interface GetPrivateLinkEndpointServiceArgs {
    /**
     * Unique identifier of the `AWS` or `AZURE` resource.
     */
    readonly endpointServiceId: string;
    /**
     * Unique identifier of the private endpoint service for which you want to retrieve a private endpoint.
     */
    readonly privateLinkId: string;
    /**
     * Unique identifier for the project.
     */
    readonly projectId: string;
    /**
     * Cloud provider for which you want to create a private endpoint. Atlas accepts `AWS` or `AZURE`.
     */
    readonly providerName: string;
}

/**
 * A collection of values returned by getPrivateLinkEndpointService.
 */
export interface GetPrivateLinkEndpointServiceResult {
    /**
     * Status of the interface endpoint for AWS.
     * Returns one of the following values:
     */
    readonly awsConnectionStatus: string;
    /**
     * Status of the interface endpoint for AZURE.
     * Returns one of the following values:
     */
    readonly azureStatus: string;
    /**
     * Indicates if Atlas received a request to remove the interface endpoint from the private endpoint connection.
     */
    readonly deleteRequested: boolean;
    readonly endpointServiceId: string;
    /**
     * Error message pertaining to the interface endpoint. Returns null if there are no errors.
     */
    readonly errorMessage: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Unique identifier of the interface endpoint.
     */
    readonly interfaceEndpointId: string;
    /**
     * Name of the connection for this private endpoint that Atlas generates.
     */
    readonly privateEndpointConnectionName: string;
    /**
     * Private IP address of the private endpoint network interface.
     */
    readonly privateEndpointIpAddress: string;
    /**
     * Unique identifier of the private endpoint.
     */
    readonly privateEndpointResourceId: string;
    readonly privateLinkId: string;
    readonly projectId: string;
    readonly providerName: string;
}
