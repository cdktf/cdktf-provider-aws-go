// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ssmpatchbaseline


type SsmPatchBaselineGlobalFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.1/docs/resources/ssm_patch_baseline#key SsmPatchBaseline#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.1/docs/resources/ssm_patch_baseline#values SsmPatchBaseline#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

