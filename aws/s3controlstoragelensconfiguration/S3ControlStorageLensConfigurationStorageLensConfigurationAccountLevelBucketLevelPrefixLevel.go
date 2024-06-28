// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package s3controlstoragelensconfiguration


type S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelPrefixLevel struct {
	// storage_metrics block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.56.0/docs/resources/s3control_storage_lens_configuration#storage_metrics S3ControlStorageLensConfiguration#storage_metrics}
	StorageMetrics *S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelPrefixLevelStorageMetrics `field:"required" json:"storageMetrics" yaml:"storageMetrics"`
}

