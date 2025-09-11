// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dbinstance


type DbInstanceS3Import struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.13.0/docs/resources/db_instance#bucket_name DbInstance#bucket_name}.
	BucketName *string `field:"required" json:"bucketName" yaml:"bucketName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.13.0/docs/resources/db_instance#ingestion_role DbInstance#ingestion_role}.
	IngestionRole *string `field:"required" json:"ingestionRole" yaml:"ingestionRole"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.13.0/docs/resources/db_instance#source_engine DbInstance#source_engine}.
	SourceEngine *string `field:"required" json:"sourceEngine" yaml:"sourceEngine"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.13.0/docs/resources/db_instance#source_engine_version DbInstance#source_engine_version}.
	SourceEngineVersion *string `field:"required" json:"sourceEngineVersion" yaml:"sourceEngineVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.13.0/docs/resources/db_instance#bucket_prefix DbInstance#bucket_prefix}.
	BucketPrefix *string `field:"optional" json:"bucketPrefix" yaml:"bucketPrefix"`
}

