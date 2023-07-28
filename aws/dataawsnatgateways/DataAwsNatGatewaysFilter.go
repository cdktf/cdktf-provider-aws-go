package dataawsnatgateways


type DataAwsNatGatewaysFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.10.0/docs/data-sources/nat_gateways#name DataAwsNatGateways#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.10.0/docs/data-sources/nat_gateways#values DataAwsNatGateways#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

