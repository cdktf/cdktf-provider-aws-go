package dbsnapshotcopy


type DbSnapshotCopyTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/db_snapshot_copy#create DbSnapshotCopy#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
}

