// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sagemakeruserprofile


type SagemakerUserProfileUserSettingsCanvasAppSettingsIdentityProviderOauthSettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.12.0/docs/resources/sagemaker_user_profile#secret_arn SagemakerUserProfile#secret_arn}.
	SecretArn *string `field:"required" json:"secretArn" yaml:"secretArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.12.0/docs/resources/sagemaker_user_profile#data_source_name SagemakerUserProfile#data_source_name}.
	DataSourceName *string `field:"optional" json:"dataSourceName" yaml:"dataSourceName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.12.0/docs/resources/sagemaker_user_profile#status SagemakerUserProfile#status}.
	Status *string `field:"optional" json:"status" yaml:"status"`
}

