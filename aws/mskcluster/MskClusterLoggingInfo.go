package mskcluster


type MskClusterLoggingInfo struct {
	// broker_logs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.8.0/docs/resources/msk_cluster#broker_logs MskCluster#broker_logs}
	BrokerLogs *MskClusterLoggingInfoBrokerLogs `field:"required" json:"brokerLogs" yaml:"brokerLogs"`
}

