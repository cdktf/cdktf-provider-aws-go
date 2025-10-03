// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lambdafunction


type LambdaFunctionEphemeralStorage struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/lambda_function#size LambdaFunction#size}.
	Size *float64 `field:"optional" json:"size" yaml:"size"`
}

