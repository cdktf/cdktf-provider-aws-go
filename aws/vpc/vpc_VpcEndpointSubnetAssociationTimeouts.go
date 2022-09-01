package vpc


type VpcEndpointSubnetAssociationTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/vpc_endpoint_subnet_association#create VpcEndpointSubnetAssociation#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/vpc_endpoint_subnet_association#delete VpcEndpointSubnetAssociation#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

