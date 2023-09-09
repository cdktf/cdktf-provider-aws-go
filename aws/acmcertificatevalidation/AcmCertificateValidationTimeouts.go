// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package acmcertificatevalidation


type AcmCertificateValidationTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.16.1/docs/resources/acm_certificate_validation#create AcmCertificateValidation#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
}

