// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package signersigningprofile


type SignerSigningProfileSigningMaterial struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/signer_signing_profile#certificate_arn SignerSigningProfile#certificate_arn}.
	CertificateArn *string `field:"required" json:"certificateArn" yaml:"certificateArn"`
}

