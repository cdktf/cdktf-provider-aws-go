package imagebuilder


type DataAwsImagebuilderImagePipelinesFilter struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/imagebuilder_image_pipelines#name DataAwsImagebuilderImagePipelines#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/imagebuilder_image_pipelines#values DataAwsImagebuilderImagePipelines#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

