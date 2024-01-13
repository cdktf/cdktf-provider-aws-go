// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appmeshvirtualgateway


type AppmeshVirtualGatewaySpecListenerTlsValidation struct {
	// trust block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.32.1/docs/resources/appmesh_virtual_gateway#trust AppmeshVirtualGateway#trust}
	Trust *AppmeshVirtualGatewaySpecListenerTlsValidationTrust `field:"required" json:"trust" yaml:"trust"`
	// subject_alternative_names block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.32.1/docs/resources/appmesh_virtual_gateway#subject_alternative_names AppmeshVirtualGateway#subject_alternative_names}
	SubjectAlternativeNames *AppmeshVirtualGatewaySpecListenerTlsValidationSubjectAlternativeNames `field:"optional" json:"subjectAlternativeNames" yaml:"subjectAlternativeNames"`
}

