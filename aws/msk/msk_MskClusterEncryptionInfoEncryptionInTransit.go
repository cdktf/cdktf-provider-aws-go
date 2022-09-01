package msk


type MskClusterEncryptionInfoEncryptionInTransit struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/msk_cluster#client_broker MskCluster#client_broker}.
	ClientBroker *string `field:"optional" json:"clientBroker" yaml:"clientBroker"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/msk_cluster#in_cluster MskCluster#in_cluster}.
	InCluster interface{} `field:"optional" json:"inCluster" yaml:"inCluster"`
}

