// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ceanomalysubscription


type CeAnomalySubscriptionThresholdExpressionAnd struct {
	// cost_category block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.21.0/docs/resources/ce_anomaly_subscription#cost_category CeAnomalySubscription#cost_category}
	CostCategory *CeAnomalySubscriptionThresholdExpressionAndCostCategory `field:"optional" json:"costCategory" yaml:"costCategory"`
	// dimension block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.21.0/docs/resources/ce_anomaly_subscription#dimension CeAnomalySubscription#dimension}
	Dimension *CeAnomalySubscriptionThresholdExpressionAndDimension `field:"optional" json:"dimension" yaml:"dimension"`
	// tags block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.21.0/docs/resources/ce_anomaly_subscription#tags CeAnomalySubscription#tags}
	Tags *CeAnomalySubscriptionThresholdExpressionAndTags `field:"optional" json:"tags" yaml:"tags"`
}

