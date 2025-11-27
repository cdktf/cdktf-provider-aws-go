// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package elb


type ElbAccessLogs struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/elb#bucket Elb#bucket}.
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/elb#bucket_prefix Elb#bucket_prefix}.
	BucketPrefix *string `field:"optional" json:"bucketPrefix" yaml:"bucketPrefix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/elb#enabled Elb#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/elb#interval Elb#interval}.
	Interval *float64 `field:"optional" json:"interval" yaml:"interval"`
}

