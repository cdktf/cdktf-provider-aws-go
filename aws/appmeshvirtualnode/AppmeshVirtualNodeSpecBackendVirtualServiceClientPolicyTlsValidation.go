// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appmeshvirtualnode


type AppmeshVirtualNodeSpecBackendVirtualServiceClientPolicyTlsValidation struct {
	// trust block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.89.0/docs/resources/appmesh_virtual_node#trust AppmeshVirtualNode#trust}
	Trust *AppmeshVirtualNodeSpecBackendVirtualServiceClientPolicyTlsValidationTrust `field:"required" json:"trust" yaml:"trust"`
	// subject_alternative_names block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.89.0/docs/resources/appmesh_virtual_node#subject_alternative_names AppmeshVirtualNode#subject_alternative_names}
	SubjectAlternativeNames *AppmeshVirtualNodeSpecBackendVirtualServiceClientPolicyTlsValidationSubjectAlternativeNames `field:"optional" json:"subjectAlternativeNames" yaml:"subjectAlternativeNames"`
}

