package vpc


type DataAwsVpcPeeringConnectionsFilter struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/vpc_peering_connections#name DataAwsVpcPeeringConnections#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/vpc_peering_connections#values DataAwsVpcPeeringConnections#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

