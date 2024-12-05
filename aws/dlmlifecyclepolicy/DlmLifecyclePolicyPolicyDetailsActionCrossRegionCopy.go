// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dlmlifecyclepolicy


type DlmLifecyclePolicyPolicyDetailsActionCrossRegionCopy struct {
	// encryption_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.80.0/docs/resources/dlm_lifecycle_policy#encryption_configuration DlmLifecyclePolicy#encryption_configuration}
	EncryptionConfiguration *DlmLifecyclePolicyPolicyDetailsActionCrossRegionCopyEncryptionConfiguration `field:"required" json:"encryptionConfiguration" yaml:"encryptionConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.80.0/docs/resources/dlm_lifecycle_policy#target DlmLifecyclePolicy#target}.
	Target *string `field:"required" json:"target" yaml:"target"`
	// retain_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.80.0/docs/resources/dlm_lifecycle_policy#retain_rule DlmLifecyclePolicy#retain_rule}
	RetainRule *DlmLifecyclePolicyPolicyDetailsActionCrossRegionCopyRetainRule `field:"optional" json:"retainRule" yaml:"retainRule"`
}

