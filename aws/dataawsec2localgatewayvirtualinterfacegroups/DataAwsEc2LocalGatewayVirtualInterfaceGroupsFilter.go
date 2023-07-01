package dataawsec2localgatewayvirtualinterfacegroups


type DataAwsEc2LocalGatewayVirtualInterfaceGroupsFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.2/docs/data-sources/ec2_local_gateway_virtual_interface_groups#name DataAwsEc2LocalGatewayVirtualInterfaceGroups#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.2/docs/data-sources/ec2_local_gateway_virtual_interface_groups#values DataAwsEc2LocalGatewayVirtualInterfaceGroups#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

