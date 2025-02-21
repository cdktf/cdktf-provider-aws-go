// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ceanomalysubscription


type CeAnomalySubscriptionThresholdExpression struct {
	// and block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.88.0/docs/resources/ce_anomaly_subscription#and CeAnomalySubscription#and}
	And interface{} `field:"optional" json:"and" yaml:"and"`
	// cost_category block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.88.0/docs/resources/ce_anomaly_subscription#cost_category CeAnomalySubscription#cost_category}
	CostCategory *CeAnomalySubscriptionThresholdExpressionCostCategory `field:"optional" json:"costCategory" yaml:"costCategory"`
	// dimension block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.88.0/docs/resources/ce_anomaly_subscription#dimension CeAnomalySubscription#dimension}
	Dimension *CeAnomalySubscriptionThresholdExpressionDimension `field:"optional" json:"dimension" yaml:"dimension"`
	// not block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.88.0/docs/resources/ce_anomaly_subscription#not CeAnomalySubscription#not}
	Not *CeAnomalySubscriptionThresholdExpressionNot `field:"optional" json:"not" yaml:"not"`
	// or block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.88.0/docs/resources/ce_anomaly_subscription#or CeAnomalySubscription#or}
	Or interface{} `field:"optional" json:"or" yaml:"or"`
	// tags block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.88.0/docs/resources/ce_anomaly_subscription#tags CeAnomalySubscription#tags}
	Tags *CeAnomalySubscriptionThresholdExpressionTags `field:"optional" json:"tags" yaml:"tags"`
}

