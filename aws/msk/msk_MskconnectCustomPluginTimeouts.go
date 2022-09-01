package msk


type MskconnectCustomPluginTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/mskconnect_custom_plugin#create MskconnectCustomPlugin#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/mskconnect_custom_plugin#delete MskconnectCustomPlugin#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

