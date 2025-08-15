// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apprunnerservice


type ApprunnerServiceSourceConfigurationImageRepositoryImageConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.9.0/docs/resources/apprunner_service#port ApprunnerService#port}.
	Port *string `field:"optional" json:"port" yaml:"port"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.9.0/docs/resources/apprunner_service#runtime_environment_secrets ApprunnerService#runtime_environment_secrets}.
	RuntimeEnvironmentSecrets *map[string]*string `field:"optional" json:"runtimeEnvironmentSecrets" yaml:"runtimeEnvironmentSecrets"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.9.0/docs/resources/apprunner_service#runtime_environment_variables ApprunnerService#runtime_environment_variables}.
	RuntimeEnvironmentVariables *map[string]*string `field:"optional" json:"runtimeEnvironmentVariables" yaml:"runtimeEnvironmentVariables"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.9.0/docs/resources/apprunner_service#start_command ApprunnerService#start_command}.
	StartCommand *string `field:"optional" json:"startCommand" yaml:"startCommand"`
}

