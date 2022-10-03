package sagemakeruserprofile


type SagemakerUserProfileUserSettingsKernelGatewayAppSettings struct {
	// default_resource_spec block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/sagemaker_user_profile#default_resource_spec SagemakerUserProfile#default_resource_spec}
	DefaultResourceSpec *SagemakerUserProfileUserSettingsKernelGatewayAppSettingsDefaultResourceSpec `field:"required" json:"defaultResourceSpec" yaml:"defaultResourceSpec"`
	// custom_image block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/sagemaker_user_profile#custom_image SagemakerUserProfile#custom_image}
	CustomImage interface{} `field:"optional" json:"customImage" yaml:"customImage"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/sagemaker_user_profile#lifecycle_config_arns SagemakerUserProfile#lifecycle_config_arns}.
	LifecycleConfigArns *[]*string `field:"optional" json:"lifecycleConfigArns" yaml:"lifecycleConfigArns"`
}

