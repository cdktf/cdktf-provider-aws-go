// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lambdacapacityprovider


type LambdaCapacityProviderCapacityProviderScalingConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/lambda_capacity_provider#max_vcpu_count LambdaCapacityProvider#max_vcpu_count}.
	MaxVcpuCount *float64 `field:"optional" json:"maxVcpuCount" yaml:"maxVcpuCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/lambda_capacity_provider#scaling_mode LambdaCapacityProvider#scaling_mode}.
	ScalingMode *string `field:"optional" json:"scalingMode" yaml:"scalingMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/lambda_capacity_provider#scaling_policies LambdaCapacityProvider#scaling_policies}.
	ScalingPolicies interface{} `field:"optional" json:"scalingPolicies" yaml:"scalingPolicies"`
}

