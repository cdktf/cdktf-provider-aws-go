// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package securityhuborganizationconfiguration


type SecurityhubOrganizationConfigurationTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.80.0/docs/resources/securityhub_organization_configuration#create SecurityhubOrganizationConfiguration#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.80.0/docs/resources/securityhub_organization_configuration#delete SecurityhubOrganizationConfiguration#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.80.0/docs/resources/securityhub_organization_configuration#update SecurityhubOrganizationConfiguration#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

