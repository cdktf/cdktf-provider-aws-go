// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawsimagebuilderimagerecipes


type DataAwsImagebuilderImageRecipesFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.100.0/docs/data-sources/imagebuilder_image_recipes#name DataAwsImagebuilderImageRecipes#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.100.0/docs/data-sources/imagebuilder_image_recipes#values DataAwsImagebuilderImageRecipes#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

