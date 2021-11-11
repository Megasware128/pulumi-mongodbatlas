# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from . import _utilities
import typing
# Export this package's modules as members:
from .alert_configuration import *
from .auditing import *
from .cloud_backup_schedule import *
from .cloud_provider_access import *
from .cloud_provider_access_authorization import *
from .cloud_provider_access_setup import *
from .cloud_provider_snapshot import *
from .cloud_provider_snapshot_backup_policy import *
from .cloud_provider_snapshot_restore_job import *
from .cluster import *
from .custom_db_role import *
from .custom_dns_configuration_cluster_aws import *
from .data_lake import *
from .database_user import *
from .encryption_at_rest import *
from .event_trigger import *
from .get509_authentication_database_user import *
from .get_alert_configuration import *
from .get_auditing import *
from .get_cloud_backup_schedule import *
from .get_cloud_provider_access import *
from .get_cloud_provider_access_setup import *
from .get_cloud_provider_snapshot import *
from .get_cloud_provider_snapshot_backup_policy import *
from .get_cloud_provider_snapshot_restore_job import *
from .get_cloud_provider_snapshot_restore_jobs import *
from .get_cloud_provider_snapshots import *
from .get_cluster import *
from .get_clusters import *
from .get_custom_db_role import *
from .get_custom_db_roles import *
from .get_custom_dns_configuration_cluster_aws import *
from .get_data_lake import *
from .get_data_lakes import *
from .get_database_user import *
from .get_database_users import *
from .get_event_trigger import *
from .get_event_triggers import *
from .get_global_cluster_config import *
from .get_ldap_configuration import *
from .get_ldap_verify import *
from .get_maintenance_window import *
from .get_network_container import *
from .get_network_containers import *
from .get_network_peering import *
from .get_network_peerings import *
from .get_online_archive import *
from .get_online_archives import *
from .get_private_link_endpoint import *
from .get_private_link_endpoint_service import *
from .get_project import *
from .get_project_ip_access_list import *
from .get_projects import *
from .get_search_index import *
from .get_search_indexes import *
from .get_team import *
from .get_teams import *
from .get_third_party_integration import *
from .get_third_party_integrations import *
from .global_cluster_config import *
from .ldap_configuration import *
from .ldap_verify import *
from .maintenance_window import *
from .network_container import *
from .network_peering import *
from .online_archive import *
from .private_ip_mode import *
from .private_link_endpoint import *
from .private_link_endpoint_service import *
from .project import *
from .project_ip_access_list import *
from .provider import *
from .search_index import *
from .team import *
from .teams import *
from .third_party_integration import *
from .x509_authentication_database_user import *
from ._inputs import *
from . import outputs

# Make subpackages available:
if typing.TYPE_CHECKING:
    import pulumi_mongodbatlas.config as __config
    config = __config
else:
    config = _utilities.lazy_import('pulumi_mongodbatlas.config')

_utilities.register(
    resource_modules="""
[
 {
  "pkg": "mongodbatlas",
  "mod": "index/alertConfiguration",
  "fqn": "pulumi_mongodbatlas",
  "classes": {
   "mongodbatlas:index/alertConfiguration:AlertConfiguration": "AlertConfiguration"
  }
 },
 {
  "pkg": "mongodbatlas",
  "mod": "index/auditing",
  "fqn": "pulumi_mongodbatlas",
  "classes": {
   "mongodbatlas:index/auditing:Auditing": "Auditing"
  }
 },
 {
  "pkg": "mongodbatlas",
  "mod": "index/cloudBackupSchedule",
  "fqn": "pulumi_mongodbatlas",
  "classes": {
   "mongodbatlas:index/cloudBackupSchedule:CloudBackupSchedule": "CloudBackupSchedule"
  }
 },
 {
  "pkg": "mongodbatlas",
  "mod": "index/cloudProviderAccess",
  "fqn": "pulumi_mongodbatlas",
  "classes": {
   "mongodbatlas:index/cloudProviderAccess:CloudProviderAccess": "CloudProviderAccess"
  }
 },
 {
  "pkg": "mongodbatlas",
  "mod": "index/cloudProviderAccessAuthorization",
  "fqn": "pulumi_mongodbatlas",
  "classes": {
   "mongodbatlas:index/cloudProviderAccessAuthorization:CloudProviderAccessAuthorization": "CloudProviderAccessAuthorization"
  }
 },
 {
  "pkg": "mongodbatlas",
  "mod": "index/cloudProviderAccessSetup",
  "fqn": "pulumi_mongodbatlas",
  "classes": {
   "mongodbatlas:index/cloudProviderAccessSetup:CloudProviderAccessSetup": "CloudProviderAccessSetup"
  }
 },
 {
  "pkg": "mongodbatlas",
  "mod": "index/cloudProviderSnapshot",
  "fqn": "pulumi_mongodbatlas",
  "classes": {
   "mongodbatlas:index/cloudProviderSnapshot:CloudProviderSnapshot": "CloudProviderSnapshot"
  }
 },
 {
  "pkg": "mongodbatlas",
  "mod": "index/cloudProviderSnapshotBackupPolicy",
  "fqn": "pulumi_mongodbatlas",
  "classes": {
   "mongodbatlas:index/cloudProviderSnapshotBackupPolicy:CloudProviderSnapshotBackupPolicy": "CloudProviderSnapshotBackupPolicy"
  }
 },
 {
  "pkg": "mongodbatlas",
  "mod": "index/cloudProviderSnapshotRestoreJob",
  "fqn": "pulumi_mongodbatlas",
  "classes": {
   "mongodbatlas:index/cloudProviderSnapshotRestoreJob:CloudProviderSnapshotRestoreJob": "CloudProviderSnapshotRestoreJob"
  }
 },
 {
  "pkg": "mongodbatlas",
  "mod": "index/cluster",
  "fqn": "pulumi_mongodbatlas",
  "classes": {
   "mongodbatlas:index/cluster:Cluster": "Cluster"
  }
 },
 {
  "pkg": "mongodbatlas",
  "mod": "index/customDbRole",
  "fqn": "pulumi_mongodbatlas",
  "classes": {
   "mongodbatlas:index/customDbRole:CustomDbRole": "CustomDbRole"
  }
 },
 {
  "pkg": "mongodbatlas",
  "mod": "index/customDnsConfigurationClusterAws",
  "fqn": "pulumi_mongodbatlas",
  "classes": {
   "mongodbatlas:index/customDnsConfigurationClusterAws:CustomDnsConfigurationClusterAws": "CustomDnsConfigurationClusterAws"
  }
 },
 {
  "pkg": "mongodbatlas",
  "mod": "index/dataLake",
  "fqn": "pulumi_mongodbatlas",
  "classes": {
   "mongodbatlas:index/dataLake:DataLake": "DataLake"
  }
 },
 {
  "pkg": "mongodbatlas",
  "mod": "index/databaseUser",
  "fqn": "pulumi_mongodbatlas",
  "classes": {
   "mongodbatlas:index/databaseUser:DatabaseUser": "DatabaseUser"
  }
 },
 {
  "pkg": "mongodbatlas",
  "mod": "index/encryptionAtRest",
  "fqn": "pulumi_mongodbatlas",
  "classes": {
   "mongodbatlas:index/encryptionAtRest:EncryptionAtRest": "EncryptionAtRest"
  }
 },
 {
  "pkg": "mongodbatlas",
  "mod": "index/eventTrigger",
  "fqn": "pulumi_mongodbatlas",
  "classes": {
   "mongodbatlas:index/eventTrigger:EventTrigger": "EventTrigger"
  }
 },
 {
  "pkg": "mongodbatlas",
  "mod": "index/globalClusterConfig",
  "fqn": "pulumi_mongodbatlas",
  "classes": {
   "mongodbatlas:index/globalClusterConfig:GlobalClusterConfig": "GlobalClusterConfig"
  }
 },
 {
  "pkg": "mongodbatlas",
  "mod": "index/ldapConfiguration",
  "fqn": "pulumi_mongodbatlas",
  "classes": {
   "mongodbatlas:index/ldapConfiguration:LdapConfiguration": "LdapConfiguration"
  }
 },
 {
  "pkg": "mongodbatlas",
  "mod": "index/ldapVerify",
  "fqn": "pulumi_mongodbatlas",
  "classes": {
   "mongodbatlas:index/ldapVerify:LdapVerify": "LdapVerify"
  }
 },
 {
  "pkg": "mongodbatlas",
  "mod": "index/maintenanceWindow",
  "fqn": "pulumi_mongodbatlas",
  "classes": {
   "mongodbatlas:index/maintenanceWindow:MaintenanceWindow": "MaintenanceWindow"
  }
 },
 {
  "pkg": "mongodbatlas",
  "mod": "index/networkContainer",
  "fqn": "pulumi_mongodbatlas",
  "classes": {
   "mongodbatlas:index/networkContainer:NetworkContainer": "NetworkContainer"
  }
 },
 {
  "pkg": "mongodbatlas",
  "mod": "index/networkPeering",
  "fqn": "pulumi_mongodbatlas",
  "classes": {
   "mongodbatlas:index/networkPeering:NetworkPeering": "NetworkPeering"
  }
 },
 {
  "pkg": "mongodbatlas",
  "mod": "index/onlineArchive",
  "fqn": "pulumi_mongodbatlas",
  "classes": {
   "mongodbatlas:index/onlineArchive:OnlineArchive": "OnlineArchive"
  }
 },
 {
  "pkg": "mongodbatlas",
  "mod": "index/privateIpMode",
  "fqn": "pulumi_mongodbatlas",
  "classes": {
   "mongodbatlas:index/privateIpMode:PrivateIpMode": "PrivateIpMode"
  }
 },
 {
  "pkg": "mongodbatlas",
  "mod": "index/privateLinkEndpoint",
  "fqn": "pulumi_mongodbatlas",
  "classes": {
   "mongodbatlas:index/privateLinkEndpoint:PrivateLinkEndpoint": "PrivateLinkEndpoint"
  }
 },
 {
  "pkg": "mongodbatlas",
  "mod": "index/privateLinkEndpointService",
  "fqn": "pulumi_mongodbatlas",
  "classes": {
   "mongodbatlas:index/privateLinkEndpointService:PrivateLinkEndpointService": "PrivateLinkEndpointService"
  }
 },
 {
  "pkg": "mongodbatlas",
  "mod": "index/project",
  "fqn": "pulumi_mongodbatlas",
  "classes": {
   "mongodbatlas:index/project:Project": "Project"
  }
 },
 {
  "pkg": "mongodbatlas",
  "mod": "index/projectIpAccessList",
  "fqn": "pulumi_mongodbatlas",
  "classes": {
   "mongodbatlas:index/projectIpAccessList:ProjectIpAccessList": "ProjectIpAccessList"
  }
 },
 {
  "pkg": "mongodbatlas",
  "mod": "index/searchIndex",
  "fqn": "pulumi_mongodbatlas",
  "classes": {
   "mongodbatlas:index/searchIndex:SearchIndex": "SearchIndex"
  }
 },
 {
  "pkg": "mongodbatlas",
  "mod": "index/team",
  "fqn": "pulumi_mongodbatlas",
  "classes": {
   "mongodbatlas:index/team:Team": "Team"
  }
 },
 {
  "pkg": "mongodbatlas",
  "mod": "index/teams",
  "fqn": "pulumi_mongodbatlas",
  "classes": {
   "mongodbatlas:index/teams:Teams": "Teams"
  }
 },
 {
  "pkg": "mongodbatlas",
  "mod": "index/thirdPartyIntegration",
  "fqn": "pulumi_mongodbatlas",
  "classes": {
   "mongodbatlas:index/thirdPartyIntegration:ThirdPartyIntegration": "ThirdPartyIntegration"
  }
 },
 {
  "pkg": "mongodbatlas",
  "mod": "index/x509AuthenticationDatabaseUser",
  "fqn": "pulumi_mongodbatlas",
  "classes": {
   "mongodbatlas:index/x509AuthenticationDatabaseUser:X509AuthenticationDatabaseUser": "X509AuthenticationDatabaseUser"
  }
 }
]
""",
    resource_packages="""
[
 {
  "pkg": "mongodbatlas",
  "token": "pulumi:providers:mongodbatlas",
  "fqn": "pulumi_mongodbatlas",
  "class": "Provider"
 }
]
"""
)
