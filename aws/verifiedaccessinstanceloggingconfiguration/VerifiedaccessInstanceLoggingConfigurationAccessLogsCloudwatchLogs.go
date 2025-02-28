// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package verifiedaccessinstanceloggingconfiguration


type VerifiedaccessInstanceLoggingConfigurationAccessLogsCloudwatchLogs struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.89.0/docs/resources/verifiedaccess_instance_logging_configuration#enabled VerifiedaccessInstanceLoggingConfiguration#enabled}.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.89.0/docs/resources/verifiedaccess_instance_logging_configuration#log_group VerifiedaccessInstanceLoggingConfiguration#log_group}.
	LogGroup *string `field:"optional" json:"logGroup" yaml:"logGroup"`
}

