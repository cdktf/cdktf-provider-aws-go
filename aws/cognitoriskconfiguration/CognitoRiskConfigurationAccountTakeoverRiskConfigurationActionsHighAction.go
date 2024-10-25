// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cognitoriskconfiguration


type CognitoRiskConfigurationAccountTakeoverRiskConfigurationActionsHighAction struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.73.0/docs/resources/cognito_risk_configuration#event_action CognitoRiskConfiguration#event_action}.
	EventAction *string `field:"required" json:"eventAction" yaml:"eventAction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.73.0/docs/resources/cognito_risk_configuration#notify CognitoRiskConfiguration#notify}.
	Notify interface{} `field:"required" json:"notify" yaml:"notify"`
}

