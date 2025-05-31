// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package rdscluster


type RdsClusterScalingConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/rds_cluster#auto_pause RdsCluster#auto_pause}.
	AutoPause interface{} `field:"optional" json:"autoPause" yaml:"autoPause"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/rds_cluster#max_capacity RdsCluster#max_capacity}.
	MaxCapacity *float64 `field:"optional" json:"maxCapacity" yaml:"maxCapacity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/rds_cluster#min_capacity RdsCluster#min_capacity}.
	MinCapacity *float64 `field:"optional" json:"minCapacity" yaml:"minCapacity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/rds_cluster#seconds_before_timeout RdsCluster#seconds_before_timeout}.
	SecondsBeforeTimeout *float64 `field:"optional" json:"secondsBeforeTimeout" yaml:"secondsBeforeTimeout"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/rds_cluster#seconds_until_auto_pause RdsCluster#seconds_until_auto_pause}.
	SecondsUntilAutoPause *float64 `field:"optional" json:"secondsUntilAutoPause" yaml:"secondsUntilAutoPause"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/rds_cluster#timeout_action RdsCluster#timeout_action}.
	TimeoutAction *string `field:"optional" json:"timeoutAction" yaml:"timeoutAction"`
}

