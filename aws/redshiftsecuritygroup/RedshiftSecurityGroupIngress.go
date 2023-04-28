package redshiftsecuritygroup


type RedshiftSecurityGroupIngress struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.65.0/docs/resources/redshift_security_group#cidr RedshiftSecurityGroup#cidr}.
	Cidr *string `field:"optional" json:"cidr" yaml:"cidr"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.65.0/docs/resources/redshift_security_group#security_group_name RedshiftSecurityGroup#security_group_name}.
	SecurityGroupName *string `field:"optional" json:"securityGroupName" yaml:"securityGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.65.0/docs/resources/redshift_security_group#security_group_owner_id RedshiftSecurityGroup#security_group_owner_id}.
	SecurityGroupOwnerId *string `field:"optional" json:"securityGroupOwnerId" yaml:"securityGroupOwnerId"`
}

