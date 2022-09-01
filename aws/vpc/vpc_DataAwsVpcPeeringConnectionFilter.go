package vpc


type DataAwsVpcPeeringConnectionFilter struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/vpc_peering_connection#name DataAwsVpcPeeringConnection#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/vpc_peering_connection#values DataAwsVpcPeeringConnection#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

