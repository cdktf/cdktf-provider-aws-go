// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cognitoriskconfiguration


type CognitoRiskConfigurationAccountTakeoverRiskConfiguration struct {
	// actions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.0/docs/resources/cognito_risk_configuration#actions CognitoRiskConfiguration#actions}
	Actions *CognitoRiskConfigurationAccountTakeoverRiskConfigurationActions `field:"required" json:"actions" yaml:"actions"`
	// notify_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.0/docs/resources/cognito_risk_configuration#notify_configuration CognitoRiskConfiguration#notify_configuration}
	NotifyConfiguration *CognitoRiskConfigurationAccountTakeoverRiskConfigurationNotifyConfiguration `field:"required" json:"notifyConfiguration" yaml:"notifyConfiguration"`
}

