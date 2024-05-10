// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ceanomalysubscription


type CeAnomalySubscriptionSubscriber struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.49.0/docs/resources/ce_anomaly_subscription#address CeAnomalySubscription#address}.
	Address *string `field:"required" json:"address" yaml:"address"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.49.0/docs/resources/ce_anomaly_subscription#type CeAnomalySubscription#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
}

