package ec2


type DataAwsEc2CoipPoolsFilter struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/ec2_coip_pools#name DataAwsEc2CoipPools#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/ec2_coip_pools#values DataAwsEc2CoipPools#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

