// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ssmpatchbaseline


type SsmPatchBaselineApprovalRule struct {
	// patch_filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.70.0/docs/resources/ssm_patch_baseline#patch_filter SsmPatchBaseline#patch_filter}
	PatchFilter interface{} `field:"required" json:"patchFilter" yaml:"patchFilter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.70.0/docs/resources/ssm_patch_baseline#approve_after_days SsmPatchBaseline#approve_after_days}.
	ApproveAfterDays *float64 `field:"optional" json:"approveAfterDays" yaml:"approveAfterDays"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.70.0/docs/resources/ssm_patch_baseline#approve_until_date SsmPatchBaseline#approve_until_date}.
	ApproveUntilDate *string `field:"optional" json:"approveUntilDate" yaml:"approveUntilDate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.70.0/docs/resources/ssm_patch_baseline#compliance_level SsmPatchBaseline#compliance_level}.
	ComplianceLevel *string `field:"optional" json:"complianceLevel" yaml:"complianceLevel"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.70.0/docs/resources/ssm_patch_baseline#enable_non_security SsmPatchBaseline#enable_non_security}.
	EnableNonSecurity interface{} `field:"optional" json:"enableNonSecurity" yaml:"enableNonSecurity"`
}

