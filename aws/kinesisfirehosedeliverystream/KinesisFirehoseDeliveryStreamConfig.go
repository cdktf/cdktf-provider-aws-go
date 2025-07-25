// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kinesisfirehosedeliverystream

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type KinesisFirehoseDeliveryStreamConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.5.0/docs/resources/kinesis_firehose_delivery_stream#destination KinesisFirehoseDeliveryStream#destination}.
	Destination *string `field:"required" json:"destination" yaml:"destination"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.5.0/docs/resources/kinesis_firehose_delivery_stream#name KinesisFirehoseDeliveryStream#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.5.0/docs/resources/kinesis_firehose_delivery_stream#arn KinesisFirehoseDeliveryStream#arn}.
	Arn *string `field:"optional" json:"arn" yaml:"arn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.5.0/docs/resources/kinesis_firehose_delivery_stream#destination_id KinesisFirehoseDeliveryStream#destination_id}.
	DestinationId *string `field:"optional" json:"destinationId" yaml:"destinationId"`
	// elasticsearch_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.5.0/docs/resources/kinesis_firehose_delivery_stream#elasticsearch_configuration KinesisFirehoseDeliveryStream#elasticsearch_configuration}
	ElasticsearchConfiguration *KinesisFirehoseDeliveryStreamElasticsearchConfiguration `field:"optional" json:"elasticsearchConfiguration" yaml:"elasticsearchConfiguration"`
	// extended_s3_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.5.0/docs/resources/kinesis_firehose_delivery_stream#extended_s3_configuration KinesisFirehoseDeliveryStream#extended_s3_configuration}
	ExtendedS3Configuration *KinesisFirehoseDeliveryStreamExtendedS3Configuration `field:"optional" json:"extendedS3Configuration" yaml:"extendedS3Configuration"`
	// http_endpoint_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.5.0/docs/resources/kinesis_firehose_delivery_stream#http_endpoint_configuration KinesisFirehoseDeliveryStream#http_endpoint_configuration}
	HttpEndpointConfiguration *KinesisFirehoseDeliveryStreamHttpEndpointConfiguration `field:"optional" json:"httpEndpointConfiguration" yaml:"httpEndpointConfiguration"`
	// iceberg_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.5.0/docs/resources/kinesis_firehose_delivery_stream#iceberg_configuration KinesisFirehoseDeliveryStream#iceberg_configuration}
	IcebergConfiguration *KinesisFirehoseDeliveryStreamIcebergConfiguration `field:"optional" json:"icebergConfiguration" yaml:"icebergConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.5.0/docs/resources/kinesis_firehose_delivery_stream#id KinesisFirehoseDeliveryStream#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// kinesis_source_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.5.0/docs/resources/kinesis_firehose_delivery_stream#kinesis_source_configuration KinesisFirehoseDeliveryStream#kinesis_source_configuration}
	KinesisSourceConfiguration *KinesisFirehoseDeliveryStreamKinesisSourceConfiguration `field:"optional" json:"kinesisSourceConfiguration" yaml:"kinesisSourceConfiguration"`
	// msk_source_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.5.0/docs/resources/kinesis_firehose_delivery_stream#msk_source_configuration KinesisFirehoseDeliveryStream#msk_source_configuration}
	MskSourceConfiguration *KinesisFirehoseDeliveryStreamMskSourceConfiguration `field:"optional" json:"mskSourceConfiguration" yaml:"mskSourceConfiguration"`
	// opensearch_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.5.0/docs/resources/kinesis_firehose_delivery_stream#opensearch_configuration KinesisFirehoseDeliveryStream#opensearch_configuration}
	OpensearchConfiguration *KinesisFirehoseDeliveryStreamOpensearchConfiguration `field:"optional" json:"opensearchConfiguration" yaml:"opensearchConfiguration"`
	// opensearchserverless_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.5.0/docs/resources/kinesis_firehose_delivery_stream#opensearchserverless_configuration KinesisFirehoseDeliveryStream#opensearchserverless_configuration}
	OpensearchserverlessConfiguration *KinesisFirehoseDeliveryStreamOpensearchserverlessConfiguration `field:"optional" json:"opensearchserverlessConfiguration" yaml:"opensearchserverlessConfiguration"`
	// redshift_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.5.0/docs/resources/kinesis_firehose_delivery_stream#redshift_configuration KinesisFirehoseDeliveryStream#redshift_configuration}
	RedshiftConfiguration *KinesisFirehoseDeliveryStreamRedshiftConfiguration `field:"optional" json:"redshiftConfiguration" yaml:"redshiftConfiguration"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.5.0/docs/resources/kinesis_firehose_delivery_stream#region KinesisFirehoseDeliveryStream#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// server_side_encryption block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.5.0/docs/resources/kinesis_firehose_delivery_stream#server_side_encryption KinesisFirehoseDeliveryStream#server_side_encryption}
	ServerSideEncryption *KinesisFirehoseDeliveryStreamServerSideEncryption `field:"optional" json:"serverSideEncryption" yaml:"serverSideEncryption"`
	// snowflake_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.5.0/docs/resources/kinesis_firehose_delivery_stream#snowflake_configuration KinesisFirehoseDeliveryStream#snowflake_configuration}
	SnowflakeConfiguration *KinesisFirehoseDeliveryStreamSnowflakeConfiguration `field:"optional" json:"snowflakeConfiguration" yaml:"snowflakeConfiguration"`
	// splunk_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.5.0/docs/resources/kinesis_firehose_delivery_stream#splunk_configuration KinesisFirehoseDeliveryStream#splunk_configuration}
	SplunkConfiguration *KinesisFirehoseDeliveryStreamSplunkConfiguration `field:"optional" json:"splunkConfiguration" yaml:"splunkConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.5.0/docs/resources/kinesis_firehose_delivery_stream#tags KinesisFirehoseDeliveryStream#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.5.0/docs/resources/kinesis_firehose_delivery_stream#tags_all KinesisFirehoseDeliveryStream#tags_all}.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.5.0/docs/resources/kinesis_firehose_delivery_stream#timeouts KinesisFirehoseDeliveryStream#timeouts}
	Timeouts *KinesisFirehoseDeliveryStreamTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.5.0/docs/resources/kinesis_firehose_delivery_stream#version_id KinesisFirehoseDeliveryStream#version_id}.
	VersionId *string `field:"optional" json:"versionId" yaml:"versionId"`
}

