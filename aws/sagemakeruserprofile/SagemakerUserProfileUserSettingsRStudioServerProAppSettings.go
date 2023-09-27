// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sagemakeruserprofile


type SagemakerUserProfileUserSettingsRStudioServerProAppSettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/resources/sagemaker_user_profile#access_status SagemakerUserProfile#access_status}.
	AccessStatus *string `field:"optional" json:"accessStatus" yaml:"accessStatus"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/resources/sagemaker_user_profile#user_group SagemakerUserProfile#user_group}.
	UserGroup *string `field:"optional" json:"userGroup" yaml:"userGroup"`
}

