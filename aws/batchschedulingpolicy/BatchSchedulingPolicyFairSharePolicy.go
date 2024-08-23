// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package batchschedulingpolicy


type BatchSchedulingPolicyFairSharePolicy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.64.0/docs/resources/batch_scheduling_policy#compute_reservation BatchSchedulingPolicy#compute_reservation}.
	ComputeReservation *float64 `field:"optional" json:"computeReservation" yaml:"computeReservation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.64.0/docs/resources/batch_scheduling_policy#share_decay_seconds BatchSchedulingPolicy#share_decay_seconds}.
	ShareDecaySeconds *float64 `field:"optional" json:"shareDecaySeconds" yaml:"shareDecaySeconds"`
	// share_distribution block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.64.0/docs/resources/batch_scheduling_policy#share_distribution BatchSchedulingPolicy#share_distribution}
	ShareDistribution interface{} `field:"optional" json:"shareDistribution" yaml:"shareDistribution"`
}

