// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cognitouserpool


type CognitoUserPoolAccountRecoverySetting struct {
	// recovery_mechanism block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.51.0/docs/resources/cognito_user_pool#recovery_mechanism CognitoUserPool#recovery_mechanism}
	RecoveryMechanism interface{} `field:"optional" json:"recoveryMechanism" yaml:"recoveryMechanism"`
}

