// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package imagebuildercontainerrecipe


type ImagebuilderContainerRecipeComponent struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.79.0/docs/resources/imagebuilder_container_recipe#component_arn ImagebuilderContainerRecipe#component_arn}.
	ComponentArn *string `field:"required" json:"componentArn" yaml:"componentArn"`
	// parameter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.79.0/docs/resources/imagebuilder_container_recipe#parameter ImagebuilderContainerRecipe#parameter}
	Parameter interface{} `field:"optional" json:"parameter" yaml:"parameter"`
}

