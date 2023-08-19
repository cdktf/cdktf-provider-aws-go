package mskcluster


type MskClusterBrokerNodeGroupInfoConnectivityInfoVpcConnectivity struct {
	// client_authentication block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/msk_cluster#client_authentication MskCluster#client_authentication}
	ClientAuthentication *MskClusterBrokerNodeGroupInfoConnectivityInfoVpcConnectivityClientAuthentication `field:"optional" json:"clientAuthentication" yaml:"clientAuthentication"`
}

