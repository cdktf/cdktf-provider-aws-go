// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudformationstacksetinstance


type CloudformationStackSetInstanceDeploymentTargets struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.6.0/docs/resources/cloudformation_stack_set_instance#account_filter_type CloudformationStackSetInstance#account_filter_type}.
	AccountFilterType *string `field:"optional" json:"accountFilterType" yaml:"accountFilterType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.6.0/docs/resources/cloudformation_stack_set_instance#accounts CloudformationStackSetInstance#accounts}.
	Accounts *[]*string `field:"optional" json:"accounts" yaml:"accounts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.6.0/docs/resources/cloudformation_stack_set_instance#accounts_url CloudformationStackSetInstance#accounts_url}.
	AccountsUrl *string `field:"optional" json:"accountsUrl" yaml:"accountsUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.6.0/docs/resources/cloudformation_stack_set_instance#organizational_unit_ids CloudformationStackSetInstance#organizational_unit_ids}.
	OrganizationalUnitIds *[]*string `field:"optional" json:"organizationalUnitIds" yaml:"organizationalUnitIds"`
}

