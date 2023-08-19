package mskconnectconnector


type MskconnectConnectorWorkerConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/mskconnect_connector#arn MskconnectConnector#arn}.
	Arn *string `field:"required" json:"arn" yaml:"arn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/mskconnect_connector#revision MskconnectConnector#revision}.
	Revision *float64 `field:"required" json:"revision" yaml:"revision"`
}

