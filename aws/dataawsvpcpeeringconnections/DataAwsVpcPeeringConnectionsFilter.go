package dataawsvpcpeeringconnections


type DataAwsVpcPeeringConnectionsFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.10.0/docs/data-sources/vpc_peering_connections#name DataAwsVpcPeeringConnections#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.10.0/docs/data-sources/vpc_peering_connections#values DataAwsVpcPeeringConnections#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

