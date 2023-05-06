package ekscluster


type EksClusterOutpostConfigControlPlanePlacement struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.1/docs/resources/eks_cluster#group_name EksCluster#group_name}.
	GroupName *string `field:"required" json:"groupName" yaml:"groupName"`
}

