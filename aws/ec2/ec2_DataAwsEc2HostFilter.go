package ec2


type DataAwsEc2HostFilter struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/ec2_host#name DataAwsEc2Host#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/ec2_host#values DataAwsEc2Host#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

