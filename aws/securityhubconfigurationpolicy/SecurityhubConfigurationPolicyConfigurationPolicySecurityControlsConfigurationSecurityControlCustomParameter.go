// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package securityhubconfigurationpolicy


type SecurityhubConfigurationPolicyConfigurationPolicySecurityControlsConfigurationSecurityControlCustomParameter struct {
	// parameter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/securityhub_configuration_policy#parameter SecurityhubConfigurationPolicy#parameter}
	Parameter interface{} `field:"required" json:"parameter" yaml:"parameter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/securityhub_configuration_policy#security_control_id SecurityhubConfigurationPolicy#security_control_id}.
	SecurityControlId *string `field:"required" json:"securityControlId" yaml:"securityControlId"`
}

