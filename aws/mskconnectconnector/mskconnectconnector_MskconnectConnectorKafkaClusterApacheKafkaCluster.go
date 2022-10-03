package mskconnectconnector


type MskconnectConnectorKafkaClusterApacheKafkaCluster struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/mskconnect_connector#bootstrap_servers MskconnectConnector#bootstrap_servers}.
	BootstrapServers *string `field:"required" json:"bootstrapServers" yaml:"bootstrapServers"`
	// vpc block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/mskconnect_connector#vpc MskconnectConnector#vpc}
	Vpc *MskconnectConnectorKafkaClusterApacheKafkaClusterVpc `field:"required" json:"vpc" yaml:"vpc"`
}

