// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ceanomalysubscription


type CeAnomalySubscriptionThresholdExpressionOr struct {
	// cost_category block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.1/docs/resources/ce_anomaly_subscription#cost_category CeAnomalySubscription#cost_category}
	CostCategory *CeAnomalySubscriptionThresholdExpressionOrCostCategory `field:"optional" json:"costCategory" yaml:"costCategory"`
	// dimension block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.1/docs/resources/ce_anomaly_subscription#dimension CeAnomalySubscription#dimension}
	Dimension *CeAnomalySubscriptionThresholdExpressionOrDimension `field:"optional" json:"dimension" yaml:"dimension"`
	// tags block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.1/docs/resources/ce_anomaly_subscription#tags CeAnomalySubscription#tags}
	Tags *CeAnomalySubscriptionThresholdExpressionOrTags `field:"optional" json:"tags" yaml:"tags"`
}

