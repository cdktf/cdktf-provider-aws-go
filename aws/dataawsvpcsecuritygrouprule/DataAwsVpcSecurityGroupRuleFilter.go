package dataawsvpcsecuritygrouprule


type DataAwsVpcSecurityGroupRuleFilter struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/vpc_security_group_rule#name DataAwsVpcSecurityGroupRule#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/vpc_security_group_rule#values DataAwsVpcSecurityGroupRule#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

