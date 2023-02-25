package dataawsvpcsecuritygrouprules


type DataAwsVpcSecurityGroupRulesFilter struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/vpc_security_group_rules#name DataAwsVpcSecurityGroupRules#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/vpc_security_group_rules#values DataAwsVpcSecurityGroupRules#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

