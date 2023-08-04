package mskcluster


type MskClusterBrokerNodeGroupInfoConnectivityInfo struct {
	// public_access block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.11.0/docs/resources/msk_cluster#public_access MskCluster#public_access}
	PublicAccess *MskClusterBrokerNodeGroupInfoConnectivityInfoPublicAccess `field:"optional" json:"publicAccess" yaml:"publicAccess"`
}

