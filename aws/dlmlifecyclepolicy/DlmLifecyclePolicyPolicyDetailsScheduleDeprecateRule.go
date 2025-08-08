// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dlmlifecyclepolicy


type DlmLifecyclePolicyPolicyDetailsScheduleDeprecateRule struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/dlm_lifecycle_policy#count DlmLifecyclePolicy#count}.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/dlm_lifecycle_policy#interval DlmLifecyclePolicy#interval}.
	Interval *float64 `field:"optional" json:"interval" yaml:"interval"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/dlm_lifecycle_policy#interval_unit DlmLifecyclePolicy#interval_unit}.
	IntervalUnit *string `field:"optional" json:"intervalUnit" yaml:"intervalUnit"`
}

