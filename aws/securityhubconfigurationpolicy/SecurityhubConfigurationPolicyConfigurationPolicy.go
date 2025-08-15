// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package securityhubconfigurationpolicy


type SecurityhubConfigurationPolicyConfigurationPolicy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.9.0/docs/resources/securityhub_configuration_policy#service_enabled SecurityhubConfigurationPolicy#service_enabled}.
	ServiceEnabled interface{} `field:"required" json:"serviceEnabled" yaml:"serviceEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.9.0/docs/resources/securityhub_configuration_policy#enabled_standard_arns SecurityhubConfigurationPolicy#enabled_standard_arns}.
	EnabledStandardArns *[]*string `field:"optional" json:"enabledStandardArns" yaml:"enabledStandardArns"`
	// security_controls_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.9.0/docs/resources/securityhub_configuration_policy#security_controls_configuration SecurityhubConfigurationPolicy#security_controls_configuration}
	SecurityControlsConfiguration *SecurityhubConfigurationPolicyConfigurationPolicySecurityControlsConfiguration `field:"optional" json:"securityControlsConfiguration" yaml:"securityControlsConfiguration"`
}

