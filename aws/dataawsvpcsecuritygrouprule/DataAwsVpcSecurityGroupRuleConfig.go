package dataawsvpcsecuritygrouprule

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataAwsVpcSecurityGroupRuleConfig struct {
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
	// filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.0/docs/data-sources/vpc_security_group_rule#filter DataAwsVpcSecurityGroupRule#filter}
	Filter interface{} `field:"optional" json:"filter" yaml:"filter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.0/docs/data-sources/vpc_security_group_rule#security_group_rule_id DataAwsVpcSecurityGroupRule#security_group_rule_id}.
	SecurityGroupRuleId *string `field:"optional" json:"securityGroupRuleId" yaml:"securityGroupRuleId"`
}

