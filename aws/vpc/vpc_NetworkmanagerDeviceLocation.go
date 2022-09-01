package vpc


type NetworkmanagerDeviceLocation struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/networkmanager_device#address NetworkmanagerDevice#address}.
	Address *string `field:"optional" json:"address" yaml:"address"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/networkmanager_device#latitude NetworkmanagerDevice#latitude}.
	Latitude *string `field:"optional" json:"latitude" yaml:"latitude"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/networkmanager_device#longitude NetworkmanagerDevice#longitude}.
	Longitude *string `field:"optional" json:"longitude" yaml:"longitude"`
}

