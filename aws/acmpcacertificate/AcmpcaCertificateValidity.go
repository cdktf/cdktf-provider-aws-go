// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package acmpcacertificate


type AcmpcaCertificateValidity struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.26.0/docs/resources/acmpca_certificate#type AcmpcaCertificate#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.26.0/docs/resources/acmpca_certificate#value AcmpcaCertificate#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

