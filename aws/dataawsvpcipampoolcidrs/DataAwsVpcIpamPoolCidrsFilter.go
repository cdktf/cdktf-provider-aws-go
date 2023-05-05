package dataawsvpcipampoolcidrs


type DataAwsVpcIpamPoolCidrsFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.0/docs/data-sources/vpc_ipam_pool_cidrs#name DataAwsVpcIpamPoolCidrs#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.0/docs/data-sources/vpc_ipam_pool_cidrs#values DataAwsVpcIpamPoolCidrs#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

