// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package configorganizationconformancepack


type ConfigOrganizationConformancePackTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.0/docs/resources/config_organization_conformance_pack#create ConfigOrganizationConformancePack#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.0/docs/resources/config_organization_conformance_pack#delete ConfigOrganizationConformancePack#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.0/docs/resources/config_organization_conformance_pack#update ConfigOrganizationConformancePack#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

