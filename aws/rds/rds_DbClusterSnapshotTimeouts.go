package rds


type DbClusterSnapshotTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/db_cluster_snapshot#create DbClusterSnapshot#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
}

