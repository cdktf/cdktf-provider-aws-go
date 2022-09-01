package vpc


type DataAwsSecurityGroupsFilter struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/security_groups#name DataAwsSecurityGroups#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/security_groups#values DataAwsSecurityGroups#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

