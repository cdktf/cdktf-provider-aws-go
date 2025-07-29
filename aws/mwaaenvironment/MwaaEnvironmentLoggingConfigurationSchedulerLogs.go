// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mwaaenvironment


type MwaaEnvironmentLoggingConfigurationSchedulerLogs struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.6.0/docs/resources/mwaa_environment#enabled MwaaEnvironment#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.6.0/docs/resources/mwaa_environment#log_level MwaaEnvironment#log_level}.
	LogLevel *string `field:"optional" json:"logLevel" yaml:"logLevel"`
}

