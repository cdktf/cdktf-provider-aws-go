// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package emrserverlessapplication


type EmrserverlessApplicationMonitoringConfiguration struct {
	// cloudwatch_logging_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.19.0/docs/resources/emrserverless_application#cloudwatch_logging_configuration EmrserverlessApplication#cloudwatch_logging_configuration}
	CloudwatchLoggingConfiguration *EmrserverlessApplicationMonitoringConfigurationCloudwatchLoggingConfiguration `field:"optional" json:"cloudwatchLoggingConfiguration" yaml:"cloudwatchLoggingConfiguration"`
	// managed_persistence_monitoring_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.19.0/docs/resources/emrserverless_application#managed_persistence_monitoring_configuration EmrserverlessApplication#managed_persistence_monitoring_configuration}
	ManagedPersistenceMonitoringConfiguration *EmrserverlessApplicationMonitoringConfigurationManagedPersistenceMonitoringConfiguration `field:"optional" json:"managedPersistenceMonitoringConfiguration" yaml:"managedPersistenceMonitoringConfiguration"`
	// prometheus_monitoring_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.19.0/docs/resources/emrserverless_application#prometheus_monitoring_configuration EmrserverlessApplication#prometheus_monitoring_configuration}
	PrometheusMonitoringConfiguration *EmrserverlessApplicationMonitoringConfigurationPrometheusMonitoringConfiguration `field:"optional" json:"prometheusMonitoringConfiguration" yaml:"prometheusMonitoringConfiguration"`
	// s3_monitoring_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.19.0/docs/resources/emrserverless_application#s3_monitoring_configuration EmrserverlessApplication#s3_monitoring_configuration}
	S3MonitoringConfiguration *EmrserverlessApplicationMonitoringConfigurationS3MonitoringConfiguration `field:"optional" json:"s3MonitoringConfiguration" yaml:"s3MonitoringConfiguration"`
}

