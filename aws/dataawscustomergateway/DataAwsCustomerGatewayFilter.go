package dataawscustomergateway


type DataAwsCustomerGatewayFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.8.0/docs/data-sources/customer_gateway#name DataAwsCustomerGateway#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.8.0/docs/data-sources/customer_gateway#values DataAwsCustomerGateway#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

