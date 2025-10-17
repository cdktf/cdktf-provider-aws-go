// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package securityhubconfigurationpolicy


type SecurityhubConfigurationPolicyConfigurationPolicySecurityControlsConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.17.0/docs/resources/securityhub_configuration_policy#disabled_control_identifiers SecurityhubConfigurationPolicy#disabled_control_identifiers}.
	DisabledControlIdentifiers *[]*string `field:"optional" json:"disabledControlIdentifiers" yaml:"disabledControlIdentifiers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.17.0/docs/resources/securityhub_configuration_policy#enabled_control_identifiers SecurityhubConfigurationPolicy#enabled_control_identifiers}.
	EnabledControlIdentifiers *[]*string `field:"optional" json:"enabledControlIdentifiers" yaml:"enabledControlIdentifiers"`
	// security_control_custom_parameter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.17.0/docs/resources/securityhub_configuration_policy#security_control_custom_parameter SecurityhubConfigurationPolicy#security_control_custom_parameter}
	SecurityControlCustomParameter interface{} `field:"optional" json:"securityControlCustomParameter" yaml:"securityControlCustomParameter"`
}

