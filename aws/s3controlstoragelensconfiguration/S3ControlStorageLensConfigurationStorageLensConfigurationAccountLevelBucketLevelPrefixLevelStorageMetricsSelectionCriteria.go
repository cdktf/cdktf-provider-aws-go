// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package s3controlstoragelensconfiguration


type S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelPrefixLevelStorageMetricsSelectionCriteria struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.0/docs/resources/s3control_storage_lens_configuration#delimiter S3ControlStorageLensConfiguration#delimiter}.
	Delimiter *string `field:"optional" json:"delimiter" yaml:"delimiter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.0/docs/resources/s3control_storage_lens_configuration#max_depth S3ControlStorageLensConfiguration#max_depth}.
	MaxDepth *float64 `field:"optional" json:"maxDepth" yaml:"maxDepth"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.0/docs/resources/s3control_storage_lens_configuration#min_storage_bytes_percentage S3ControlStorageLensConfiguration#min_storage_bytes_percentage}.
	MinStorageBytesPercentage *float64 `field:"optional" json:"minStorageBytesPercentage" yaml:"minStorageBytesPercentage"`
}

