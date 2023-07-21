package networkmanagersite


type NetworkmanagerSiteLocation struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/networkmanager_site#address NetworkmanagerSite#address}.
	Address *string `field:"optional" json:"address" yaml:"address"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/networkmanager_site#latitude NetworkmanagerSite#latitude}.
	Latitude *string `field:"optional" json:"latitude" yaml:"latitude"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/networkmanager_site#longitude NetworkmanagerSite#longitude}.
	Longitude *string `field:"optional" json:"longitude" yaml:"longitude"`
}

