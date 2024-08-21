// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package batchcomputeenvironment


type BatchComputeEnvironmentUpdatePolicy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.1/docs/resources/batch_compute_environment#job_execution_timeout_minutes BatchComputeEnvironment#job_execution_timeout_minutes}.
	JobExecutionTimeoutMinutes *float64 `field:"required" json:"jobExecutionTimeoutMinutes" yaml:"jobExecutionTimeoutMinutes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.1/docs/resources/batch_compute_environment#terminate_jobs_on_update BatchComputeEnvironment#terminate_jobs_on_update}.
	TerminateJobsOnUpdate interface{} `field:"required" json:"terminateJobsOnUpdate" yaml:"terminateJobsOnUpdate"`
}

