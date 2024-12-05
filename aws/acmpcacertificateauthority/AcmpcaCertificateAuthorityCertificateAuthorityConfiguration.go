// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package acmpcacertificateauthority


type AcmpcaCertificateAuthorityCertificateAuthorityConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.80.0/docs/resources/acmpca_certificate_authority#key_algorithm AcmpcaCertificateAuthority#key_algorithm}.
	KeyAlgorithm *string `field:"required" json:"keyAlgorithm" yaml:"keyAlgorithm"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.80.0/docs/resources/acmpca_certificate_authority#signing_algorithm AcmpcaCertificateAuthority#signing_algorithm}.
	SigningAlgorithm *string `field:"required" json:"signingAlgorithm" yaml:"signingAlgorithm"`
	// subject block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.80.0/docs/resources/acmpca_certificate_authority#subject AcmpcaCertificateAuthority#subject}
	Subject *AcmpcaCertificateAuthorityCertificateAuthorityConfigurationSubject `field:"required" json:"subject" yaml:"subject"`
}

