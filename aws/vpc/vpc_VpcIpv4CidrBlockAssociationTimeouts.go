package vpc


type VpcIpv4CidrBlockAssociationTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/vpc_ipv4_cidr_block_association#create VpcIpv4CidrBlockAssociation#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/vpc_ipv4_cidr_block_association#delete VpcIpv4CidrBlockAssociation#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

