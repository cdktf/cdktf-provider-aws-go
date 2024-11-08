// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sagemakerdomain


type SagemakerDomainDomainSettingsDockerSettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.0/docs/resources/sagemaker_domain#enable_docker_access SagemakerDomain#enable_docker_access}.
	EnableDockerAccess *string `field:"optional" json:"enableDockerAccess" yaml:"enableDockerAccess"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.0/docs/resources/sagemaker_domain#vpc_only_trusted_accounts SagemakerDomain#vpc_only_trusted_accounts}.
	VpcOnlyTrustedAccounts *[]*string `field:"optional" json:"vpcOnlyTrustedAccounts" yaml:"vpcOnlyTrustedAccounts"`
}

