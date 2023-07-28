package fsxfilecache


type FsxFileCacheDataRepositoryAssociation struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.10.0/docs/resources/fsx_file_cache#data_repository_path FsxFileCache#data_repository_path}.
	DataRepositoryPath *string `field:"required" json:"dataRepositoryPath" yaml:"dataRepositoryPath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.10.0/docs/resources/fsx_file_cache#file_cache_path FsxFileCache#file_cache_path}.
	FileCachePath *string `field:"required" json:"fileCachePath" yaml:"fileCachePath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.10.0/docs/resources/fsx_file_cache#data_repository_subdirectories FsxFileCache#data_repository_subdirectories}.
	DataRepositorySubdirectories *[]*string `field:"optional" json:"dataRepositorySubdirectories" yaml:"dataRepositorySubdirectories"`
	// nfs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.10.0/docs/resources/fsx_file_cache#nfs FsxFileCache#nfs}
	Nfs interface{} `field:"optional" json:"nfs" yaml:"nfs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.10.0/docs/resources/fsx_file_cache#tags FsxFileCache#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

