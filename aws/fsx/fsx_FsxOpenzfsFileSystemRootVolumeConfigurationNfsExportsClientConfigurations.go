package fsx


type FsxOpenzfsFileSystemRootVolumeConfigurationNfsExportsClientConfigurations struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_openzfs_file_system#clients FsxOpenzfsFileSystem#clients}.
	Clients *string `field:"required" json:"clients" yaml:"clients"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_openzfs_file_system#options FsxOpenzfsFileSystem#options}.
	Options *[]*string `field:"required" json:"options" yaml:"options"`
}

