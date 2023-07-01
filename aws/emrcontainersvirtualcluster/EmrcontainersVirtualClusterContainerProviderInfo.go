package emrcontainersvirtualcluster


type EmrcontainersVirtualClusterContainerProviderInfo struct {
	// eks_info block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.2/docs/resources/emrcontainers_virtual_cluster#eks_info EmrcontainersVirtualCluster#eks_info}
	EksInfo *EmrcontainersVirtualClusterContainerProviderInfoEksInfo `field:"required" json:"eksInfo" yaml:"eksInfo"`
}

