package mskconnectconnector


type MskconnectConnectorPluginCustomPlugin struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.0/docs/resources/mskconnect_connector#arn MskconnectConnector#arn}.
	Arn *string `field:"required" json:"arn" yaml:"arn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.0/docs/resources/mskconnect_connector#revision MskconnectConnector#revision}.
	Revision *float64 `field:"required" json:"revision" yaml:"revision"`
}

