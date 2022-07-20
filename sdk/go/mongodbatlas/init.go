// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mongodbatlas

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "mongodbatlas:index/advancedCluster:AdvancedCluster":
		r = &AdvancedCluster{}
	case "mongodbatlas:index/alertConfiguration:AlertConfiguration":
		r = &AlertConfiguration{}
	case "mongodbatlas:index/auditing:Auditing":
		r = &Auditing{}
	case "mongodbatlas:index/cloudBackupSchedule:CloudBackupSchedule":
		r = &CloudBackupSchedule{}
	case "mongodbatlas:index/cloudBackupSnapshot:CloudBackupSnapshot":
		r = &CloudBackupSnapshot{}
	case "mongodbatlas:index/cloudBackupSnapshotExportBucket:CloudBackupSnapshotExportBucket":
		r = &CloudBackupSnapshotExportBucket{}
	case "mongodbatlas:index/cloudBackupSnapshotExportJob:CloudBackupSnapshotExportJob":
		r = &CloudBackupSnapshotExportJob{}
	case "mongodbatlas:index/cloudBackupSnapshotRestoreJob:CloudBackupSnapshotRestoreJob":
		r = &CloudBackupSnapshotRestoreJob{}
	case "mongodbatlas:index/cloudProviderAccess:CloudProviderAccess":
		r = &CloudProviderAccess{}
	case "mongodbatlas:index/cloudProviderAccessAuthorization:CloudProviderAccessAuthorization":
		r = &CloudProviderAccessAuthorization{}
	case "mongodbatlas:index/cloudProviderAccessSetup:CloudProviderAccessSetup":
		r = &CloudProviderAccessSetup{}
	case "mongodbatlas:index/cloudProviderSnapshot:CloudProviderSnapshot":
		r = &CloudProviderSnapshot{}
	case "mongodbatlas:index/cloudProviderSnapshotBackupPolicy:CloudProviderSnapshotBackupPolicy":
		r = &CloudProviderSnapshotBackupPolicy{}
	case "mongodbatlas:index/cloudProviderSnapshotRestoreJob:CloudProviderSnapshotRestoreJob":
		r = &CloudProviderSnapshotRestoreJob{}
	case "mongodbatlas:index/cluster:Cluster":
		r = &Cluster{}
	case "mongodbatlas:index/customDbRole:CustomDbRole":
		r = &CustomDbRole{}
	case "mongodbatlas:index/customDnsConfigurationClusterAws:CustomDnsConfigurationClusterAws":
		r = &CustomDnsConfigurationClusterAws{}
	case "mongodbatlas:index/dataLake:DataLake":
		r = &DataLake{}
	case "mongodbatlas:index/databaseUser:DatabaseUser":
		r = &DatabaseUser{}
	case "mongodbatlas:index/encryptionAtRest:EncryptionAtRest":
		r = &EncryptionAtRest{}
	case "mongodbatlas:index/eventTrigger:EventTrigger":
		r = &EventTrigger{}
	case "mongodbatlas:index/federatedSettingsIdentityProvider:FederatedSettingsIdentityProvider":
		r = &FederatedSettingsIdentityProvider{}
	case "mongodbatlas:index/federatedSettingsOrgConfig:FederatedSettingsOrgConfig":
		r = &FederatedSettingsOrgConfig{}
	case "mongodbatlas:index/federatedSettingsOrgRoleMapping:FederatedSettingsOrgRoleMapping":
		r = &FederatedSettingsOrgRoleMapping{}
	case "mongodbatlas:index/globalClusterConfig:GlobalClusterConfig":
		r = &GlobalClusterConfig{}
	case "mongodbatlas:index/ldapConfiguration:LdapConfiguration":
		r = &LdapConfiguration{}
	case "mongodbatlas:index/ldapVerify:LdapVerify":
		r = &LdapVerify{}
	case "mongodbatlas:index/maintenanceWindow:MaintenanceWindow":
		r = &MaintenanceWindow{}
	case "mongodbatlas:index/networkContainer:NetworkContainer":
		r = &NetworkContainer{}
	case "mongodbatlas:index/networkPeering:NetworkPeering":
		r = &NetworkPeering{}
	case "mongodbatlas:index/onlineArchive:OnlineArchive":
		r = &OnlineArchive{}
	case "mongodbatlas:index/orgInvitation:OrgInvitation":
		r = &OrgInvitation{}
	case "mongodbatlas:index/privateEndpointRegionalMode:PrivateEndpointRegionalMode":
		r = &PrivateEndpointRegionalMode{}
	case "mongodbatlas:index/privateIpMode:PrivateIpMode":
		r = &PrivateIpMode{}
	case "mongodbatlas:index/privateLinkEndpoint:PrivateLinkEndpoint":
		r = &PrivateLinkEndpoint{}
	case "mongodbatlas:index/privateLinkEndpointService:PrivateLinkEndpointService":
		r = &PrivateLinkEndpointService{}
	case "mongodbatlas:index/privatelinkEndpointServiceAdl:PrivatelinkEndpointServiceAdl":
		r = &PrivatelinkEndpointServiceAdl{}
	case "mongodbatlas:index/project:Project":
		r = &Project{}
	case "mongodbatlas:index/projectInvitation:ProjectInvitation":
		r = &ProjectInvitation{}
	case "mongodbatlas:index/projectIpAccessList:ProjectIpAccessList":
		r = &ProjectIpAccessList{}
	case "mongodbatlas:index/searchIndex:SearchIndex":
		r = &SearchIndex{}
	case "mongodbatlas:index/serverlessInstance:ServerlessInstance":
		r = &ServerlessInstance{}
	case "mongodbatlas:index/team:Team":
		r = &Team{}
	case "mongodbatlas:index/teams:Teams":
		r = &Teams{}
	case "mongodbatlas:index/thirdPartyIntegration:ThirdPartyIntegration":
		r = &ThirdPartyIntegration{}
	case "mongodbatlas:index/x509AuthenticationDatabaseUser:X509AuthenticationDatabaseUser":
		r = &X509AuthenticationDatabaseUser{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

type pkg struct {
	version semver.Version
}

func (p *pkg) Version() semver.Version {
	return p.version
}

func (p *pkg) ConstructProvider(ctx *pulumi.Context, name, typ, urn string) (pulumi.ProviderResource, error) {
	if typ != "pulumi:providers:mongodbatlas" {
		return nil, fmt.Errorf("unknown provider type: %s", typ)
	}

	r := &Provider{}
	err := ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return r, err
}

func init() {
	version, _ := PkgVersion()
	pulumi.RegisterResourceModule(
		"mongodbatlas",
		"index/advancedCluster",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"mongodbatlas",
		"index/alertConfiguration",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"mongodbatlas",
		"index/auditing",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"mongodbatlas",
		"index/cloudBackupSchedule",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"mongodbatlas",
		"index/cloudBackupSnapshot",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"mongodbatlas",
		"index/cloudBackupSnapshotExportBucket",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"mongodbatlas",
		"index/cloudBackupSnapshotExportJob",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"mongodbatlas",
		"index/cloudBackupSnapshotRestoreJob",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"mongodbatlas",
		"index/cloudProviderAccess",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"mongodbatlas",
		"index/cloudProviderAccessAuthorization",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"mongodbatlas",
		"index/cloudProviderAccessSetup",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"mongodbatlas",
		"index/cloudProviderSnapshot",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"mongodbatlas",
		"index/cloudProviderSnapshotBackupPolicy",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"mongodbatlas",
		"index/cloudProviderSnapshotRestoreJob",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"mongodbatlas",
		"index/cluster",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"mongodbatlas",
		"index/customDbRole",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"mongodbatlas",
		"index/customDnsConfigurationClusterAws",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"mongodbatlas",
		"index/dataLake",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"mongodbatlas",
		"index/databaseUser",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"mongodbatlas",
		"index/encryptionAtRest",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"mongodbatlas",
		"index/eventTrigger",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"mongodbatlas",
		"index/federatedSettingsIdentityProvider",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"mongodbatlas",
		"index/federatedSettingsOrgConfig",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"mongodbatlas",
		"index/federatedSettingsOrgRoleMapping",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"mongodbatlas",
		"index/globalClusterConfig",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"mongodbatlas",
		"index/ldapConfiguration",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"mongodbatlas",
		"index/ldapVerify",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"mongodbatlas",
		"index/maintenanceWindow",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"mongodbatlas",
		"index/networkContainer",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"mongodbatlas",
		"index/networkPeering",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"mongodbatlas",
		"index/onlineArchive",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"mongodbatlas",
		"index/orgInvitation",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"mongodbatlas",
		"index/privateEndpointRegionalMode",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"mongodbatlas",
		"index/privateIpMode",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"mongodbatlas",
		"index/privateLinkEndpoint",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"mongodbatlas",
		"index/privateLinkEndpointService",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"mongodbatlas",
		"index/privatelinkEndpointServiceAdl",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"mongodbatlas",
		"index/project",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"mongodbatlas",
		"index/projectInvitation",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"mongodbatlas",
		"index/projectIpAccessList",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"mongodbatlas",
		"index/searchIndex",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"mongodbatlas",
		"index/serverlessInstance",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"mongodbatlas",
		"index/team",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"mongodbatlas",
		"index/teams",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"mongodbatlas",
		"index/thirdPartyIntegration",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"mongodbatlas",
		"index/x509AuthenticationDatabaseUser",
		&module{version},
	)
	pulumi.RegisterResourcePackage(
		"mongodbatlas",
		&pkg{version},
	)
}
