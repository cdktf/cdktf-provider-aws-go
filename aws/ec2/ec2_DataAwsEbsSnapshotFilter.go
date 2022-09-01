package ec2


type DataAwsEbsSnapshotFilter struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/ebs_snapshot#name DataAwsEbsSnapshot#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/ebs_snapshot#values DataAwsEbsSnapshot#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

