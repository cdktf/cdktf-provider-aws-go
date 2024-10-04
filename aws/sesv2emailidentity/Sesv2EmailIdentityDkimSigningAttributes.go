// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sesv2emailidentity


type Sesv2EmailIdentityDkimSigningAttributes struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.70.0/docs/resources/sesv2_email_identity#domain_signing_private_key Sesv2EmailIdentity#domain_signing_private_key}.
	DomainSigningPrivateKey *string `field:"optional" json:"domainSigningPrivateKey" yaml:"domainSigningPrivateKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.70.0/docs/resources/sesv2_email_identity#domain_signing_selector Sesv2EmailIdentity#domain_signing_selector}.
	DomainSigningSelector *string `field:"optional" json:"domainSigningSelector" yaml:"domainSigningSelector"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.70.0/docs/resources/sesv2_email_identity#next_signing_key_length Sesv2EmailIdentity#next_signing_key_length}.
	NextSigningKeyLength *string `field:"optional" json:"nextSigningKeyLength" yaml:"nextSigningKeyLength"`
}

