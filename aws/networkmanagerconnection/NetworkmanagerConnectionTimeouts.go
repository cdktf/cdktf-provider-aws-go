package networkmanagerconnection


type NetworkmanagerConnectionTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.10.0/docs/resources/networkmanager_connection#create NetworkmanagerConnection#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.10.0/docs/resources/networkmanager_connection#delete NetworkmanagerConnection#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.10.0/docs/resources/networkmanager_connection#update NetworkmanagerConnection#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

