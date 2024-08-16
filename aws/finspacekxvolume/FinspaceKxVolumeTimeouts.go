// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package finspacekxvolume


type FinspaceKxVolumeTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.0/docs/resources/finspace_kx_volume#create FinspaceKxVolume#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.0/docs/resources/finspace_kx_volume#delete FinspaceKxVolume#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.0/docs/resources/finspace_kx_volume#update FinspaceKxVolume#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

