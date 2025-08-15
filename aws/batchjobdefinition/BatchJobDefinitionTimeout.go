// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package batchjobdefinition


type BatchJobDefinitionTimeout struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.9.0/docs/resources/batch_job_definition#attempt_duration_seconds BatchJobDefinition#attempt_duration_seconds}.
	AttemptDurationSeconds *float64 `field:"optional" json:"attemptDurationSeconds" yaml:"attemptDurationSeconds"`
}

