package ekscluster


type EksClusterOutpostConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.1/docs/resources/eks_cluster#control_plane_instance_type EksCluster#control_plane_instance_type}.
	ControlPlaneInstanceType *string `field:"required" json:"controlPlaneInstanceType" yaml:"controlPlaneInstanceType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.1/docs/resources/eks_cluster#outpost_arns EksCluster#outpost_arns}.
	OutpostArns *[]*string `field:"required" json:"outpostArns" yaml:"outpostArns"`
	// control_plane_placement block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.1/docs/resources/eks_cluster#control_plane_placement EksCluster#control_plane_placement}
	ControlPlanePlacement *EksClusterOutpostConfigControlPlanePlacement `field:"optional" json:"controlPlanePlacement" yaml:"controlPlanePlacement"`
}

