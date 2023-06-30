package dataawsinternetgateway


type DataAwsInternetGatewayFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.1/docs/data-sources/internet_gateway#name DataAwsInternetGateway#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.1/docs/data-sources/internet_gateway#values DataAwsInternetGateway#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

