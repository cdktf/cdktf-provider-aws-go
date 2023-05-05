package sagemakerapp


type SagemakerAppResourceSpec struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.0/docs/resources/sagemaker_app#instance_type SagemakerApp#instance_type}.
	InstanceType *string `field:"optional" json:"instanceType" yaml:"instanceType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.0/docs/resources/sagemaker_app#lifecycle_config_arn SagemakerApp#lifecycle_config_arn}.
	LifecycleConfigArn *string `field:"optional" json:"lifecycleConfigArn" yaml:"lifecycleConfigArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.0/docs/resources/sagemaker_app#sagemaker_image_arn SagemakerApp#sagemaker_image_arn}.
	SagemakerImageArn *string `field:"optional" json:"sagemakerImageArn" yaml:"sagemakerImageArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.0/docs/resources/sagemaker_app#sagemaker_image_version_arn SagemakerApp#sagemaker_image_version_arn}.
	SagemakerImageVersionArn *string `field:"optional" json:"sagemakerImageVersionArn" yaml:"sagemakerImageVersionArn"`
}

