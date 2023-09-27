// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package inspector2organizationconfiguration


type Inspector2OrganizationConfigurationTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/resources/inspector2_organization_configuration#create Inspector2OrganizationConfiguration#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/resources/inspector2_organization_configuration#delete Inspector2OrganizationConfiguration#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/resources/inspector2_organization_configuration#update Inspector2OrganizationConfiguration#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

