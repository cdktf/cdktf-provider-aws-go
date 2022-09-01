package ec2


type DataAwsEc2LocalGatewaysFilter struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/ec2_local_gateways#name DataAwsEc2LocalGateways#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/ec2_local_gateways#values DataAwsEc2LocalGateways#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

