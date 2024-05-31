// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lambdafunctioneventinvokeconfig


type LambdaFunctionEventInvokeConfigDestinationConfig struct {
	// on_failure block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.52.0/docs/resources/lambda_function_event_invoke_config#on_failure LambdaFunctionEventInvokeConfig#on_failure}
	OnFailure *LambdaFunctionEventInvokeConfigDestinationConfigOnFailure `field:"optional" json:"onFailure" yaml:"onFailure"`
	// on_success block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.52.0/docs/resources/lambda_function_event_invoke_config#on_success LambdaFunctionEventInvokeConfig#on_success}
	OnSuccess *LambdaFunctionEventInvokeConfigDestinationConfigOnSuccess `field:"optional" json:"onSuccess" yaml:"onSuccess"`
}

