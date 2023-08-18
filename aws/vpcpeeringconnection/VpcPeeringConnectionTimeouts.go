package vpcpeeringconnection


type VpcPeeringConnectionTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.0/docs/resources/vpc_peering_connection#create VpcPeeringConnection#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.0/docs/resources/vpc_peering_connection#delete VpcPeeringConnection#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.0/docs/resources/vpc_peering_connection#update VpcPeeringConnection#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

