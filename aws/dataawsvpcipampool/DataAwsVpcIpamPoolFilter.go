package dataawsvpcipampool


type DataAwsVpcIpamPoolFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/data-sources/vpc_ipam_pool#name DataAwsVpcIpamPool#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/data-sources/vpc_ipam_pool#values DataAwsVpcIpamPool#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

