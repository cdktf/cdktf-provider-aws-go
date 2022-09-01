package fsx


type DataAwsFsxOpenzfsSnapshotFilter struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/fsx_openzfs_snapshot#name DataAwsFsxOpenzfsSnapshot#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/fsx_openzfs_snapshot#values DataAwsFsxOpenzfsSnapshot#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

