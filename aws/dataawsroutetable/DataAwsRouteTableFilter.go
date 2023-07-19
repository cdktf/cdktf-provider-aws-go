package dataawsroutetable


type DataAwsRouteTableFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.8.0/docs/data-sources/route_table#name DataAwsRouteTable#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.8.0/docs/data-sources/route_table#values DataAwsRouteTable#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

