// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package syntheticscanary


type SyntheticsCanarySchedule struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.0/docs/resources/synthetics_canary#expression SyntheticsCanary#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.0/docs/resources/synthetics_canary#duration_in_seconds SyntheticsCanary#duration_in_seconds}.
	DurationInSeconds *float64 `field:"optional" json:"durationInSeconds" yaml:"durationInSeconds"`
	// retry_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.0/docs/resources/synthetics_canary#retry_config SyntheticsCanary#retry_config}
	RetryConfig *SyntheticsCanaryScheduleRetryConfig `field:"optional" json:"retryConfig" yaml:"retryConfig"`
}

