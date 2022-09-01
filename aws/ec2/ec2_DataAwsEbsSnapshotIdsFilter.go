package ec2


type DataAwsEbsSnapshotIdsFilter struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/ebs_snapshot_ids#name DataAwsEbsSnapshotIds#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/ebs_snapshot_ids#values DataAwsEbsSnapshotIds#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

