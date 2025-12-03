// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package timestreaminfluxdbdbcluster


type TimestreaminfluxdbDbClusterLogDeliveryConfigurationS3Configuration struct {
	// The name of the S3 bucket to deliver logs to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.24.0/docs/resources/timestreaminfluxdb_db_cluster#bucket_name TimestreaminfluxdbDbCluster#bucket_name}
	BucketName *string `field:"required" json:"bucketName" yaml:"bucketName"`
	// Indicates whether log delivery to the S3 bucket is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.24.0/docs/resources/timestreaminfluxdb_db_cluster#enabled TimestreaminfluxdbDbCluster#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
}

