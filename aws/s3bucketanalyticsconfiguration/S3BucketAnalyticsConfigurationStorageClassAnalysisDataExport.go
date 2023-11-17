// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package s3bucketanalyticsconfiguration


type S3BucketAnalyticsConfigurationStorageClassAnalysisDataExport struct {
	// destination block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.26.0/docs/resources/s3_bucket_analytics_configuration#destination S3BucketAnalyticsConfiguration#destination}
	Destination *S3BucketAnalyticsConfigurationStorageClassAnalysisDataExportDestination `field:"required" json:"destination" yaml:"destination"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.26.0/docs/resources/s3_bucket_analytics_configuration#output_schema_version S3BucketAnalyticsConfiguration#output_schema_version}.
	OutputSchemaVersion *string `field:"optional" json:"outputSchemaVersion" yaml:"outputSchemaVersion"`
}

