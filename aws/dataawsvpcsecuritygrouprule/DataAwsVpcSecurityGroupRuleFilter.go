// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawsvpcsecuritygrouprule


type DataAwsVpcSecurityGroupRuleFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/data-sources/vpc_security_group_rule#name DataAwsVpcSecurityGroupRule#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/data-sources/vpc_security_group_rule#values DataAwsVpcSecurityGroupRule#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

