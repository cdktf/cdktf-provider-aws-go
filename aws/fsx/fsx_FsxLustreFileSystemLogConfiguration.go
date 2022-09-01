package fsx


type FsxLustreFileSystemLogConfiguration struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_lustre_file_system#destination FsxLustreFileSystem#destination}.
	Destination *string `field:"optional" json:"destination" yaml:"destination"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_lustre_file_system#level FsxLustreFileSystem#level}.
	Level *string `field:"optional" json:"level" yaml:"level"`
}

