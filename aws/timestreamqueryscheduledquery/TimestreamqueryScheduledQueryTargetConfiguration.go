// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package timestreamqueryscheduledquery


type TimestreamqueryScheduledQueryTargetConfiguration struct {
	// timestream_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/timestreamquery_scheduled_query#timestream_configuration TimestreamqueryScheduledQuery#timestream_configuration}
	TimestreamConfiguration interface{} `field:"optional" json:"timestreamConfiguration" yaml:"timestreamConfiguration"`
}

