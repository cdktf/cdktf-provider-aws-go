// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sagemakeruserprofile


type SagemakerUserProfileUserSettingsJupyterLabAppSettingsEmrSettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.5.0/docs/resources/sagemaker_user_profile#assumable_role_arns SagemakerUserProfile#assumable_role_arns}.
	AssumableRoleArns *[]*string `field:"optional" json:"assumableRoleArns" yaml:"assumableRoleArns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.5.0/docs/resources/sagemaker_user_profile#execution_role_arns SagemakerUserProfile#execution_role_arns}.
	ExecutionRoleArns *[]*string `field:"optional" json:"executionRoleArns" yaml:"executionRoleArns"`
}

