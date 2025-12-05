// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lambdafunction


type LambdaFunctionDurableConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/lambda_function#execution_timeout LambdaFunction#execution_timeout}.
	ExecutionTimeout *float64 `field:"required" json:"executionTimeout" yaml:"executionTimeout"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/lambda_function#retention_period LambdaFunction#retention_period}.
	RetentionPeriod *float64 `field:"optional" json:"retentionPeriod" yaml:"retentionPeriod"`
}

