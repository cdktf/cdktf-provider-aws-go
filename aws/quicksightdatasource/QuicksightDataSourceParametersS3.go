// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package quicksightdatasource


type QuicksightDataSourceParametersS3 struct {
	// manifest_file_location block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.26.0/docs/resources/quicksight_data_source#manifest_file_location QuicksightDataSource#manifest_file_location}
	ManifestFileLocation *QuicksightDataSourceParametersS3ManifestFileLocation `field:"required" json:"manifestFileLocation" yaml:"manifestFileLocation"`
}

