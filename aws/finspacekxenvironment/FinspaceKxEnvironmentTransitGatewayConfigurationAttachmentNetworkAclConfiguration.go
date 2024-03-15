// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package finspacekxenvironment


type FinspaceKxEnvironmentTransitGatewayConfigurationAttachmentNetworkAclConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.41.0/docs/resources/finspace_kx_environment#cidr_block FinspaceKxEnvironment#cidr_block}.
	CidrBlock *string `field:"required" json:"cidrBlock" yaml:"cidrBlock"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.41.0/docs/resources/finspace_kx_environment#protocol FinspaceKxEnvironment#protocol}.
	Protocol *string `field:"required" json:"protocol" yaml:"protocol"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.41.0/docs/resources/finspace_kx_environment#rule_action FinspaceKxEnvironment#rule_action}.
	RuleAction *string `field:"required" json:"ruleAction" yaml:"ruleAction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.41.0/docs/resources/finspace_kx_environment#rule_number FinspaceKxEnvironment#rule_number}.
	RuleNumber *float64 `field:"required" json:"ruleNumber" yaml:"ruleNumber"`
	// icmp_type_code block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.41.0/docs/resources/finspace_kx_environment#icmp_type_code FinspaceKxEnvironment#icmp_type_code}
	IcmpTypeCode *FinspaceKxEnvironmentTransitGatewayConfigurationAttachmentNetworkAclConfigurationIcmpTypeCode `field:"optional" json:"icmpTypeCode" yaml:"icmpTypeCode"`
	// port_range block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.41.0/docs/resources/finspace_kx_environment#port_range FinspaceKxEnvironment#port_range}
	PortRange *FinspaceKxEnvironmentTransitGatewayConfigurationAttachmentNetworkAclConfigurationPortRange `field:"optional" json:"portRange" yaml:"portRange"`
}

