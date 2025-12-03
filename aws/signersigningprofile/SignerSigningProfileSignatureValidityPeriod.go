// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package signersigningprofile


type SignerSigningProfileSignatureValidityPeriod struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.24.0/docs/resources/signer_signing_profile#type SignerSigningProfile#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.24.0/docs/resources/signer_signing_profile#value SignerSigningProfile#value}.
	Value *float64 `field:"required" json:"value" yaml:"value"`
}

