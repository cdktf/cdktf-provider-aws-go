// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package timestreaminfluxdbdbinstance


type TimestreaminfluxdbDbInstanceLogDeliveryConfigurationS3Configuration struct {
	// The name of the S3 bucket to deliver logs to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.0/docs/resources/timestreaminfluxdb_db_instance#bucket_name TimestreaminfluxdbDbInstance#bucket_name}
	BucketName *string `field:"required" json:"bucketName" yaml:"bucketName"`
	// Indicates whether log delivery to the S3 bucket is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.0/docs/resources/timestreaminfluxdb_db_instance#enabled TimestreaminfluxdbDbInstance#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
}

