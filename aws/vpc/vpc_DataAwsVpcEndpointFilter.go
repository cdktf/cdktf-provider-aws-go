package vpc


type DataAwsVpcEndpointFilter struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/vpc_endpoint#name DataAwsVpcEndpoint#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/vpc_endpoint#values DataAwsVpcEndpoint#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

