// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package globalacceleratorcustomroutingaccelerator


type GlobalacceleratorCustomRoutingAcceleratorAttributes struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.29.0/docs/resources/globalaccelerator_custom_routing_accelerator#flow_logs_enabled GlobalacceleratorCustomRoutingAccelerator#flow_logs_enabled}.
	FlowLogsEnabled interface{} `field:"optional" json:"flowLogsEnabled" yaml:"flowLogsEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.29.0/docs/resources/globalaccelerator_custom_routing_accelerator#flow_logs_s3_bucket GlobalacceleratorCustomRoutingAccelerator#flow_logs_s3_bucket}.
	FlowLogsS3Bucket *string `field:"optional" json:"flowLogsS3Bucket" yaml:"flowLogsS3Bucket"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.29.0/docs/resources/globalaccelerator_custom_routing_accelerator#flow_logs_s3_prefix GlobalacceleratorCustomRoutingAccelerator#flow_logs_s3_prefix}.
	FlowLogsS3Prefix *string `field:"optional" json:"flowLogsS3Prefix" yaml:"flowLogsS3Prefix"`
}

