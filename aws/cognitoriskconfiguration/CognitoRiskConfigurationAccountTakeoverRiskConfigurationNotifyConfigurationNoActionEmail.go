// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cognitoriskconfiguration


type CognitoRiskConfigurationAccountTakeoverRiskConfigurationNotifyConfigurationNoActionEmail struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/resources/cognito_risk_configuration#html_body CognitoRiskConfiguration#html_body}.
	HtmlBody *string `field:"required" json:"htmlBody" yaml:"htmlBody"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/resources/cognito_risk_configuration#subject CognitoRiskConfiguration#subject}.
	Subject *string `field:"required" json:"subject" yaml:"subject"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/resources/cognito_risk_configuration#text_body CognitoRiskConfiguration#text_body}.
	TextBody *string `field:"required" json:"textBody" yaml:"textBody"`
}

