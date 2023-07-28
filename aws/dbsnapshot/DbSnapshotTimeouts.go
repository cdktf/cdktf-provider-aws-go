package dbsnapshot


type DbSnapshotTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.10.0/docs/resources/db_snapshot#create DbSnapshot#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
}

