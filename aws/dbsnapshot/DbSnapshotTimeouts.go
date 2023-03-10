package dbsnapshot


type DbSnapshotTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/db_snapshot#create DbSnapshot#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
}

