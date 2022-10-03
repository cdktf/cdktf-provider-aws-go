package mskcluster


type MskClusterBrokerNodeGroupInfoConnectivityInfo struct {
	// public_access block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/msk_cluster#public_access MskCluster#public_access}
	PublicAccess *MskClusterBrokerNodeGroupInfoConnectivityInfoPublicAccess `field:"optional" json:"publicAccess" yaml:"publicAccess"`
}

