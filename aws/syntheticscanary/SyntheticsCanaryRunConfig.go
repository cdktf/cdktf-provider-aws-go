// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package syntheticscanary


type SyntheticsCanaryRunConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/synthetics_canary#active_tracing SyntheticsCanary#active_tracing}.
	ActiveTracing interface{} `field:"optional" json:"activeTracing" yaml:"activeTracing"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/synthetics_canary#environment_variables SyntheticsCanary#environment_variables}.
	EnvironmentVariables *map[string]*string `field:"optional" json:"environmentVariables" yaml:"environmentVariables"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/synthetics_canary#ephemeral_storage SyntheticsCanary#ephemeral_storage}.
	EphemeralStorage *float64 `field:"optional" json:"ephemeralStorage" yaml:"ephemeralStorage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/synthetics_canary#memory_in_mb SyntheticsCanary#memory_in_mb}.
	MemoryInMb *float64 `field:"optional" json:"memoryInMb" yaml:"memoryInMb"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/synthetics_canary#timeout_in_seconds SyntheticsCanary#timeout_in_seconds}.
	TimeoutInSeconds *float64 `field:"optional" json:"timeoutInSeconds" yaml:"timeoutInSeconds"`
}

