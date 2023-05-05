package datasynclocationfsxopenzfsfilesystem


type DatasyncLocationFsxOpenzfsFileSystemProtocol struct {
	// nfs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.0/docs/resources/datasync_location_fsx_openzfs_file_system#nfs DatasyncLocationFsxOpenzfsFileSystem#nfs}
	Nfs *DatasyncLocationFsxOpenzfsFileSystemProtocolNfs `field:"required" json:"nfs" yaml:"nfs"`
}

