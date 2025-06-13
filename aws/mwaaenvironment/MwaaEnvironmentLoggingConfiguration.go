// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mwaaenvironment


type MwaaEnvironmentLoggingConfiguration struct {
	// dag_processing_logs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.100.0/docs/resources/mwaa_environment#dag_processing_logs MwaaEnvironment#dag_processing_logs}
	DagProcessingLogs *MwaaEnvironmentLoggingConfigurationDagProcessingLogs `field:"optional" json:"dagProcessingLogs" yaml:"dagProcessingLogs"`
	// scheduler_logs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.100.0/docs/resources/mwaa_environment#scheduler_logs MwaaEnvironment#scheduler_logs}
	SchedulerLogs *MwaaEnvironmentLoggingConfigurationSchedulerLogs `field:"optional" json:"schedulerLogs" yaml:"schedulerLogs"`
	// task_logs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.100.0/docs/resources/mwaa_environment#task_logs MwaaEnvironment#task_logs}
	TaskLogs *MwaaEnvironmentLoggingConfigurationTaskLogs `field:"optional" json:"taskLogs" yaml:"taskLogs"`
	// webserver_logs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.100.0/docs/resources/mwaa_environment#webserver_logs MwaaEnvironment#webserver_logs}
	WebserverLogs *MwaaEnvironmentLoggingConfigurationWebserverLogs `field:"optional" json:"webserverLogs" yaml:"webserverLogs"`
	// worker_logs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.100.0/docs/resources/mwaa_environment#worker_logs MwaaEnvironment#worker_logs}
	WorkerLogs *MwaaEnvironmentLoggingConfigurationWorkerLogs `field:"optional" json:"workerLogs" yaml:"workerLogs"`
}

