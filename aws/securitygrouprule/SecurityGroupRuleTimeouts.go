// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package securitygrouprule


type SecurityGroupRuleTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/security_group_rule#create SecurityGroupRule#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
}

