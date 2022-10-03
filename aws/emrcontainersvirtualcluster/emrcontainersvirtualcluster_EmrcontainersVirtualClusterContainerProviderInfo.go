package emrcontainersvirtualcluster


type EmrcontainersVirtualClusterContainerProviderInfo struct {
	// eks_info block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emrcontainers_virtual_cluster#eks_info EmrcontainersVirtualCluster#eks_info}
	EksInfo *EmrcontainersVirtualClusterContainerProviderInfoEksInfo `field:"required" json:"eksInfo" yaml:"eksInfo"`
}

