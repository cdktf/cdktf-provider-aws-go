package fsxontapfilesystem


type FsxOntapFileSystemDiskIopsConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.0/docs/resources/fsx_ontap_file_system#iops FsxOntapFileSystem#iops}.
	Iops *float64 `field:"optional" json:"iops" yaml:"iops"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.0/docs/resources/fsx_ontap_file_system#mode FsxOntapFileSystem#mode}.
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
}

