// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package acmcertificate


type AcmCertificateValidationOption struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/resources/acm_certificate#domain_name AcmCertificate#domain_name}.
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/resources/acm_certificate#validation_domain AcmCertificate#validation_domain}.
	ValidationDomain *string `field:"required" json:"validationDomain" yaml:"validationDomain"`
}

