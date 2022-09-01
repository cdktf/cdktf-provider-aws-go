package vpc


type DataAwsRouteTableFilter struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/route_table#name DataAwsRouteTable#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/route_table#values DataAwsRouteTable#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

