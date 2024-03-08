// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawsvpcsecuritygrouprules


type DataAwsVpcSecurityGroupRulesFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.40.0/docs/data-sources/vpc_security_group_rules#name DataAwsVpcSecurityGroupRules#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.40.0/docs/data-sources/vpc_security_group_rules#values DataAwsVpcSecurityGroupRules#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

