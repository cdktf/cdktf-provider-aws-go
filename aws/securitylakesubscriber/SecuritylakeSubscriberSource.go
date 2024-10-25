// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package securitylakesubscriber


type SecuritylakeSubscriberSource struct {
	// aws_log_source_resource block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.73.0/docs/resources/securitylake_subscriber#aws_log_source_resource SecuritylakeSubscriber#aws_log_source_resource}
	AwsLogSourceResource interface{} `field:"optional" json:"awsLogSourceResource" yaml:"awsLogSourceResource"`
	// custom_log_source_resource block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.73.0/docs/resources/securitylake_subscriber#custom_log_source_resource SecuritylakeSubscriber#custom_log_source_resource}
	CustomLogSourceResource interface{} `field:"optional" json:"customLogSourceResource" yaml:"customLogSourceResource"`
}

