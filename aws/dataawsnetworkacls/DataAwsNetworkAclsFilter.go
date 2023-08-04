package dataawsnetworkacls


type DataAwsNetworkAclsFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.11.0/docs/data-sources/network_acls#name DataAwsNetworkAcls#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.11.0/docs/data-sources/network_acls#values DataAwsNetworkAcls#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

