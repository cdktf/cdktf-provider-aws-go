package sagemaker


type SagemakerAppImageConfigKernelGatewayImageConfigKernelSpec struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/sagemaker_app_image_config#name SagemakerAppImageConfig#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/sagemaker_app_image_config#display_name SagemakerAppImageConfig#display_name}.
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
}

