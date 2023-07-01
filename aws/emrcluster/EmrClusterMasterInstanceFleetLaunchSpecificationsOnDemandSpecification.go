package emrcluster


type EmrClusterMasterInstanceFleetLaunchSpecificationsOnDemandSpecification struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.2/docs/resources/emr_cluster#allocation_strategy EmrCluster#allocation_strategy}.
	AllocationStrategy *string `field:"required" json:"allocationStrategy" yaml:"allocationStrategy"`
}

