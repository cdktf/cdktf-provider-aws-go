// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apprunnerservice


type ApprunnerServiceSourceConfiguration struct {
	// authentication_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/apprunner_service#authentication_configuration ApprunnerService#authentication_configuration}
	AuthenticationConfiguration *ApprunnerServiceSourceConfigurationAuthenticationConfiguration `field:"optional" json:"authenticationConfiguration" yaml:"authenticationConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/apprunner_service#auto_deployments_enabled ApprunnerService#auto_deployments_enabled}.
	AutoDeploymentsEnabled interface{} `field:"optional" json:"autoDeploymentsEnabled" yaml:"autoDeploymentsEnabled"`
	// code_repository block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/apprunner_service#code_repository ApprunnerService#code_repository}
	CodeRepository *ApprunnerServiceSourceConfigurationCodeRepository `field:"optional" json:"codeRepository" yaml:"codeRepository"`
	// image_repository block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/apprunner_service#image_repository ApprunnerService#image_repository}
	ImageRepository *ApprunnerServiceSourceConfigurationImageRepository `field:"optional" json:"imageRepository" yaml:"imageRepository"`
}

