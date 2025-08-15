// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package timestreamqueryscheduledquery


type TimestreamqueryScheduledQueryNotificationConfiguration struct {
	// sns_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.9.0/docs/resources/timestreamquery_scheduled_query#sns_configuration TimestreamqueryScheduledQuery#sns_configuration}
	SnsConfiguration interface{} `field:"optional" json:"snsConfiguration" yaml:"snsConfiguration"`
}

