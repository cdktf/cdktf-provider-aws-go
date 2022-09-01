package vpc


type NetworkmanagerVpcAttachmentOptions struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/networkmanager_vpc_attachment#ipv6_support NetworkmanagerVpcAttachment#ipv6_support}.
	Ipv6Support interface{} `field:"required" json:"ipv6Support" yaml:"ipv6Support"`
}

