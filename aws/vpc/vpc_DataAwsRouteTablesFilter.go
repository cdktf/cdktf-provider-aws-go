package vpc


type DataAwsRouteTablesFilter struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/route_tables#name DataAwsRouteTables#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/route_tables#values DataAwsRouteTables#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

