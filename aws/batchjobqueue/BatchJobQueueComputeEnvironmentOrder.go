// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package batchjobqueue


type BatchJobQueueComputeEnvironmentOrder struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.82.1/docs/resources/batch_job_queue#compute_environment BatchJobQueue#compute_environment}.
	ComputeEnvironment *string `field:"required" json:"computeEnvironment" yaml:"computeEnvironment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.82.1/docs/resources/batch_job_queue#order BatchJobQueue#order}.
	Order *float64 `field:"required" json:"order" yaml:"order"`
}

