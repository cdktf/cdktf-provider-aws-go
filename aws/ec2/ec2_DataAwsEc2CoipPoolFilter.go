package ec2


type DataAwsEc2CoipPoolFilter struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/ec2_coip_pool#name DataAwsEc2CoipPool#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/ec2_coip_pool#values DataAwsEc2CoipPool#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

