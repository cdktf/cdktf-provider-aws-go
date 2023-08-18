package fsxopenzfsvolume


type FsxOpenzfsVolumeNfsExports struct {
	// client_configurations block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.0/docs/resources/fsx_openzfs_volume#client_configurations FsxOpenzfsVolume#client_configurations}
	ClientConfigurations interface{} `field:"required" json:"clientConfigurations" yaml:"clientConfigurations"`
}

