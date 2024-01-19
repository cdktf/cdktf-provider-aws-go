// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawsimagebuildercomponents


type DataAwsImagebuilderComponentsFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/data-sources/imagebuilder_components#name DataAwsImagebuilderComponents#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/data-sources/imagebuilder_components#values DataAwsImagebuilderComponents#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

