package fsxopenzfsfilesystem


type FsxOpenzfsFileSystemRootVolumeConfigurationNfsExports struct {
	// client_configurations block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_openzfs_file_system#client_configurations FsxOpenzfsFileSystem#client_configurations}
	ClientConfigurations interface{} `field:"required" json:"clientConfigurations" yaml:"clientConfigurations"`
}

