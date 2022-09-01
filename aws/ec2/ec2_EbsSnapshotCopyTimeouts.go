package ec2


type EbsSnapshotCopyTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ebs_snapshot_copy#create EbsSnapshotCopy#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ebs_snapshot_copy#delete EbsSnapshotCopy#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

