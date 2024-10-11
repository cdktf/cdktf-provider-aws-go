// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sagemakercoderepository


type SagemakerCodeRepositoryGitConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.71.0/docs/resources/sagemaker_code_repository#repository_url SagemakerCodeRepository#repository_url}.
	RepositoryUrl *string `field:"required" json:"repositoryUrl" yaml:"repositoryUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.71.0/docs/resources/sagemaker_code_repository#branch SagemakerCodeRepository#branch}.
	Branch *string `field:"optional" json:"branch" yaml:"branch"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.71.0/docs/resources/sagemaker_code_repository#secret_arn SagemakerCodeRepository#secret_arn}.
	SecretArn *string `field:"optional" json:"secretArn" yaml:"secretArn"`
}

