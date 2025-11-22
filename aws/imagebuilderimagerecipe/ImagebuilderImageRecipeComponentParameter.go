// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package imagebuilderimagerecipe


type ImagebuilderImageRecipeComponentParameter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/imagebuilder_image_recipe#name ImagebuilderImageRecipe#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/imagebuilder_image_recipe#value ImagebuilderImageRecipe#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

