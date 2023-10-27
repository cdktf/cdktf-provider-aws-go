// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vpcsecuritygroupegressrule

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VpcSecurityGroupEgressRuleConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.0/docs/resources/vpc_security_group_egress_rule#ip_protocol VpcSecurityGroupEgressRule#ip_protocol}.
	IpProtocol *string `field:"required" json:"ipProtocol" yaml:"ipProtocol"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.0/docs/resources/vpc_security_group_egress_rule#security_group_id VpcSecurityGroupEgressRule#security_group_id}.
	SecurityGroupId *string `field:"required" json:"securityGroupId" yaml:"securityGroupId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.0/docs/resources/vpc_security_group_egress_rule#cidr_ipv4 VpcSecurityGroupEgressRule#cidr_ipv4}.
	CidrIpv4 *string `field:"optional" json:"cidrIpv4" yaml:"cidrIpv4"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.0/docs/resources/vpc_security_group_egress_rule#cidr_ipv6 VpcSecurityGroupEgressRule#cidr_ipv6}.
	CidrIpv6 *string `field:"optional" json:"cidrIpv6" yaml:"cidrIpv6"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.0/docs/resources/vpc_security_group_egress_rule#description VpcSecurityGroupEgressRule#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.0/docs/resources/vpc_security_group_egress_rule#from_port VpcSecurityGroupEgressRule#from_port}.
	FromPort *float64 `field:"optional" json:"fromPort" yaml:"fromPort"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.0/docs/resources/vpc_security_group_egress_rule#prefix_list_id VpcSecurityGroupEgressRule#prefix_list_id}.
	PrefixListId *string `field:"optional" json:"prefixListId" yaml:"prefixListId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.0/docs/resources/vpc_security_group_egress_rule#referenced_security_group_id VpcSecurityGroupEgressRule#referenced_security_group_id}.
	ReferencedSecurityGroupId *string `field:"optional" json:"referencedSecurityGroupId" yaml:"referencedSecurityGroupId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.0/docs/resources/vpc_security_group_egress_rule#tags VpcSecurityGroupEgressRule#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.0/docs/resources/vpc_security_group_egress_rule#to_port VpcSecurityGroupEgressRule#to_port}.
	ToPort *float64 `field:"optional" json:"toPort" yaml:"toPort"`
}

