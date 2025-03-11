// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lambdafunctioneventinvokeconfig

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LambdaFunctionEventInvokeConfigConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.1/docs/resources/lambda_function_event_invoke_config#function_name LambdaFunctionEventInvokeConfig#function_name}.
	FunctionName *string `field:"required" json:"functionName" yaml:"functionName"`
	// destination_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.1/docs/resources/lambda_function_event_invoke_config#destination_config LambdaFunctionEventInvokeConfig#destination_config}
	DestinationConfig *LambdaFunctionEventInvokeConfigDestinationConfig `field:"optional" json:"destinationConfig" yaml:"destinationConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.1/docs/resources/lambda_function_event_invoke_config#id LambdaFunctionEventInvokeConfig#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.1/docs/resources/lambda_function_event_invoke_config#maximum_event_age_in_seconds LambdaFunctionEventInvokeConfig#maximum_event_age_in_seconds}.
	MaximumEventAgeInSeconds *float64 `field:"optional" json:"maximumEventAgeInSeconds" yaml:"maximumEventAgeInSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.1/docs/resources/lambda_function_event_invoke_config#maximum_retry_attempts LambdaFunctionEventInvokeConfig#maximum_retry_attempts}.
	MaximumRetryAttempts *float64 `field:"optional" json:"maximumRetryAttempts" yaml:"maximumRetryAttempts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.1/docs/resources/lambda_function_event_invoke_config#qualifier LambdaFunctionEventInvokeConfig#qualifier}.
	Qualifier *string `field:"optional" json:"qualifier" yaml:"qualifier"`
}

