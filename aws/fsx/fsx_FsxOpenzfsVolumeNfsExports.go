package fsx


type FsxOpenzfsVolumeNfsExports struct {
	// client_configurations block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_openzfs_volume#client_configurations FsxOpenzfsVolume#client_configurations}
	ClientConfigurations interface{} `field:"required" json:"clientConfigurations" yaml:"clientConfigurations"`
}

