package vpc


type NetworkmanagerConnectionTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/networkmanager_connection#create NetworkmanagerConnection#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/networkmanager_connection#delete NetworkmanagerConnection#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/networkmanager_connection#update NetworkmanagerConnection#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

