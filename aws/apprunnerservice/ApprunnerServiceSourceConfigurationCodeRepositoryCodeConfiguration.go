// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apprunnerservice


type ApprunnerServiceSourceConfigurationCodeRepositoryCodeConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.69.0/docs/resources/apprunner_service#configuration_source ApprunnerService#configuration_source}.
	ConfigurationSource *string `field:"required" json:"configurationSource" yaml:"configurationSource"`
	// code_configuration_values block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.69.0/docs/resources/apprunner_service#code_configuration_values ApprunnerService#code_configuration_values}
	CodeConfigurationValues *ApprunnerServiceSourceConfigurationCodeRepositoryCodeConfigurationCodeConfigurationValues `field:"optional" json:"codeConfigurationValues" yaml:"codeConfigurationValues"`
}

