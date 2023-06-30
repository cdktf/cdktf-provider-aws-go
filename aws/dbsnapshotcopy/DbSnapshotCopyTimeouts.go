package dbsnapshotcopy


type DbSnapshotCopyTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.1/docs/resources/db_snapshot_copy#create DbSnapshotCopy#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
}

