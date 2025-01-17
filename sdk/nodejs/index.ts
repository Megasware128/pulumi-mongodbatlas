// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

// Export members:
export * from "./advancedCluster";
export * from "./alertConfiguration";
export * from "./auditing";
export * from "./cloudBackupSchedule";
export * from "./cloudBackupSnapshot";
export * from "./cloudBackupSnapshotExportBucket";
export * from "./cloudBackupSnapshotExportJob";
export * from "./cloudBackupSnapshotRestoreJob";
export * from "./cloudProviderAccess";
export * from "./cloudProviderAccessAuthorization";
export * from "./cloudProviderAccessSetup";
export * from "./cloudProviderSnapshot";
export * from "./cloudProviderSnapshotBackupPolicy";
export * from "./cloudProviderSnapshotRestoreJob";
export * from "./cluster";
export * from "./customDbRole";
export * from "./customDnsConfigurationClusterAws";
export * from "./dataLake";
export * from "./databaseUser";
export * from "./encryptionAtRest";
export * from "./eventTrigger";
export * from "./federatedSettingsIdentityProvider";
export * from "./federatedSettingsOrgConfig";
export * from "./federatedSettingsOrgRoleMapping";
export * from "./get509AuthenticationDatabaseUser";
export * from "./getAdvancedCluster";
export * from "./getAdvancedClusters";
export * from "./getAlertConfiguration";
export * from "./getAuditing";
export * from "./getCloudBackupSchedule";
export * from "./getCloudBackupSnapshot";
export * from "./getCloudBackupSnapshotExportBucket";
export * from "./getCloudBackupSnapshotExportBuckets";
export * from "./getCloudBackupSnapshotExportJob";
export * from "./getCloudBackupSnapshotExportJobs";
export * from "./getCloudBackupSnapshotRestoreJob";
export * from "./getCloudBackupSnapshotRestoreJobs";
export * from "./getCloudBackupSnapshots";
export * from "./getCloudProviderAccess";
export * from "./getCloudProviderAccessSetup";
export * from "./getCloudProviderSnapshot";
export * from "./getCloudProviderSnapshotBackupPolicy";
export * from "./getCloudProviderSnapshotRestoreJob";
export * from "./getCloudProviderSnapshotRestoreJobs";
export * from "./getCloudProviderSnapshots";
export * from "./getCluster";
export * from "./getClusters";
export * from "./getCustomDbRole";
export * from "./getCustomDbRoles";
export * from "./getCustomDnsConfigurationClusterAws";
export * from "./getDataLake";
export * from "./getDataLakes";
export * from "./getDatabaseUser";
export * from "./getDatabaseUsers";
export * from "./getEventTrigger";
export * from "./getEventTriggers";
export * from "./getFederatedSettings";
export * from "./getFederatedSettingsIdentityProvider";
export * from "./getFederatedSettingsIdentityProviders";
export * from "./getFederatedSettingsOrgConfig";
export * from "./getFederatedSettingsOrgConfigs";
export * from "./getFederatedSettingsOrgRoleMapping";
export * from "./getFederatedSettingsOrgRoleMappings";
export * from "./getGlobalClusterConfig";
export * from "./getLdapConfiguration";
export * from "./getLdapVerify";
export * from "./getMaintenanceWindow";
export * from "./getNetworkContainer";
export * from "./getNetworkContainers";
export * from "./getNetworkPeering";
export * from "./getNetworkPeerings";
export * from "./getOnlineArchive";
export * from "./getOnlineArchives";
export * from "./getOrgInvitation";
export * from "./getPrivateEndpointRegionalMode";
export * from "./getPrivateLinkEndpoint";
export * from "./getPrivateLinkEndpointService";
export * from "./getPrivatelinkEndpointServiceAdl";
export * from "./getPrivatelinkEndpointsServiceAdl";
export * from "./getProject";
export * from "./getProjectInvitation";
export * from "./getProjectIpAccessList";
export * from "./getProjects";
export * from "./getSearchIndex";
export * from "./getSearchIndexes";
export * from "./getServerlessInstance";
export * from "./getServerlessInstances";
export * from "./getTeam";
export * from "./getTeams";
export * from "./getThirdPartyIntegration";
export * from "./getThirdPartyIntegrations";
export * from "./globalClusterConfig";
export * from "./ldapConfiguration";
export * from "./ldapVerify";
export * from "./maintenanceWindow";
export * from "./networkContainer";
export * from "./networkPeering";
export * from "./onlineArchive";
export * from "./orgInvitation";
export * from "./privateEndpointRegionalMode";
export * from "./privateIpMode";
export * from "./privateLinkEndpoint";
export * from "./privateLinkEndpointService";
export * from "./privatelinkEndpointServiceAdl";
export * from "./project";
export * from "./projectInvitation";
export * from "./projectIpAccessList";
export * from "./provider";
export * from "./searchIndex";
export * from "./serverlessInstance";
export * from "./team";
export * from "./teams";
export * from "./thirdPartyIntegration";
export * from "./x509authenticationDatabaseUser";

// Export sub-modules:
import * as config from "./config";
import * as types from "./types";

export {
    config,
    types,
};

// Import resources to register:
import { AdvancedCluster } from "./advancedCluster";
import { AlertConfiguration } from "./alertConfiguration";
import { Auditing } from "./auditing";
import { CloudBackupSchedule } from "./cloudBackupSchedule";
import { CloudBackupSnapshot } from "./cloudBackupSnapshot";
import { CloudBackupSnapshotExportBucket } from "./cloudBackupSnapshotExportBucket";
import { CloudBackupSnapshotExportJob } from "./cloudBackupSnapshotExportJob";
import { CloudBackupSnapshotRestoreJob } from "./cloudBackupSnapshotRestoreJob";
import { CloudProviderAccess } from "./cloudProviderAccess";
import { CloudProviderAccessAuthorization } from "./cloudProviderAccessAuthorization";
import { CloudProviderAccessSetup } from "./cloudProviderAccessSetup";
import { CloudProviderSnapshot } from "./cloudProviderSnapshot";
import { CloudProviderSnapshotBackupPolicy } from "./cloudProviderSnapshotBackupPolicy";
import { CloudProviderSnapshotRestoreJob } from "./cloudProviderSnapshotRestoreJob";
import { Cluster } from "./cluster";
import { CustomDbRole } from "./customDbRole";
import { CustomDnsConfigurationClusterAws } from "./customDnsConfigurationClusterAws";
import { DataLake } from "./dataLake";
import { DatabaseUser } from "./databaseUser";
import { EncryptionAtRest } from "./encryptionAtRest";
import { EventTrigger } from "./eventTrigger";
import { FederatedSettingsIdentityProvider } from "./federatedSettingsIdentityProvider";
import { FederatedSettingsOrgConfig } from "./federatedSettingsOrgConfig";
import { FederatedSettingsOrgRoleMapping } from "./federatedSettingsOrgRoleMapping";
import { GlobalClusterConfig } from "./globalClusterConfig";
import { LdapConfiguration } from "./ldapConfiguration";
import { LdapVerify } from "./ldapVerify";
import { MaintenanceWindow } from "./maintenanceWindow";
import { NetworkContainer } from "./networkContainer";
import { NetworkPeering } from "./networkPeering";
import { OnlineArchive } from "./onlineArchive";
import { OrgInvitation } from "./orgInvitation";
import { PrivateEndpointRegionalMode } from "./privateEndpointRegionalMode";
import { PrivateIpMode } from "./privateIpMode";
import { PrivateLinkEndpoint } from "./privateLinkEndpoint";
import { PrivateLinkEndpointService } from "./privateLinkEndpointService";
import { PrivatelinkEndpointServiceAdl } from "./privatelinkEndpointServiceAdl";
import { Project } from "./project";
import { ProjectInvitation } from "./projectInvitation";
import { ProjectIpAccessList } from "./projectIpAccessList";
import { SearchIndex } from "./searchIndex";
import { ServerlessInstance } from "./serverlessInstance";
import { Team } from "./team";
import { Teams } from "./teams";
import { ThirdPartyIntegration } from "./thirdPartyIntegration";
import { X509AuthenticationDatabaseUser } from "./x509authenticationDatabaseUser";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "mongodbatlas:index/advancedCluster:AdvancedCluster":
                return new AdvancedCluster(name, <any>undefined, { urn })
            case "mongodbatlas:index/alertConfiguration:AlertConfiguration":
                return new AlertConfiguration(name, <any>undefined, { urn })
            case "mongodbatlas:index/auditing:Auditing":
                return new Auditing(name, <any>undefined, { urn })
            case "mongodbatlas:index/cloudBackupSchedule:CloudBackupSchedule":
                return new CloudBackupSchedule(name, <any>undefined, { urn })
            case "mongodbatlas:index/cloudBackupSnapshot:CloudBackupSnapshot":
                return new CloudBackupSnapshot(name, <any>undefined, { urn })
            case "mongodbatlas:index/cloudBackupSnapshotExportBucket:CloudBackupSnapshotExportBucket":
                return new CloudBackupSnapshotExportBucket(name, <any>undefined, { urn })
            case "mongodbatlas:index/cloudBackupSnapshotExportJob:CloudBackupSnapshotExportJob":
                return new CloudBackupSnapshotExportJob(name, <any>undefined, { urn })
            case "mongodbatlas:index/cloudBackupSnapshotRestoreJob:CloudBackupSnapshotRestoreJob":
                return new CloudBackupSnapshotRestoreJob(name, <any>undefined, { urn })
            case "mongodbatlas:index/cloudProviderAccess:CloudProviderAccess":
                return new CloudProviderAccess(name, <any>undefined, { urn })
            case "mongodbatlas:index/cloudProviderAccessAuthorization:CloudProviderAccessAuthorization":
                return new CloudProviderAccessAuthorization(name, <any>undefined, { urn })
            case "mongodbatlas:index/cloudProviderAccessSetup:CloudProviderAccessSetup":
                return new CloudProviderAccessSetup(name, <any>undefined, { urn })
            case "mongodbatlas:index/cloudProviderSnapshot:CloudProviderSnapshot":
                return new CloudProviderSnapshot(name, <any>undefined, { urn })
            case "mongodbatlas:index/cloudProviderSnapshotBackupPolicy:CloudProviderSnapshotBackupPolicy":
                return new CloudProviderSnapshotBackupPolicy(name, <any>undefined, { urn })
            case "mongodbatlas:index/cloudProviderSnapshotRestoreJob:CloudProviderSnapshotRestoreJob":
                return new CloudProviderSnapshotRestoreJob(name, <any>undefined, { urn })
            case "mongodbatlas:index/cluster:Cluster":
                return new Cluster(name, <any>undefined, { urn })
            case "mongodbatlas:index/customDbRole:CustomDbRole":
                return new CustomDbRole(name, <any>undefined, { urn })
            case "mongodbatlas:index/customDnsConfigurationClusterAws:CustomDnsConfigurationClusterAws":
                return new CustomDnsConfigurationClusterAws(name, <any>undefined, { urn })
            case "mongodbatlas:index/dataLake:DataLake":
                return new DataLake(name, <any>undefined, { urn })
            case "mongodbatlas:index/databaseUser:DatabaseUser":
                return new DatabaseUser(name, <any>undefined, { urn })
            case "mongodbatlas:index/encryptionAtRest:EncryptionAtRest":
                return new EncryptionAtRest(name, <any>undefined, { urn })
            case "mongodbatlas:index/eventTrigger:EventTrigger":
                return new EventTrigger(name, <any>undefined, { urn })
            case "mongodbatlas:index/federatedSettingsIdentityProvider:FederatedSettingsIdentityProvider":
                return new FederatedSettingsIdentityProvider(name, <any>undefined, { urn })
            case "mongodbatlas:index/federatedSettingsOrgConfig:FederatedSettingsOrgConfig":
                return new FederatedSettingsOrgConfig(name, <any>undefined, { urn })
            case "mongodbatlas:index/federatedSettingsOrgRoleMapping:FederatedSettingsOrgRoleMapping":
                return new FederatedSettingsOrgRoleMapping(name, <any>undefined, { urn })
            case "mongodbatlas:index/globalClusterConfig:GlobalClusterConfig":
                return new GlobalClusterConfig(name, <any>undefined, { urn })
            case "mongodbatlas:index/ldapConfiguration:LdapConfiguration":
                return new LdapConfiguration(name, <any>undefined, { urn })
            case "mongodbatlas:index/ldapVerify:LdapVerify":
                return new LdapVerify(name, <any>undefined, { urn })
            case "mongodbatlas:index/maintenanceWindow:MaintenanceWindow":
                return new MaintenanceWindow(name, <any>undefined, { urn })
            case "mongodbatlas:index/networkContainer:NetworkContainer":
                return new NetworkContainer(name, <any>undefined, { urn })
            case "mongodbatlas:index/networkPeering:NetworkPeering":
                return new NetworkPeering(name, <any>undefined, { urn })
            case "mongodbatlas:index/onlineArchive:OnlineArchive":
                return new OnlineArchive(name, <any>undefined, { urn })
            case "mongodbatlas:index/orgInvitation:OrgInvitation":
                return new OrgInvitation(name, <any>undefined, { urn })
            case "mongodbatlas:index/privateEndpointRegionalMode:PrivateEndpointRegionalMode":
                return new PrivateEndpointRegionalMode(name, <any>undefined, { urn })
            case "mongodbatlas:index/privateIpMode:PrivateIpMode":
                return new PrivateIpMode(name, <any>undefined, { urn })
            case "mongodbatlas:index/privateLinkEndpoint:PrivateLinkEndpoint":
                return new PrivateLinkEndpoint(name, <any>undefined, { urn })
            case "mongodbatlas:index/privateLinkEndpointService:PrivateLinkEndpointService":
                return new PrivateLinkEndpointService(name, <any>undefined, { urn })
            case "mongodbatlas:index/privatelinkEndpointServiceAdl:PrivatelinkEndpointServiceAdl":
                return new PrivatelinkEndpointServiceAdl(name, <any>undefined, { urn })
            case "mongodbatlas:index/project:Project":
                return new Project(name, <any>undefined, { urn })
            case "mongodbatlas:index/projectInvitation:ProjectInvitation":
                return new ProjectInvitation(name, <any>undefined, { urn })
            case "mongodbatlas:index/projectIpAccessList:ProjectIpAccessList":
                return new ProjectIpAccessList(name, <any>undefined, { urn })
            case "mongodbatlas:index/searchIndex:SearchIndex":
                return new SearchIndex(name, <any>undefined, { urn })
            case "mongodbatlas:index/serverlessInstance:ServerlessInstance":
                return new ServerlessInstance(name, <any>undefined, { urn })
            case "mongodbatlas:index/team:Team":
                return new Team(name, <any>undefined, { urn })
            case "mongodbatlas:index/teams:Teams":
                return new Teams(name, <any>undefined, { urn })
            case "mongodbatlas:index/thirdPartyIntegration:ThirdPartyIntegration":
                return new ThirdPartyIntegration(name, <any>undefined, { urn })
            case "mongodbatlas:index/x509AuthenticationDatabaseUser:X509AuthenticationDatabaseUser":
                return new X509AuthenticationDatabaseUser(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("mongodbatlas", "index/advancedCluster", _module)
pulumi.runtime.registerResourceModule("mongodbatlas", "index/alertConfiguration", _module)
pulumi.runtime.registerResourceModule("mongodbatlas", "index/auditing", _module)
pulumi.runtime.registerResourceModule("mongodbatlas", "index/cloudBackupSchedule", _module)
pulumi.runtime.registerResourceModule("mongodbatlas", "index/cloudBackupSnapshot", _module)
pulumi.runtime.registerResourceModule("mongodbatlas", "index/cloudBackupSnapshotExportBucket", _module)
pulumi.runtime.registerResourceModule("mongodbatlas", "index/cloudBackupSnapshotExportJob", _module)
pulumi.runtime.registerResourceModule("mongodbatlas", "index/cloudBackupSnapshotRestoreJob", _module)
pulumi.runtime.registerResourceModule("mongodbatlas", "index/cloudProviderAccess", _module)
pulumi.runtime.registerResourceModule("mongodbatlas", "index/cloudProviderAccessAuthorization", _module)
pulumi.runtime.registerResourceModule("mongodbatlas", "index/cloudProviderAccessSetup", _module)
pulumi.runtime.registerResourceModule("mongodbatlas", "index/cloudProviderSnapshot", _module)
pulumi.runtime.registerResourceModule("mongodbatlas", "index/cloudProviderSnapshotBackupPolicy", _module)
pulumi.runtime.registerResourceModule("mongodbatlas", "index/cloudProviderSnapshotRestoreJob", _module)
pulumi.runtime.registerResourceModule("mongodbatlas", "index/cluster", _module)
pulumi.runtime.registerResourceModule("mongodbatlas", "index/customDbRole", _module)
pulumi.runtime.registerResourceModule("mongodbatlas", "index/customDnsConfigurationClusterAws", _module)
pulumi.runtime.registerResourceModule("mongodbatlas", "index/dataLake", _module)
pulumi.runtime.registerResourceModule("mongodbatlas", "index/databaseUser", _module)
pulumi.runtime.registerResourceModule("mongodbatlas", "index/encryptionAtRest", _module)
pulumi.runtime.registerResourceModule("mongodbatlas", "index/eventTrigger", _module)
pulumi.runtime.registerResourceModule("mongodbatlas", "index/federatedSettingsIdentityProvider", _module)
pulumi.runtime.registerResourceModule("mongodbatlas", "index/federatedSettingsOrgConfig", _module)
pulumi.runtime.registerResourceModule("mongodbatlas", "index/federatedSettingsOrgRoleMapping", _module)
pulumi.runtime.registerResourceModule("mongodbatlas", "index/globalClusterConfig", _module)
pulumi.runtime.registerResourceModule("mongodbatlas", "index/ldapConfiguration", _module)
pulumi.runtime.registerResourceModule("mongodbatlas", "index/ldapVerify", _module)
pulumi.runtime.registerResourceModule("mongodbatlas", "index/maintenanceWindow", _module)
pulumi.runtime.registerResourceModule("mongodbatlas", "index/networkContainer", _module)
pulumi.runtime.registerResourceModule("mongodbatlas", "index/networkPeering", _module)
pulumi.runtime.registerResourceModule("mongodbatlas", "index/onlineArchive", _module)
pulumi.runtime.registerResourceModule("mongodbatlas", "index/orgInvitation", _module)
pulumi.runtime.registerResourceModule("mongodbatlas", "index/privateEndpointRegionalMode", _module)
pulumi.runtime.registerResourceModule("mongodbatlas", "index/privateIpMode", _module)
pulumi.runtime.registerResourceModule("mongodbatlas", "index/privateLinkEndpoint", _module)
pulumi.runtime.registerResourceModule("mongodbatlas", "index/privateLinkEndpointService", _module)
pulumi.runtime.registerResourceModule("mongodbatlas", "index/privatelinkEndpointServiceAdl", _module)
pulumi.runtime.registerResourceModule("mongodbatlas", "index/project", _module)
pulumi.runtime.registerResourceModule("mongodbatlas", "index/projectInvitation", _module)
pulumi.runtime.registerResourceModule("mongodbatlas", "index/projectIpAccessList", _module)
pulumi.runtime.registerResourceModule("mongodbatlas", "index/searchIndex", _module)
pulumi.runtime.registerResourceModule("mongodbatlas", "index/serverlessInstance", _module)
pulumi.runtime.registerResourceModule("mongodbatlas", "index/team", _module)
pulumi.runtime.registerResourceModule("mongodbatlas", "index/teams", _module)
pulumi.runtime.registerResourceModule("mongodbatlas", "index/thirdPartyIntegration", _module)
pulumi.runtime.registerResourceModule("mongodbatlas", "index/x509AuthenticationDatabaseUser", _module)

import { Provider } from "./provider";

pulumi.runtime.registerResourcePackage("mongodbatlas", {
    version: utilities.getVersion(),
    constructProvider: (name: string, type: string, urn: string): pulumi.ProviderResource => {
        if (type !== "pulumi:providers:mongodbatlas") {
            throw new Error(`unknown provider type ${type}`);
        }
        return new Provider(name, <any>undefined, { urn });
    },
});
