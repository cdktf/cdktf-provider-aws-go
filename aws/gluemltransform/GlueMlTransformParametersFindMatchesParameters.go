// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gluemltransform


type GlueMlTransformParametersFindMatchesParameters struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.1/docs/resources/glue_ml_transform#accuracy_cost_trade_off GlueMlTransform#accuracy_cost_trade_off}.
	AccuracyCostTradeOff *float64 `field:"optional" json:"accuracyCostTradeOff" yaml:"accuracyCostTradeOff"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.1/docs/resources/glue_ml_transform#enforce_provided_labels GlueMlTransform#enforce_provided_labels}.
	EnforceProvidedLabels interface{} `field:"optional" json:"enforceProvidedLabels" yaml:"enforceProvidedLabels"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.1/docs/resources/glue_ml_transform#precision_recall_trade_off GlueMlTransform#precision_recall_trade_off}.
	PrecisionRecallTradeOff *float64 `field:"optional" json:"precisionRecallTradeOff" yaml:"precisionRecallTradeOff"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.1/docs/resources/glue_ml_transform#primary_key_column_name GlueMlTransform#primary_key_column_name}.
	PrimaryKeyColumnName *string `field:"optional" json:"primaryKeyColumnName" yaml:"primaryKeyColumnName"`
}

