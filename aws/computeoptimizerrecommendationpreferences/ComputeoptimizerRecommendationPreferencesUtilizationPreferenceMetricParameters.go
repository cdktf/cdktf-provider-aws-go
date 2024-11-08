// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computeoptimizerrecommendationpreferences


type ComputeoptimizerRecommendationPreferencesUtilizationPreferenceMetricParameters struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.0/docs/resources/computeoptimizer_recommendation_preferences#headroom ComputeoptimizerRecommendationPreferences#headroom}.
	Headroom *string `field:"required" json:"headroom" yaml:"headroom"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.0/docs/resources/computeoptimizer_recommendation_preferences#threshold ComputeoptimizerRecommendationPreferences#threshold}.
	Threshold *string `field:"optional" json:"threshold" yaml:"threshold"`
}

