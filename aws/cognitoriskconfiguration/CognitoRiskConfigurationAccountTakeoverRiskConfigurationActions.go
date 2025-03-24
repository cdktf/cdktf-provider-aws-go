// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cognitoriskconfiguration


type CognitoRiskConfigurationAccountTakeoverRiskConfigurationActions struct {
	// high_action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.92.0/docs/resources/cognito_risk_configuration#high_action CognitoRiskConfiguration#high_action}
	HighAction *CognitoRiskConfigurationAccountTakeoverRiskConfigurationActionsHighAction `field:"optional" json:"highAction" yaml:"highAction"`
	// low_action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.92.0/docs/resources/cognito_risk_configuration#low_action CognitoRiskConfiguration#low_action}
	LowAction *CognitoRiskConfigurationAccountTakeoverRiskConfigurationActionsLowAction `field:"optional" json:"lowAction" yaml:"lowAction"`
	// medium_action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.92.0/docs/resources/cognito_risk_configuration#medium_action CognitoRiskConfiguration#medium_action}
	MediumAction *CognitoRiskConfigurationAccountTakeoverRiskConfigurationActionsMediumAction `field:"optional" json:"mediumAction" yaml:"mediumAction"`
}

