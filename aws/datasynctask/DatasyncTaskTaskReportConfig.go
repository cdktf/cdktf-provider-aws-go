// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datasynctask


type DatasyncTaskTaskReportConfig struct {
	// s3_destination block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.45.0/docs/resources/datasync_task#s3_destination DatasyncTask#s3_destination}
	S3Destination *DatasyncTaskTaskReportConfigS3Destination `field:"required" json:"s3Destination" yaml:"s3Destination"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.45.0/docs/resources/datasync_task#output_type DatasyncTask#output_type}.
	OutputType *string `field:"optional" json:"outputType" yaml:"outputType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.45.0/docs/resources/datasync_task#report_level DatasyncTask#report_level}.
	ReportLevel *string `field:"optional" json:"reportLevel" yaml:"reportLevel"`
	// report_overrides block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.45.0/docs/resources/datasync_task#report_overrides DatasyncTask#report_overrides}
	ReportOverrides *DatasyncTaskTaskReportConfigReportOverrides `field:"optional" json:"reportOverrides" yaml:"reportOverrides"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.45.0/docs/resources/datasync_task#s3_object_versioning DatasyncTask#s3_object_versioning}.
	S3ObjectVersioning *string `field:"optional" json:"s3ObjectVersioning" yaml:"s3ObjectVersioning"`
}

