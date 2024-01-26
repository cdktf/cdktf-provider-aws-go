// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package finspacekxvolume


type FinspaceKxVolumeNas1Configuration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.34.0/docs/resources/finspace_kx_volume#size FinspaceKxVolume#size}.
	Size *float64 `field:"required" json:"size" yaml:"size"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.34.0/docs/resources/finspace_kx_volume#type FinspaceKxVolume#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
}

