package emrcluster


type EmrClusterPlacementGroupConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.0/docs/resources/emr_cluster#instance_role EmrCluster#instance_role}.
	InstanceRole *string `field:"optional" json:"instanceRole" yaml:"instanceRole"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.0/docs/resources/emr_cluster#placement_strategy EmrCluster#placement_strategy}.
	PlacementStrategy *string `field:"optional" json:"placementStrategy" yaml:"placementStrategy"`
}

