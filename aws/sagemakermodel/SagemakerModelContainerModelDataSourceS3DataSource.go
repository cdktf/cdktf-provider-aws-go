// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sagemakermodel


type SagemakerModelContainerModelDataSourceS3DataSource struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.9.0/docs/resources/sagemaker_model#compression_type SagemakerModel#compression_type}.
	CompressionType *string `field:"required" json:"compressionType" yaml:"compressionType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.9.0/docs/resources/sagemaker_model#s3_data_type SagemakerModel#s3_data_type}.
	S3DataType *string `field:"required" json:"s3DataType" yaml:"s3DataType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.9.0/docs/resources/sagemaker_model#s3_uri SagemakerModel#s3_uri}.
	S3Uri *string `field:"required" json:"s3Uri" yaml:"s3Uri"`
	// model_access_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.9.0/docs/resources/sagemaker_model#model_access_config SagemakerModel#model_access_config}
	ModelAccessConfig *SagemakerModelContainerModelDataSourceS3DataSourceModelAccessConfig `field:"optional" json:"modelAccessConfig" yaml:"modelAccessConfig"`
}

