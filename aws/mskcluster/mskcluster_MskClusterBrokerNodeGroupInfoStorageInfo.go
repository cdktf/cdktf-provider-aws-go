package mskcluster


type MskClusterBrokerNodeGroupInfoStorageInfo struct {
	// ebs_storage_info block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/msk_cluster#ebs_storage_info MskCluster#ebs_storage_info}
	EbsStorageInfo *MskClusterBrokerNodeGroupInfoStorageInfoEbsStorageInfo `field:"optional" json:"ebsStorageInfo" yaml:"ebsStorageInfo"`
}

