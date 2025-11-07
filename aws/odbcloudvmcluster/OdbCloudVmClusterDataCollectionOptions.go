// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package odbcloudvmcluster


type OdbCloudVmClusterDataCollectionOptions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.20.0/docs/resources/odb_cloud_vm_cluster#is_diagnostics_events_enabled OdbCloudVmCluster#is_diagnostics_events_enabled}.
	IsDiagnosticsEventsEnabled interface{} `field:"required" json:"isDiagnosticsEventsEnabled" yaml:"isDiagnosticsEventsEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.20.0/docs/resources/odb_cloud_vm_cluster#is_health_monitoring_enabled OdbCloudVmCluster#is_health_monitoring_enabled}.
	IsHealthMonitoringEnabled interface{} `field:"required" json:"isHealthMonitoringEnabled" yaml:"isHealthMonitoringEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.20.0/docs/resources/odb_cloud_vm_cluster#is_incident_logs_enabled OdbCloudVmCluster#is_incident_logs_enabled}.
	IsIncidentLogsEnabled interface{} `field:"required" json:"isIncidentLogsEnabled" yaml:"isIncidentLogsEnabled"`
}

