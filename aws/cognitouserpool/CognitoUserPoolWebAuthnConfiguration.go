// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cognitouserpool


type CognitoUserPoolWebAuthnConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.17.0/docs/resources/cognito_user_pool#relying_party_id CognitoUserPool#relying_party_id}.
	RelyingPartyId *string `field:"optional" json:"relyingPartyId" yaml:"relyingPartyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.17.0/docs/resources/cognito_user_pool#user_verification CognitoUserPool#user_verification}.
	UserVerification *string `field:"optional" json:"userVerification" yaml:"userVerification"`
}

