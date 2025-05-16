// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawsimagebuildercontainerrecipes


type DataAwsImagebuilderContainerRecipesFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.98.0/docs/data-sources/imagebuilder_container_recipes#name DataAwsImagebuilderContainerRecipes#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.98.0/docs/data-sources/imagebuilder_container_recipes#values DataAwsImagebuilderContainerRecipes#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

