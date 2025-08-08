// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package acmpcacertificateauthority


type AcmpcaCertificateAuthorityRevocationConfiguration struct {
	// crl_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/acmpca_certificate_authority#crl_configuration AcmpcaCertificateAuthority#crl_configuration}
	CrlConfiguration *AcmpcaCertificateAuthorityRevocationConfigurationCrlConfiguration `field:"optional" json:"crlConfiguration" yaml:"crlConfiguration"`
	// ocsp_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/acmpca_certificate_authority#ocsp_configuration AcmpcaCertificateAuthority#ocsp_configuration}
	OcspConfiguration *AcmpcaCertificateAuthorityRevocationConfigurationOcspConfiguration `field:"optional" json:"ocspConfiguration" yaml:"ocspConfiguration"`
}

