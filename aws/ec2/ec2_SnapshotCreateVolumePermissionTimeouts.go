package ec2


type SnapshotCreateVolumePermissionTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/snapshot_create_volume_permission#create SnapshotCreateVolumePermission#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/snapshot_create_volume_permission#delete SnapshotCreateVolumePermission#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

