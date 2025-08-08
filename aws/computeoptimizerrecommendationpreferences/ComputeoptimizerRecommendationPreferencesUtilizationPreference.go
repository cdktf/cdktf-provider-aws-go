// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computeoptimizerrecommendationpreferences


type ComputeoptimizerRecommendationPreferencesUtilizationPreference struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/computeoptimizer_recommendation_preferences#metric_name ComputeoptimizerRecommendationPreferences#metric_name}.
	MetricName *string `field:"required" json:"metricName" yaml:"metricName"`
	// metric_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/computeoptimizer_recommendation_preferences#metric_parameters ComputeoptimizerRecommendationPreferences#metric_parameters}
	MetricParameters interface{} `field:"optional" json:"metricParameters" yaml:"metricParameters"`
}

