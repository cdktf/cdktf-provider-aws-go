package rds


type DbSnapshotTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/db_snapshot#read DbSnapshot#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
}

