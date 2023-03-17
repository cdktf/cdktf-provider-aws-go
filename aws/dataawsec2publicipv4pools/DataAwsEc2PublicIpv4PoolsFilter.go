package dataawsec2publicipv4pools


type DataAwsEc2PublicIpv4PoolsFilter struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/ec2_public_ipv4_pools#name DataAwsEc2PublicIpv4Pools#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/ec2_public_ipv4_pools#values DataAwsEc2PublicIpv4Pools#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

