// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cognitouserpool


type CognitoUserPoolUserPoolAddOns struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.9.0/docs/resources/cognito_user_pool#advanced_security_mode CognitoUserPool#advanced_security_mode}.
	AdvancedSecurityMode *string `field:"required" json:"advancedSecurityMode" yaml:"advancedSecurityMode"`
	// advanced_security_additional_flows block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.9.0/docs/resources/cognito_user_pool#advanced_security_additional_flows CognitoUserPool#advanced_security_additional_flows}
	AdvancedSecurityAdditionalFlows *CognitoUserPoolUserPoolAddOnsAdvancedSecurityAdditionalFlows `field:"optional" json:"advancedSecurityAdditionalFlows" yaml:"advancedSecurityAdditionalFlows"`
}

