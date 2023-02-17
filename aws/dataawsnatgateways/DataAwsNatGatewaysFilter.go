package dataawsnatgateways


type DataAwsNatGatewaysFilter struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/nat_gateways#name DataAwsNatGateways#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/nat_gateways#values DataAwsNatGateways#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

