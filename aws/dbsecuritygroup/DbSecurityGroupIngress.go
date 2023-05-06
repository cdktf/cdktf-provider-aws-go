package dbsecuritygroup


type DbSecurityGroupIngress struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.1/docs/resources/db_security_group#cidr DbSecurityGroup#cidr}.
	Cidr *string `field:"optional" json:"cidr" yaml:"cidr"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.1/docs/resources/db_security_group#security_group_id DbSecurityGroup#security_group_id}.
	SecurityGroupId *string `field:"optional" json:"securityGroupId" yaml:"securityGroupId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.1/docs/resources/db_security_group#security_group_name DbSecurityGroup#security_group_name}.
	SecurityGroupName *string `field:"optional" json:"securityGroupName" yaml:"securityGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.1/docs/resources/db_security_group#security_group_owner_id DbSecurityGroup#security_group_owner_id}.
	SecurityGroupOwnerId *string `field:"optional" json:"securityGroupOwnerId" yaml:"securityGroupOwnerId"`
}

