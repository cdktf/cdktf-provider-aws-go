package dataawsvpcpeeringconnection


type DataAwsVpcPeeringConnectionFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/data-sources/vpc_peering_connection#name DataAwsVpcPeeringConnection#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/data-sources/vpc_peering_connection#values DataAwsVpcPeeringConnection#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

