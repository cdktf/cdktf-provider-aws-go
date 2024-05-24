// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ssmpatchbaseline


type SsmPatchBaselineApprovalRulePatchFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.51.0/docs/resources/ssm_patch_baseline#key SsmPatchBaseline#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.51.0/docs/resources/ssm_patch_baseline#values SsmPatchBaseline#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

