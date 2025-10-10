// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package emrserverlessapplication


type EmrserverlessApplicationSchedulerConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/emrserverless_application#max_concurrent_runs EmrserverlessApplication#max_concurrent_runs}.
	MaxConcurrentRuns *float64 `field:"optional" json:"maxConcurrentRuns" yaml:"maxConcurrentRuns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/emrserverless_application#queue_timeout_minutes EmrserverlessApplication#queue_timeout_minutes}.
	QueueTimeoutMinutes *float64 `field:"optional" json:"queueTimeoutMinutes" yaml:"queueTimeoutMinutes"`
}

