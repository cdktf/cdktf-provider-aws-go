// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dynamodbtable


type DynamodbTablePointInTimeRecovery struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.0/docs/resources/dynamodb_table#enabled DynamodbTable#enabled}.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.0/docs/resources/dynamodb_table#recovery_period_in_days DynamodbTable#recovery_period_in_days}.
	RecoveryPeriodInDays *float64 `field:"optional" json:"recoveryPeriodInDays" yaml:"recoveryPeriodInDays"`
}

