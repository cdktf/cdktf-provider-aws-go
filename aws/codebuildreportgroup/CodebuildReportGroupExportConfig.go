// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package codebuildreportgroup


type CodebuildReportGroupExportConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.94.0/docs/resources/codebuild_report_group#type CodebuildReportGroup#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// s3_destination block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.94.0/docs/resources/codebuild_report_group#s3_destination CodebuildReportGroup#s3_destination}
	S3Destination *CodebuildReportGroupExportConfigS3Destination `field:"optional" json:"s3Destination" yaml:"s3Destination"`
}

