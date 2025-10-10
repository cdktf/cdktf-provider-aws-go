// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package globalacceleratoraccelerator


type GlobalacceleratorAcceleratorAttributes struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/globalaccelerator_accelerator#flow_logs_enabled GlobalacceleratorAccelerator#flow_logs_enabled}.
	FlowLogsEnabled interface{} `field:"optional" json:"flowLogsEnabled" yaml:"flowLogsEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/globalaccelerator_accelerator#flow_logs_s3_bucket GlobalacceleratorAccelerator#flow_logs_s3_bucket}.
	FlowLogsS3Bucket *string `field:"optional" json:"flowLogsS3Bucket" yaml:"flowLogsS3Bucket"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/globalaccelerator_accelerator#flow_logs_s3_prefix GlobalacceleratorAccelerator#flow_logs_s3_prefix}.
	FlowLogsS3Prefix *string `field:"optional" json:"flowLogsS3Prefix" yaml:"flowLogsS3Prefix"`
}

