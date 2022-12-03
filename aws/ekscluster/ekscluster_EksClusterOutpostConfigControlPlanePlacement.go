package ekscluster


type EksClusterOutpostConfigControlPlanePlacement struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/eks_cluster#group_name EksCluster#group_name}.
	GroupName *string `field:"required" json:"groupName" yaml:"groupName"`
}

