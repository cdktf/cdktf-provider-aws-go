// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computeoptimizerrecommendationpreferences


type ComputeoptimizerRecommendationPreferencesPreferredResource struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/computeoptimizer_recommendation_preferences#name ComputeoptimizerRecommendationPreferences#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/computeoptimizer_recommendation_preferences#exclude_list ComputeoptimizerRecommendationPreferences#exclude_list}.
	ExcludeList *[]*string `field:"optional" json:"excludeList" yaml:"excludeList"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/computeoptimizer_recommendation_preferences#include_list ComputeoptimizerRecommendationPreferences#include_list}.
	IncludeList *[]*string `field:"optional" json:"includeList" yaml:"includeList"`
}

