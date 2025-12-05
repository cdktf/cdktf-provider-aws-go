// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package timestreamqueryscheduledquery


type TimestreamqueryScheduledQueryRecentlyFailedRunsQueryInsightsResponse struct {
	// query_spatial_coverage block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/timestreamquery_scheduled_query#query_spatial_coverage TimestreamqueryScheduledQuery#query_spatial_coverage}
	QuerySpatialCoverage interface{} `field:"optional" json:"querySpatialCoverage" yaml:"querySpatialCoverage"`
	// query_temporal_range block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/timestreamquery_scheduled_query#query_temporal_range TimestreamqueryScheduledQuery#query_temporal_range}
	QueryTemporalRange interface{} `field:"optional" json:"queryTemporalRange" yaml:"queryTemporalRange"`
}

