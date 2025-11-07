// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dlmlifecyclepolicy


type DlmLifecyclePolicyPolicyDetailsScheduleCreateRule struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.20.0/docs/resources/dlm_lifecycle_policy#cron_expression DlmLifecyclePolicy#cron_expression}.
	CronExpression *string `field:"optional" json:"cronExpression" yaml:"cronExpression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.20.0/docs/resources/dlm_lifecycle_policy#interval DlmLifecyclePolicy#interval}.
	Interval *float64 `field:"optional" json:"interval" yaml:"interval"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.20.0/docs/resources/dlm_lifecycle_policy#interval_unit DlmLifecyclePolicy#interval_unit}.
	IntervalUnit *string `field:"optional" json:"intervalUnit" yaml:"intervalUnit"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.20.0/docs/resources/dlm_lifecycle_policy#location DlmLifecyclePolicy#location}.
	Location *string `field:"optional" json:"location" yaml:"location"`
	// scripts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.20.0/docs/resources/dlm_lifecycle_policy#scripts DlmLifecyclePolicy#scripts}
	Scripts *DlmLifecyclePolicyPolicyDetailsScheduleCreateRuleScripts `field:"optional" json:"scripts" yaml:"scripts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.20.0/docs/resources/dlm_lifecycle_policy#times DlmLifecyclePolicy#times}.
	Times *[]*string `field:"optional" json:"times" yaml:"times"`
}

