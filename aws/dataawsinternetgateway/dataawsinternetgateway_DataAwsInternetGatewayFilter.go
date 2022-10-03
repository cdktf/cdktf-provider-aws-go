package dataawsinternetgateway


type DataAwsInternetGatewayFilter struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/internet_gateway#name DataAwsInternetGateway#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/internet_gateway#values DataAwsInternetGateway#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

