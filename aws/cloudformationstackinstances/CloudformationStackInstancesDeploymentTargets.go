// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudformationstackinstances


type CloudformationStackInstancesDeploymentTargets struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.3.0/docs/resources/cloudformation_stack_instances#account_filter_type CloudformationStackInstances#account_filter_type}.
	AccountFilterType *string `field:"optional" json:"accountFilterType" yaml:"accountFilterType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.3.0/docs/resources/cloudformation_stack_instances#accounts CloudformationStackInstances#accounts}.
	Accounts *[]*string `field:"optional" json:"accounts" yaml:"accounts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.3.0/docs/resources/cloudformation_stack_instances#accounts_url CloudformationStackInstances#accounts_url}.
	AccountsUrl *string `field:"optional" json:"accountsUrl" yaml:"accountsUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.3.0/docs/resources/cloudformation_stack_instances#organizational_unit_ids CloudformationStackInstances#organizational_unit_ids}.
	OrganizationalUnitIds *[]*string `field:"optional" json:"organizationalUnitIds" yaml:"organizationalUnitIds"`
}

