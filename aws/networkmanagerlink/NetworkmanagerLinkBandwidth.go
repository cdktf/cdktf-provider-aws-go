package networkmanagerlink


type NetworkmanagerLinkBandwidth struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.0/docs/resources/networkmanager_link#download_speed NetworkmanagerLink#download_speed}.
	DownloadSpeed *float64 `field:"optional" json:"downloadSpeed" yaml:"downloadSpeed"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.0/docs/resources/networkmanager_link#upload_speed NetworkmanagerLink#upload_speed}.
	UploadSpeed *float64 `field:"optional" json:"uploadSpeed" yaml:"uploadSpeed"`
}

