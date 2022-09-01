package ram


type RamResourceShareAccepterTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ram_resource_share_accepter#create RamResourceShareAccepter#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ram_resource_share_accepter#delete RamResourceShareAccepter#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

