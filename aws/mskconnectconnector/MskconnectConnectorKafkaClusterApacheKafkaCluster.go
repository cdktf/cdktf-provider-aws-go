package mskconnectconnector


type MskconnectConnectorKafkaClusterApacheKafkaCluster struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.1/docs/resources/mskconnect_connector#bootstrap_servers MskconnectConnector#bootstrap_servers}.
	BootstrapServers *string `field:"required" json:"bootstrapServers" yaml:"bootstrapServers"`
	// vpc block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.1/docs/resources/mskconnect_connector#vpc MskconnectConnector#vpc}
	Vpc *MskconnectConnectorKafkaClusterApacheKafkaClusterVpc `field:"required" json:"vpc" yaml:"vpc"`
}

