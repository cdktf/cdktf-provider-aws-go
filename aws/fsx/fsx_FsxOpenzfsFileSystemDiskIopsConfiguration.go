package fsx


type FsxOpenzfsFileSystemDiskIopsConfiguration struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_openzfs_file_system#iops FsxOpenzfsFileSystem#iops}.
	Iops *float64 `field:"optional" json:"iops" yaml:"iops"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_openzfs_file_system#mode FsxOpenzfsFileSystem#mode}.
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
}

