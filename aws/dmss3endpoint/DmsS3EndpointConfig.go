// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dmss3endpoint

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DmsS3EndpointConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/resources/dms_s3_endpoint#bucket_name DmsS3Endpoint#bucket_name}.
	BucketName *string `field:"required" json:"bucketName" yaml:"bucketName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/resources/dms_s3_endpoint#endpoint_id DmsS3Endpoint#endpoint_id}.
	EndpointId *string `field:"required" json:"endpointId" yaml:"endpointId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/resources/dms_s3_endpoint#endpoint_type DmsS3Endpoint#endpoint_type}.
	EndpointType *string `field:"required" json:"endpointType" yaml:"endpointType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/resources/dms_s3_endpoint#service_access_role_arn DmsS3Endpoint#service_access_role_arn}.
	ServiceAccessRoleArn *string `field:"required" json:"serviceAccessRoleArn" yaml:"serviceAccessRoleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/resources/dms_s3_endpoint#add_column_name DmsS3Endpoint#add_column_name}.
	AddColumnName interface{} `field:"optional" json:"addColumnName" yaml:"addColumnName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/resources/dms_s3_endpoint#add_trailing_padding_character DmsS3Endpoint#add_trailing_padding_character}.
	AddTrailingPaddingCharacter interface{} `field:"optional" json:"addTrailingPaddingCharacter" yaml:"addTrailingPaddingCharacter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/resources/dms_s3_endpoint#bucket_folder DmsS3Endpoint#bucket_folder}.
	BucketFolder *string `field:"optional" json:"bucketFolder" yaml:"bucketFolder"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/resources/dms_s3_endpoint#canned_acl_for_objects DmsS3Endpoint#canned_acl_for_objects}.
	CannedAclForObjects *string `field:"optional" json:"cannedAclForObjects" yaml:"cannedAclForObjects"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/resources/dms_s3_endpoint#cdc_inserts_and_updates DmsS3Endpoint#cdc_inserts_and_updates}.
	CdcInsertsAndUpdates interface{} `field:"optional" json:"cdcInsertsAndUpdates" yaml:"cdcInsertsAndUpdates"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/resources/dms_s3_endpoint#cdc_inserts_only DmsS3Endpoint#cdc_inserts_only}.
	CdcInsertsOnly interface{} `field:"optional" json:"cdcInsertsOnly" yaml:"cdcInsertsOnly"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/resources/dms_s3_endpoint#cdc_max_batch_interval DmsS3Endpoint#cdc_max_batch_interval}.
	CdcMaxBatchInterval *float64 `field:"optional" json:"cdcMaxBatchInterval" yaml:"cdcMaxBatchInterval"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/resources/dms_s3_endpoint#cdc_min_file_size DmsS3Endpoint#cdc_min_file_size}.
	CdcMinFileSize *float64 `field:"optional" json:"cdcMinFileSize" yaml:"cdcMinFileSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/resources/dms_s3_endpoint#cdc_path DmsS3Endpoint#cdc_path}.
	CdcPath *string `field:"optional" json:"cdcPath" yaml:"cdcPath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/resources/dms_s3_endpoint#certificate_arn DmsS3Endpoint#certificate_arn}.
	CertificateArn *string `field:"optional" json:"certificateArn" yaml:"certificateArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/resources/dms_s3_endpoint#compression_type DmsS3Endpoint#compression_type}.
	CompressionType *string `field:"optional" json:"compressionType" yaml:"compressionType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/resources/dms_s3_endpoint#csv_delimiter DmsS3Endpoint#csv_delimiter}.
	CsvDelimiter *string `field:"optional" json:"csvDelimiter" yaml:"csvDelimiter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/resources/dms_s3_endpoint#csv_no_sup_value DmsS3Endpoint#csv_no_sup_value}.
	CsvNoSupValue *string `field:"optional" json:"csvNoSupValue" yaml:"csvNoSupValue"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/resources/dms_s3_endpoint#csv_null_value DmsS3Endpoint#csv_null_value}.
	CsvNullValue *string `field:"optional" json:"csvNullValue" yaml:"csvNullValue"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/resources/dms_s3_endpoint#csv_row_delimiter DmsS3Endpoint#csv_row_delimiter}.
	CsvRowDelimiter *string `field:"optional" json:"csvRowDelimiter" yaml:"csvRowDelimiter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/resources/dms_s3_endpoint#data_format DmsS3Endpoint#data_format}.
	DataFormat *string `field:"optional" json:"dataFormat" yaml:"dataFormat"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/resources/dms_s3_endpoint#data_page_size DmsS3Endpoint#data_page_size}.
	DataPageSize *float64 `field:"optional" json:"dataPageSize" yaml:"dataPageSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/resources/dms_s3_endpoint#date_partition_delimiter DmsS3Endpoint#date_partition_delimiter}.
	DatePartitionDelimiter *string `field:"optional" json:"datePartitionDelimiter" yaml:"datePartitionDelimiter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/resources/dms_s3_endpoint#date_partition_enabled DmsS3Endpoint#date_partition_enabled}.
	DatePartitionEnabled interface{} `field:"optional" json:"datePartitionEnabled" yaml:"datePartitionEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/resources/dms_s3_endpoint#date_partition_sequence DmsS3Endpoint#date_partition_sequence}.
	DatePartitionSequence *string `field:"optional" json:"datePartitionSequence" yaml:"datePartitionSequence"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/resources/dms_s3_endpoint#date_partition_timezone DmsS3Endpoint#date_partition_timezone}.
	DatePartitionTimezone *string `field:"optional" json:"datePartitionTimezone" yaml:"datePartitionTimezone"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/resources/dms_s3_endpoint#detach_target_on_lob_lookup_failure_parquet DmsS3Endpoint#detach_target_on_lob_lookup_failure_parquet}.
	DetachTargetOnLobLookupFailureParquet interface{} `field:"optional" json:"detachTargetOnLobLookupFailureParquet" yaml:"detachTargetOnLobLookupFailureParquet"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/resources/dms_s3_endpoint#dict_page_size_limit DmsS3Endpoint#dict_page_size_limit}.
	DictPageSizeLimit *float64 `field:"optional" json:"dictPageSizeLimit" yaml:"dictPageSizeLimit"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/resources/dms_s3_endpoint#enable_statistics DmsS3Endpoint#enable_statistics}.
	EnableStatistics interface{} `field:"optional" json:"enableStatistics" yaml:"enableStatistics"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/resources/dms_s3_endpoint#encoding_type DmsS3Endpoint#encoding_type}.
	EncodingType *string `field:"optional" json:"encodingType" yaml:"encodingType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/resources/dms_s3_endpoint#encryption_mode DmsS3Endpoint#encryption_mode}.
	EncryptionMode *string `field:"optional" json:"encryptionMode" yaml:"encryptionMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/resources/dms_s3_endpoint#expected_bucket_owner DmsS3Endpoint#expected_bucket_owner}.
	ExpectedBucketOwner *string `field:"optional" json:"expectedBucketOwner" yaml:"expectedBucketOwner"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/resources/dms_s3_endpoint#external_table_definition DmsS3Endpoint#external_table_definition}.
	ExternalTableDefinition *string `field:"optional" json:"externalTableDefinition" yaml:"externalTableDefinition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/resources/dms_s3_endpoint#glue_catalog_generation DmsS3Endpoint#glue_catalog_generation}.
	GlueCatalogGeneration interface{} `field:"optional" json:"glueCatalogGeneration" yaml:"glueCatalogGeneration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/resources/dms_s3_endpoint#id DmsS3Endpoint#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/resources/dms_s3_endpoint#ignore_header_rows DmsS3Endpoint#ignore_header_rows}.
	IgnoreHeaderRows *float64 `field:"optional" json:"ignoreHeaderRows" yaml:"ignoreHeaderRows"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/resources/dms_s3_endpoint#include_op_for_full_load DmsS3Endpoint#include_op_for_full_load}.
	IncludeOpForFullLoad interface{} `field:"optional" json:"includeOpForFullLoad" yaml:"includeOpForFullLoad"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/resources/dms_s3_endpoint#kms_key_arn DmsS3Endpoint#kms_key_arn}.
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/resources/dms_s3_endpoint#max_file_size DmsS3Endpoint#max_file_size}.
	MaxFileSize *float64 `field:"optional" json:"maxFileSize" yaml:"maxFileSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/resources/dms_s3_endpoint#parquet_timestamp_in_millisecond DmsS3Endpoint#parquet_timestamp_in_millisecond}.
	ParquetTimestampInMillisecond interface{} `field:"optional" json:"parquetTimestampInMillisecond" yaml:"parquetTimestampInMillisecond"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/resources/dms_s3_endpoint#parquet_version DmsS3Endpoint#parquet_version}.
	ParquetVersion *string `field:"optional" json:"parquetVersion" yaml:"parquetVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/resources/dms_s3_endpoint#preserve_transactions DmsS3Endpoint#preserve_transactions}.
	PreserveTransactions interface{} `field:"optional" json:"preserveTransactions" yaml:"preserveTransactions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/resources/dms_s3_endpoint#rfc_4180 DmsS3Endpoint#rfc_4180}.
	Rfc4180 interface{} `field:"optional" json:"rfc4180" yaml:"rfc4180"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/resources/dms_s3_endpoint#row_group_length DmsS3Endpoint#row_group_length}.
	RowGroupLength *float64 `field:"optional" json:"rowGroupLength" yaml:"rowGroupLength"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/resources/dms_s3_endpoint#server_side_encryption_kms_key_id DmsS3Endpoint#server_side_encryption_kms_key_id}.
	ServerSideEncryptionKmsKeyId *string `field:"optional" json:"serverSideEncryptionKmsKeyId" yaml:"serverSideEncryptionKmsKeyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/resources/dms_s3_endpoint#ssl_mode DmsS3Endpoint#ssl_mode}.
	SslMode *string `field:"optional" json:"sslMode" yaml:"sslMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/resources/dms_s3_endpoint#tags DmsS3Endpoint#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/resources/dms_s3_endpoint#tags_all DmsS3Endpoint#tags_all}.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/resources/dms_s3_endpoint#timeouts DmsS3Endpoint#timeouts}
	Timeouts *DmsS3EndpointTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/resources/dms_s3_endpoint#timestamp_column_name DmsS3Endpoint#timestamp_column_name}.
	TimestampColumnName *string `field:"optional" json:"timestampColumnName" yaml:"timestampColumnName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/resources/dms_s3_endpoint#use_csv_no_sup_value DmsS3Endpoint#use_csv_no_sup_value}.
	UseCsvNoSupValue interface{} `field:"optional" json:"useCsvNoSupValue" yaml:"useCsvNoSupValue"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/resources/dms_s3_endpoint#use_task_start_time_for_full_load_timestamp DmsS3Endpoint#use_task_start_time_for_full_load_timestamp}.
	UseTaskStartTimeForFullLoadTimestamp interface{} `field:"optional" json:"useTaskStartTimeForFullLoadTimestamp" yaml:"useTaskStartTimeForFullLoadTimestamp"`
}

