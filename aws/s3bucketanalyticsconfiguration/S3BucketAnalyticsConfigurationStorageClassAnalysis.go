// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package s3bucketanalyticsconfiguration


type S3BucketAnalyticsConfigurationStorageClassAnalysis struct {
	// data_export block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.16.2/docs/resources/s3_bucket_analytics_configuration#data_export S3BucketAnalyticsConfiguration#data_export}
	DataExport *S3BucketAnalyticsConfigurationStorageClassAnalysisDataExport `field:"required" json:"dataExport" yaml:"dataExport"`
}

