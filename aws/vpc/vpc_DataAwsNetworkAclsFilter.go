package vpc


type DataAwsNetworkAclsFilter struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/network_acls#name DataAwsNetworkAcls#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/network_acls#values DataAwsNetworkAcls#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

