package dmsendpoint


type DmsEndpointS3Settings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.1/docs/resources/dms_endpoint#add_column_name DmsEndpoint#add_column_name}.
	AddColumnName interface{} `field:"optional" json:"addColumnName" yaml:"addColumnName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.1/docs/resources/dms_endpoint#bucket_folder DmsEndpoint#bucket_folder}.
	BucketFolder *string `field:"optional" json:"bucketFolder" yaml:"bucketFolder"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.1/docs/resources/dms_endpoint#bucket_name DmsEndpoint#bucket_name}.
	BucketName *string `field:"optional" json:"bucketName" yaml:"bucketName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.1/docs/resources/dms_endpoint#canned_acl_for_objects DmsEndpoint#canned_acl_for_objects}.
	CannedAclForObjects *string `field:"optional" json:"cannedAclForObjects" yaml:"cannedAclForObjects"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.1/docs/resources/dms_endpoint#cdc_inserts_and_updates DmsEndpoint#cdc_inserts_and_updates}.
	CdcInsertsAndUpdates interface{} `field:"optional" json:"cdcInsertsAndUpdates" yaml:"cdcInsertsAndUpdates"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.1/docs/resources/dms_endpoint#cdc_inserts_only DmsEndpoint#cdc_inserts_only}.
	CdcInsertsOnly interface{} `field:"optional" json:"cdcInsertsOnly" yaml:"cdcInsertsOnly"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.1/docs/resources/dms_endpoint#cdc_max_batch_interval DmsEndpoint#cdc_max_batch_interval}.
	CdcMaxBatchInterval *float64 `field:"optional" json:"cdcMaxBatchInterval" yaml:"cdcMaxBatchInterval"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.1/docs/resources/dms_endpoint#cdc_min_file_size DmsEndpoint#cdc_min_file_size}.
	CdcMinFileSize *float64 `field:"optional" json:"cdcMinFileSize" yaml:"cdcMinFileSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.1/docs/resources/dms_endpoint#cdc_path DmsEndpoint#cdc_path}.
	CdcPath *string `field:"optional" json:"cdcPath" yaml:"cdcPath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.1/docs/resources/dms_endpoint#compression_type DmsEndpoint#compression_type}.
	CompressionType *string `field:"optional" json:"compressionType" yaml:"compressionType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.1/docs/resources/dms_endpoint#csv_delimiter DmsEndpoint#csv_delimiter}.
	CsvDelimiter *string `field:"optional" json:"csvDelimiter" yaml:"csvDelimiter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.1/docs/resources/dms_endpoint#csv_no_sup_value DmsEndpoint#csv_no_sup_value}.
	CsvNoSupValue *string `field:"optional" json:"csvNoSupValue" yaml:"csvNoSupValue"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.1/docs/resources/dms_endpoint#csv_null_value DmsEndpoint#csv_null_value}.
	CsvNullValue *string `field:"optional" json:"csvNullValue" yaml:"csvNullValue"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.1/docs/resources/dms_endpoint#csv_row_delimiter DmsEndpoint#csv_row_delimiter}.
	CsvRowDelimiter *string `field:"optional" json:"csvRowDelimiter" yaml:"csvRowDelimiter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.1/docs/resources/dms_endpoint#data_format DmsEndpoint#data_format}.
	DataFormat *string `field:"optional" json:"dataFormat" yaml:"dataFormat"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.1/docs/resources/dms_endpoint#data_page_size DmsEndpoint#data_page_size}.
	DataPageSize *float64 `field:"optional" json:"dataPageSize" yaml:"dataPageSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.1/docs/resources/dms_endpoint#date_partition_delimiter DmsEndpoint#date_partition_delimiter}.
	DatePartitionDelimiter *string `field:"optional" json:"datePartitionDelimiter" yaml:"datePartitionDelimiter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.1/docs/resources/dms_endpoint#date_partition_enabled DmsEndpoint#date_partition_enabled}.
	DatePartitionEnabled interface{} `field:"optional" json:"datePartitionEnabled" yaml:"datePartitionEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.1/docs/resources/dms_endpoint#date_partition_sequence DmsEndpoint#date_partition_sequence}.
	DatePartitionSequence *string `field:"optional" json:"datePartitionSequence" yaml:"datePartitionSequence"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.1/docs/resources/dms_endpoint#dict_page_size_limit DmsEndpoint#dict_page_size_limit}.
	DictPageSizeLimit *float64 `field:"optional" json:"dictPageSizeLimit" yaml:"dictPageSizeLimit"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.1/docs/resources/dms_endpoint#enable_statistics DmsEndpoint#enable_statistics}.
	EnableStatistics interface{} `field:"optional" json:"enableStatistics" yaml:"enableStatistics"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.1/docs/resources/dms_endpoint#encoding_type DmsEndpoint#encoding_type}.
	EncodingType *string `field:"optional" json:"encodingType" yaml:"encodingType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.1/docs/resources/dms_endpoint#encryption_mode DmsEndpoint#encryption_mode}.
	EncryptionMode *string `field:"optional" json:"encryptionMode" yaml:"encryptionMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.1/docs/resources/dms_endpoint#external_table_definition DmsEndpoint#external_table_definition}.
	ExternalTableDefinition *string `field:"optional" json:"externalTableDefinition" yaml:"externalTableDefinition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.1/docs/resources/dms_endpoint#ignore_header_rows DmsEndpoint#ignore_header_rows}.
	IgnoreHeaderRows *float64 `field:"optional" json:"ignoreHeaderRows" yaml:"ignoreHeaderRows"`
	// This setting has no effect, is deprecated, and will be removed in a future version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.1/docs/resources/dms_endpoint#ignore_headers_row DmsEndpoint#ignore_headers_row}
	IgnoreHeadersRow *float64 `field:"optional" json:"ignoreHeadersRow" yaml:"ignoreHeadersRow"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.1/docs/resources/dms_endpoint#include_op_for_full_load DmsEndpoint#include_op_for_full_load}.
	IncludeOpForFullLoad interface{} `field:"optional" json:"includeOpForFullLoad" yaml:"includeOpForFullLoad"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.1/docs/resources/dms_endpoint#max_file_size DmsEndpoint#max_file_size}.
	MaxFileSize *float64 `field:"optional" json:"maxFileSize" yaml:"maxFileSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.1/docs/resources/dms_endpoint#parquet_timestamp_in_millisecond DmsEndpoint#parquet_timestamp_in_millisecond}.
	ParquetTimestampInMillisecond interface{} `field:"optional" json:"parquetTimestampInMillisecond" yaml:"parquetTimestampInMillisecond"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.1/docs/resources/dms_endpoint#parquet_version DmsEndpoint#parquet_version}.
	ParquetVersion *string `field:"optional" json:"parquetVersion" yaml:"parquetVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.1/docs/resources/dms_endpoint#preserve_transactions DmsEndpoint#preserve_transactions}.
	PreserveTransactions interface{} `field:"optional" json:"preserveTransactions" yaml:"preserveTransactions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.1/docs/resources/dms_endpoint#rfc_4180 DmsEndpoint#rfc_4180}.
	Rfc4180 interface{} `field:"optional" json:"rfc4180" yaml:"rfc4180"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.1/docs/resources/dms_endpoint#row_group_length DmsEndpoint#row_group_length}.
	RowGroupLength *float64 `field:"optional" json:"rowGroupLength" yaml:"rowGroupLength"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.1/docs/resources/dms_endpoint#server_side_encryption_kms_key_id DmsEndpoint#server_side_encryption_kms_key_id}.
	ServerSideEncryptionKmsKeyId *string `field:"optional" json:"serverSideEncryptionKmsKeyId" yaml:"serverSideEncryptionKmsKeyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.1/docs/resources/dms_endpoint#service_access_role_arn DmsEndpoint#service_access_role_arn}.
	ServiceAccessRoleArn *string `field:"optional" json:"serviceAccessRoleArn" yaml:"serviceAccessRoleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.1/docs/resources/dms_endpoint#timestamp_column_name DmsEndpoint#timestamp_column_name}.
	TimestampColumnName *string `field:"optional" json:"timestampColumnName" yaml:"timestampColumnName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.1/docs/resources/dms_endpoint#use_csv_no_sup_value DmsEndpoint#use_csv_no_sup_value}.
	UseCsvNoSupValue interface{} `field:"optional" json:"useCsvNoSupValue" yaml:"useCsvNoSupValue"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.1/docs/resources/dms_endpoint#use_task_start_time_for_full_load_timestamp DmsEndpoint#use_task_start_time_for_full_load_timestamp}.
	UseTaskStartTimeForFullLoadTimestamp interface{} `field:"optional" json:"useTaskStartTimeForFullLoadTimestamp" yaml:"useTaskStartTimeForFullLoadTimestamp"`
}

