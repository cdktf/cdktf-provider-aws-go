// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package inspector2organizationconfiguration


type Inspector2OrganizationConfigurationAutoEnable struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/inspector2_organization_configuration#ec2 Inspector2OrganizationConfiguration#ec2}.
	Ec2 interface{} `field:"required" json:"ec2" yaml:"ec2"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/inspector2_organization_configuration#ecr Inspector2OrganizationConfiguration#ecr}.
	Ecr interface{} `field:"required" json:"ecr" yaml:"ecr"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/inspector2_organization_configuration#code_repository Inspector2OrganizationConfiguration#code_repository}.
	CodeRepository interface{} `field:"optional" json:"codeRepository" yaml:"codeRepository"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/inspector2_organization_configuration#lambda Inspector2OrganizationConfiguration#lambda}.
	Lambda interface{} `field:"optional" json:"lambda" yaml:"lambda"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/inspector2_organization_configuration#lambda_code Inspector2OrganizationConfiguration#lambda_code}.
	LambdaCode interface{} `field:"optional" json:"lambdaCode" yaml:"lambdaCode"`
}

