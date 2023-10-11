// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lambdafunction


type LambdaFunctionDeadLetterConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.20.1/docs/resources/lambda_function#target_arn LambdaFunction#target_arn}.
	TargetArn *string `field:"required" json:"targetArn" yaml:"targetArn"`
}

