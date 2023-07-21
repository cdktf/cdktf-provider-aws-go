package defaultsubnet


type DefaultSubnetTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/default_subnet#create DefaultSubnet#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/default_subnet#delete DefaultSubnet#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

