// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package alblistener


type AlbListenerDefaultActionJwtValidation struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/alb_listener#issuer AlbListener#issuer}.
	Issuer *string `field:"required" json:"issuer" yaml:"issuer"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/alb_listener#jwks_endpoint AlbListener#jwks_endpoint}.
	JwksEndpoint *string `field:"required" json:"jwksEndpoint" yaml:"jwksEndpoint"`
	// additional_claim block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/alb_listener#additional_claim AlbListener#additional_claim}
	AdditionalClaim interface{} `field:"optional" json:"additionalClaim" yaml:"additionalClaim"`
}

