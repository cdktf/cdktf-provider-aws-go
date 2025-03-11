// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sagemakeruserprofile


type SagemakerUserProfileUserSettingsRSessionAppSettingsCustomImage struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.1/docs/resources/sagemaker_user_profile#app_image_config_name SagemakerUserProfile#app_image_config_name}.
	AppImageConfigName *string `field:"required" json:"appImageConfigName" yaml:"appImageConfigName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.1/docs/resources/sagemaker_user_profile#image_name SagemakerUserProfile#image_name}.
	ImageName *string `field:"required" json:"imageName" yaml:"imageName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.1/docs/resources/sagemaker_user_profile#image_version_number SagemakerUserProfile#image_version_number}.
	ImageVersionNumber *float64 `field:"optional" json:"imageVersionNumber" yaml:"imageVersionNumber"`
}

