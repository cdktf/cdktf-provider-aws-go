package elb


type LbSubnetMapping struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lb#subnet_id Lb#subnet_id}.
	SubnetId *string `field:"required" json:"subnetId" yaml:"subnetId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lb#allocation_id Lb#allocation_id}.
	AllocationId *string `field:"optional" json:"allocationId" yaml:"allocationId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lb#ipv6_address Lb#ipv6_address}.
	Ipv6Address *string `field:"optional" json:"ipv6Address" yaml:"ipv6Address"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lb#private_ipv4_address Lb#private_ipv4_address}.
	PrivateIpv4Address *string `field:"optional" json:"privateIpv4Address" yaml:"privateIpv4Address"`
}

