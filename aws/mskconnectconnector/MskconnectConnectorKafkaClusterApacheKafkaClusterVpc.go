package mskconnectconnector


type MskconnectConnectorKafkaClusterApacheKafkaClusterVpc struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.2/docs/resources/mskconnect_connector#security_groups MskconnectConnector#security_groups}.
	SecurityGroups *[]*string `field:"required" json:"securityGroups" yaml:"securityGroups"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.2/docs/resources/mskconnect_connector#subnets MskconnectConnector#subnets}.
	Subnets *[]*string `field:"required" json:"subnets" yaml:"subnets"`
}

