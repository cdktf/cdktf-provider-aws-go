// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ssoadmintrustedtokenissuer


type SsoadminTrustedTokenIssuerTrustedTokenIssuerConfiguration struct {
	// oidc_jwt_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.1/docs/resources/ssoadmin_trusted_token_issuer#oidc_jwt_configuration SsoadminTrustedTokenIssuer#oidc_jwt_configuration}
	OidcJwtConfiguration interface{} `field:"optional" json:"oidcJwtConfiguration" yaml:"oidcJwtConfiguration"`
}

