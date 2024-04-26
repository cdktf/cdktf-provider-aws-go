// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package acmpcacertificateauthority


type AcmpcaCertificateAuthorityRevocationConfigurationOcspConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.47.0/docs/resources/acmpca_certificate_authority#enabled AcmpcaCertificateAuthority#enabled}.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.47.0/docs/resources/acmpca_certificate_authority#ocsp_custom_cname AcmpcaCertificateAuthority#ocsp_custom_cname}.
	OcspCustomCname *string `field:"optional" json:"ocspCustomCname" yaml:"ocspCustomCname"`
}

