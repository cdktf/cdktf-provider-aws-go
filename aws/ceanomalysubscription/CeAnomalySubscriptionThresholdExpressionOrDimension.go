// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ceanomalysubscription


type CeAnomalySubscriptionThresholdExpressionOrDimension struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.81.0/docs/resources/ce_anomaly_subscription#key CeAnomalySubscription#key}.
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.81.0/docs/resources/ce_anomaly_subscription#match_options CeAnomalySubscription#match_options}.
	MatchOptions *[]*string `field:"optional" json:"matchOptions" yaml:"matchOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.81.0/docs/resources/ce_anomaly_subscription#values CeAnomalySubscription#values}.
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

