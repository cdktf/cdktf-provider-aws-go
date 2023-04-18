package defaultroutetable


type DefaultRouteTableTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.63.0/docs/resources/default_route_table#create DefaultRouteTable#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.63.0/docs/resources/default_route_table#update DefaultRouteTable#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

