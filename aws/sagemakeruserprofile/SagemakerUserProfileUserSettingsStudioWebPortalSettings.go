// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sagemakeruserprofile


type SagemakerUserProfileUserSettingsStudioWebPortalSettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.17.0/docs/resources/sagemaker_user_profile#hidden_app_types SagemakerUserProfile#hidden_app_types}.
	HiddenAppTypes *[]*string `field:"optional" json:"hiddenAppTypes" yaml:"hiddenAppTypes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.17.0/docs/resources/sagemaker_user_profile#hidden_instance_types SagemakerUserProfile#hidden_instance_types}.
	HiddenInstanceTypes *[]*string `field:"optional" json:"hiddenInstanceTypes" yaml:"hiddenInstanceTypes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.17.0/docs/resources/sagemaker_user_profile#hidden_ml_tools SagemakerUserProfile#hidden_ml_tools}.
	HiddenMlTools *[]*string `field:"optional" json:"hiddenMlTools" yaml:"hiddenMlTools"`
}

