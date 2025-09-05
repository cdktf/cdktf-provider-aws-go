// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dlmlifecyclepolicy


type DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRule struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.12.0/docs/resources/dlm_lifecycle_policy#encrypted DlmLifecyclePolicy#encrypted}.
	Encrypted interface{} `field:"required" json:"encrypted" yaml:"encrypted"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.12.0/docs/resources/dlm_lifecycle_policy#cmk_arn DlmLifecyclePolicy#cmk_arn}.
	CmkArn *string `field:"optional" json:"cmkArn" yaml:"cmkArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.12.0/docs/resources/dlm_lifecycle_policy#copy_tags DlmLifecyclePolicy#copy_tags}.
	CopyTags interface{} `field:"optional" json:"copyTags" yaml:"copyTags"`
	// deprecate_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.12.0/docs/resources/dlm_lifecycle_policy#deprecate_rule DlmLifecyclePolicy#deprecate_rule}
	DeprecateRule *DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleDeprecateRule `field:"optional" json:"deprecateRule" yaml:"deprecateRule"`
	// retain_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.12.0/docs/resources/dlm_lifecycle_policy#retain_rule DlmLifecyclePolicy#retain_rule}
	RetainRule *DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleRetainRule `field:"optional" json:"retainRule" yaml:"retainRule"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.12.0/docs/resources/dlm_lifecycle_policy#target DlmLifecyclePolicy#target}.
	Target *string `field:"optional" json:"target" yaml:"target"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.12.0/docs/resources/dlm_lifecycle_policy#target_region DlmLifecyclePolicy#target_region}.
	TargetRegion *string `field:"optional" json:"targetRegion" yaml:"targetRegion"`
}

