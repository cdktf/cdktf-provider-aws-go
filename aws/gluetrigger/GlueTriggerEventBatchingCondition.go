// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gluetrigger


type GlueTriggerEventBatchingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/glue_trigger#batch_size GlueTrigger#batch_size}.
	BatchSize *float64 `field:"required" json:"batchSize" yaml:"batchSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/glue_trigger#batch_window GlueTrigger#batch_window}.
	BatchWindow *float64 `field:"optional" json:"batchWindow" yaml:"batchWindow"`
}

