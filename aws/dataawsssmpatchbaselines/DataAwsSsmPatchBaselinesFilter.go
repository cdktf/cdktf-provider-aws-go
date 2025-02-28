// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawsssmpatchbaselines


type DataAwsSsmPatchBaselinesFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.89.0/docs/data-sources/ssm_patch_baselines#key DataAwsSsmPatchBaselines#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.89.0/docs/data-sources/ssm_patch_baselines#values DataAwsSsmPatchBaselines#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

