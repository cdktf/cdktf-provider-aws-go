// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cognitouserpool


type CognitoUserPoolDeviceConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/cognito_user_pool#challenge_required_on_new_device CognitoUserPool#challenge_required_on_new_device}.
	ChallengeRequiredOnNewDevice interface{} `field:"optional" json:"challengeRequiredOnNewDevice" yaml:"challengeRequiredOnNewDevice"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/cognito_user_pool#device_only_remembered_on_user_prompt CognitoUserPool#device_only_remembered_on_user_prompt}.
	DeviceOnlyRememberedOnUserPrompt interface{} `field:"optional" json:"deviceOnlyRememberedOnUserPrompt" yaml:"deviceOnlyRememberedOnUserPrompt"`
}

