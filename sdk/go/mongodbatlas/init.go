// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package mongodbatlas

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "mongodbatlas:index/alertConfiguration:AlertConfiguration":
		r, err = NewAlertConfiguration(ctx, name, nil, pulumi.URN_(urn))
	case "mongodbatlas:index/auditing:Auditing":
		r, err = NewAuditing(ctx, name, nil, pulumi.URN_(urn))
	case "mongodbatlas:index/cloudProviderAccess:CloudProviderAccess":
		r, err = NewCloudProviderAccess(ctx, name, nil, pulumi.URN_(urn))
	case "mongodbatlas:index/cloudProviderSnapshot:CloudProviderSnapshot":
		r, err = NewCloudProviderSnapshot(ctx, name, nil, pulumi.URN_(urn))
	case "mongodbatlas:index/cloudProviderSnapshotBackupPolicy:CloudProviderSnapshotBackupPolicy":
		r, err = NewCloudProviderSnapshotBackupPolicy(ctx, name, nil, pulumi.URN_(urn))
	case "mongodbatlas:index/cloudProviderSnapshotRestoreJob:CloudProviderSnapshotRestoreJob":
		r, err = NewCloudProviderSnapshotRestoreJob(ctx, name, nil, pulumi.URN_(urn))
	case "mongodbatlas:index/cluster:Cluster":
		r, err = NewCluster(ctx, name, nil, pulumi.URN_(urn))
	case "mongodbatlas:index/customDbRole:CustomDbRole":
		r, err = NewCustomDbRole(ctx, name, nil, pulumi.URN_(urn))
	case "mongodbatlas:index/databaseUser:DatabaseUser":
		r, err = NewDatabaseUser(ctx, name, nil, pulumi.URN_(urn))
	case "mongodbatlas:index/encryptionAtRest:EncryptionAtRest":
		r, err = NewEncryptionAtRest(ctx, name, nil, pulumi.URN_(urn))
	case "mongodbatlas:index/globalClusterConfig:GlobalClusterConfig":
		r, err = NewGlobalClusterConfig(ctx, name, nil, pulumi.URN_(urn))
	case "mongodbatlas:index/maintenanceWindow:MaintenanceWindow":
		r, err = NewMaintenanceWindow(ctx, name, nil, pulumi.URN_(urn))
	case "mongodbatlas:index/networkContainer:NetworkContainer":
		r, err = NewNetworkContainer(ctx, name, nil, pulumi.URN_(urn))
	case "mongodbatlas:index/networkPeering:NetworkPeering":
		r, err = NewNetworkPeering(ctx, name, nil, pulumi.URN_(urn))
	case "mongodbatlas:index/privateEndpoint:PrivateEndpoint":
		r, err = NewPrivateEndpoint(ctx, name, nil, pulumi.URN_(urn))
	case "mongodbatlas:index/privateEndpointInterfaceLink:PrivateEndpointInterfaceLink":
		r, err = NewPrivateEndpointInterfaceLink(ctx, name, nil, pulumi.URN_(urn))
	case "mongodbatlas:index/privateIpMode:PrivateIpMode":
		r, err = NewPrivateIpMode(ctx, name, nil, pulumi.URN_(urn))
	case "mongodbatlas:index/privateLinkEndpoint:PrivateLinkEndpoint":
		r, err = NewPrivateLinkEndpoint(ctx, name, nil, pulumi.URN_(urn))
	case "mongodbatlas:index/privateLinkEndpointService:PrivateLinkEndpointService":
		r, err = NewPrivateLinkEndpointService(ctx, name, nil, pulumi.URN_(urn))
	case "mongodbatlas:index/project:Project":
		r, err = NewProject(ctx, name, nil, pulumi.URN_(urn))
	case "mongodbatlas:index/projectIpAccessList:ProjectIpAccessList":
		r, err = NewProjectIpAccessList(ctx, name, nil, pulumi.URN_(urn))
	case "mongodbatlas:index/projectIpWhitelist:ProjectIpWhitelist":
		r, err = NewProjectIpWhitelist(ctx, name, nil, pulumi.URN_(urn))
	case "mongodbatlas:index/team:Team":
		r, err = NewTeam(ctx, name, nil, pulumi.URN_(urn))
	case "mongodbatlas:index/teams:Teams":
		r, err = NewTeams(ctx, name, nil, pulumi.URN_(urn))
	case "mongodbatlas:index/thirdPartyIntegration:ThirdPartyIntegration":
		r, err = NewThirdPartyIntegration(ctx, name, nil, pulumi.URN_(urn))
	case "mongodbatlas:index/x509AuthenticationDatabaseUser:X509AuthenticationDatabaseUser":
		r, err = NewX509AuthenticationDatabaseUser(ctx, name, nil, pulumi.URN_(urn))
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

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

	return NewProvider(ctx, name, nil, pulumi.URN_(urn))
}

func init() {
	version, err := PkgVersion()
	if err != nil {
		fmt.Println("failed to determine package version. defaulting to v1: %v", err)
	}
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
		"index/cloudProviderAccess",
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
		"index/globalClusterConfig",
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
		"index/privateEndpoint",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"mongodbatlas",
		"index/privateEndpointInterfaceLink",
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
		"index/project",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"mongodbatlas",
		"index/projectIpAccessList",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"mongodbatlas",
		"index/projectIpWhitelist",
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
