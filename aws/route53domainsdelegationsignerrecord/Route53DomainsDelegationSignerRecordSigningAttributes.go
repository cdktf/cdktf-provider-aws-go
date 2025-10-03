// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package route53domainsdelegationsignerrecord


type Route53DomainsDelegationSignerRecordSigningAttributes struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/route53domains_delegation_signer_record#algorithm Route53DomainsDelegationSignerRecord#algorithm}.
	Algorithm *float64 `field:"required" json:"algorithm" yaml:"algorithm"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/route53domains_delegation_signer_record#flags Route53DomainsDelegationSignerRecord#flags}.
	Flags *float64 `field:"required" json:"flags" yaml:"flags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/route53domains_delegation_signer_record#public_key Route53DomainsDelegationSignerRecord#public_key}.
	PublicKey *string `field:"required" json:"publicKey" yaml:"publicKey"`
}

