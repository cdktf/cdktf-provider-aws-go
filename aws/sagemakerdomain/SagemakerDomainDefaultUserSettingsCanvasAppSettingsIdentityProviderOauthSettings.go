// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sagemakerdomain


type SagemakerDomainDefaultUserSettingsCanvasAppSettingsIdentityProviderOauthSettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.76.0/docs/resources/sagemaker_domain#secret_arn SagemakerDomain#secret_arn}.
	SecretArn *string `field:"required" json:"secretArn" yaml:"secretArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.76.0/docs/resources/sagemaker_domain#data_source_name SagemakerDomain#data_source_name}.
	DataSourceName *string `field:"optional" json:"dataSourceName" yaml:"dataSourceName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.76.0/docs/resources/sagemaker_domain#status SagemakerDomain#status}.
	Status *string `field:"optional" json:"status" yaml:"status"`
}

