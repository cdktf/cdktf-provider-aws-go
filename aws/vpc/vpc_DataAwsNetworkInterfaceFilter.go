package vpc


type DataAwsNetworkInterfaceFilter struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/network_interface#name DataAwsNetworkInterface#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/network_interface#values DataAwsNetworkInterface#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

