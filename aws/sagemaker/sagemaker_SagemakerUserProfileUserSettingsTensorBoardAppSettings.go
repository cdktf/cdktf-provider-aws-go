package sagemaker


type SagemakerUserProfileUserSettingsTensorBoardAppSettings struct {
	// default_resource_spec block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/sagemaker_user_profile#default_resource_spec SagemakerUserProfile#default_resource_spec}
	DefaultResourceSpec *SagemakerUserProfileUserSettingsTensorBoardAppSettingsDefaultResourceSpec `field:"required" json:"defaultResourceSpec" yaml:"defaultResourceSpec"`
}

