package quicksight


type QuicksightDataSourceParametersS3 struct {
	// manifest_file_location block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/quicksight_data_source#manifest_file_location QuicksightDataSource#manifest_file_location}
	ManifestFileLocation *QuicksightDataSourceParametersS3ManifestFileLocation `field:"required" json:"manifestFileLocation" yaml:"manifestFileLocation"`
}

