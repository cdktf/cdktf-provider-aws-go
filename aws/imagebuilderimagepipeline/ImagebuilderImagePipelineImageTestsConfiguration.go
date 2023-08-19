package imagebuilderimagepipeline


type ImagebuilderImagePipelineImageTestsConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/imagebuilder_image_pipeline#image_tests_enabled ImagebuilderImagePipeline#image_tests_enabled}.
	ImageTestsEnabled interface{} `field:"optional" json:"imageTestsEnabled" yaml:"imageTestsEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/imagebuilder_image_pipeline#timeout_minutes ImagebuilderImagePipeline#timeout_minutes}.
	TimeoutMinutes *float64 `field:"optional" json:"timeoutMinutes" yaml:"timeoutMinutes"`
}

