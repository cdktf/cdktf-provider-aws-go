package imagebuildercontainerrecipe


type ImagebuilderContainerRecipeComponentParameter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.2/docs/resources/imagebuilder_container_recipe#name ImagebuilderContainerRecipe#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.2/docs/resources/imagebuilder_container_recipe#value ImagebuilderContainerRecipe#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

