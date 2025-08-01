// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cognitoriskconfiguration


type CognitoRiskConfigurationAccountTakeoverRiskConfigurationNotifyConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/cognito_risk_configuration#source_arn CognitoRiskConfiguration#source_arn}.
	SourceArn *string `field:"required" json:"sourceArn" yaml:"sourceArn"`
	// block_email block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/cognito_risk_configuration#block_email CognitoRiskConfiguration#block_email}
	BlockEmail *CognitoRiskConfigurationAccountTakeoverRiskConfigurationNotifyConfigurationBlockEmail `field:"optional" json:"blockEmail" yaml:"blockEmail"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/cognito_risk_configuration#from CognitoRiskConfiguration#from}.
	From *string `field:"optional" json:"from" yaml:"from"`
	// mfa_email block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/cognito_risk_configuration#mfa_email CognitoRiskConfiguration#mfa_email}
	MfaEmail *CognitoRiskConfigurationAccountTakeoverRiskConfigurationNotifyConfigurationMfaEmail `field:"optional" json:"mfaEmail" yaml:"mfaEmail"`
	// no_action_email block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/cognito_risk_configuration#no_action_email CognitoRiskConfiguration#no_action_email}
	NoActionEmail *CognitoRiskConfigurationAccountTakeoverRiskConfigurationNotifyConfigurationNoActionEmail `field:"optional" json:"noActionEmail" yaml:"noActionEmail"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/cognito_risk_configuration#reply_to CognitoRiskConfiguration#reply_to}.
	ReplyTo *string `field:"optional" json:"replyTo" yaml:"replyTo"`
}

