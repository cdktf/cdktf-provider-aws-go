// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package emrcontainersjobtemplate


type EmrcontainersJobTemplateJobTemplateDataConfigurationOverridesMonitoringConfiguration struct {
	// cloud_watch_monitoring_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/emrcontainers_job_template#cloud_watch_monitoring_configuration EmrcontainersJobTemplate#cloud_watch_monitoring_configuration}
	CloudWatchMonitoringConfiguration *EmrcontainersJobTemplateJobTemplateDataConfigurationOverridesMonitoringConfigurationCloudWatchMonitoringConfiguration `field:"optional" json:"cloudWatchMonitoringConfiguration" yaml:"cloudWatchMonitoringConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/emrcontainers_job_template#persistent_app_ui EmrcontainersJobTemplate#persistent_app_ui}.
	PersistentAppUi *string `field:"optional" json:"persistentAppUi" yaml:"persistentAppUi"`
	// s3_monitoring_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/emrcontainers_job_template#s3_monitoring_configuration EmrcontainersJobTemplate#s3_monitoring_configuration}
	S3MonitoringConfiguration *EmrcontainersJobTemplateJobTemplateDataConfigurationOverridesMonitoringConfigurationS3MonitoringConfiguration `field:"optional" json:"s3MonitoringConfiguration" yaml:"s3MonitoringConfiguration"`
}

