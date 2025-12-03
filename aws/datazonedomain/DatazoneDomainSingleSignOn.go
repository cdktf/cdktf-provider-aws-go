// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datazonedomain


type DatazoneDomainSingleSignOn struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.24.0/docs/resources/datazone_domain#type DatazoneDomain#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.24.0/docs/resources/datazone_domain#user_assignment DatazoneDomain#user_assignment}.
	UserAssignment *string `field:"optional" json:"userAssignment" yaml:"userAssignment"`
}

