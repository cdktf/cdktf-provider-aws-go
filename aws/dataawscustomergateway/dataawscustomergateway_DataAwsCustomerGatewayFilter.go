package dataawscustomergateway


type DataAwsCustomerGatewayFilter struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/customer_gateway#name DataAwsCustomerGateway#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/customer_gateway#values DataAwsCustomerGateway#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

