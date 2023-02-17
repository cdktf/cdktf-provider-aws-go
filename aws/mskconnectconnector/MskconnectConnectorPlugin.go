package mskconnectconnector


type MskconnectConnectorPlugin struct {
	// custom_plugin block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/mskconnect_connector#custom_plugin MskconnectConnector#custom_plugin}
	CustomPlugin *MskconnectConnectorPluginCustomPlugin `field:"required" json:"customPlugin" yaml:"customPlugin"`
}

