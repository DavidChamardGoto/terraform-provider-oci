// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterResource() {
	tfresource.RegisterResource("oci_database_application_vip", DatabaseApplicationVipResource())
	tfresource.RegisterResource("oci_database_autonomous_container_database", DatabaseAutonomousContainerDatabaseResource())
	tfresource.RegisterResource("oci_database_autonomous_container_database_dataguard_association", DatabaseAutonomousContainerDatabaseDataguardAssociationResource())
	tfresource.RegisterResource("oci_database_autonomous_database", DatabaseAutonomousDatabaseResource())
	tfresource.RegisterResource("oci_database_autonomous_database_backup", DatabaseAutonomousDatabaseBackupResource())
	tfresource.RegisterResource("oci_database_autonomous_database_instance_wallet_management", DatabaseAutonomousDatabaseInstanceWalletManagementResource())
	tfresource.RegisterResource("oci_database_autonomous_database_regional_wallet_management", DatabaseAutonomousDatabaseRegionalWalletManagementResource())
	tfresource.RegisterResource("oci_database_autonomous_database_saas_admin_user", DatabaseAutonomousDatabaseSaasAdminUserResource())
	tfresource.RegisterResource("oci_database_autonomous_database_software_image", DatabaseAutonomousDatabaseSoftwareImageResource())
	tfresource.RegisterResource("oci_database_autonomous_database_wallet", DatabaseAutonomousDatabaseWalletResource())
	tfresource.RegisterResource("oci_database_autonomous_exadata_infrastructure", DatabaseAutonomousExadataInfrastructureResource())
	tfresource.RegisterResource("oci_database_autonomous_vm_cluster", DatabaseAutonomousVmClusterResource())
	tfresource.RegisterResource("oci_database_autonomous_vm_cluster_ords_certificate_management", DatabaseAutonomousVmClusterOrdsCertificateManagementResource())
	tfresource.RegisterResource("oci_database_autonomous_vm_cluster_ssl_certificate_management", DatabaseAutonomousVmClusterSslCertificateManagementResource())
	tfresource.RegisterResource("oci_database_backup", DatabaseBackupResource())
	tfresource.RegisterResource("oci_database_backup_cancel_management", DatabaseBackupCancelManagementResource())
	tfresource.RegisterResource("oci_database_backup_destination", DatabaseBackupDestinationResource())
	tfresource.RegisterResource("oci_database_cloud_autonomous_vm_cluster", DatabaseCloudAutonomousVmClusterResource())
	tfresource.RegisterResource("oci_database_cloud_database_management", DatabaseCloudDatabaseManagementResource())
	tfresource.RegisterResource("oci_database_cloud_exadata_infrastructure", DatabaseCloudExadataInfrastructureResource())
	tfresource.RegisterResource("oci_database_cloud_vm_cluster", DatabaseCloudVmClusterResource())
	tfresource.RegisterResource("oci_database_cloud_vm_cluster_iorm_config", DatabaseCloudVmClusterIormConfigResource())
	tfresource.RegisterResource("oci_database_data_guard_association", DatabaseDataGuardAssociationResource())
	tfresource.RegisterResource("oci_database_database", DatabaseDatabaseResource())
	tfresource.RegisterResource("oci_database_database_software_image", DatabaseDatabaseSoftwareImageResource())
	tfresource.RegisterResource("oci_database_database_upgrade", DatabaseDatabaseUpgradeResource())
	tfresource.RegisterResource("oci_database_db_home", DatabaseDbHomeResource())
	tfresource.RegisterResource("oci_database_db_node", DatabaseDbNodeResource())
	tfresource.RegisterResource("oci_database_db_node_console_connection", DatabaseDbNodeConsoleConnectionResource())
	tfresource.RegisterResource("oci_database_db_node_console_history", DatabaseDbNodeConsoleHistoryResource())
	tfresource.RegisterResource("oci_database_db_system", DatabaseDbSystemResource())
	tfresource.RegisterResource("oci_database_exadata_infrastructure", DatabaseExadataInfrastructureResource())
	tfresource.RegisterResource("oci_database_exadata_iorm_config", DatabaseExadataIormConfigResource())
	tfresource.RegisterResource("oci_database_exadb_vm_cluster", DatabaseExadbVmClusterResource())
	tfresource.RegisterResource("oci_database_exascale_db_storage_vault", DatabaseExascaleDbStorageVaultResource())
	tfresource.RegisterResource("oci_database_external_container_database", DatabaseExternalContainerDatabaseResource())
	tfresource.RegisterResource("oci_database_external_container_database_management", DatabaseExternalContainerDatabaseManagementResource())
	tfresource.RegisterResource("oci_database_external_database_connector", DatabaseExternalDatabaseConnectorResource())
	tfresource.RegisterResource("oci_database_external_non_container_database", DatabaseExternalNonContainerDatabaseResource())
	tfresource.RegisterResource("oci_database_external_non_container_database_management", DatabaseExternalNonContainerDatabaseManagementResource())
	tfresource.RegisterResource("oci_database_external_non_container_database_operations_insights_management", DatabaseExternalNonContainerDatabaseOperationsInsightsManagementResource())
	tfresource.RegisterResource("oci_database_external_pluggable_database", DatabaseExternalPluggableDatabaseResource())
	tfresource.RegisterResource("oci_database_external_pluggable_database_management", DatabaseExternalPluggableDatabaseManagementResource())
	tfresource.RegisterResource("oci_database_external_pluggable_database_operations_insights_management", DatabaseExternalPluggableDatabaseOperationsInsightsManagementResource())
	tfresource.RegisterResource("oci_database_externalcontainerdatabases_stack_monitoring", DatabaseExternalcontainerdatabasesStackMonitoringResource())
	tfresource.RegisterResource("oci_database_externalnoncontainerdatabases_stack_monitoring", DatabaseExternalnoncontainerdatabasesStackMonitoringResource())
	tfresource.RegisterResource("oci_database_externalpluggabledatabases_stack_monitoring", DatabaseExternalpluggabledatabasesStackMonitoringResource())
	tfresource.RegisterResource("oci_database_key_store", DatabaseKeyStoreResource())
	tfresource.RegisterResource("oci_database_maintenance_run", DatabaseMaintenanceRunResource())
	tfresource.RegisterResource("oci_database_migration", DatabaseMigrationResource())
	tfresource.RegisterResource("oci_database_oneoff_patch", DatabaseOneoffPatchResource())
	tfresource.RegisterResource("oci_database_pluggable_database", DatabasePluggableDatabaseResource())
	tfresource.RegisterResource("oci_database_pluggable_database_pluggabledatabasemanagements_management", DatabasePluggableDatabasePluggabledatabasemanagementsManagementResource())
	tfresource.RegisterResource("oci_database_pluggable_databases_local_clone", DatabasePluggableDatabasesLocalCloneResource())
	tfresource.RegisterResource("oci_database_pluggable_databases_remote_clone", DatabasePluggableDatabasesRemoteCloneResource())
	tfresource.RegisterResource("oci_database_vm_cluster", DatabaseVmClusterResource())
	tfresource.RegisterResource("oci_database_vm_cluster_add_virtual_machine", DatabaseVmClusterAddVirtualMachineResource())
	tfresource.RegisterResource("oci_database_vm_cluster_network", DatabaseVmClusterNetworkResource())
	tfresource.RegisterResource("oci_database_vm_cluster_remove_virtual_machine", DatabaseVmClusterRemoveVirtualMachineResource())
	tfresource.RegisterResource("oci_database_autonomous_container_database_dataguard_association_operation", DatabaseAutonomousContainerDatabaseDataguardAssociationOperationResource())
	tfresource.RegisterResource("oci_database_autonomous_container_database_dataguard_role_change", DatabaseAutonomousContainerDatabaseDataguardRoleChangeResource())
	tfresource.RegisterResource("oci_database_db_systems_upgrade", DatabaseDbSystemsUpgradeResource())
	tfresource.RegisterResource("oci_database_exadata_infrastructure_storage", DatabaseExadataInfrastructureStorageResource())
	tfresource.RegisterResource("oci_database_exadata_infrastructure_compute", DatabaseExadataInfrastructureComputeManagedResource())
}
