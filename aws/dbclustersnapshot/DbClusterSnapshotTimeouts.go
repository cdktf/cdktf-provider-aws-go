package dbclustersnapshot


type DbClusterSnapshotTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.65.0/docs/resources/db_cluster_snapshot#create DbClusterSnapshot#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
}

