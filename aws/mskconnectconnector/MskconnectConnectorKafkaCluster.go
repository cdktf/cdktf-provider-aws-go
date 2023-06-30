package mskconnectconnector


type MskconnectConnectorKafkaCluster struct {
	// apache_kafka_cluster block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.1/docs/resources/mskconnect_connector#apache_kafka_cluster MskconnectConnector#apache_kafka_cluster}
	ApacheKafkaCluster *MskconnectConnectorKafkaClusterApacheKafkaCluster `field:"required" json:"apacheKafkaCluster" yaml:"apacheKafkaCluster"`
}

