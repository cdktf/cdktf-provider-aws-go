// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package timestreamqueryscheduledquery


type TimestreamqueryScheduledQueryTargetConfigurationTimestreamConfigurationDimensionMapping struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.21.0/docs/resources/timestreamquery_scheduled_query#dimension_value_type TimestreamqueryScheduledQuery#dimension_value_type}.
	DimensionValueType *string `field:"required" json:"dimensionValueType" yaml:"dimensionValueType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.21.0/docs/resources/timestreamquery_scheduled_query#name TimestreamqueryScheduledQuery#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
}

