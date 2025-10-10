// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package timestreamqueryscheduledquery


type TimestreamqueryScheduledQueryRecentlyFailedRuns struct {
	// error_report_location block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/timestreamquery_scheduled_query#error_report_location TimestreamqueryScheduledQuery#error_report_location}
	ErrorReportLocation interface{} `field:"optional" json:"errorReportLocation" yaml:"errorReportLocation"`
	// execution_stats block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/timestreamquery_scheduled_query#execution_stats TimestreamqueryScheduledQuery#execution_stats}
	ExecutionStats interface{} `field:"optional" json:"executionStats" yaml:"executionStats"`
	// query_insights_response block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/timestreamquery_scheduled_query#query_insights_response TimestreamqueryScheduledQuery#query_insights_response}
	QueryInsightsResponse interface{} `field:"optional" json:"queryInsightsResponse" yaml:"queryInsightsResponse"`
}

