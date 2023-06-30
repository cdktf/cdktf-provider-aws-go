package dataawsvpcdhcpoptions


type DataAwsVpcDhcpOptionsFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.1/docs/data-sources/vpc_dhcp_options#name DataAwsVpcDhcpOptions#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.1/docs/data-sources/vpc_dhcp_options#values DataAwsVpcDhcpOptions#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

