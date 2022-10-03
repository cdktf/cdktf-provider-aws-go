package datasynclocationfsxopenzfsfilesystem


type DatasyncLocationFsxOpenzfsFileSystemProtocol struct {
	// nfs block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/datasync_location_fsx_openzfs_file_system#nfs DatasyncLocationFsxOpenzfsFileSystem#nfs}
	Nfs *DatasyncLocationFsxOpenzfsFileSystemProtocolNfs `field:"required" json:"nfs" yaml:"nfs"`
}

