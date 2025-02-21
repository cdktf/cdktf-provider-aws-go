// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudwatchmetricstream


type CloudwatchMetricStreamStatisticsConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.88.0/docs/resources/cloudwatch_metric_stream#additional_statistics CloudwatchMetricStream#additional_statistics}.
	AdditionalStatistics *[]*string `field:"required" json:"additionalStatistics" yaml:"additionalStatistics"`
	// include_metric block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.88.0/docs/resources/cloudwatch_metric_stream#include_metric CloudwatchMetricStream#include_metric}
	IncludeMetric interface{} `field:"required" json:"includeMetric" yaml:"includeMetric"`
}

