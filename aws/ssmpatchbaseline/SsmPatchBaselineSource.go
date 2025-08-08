// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ssmpatchbaseline


type SsmPatchBaselineSource struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/ssm_patch_baseline#configuration SsmPatchBaseline#configuration}.
	Configuration *string `field:"required" json:"configuration" yaml:"configuration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/ssm_patch_baseline#name SsmPatchBaseline#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/ssm_patch_baseline#products SsmPatchBaseline#products}.
	Products *[]*string `field:"required" json:"products" yaml:"products"`
}

