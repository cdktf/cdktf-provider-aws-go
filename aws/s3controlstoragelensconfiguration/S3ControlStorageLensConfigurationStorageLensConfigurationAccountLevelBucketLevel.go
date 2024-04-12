// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package s3controlstoragelensconfiguration


type S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevel struct {
	// activity_metrics block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.45.0/docs/resources/s3control_storage_lens_configuration#activity_metrics S3ControlStorageLensConfiguration#activity_metrics}
	ActivityMetrics *S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelActivityMetrics `field:"optional" json:"activityMetrics" yaml:"activityMetrics"`
	// advanced_cost_optimization_metrics block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.45.0/docs/resources/s3control_storage_lens_configuration#advanced_cost_optimization_metrics S3ControlStorageLensConfiguration#advanced_cost_optimization_metrics}
	AdvancedCostOptimizationMetrics *S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelAdvancedCostOptimizationMetrics `field:"optional" json:"advancedCostOptimizationMetrics" yaml:"advancedCostOptimizationMetrics"`
	// advanced_data_protection_metrics block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.45.0/docs/resources/s3control_storage_lens_configuration#advanced_data_protection_metrics S3ControlStorageLensConfiguration#advanced_data_protection_metrics}
	AdvancedDataProtectionMetrics *S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelAdvancedDataProtectionMetrics `field:"optional" json:"advancedDataProtectionMetrics" yaml:"advancedDataProtectionMetrics"`
	// detailed_status_code_metrics block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.45.0/docs/resources/s3control_storage_lens_configuration#detailed_status_code_metrics S3ControlStorageLensConfiguration#detailed_status_code_metrics}
	DetailedStatusCodeMetrics *S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelDetailedStatusCodeMetrics `field:"optional" json:"detailedStatusCodeMetrics" yaml:"detailedStatusCodeMetrics"`
	// prefix_level block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.45.0/docs/resources/s3control_storage_lens_configuration#prefix_level S3ControlStorageLensConfiguration#prefix_level}
	PrefixLevel *S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelPrefixLevel `field:"optional" json:"prefixLevel" yaml:"prefixLevel"`
}

