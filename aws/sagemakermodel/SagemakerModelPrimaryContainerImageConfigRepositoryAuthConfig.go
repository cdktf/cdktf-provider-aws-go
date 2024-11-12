// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sagemakermodel


type SagemakerModelPrimaryContainerImageConfigRepositoryAuthConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.1/docs/resources/sagemaker_model#repository_credentials_provider_arn SagemakerModel#repository_credentials_provider_arn}.
	RepositoryCredentialsProviderArn *string `field:"required" json:"repositoryCredentialsProviderArn" yaml:"repositoryCredentialsProviderArn"`
}

