// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package batchjobqueue


type BatchJobQueueJobStateTimeLimitAction struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.17.0/docs/resources/batch_job_queue#action BatchJobQueue#action}.
	Action *string `field:"required" json:"action" yaml:"action"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.17.0/docs/resources/batch_job_queue#max_time_seconds BatchJobQueue#max_time_seconds}.
	MaxTimeSeconds *float64 `field:"required" json:"maxTimeSeconds" yaml:"maxTimeSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.17.0/docs/resources/batch_job_queue#reason BatchJobQueue#reason}.
	Reason *string `field:"required" json:"reason" yaml:"reason"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.17.0/docs/resources/batch_job_queue#state BatchJobQueue#state}.
	State *string `field:"required" json:"state" yaml:"state"`
}

