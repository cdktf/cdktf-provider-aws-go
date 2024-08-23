// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package imagebuildercontainerrecipe


type ImagebuilderContainerRecipeTargetRepository struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.64.0/docs/resources/imagebuilder_container_recipe#repository_name ImagebuilderContainerRecipe#repository_name}.
	RepositoryName *string `field:"required" json:"repositoryName" yaml:"repositoryName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.64.0/docs/resources/imagebuilder_container_recipe#service ImagebuilderContainerRecipe#service}.
	Service *string `field:"required" json:"service" yaml:"service"`
}

