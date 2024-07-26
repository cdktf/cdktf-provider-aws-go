// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dlmlifecyclepolicy


type DlmLifecyclePolicyPolicyDetailsScheduleShareRule struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/dlm_lifecycle_policy#target_accounts DlmLifecyclePolicy#target_accounts}.
	TargetAccounts *[]*string `field:"required" json:"targetAccounts" yaml:"targetAccounts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/dlm_lifecycle_policy#unshare_interval DlmLifecyclePolicy#unshare_interval}.
	UnshareInterval *float64 `field:"optional" json:"unshareInterval" yaml:"unshareInterval"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/dlm_lifecycle_policy#unshare_interval_unit DlmLifecyclePolicy#unshare_interval_unit}.
	UnshareIntervalUnit *string `field:"optional" json:"unshareIntervalUnit" yaml:"unshareIntervalUnit"`
}

