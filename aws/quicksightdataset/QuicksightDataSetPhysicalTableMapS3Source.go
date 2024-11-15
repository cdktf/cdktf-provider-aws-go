// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package quicksightdataset


type QuicksightDataSetPhysicalTableMapS3Source struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.76.0/docs/resources/quicksight_data_set#data_source_arn QuicksightDataSet#data_source_arn}.
	DataSourceArn *string `field:"required" json:"dataSourceArn" yaml:"dataSourceArn"`
	// input_columns block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.76.0/docs/resources/quicksight_data_set#input_columns QuicksightDataSet#input_columns}
	InputColumns interface{} `field:"required" json:"inputColumns" yaml:"inputColumns"`
	// upload_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.76.0/docs/resources/quicksight_data_set#upload_settings QuicksightDataSet#upload_settings}
	UploadSettings *QuicksightDataSetPhysicalTableMapS3SourceUploadSettings `field:"required" json:"uploadSettings" yaml:"uploadSettings"`
}

