// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appfabricingestiondestination


type AppfabricIngestionDestinationDestinationConfigurationAuditLogDestination struct {
	// firehose_stream block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.4.0/docs/resources/appfabric_ingestion_destination#firehose_stream AppfabricIngestionDestination#firehose_stream}
	FirehoseStream interface{} `field:"optional" json:"firehoseStream" yaml:"firehoseStream"`
	// s3_bucket block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.4.0/docs/resources/appfabric_ingestion_destination#s3_bucket AppfabricIngestionDestination#s3_bucket}
	S3Bucket interface{} `field:"optional" json:"s3Bucket" yaml:"s3Bucket"`
}

