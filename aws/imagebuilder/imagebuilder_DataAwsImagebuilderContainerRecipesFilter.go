package imagebuilder


type DataAwsImagebuilderContainerRecipesFilter struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/imagebuilder_container_recipes#name DataAwsImagebuilderContainerRecipes#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/imagebuilder_container_recipes#values DataAwsImagebuilderContainerRecipes#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

